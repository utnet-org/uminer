// Copyright 2014 The go-utility Authors
// This file is part of the go-utility library.
//
// The go-utility library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-utility library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-utility library. If not, see <http://www.gnu.org/licenses/>.

// Package unc implements the Utility protocol.
package unc

import (
	"errors"
	"fmt"
	"math/big"
	"runtime"
	"sync"

	"github.com/yanhuangpai/go-utility/accounts"
	"github.com/yanhuangpai/go-utility/common"
	"github.com/yanhuangpai/go-utility/common/hexutil"
	"github.com/yanhuangpai/go-utility/consensus"
	"github.com/yanhuangpai/go-utility/consensus/beacon"
	"github.com/yanhuangpai/go-utility/consensus/clique"
	"github.com/yanhuangpai/go-utility/core"
	"github.com/yanhuangpai/go-utility/core/bloombits"
	"github.com/yanhuangpai/go-utility/core/rawdb"
	"github.com/yanhuangpai/go-utility/core/state/pruner"
	"github.com/yanhuangpai/go-utility/core/txpool"
	"github.com/yanhuangpai/go-utility/core/txpool/blobpool"
	"github.com/yanhuangpai/go-utility/core/txpool/legacypool"
	"github.com/yanhuangpai/go-utility/core/types"
	"github.com/yanhuangpai/go-utility/core/vm"
	"github.com/yanhuangpai/go-utility/event"
	"github.com/yanhuangpai/go-utility/internal/shutdowncheck"
	"github.com/yanhuangpai/go-utility/internal/uncapi"
	"github.com/yanhuangpai/go-utility/log"
	"github.com/yanhuangpai/go-utility/miner"
	"github.com/yanhuangpai/go-utility/node"
	"github.com/yanhuangpai/go-utility/p2p"
	"github.com/yanhuangpai/go-utility/p2p/dnsdisc"
	"github.com/yanhuangpai/go-utility/p2p/enode"
	"github.com/yanhuangpai/go-utility/params"
	"github.com/yanhuangpai/go-utility/rlp"
	"github.com/yanhuangpai/go-utility/rpc"
	"github.com/yanhuangpai/go-utility/unc/downloader"
	"github.com/yanhuangpai/go-utility/unc/gasprice"
	"github.com/yanhuangpai/go-utility/unc/protocols/snap"
	"github.com/yanhuangpai/go-utility/unc/protocols/unc"
	"github.com/yanhuangpai/go-utility/unc/uncconfig"
	"github.com/yanhuangpai/go-utility/uncdb"
)

// Config contains the configuration options of the UNC protocol.
// Deprecated: use uncconfig.Config instead.
type Config = uncconfig.Config

// Utility implements the Utility full node service.
type Utility struct {
	config *uncconfig.Config

	// Handlers
	txPool *txpool.TxPool

	blockchain         *core.BlockChain
	handler            *handler
	ethDialCandidates  enode.Iterator
	snapDialCandidates enode.Iterator
	merger             *consensus.Merger

	// DB interfaces
	chainDb uncdb.Database // Block chain database

	eventMux       *event.TypeMux
	engine         consensus.Engine
	accountManager *accounts.Manager

	bloomRequests     chan chan *bloombits.Retrieval // Channel receiving bloom data retrieval requests
	bloomIndexer      *core.ChainIndexer             // Bloom indexer operating during block imports
	closeBloomHandler chan struct{}

	APIBackend *UncAPIBackend

	miner        *miner.Miner
	gasPrice     *big.Int
	unicrpytbase common.Address

	networkID     uint64
	netRPCService *uncapi.NetAPI

	p2pServer *p2p.Server

	lock sync.RWMutex // Protects the variadic fields (e.g. gas price and unicrpytbase)

	shutdownTracker *shutdowncheck.ShutdownTracker // Tracks if and when the node has shutdown ungracefully
}

