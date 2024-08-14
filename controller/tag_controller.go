package controller

import "github.com/glinboy/gin-jwt-secured-demo/service"

type TagController struct {
	tagService service.TagsService
}

func NewTagController(service service.TagsService) *TagController {
	return &TagController{tagService: service}
}
