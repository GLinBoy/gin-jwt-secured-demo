package service

import (
	"github.com/glinboy/gin-jwt-secured-demo/repository"
	"github.com/go-playground/validator/v10"
)

type TagsServiceImpl struct {
	TagsRepository repository.TagsRepository
	Validate       *validator.Validate
}
