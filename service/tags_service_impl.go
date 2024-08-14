package service

import (
	"github.com/glinboy/gin-jwt-secured-demo/data/request"
	"github.com/glinboy/gin-jwt-secured-demo/data/response"
	"github.com/glinboy/gin-jwt-secured-demo/model"
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

func (t TagsServiceImpl) Create(tag request.CreateTagsRequest) {
	err := t.Validate.Struct(tag)
	if err != nil {
		panic(err)
	}
	tagModel := model.Tags{
		Name: tag.Name,
	}
	t.TagsRepository.Save(tagModel)
}

func (t TagsServiceImpl) Update(tag request.UpdateTagsRequest) {
	tagData, err := t.TagsRepository.FindById(tag.Id)
	if err != nil {
		panic(err)
	}
	tagData.Name = tag.Name
	t.TagsRepository.Update(tagData)
}

func (t TagsServiceImpl) Delete(tagId int) {
	t.TagsRepository.Delete(tagId)
}

func (t TagsServiceImpl) FindById(tagId int) response.TagsResponse {
	tagData, err := t.TagsRepository.FindById(tagId)
	if err != nil {
		panic(err)
	}
	tagResponse := response.TagsResponse{
		Id:   tagData.Id,
		Name: tagData.Name,
	}
	return tagResponse
}
