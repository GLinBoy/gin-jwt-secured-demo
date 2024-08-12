package repository

import (
	"github.com/glinboy/gin-jwt-secured-demo/model"
	"gorm.io/gorm"
)

type TagsRepositoryImpl struct {
	Db *gorm.DB
}

func NewTagsRepositoryImpl(Db *gorm.DB) TagsRepository {
	return &TagsRepositoryImpl{Db: Db}
}

func (t TagsRepositoryImpl) Save(tags model.Tags) {
	t.Db.Create(&tags)
}
