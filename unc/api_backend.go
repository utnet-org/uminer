// Copyright 2015 The go-utility Authors
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

package unc

import (
	"context"
	"errors"
	"math/big"
	"time"

	"github.com/yanhuangpai/go-utility"
	"github.com/yanhuangpai/go-utility/accounts"
	"github.com/yanhuangpai/go-utility/common"
	"github.com/yanhuangpai/go-utility/consensus"
	"github.com/yanhuangpai/go-utility/core"
	"github.com/yanhuangpai/go-utility/core/bloombits"
	"github.com/yanhuangpai/go-utility/core/rawdb"
	"github.com/yanhuangpai/go-utility/core/state"
	"github.com/yanhuangpai/go-utility/core/txpool"
	"github.com/yanhuangpai/go-utility/core/types"
	"github.com/yanhuangpai/go-utility/core/vm"
	"github.com/yanhuangpai/go-utility/event"
	"github.com/yanhuangpai/go-utility/miner"
	"github.com/yanhuangpai/go-utility/params"
	"github.com/yanhuangpai/go-utility/rpc"
	"github.com/yanhuangpai/go-utility/unc/gasprice"
	"github.com/yanhuangpai/go-utility/unc/tracers"
	"github.com/yanhuangpai/go-utility/uncdb"
)

// UncAPIBackend implements uncapi.Backend and tracers.Backend for full nodes
type UncAPIBackend struct {
	extRPCEnabled       bool
	allowUnprotectedTxs bool
	unc                 *Utility
	gpo                 *gasprice.Oracle
}

// ChainConfig returns the active chain configuration.
func (b *UncAPIBackend) ChainConfig() *params.ChainConfig {
	return b.unc.blockchain.Config()
}

func (b *UncAPIBackend) CurrentBlock() *types.Header {
	return b.unc.blockchain.CurrentBlock()
}

func (b *UncAPIBackend) SetHead(number uint64) {
	b.unc.handler.downloader.Cancel()
	b.unc.blockchain.SetHead(number)
}

func (b *UncAPIBackend) HeaderByNumber(ctx context.Context, number rpc.BlockNumber) (*types.Header, error) {
	// Pending block is only known by the miner
	if number == rpc.PendingBlockNumber {
		block := b.unc.miner.PendingBlock()
		if block == nil {
			return nil, errors.New("pending block is not available")
		}
		return block.Header(), nil
	}
	// Otherwise resolve and return the block
	if number == rpc.LatestBlockNumber {
		return b.unc.blockchain.CurrentBlock(), nil
	}
	if number == rpc.FinalizedBlockNumber {
		block := b.unc.blockchain.CurrentFinalBlock()
		if block == nil {
			return nil, errors.New("finalized block not found")
		}
		return block, nil
	}
	if number == rpc.SafeBlockNumber {
		block := b.unc.blockchain.CurrentSafeBlock()
		if block == nil {
			return nil, errors.New("safe block not found")
		}
		return block, nil
	}
	return b.unc.blockchain.GetHeaderByNumber(uint64(number)), nil
}

func (b *UncAPIBackend) HeaderByNumberOrHash(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (*types.Header, error) {
	if blockNr, ok := blockNrOrHash.Number(); ok {
		return b.HeaderByNumber(ctx, blockNr)
	}
	if hash, ok := blockNrOrHash.Hash(); ok {
		header := b.unc.blockchain.GetHeaderByHash(hash)
		if header == nil {
			return nil, errors.New("header for hash not found")
		}
		if blockNrOrHash.RequireCanonical && b.unc.blockchain.GetCanonicalHash(header.Number.Uint64()) != hash {
			return nil, errors.New("hash is not currently canonical")
		}
		return header, nil
	}
	return nil, errors.New("invalid arguments; neither block nor hash specified")
}

func (b *UncAPIBackend) HeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error) {
	return b.unc.blockchain.GetHeaderByHash(hash), nil
}

