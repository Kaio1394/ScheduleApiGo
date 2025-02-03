package model

type User struct {
	Id       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Password string
}
