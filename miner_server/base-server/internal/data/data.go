package data

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"uminer/common/log"
	"uminer/miner_server/base-server/internal/conf"
	"uminer/miner_server/base-server/internal/data/model"
	"uminer/miner_server/base-server/internal/data/model/fun"
)

type Data struct {
	UserM model.UserMethod
}

func NewData(bc *conf.Bootstrap, logger log.Logger) (*Data, func(), error) {
	confData := bc.Data
	d := &Data{}

	db, err := dbInit(confData)
	if err != nil {
		return nil, nil, err
	}

	d.UserM = model.NewUserMethod(db, logger)

	return d, func() {
	}, nil
}

func dbInit(confData *conf.Data) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(confData.Database.Source), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   log.DefaultGormLogger,
	})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&fun.User{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&fun.UserConfig{})
	if err != nil {
		return nil, err
	}

	return db, err
}
