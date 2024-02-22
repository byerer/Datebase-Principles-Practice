package database

import (
	"gorm.io/gorm"
	"time"
)

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

type Class struct {
	gorm.Model
	ClassName string    `gorm:"type:varchar(20);not null"`
	Teacher   Teacher   `gorm:"foreignKey:ClassID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Student   []Student `gorm:"foreignKey:ClassID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Essay struct {
	gorm.Model
	TeacherID int64
	Teacher   Teacher `gorm:"references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // 外键 belongs to
	Topic     string  `gorm:"type:text;not null"`
	Kind      string  `gorm:"type:varchar(20);not null"`
	Standard  string  `gorm:"type:text;not null"`
}

type StudentAnswer struct {
	gorm.Model
	StudentID      int64
	Student        Student `gorm:"references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // 外键 belongs to
	EssayID        int64
	Essay          Essay `gorm:"references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // 外键 belongs to
	ReferenceScore int
	Score          int
	Answer         string
	Comment        string `gorm:"type:text"`
}
