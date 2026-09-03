package redis

import "github.com/redis/go-redis/v9"

func NewRedisClient(conn string, db int) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: conn,
		DB:   db,
	})
}
