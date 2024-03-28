package rdsdrv

import "github.com/redis/go-redis/v9"

func MustLoad() *redis.Client {
	opt, err := redis.ParseURL("redis://localhost:6379")
	if err != nil {
		panic(err)
	}

	return redis.NewClient(opt)
}
