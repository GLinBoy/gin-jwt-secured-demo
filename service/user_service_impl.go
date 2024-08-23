package service

import (
	"time"

	"github.com/glinboy/gin-jwt-secured-demo/data/request"
	"github.com/glinboy/gin-jwt-secured-demo/data/response"
	"github.com/glinboy/gin-jwt-secured-demo/model"
	"github.com/glinboy/gin-jwt-secured-demo/repository"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
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

func (u UserServiceImpl) Signin(signin request.SigninRequest) response.TokenResponse {
	userData, err := u.UserRepository.FindByEmail(signin.Email)
	if err != nil {
		panic(err)
	}
	if bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(signin.Password)) != nil {
		panic("Invalid email or password")
	}

	expiration := time.Now().Add(time.Hour * 24 * 30).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userData.Email,
		"exp": expiration,
	})

	tokenString, err := token.SignedString([]byte("SECURE_SECRET"))

	if err != nil {
		panic(err)
	}

	return response.TokenResponse{
		Token:      tokenString,
		Expiration: expiration,
	}
}
