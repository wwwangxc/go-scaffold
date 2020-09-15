package xredis

import (
	"time"

	"github.com/go-redis/redis"
)

// Client ..
type Client struct {
	cli *redis.Client
}

// Close ..
func (t *Client) Close() error {
	if t.cli == nil {
		return nil
	}
	return t.cli.Close()
}

// -------------------------------------------------------------------------------- String Commands

// Redis `GET key` command.
func (t *Client) Get(key string) String {
	v, err := t.cli.Get(key).Result()
	if err != nil {
		return ""
	}
	return String(v)
}

// Redis `SET key value [expiration]` command.
func (t *Client) Set(key string, value interface{}, expire time.Duration) bool {
	return t.cli.Set(key, value, expire).Err() == nil
}

// -------------------------------------------------------------------------------- Hash Commands

// Redis `HGET key field` command.
func (t *Client) HGet(key string, field string) String {
	v, err := t.cli.HGet(key, field).Result()
	if err != nil {
		return ""
	}
	return String(v)
}

// Redis `HMGET key field [field ...]` command.
func (t *Client) HMGet(key string, fields []string) map[string]String {
	if len(fields) == 0 {
		return nil
	}
	v, err := t.cli.HMGet(key, fields...).Result()
	if err != nil {
		return nil
	}
	retMap := make(map[string]String, len(fields))
	for index, value := range fields {
		tmp := v[index]
		if tmp != nil {
			retMap[value] = String(tmp.(string))
		} else {
			retMap[value] = ""
		}
	}
	return retMap
}

// Redis `HGETALL key` command.
func (t *Client) HGetAll(key string) map[string]String {
	v, err := t.cli.HGetAll(key).Result()
	if err != nil {
		return nil
	}
	retMap := make(map[string]String, len(v))
	for k, value := range v {
		retMap[k] = String(value)
	}
	return retMap
}

// Redis `HMSet` `Expire` command with pipeline.
func (t *Client) HMSet(key string, value map[string]interface{}, expire time.Duration) bool {
	if len(value) == 0 {
		return false
	}
	pipe := t.cli.Pipeline()
	if pipe.HMSet(key, value).Err() != nil {
		return false
	}
	if expire > 0 && pipe.Expire(key, expire).Err() != nil {
		return false
	}
	_, err := pipe.Exec()
	return err == nil
}

// Redis `HDEL key field [field ...]` command.
func (t *Client) HDel(key string, fields []string) int64 {
	if len(fields) == 0 {
		return 0
	}
	num, err := t.cli.HDel(key, fields...).Result()
	if err != nil {
		return 0
	}
	return num
}

// -------------------------------------------------------------------------------- List Commands

// -------------------------------------------------------------------------------- Set Commands

// -------------------------------------------------------------------------------- Sorted Set Commands

// -------------------------------------------------------------------------------- Global Commands

// Redis `DEL` command.
func (t *Client) Del(key ...string) int64 {
	ret, err := t.cli.Del(key...).Result()
	if err != nil {
		return 0
	}
	return ret
}
