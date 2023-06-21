package utils

import "github.com/redis/go-redis/v9"

var RDB *redis.Client

func InitRedis() {
	RDB = redis.NewClient(
		&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		},
	)
}
