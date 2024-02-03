package redis

import (
	"GradingSystem/global"
	"time"
)

func SetCode(email string, code string) error {
	result := global.RedisClient.Set(global.Ctx, email, code, 5*time.Minute)
	if result.Err() != nil {
		return result.Err()
	}
	return nil
}
