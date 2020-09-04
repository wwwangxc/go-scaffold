package xredis

import "github.com/go-redis/redis"

var (
	// Cli xredis client
	Cli *Client
)

func initialize(conf *Config) {
	Cli = newClient(conf)
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
		cli: client,
	}
}
