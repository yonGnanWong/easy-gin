package App

import "github.com/go-redis/redis"

var Redis RedisParam

type RedisParam struct {
	Addr string
	Password string
	Database int
	Redis *redis.Client
}
