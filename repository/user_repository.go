package repository

import "github.com/glinboy/gin-jwt-secured-demo/model"

type UserRepository interface {
	Save(user model.User)
	Update(user model.User)
	FindById(userId int) (user model.User, err error)
	FindByEmail(email string) (user model.User, err error)
}
