package types

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"github.com/tidwall/gjson"
	"google.golang.org/grpc"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strings"
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

	//arg := `{"serial":` + req.SerialNumber + `,"busid":` + req.BusId + `, "power":` + req.Power + `,"p2key":` + req.P2 + `,"pubkey":` + req.PublicKey + `,"p2keysize":` + req.P2Size + `,"pubkeysize":` + req.PublicKeySize + `}`

	// command on near nodes at utility-cli-js  (KeyPath for chips.json)
	err := os.Setenv("unc", req.NearPath)
	if err != nil {
		fmt.Println("设置环境变量失败:", err)
		return nil, err
	}
	cmdString := os.Getenv("unc") + " extensions register-rsa-keys unc use-file " + req.ChipFilePath + " with-init-call network-config custom sign-with-plaintext-private-key --signer-public-key " + req.FounderPubK +
		" --signer-private-key " + req.FounderPrivK + " send"
	parts := strings.Fields(cmdString)
	order := exec.Command(parts[0], parts[1:]...)
	output, err := order.CombinedOutput()
	fmt.Println(string(output))
	//if err != nil {
	//	fmt.Println("Error executing command:", err)
	//	return nil, err
	//}
	// get error
	geterror := regexp.MustCompile(`Error:\s*(.+)`)
	matches := geterror.FindStringSubmatch(string(output))
	if len(matches) != 0 {
		errMsg := strings.Join(matches, ", ")
		return nil, errors.New(errMsg)
	}

	// get tx id
	re := regexp.MustCompile(`Transaction ID:\s*(.+)`)
	matches = re.FindStringSubmatch(string(output))
	// 提取signature字段的值
	txhash := matches[1]
	fmt.Println("tx id:", txhash)

	return &rpc.ReportChipReply{TxHash: txhash}, nil

}

