package types

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tidwall/gjson"
	"google.golang.org/grpc"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	chipRPC "uminer/miner-server/api/chipApi/rpc"
	"uminer/miner-server/cmd"
	"uminer/miner-server/service/connect"

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
	r, err := http.NewRequestWithContext(ctx, http.MethodPost, cmd.NodeURL, bytes.NewReader(jsonStr))
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

// ReportChip foundation report chips uploading to chain
func (s *ChainService) ReportChip(ctx context.Context, req *rpc.ReportChipRequest) (*rpc.ReportChipReply, error) {

	arg := `{"serial":` + req.SerialNumber + `,"busid":` + req.BusId + `, "power":` + req.Power + `,"p2key":` + req.P2 + `,"pubkey":` + req.PublicKey + `,"p2keysize":` + req.P2Size + `,"pubkeysize":` + req.PublicKeySize + `}`

	// command on near nodes at near-cli-js  (KeyPath for validator_key.json)
	order := exec.Command(req.NearPath, "extensions register-rsa-keys ", req.Founder, "use-file", req.KeyPath, " with-init-call json-args ", arg, "network-config my-private-chain-id sign-with-plaintext-private-key --signer-public-key ", req.FounderPubK,
		" --signer-private-key ", req.FounderPrivK, " send")
	output, err := order.CombinedOutput()
	if err != nil {
		fmt.Println("Error executing command:", err)
		return nil, err
	}
	fmt.Println("output:", output)

	return &rpc.ReportChipReply{TxHash: ""}, nil

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
	/* obtain pubKey and fill it into the ' "challenge_key": "ed25519..." ' of miner_key.json */
	return &rpc.GetMinerKeysReply{
		PrivateKey: privKey,
		PubKey:     pubKey,
	}, nil

}

// ClaimStake claim amount of token deposit to the chain as stake before start mining
func (s *ChainService) ClaimStake(ctx context.Context, req *rpc.ClaimStakeRequest) (*rpc.ClaimStakeReply, error) {

	// check if access key exist
	pubKeyBytes, err := ioutil.ReadFile("public.pem")
	if err != nil {
		return nil, err
	}
	pubKey := "ed25519:" + string(pubKeyBytes)
	jsonData := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      "dontcare",
		"method":  "query",
		"params":  map[string]interface{}{"request_type": "view_access_key", "finality": "final", "account_id": req.AccountId, "public_key": pubKey},
	}
	jsonStr, _ := json.Marshal(jsonData)
	clientDeadline := time.Now().Add(time.Duration(connect.Delay * time.Second))
	ctx, cancel := context.WithDeadline(context.Background(), clientDeadline)
	defer cancel()
	r, err := http.NewRequestWithContext(ctx, http.MethodPost, cmd.NodeURL, bytes.NewReader(jsonStr))
	if err != nil {
		return nil, err
	}
	r.Header.Add("Content-Type", "application/json; charset=utf-8")
	r.Header.Add("accept-encoding", "gzip,deflate")

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	gzipBytes := util.GzipApi(resp)
	// no such account
	if gjson.Get(string(gzipBytes), "error").String() != "" {
		errs := gjson.Get(string(gzipBytes), "error").String()
		datas := gjson.Get(errs, "data").String()
		return &rpc.ClaimStakeReply{}, errors.New(datas)
	}
	res := gjson.Get(string(gzipBytes), "result").String()
	permission := gjson.Get(res, "permission").String()
	if permission != "FullAccess" {
		return &rpc.ClaimStakeReply{}, errors.New("miner account is not accessible")
	}

	// command on near nodes at near-cli-js  (KeyPath for validator_key.json)
	order := exec.Command(req.NearPath, "stake", req.AccountId, pubKey, req.Amount, "--keyPath", req.KeyPath)
	output, err := order.CombinedOutput()
	if err != nil {
		fmt.Println("Error executing command:", err)
		return nil, err
	}
	fmt.Println("output:", output)

	return &rpc.ClaimStakeReply{
		TransId: "",
		Status:  "1",
	}, nil
}