// New creates a new Utility object (including the
// initialisation of the common Utility object)
func New(stack *node.Node, config *uncconfig.Config) (*Utility, error) {
	// Ensure configuration values are compatible and sane
	if config.SyncMode == downloader.LightSync {
		return nil, errors.New("can't run unc.Utility in light sync mode, use les.LightUnility")
	}
	if !config.SyncMode.IsValid() {
		return nil, fmt.Errorf("invalid sync mode %d", config.SyncMode)
	}
	if config.Miner.GasPrice == nil || config.Miner.GasPrice.Cmp(common.Big0) <= 0 {
		log.Warn("Sanitizing invalid miner gas price", "provided", config.Miner.GasPrice, "updated", uncconfig.Defaults.Miner.GasPrice)
		config.Miner.GasPrice = new(big.Int).Set(uncconfig.Defaults.Miner.GasPrice)
	}
	if config.NoPruning && config.TrieDirtyCache > 0 {
		if config.SnapshotCache > 0 {
			config.TrieCleanCache += config.TrieDirtyCache * 3 / 5
			config.SnapshotCache += config.TrieDirtyCache * 2 / 5
		} else {
			config.TrieCleanCache += config.TrieDirtyCache
		}
		config.TrieDirtyCache = 0
	}
	log.Info("Allocated trie memory caches", "clean", common.StorageSize(config.TrieCleanCache)*1024*1024, "dirty", common.StorageSize(config.TrieDirtyCache)*1024*1024)

	// Assemble the Utility object
	chainDb, err := stack.OpenDatabaseWithFreezer("chaindata", config.DatabaseCache, config.DatabaseHandles, config.DatabaseFreezer, "unc/db/chaindata/", false)
	if err != nil {
		return nil, err
	}
	scheme, err := rawdb.ParseStateScheme(config.StateScheme, chainDb)
	if err != nil {
		return nil, err
	}
	// Try to recover offline state pruning only in hash-based.
	if scheme == rawdb.HashScheme {
		if err := pruner.RecoverPruning(stack.ResolvePath(""), chainDb); err != nil {
			log.Error("Failed to recover state", "error", err)
		}
	}
	// Transfer mining-related config to the ethash config.
	chainConfig, err := core.LoadChainConfig(chainDb, config.Genesis)
	if err != nil {
		return nil, err
	}
	engine, err := uncconfig.CreateConsensusEngine(chainConfig, chainDb)
	if err != nil {
		return nil, err
	}
	networkID := config.NetworkId
	if networkID == 0 {
		networkID = chainConfig.ChainID.Uint64()
	}
	unc := &Utility{
		config:            config,
		merger:            consensus.NewMerger(chainDb),
		chainDb:           chainDb,
		eventMux:          stack.EventMux(),
		accountManager:    stack.AccountManager(),
		engine:            engine,
		closeBloomHandler: make(chan struct{}),
		networkID:         networkID,
		gasPrice:          config.Miner.GasPrice,
		unicrpytbase:      config.Miner.Unicrpytbase,
		bloomRequests:     make(chan chan *bloombits.Retrieval),
		bloomIndexer:      core.NewBloomIndexer(chainDb, params.BloomBitsBlocks, params.BloomConfirms),
		p2pServer:         stack.Server(),
		shutdownTracker:   shutdowncheck.NewShutdownTracker(chainDb),
	}
	bcVersion := rawdb.ReadDatabaseVersion(chainDb)
	var dbVer = "<nil>"
	if bcVersion != nil {
		dbVer = fmt.Sprintf("%d", *bcVersion)
	}
	log.Info("Initialising Utility protocol", "network", networkID, "dbversion", dbVer)

	if !config.SkipBcVersionCheck {
		if bcVersion != nil && *bcVersion > core.BlockChainVersion {
			return nil, fmt.Errorf("database version is v%d, Gunc %s only supports v%d", *bcVersion, params.VersionWithMeta, core.BlockChainVersion)
		} else if bcVersion == nil || *bcVersion < core.BlockChainVersion {
			if bcVersion != nil { // only print warning on upgrade, not on init
				log.Warn("Upgrade blockchain database version", "from", dbVer, "to", core.BlockChainVersion)
			}
			rawdb.WriteDatabaseVersion(chainDb, core.BlockChainVersion)
		}
	}
	var (
		vmConfig = vm.Config{
			EnablePreimageRecording: config.EnablePreimageRecording,
		}
		cacheConfig = &core.CacheConfig{
			TrieCleanLimit:      config.TrieCleanCache,
			TrieCleanNoPrefetch: config.NoPrefetch,
			TrieDirtyLimit:      config.TrieDirtyCache,
			TrieDirtyDisabled:   config.NoPruning,
			TrieTimeLimit:       config.TrieTimeout,
			SnapshotLimit:       config.SnapshotCache,
			Preimages:           config.Preimages,
			StateHistory:        config.StateHistory,
			StateScheme:         scheme,
		}
	)
	// Override the chain config with provided settings.
	var overrides core.ChainOverrides
	if config.OverrideCancun != nil {
		overrides.OverrideCancun = config.OverrideCancun
	}
	if config.OverrideVerkle != nil {
		overrides.OverrideVerkle = config.OverrideVerkle
	}
	unc.blockchain, err = core.NewBlockChain(chainDb, cacheConfig, config.Genesis, &overrides, unc.engine, vmConfig, unc.shouldPreserve, &config.TransactionHistory)
	if err != nil {
		return nil, err
	}
	unc.bloomIndexer.Start(unc.blockchain)

	if config.BlobPool.Datadir != "" {
		config.BlobPool.Datadir = stack.ResolvePath(config.BlobPool.Datadir)
	}
	blobPool := blobpool.New(config.BlobPool, unc.blockchain)

	if config.TxPool.Journal != "" {
		config.TxPool.Journal = stack.ResolvePath(config.TxPool.Journal)
	}
	legacyPool := legacypool.New(config.TxPool, unc.blockchain)

	unc.txPool, err = txpool.New(new(big.Int).SetUint64(config.TxPool.PriceLimit), unc.blockchain, []txpool.SubPool{legacyPool, blobPool})
	if err != nil {
		return nil, err
	}
	// Permit the downloader to use the trie cache allowance during fast sync
	cacheLimit := cacheConfig.TrieCleanLimit + cacheConfig.TrieDirtyLimit + cacheConfig.SnapshotLimit
	if unc.handler, err = newHandler(&handlerConfig{
		Database:       chainDb,
		Chain:          unc.blockchain,
		TxPool:         unc.txPool,
		Merger:         unc.merger,
		Network:        networkID,
		Sync:           config.SyncMode,
		BloomCache:     uint64(cacheLimit),
		EventMux:       unc.eventMux,
		RequiredBlocks: config.RequiredBlocks,
	}); err != nil {
		return nil, err
	}

	unc.miner = miner.New(unc, &config.Miner, unc.blockchain.Config(), unc.EventMux(), unc.engine, unc.isLocalBlock)
	unc.miner.SetExtra(makeExtraData(config.Miner.ExtraData))

	unc.APIBackend = &UncAPIBackend{stack.Config().ExtRPCEnabled(), stack.Config().AllowUnprotectedTxs, unc, nil}
	if unc.APIBackend.allowUnprotectedTxs {
		log.Info("Unprotected transactions allowed")
	}
	gpoParams := config.GPO
	if gpoParams.Default == nil {
		gpoParams.Default = config.Miner.GasPrice
	}
	unc.APIBackend.gpo = gasprice.NewOracle(unc.APIBackend, gpoParams)

	// Setup DNS discovery iterators.
	dnsclient := dnsdisc.NewClient(dnsdisc.Config{})
	unc.ethDialCandidates, err = dnsclient.NewIterator(unc.config.EthDiscoveryURLs...)
	if err != nil {
		return nil, err
	}
	unc.snapDialCandidates, err = dnsclient.NewIterator(unc.config.SnapDiscoveryURLs...)
	if err != nil {
		return nil, err
	}

	// Start the RPC service
	unc.netRPCService = uncapi.NewNetAPI(unc.p2pServer, networkID)

	// Register the backend on the node
	stack.RegisterAPIs(unc.APIs())
	stack.RegisterProtocols(unc.Protocols())
	stack.RegisterLifecycle(unc)

	// Successful startup; push a marker and check previous unclean shutdowns.
	unc.shutdownTracker.MarkStartup()

	return unc, nil
}

