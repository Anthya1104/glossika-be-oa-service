package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var Client *redis.Client

func InitRedis(addr, password string, db int) {
	Client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
}

func Ping() error {
	return Client.Ping(context.Background()).Err()
}
