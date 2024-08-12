package service

import "github.com/glinboy/gin-jwt-secured-demo/model"

type TagsService interface {
	Create(tags model.Tags)
	Update(tags model.Tags)
	Delete(tagsId int)
	FindById(tagsId int) model.Tags
	FindAll() []model.Tags
}
