package service

import (
	"context"
	api "uminer/base-server/api/v1"
	"uminer/base-server/internal/conf"
	"uminer/base-server/internal/data"
	"uminer/base-server/internal/service/user"
	"uminer/common/log"
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
