package service

import (
	"github.com/glinboy/gin-jwt-secured-demo/repository"
	"github.com/go-playground/validator/v10"
)

type TagsServiceImpl struct {
	TagsRepository repository.TagsRepository
	Validate       *validator.Validate
}

func NewTagServiceImpl(tagRepository repository.TagsRepository, valitate *validator.Validate) TagsService {
	return &TagsServiceImpl{
		TagsRepository: tagRepository,
		Validate:       valitate,
	}
}
