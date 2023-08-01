package db

import (
	"context"
	"sync"

	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client
var redisOnce sync.Once

// RedisConnect returns a redis client
func RedisConnect() *redis.Client {

	redisOnce.Do(func() {
		rdb = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})
	})

	return rdb
}

// RedisDisconnect disconnects from redis
func RedisDisconnect() {
	rdb.Close()
}

// RedisPing pings redis
func RedisPing() (string, error) {
	return rdb.Ping(context.TODO()).Result()
}

// RedisSet sets a key-value pair in redis
func RedisSet(key string, value interface{}) error {
	return rdb.Set(context.TODO(), key, value, 0).Err()
}

// RedisGet gets a value from redis
func RedisGet(key string) (string, error) {
	return rdb.Get(context.TODO(), key).Result()
}
