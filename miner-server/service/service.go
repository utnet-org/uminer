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
	// http miner service for miner UI software call
	MinerLoginServiceH     types.MinerLoginServiceHTTP
	MinerStatusServiceH    types.MinerStatusServiceHTTP
	MinerContainerServiceH types.ContainerServiceHTTP
	// grpc Chip service for calling chip using bm-sophon driver: start/burn/generate key pairs/sign message
	ChipServiceG chipApi.ChipServiceServer
	// grpc chain service for communicating with blockchain nodes, asking for status, get informed of burst block or get transaction to pack on blocks
	ChainService chainApi.ChainServiceServer
	// grpc image service for communicating with container cloud server based on k8s framework, miner will manage or coordinate AI task based on images provided by users
	ImageService containerApi.ImageServiceServer
	// grpc notebook service for communicating with container cloud server based on k8s framework, miner will manage or offer notebooks to users, who can do AI coding online conveniently
	NotebookService containerApi.NotebookServiceServer
}

func NewMinerService(ctx context.Context, conf *serverConf.Bootstrap, logger log.Logger, data *data.Data) (*Service, error) {
	var err error
	service := &Service{}

	if err != nil {
		return nil, err
	}

	service.MinerLoginServiceH = *types.NewMinerLoginServiceHTTP(conf, logger, data)
	service.MinerStatusServiceH = *types.NewMinerStatusServiceHTTP(conf, logger, data)
	service.MinerContainerServiceH = *types.NewMinerContainerServiceHTTP(conf, logger, data)
	service.ChainService = types.NewChainService(conf, logger, data)
	service.ImageService = types.NewImageService(conf, logger, data)
	service.NotebookService = types.NewNotebookService(conf, logger, data)

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
