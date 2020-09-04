package xgorm

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	// Cli mysql client
	Cli *gorm.DB
)

func initialize(conf *Config) {
	Cli = newClient(conf)
}

func newClient(conf *Config) *gorm.DB {
	client, err := gorm.Open("mysql", conf.DSN)
	if err != nil {
		panic(err.Error())
	}
	if err = client.DB().Ping(); err != nil {
		panic(err.Error())
	}
	return client
}
