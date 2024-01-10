package types

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
	"uminer/common/log"
	"uminer/miner-server/chainApi/rpc"
	"uminer/miner-server/chipApi"
	"uminer/miner-server/data"
	"uminer/miner-server/serverConf"
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
		log:  log.NewHelper("ChipService", logger),
		data: data,
	}
}

// ClaimComputation claim server/chips to the chain, binding miner address, obtain container cloud connection
func (s *ChipService) ClaimComputation(ctx context.Context, req *rpc.ClaimComputationRequest) (*rpc.ClaimComputationReply, error) {

	bmchips := make([]MinerChip, 0)
	for _, item := range req.ChipSets {
		bmchips = append(bmchips, MinerChip{
			SN:    item.SerialNumber,
			BusID: item.BusID,
		})
	}
	reportData := reportMinerComputation{
		Address:    req.Address,
		ServerIP:   req.ServerIP,
		BMChips:    bmchips,
		totalPower: 10,
	}

	jsonData, _ := json.Marshal(reportData)
	timeNow := strconv.FormatInt(time.Now().Unix(), 10)
	joinnData := req.Address + string(jsonData) + timeNow
	hash := sha256.New()
	hash.Write([]byte(joinnData))
	result := hash.Sum(nil)
	newTx := transaction{
		Address:   req.Address,
		From:      "",
		To:        "",
		Amount:    0,
		txData:    string(jsonData),
		TimeStamp: timeNow,
		GasFee:    0,
		Hash:      string(result),
	}

	// packed as transaction, upload to the chain
	fmt.Println(newTx)

	return &rpc.ClaimComputationReply{
		BlockHeight: 0,
		ContainerId: "",
		RangeSet:    []int64{1000, 1001},
	}, nil

}

// ChallengeComputation accept challenge by the blockchain to sign chips
func (s *ChipService) ChallengeComputation(ctx context.Context, req *rpc.ChallengeComputationRequest) (*rpc.ChallengeComputationReply, error) {

	signatures := make([]*rpc.SignatureSets, 0)

	// read data from chains db ...
	requiredChips := make([]MinerChip, 0)
	containerID := ""

	for _, item := range requiredChips {
		sign := chipApi.SignMinerChips(item.SN, item.BusID, 0, item.P2, item.PublicKey, int(item.P2Size), int(item.PublicKeySize), req.Message)
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
