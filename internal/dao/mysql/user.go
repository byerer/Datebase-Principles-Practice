package mysql

import (
	"GradingSystem/global"
	"GradingSystem/internal/model/database"
)

func InsertUser(user database.User) error {
	result := global.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