// ClaimChipComputation claim server/chips to the chain, binding miner address, obtain container cloud connection
func (s *ChainService) ClaimChipComputation(ctx context.Context, req *rpc.ClaimChipComputationRequest) (*rpc.ClaimChipComputationReply, error) {

	//bmchips := make([]MinerChip, 0)
	//for _, item := range req.ChipSets {
	//	bmchips = append(bmchips, MinerChip{
	//		SN:    item.SerialNumber,
	//		BusID: item.BusID,
	//	})
	//}

	// check if miner account exist
	jsonData := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      "czROwmnXE",
		"method":  "query",
		"params":  map[string]interface{}{"account_id": req.AccountId, "finality": "final", "request_type": "view_account"},
	}
	jsonStr, _ := json.Marshal(jsonData)
	clientDeadline := time.Now().Add(time.Duration(connect.Delay * time.Second))
	ctx, cancel := context.WithDeadline(context.Background(), clientDeadline)
	defer cancel()
	r, err := http.NewRequestWithContext(ctx, http.MethodPost, cmd.NodeURL, bytes.NewReader(jsonStr))
	if err != nil {
		return nil, err
	}
	r.Header.Add("Content-Type", "application/json; charset=utf-8")
	r.Header.Add("accept-encoding", "gzip,deflate")

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	gzipBytes := util.GzipApi(resp)
	// no such account
	if gjson.Get(string(gzipBytes), "error").String() != "" {
		return &rpc.ClaimChipComputationReply{}, errors.New("miner account is not registered yet")
	}

	/* miner signature */
	pubKeyBytes, err := ioutil.ReadFile("public.pem")
	if err != nil {
		return nil, err
	}
	pubKey := "ed25519:" + string(pubKeyBytes)
	privKeyBytes, err := ioutil.ReadFile("private.pem")
	if err != nil {
		return nil, err
	}
	privKey := "ed25519:" + string(privKeyBytes)
	// command on near nodes at near-cli-js (KeyPath for miner_key.json)
	order := exec.Command(req.NearPath, "extensions", "create-challenge-rsa", req.AccountId, "use-file", req.KeyPath, "without-init-call", "network-config", "my-private-chain-id", "sign-with-plaintext-private-key", "--signer-public-key", pubKey, "--signer-private-key", privKey, "display")
	output, err := order.CombinedOutput()
	if err != nil {
		fmt.Println("Error executing command:", err)
		return nil, err
	}
	// get signature
	re := regexp.MustCompile(`Signed transaction \(serialized as base64\):\s*(.+)`)
	matches := re.FindStringSubmatch(string(output))
	// 提取signature字段的值
	signature := matches[1]
	fmt.Println("signature:", signature)

	timeNow := strconv.FormatInt(time.Now().Unix(), 10)
	joinData := req.AccountId + req.ChipPubK + timeNow
	txStr := fmt.Sprintf("%+v", joinData)
	_ = util.MinerSignTx(privKey, txStr)

	// packed as transaction, upload to the chain
	txhash, err := connect.SendTransactionAsync(ctx, "")
	if err != nil {
		return &rpc.ClaimChipComputationReply{}, err
	}

	return &rpc.ClaimChipComputationReply{
		TxHash: txhash,
	}, nil

}

// GetMinerChipsList miner get all chips from chain
func (s *ChainService) GetMinerChipsList(ctx context.Context, req *rpc.GetMinerChipsListRequest) (*rpc.GetMinerChipsListReply, error) {

	chips := make([]*rpc.ChipDetails, 0)
	chips = append(chips, &rpc.ChipDetails{
		SerialNumber:  "test",
		BusId:         "test",
		P2:            "test",
		PublicKey:     "test",
		P2Size:        1,
		PublicKeySize: 1,
	})

	return &rpc.GetMinerChipsListReply{Chips: chips}, nil

}

// ChallengeComputation accept challenge by the blockchain to sign chips
func (s *ChainService) ChallengeComputation(ctx context.Context, req *rpc.ChallengeComputationRequest) (*rpc.ChallengeComputationReply, error) {

	// read data from chains db ...
	signatures := make([]*rpc.SignatureSets, 0)
	requiredChips := make([]connect.MinerChip, 0)
	if len(requiredChips) == 0 {
		return &rpc.ChallengeComputationReply{
			SignatureSets: signatures,
			TxHash:        "",
		}, errors.New("no chip is selected")
	}
	// call rpc of every worker and sign the chips
	for _, each := range req.Url {

		conn, err := grpc.DialContext(ctx, each+":7001", grpc.WithInsecure())
		if err != nil {
			fmt.Println("Error connecting to RPC server:", err)
			continue
		}
		request := &chipRPC.ChipsRequest{
			Url:       "http://119.120.92.239" + ":30345",
			SerialNum: "",
			BusId:     "",
		}
		client := chipRPC.NewChipServiceClient(conn)
		// list all chips
		var response *chipRPC.ListChipsReply
		response, err = client.ListAllChips(ctx, request, grpc.WaitForReady(true))
		if err != nil {
			fmt.Println("Error query chip information:", err)
			continue
		}

		// begin to sign

		cardLists := response.Cards
		conn.Close()
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
					SignatureSets: signatures,
					TxHash:        "",
				}, errors.New("no chip is selected")
			}

			sign := chipApi.SignMinerChips(devId, item.P2, item.PublicKey, int(item.P2Size), int(item.PublicKeySize), req.Message)
			signatures = append(signatures, &rpc.SignatureSets{
				SerialNumber: item.SN,
				BusID:        item.BusID,
				Signature:    sign.Signature,
			})
		}
	}

	// broadcast to nodes
	privKeyBytes, err := ioutil.ReadFile("private.pem")
	if err != nil {
		return nil, err
	}
	privKey := string(privKeyBytes)
	txStr := fmt.Sprintf("%+v", signatures)
	sign := util.MinerSignTx(privKey, txStr)
	txhash, err := connect.SendTransactionAsync(ctx, sign)
	if err != nil {
		return nil, err
	}

	return &rpc.ChallengeComputationReply{
		SignatureSets: signatures,
		TxHash:        txhash,
	}, nil

}
