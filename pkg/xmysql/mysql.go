package xmysql

import (
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	// Cli mysql client
	Cli *sqlx.DB

	_initOnce  sync.Once
	_closeOnce sync.Once
)

func initialize(conf *Config) {
	_initOnce.Do(func() {
		Cli = newClient(conf)
	})
}

func newClient(conf *Config) *sqlx.DB {
	cli, err := sqlx.Connect("mysql", conf.DSN)
	if err != nil {
		panic(err.Error())
	}
	return cli
}

// Close ..
func Close() error {
	if Cli == nil {
		return nil
	}
	var err error
	_closeOnce.Do(func() {
		err = Cli.Close()
	})
	return err
}
