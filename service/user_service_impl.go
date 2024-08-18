package service

import (
	"github.com/glinboy/gin-jwt-secured-demo/data/request"
	"github.com/glinboy/gin-jwt-secured-demo/model"
	"github.com/glinboy/gin-jwt-secured-demo/repository"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	Validate       *validator.Validate
}

func NewUserServiceImpl(userRepository repository.UserRepository, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		Validate:       validate,
	}
}

func (u UserServiceImpl) Signup(user request.CreateUserRequest) {
	err := u.Validate.Struct(user)
	if err != nil {
		panic(err)
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		panic(err)
	}
	userModel := model.User{
		FullName: user.FullName,
		Email:    user.Email,
		Password: string(hash),
	}
	u.UserRepository.Save(userModel)
}
