package types

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"os"

	//"github.com/ethereum/go-ethereum/rpc"
	"strconv"
	"time"
	"uminer/common/log"
	"uminer/miner-server/api/chainApi/rpc"
	"uminer/miner-server/api/chipApi"
	"uminer/miner-server/data"
	"uminer/miner-server/serverConf"
	"uminer/miner-server/util"
)

type ChainService struct {
	rpc.UnimplementedChainServiceServer
	conf *serverConf.Bootstrap
	log  *log.Helper
	data *data.Data
}

func NewChainService(conf *serverConf.Bootstrap, logger log.Logger, data *data.Data) rpc.ChainServiceServer {
	return &ChainService{
		conf: conf,
		log:  log.NewHelper("ChainService", logger),
		data: data,
	}
}

// UpdateChainsStatus get basic info of blockchain
func (s *ChainService) UpdateChainsStatus(ctx context.Context, req *rpc.ReportChainsStatusRequest) (*rpc.ReportChainsStatusReply, error) {

	jsonData := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      "dontcare",
		"method":  "status",
		"params":  make([]interface{}, 0),
	}
	jsonStr, _ := json.Marshal(jsonData)

	// POST request
	r, err := http.NewRequestWithContext(ctx, http.MethodPost, nodeURL, bytes.NewReader(jsonStr))
	if err != nil {
		return &rpc.ReportChainsStatusReply{}, err
	}
	r.Header.Add("Content-Type", "application/json; charset=utf-8")
	r.Header.Add("accept-encoding", "gzip,deflate")

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(r)
	if err != nil {
		return &rpc.ReportChainsStatusReply{}, err
	}
	defer resp.Body.Close()
	gzipBytes := util.GzipApi(resp)

	res := gjson.Get(string(gzipBytes), "result").String()
	sync := gjson.Get(res, "sync_info").String()
	latestHash := gjson.Get(sync, "latest_block_hash").String()
	latestHeight := gjson.Get(sync, "latest_block_height").Int()
	latestTime := gjson.Get(sync, "latest_block_time").String()

	return &rpc.ReportChainsStatusReply{
		Computation:       "",
		Rewards:           "",
		LatestBlockHash:   latestHash,
		LatestBlockHeight: strconv.FormatInt(latestHeight, 10),
		LatestBlockTime:   latestTime,
		NumberOfMiners:    "",
	}, nil
}

// UpdateMinerStatus get basic info of every miner
func (s *ChainService) UpdateMinerStatus(ctx context.Context, req *rpc.ReportMinerStatusRequest) (*rpc.ReportMinerStatusReply, error) {

	//client, err := rpc.Dial("http://node-url")
	//if err != nil {
	//	return nil, err
	//}

	//blockHeight, err := getBlockHeight(client)
	//if err != nil {
	//	return nil, err
	//}

	return &rpc.ReportMinerStatusReply{
		Computation:     "",
		Rewards:         "",
		NumberOfBlock:   "",
		NumberOfWorkers: "",
	}, nil
}

// GetMinerKeys generate miner pri/pubK pairs
func (s *ChainService) GetMinerKeys(ctx context.Context, req *rpc.GetMinerKeysRequest) (*rpc.GetMinerKeysReply, error) {
	// first check if there is stored key pairs
	_, pubErr := os.Stat("public.pem")
	_, privErr := os.Stat("private.pem")

	var privKey, pubKey string
	if pubErr == nil && privErr == nil {
		// Read the key files
		pubKeyBytes, err := ioutil.ReadFile("public.pem")
		if err != nil {
			return nil, err
		}
		privKeyBytes, err := ioutil.ReadFile("private.pem")
		if err != nil {
			return nil, err
		}
		pubKey = string(pubKeyBytes)
		privKey = string(privKeyBytes)
	} else {
		// Generate new key pair
		pubKey, privKey = util.ED25519KeysGeneration()
		// Save the newly generated key pair
		err := ioutil.WriteFile("public.pem", []byte(pubKey), 0644)
		if err != nil {
			return nil, err
		}
		// Save the private key securely, you might want to handle this differently
		err = ioutil.WriteFile("private.pem", []byte(privKey), 0644)
		if err != nil {
			return nil, err
		}
	}

	return &rpc.GetMinerKeysReply{
		PrivateKey: privKey,
		PubKey:     pubKey,
	}, nil

}

// ClaimComputation claim server/chips to the chain, binding miner address, obtain container cloud connection
func (s *ChainService) ClaimComputation(ctx context.Context, req *rpc.ClaimComputationRequest) (*rpc.ClaimComputationReply, error) {

	//bmchips := make([]MinerChip, 0)
	//for _, item := range req.ChipSets {
	//	bmchips = append(bmchips, MinerChip{
	//		SN:    item.SerialNumber,
	//		BusID: item.BusID,
	//	})
	//}

	// miner signature
	timeNow := strconv.FormatInt(time.Now().Unix(), 10)
	joinData := req.ChipPubK + req.MinerKey + timeNow
	txStr := fmt.Sprintf("%+v", joinData)
	privKeyBytes, err := ioutil.ReadFile("private.pem")
	if err != nil {
		return nil, err
	}
	privKey := string(privKeyBytes)
	_ = util.MinerSignTx(privKey, txStr)

	// packed as transaction, upload to the chain
	txhash, err := sendTransactionAsync(ctx, req.Signature)
	if err != nil {
		return &rpc.ClaimComputationReply{}, err
	}

	return &rpc.ClaimComputationReply{
		TxHash: txhash,
	}, nil

}

// ChallengeComputation accept challenge by the blockchain to sign chips
func (s *ChainService) ChallengeComputation(ctx context.Context, req *rpc.ChallengeComputationRequest) (*rpc.ChallengeComputationReply, error) {

	signatures := make([]*rpc.SignatureSets, 0)

	// read data from chains db ...
	requiredChips := make([]MinerChip, 0)
	containerID := ""
	if len(requiredChips) == 0 {
		return &rpc.ChallengeComputationReply{
			ContainerID:   containerID,
			SignatureSets: signatures,
			Status:        false,
		}, errors.New("no chip is selected")
	}

	// obtain the devId of the chip
	cardLists := chipApi.RemoteGetChipInfo(req.Url)
	for _, item := range requiredChips {
		devId := -1
		for _, card := range cardLists {
			if card.SerialNum == item.SN {
				for _, chip := range card.Chips {
					if chip.BusId == item.BusID {
						id, _ := strconv.ParseInt(chip.DevId, 10, 64)
						devId = int(id)
					}
				}
			}

		}
		if devId == -1 {
			return &rpc.ChallengeComputationReply{
				ContainerID:   containerID,
				SignatureSets: signatures,
				Status:        false,
			}, errors.New("no chip is selected")
		}

		sign := chipApi.SignMinerChips(devId, item.P2, item.PublicKey, int(item.P2Size), int(item.PublicKeySize), req.Message)
		signatures = append(signatures, &rpc.SignatureSets{
			SerialNumber: item.SN,
			BusID:        item.BusID,
			Signature:    sign.Signature,
		})
	}

	return &rpc.ChallengeComputationReply{
		ContainerID:   containerID,
		SignatureSets: signatures,
		Status:        true,
	}, nil
}