func makeExtraData(extra []byte) []byte {
	if len(extra) == 0 {
		// create default extradata
		extra, _ = rlp.EncodeToBytes([]interface{}{
			uint(params.VersionMajor<<16 | params.VersionMinor<<8 | params.VersionPatch),
			"gunc",
			runtime.Version(),
			runtime.GOOS,
		})
	}
	if uint64(len(extra)) > params.MaximumExtraDataSize {
		log.Warn("Miner extra data exceed limit", "extra", hexutil.Bytes(extra), "limit", params.MaximumExtraDataSize)
		extra = nil
	}
	return extra
}

// APIs return the collection of RPC services the utility package offers.
// NOTE, some of these services probably need to be moved to somewhere else.
func (s *Utility) APIs() []rpc.API {
	apis := uncapi.GetAPIs(s.APIBackend)

	// Append any APIs exposed explicitly by the consensus engine
	apis = append(apis, s.engine.APIs(s.BlockChain())...)

	// Append all the local APIs and return
	return append(apis, []rpc.API{
		{
			Namespace: "unc",
			Service:   NewUnilityAPI(s),
		}, {
			Namespace: "miner",
			Service:   NewMinerAPI(s),
		}, {
			Namespace: "unc",
			Service:   downloader.NewDownloaderAPI(s.handler.downloader, s.eventMux),
		}, {
			Namespace: "admin",
			Service:   NewAdminAPI(s),
		}, {
			Namespace: "debug",
			Service:   NewDebugAPI(s),
		}, {
			Namespace: "net",
			Service:   s.netRPCService,
		},
	}...)
}

