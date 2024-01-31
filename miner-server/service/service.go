package service

import (
	"context"
	"uminer/common/log"
	chainApi "uminer/miner-server/api/chainApi/rpc"
	chipApi "uminer/miner-server/api/chipApi/rpc"
	"uminer/miner-server/api/containerApi"
	"uminer/miner-server/data"
	"uminer/miner-server/serverConf"
	"uminer/miner-server/service/types"
)

type Service struct {
	ChipServiceH    types.ChipServiceHTTP
	ChipServiceG    chipApi.ChipServiceServer
	ChainService    chainApi.ChainServiceServer
	ImageService    containerApi.ImageServiceServer
	NotebookService containerApi.NotebookServiceServer
}

func NewMinerService(ctx context.Context, conf *serverConf.Bootstrap, logger log.Logger, data *data.Data) (*Service, error) {
	var err error
	service := &Service{}

	if err != nil {
		return nil, err
	}

	service.ChipServiceH = *types.NewChipServiceHTTP(conf, logger, data)
	service.ChainService = types.NewChainService(conf, logger, data)
	service.ImageService = types.NewImageService(conf, logger, data)
	service.NotebookService = types.NewRentalService(conf, logger, data)

	return service, nil
}

func NewWorkerService(ctx context.Context, conf *serverConf.Bootstrap, logger log.Logger, data *data.Data) (*Service, error) {
	var err error
	service := &Service{}

	if err != nil {
		return nil, err
	}

	service.ChipServiceG = types.NewChipService(conf, logger, data)

	return service, nil
}