func (b *UncAPIBackend) BlockByNumber(ctx context.Context, number rpc.BlockNumber) (*types.Block, error) {
	// Pending block is only known by the miner
	if number == rpc.PendingBlockNumber {
		block := b.unc.miner.PendingBlock()
		if block == nil {
			return nil, errors.New("pending block is not available")
		}
		return block, nil
	}
	// Otherwise resolve and return the block
	if number == rpc.LatestBlockNumber {
		header := b.unc.blockchain.CurrentBlock()
		return b.unc.blockchain.GetBlock(header.Hash(), header.Number.Uint64()), nil
	}
	if number == rpc.FinalizedBlockNumber {
		header := b.unc.blockchain.CurrentFinalBlock()
		if header == nil {
			return nil, errors.New("finalized block not found")
		}
		return b.unc.blockchain.GetBlock(header.Hash(), header.Number.Uint64()), nil
	}
	if number == rpc.SafeBlockNumber {
		header := b.unc.blockchain.CurrentSafeBlock()
		if header == nil {
			return nil, errors.New("safe block not found")
		}
		return b.unc.blockchain.GetBlock(header.Hash(), header.Number.Uint64()), nil
	}
	return b.unc.blockchain.GetBlockByNumber(uint64(number)), nil
}

func (b *UncAPIBackend) BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error) {
	return b.unc.blockchain.GetBlockByHash(hash), nil
}

// GetBody returns body of a block. It does not resolve special block numbers.
func (b *UncAPIBackend) GetBody(ctx context.Context, hash common.Hash, number rpc.BlockNumber) (*types.Body, error) {
	if number < 0 || hash == (common.Hash{}) {
		return nil, errors.New("invalid arguments; expect hash and no special block numbers")
	}
	if body := b.unc.blockchain.GetBody(hash); body != nil {
		return body, nil
	}
	return nil, errors.New("block body not found")
}

