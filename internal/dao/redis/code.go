package redis

import (
	"GradingSystem/global"
	"errors"
	"github.com/redis/go-redis/v9"
	"time"
)

func SetCode(email string, code string) error {
	result := global.RedisClient.Set(global.Ctx, email, code, 5*time.Minute)
	global.SugarLogger.Infof("email: %s, code: %s", email, code)
	if result.Err() != nil {
		return result.Err()
	}
	return nil
}

func GetCode(email string) (code string, err error) {
	val, err := global.RedisClient.Get(global.Ctx, email).Result()
	switch {
	case err == redis.Nil:
		return "", errors.New("code not found")
	case err != nil:
		return "", err
	default:
		return val, nil
	}

}

func DelCode(email string) error {
	result := global.RedisClient.Del(global.Ctx, email)
	if result.Err() != nil {
		return result.Err()
	}
	return nil
}
