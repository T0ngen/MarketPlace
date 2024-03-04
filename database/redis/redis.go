package redis


import "github.com/redis/go-redis/v9"
 



func OpenRedis() *redis.Client {
    redisClient := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // если есть пароль, укажите его здесь
        DB:       0,
       })
    return redisClient
}