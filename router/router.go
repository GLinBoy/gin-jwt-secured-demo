package router

import (
	"github.com/gin-gonic/gin"
	"github.com/glinboy/gin-jwt-secured-demo/controller"
)

func NewRouter(tagController *controller.TagController) *gin.Engine {
	service := gin.Default()

	return service
}
