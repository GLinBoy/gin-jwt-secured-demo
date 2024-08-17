package repository

import (
	"github.com/glinboy/gin-jwt-secured-demo/model"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func NewUserRepositoryImpl(Db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{Db: Db}
}

func (u UserRepositoryImpl) Save(user model.User) {
	result := u.Db.Create(&user)
	if result.Error != nil {
		panic(result.Error)
	}
}
