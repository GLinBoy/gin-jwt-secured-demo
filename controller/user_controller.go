package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/glinboy/gin-jwt-secured-demo/data/request"
	"github.com/glinboy/gin-jwt-secured-demo/data/response"
	"github.com/glinboy/gin-jwt-secured-demo/service"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(service service.UserService) *UserController {
	return &UserController{userService: service}
}

func (c *UserController) Signup(ctx *gin.Context) {
	signupRequest := request.CreateUserRequest{}
	err := ctx.ShouldBindJSON(&signupRequest)
	if err != nil {
		panic(err)
	}

	c.userService.Signup(signupRequest)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func (c *UserController) Signin(ctx *gin.Context) {
	signinRequest := request.SigninRequest{}
	err := ctx.ShouldBindJSON(&signinRequest)
	if err != nil {
		panic(err)
	}
	c.userService.Signin(signinRequest)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   signinRequest,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
