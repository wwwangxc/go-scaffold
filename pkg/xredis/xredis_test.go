package xredis

import (
	"testing"
)

func TestRedis(t *testing.T) {
	conf := &Config{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	}
	conf.Init()
}
