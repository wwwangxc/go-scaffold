package mysql

import (
	"fmt"
	"testing"
)

func TestBuild(t *testing.T) {
	conf := &Config{
		DSN: "root:root@tcp(127.0.0.1:3306)/gateway?charset=utf8&parseTime=True&loc=Local",
	}
	db := conf.Build()
	defer db.Close()
	fmt.Println(db.StoreName, db.Ping())
	_, _ = Append("dbname", db)
	fmt.Println(db.StoreName, db.Ping())
}

func TestAppend(t *testing.T) {
	conf := &Config{
		DSN: "root:root@tcp(127.0.0.1:3306)/gateway?charset=utf8&parseTime=True&loc=Local",
	}
	conf.Append("db0")
	defer Close("db0")
	fmt.Println(Store("db0").StoreName, Store("db0").Ping())
}
