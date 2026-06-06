package redisstore

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var (
	Ctx    = context.Background()
	Client *redis.Client
)

func Init() {

	Client = redis.NewClient(
		&redis.Options{
			Addr: "localhost:6379",
		},
	)

	_, err := Client.Ping(Ctx).Result()

	if err != nil {
		panic(err)
	}
}
