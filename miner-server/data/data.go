package data

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"uminer/common/log"
	"uminer/miner-server/serverConf"
)

type Data struct {
	ChipsM ChipsMethod
}

func NewData(bc *serverConf.Bootstrap, logger log.Logger) (*Data, func(), error) {
	confData := bc.Data
	d := &Data{}

	db, err := dbInit(confData)
	if err != nil {
		return nil, nil, err
	}

	d.ChipsM = NewChipMethod(db, logger)

	return d, func() {
	}, nil
}

func dbInit(confData *serverConf.Data) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(confData.Database.Source), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   log.DefaultGormLogger,
	})
	if err != nil {
		return nil, err
	}

	//err = db.AutoMigrate(&fun.User{})
	//if err != nil {
	//	return nil, err
	//}

	return db, err
}
