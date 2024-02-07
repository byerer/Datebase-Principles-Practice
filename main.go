package main

import (
	"GradingSystem/global"
	"GradingSystem/internal/router"
	"GradingSystem/pkg/setting"
	"fmt"
	"time"
)

func init() {
	global.InitLogger()
	defer global.SugarLogger.Sync()
	err := setupSetting()
	if err != nil {
		global.SugarLogger.Fatalf("init.setupSetting err: %v", err)
	}
	err = setupDBEngine()
	if err != nil {
		global.SugarLogger.Fatalf("init.setupDBEngine err: %v", err)
	}
	err = setupSMTP()
	if err != nil {
		global.SugarLogger.Fatalf("init.setupSMTP err: %v", err)
	}
	err = setupRedis()
	if err != nil {
		global.SugarLogger.Fatalf("init.setupRedis err: %v", err)
	}
	global.SugarLogger.Info("init setting success")
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
