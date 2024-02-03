package database

type User struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `gorm:"type:varchar(20);not null" binding:"required"`
	Password string
	Email    string
}
