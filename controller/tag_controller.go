package controller

import "github.com/glinboy/gin-jwt-secured-demo/service"

type TagController struct {
	tagService service.TagsService
}
