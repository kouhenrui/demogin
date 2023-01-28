package global

import (
	"context"
	"github.com/go-redis/redis/v8"

	"log"
)

type RedisClient struct {
	Client *redis.Client
}

var (
	Redis *redis.Client
)

func init() {
	//log.Println("redis connecting")
	Redis = redis.NewClient(&redis.Options{
		Addr:     RedisConfig.address,
		Password: RedisConfig.password,
		DB:       RedisConfig.db,
		PoolSize: RedisConfig.poolSize,
	})

	_, err := Redis.Ping(context.Background()).Result()
	if err != nil {
		log.Printf("redis connect get failed.%v", err)
		return
	}
	log.Printf("redis init success")
}
