package types

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kratos/kratos/v2/transport/http"
	"google.golang.org/grpc"
	http2 "net/http"
	"strings"
	"uminer/common/log"
	"uminer/miner-server/api/chipApi/HTTP"
	chipRPC "uminer/miner-server/api/chipApi/rpc"
	"uminer/miner-server/data"
	"uminer/miner-server/serverConf"
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

	workers := &HTTP.ListWorkersReply{Workers: make([]HTTP.ListCards, 0)}
	for _, each := range req.Addr {

		cardLists := make([]*chipRPC.CardItem, 0)
		cards := make([]*HTTP.CardItem, 0)

		// connect to each worker
		conn, err := grpc.Dial(each+":7001", grpc.WithInsecure())
		if err != nil {
			fmt.Println("Error connecting to RPC server:", err)
			workers.Workers = append(workers.Workers, HTTP.ListCards{
				TotalSize: 0,
				Addr:      each,
				Cards:     cards,
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
		response, err = client.ListAllChips(context.Background(), request, grpc.WaitForReady(true))
		if err != nil {
			fmt.Println("Error query chip information:", err)
			workers.Workers = append(workers.Workers, HTTP.ListCards{
				TotalSize: 0,
				Addr:      each,
				Cards:     cards,
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
	conn, err := grpc.Dial(req.Addr+":7001", grpc.WithInsecure())
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
	response, err = client.StartChipCPU(context.Background(), request, grpc.WaitForReady(true))
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http2.Error(w, err.Error(), http2.StatusInternalServerError)
	}

}
