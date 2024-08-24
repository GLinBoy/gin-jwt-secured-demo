package main

import (
	"net/http"
	"time"

	"github.com/glinboy/gin-jwt-secured-demo/config"
	"github.com/glinboy/gin-jwt-secured-demo/controller"
	"github.com/glinboy/gin-jwt-secured-demo/model"
	"github.com/glinboy/gin-jwt-secured-demo/repository"
	"github.com/glinboy/gin-jwt-secured-demo/router"
	"github.com/glinboy/gin-jwt-secured-demo/service"
	"github.com/go-playground/validator/v10"
)

func main() {
	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("tags").AutoMigrate(&model.Tags{})
	db.Table("user").AutoMigrate(&model.User{})

	tagRepository := repository.NewTagsRepositoryImpl(db)
	userRepository := repository.NewUserRepositoryImpl(db)

	tagService := service.NewTagServiceImpl(tagRepository, validate)
	userService := service.NewUserServiceImpl(userRepository, validate)

	tagController := controller.NewTagController(tagService)
	userController := controller.NewUserController(userService)

	routes := router.NewRouter(tagController, userController)

	server := &http.Server{
		Addr:           ":8888",
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
