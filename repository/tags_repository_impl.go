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

func (t TagsRepositoryImpl) Update(tags model.Tags) {
	// t.Db.Model(&tags).Update()
}

func (t TagsRepositoryImpl) Delete(tagsId int) {
	var tags model.Tags
	t.Db.Where("id = ?", tagsId).Delete(&tags)
}
