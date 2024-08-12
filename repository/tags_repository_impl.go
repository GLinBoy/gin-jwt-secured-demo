package repository

import "gorm.io/gorm"

type TagsRepositoryImpl struct {
	Db *gorm.DB
}
