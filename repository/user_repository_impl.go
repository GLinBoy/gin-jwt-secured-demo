package repository

import (
	"errors"

	"github.com/glinboy/gin-jwt-secured-demo/data/request"
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

func (u UserRepositoryImpl) Update(user model.User) {
	var updateUser = request.UpdateUserRequest{
		Id:       user.Id,
		FullName: user.FullName,
		Password: user.Password,
	}
	result := u.Db.Model(&user).Updates(updateUser)
	if result.Error != nil {
		panic(result.Error)
	}
}

func (u UserRepositoryImpl) Delete(userId int) {
	var user model.User
	result := u.Db.Where("id = ?").Delete(&user)
	if result.Error != nil {
		panic(result.Error)
	}
}

func (u UserRepositoryImpl) FindById(userId int) (model.User, error) {
	var user model.User
	result := u.Db.Find(&user, userId)
	if result != nil {
		return user, nil
	} else {
		return user, errors.New("user is not found")
	}
}

func (u UserRepositoryImpl) FindByEmail(email string) (model.User, error) {
	var user model.User
	result := u.Db.Find(&user, email)
	if result != nil {
		return user, nil
	} else {
		return user, errors.New("user is not found")
	}
}
