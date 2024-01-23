package types

import (
	"context"
	"encoding/json"
	"uminer/common/log"
	"uminer/miner-server/api/rentalApi"
	"uminer/miner-server/data"
	"uminer/miner-server/serverConf"
)

type RentalService struct {
	rentalApi.UnimplementedRentalServiceServer
	conf *serverConf.Bootstrap
	log  *log.Helper
	data *data.Data
}

func NewRentalService(conf *serverConf.Bootstrap, logger log.Logger, data *data.Data) rentalApi.RentalServiceServer {
	return &RentalService{
		conf: conf,
		log:  log.NewHelper("RentalService", logger),
		data: data,
	}
}

// CreateNoteBook create notebook by miner after receiving request from renter
func (s *RentalService) CreateNoteBook(ctx context.Context, req *rentalApi.CreateNoteBookRequest) (*rentalApi.CreateNoteBookReply, error) {

	reply := &rentalApi.CreateNoteBookReply{
		Id:     "",
		Status: false,
	}

	// http addr request
	requestUrl := "https://console.utlab.io/openaiserver/v1/developmanage/notebook"
	jsonData := map[string]interface{}{
		"algorithmId":      req.AlgorithmId,
		"algorithmVersion": "V1",
		"datasetId":        "",
		"datasetVersion":   "",
		"desc":             req.Description,
		"imageId":          req.ImageId,
		"imageVersion":     "v1",
		"name":             req.Name,
		"resourcePool":     "common-pool",
		"resourceSpecId":   req.ResourceSpecId,
		"taskNumber":       1,
	}
	resp := Post(requestUrl, jsonData, "application/json", req.Token)

	var response struct {
		Success bool `json:"success"`
		Payload struct {
			ID string `json:"id"`
		} `json:"payload"`
		Error interface{} `json:"error"`
	}
	err := json.Unmarshal(resp, &response)
	if err != nil {
		return reply, err
	}

	reply.Id = response.Payload.ID
	reply.Status = true
	return reply, nil
}
