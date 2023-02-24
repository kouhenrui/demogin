package global

import (
	"context"
	"github.com/go-redis/redis/v8"
	"gopkg.in/ini.v1"

	"log"
)

type RedisClient struct {
	Client *redis.Client
}

var (
	Redis *redis.Client
)

func init() {
	Cfg, _ := ini.Load("conf/conf.ini")
	var (
		reAddr          = Cfg.Section("redis").Key("address").String()
		rePwd           = Cfg.Section("mysql").Key("passWord").String()
		reName          = Cfg.Section("mysql").Key("username").String()
		reDb, _         = Cfg.Section("mysql").Key("db").Int()
		rePool, _       = Cfg.Section("mysql").Key("poolSize").Int()
		reMaxRetries, _ = Cfg.Section("mysql").Key("maxRetries").Int()
	)
	Redis = redis.NewClient(&redis.Options{
		Addr:       reAddr,
		Username:   reName,
		Password:   rePwd,
		DB:         reDb,
		PoolSize:   rePool,
		MaxRetries: reMaxRetries,
	})

	_, err := Redis.Ping(context.Background()).Result()
	if err != nil {
		log.Printf("redis connect get failed.%v", err)
		return
	}
	log.Printf("redis init success")
}
