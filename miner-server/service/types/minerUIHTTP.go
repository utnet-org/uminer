package types

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/tidwall/gjson"
	"google.golang.org/grpc"
	http2 "net/http"
	"strconv"
	"strings"
	"time"
	"uminer/common/log"
	"uminer/miner-server/api/HTTP"
	chipRPC "uminer/miner-server/api/chipApi/rpc"
	"uminer/miner-server/data"
	"uminer/miner-server/serverConf"
	"uminer/miner-server/util"
)

type MinerUIServiceHTTP struct {
	conf *serverConf.Bootstrap
	log  *log.Helper
	data *data.Data
}

func NewChipServiceHTTP(conf *serverConf.Bootstrap, logger log.Logger, data *data.Data) *MinerUIServiceHTTP {
	return &MinerUIServiceHTTP{
		conf: conf,
		log:  log.NewHelper("ChipService", logger),
		data: data,
	}
}

// MapWorkersAddrHandler get all worker address of a miner
func (s *MinerUIServiceHTTP) MapWorkersAddrHandler(w http.ResponseWriter, r *http.Request) {

	// method Post
	if r.Method != http2.MethodGet {
		http2.Error(w, "Method Not Allowed", http2.StatusMethodNotAllowed)
		return
	}
	// get params
	query := r.URL.Query()
	req := &HTTP.MapWorkersAddressRequest{
		MinerAddr: query.Get("minerAddr"),
	}

	// get mapping
	workers := make([]string, 0)
	workers = append(workers, "192.168.10.49")
	workers = append(workers, "192.168.10.50")
	response := HTTP.MapWorkersAddressReply{
		MinerAddr:  req.MinerAddr,
		WorkerAddr: workers,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http2.Error(w, err.Error(), http2.StatusInternalServerError)
	}

}

// GetNodesStatusHandler get the latest information about node and about miner himself as well
func (s *MinerUIServiceHTTP) GetNodesStatusHandler(w http.ResponseWriter, r *http.Request) {

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
	clientDeadline := time.Now().Add(time.Duration(delay * time.Second))
	ctx, cancel := context.WithDeadline(context.Background(), clientDeadline)
	defer cancel()
	r, err := http2.NewRequestWithContext(ctx, http2.MethodPost, nodeURL, bytes.NewReader(jsonStr))
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
	r, err = http2.NewRequestWithContext(ctx, http2.MethodPost, nodeURL, bytes.NewReader(jsonStr))
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
		MyComputation:     "10",
		MyRewards:         "0.1",
		MyBlocks:          "1",
		MyWorkerNum:       "1",
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http2.Error(w, err.Error(), http2.StatusInternalServerError)
	}

}

// ListAllChipsHTTPHandler show details of all chips
func (s *MinerUIServiceHTTP) ListAllChipsHTTPHandler(w http.ResponseWriter, r *http.Request) {
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

	clientDeadline := time.Now().Add(time.Duration(delay * time.Second))
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
func (s *MinerUIServiceHTTP) StartChipCPUHandler(w http.ResponseWriter, r *http.Request) {

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
