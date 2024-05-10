package types

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tidwall/gjson"
	"github.com/tyler-smith/go-bip39"
	"google.golang.org/grpc"
	"io/ioutil"
	"math/big"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
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

type KeyData struct {
	AccountID  string `json:"implicit_account_id"`
	PublicKey  string `json:"public_key"`
	PrivateKey string `json:"private_key"`
}

// UpdateChainsStatus get basic information and status of blockchain from utility node
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

// ReportChip foundation of utility report chips, uploading to chain
func (s *ChainService) ReportChip(ctx context.Context, req *rpc.ReportChipRequest) (*rpc.ReportChipReply, error) {

	// command on utility nodes at utility-cli-js (KeyPath for chips.json)
	err := os.Setenv("unc", req.NodePath)
	if err != nil {
		fmt.Println("setting environment variable fail:", err)
		return nil, err
	}
	// read foundation keys from unc.json file
	//type UNCKeyPair struct {
	//	FounderPubK  string `json:"public_key"`
	//	FounderPrivK string `json:"private_key"`
	//}
	//js, err := ioutil.ReadFile("../jsonfile/unc.json")
	//if err != nil {
	//	fmt.Println("Failed to read file:", err)
	//	return nil, err
	//}
	//var keyPair UNCKeyPair
	//if err := json.Unmarshal(js, &keyPair); err != nil {
	//	fmt.Println("Failed to unmarshal JSON:", err)
	//	return nil, err
	//}
	cmdString := os.Getenv("unc") + " extensions register-rsa-keys " + req.Founder + " use-file " + req.ChipFilePath + " with-init-call network-config custom sign-with-access-key-file " + req.FounderKeyPath + " send"
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
	if len(matches) == 0 {
		return nil, errors.New("transaction failed")
	}
	txhash := matches[1]
	fmt.Println("tx id:", txhash)

	return &rpc.ReportChipReply{TxHash: txhash}, nil

}

// GetMinerAccountKeys miner generate miner private key/public key pairs with its account address (if no private ket is sent, it will generate automatically)
func (s *ChainService) GetMinerAccountKeys(ctx context.Context, req *rpc.GetMinerAccountKeysRequest) (*rpc.GetMinerAccountKeysReply, error) {

	// check if account access key exists
	file, fileErr := filepath.Glob("*.json")
	var address, pubKey, privateKey string
	if (fileErr == nil && len(file) != 0) && (req.Mnemonic == "" || len(strings.Fields(req.Mnemonic)) != 12) {
		// Read the key files
		fileContent, err := os.Open(file[0])
		if err != nil {
			return nil, err
		}
		var keyData KeyData
		err = json.NewDecoder(fileContent).Decode(&keyData)
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
			if err != nil {
				return nil, err
			}
		}
		pubKey = keyData.PublicKey
		privateKey = keyData.PrivateKey
		address = keyData.AccountID
	} else {
		// Generate new key pair using unc-rs client
		entropy, err := bip39.NewEntropy(128)
		if err != nil {
			return nil, err
		}
		mnemonic, _ := bip39.NewMnemonic(entropy)
		if req.Mnemonic != "" && len(strings.Fields(req.Mnemonic)) == 12 {
			mnemonic = req.Mnemonic
		}
		err = os.Setenv("unc", req.NodePath)
		if err != nil {
			fmt.Println("setting environment variable fail:", err)
			return nil, err
		}
		command := os.Getenv("unc")
		args := []string{
			"account", "create-account", "fund-later", "use-seed-phrase", mnemonic,
			"--seed-phrase-hd-path", "m/44'/397'/0'",
			"save-to-folder", ".",
		}
		// execute the command
		/* mnemonic, private key, public Key and address will be stored into the 'address'.json */
		order := exec.Command(command, args...)
		output, err := order.CombinedOutput()
		fmt.Println(string(output))
		// get error
		geterror := regexp.MustCompile(`Error:\s*(.+)`)
		matches := geterror.FindStringSubmatch(string(output))
		if len(matches) != 0 {
			errMsg := strings.Join(matches, ", ")
			return nil, errors.New(errMsg)
		}
		// Read the key files
		file, fileErr = filepath.Glob("*.json")
		fileContent, err := os.Open(file[0])
		if err != nil {
			return nil, err
		}
		var keyData KeyData
		err = json.NewDecoder(fileContent).Decode(&keyData)
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
			if err != nil {
				return nil, err
			}
		}
		pubKey = keyData.PublicKey
		privateKey = keyData.PrivateKey
		address = keyData.AccountID
	}

	// generate validator_key.json for signing keys at claim computation
	rawdata := map[string]string{
		"account_id":  address,
		"public_key":  pubKey,
		"private_key": privateKey,
	}
	_, err := ioutil.ReadFile("../jsonfile/validator_key.json")
	if err == nil {
		file, err := os.OpenFile("../jsonfile/validator_key.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return nil, err
		}
		encoder := json.NewEncoder(file)
		err = encoder.Encode(rawdata)
		if err != nil {
			fmt.Println("Error writing JSON to file:", err)
			return nil, err
		}
	} else {
		file, err := os.Create("../jsonfile/validator_key.json")
		if err != nil {
			fmt.Println("Error creating JSON file:", err)
			return nil, err
		}
		encoder := json.NewEncoder(file)
		err = encoder.Encode(rawdata)
		if err != nil {
			fmt.Println("Error encoding JSON:", err)
			return nil, err
		}
	}

	return &rpc.GetMinerAccountKeysReply{
		PubKey:  pubKey,
		Address: address,
	}, nil

}

