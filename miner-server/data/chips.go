package data

import (
	"context"
	"uminer/common/log"
	"uminer/miner-server/api/chipApi"

	"gorm.io/gorm"
)

// SelectChipList is for condition selection
type SelectChipList struct {
	devId string
}

type ChipsMethod interface {
	List(ctx context.Context, condition *SelectChipList) ([]*chipApi.TPUCards, error)
}

type chipMethod struct {
	log *log.Helper
	db  *gorm.DB
}

func NewChipMethod(db *gorm.DB, logger log.Logger) ChipsMethod {
	return &chipMethod{
		log: log.NewHelper("ChipsMethod", logger),
		db:  db,
	}
}

func (d *chipMethod) List(ctx context.Context, condition *SelectChipList) ([]*chipApi.TPUCards, error) {
	//db := d.db

	cards := make([]*chipApi.TPUCards, 0)

	return cards, nil
}
