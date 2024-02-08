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

func UpdatePassword(email string, password string) error {
	result := global.DB.Model(&database.User{}).Where("email = ?", email).Update("password", password)
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

func FindUserByEmail(email string) bool {
	_, err := GetUserByEmail(email)
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

func GetUserByEmail(email string) (user database.User, err error) {
	result := global.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}