// FaucetNewAccount founder send unc to a new miner for activation
func (s *ChainService) FaucetNewAccount(ctx context.Context, req *rpc.FaucetNewAccountRequest) (*rpc.FaucetNewAccountReply, error) {

	// load founder key
	file, fileErr := filepath.Glob("../jsonfile/unc.json")
	if fileErr != nil {
		return nil, fileErr
	}
	fileContent, err := os.Open(file[0])
	if err != nil {
		return nil, err
	}
	var keyData KeyData
	err = json.NewDecoder(fileContent).Decode(&keyData)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		if err != nil {
			return nil, err
		}
	}
	pubKey := keyData.PublicKey
	privateKey := keyData.PrivateKey
	// start transfer
	amount := req.Amount + " unc"
	order := exec.Command(req.NodePath, "tokens", req.Sender, "send-unc", req.AccountId, amount, "network-config", req.Net, "sign-with-plaintext-private-key",
		"--signer-public-key", pubKey, "--signer-private-key", privateKey, "send")
	output, _ := order.CombinedOutput()
	fmt.Println(string(output))
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
	if len(matches) == 0 {
		return nil, errors.New("transaction failed")
	}
	txhash := matches[1]
	fmt.Println("tx id:", txhash)
	return &rpc.FaucetNewAccountReply{TxHash: txhash}, nil

}

// ClaimStake claim amount of token deposit to the chain as stake before start mining
func (s *ChainService) ClaimStake(ctx context.Context, req *rpc.ClaimStakeRequest) (*rpc.ClaimStakeReply, error) {

	// check if account access key exists
	file, fileErr := filepath.Glob("*.json")
	if fileErr != nil {
		return nil, fileErr
	}
	fileContent, err := os.Open(file[0])
	if err != nil {
		return nil, err
	}
	var keyData KeyData
	err = json.NewDecoder(fileContent).Decode(&keyData)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		if err != nil {
			return nil, err
		}
	}
	pubKey := keyData.PublicKey
	privateKey := keyData.PrivateKey

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

	// command on utility nodes at utility:unc-validator
	err = os.Setenv("unc-validator", req.NodePath)
	if err != nil {
		fmt.Println("setting environment variable fail:", err)
		return nil, err
	}
	command := os.Getenv("unc-validator")
	amount := req.Amount + " unc"
	args := []string{
		"pledging", "pledge-proposal", req.AccountId, pubKey, amount, "network-config", req.Net, "sign-with-plaintext-private-key", "--signer-public-key",
		pubKey, "--signer-private-key", privateKey, "send",
	}
	// execute the command
	order := exec.Command(command, args...)
	output, err := order.CombinedOutput()
	fmt.Println(string(output))
	// get error
	geterror := regexp.MustCompile(`error:\s*(.+)`)
	matches := geterror.FindStringSubmatch(string(output))
	if len(matches) != 0 {
		errMsg := strings.Join(matches, ", ")
		return nil, errors.New(errMsg)
	}
	geterror = regexp.MustCompile(`Error:\s*(.+)`)
	matches = geterror.FindStringSubmatch(string(output))
	if len(matches) != 0 {
		errMsg := strings.Join(matches, ", ")
		return nil, errors.New(errMsg)
	}
	// get tx id
	re := regexp.MustCompile(`Transaction ID:\s*(.+)`)
	matches = re.FindStringSubmatch(string(output))
	if len(matches) == 0 {
		return nil, errors.New("transaction failed")
	}
	txhash := matches[1]
	fmt.Println("tx id:", txhash)

	return &rpc.ClaimStakeReply{
		TransId: txhash,
		Status:  "1",
	}, nil
}

