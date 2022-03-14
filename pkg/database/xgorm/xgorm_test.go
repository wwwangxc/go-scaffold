package xgorm

import (
	"testing"
	"time"
)

func TestBuild(t *testing.T) {
	conf := &Config{
		DSN:             "root:root@tcp(127.0.0.1:3306)/scaffold?charset=utf8&parseTime=True&loc=Local",
		ConnMaxLifetime: 300 * time.Second,
		MaxIdleConns:    50,
		MaxOpenConns:    100,
	}
	_ = conf.Build()
}

func TestAppend(t *testing.T) {
	conf := &Config{
		DSN:             "root:root@tcp(127.0.0.1:3306)/scaffold?charset=utf8&parseTime=True&loc=Local",
		ConnMaxLifetime: 300 * time.Second,
		MaxIdleConns:    50,
		MaxOpenConns:    100,
	}
	conf.Append("test")
}
