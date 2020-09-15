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

func TestSet(t *testing.T) {
	cli := getClient()
	defer cli.Close()
	t.Log(cli.Set("k1", "v1", 0))
	t.Log(cli.Set("k2", "v2", 0))
	t.Log(cli.Set("k3", "v3", 0))
}

func TestGet(t *testing.T) {
	cli := getClient()
	defer cli.Close()
	t.Log(cli.Get("k1"))
	t.Log(cli.Get("k2"))
	t.Log(cli.Get("k3"))
	t.Log(cli.Get("k4"))
}

func TestHGet(t *testing.T) {
	cli := getClient()
	defer cli.Close()
	t.Log(cli.HGet("hKey", "field1"))
	t.Log(cli.HGet("key", "field1"))
	t.Log(cli.HGet("hKey", "field4"))
}

func TestHMSet(t *testing.T) {
	cli := getClient()
	defer cli.Close()
	arg := map[string]interface{}{
		"f1": "v1",
		"f2": "v2",
		"f3": "v3",
	}
	t.Log(cli.HMSet("hKey", arg, 0))
}

func TestHMGet(t *testing.T) {
	cli := getClient()
	defer cli.Close()
	hashMap := cli.HMGet("hKey", []string{"f1", "f2", "field3", "field4"})
	for k, v := range hashMap {
		t.Log(k, "：", v)
	}
}

func TestHGetALL(t *testing.T) {
	cli := getClient()
	defer cli.Close()
	hashMap := cli.HGetAll("hKey")
	for k, v := range hashMap {
		t.Log(k, "：", v)
	}
	hashMap2 := cli.HGetAll("hkey")
	for k, v := range hashMap2 {
		t.Log(k, "：", v)
	}
}

func TestHDel(t *testing.T) {
	cli := getClient()
	defer cli.Close()
	t.Log(cli.HDel("hKey", []string{"f1", "f2", "f3", "f4"}))
}

func TestDel(t *testing.T) {
	cli := getClient()
	defer cli.Close()
	t.Log(cli.Del("hKey"))
}

func getClient() *Client {
	conf := &Config{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	}
	return conf.Build()
}
