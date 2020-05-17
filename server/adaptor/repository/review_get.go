package repository

import "github.com/jinzhu/gorm"

type (
	ReviewGetRepository struct {
		DB *gorm.DB
	}
)
