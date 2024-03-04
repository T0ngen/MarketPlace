package redis


import "github.com/gomodule/redigo/redis"
 
func OpenRedis() (redis.Conn, error) {
    c, err := redis.DialURL("redis://localhost:6379/0")
    if err != nil {
        return nil, err
    }
    return c, nil
}
