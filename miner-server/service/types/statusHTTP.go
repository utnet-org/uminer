package types

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/tidwall/gjson"
	"google.golang.org/grpc"
	"math/big"
	http2 "net/http"
	"strconv"
	"strings"
	"time"
	"uminer/common/log"
	"uminer/miner-server/api/HTTP"
	chipRPC "uminer/miner-server/api/chipApi/rpc"
	"uminer/miner-server/cmd"
	"uminer/miner-server/data"
	"uminer/miner-server/serverConf"
	"uminer/miner-server/service/connect"
	"uminer/miner-server/util"
)

type MinerStatusServiceHTTP struct {
	conf *serverConf.Bootstrap
	log  *log.Helper
	data *data.Data
}

func NewMinerStatusServiceHTTP(conf *serverConf.Bootstrap, logger log.Logger, data *data.Data) *MinerStatusServiceHTTP {
	return &MinerStatusServiceHTTP{
		conf: conf,
		log:  log.NewHelper("MinerStatusService", logger),
		data: data,
	}
}

// GetNodesStatusHandler get the latest information about node and about miner himself as well
func (s *MinerStatusServiceHTTP) GetNodesStatusHandler(w http.ResponseWriter, r *http.Request) {

	// method Post
	if r.Method != http2.MethodPost {
		http2.Error(w, "Method Not Allowed", http2.StatusMethodNotAllowed)
		return
	}
	// get status
	jsonData := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      "dontcare",
		"method":  "status",
		"params":  make([]interface{}, 0),
	}
	jsonStr, _ := json.Marshal(jsonData)
	// POST request
	clientDeadline := time.Now().Add(time.Duration(connect.Delay * time.Second))
	ctx, cancel := context.WithDeadline(context.Background(), clientDeadline)
	defer cancel()
	r, err := http2.NewRequestWithContext(ctx, http2.MethodPost, cmd.NodeURL, bytes.NewReader(jsonStr))
	if err != nil {
		http2.Error(w, err.Error(), http2.StatusInternalServerError)
		return
	}
	r.Header.Add("Content-Type", "application/json; charset=utf-8")
	r.Header.Add("accept-encoding", "gzip,deflate")

	client := &http2.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(r)
	if err != nil {
		http2.Error(w, err.Error(), http2.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	gzipBytes := util.GzipApi(resp)

	res := gjson.Get(string(gzipBytes), "result").String()
	sync := gjson.Get(res, "sync_info").String()
	latestHeight := gjson.Get(sync, "latest_block_height").Int()
	//latestHash := gjson.Get(sync, "latest_block_hash").String()
	//latestTime := gjson.Get(sync, "latest_block_time").String()

	// get gas fee
	jsonData = map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      "dontcare",
		"method":  "gas_price",
		"params":  []int{int(latestHeight)},
	}
	jsonStr, _ = json.Marshal(jsonData)
	r, err = http2.NewRequestWithContext(ctx, http2.MethodPost, cmd.NodeURL, bytes.NewReader(jsonStr))
	if err != nil {
		http2.Error(w, err.Error(), http2.StatusInternalServerError)
		return
	}
	r.Header.Add("Content-Type", "application/json; charset=utf-8")
	r.Header.Add("accept-encoding", "gzip,deflate")
	client = &http2.Client{Timeout: 5 * time.Second}
	resp, err = client.Do(r)
	if err != nil {
		http2.Error(w, err.Error(), http2.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	res = gjson.Get(string(util.GzipApi(resp)), "result").String()
	gas := gjson.Get(res, "gas_price").String()

	response := HTTP.ReportNodesStatusReply{
		Computation:       "1000",
		NumberOfMiners:    "100",
		Rewards:           "10",
		LatestBlockHeight: strconv.FormatInt(latestHeight, 10),
		GasFee:            gas,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http2.Error(w, err.Error(), http2.StatusInternalServerError)
	}

}

// ListAllChipsHTTPHandler show details of all chips
func (s *MinerStatusServiceHTTP) ListAllChipsHTTPHandler(w http.ResponseWriter, r *http.Request) {
	// method GET
	if r.Method != http2.MethodGet {
		http2.Error(w, "Method Not Allowed", http2.StatusMethodNotAllowed)
		return
	}
	// get params
	query := r.URL.Query()
	req := &HTTP.ListChipsRequest{
		Addr:      strings.Split(query.Get("url"), ","),
		SerialNum: query.Get("serialNum"),
		BusId:     query.Get("busId"),
	}

	clientDeadline := time.Now().Add(time.Duration(connect.Delay * time.Second))
	ctx, cancel := context.WithDeadline(context.Background(), clientDeadline)
	defer cancel()
	workers := &HTTP.ListWorkersReply{Workers: make([]HTTP.ListCards, 0)}
	for _, each := range req.Addr {

		cardLists := make([]*chipRPC.CardItem, 0)
		cards := make([]*HTTP.CardItem, 0)

		// connect to each worker
		conn, err := grpc.DialContext(ctx, each+":7001", grpc.WithInsecure())
		if err != nil {
			fmt.Println("Error connecting to RPC server:", err)
			workers.Workers = append(workers.Workers, HTTP.ListCards{
				TotalSize: 0,
				Addr:      each,
				Cards:     cards,
				Status:    "Disconnected",
			})
			workers.NumOfWorkers += 1
			continue
		}
		// Prepare the request
		request := &chipRPC.ChipsRequest{
			Url:       "http://119.120.92.239" + ":30345",
			SerialNum: req.SerialNum,
			BusId:     req.SerialNum,
		}
		client := chipRPC.NewChipServiceClient(conn)
		// Call the RPC method
		var response *chipRPC.ListChipsReply
		response, err = client.ListAllChips(ctx, request, grpc.WaitForReady(true))
		if err != nil {
			fmt.Println("Error query chip information:", err)
			workers.Workers = append(workers.Workers, HTTP.ListCards{
				TotalSize: 0,
				Addr:      each,
				Cards:     cards,
				Status:    "Disconnected",
			})
			workers.NumOfWorkers += 1
			continue
		}

		// deal with response
		cardLists = response.Cards
		listLen := 0
		conn.Close()
		for _, card := range cardLists {
			tpus := make([]*HTTP.ChipItem, 0)
			if req.SerialNum != "" && card.SerialNum != req.SerialNum {
				continue
			}
			// tpu chips
			for _, chip := range card.Chips {
				if req.BusId != "" && chip.BusId != req.BusId {
					continue
				}
				tpus = append(tpus, &HTTP.ChipItem{
					DevId:  chip.DevId,
					BusId:  chip.BusId,
					Memory: chip.Memory,
					Tpuuti: chip.Tpuuti,
					//BoardT:  chip.BoardT,
					ChipT:   chip.ChipT,
					TpuP:    chip.TpuP,
					TpuV:    chip.TpuV,
					TpuC:    chip.TpuC,
					Currclk: chip.Currclk,
					Status:  chip.Status,
				})
				listLen += 1
			}
			// all card infos
			// get claim status
			claimStatus := "unclaimed"
			// claimStatus := getStatusOfCard(card.SerialNum)
			cards = append(cards, &HTTP.CardItem{
				CardID:      card.CardID,
				Name:        card.Name,
				Mode:        card.Mode,
				SerialNum:   card.SerialNum,
				Atx:         card.Atx,
				MaxP:        card.MaxP,
				BoardP:      card.BoardP,
				BoardT:      card.BoardT,
				Minclk:      card.Minclk,
				Maxclk:      card.Maxclk,
				Chips:       tpus,
				ClaimStatus: claimStatus,
			})
		}

		workers.Workers = append(workers.Workers, HTTP.ListCards{
			TotalSize: int64(listLen),
			Addr:      each,
			Cards:     cards,
			Status:    "Connected",
		})
		workers.NumOfWorkers += 1

	}

	//if err != nil {
	//	http2.Error(w, err.Error(), http2.StatusInternalServerError)
	//	return
	//}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(workers); err != nil {
		http2.Error(w, err.Error(), http2.StatusInternalServerError)
	}

}

// StartChipCPUHandler start chip CPU before asking chips to perform their task
func (s *MinerStatusServiceHTTP) StartChipCPUHandler(w http.ResponseWriter, r *http.Request) {

	// method Post
	if r.Method != http2.MethodPost {
		http2.Error(w, "Method Not Allowed", http2.StatusMethodNotAllowed)
		return
	}
	// get params
	query := r.URL.Query()
	req := &HTTP.StartChipsRequest{
		Addr:  query.Get("url"),
		DevId: query.Get("devId"),
	}

	// connect to worker
	clientDeadline := time.Now().Add(time.Duration(15 * time.Second)) // for the first time to start a chip, it can take up to 10s
	ctx, cancel := context.WithDeadline(context.Background(), clientDeadline)
	defer cancel()
	conn, err := grpc.DialContext(ctx, req.Addr+":7001", grpc.WithInsecure())
	if err != nil {
		http2.Error(w, err.Error(), http2.StatusInternalServerError)
		return
	}
	// Prepare the request
	request := &chipRPC.ChipsRequest{
		DevId: req.DevId,
	}
	client := chipRPC.NewChipServiceClient(conn)
	// Call the RPC method
	var response *chipRPC.ChipStatusReply
	fmt.Println("worker", req.Addr, ": start chip", req.DevId)
	response, err = client.StartChipCPU(ctx, request, grpc.WaitForReady(true))
	if err != nil {
		http2.Error(w, err.Error(), http2.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http2.Error(w, err.Error(), http2.StatusInternalServerError)
	}

}

// ViewAccountHandler get the latest information about account of miner and balance
func (s *MinerStatusServiceHTTP) ViewAccountHandler(w http.ResponseWriter, r *http.Request) {

	// method Get
	if r.Method != http2.MethodGet {
		http2.Error(w, "Method Not Allowed", http2.StatusMethodNotAllowed)
		return
	}
	// get params
	query := r.URL.Query()
	accountId := query.Get("accountId")

	// get status
	jsonData := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      "czROwmnXE",
		"method":  "query",
		"params":  map[string]interface{}{"account_id": accountId, "finality": "final", "request_type": "view_account"},
	}
	jsonStr, _ := json.Marshal(jsonData)
	// POST request
	clientDeadline := time.Now().Add(time.Duration(connect.Delay * time.Second))
	ctx, cancel := context.WithDeadline(context.Background(), clientDeadline)
	defer cancel()
	r, err := http2.NewRequestWithContext(ctx, http2.MethodPost, cmd.NodeURL, bytes.NewReader(jsonStr))
	if err != nil {
		http2.Error(w, err.Error(), http2.StatusInternalServerError)
		return
	}
	r.Header.Add("Content-Type", "application/json; charset=utf-8")
	r.Header.Add("accept-encoding", "gzip,deflate")

	client := &http2.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(r)
	if err != nil {
		http2.Error(w, err.Error(), http2.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	gzipBytes := util.GzipApi(resp)
	res := gjson.Get(string(gzipBytes), "result").String()
	power := gjson.Get(res, "power").Int() / (1e12)
	balance := gjson.Get(res, "amount").String()

	num := new(big.Float)
	num.SetString(balance)
	divisor := new(big.Int)
	divisor.SetString("1000000000000000000000000", 10)
	amount := new(big.Float).Quo(num, new(big.Float).SetInt(divisor))
	amountString := amount.Text('f', 9)

	if gjson.Get(string(gzipBytes), "error").String() != "" {
		amountString = "--"
	}

	response := HTTP.ViewAccountReply{
		Total:   amountString,
		Rewards: "1",
		Slashed: "-1",
		Power:   strconv.FormatInt(power, 10),
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http2.Error(w, err.Error(), http2.StatusInternalServerError)
	}

}
