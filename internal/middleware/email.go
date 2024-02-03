package middleware

import (
	"GradingSystem/global"
	"golang.org/x/exp/rand"
	"gopkg.in/gomail.v2"
)

func SendCode(email string, code string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", global.SMTPSetting.User)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello, your certification code is "+code)

	d := gomail.NewDialer(global.SMTPSetting.Host, global.SMTPSetting.Port, global.SMTPSetting.User, global.SMTPSetting.Password)

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

// GenerateCode 生成6位验证码
func GenerateCode() string {
	const number = "0123456789"
	rand.Seed(1)
	code := make([]byte, 6)
	for i := range code {
		code[i] = number[rand.Intn(len(number))]
	}
	return string(code)
}
