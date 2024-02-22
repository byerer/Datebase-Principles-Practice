package database

import "time"

type User struct {
	ID        int64  `gorm:"primary_key"`
	Username  string `gorm:"type:varchar(20);not null"`
	Password  string `gorm:"type:varchar(30);not null"`
	Email     string `gorm:"type:varchar(30);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Teacher struct {
	User
	TeacherName string `gorm:"type:varchar(20);not null"`
}

type Student struct {
	User
	ClassID     int64  `gorm:"not null"`
	StudentID   int64  `gorm:"not null"`
	StudentName string `gorm:"type:varchar(20);not null"`
}

type Essay struct {
	ID        int64  `gorm:"primary_key"`
	TeacherID int64  `gorm:"not null"`
	Topic     string `gorm:"type:text;not null"`
	Kind      string `gorm:"type:varchar(20);not null"`
	Standard  string `gorm:"type:text;not null"`
}

type StudentAnswer struct {
	ID             int64 `gorm:"primary_key"`
	StudentID      int64 `gorm:"not null"`
	EssayID        int64 `gorm:"not null"`
	ReferenceScore int
	Score          int
	Answer         string
	Comment        string `gorm:"type:text"`
}
