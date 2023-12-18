package service

import (
	"context"
	"uminer/common/log"
	api "uminer/miner_server/base-server/api/v1"
	"uminer/miner_server/base-server/internal/conf"
	"uminer/miner_server/base-server/internal/data"
	"uminer/miner_server/base-server/internal/service/user"
)

type Service struct {
	UserService api.UserServiceServer
}

func NewService(ctx context.Context, conf *conf.Bootstrap, logger log.Logger, data *data.Data) (*Service, error) {
	var err error
	service := &Service{}

	if err != nil {
		return nil, err
	}

	service.UserService = user.NewUserService(conf, logger, data)

	return service, nil
}
