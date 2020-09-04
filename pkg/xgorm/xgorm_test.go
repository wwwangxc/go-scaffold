package xgorm

import (
	"testing"
)

func TestGorm(t *testing.T) {
	conf := &Config{
		DSN: "root:root@tcp(127.0.0.1:3306)/gateway?charset=utf8&parseTime=True&loc=Local",
	}
	cli := conf.Build()
	defer cli.Close()
}
