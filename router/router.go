package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/glinboy/gin-jwt-secured-demo/controller"
	"github.com/glinboy/gin-jwt-secured-demo/middleware"
)

func NewRouter(tagController *controller.TagController,
	userController *controller.UserController) *gin.Engine {
	service := gin.Default()
	service.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Welcome Home!")
	})

	service.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	router := service.Group("/api")
	tagRouter := router.Group("/tag")
	tagRouter.GET("", middleware.RequireAuth, tagController.FindAll)
	tagRouter.GET("/:tagId", middleware.RequireAuth, tagController.FindById)
	tagRouter.POST("", middleware.RequireAuth, tagController.Create)
	tagRouter.PATCH("/:tagId", middleware.RequireAuth, tagController.Update)
	tagRouter.DELETE("/:tagId", middleware.RequireAuth, tagController.Delete)

	authRouter := router.Group("/auth")
	authRouter.POST("/signin", userController.Signin)
	authRouter.POST("/signup", userController.Signup)

	return service
}
