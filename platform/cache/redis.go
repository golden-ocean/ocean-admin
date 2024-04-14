package cache

import (
	"log"

	"github.com/gofiber/storage/redis/v3"
	"github.com/golden-ocean/ocean-admin/pkg/utils"
)

func RedisConnection() (*redis.Storage, error) {

	redisConnURL, err := utils.ConnectionURLBuilder("redis")
	if err != nil {
		return nil, err
	}

	store := redis.New(redis.Config{
		URL:   redisConnURL,
		Reset: false,
	})
	return store, nil
}

func OpenRedisConnection() *redis.Storage {
	rdb, err := RedisConnection()
	if err != nil {
		log.Fatal("Redis数据库连接出错!", err)
	}
	return rdb
}
