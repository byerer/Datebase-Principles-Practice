package database

import "time"

type User struct {
	ID        int64  `gorm:"primary_key"`
	Username  string `gorm:"type:varchar(20);not null" binding:"required"`
	Password  string `gorm:"type:varchar(60);not null" binding:"required"`
	Email     string `gorm:"type:varchar(100);not null" binding:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
