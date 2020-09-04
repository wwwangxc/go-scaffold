package xmysql

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	conf := &Config{
		DSN: "root:root@tcp(127.0.0.1:3306)/gateway?charset=utf8&parseTime=True&loc=Local",
	}
	conf.Init()
	defer Close()
	fmt.Println(Cli.Ping())
}

func TestBuild(t *testing.T) {
	conf := &Config{
		DSN: "root:root@tcp(127.0.0.1:3306)/gateway?charset=utf8&parseTime=True&loc=Local",
	}
	cli := conf.Build()
	defer cli.Close()
	fmt.Println(cli.Ping())
}