func (s *Utility) ResetWithGenesisBlock(gb *types.Block) {
	s.blockchain.ResetWithGenesisBlock(gb)
}

func (s *Utility) Unicrpytbase() (eb common.Address, err error) {
	s.lock.RLock()
	unicrpytbase := s.unicrpytbase
	s.lock.RUnlock()

	if unicrpytbase != (common.Address{}) {
		return unicrpytbase, nil
	}
	return common.Address{}, errors.New("unicrpytbase must be explicitly specified")
}

// isLocalBlock checks if the specified block is mined
// by local miner accounts.
//
// We regard two types of accounts as local miner account: unicrpytbase
// and accounts specified via `txpool.locals` flag.
func (s *Utility) isLocalBlock(header *types.Header) bool {
	author, err := s.engine.Author(header)
	if err != nil {
		log.Warn("Failed to retrieve block author", "number", header.Number.Uint64(), "hash", header.Hash(), "err", err)
		return false
	}
	// Check if the given address is unicrpytbase.
	s.lock.RLock()
	unicrpytbase := s.unicrpytbase
	s.lock.RUnlock()
	if author == unicrpytbase {
		return true
	}
	// Check if the given address is specified by `txpool.local`
	// CLI flag.
	for _, account := range s.config.TxPool.Locals {
		if account == author {
			return true
		}
	}
	return false
}

// shouldPreserve checks if we should preserve the given block
// during the chain reorg depending on if the author of block
// is a local account.
func (s *Utility) shouldPreserve(header *types.Header) bool {
	// The reason we need to disable the self-reorg preserving for clique
	// is it can be probable to introduce a deadlock.
	//
	// e.g. If there are 7 available signers
	//
	// r1   A
	// r2     B
	// r3       C
	// r4         D
	// r5   A      [X] F G
	// r6    [X]
	//
	// In the round5, the in-turn signer E is offline, so the worst case
	// is A, F and G sign the block of round5 and reject the block of opponents
	// and in the round6, the last available signer B is offline, the whole
	// network is stuck.
	if _, ok := s.engine.(*clique.Clique); ok {
		return false
	}
	return s.isLocalBlock(header)
}

// SetUnicrpytbase sets the mining reward address.
func (s *Utility) SetUnicrpytbase(unicrpytbase common.Address) {
	s.lock.Lock()
	s.unicrpytbase = unicrpytbase
	s.lock.Unlock()

	s.miner.SetUnicrpytbase(unicrpytbase)
}

// StartMining starts the miner with the given number of CPU threads. If mining
// is already running, this method adjust the number of threads allowed to use
// and updates the minimum price required by the transaction pool.
func (s *Utility) StartMining() error {
	// If the miner was not running, initialize it
	if !s.IsMining() {
		// Propagate the initial price point to the transaction pool
		s.lock.RLock()
		price := s.gasPrice
		s.lock.RUnlock()
		s.txPool.SetGasTip(price)

		// Configure the local mining address
		eb, err := s.Unicrpytbase()
		if err != nil {
			log.Error("Cannot start mining without unicrpytbase", "err", err)
			return fmt.Errorf("unicrpytbase missing: %v", err)
		}
		var cli *clique.Clique
		if c, ok := s.engine.(*clique.Clique); ok {
			cli = c
		} else if cl, ok := s.engine.(*beacon.Beacon); ok {
			if c, ok := cl.InnerEngine().(*clique.Clique); ok {
				cli = c
			}
		}
		if cli != nil {
			wallet, err := s.accountManager.Find(accounts.Account{Address: eb})
			if wallet == nil || err != nil {
				log.Error("Unicrpytbase account unavailable locally", "err", err)
				return fmt.Errorf("signer missing: %v", err)
			}
			cli.Authorize(eb, wallet.SignData)
		}
		// If mining is started, we can disable the transaction rejection mechanism
		// introduced to speed sync times.
		s.handler.enableSyncedFeatures()

		go s.miner.Start()
	}
	return nil
}

