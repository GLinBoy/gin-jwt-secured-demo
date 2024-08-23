package service

import "github.com/glinboy/gin-jwt-secured-demo/data/request"

type UserService interface {
	Signup(user request.CreateUserRequest)
	Signin(signin request.SigninRequest)
}
