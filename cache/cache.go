package cache

import (
	"context"
	"strconv"

	"github.com/ChenMiaoQiu/go-cloud-disk/conf"
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

// init redis
func Redis() {
	db, _ := strconv.Atoi(conf.RedisDB)
	client := redis.NewClient(&redis.Options{
		Addr:       conf.RedisAddr,
		Password:   conf.RedisPassword,
		DB:         db,
		MaxRetries: 1,
	})

	_, err := client.Ping(context.Background()).Result()

	if err != nil {
		panic("can't connect redis")
	}

	RedisClient = client
}