// AddChipOwnership add a key ownership to a miner identified by chip public key
func (s *ChainService) AddChipOwnership(ctx context.Context, req *rpc.AddChipOwnershipRequest) (*rpc.AddChipOwnershipReply, error) {

	// get keys
	file, fileErr := filepath.Glob("*.json")
	if fileErr != nil {
		return nil, fileErr
	}
	fileContent, err := os.Open(file[0])
	if err != nil {
		return nil, err
	}
	var keyData KeyData
	err = json.NewDecoder(fileContent).Decode(&keyData)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		if err != nil {
			return nil, err
		}
	}
	pubKey := keyData.PublicKey
	privateKey := keyData.PrivateKey

	// check if miner account exist
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
		return &rpc.AddChipOwnershipReply{}, errors.New(datas)
	}
	res := gjson.Get(string(gzipBytes), "result").String()
	permission := gjson.Get(res, "permission").String()
	if permission != "FullAccess" {
		return &rpc.AddChipOwnershipReply{}, errors.New("miner account is not accessible")
	}

	// generate miner_challengeKey.json for challenge at claim computation
	rawdata := map[string]string{
		"public_key":    "rsa2048:" + req.ChipPubK,
		"challenge_key": pubKey,
	}
	_, err = ioutil.ReadFile("../jsonfile/miner_challengeKey.json")
	if err == nil {
		file, err := os.OpenFile("../jsonfile/miner_challengeKey.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return nil, err
		}
		encoder := json.NewEncoder(file)
		err = encoder.Encode(rawdata)
		if err != nil {
			fmt.Println("Error writing JSON to file:", err)
			return nil, err
		}
	} else {
		file, err := os.Create("../jsonfile/miner_challengeKey.json")
		if err != nil {
			fmt.Println("Error creating JSON file:", err)
			return nil, err
		}
		encoder := json.NewEncoder(file)
		err = encoder.Encode(rawdata)
		if err != nil {
			fmt.Println("Error encoding JSON:", err)
			return nil, err
		}
	}

	/* miner signature: command on unc node */
	rsaKey := "rsa2048:" + req.ChipPubK
	order := exec.Command(req.NodePath, "account", "add-key", req.AccountId, "grant-full-access", "use-manually-provided-public-key", rsaKey, "network-config", req.Net, "sign-with-plaintext-private-key",
		"--signer-public-key", pubKey, "--signer-private-key", privateKey, "send")
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
	if len(matches) == 0 {
		return nil, errors.New("transaction failed")
	}
	txhash := matches[1]
	fmt.Println("tx id:", txhash)

	return &rpc.AddChipOwnershipReply{
		TxHash: txhash,
	}, nil

}

// ClaimChipComputation claim server/chips to the chain, binding miner address
func (s *ChainService) ClaimChipComputation(ctx context.Context, req *rpc.ClaimChipComputationRequest) (*rpc.ClaimChipComputationReply, error) {

	// get keys
	file, fileErr := filepath.Glob("*.json")
	if fileErr != nil {
		return nil, fileErr
	}
	fileContent, err := os.Open(file[0])
	if err != nil {
		return nil, err
	}
	var keyData KeyData
	err = json.NewDecoder(fileContent).Decode(&keyData)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		if err != nil {
			return nil, err
		}
	}
	pubKey := keyData.PublicKey

	// check if miner account exist
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
		return &rpc.ClaimChipComputationReply{}, errors.New(datas)
	}
	res := gjson.Get(string(gzipBytes), "result").String()
	permission := gjson.Get(res, "permission").String()
	if permission != "FullAccess" {
		return &rpc.ClaimChipComputationReply{}, errors.New("miner account is not accessible")
	}

	/* miner signature: command on unc node (KeyPath: miner_challengeKey.json + validator_key.json) */
	order := exec.Command(req.NodePath, "extensions", "create-challenge-rsa", req.AccountId, "use-file", req.ChallengeKeyPath, "without-init-call", "network-config", req.Net, "sign-with-access-key-file", req.SignerKeyPath, "send")
	output, err := order.CombinedOutput()
	fmt.Println(string(output))
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
	if len(matches) == 0 {
		return nil, errors.New("transaction failed")
	}
	txhash := matches[1]
	fmt.Println("tx id:", txhash)

	return &rpc.ClaimChipComputationReply{
		TxHash: txhash,
	}, nil

}

