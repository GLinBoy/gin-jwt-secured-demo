package controller

import "github.com/glinboy/gin-jwt-secured-demo/service"

type UserController struct {
	userService service.UserService
}

func NewUserController(service service.UserService) *UserController {
	return &UserController{userService: service}
}
