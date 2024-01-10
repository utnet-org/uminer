package service

import (
	"context"
	"uminer/common/log"
	chainApi "uminer/miner-server/chainApi/rpc"
	chipApi "uminer/miner-server/chipApi/rpc"
	"uminer/miner-server/data"
	"uminer/miner-server/serverConf"
	"uminer/miner-server/service/types"
)

type Service struct {
	ChipService  chipApi.ChipServiceServer
	ChainService chainApi.ChainServiceServer
}

func NewService(ctx context.Context, conf *serverConf.Bootstrap, logger log.Logger, data *data.Data) (*Service, error) {
	var err error
	service := &Service{}

	if err != nil {
		return nil, err
	}

	service.ChipService = types.NewChipService(conf, logger, data)
	service.ChainService = types.NewChainService(conf, logger, data)

	return service, nil
}
