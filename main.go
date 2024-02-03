package main

import (
	"GradingSystem/global"
	"GradingSystem/internal/router"
	"GradingSystem/pkg/setting"
	"fmt"
	"log"
	"time"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
	err = setupSMTP()
	if err != nil {
		log.Fatalf("init.setupSMTP err: %v", err)
	}
	err = setupRedis()
	if err != nil {
		log.Fatalf("init.setupRedis err: %v", err)
	}

	log.Println("init success")
}

func main() {
	router.Run()
}

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}

	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("MySQL", &global.MySQLSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Redis", &global.RedisSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

func setupDBEngine() error {
	var err error
	global.DB, err = global.NewDBEngine(global.MySQLSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupSMTP() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("SMTP", &global.SMTPSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupRedis() error {
	global.RedisClient = global.NewRedisClient()
	s, err := global.RedisClient.Ping(global.Ctx).Result()
	fmt.Println(s)
	if err != nil {
		return err
	}
	return nil
}