func (b *UncAPIBackend) BlockByNumberOrHash(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (*types.Block, error) {
	if blockNr, ok := blockNrOrHash.Number(); ok {
		return b.BlockByNumber(ctx, blockNr)
	}
	if hash, ok := blockNrOrHash.Hash(); ok {
		header := b.unc.blockchain.GetHeaderByHash(hash)
		if header == nil {
			return nil, errors.New("header for hash not found")
		}
		if blockNrOrHash.RequireCanonical && b.unc.blockchain.GetCanonicalHash(header.Number.Uint64()) != hash {
			return nil, errors.New("hash is not currently canonical")
		}
		block := b.unc.blockchain.GetBlock(hash, header.Number.Uint64())
		if block == nil {
			return nil, errors.New("header found, but block body is missing")
		}
		return block, nil
	}
	return nil, errors.New("invalid arguments; neither block nor hash specified")
}

func (b *UncAPIBackend) PendingBlockAndReceipts() (*types.Block, types.Receipts) {
	return b.unc.miner.PendingBlockAndReceipts()
}

func (b *UncAPIBackend) StateAndHeaderByNumber(ctx context.Context, number rpc.BlockNumber) (*state.StateDB, *types.Header, error) {
	// Pending state is only known by the miner
	if number == rpc.PendingBlockNumber {
		block, state := b.unc.miner.Pending()
		if block == nil || state == nil {
			return nil, nil, errors.New("pending state is not available")
		}
		return state, block.Header(), nil
	}
	// Otherwise resolve the block number and return its state
	header, err := b.HeaderByNumber(ctx, number)
	if err != nil {
		return nil, nil, err
	}
	if header == nil {
		return nil, nil, errors.New("header not found")
	}
	stateDb, err := b.unc.BlockChain().StateAt(header.Root)
	if err != nil {
		return nil, nil, err
	}
	return stateDb, header, nil
}

func (b *UncAPIBackend) StateAndHeaderByNumberOrHash(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (*state.StateDB, *types.Header, error) {
	if blockNr, ok := blockNrOrHash.Number(); ok {
		return b.StateAndHeaderByNumber(ctx, blockNr)
	}
	if hash, ok := blockNrOrHash.Hash(); ok {
		header, err := b.HeaderByHash(ctx, hash)
		if err != nil {
			return nil, nil, err
		}
		if header == nil {
			return nil, nil, errors.New("header for hash not found")
		}
		if blockNrOrHash.RequireCanonical && b.unc.blockchain.GetCanonicalHash(header.Number.Uint64()) != hash {
			return nil, nil, errors.New("hash is not currently canonical")
		}
		stateDb, err := b.unc.BlockChain().StateAt(header.Root)
		if err != nil {
			return nil, nil, err
		}
		return stateDb, header, nil
	}
	return nil, nil, errors.New("invalid arguments; neither block nor hash specified")
}

func (b *UncAPIBackend) GetReceipts(ctx context.Context, hash common.Hash) (types.Receipts, error) {
	return b.unc.blockchain.GetReceiptsByHash(hash), nil
}

func (b *UncAPIBackend) GetLogs(ctx context.Context, hash common.Hash, number uint64) ([][]*types.Log, error) {
	return rawdb.ReadLogs(b.unc.chainDb, hash, number), nil
}

func (b *UncAPIBackend) GetTd(ctx context.Context, hash common.Hash) *big.Int {
	if header := b.unc.blockchain.GetHeaderByHash(hash); header != nil {
		return b.unc.blockchain.GetTd(hash, header.Number.Uint64())
	}
	return nil
}

func (b *UncAPIBackend) GetEVM(ctx context.Context, msg *core.Message, state *state.StateDB, header *types.Header, vmConfig *vm.Config, blockCtx *vm.BlockContext) (*vm.UVM, func() error) {
	if vmConfig == nil {
		vmConfig = b.unc.blockchain.GetVMConfig()
	}
	txContext := core.NewEVMTxContext(msg)
	var context vm.BlockContext
	if blockCtx != nil {
		context = *blockCtx
	} else {
		context = core.NewEVMBlockContext(header, b.unc.BlockChain(), nil)
	}
	return vm.NewEVM(context, txContext, state, b.unc.blockchain.Config(), *vmConfig), state.Error
}

func (b *UncAPIBackend) SubscribeRemovedLogsEvent(ch chan<- core.RemovedLogsEvent) event.Subscription {
	return b.unc.BlockChain().SubscribeRemovedLogsEvent(ch)
}

func (b *UncAPIBackend) SubscribePendingLogsEvent(ch chan<- []*types.Log) event.Subscription {
	return b.unc.miner.SubscribePendingLogs(ch)
}

func (b *UncAPIBackend) SubscribeChainEvent(ch chan<- core.ChainEvent) event.Subscription {
	return b.unc.BlockChain().SubscribeChainEvent(ch)
}

func (b *UncAPIBackend) SubscribeChainHeadEvent(ch chan<- core.ChainHeadEvent) event.Subscription {
	return b.unc.BlockChain().SubscribeChainHeadEvent(ch)
}

func (b *UncAPIBackend) SubscribeChainSideEvent(ch chan<- core.ChainSideEvent) event.Subscription {
	return b.unc.BlockChain().SubscribeChainSideEvent(ch)
}

func (b *UncAPIBackend) SubscribeLogsEvent(ch chan<- []*types.Log) event.Subscription {
	return b.unc.BlockChain().SubscribeLogsEvent(ch)
}

func (b *UncAPIBackend) SendTx(ctx context.Context, signedTx *types.Transaction) error {
	return b.unc.txPool.Add([]*types.Transaction{signedTx}, true, false)[0]
}

func (b *UncAPIBackend) GetPoolTransactions() (types.Transactions, error) {
	pending := b.unc.txPool.Pending(false)
	var txs types.Transactions
	for _, batch := range pending {
		for _, lazy := range batch {
			if tx := lazy.Resolve(); tx != nil {
				txs = append(txs, tx)
			}
		}
	}
	return txs, nil
}

func (b *UncAPIBackend) GetPoolTransaction(hash common.Hash) *types.Transaction {
	return b.unc.txPool.Get(hash)
}

func (b *UncAPIBackend) GetTransaction(ctx context.Context, txHash common.Hash) (*types.Transaction, common.Hash, uint64, uint64, error) {
	tx, blockHash, blockNumber, index := rawdb.ReadTransaction(b.unc.ChainDb(), txHash)
	return tx, blockHash, blockNumber, index, nil
}

func (b *UncAPIBackend) GetPoolNonce(ctx context.Context, addr common.Address) (uint64, error) {
	return b.unc.txPool.Nonce(addr), nil
}

func (b *UncAPIBackend) Stats() (runnable int, blocked int) {
	return b.unc.txPool.Stats()
}

func (b *UncAPIBackend) TxPoolContent() (map[common.Address][]*types.Transaction, map[common.Address][]*types.Transaction) {
	return b.unc.txPool.Content()
}

func (b *UncAPIBackend) TxPoolContentFrom(addr common.Address) ([]*types.Transaction, []*types.Transaction) {
	return b.unc.txPool.ContentFrom(addr)
}

func (b *UncAPIBackend) TxPool() *txpool.TxPool {
	return b.unc.txPool
}

func (b *UncAPIBackend) SubscribeNewTxsEvent(ch chan<- core.NewTxsEvent) event.Subscription {
	return b.unc.txPool.SubscribeTransactions(ch, true)
}

func (b *UncAPIBackend) SyncProgress() utility.SyncProgress {
	return b.unc.Downloader().Progress()
}

func (b *UncAPIBackend) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	return b.gpo.SuggestTipCap(ctx)
}

func (b *UncAPIBackend) FeeHistory(ctx context.Context, blockCount uint64, lastBlock rpc.BlockNumber, rewardPercentiles []float64) (firstBlock *big.Int, reward [][]*big.Int, baseFee []*big.Int, gasUsedRatio []float64, err error) {
	return b.gpo.FeeHistory(ctx, blockCount, lastBlock, rewardPercentiles)
}

func (b *UncAPIBackend) ChainDb() uncdb.Database {
	return b.unc.ChainDb()
}

func (b *UncAPIBackend) EventMux() *event.TypeMux {
	return b.unc.EventMux()
}

func (b *UncAPIBackend) AccountManager() *accounts.Manager {
	return b.unc.AccountManager()
}

func (b *UncAPIBackend) ExtRPCEnabled() bool {
	return b.extRPCEnabled
}

func (b *UncAPIBackend) UnprotectedAllowed() bool {
	return b.allowUnprotectedTxs
}

func (b *UncAPIBackend) RPCGasCap() uint64 {
	return b.unc.config.RPCGasCap
}

func (b *UncAPIBackend) RPCEVMTimeout() time.Duration {
	return b.unc.config.RPCEVMTimeout
}

func (b *UncAPIBackend) RPCTxFeeCap() float64 {
	return b.unc.config.RPCTxFeeCap
}

func (b *UncAPIBackend) BloomStatus() (uint64, uint64) {
	sections, _, _ := b.unc.bloomIndexer.Sections()
	return params.BloomBitsBlocks, sections
}

func (b *UncAPIBackend) ServiceFilter(ctx context.Context, session *bloombits.MatcherSession) {
	for i := 0; i < bloomFilterThreads; i++ {
		go session.Multiplex(bloomRetrievalBatch, bloomRetrievalWait, b.unc.bloomRequests)
	}
}

func (b *UncAPIBackend) Engine() consensus.Engine {
	return b.unc.engine
}

func (b *UncAPIBackend) CurrentHeader() *types.Header {
	return b.unc.blockchain.CurrentHeader()
}

func (b *UncAPIBackend) Miner() *miner.Miner {
	return b.unc.Miner()
}

func (b *UncAPIBackend) StartMining() error {
	return b.unc.StartMining()
}

func (b *UncAPIBackend) StateAtBlock(ctx context.Context, block *types.Block, reexec uint64, base *state.StateDB, readOnly bool, preferDisk bool) (*state.StateDB, tracers.StateReleaseFunc, error) {
	return b.unc.stateAtBlock(ctx, block, reexec, base, readOnly, preferDisk)
}

func (b *UncAPIBackend) StateAtTransaction(ctx context.Context, block *types.Block, txIndex int, reexec uint64) (*core.Message, vm.BlockContext, *state.StateDB, tracers.StateReleaseFunc, error) {
	return b.unc.stateAtTransaction(ctx, block, txIndex, reexec)
}
