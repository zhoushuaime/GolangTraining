package client

import (

	"github.com/go-redis/redis"
)

var clt *redis.Client

func init() {
	//conf := config.NewConfig()
	clt = redis.NewClient(&redis.Options{
		Addr:     "",
		Password: "", // no password set
		DB:       0,                  // use default DB
	})
}

// NewClient 获取redis客户端
func NewClient() *redis.Client {
	return clt
}
