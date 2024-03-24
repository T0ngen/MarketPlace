package redis

import (
	"marketplace/pkg/common/config"

	"github.com/redis/go-redis/v9"
)
 



func OpenRedis(conf config.Config) *redis.Client {
    redisClient := redis.NewClient(&redis.Options{
        Addr:     conf.RedisString,
        Password: conf.RedisPassword, 
        DB:       conf.RedisDB,
       })
    return redisClient
}