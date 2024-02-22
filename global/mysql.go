package global

import (
	"GradingSystem/internal/model/database"
	"GradingSystem/pkg/setting"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	DB *gorm.DB
)

func NewDBEngine(databaseSetting *setting.MySQLSettingS) (*gorm.DB, error) {
	s := "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(s,
		databaseSetting.User,
		databaseSetting.Password,
		databaseSetting.IP,
		databaseSetting.Port,
		databaseSetting.Database,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(
		database.User{},
		database.Teacher{},
		database.Student{},
		database.Essay{},
	)
	return db, nil
}
