package database

import "time"

type User struct {
	ID        int64  `gorm:"primary_key"`
	Username  string `gorm:"type:varchar(20);not null" binding:"required"`
	Password  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
