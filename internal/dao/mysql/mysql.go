package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMYSQL() (*gorm.DB, error) {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "user:123456@tcp(127.0.0.1:3306)/GradingSystem?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}