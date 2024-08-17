package model

type User struct {
	Id       int    `gorm:"type:int;primary_key"`
	FullName string `gorm:"type:varchar(255);"`
	Email    string `gorm:"type:varchar(255);unique;"`
	Password string `gorm:"type:varchar(255)"`
}
