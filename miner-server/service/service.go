package service

import (
	"context"
	"uminer/common/log"
	api "uminer/miner-server/chipApi"
	"uminer/miner-server/data"
	"uminer/miner-server/serverConf"
	"uminer/miner-server/service/types"
)

type Service struct {
	ChipService api.ChipServiceServer
}

func NewService(ctx context.Context, conf *serverConf.Bootstrap, logger log.Logger, data *data.Data) (*Service, error) {
	var err error
	service := &Service{}

	if err != nil {
		return nil, err
	}

	service.ChipService = types.NewChipService(conf, logger, data)

	return service, nil
}
