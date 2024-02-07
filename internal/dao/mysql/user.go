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

func FindUserByName(username string) bool {
	_, err := GetUserByUsername(username)
	if err != nil {
		return false
	}
	return true
}

func GetUserByUsername(username string) (user database.User, err error) {
	result := global.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}
