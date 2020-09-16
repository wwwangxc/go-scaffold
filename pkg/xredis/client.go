package xredis

import (
	"sync"
	"time"

	"github.com/go-redis/redis"
)

// Client ..
type Client struct {
	cli    *redis.Client
	Closed bool
	mu     sync.Mutex
}

func newClient(conf *Config) *Client {
	client := redis.NewClient(&redis.Options{
		Addr:       conf.Addr,
		Password:   conf.Password,
		DB:         conf.DB,
		MaxRetries: conf.MaxRetries,
	})
	if _, err := client.Ping().Result(); err != nil {
		panic(err.Error())
	}
	return &Client{
		cli:    client,
		Closed: false,
	}
}

// Close ..
func (t *Client) Close() error {
	if t.cli == nil || t.Closed {
		return nil
	}
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.Closed {
		return nil
	}
	err := t.cli.Close()
	if err != nil {
		return nil
	}
	t.Closed = true
	return nil
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

// Redis `LLEN key` command.
func (t *Client) LLen(key string) int64 {
	len, err := t.cli.LLen(key).Result()
	if err != nil {
		return 0
	}
	return len
}

// Redis `LPOP key` command.
func (t *Client) LPop(key string) String {
	v, err := t.cli.LPop(key).Result()
	if err != nil {
		return ""
	}
	return String(v)
}

// Redis `LPUSH key value [value ...]` command.
func (t *Client) LPush(key string, values ...interface{}) (bool, error) {
	v, err := t.cli.LPush(key, values...).Result()
	if err != nil {
		return false, err
	}
	return int(v) == len(values), err
}

// Redis `RPOP key` command.
func (t *Client) RPop(key string) String {
	v, err := t.cli.RPop(key).Result()
	if err != nil {
		return ""
	}
	return String(v)
}

// Redis `RPUSH key value [value ...]` command.
func (t *Client) RPush(key string, values ...interface{}) (bool, error) {
	v, err := t.cli.RPush(key, values...).Result()
	if err != nil {
		return false, err
	}
	return int(v) == len(values), err
}

// -------------------------------------------------------------------------------- Set Commands

// Redis `SADD key member [member ...]` command.
func (t *Client) SAdd(key string, members ...interface{}) (int64, error) {
	if len(members) == 0 {
		return 0, nil
	}
	return t.cli.SAdd(key, members...).Result()
}

// Redis `SPOP key` command.
func (t *Client) SPop(key string) String {
	v, err := t.cli.SPop(key).Result()
	if err != nil {
		return ""
	}
	return String(v)
}

// Redis `SISMEMBER key member` command.
func (t *Client) SIsMember(key string, member interface{}) (bool, error) {
	return t.cli.SIsMember(key, member).Result()
}

// -------------------------------------------------------------------------------- Sorted Set Commands

// Redis `ZADD key score member [[score member] [score member] ...]` command.
func (t *Client) ZAdd(key string, members ...redis.Z) (int64, error) {
	if len(members) == 0 {
		return 0, nil
	}
	return t.cli.ZAdd(key, members...).Result()
}

// -------------------------------------------------------------------------------- Global Commands

// Redis `EXISTS` command.
func (t *Client) Exists(key string) bool {
	v, err := t.cli.Exists(key).Result()
	if err != nil {
		return false
	}
	return v == 1
}

// Redis `DEL` command.
func (t *Client) Del(key ...string) int64 {
	ret, err := t.cli.Del(key...).Result()
	if err != nil {
		return 0
	}
	return ret
}

// Redis `TTL key` command.
func (t *Client) TTL(key string) (int64, error) {
	v, err := t.cli.TTL(key).Result()
	if err != nil {
		return 0, err
	}
	return int64(v.Seconds()), nil
}
