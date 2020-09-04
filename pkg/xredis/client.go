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

/////////////////////////////////////////////////////////////////////// Redis Commands String

// Get ->
// Redis `GET key` command.
func (t *Client) Get(key string) String {
	v, err := t.cli.Get(key).Result()
	if err != nil {
		return ""
	}
	return String(v)
}

// Set ->
// Redis `SET key value [expiration]` command.
func (t *Client) Set(key string, value interface{}, expire time.Duration) bool {
	return t.cli.Set(key, value, expire).Err() == redis.Nil
}

/////////////////////////////////////////////////////////////////////// Redis Commands Hash

// HMSet ->
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

// HGetAll ->
// Redis `HGETALL key` command.
func (t *Client) HGetAll(key string) map[string]string {
	return t.cli.HGetAll(key).Val()
}

// Del ->
// Redis `DEL` command.
func (t *Client) Del(key ...string) int64 {
	ret, err := t.cli.Del(key...).Result()
	if err != nil {
		return 0
	}
	return ret
}
