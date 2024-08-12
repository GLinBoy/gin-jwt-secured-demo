package repository

import (
	"errors"

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

func (t TagsRepositoryImpl) FindById(tagsId int) (model.Tags, error) {
	var tag model.Tags
	result := t.Db.Find(&tag, tagsId)
	if result != nil {
		return tag, nil
	} else {
		return tag, errors.New("tag is not found")
	}
}

func (t TagsRepositoryImpl) FindAll() []model.Tags {
	var tags []model.Tags
	t.Db.Find(&tags)
	return tags
}