// GetMinerChipsList miner get all claimed chips by rpc method at the utility node
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
	total := gjson.Get(res, "total_power").String()
	num := new(big.Float)
	num.SetString(total)
	divisor := new(big.Int)
	divisor.SetString("1000000000000", 10) // = 1e12
	amount := new(big.Float).Quo(num, new(big.Float).SetInt(divisor))
	amountString := amount.Text('f', 1)
	chipLists := gjson.Get(res, "chips").Array()

	chips := make([]*rpc.ChipDetails, 0)
	for _, item := range chipLists {
		chips = append(chips, &rpc.ChipDetails{
			SerialNumber:  gjson.Get(item.String(), "sn").String(),
			BusId:         gjson.Get(item.String(), "bus_id").String(),
			Power:         gjson.Get(item.String(), "power").Int(),
			P2:            gjson.Get(item.String(), "p2key").String(),
			PublicKey:     gjson.Get(item.String(), "public_key").String(),
			P2Size:        1680, //gjson.Get(item.String(), "p2_size").Int(),
			PublicKeySize: 426,  //gjson.Get(item.String(), "public_key_size").Int(),
		})
	}

	return &rpc.GetMinerChipsListReply{Chips: chips, TotalPower: amountString}, nil

}

// ChallengeComputation accept challenge by the blockchain nodes to sign chips when chosen to report a new block
func (s *ChainService) ChallengeComputation(ctx context.Context, req *rpc.ChallengeComputationRequest) (*rpc.ChallengeComputationReply, error) {

	// read data from chains db ...
	signatures := make([]*rpc.SignatureSets, 0)
	if len(req.Chips) == 0 {
		return &rpc.ChallengeComputationReply{
			SignatureSets: signatures,
			TxHash:        "",
		}, errors.New("no chip is selected")
	}
	// call every worker and sign the corresponding chips in rpc method
	for _, each := range req.Url {

		clientDeadline := time.Now().Add(time.Duration(connect.Delay * time.Second))
		ctx1, cancel := context.WithDeadline(context.Background(), clientDeadline)
		defer cancel()
		conn, err := grpc.DialContext(ctx1, each+":7001", grpc.WithInsecure())
		if err != nil {
			fmt.Println("Error connecting to RPC server:", err)
			continue
		}
		request := &chipRPC.ChipsRequest{
			// the port 30345 of each worker provides the information of all the chips
			Url:       each + ":30345",
			SerialNum: "",
			BusId:     "",
		}
		client := chipRPC.NewChipServiceClient(conn)
		// list all chips
		var response *chipRPC.ListChipsReply
		response, err = client.ListAllChips(ctx1, request, grpc.WaitForReady(true))
		if err != nil {
			fmt.Println("Error query chip information:", err)
			continue
		}

		// begin to search chip devId
		cardLists := make([]*chipRPC.CardItem, 0)
		if response != nil {
			cardLists = response.Cards
		}
		conn.Close()
		for _, item := range req.Chips {
			devId := -1
			for _, card := range cardLists {
				if card.SerialNum == item.SerialNumber {
					for _, chip := range card.Chips {
						if chip.BusId == item.BusId {
							id, _ := strconv.ParseInt(chip.DevId, 10, 64)
							devId = int(id)
						}
					}
				}
			}
			// if the current chip not found, proceed to next chip
			if devId == -1 {
				fmt.Println("chip", item.SerialNumber, item.BusId, "is not found in ", each)
				continue
			}
			// found the right chip and sign
			sign := chipApi.SignMinerChips(devId, item.P2, item.PublicKey, int(item.P2Size), int(item.PublicKeySize), req.Message)
			signatures = append(signatures, &rpc.SignatureSets{
				SerialNumber: item.SerialNumber,
				BusID:        item.BusId,
				Signature:    sign.Signature,
			})
		}
	}

	// broadcast the signature to nodes (now under development)
	privKeyBytes, err := ioutil.ReadFile("private.pem")
	if err != nil {
		return nil, err
	}
	privKey := string(privKeyBytes)
	txStr := fmt.Sprintf("%+v", signatures)
	txhash, err := connect.SendTransactionAsync(ctx, privKey+txStr)
	if err != nil {
		return nil, err
	}

	return &rpc.ChallengeComputationReply{
		SignatureSets: signatures,
		TxHash:        txhash,
	}, nil

}