// StopMining terminates the miner, both at the consensus engine level as well as
// at the block creation level.
func (s *Utility) StopMining() {
	// Update the thread count within the consensus engine
	type threaded interface {
		SetThreads(threads int)
	}
	if th, ok := s.engine.(threaded); ok {
		th.SetThreads(-1)
	}
	// Stop the block creating itself
	s.miner.Stop()
}

func (s *Utility) IsMining() bool      { return s.miner.Mining() }
func (s *Utility) Miner() *miner.Miner { return s.miner }

func (s *Utility) AccountManager() *accounts.Manager  { return s.accountManager }
func (s *Utility) BlockChain() *core.BlockChain       { return s.blockchain }
func (s *Utility) TxPool() *txpool.TxPool             { return s.txPool }
func (s *Utility) EventMux() *event.TypeMux           { return s.eventMux }
func (s *Utility) Engine() consensus.Engine           { return s.engine }
func (s *Utility) ChainDb() uncdb.Database            { return s.chainDb }
func (s *Utility) IsListening() bool                  { return true } // Always listening
func (s *Utility) Downloader() *downloader.Downloader { return s.handler.downloader }
func (s *Utility) Synced() bool                       { return s.handler.synced.Load() }
func (s *Utility) SetSynced()                         { s.handler.enableSyncedFeatures() }
func (s *Utility) ArchiveMode() bool                  { return s.config.NoPruning }
func (s *Utility) BloomIndexer() *core.ChainIndexer   { return s.bloomIndexer }
func (s *Utility) Merger() *consensus.Merger          { return s.merger }
func (s *Utility) SyncMode() downloader.SyncMode {
	mode, _ := s.handler.chainSync.modeAndLocalHead()
	return mode
}

// Protocols returns all the currently configured
// network protocols to start.
func (s *Utility) Protocols() []p2p.Protocol {
	protos := unc.MakeProtocols((*ethHandler)(s.handler), s.networkID, s.ethDialCandidates)
	if s.config.SnapshotCache > 0 {
		protos = append(protos, snap.MakeProtocols((*snapHandler)(s.handler), s.snapDialCandidates)...)
	}
	return protos
}

// Start implements node.Lifecycle, starting all internal goroutines needed by the
// Utility protocol implementation.
func (s *Utility) Start() error {
	unc.StartENRUpdater(s.blockchain, s.p2pServer.LocalNode())

	// Start the bloom bits servicing goroutines
	s.startBloomHandlers(params.BloomBitsBlocks)

	// Regularly update shutdown marker
	s.shutdownTracker.Start()

	// Figure out a max peers count based on the server limits
	maxPeers := s.p2pServer.MaxPeers
	if s.config.LightServ > 0 {
		if s.config.LightPeers >= s.p2pServer.MaxPeers {
			return fmt.Errorf("invalid peer config: light peer count (%d) >= total peer count (%d)", s.config.LightPeers, s.p2pServer.MaxPeers)
		}
		maxPeers -= s.config.LightPeers
	}
	// Start the networking layer and the light server if requested
	s.handler.Start(maxPeers)
	return nil
}

// Stop implements node.Lifecycle, terminating all internal goroutines used by the
// Utility protocol.
func (s *Utility) Stop() error {
	// Stop all the peer-related stuff first.
	s.ethDialCandidates.Close()
	s.snapDialCandidates.Close()
	s.handler.Stop()

	// Then stop everything else.
	s.bloomIndexer.Close()
	close(s.closeBloomHandler)
	s.txPool.Close()
	s.miner.Close()
	s.blockchain.Stop()
	s.engine.Close()

	// Clean shutdown marker as the last thing before closing db
	s.shutdownTracker.Stop()

	s.chainDb.Close()
	s.eventMux.Stop()

	return nil
}
