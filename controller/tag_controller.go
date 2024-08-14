package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/glinboy/gin-jwt-secured-demo/data/request"
	"github.com/glinboy/gin-jwt-secured-demo/data/response"
	"github.com/glinboy/gin-jwt-secured-demo/service"
)

type TagController struct {
	tagService service.TagsService
}

func NewTagController(service service.TagsService) *TagController {
	return &TagController{tagService: service}
}

func (controller *TagController) Create(ctx *gin.Context) {
	createTagRequest := request.CreateTagsRequest{}
	err := ctx.ShouldBindJSON(&createTagRequest)
	if err != nil {
		panic(err)
	}
	controller.tagService.Create(createTagRequest)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}
	ctx.JSON(http.StatusOK, webResponse)
}