// GetMinerKeys generate miner pri/pubK pairs(if no private ket is sent, it will generate automatically)
func (s *ChainService) GetMinerKeys(ctx context.Context, req *rpc.GetMinerKeysRequest) (*rpc.GetMinerKeysReply, error) {
	// first check if there is stored key pairs
	_, pubErr := os.Stat("public.pem")
	_, privErr := os.Stat("private.pem")

	var mnemonic, privKey, pubKey string
	if pubErr == nil && privErr == nil && req.PrivateKey == "" {
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
		mnemonic, pubKey, privKey = util.ED25519AddressGeneration(req.PrivateKey)
		// Save the newly generated key pair
		err := ioutil.WriteFile("mnemonic.pem", []byte(mnemonic), 0644)
		if err != nil {
			return nil, err
		}
		err = ioutil.WriteFile("public.pem", []byte(pubKey), 0644)
		if err != nil {
			return nil, err
		}
		err = ioutil.WriteFile("private.pem", []byte(privKey), 0644)
		if err != nil {
			return nil, err
		}
	}
	/* obtain pubKey and fill it into the ' "challenge_key": "ed25519..." ' of miner_key.json */

	publicKeyBytes := base58.Decode(pubKey)
	if len(publicKeyBytes) != 32 {
		return nil, errors.New("Invalid public key length")
	}
	// 将Ed25519公钥转换为地址
	publicKeyHex := hex.EncodeToString(publicKeyBytes)

	return &rpc.GetMinerKeysReply{
		PrivateKey: "privKey",
		PubKey:     pubKey,
		Address:    publicKeyHex,
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
	order := exec.Command(req.NearPath, "extensions", "create-challenge-rsa", req.AccountId, "use-file", req.KeyPath, "without-init-call", "network-config", "custom", "sign-with-plaintext-private-key", "--signer-public-key", pubKey, "--signer-private-key", privKey, "display")
	output, err := order.CombinedOutput()
	fmt.Println(string(output))
	//if err != nil {
	//	fmt.Println("Error executing command:", err)
	//	return nil, err
	//}
	// get signature
	//re := regexp.MustCompile(`Signed transaction \(serialized as base64\):\s*(.+)`)
	//txhash, err := connect.SendTransactionAsync(ctx, signature)
	//if err != nil {
	//	return &rpc.ClaimChipComputationReply{}, err
	//}

	// get error
	geterror := regexp.MustCompile(`Error:\s*(.+)`)
	matches := geterror.FindStringSubmatch(string(output))
	if len(matches) != 0 {
		errMsg := strings.Join(matches, ", ")
		return nil, errors.New(errMsg)
	}

	// get tx id
	re := regexp.MustCompile(`Transaction ID:\s*(.+)`)
	matches = re.FindStringSubmatch(string(output))
	// 提取signature字段的值
	txhash := matches[1]
	fmt.Println("tx id:", txhash)

	return &rpc.ClaimChipComputationReply{
		TxHash: txhash,
	}, nil

}

// GetMinerChipsList miner get all chips from chain
func (s *ChainService) GetMinerChipsList(ctx context.Context, req *rpc.GetMinerChipsListRequest) (*rpc.GetMinerChipsListReply, error) {

	jsonData := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      "dontcare",
		"method":  "query",
		"params":  map[string]interface{}{"request_type": "view_chip_list", "finality": "final", "account_id": req.AccountId},
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
		return &rpc.GetMinerChipsListReply{}, err
	}
	defer resp.Body.Close()
	gzipBytes := util.GzipApi(resp)
	if gjson.Get(string(gzipBytes), "error").String() != "" {
		return &rpc.GetMinerChipsListReply{}, errors.New(gjson.Get(string(gzipBytes), "error").String())
	}

	res := gjson.Get(string(gzipBytes), "result").String()
	chipLists := gjson.Get(res, "chips").Array()

	chips := make([]*rpc.ChipDetails, 0)
	for _, item := range chipLists {
		chips = append(chips, &rpc.ChipDetails{
			SerialNumber:  gjson.Get(item.String(), "serial_number").String(),
			BusId:         gjson.Get(item.String(), "bus_id").String(),
			Power:         gjson.Get(item.String(), "power").Int(),
			P2:            gjson.Get(item.String(), "p2").String(),
			PublicKey:     gjson.Get(item.String(), "public_key").String(),
			P2Size:        1680, //gjson.Get(item.String(), "p2_size").Int(),
			PublicKeySize: 426,  //gjson.Get(item.String(), "public_key_size").Int(),
		})
	}

	return &rpc.GetMinerChipsListReply{Chips: chips}, nil

}

// ChallengeComputation accept challenge by the blockchain to sign chips
func (s *ChainService) ChallengeComputation(ctx context.Context, req *rpc.ChallengeComputationRequest) (*rpc.ChallengeComputationReply, error) {

	// read data from chains db ...
	signatures := make([]*rpc.SignatureSets, 0)
	if len(req.Chips) == 0 {
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
			Url:       each + ":30345",
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

		// begin to search
		cardLists := response.Cards
		conn.Close()
		for _, card := range cardLists {
			devId := -1
			for _, item := range req.Chips {
				if card.SerialNum == item.SerialNumber {
					for _, chip := range card.Chips {
						if chip.BusId == item.BusId {
							id, _ := strconv.ParseInt(chip.DevId, 10, 64)
							devId = int(id)
						}
					}
				}
				// not found, proceed to next chip
				if devId == -1 {
					continue
					//return &rpc.ChallengeComputationReply{
					//	SignatureSets: signatures,
					//	TxHash:        "",
					//}, errors.New("no chip is selected")
				}
				// found, sign
				sign := chipApi.SignMinerChips(devId, item.P2, item.PublicKey, int(item.P2Size), int(item.PublicKeySize), req.Message)
				signatures = append(signatures, &rpc.SignatureSets{
					SerialNumber: item.SerialNumber,
					BusID:        item.BusId,
					Signature:    sign.Signature,
				})
			}
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
