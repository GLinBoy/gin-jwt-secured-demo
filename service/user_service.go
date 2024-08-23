package service

import (
	"github.com/glinboy/gin-jwt-secured-demo/data/request"
	"github.com/glinboy/gin-jwt-secured-demo/data/response"
)

type UserService interface {
	Signup(user request.CreateUserRequest)
	Signin(signin request.SigninRequest) response.TokenResponse
}
