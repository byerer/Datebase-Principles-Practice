package global

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var (
	RedisClient *redis.Client
	Ctx         = context.Background()
)

func NewRedisClient() *redis.Client {
	addr := fmt.Sprintf("%s:%s", RedisSetting.IP, RedisSetting.Port)
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
