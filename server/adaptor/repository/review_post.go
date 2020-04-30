package repository

import "github.com/jinzhu/gorm"

type (
	ReviewPostRepository struct {
		DB *gorm.DB
	}
)
