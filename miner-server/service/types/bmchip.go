package types

import (
	"context"
	"uminer/common/log"
	"uminer/miner-server/chipApi"
	"uminer/miner-server/data"
	"uminer/miner-server/serverConf"
)

type ChipService struct {
	chipApi.UnimplementedChipServiceServer
	conf *serverConf.Bootstrap
	log  *log.Helper
	data *data.Data
}

func NewChipService(conf *serverConf.Bootstrap, logger log.Logger, data *data.Data) chipApi.ChipServiceServer {
	return &ChipService{
		conf: conf,
		log:  log.NewHelper("ChipService", logger),
		data: data,
	}
}

func (s *ChipService) ListAllChips(ctx context.Context, req *chipApi.ListChipsRequest) (*chipApi.ListChipsReply, error) {

	cards := make([]*chipApi.CardItem, 0)

	return &chipApi.ListChipsReply{
		TotalSize: 0,
		Cards:     cards,
	}, nil
}
