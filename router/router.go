package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/glinboy/gin-jwt-secured-demo/controller"
)

func NewRouter(tagController *controller.TagController) *gin.Engine {
	service := gin.Default()
	service.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Welcome Home!")
	})

	service.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	router := service.Group("/api")
	tagRouter := router.Group("/tag")
	tagRouter.GET("", tagController.FindAll)
	tagRouter.GET("/:tagId", tagController.FindById)
	tagRouter.POST("", tagController.Create)
	tagRouter.PATCH("/:tagId", tagController.Update)
	tagRouter.DELETE("/:tagId", tagController.Delete)

	return service
}
