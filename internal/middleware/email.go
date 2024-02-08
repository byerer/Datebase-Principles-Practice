package middleware

import (
	"GradingSystem/global"
	"GradingSystem/internal/dao/redis"
	"GradingSystem/internal/model/api"
	"golang.org/x/exp/rand"
	"gopkg.in/gomail.v2"
	"time"
)

func SendCode(email string, code string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", global.SMTPSetting.User)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello, your certification code is "+code)
	global.SugarLogger.Infof("email: %s, code: %s", email, code)
	d := gomail.NewDialer(global.SMTPSetting.Host, global.SMTPSetting.Port, global.SMTPSetting.User, global.SMTPSetting.Password)
	err := redis.SetCode(email, code)
	if err != nil {
		return err
	}
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

// GenerateCode 生成6位验证码
func GenerateCode() string {
	const number = "0123456789"
	rand.Seed(uint64(time.Now().Unix()))
	code := make([]byte, 6)
	for i := range code {
		code[i] = number[rand.Intn(len(number))]
	}
	return string(code)
}

func ValidateCode(emailInfo api.EmailInfo) bool {
	global.SugarLogger.Infof("email: %s, code: %s", emailInfo.Email, emailInfo.Code)
	if redisCode, err := redis.GetCode(emailInfo.Email); err != nil || redisCode != emailInfo.Code {
		return false
	}
	err := redis.DelCode(emailInfo.Email)
	if err != nil {
		global.SugarLogger.Errorf("delete code failed: %v", err)
	}
	return true
}
