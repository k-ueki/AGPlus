package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/k-ueki/AGPlus/server/domain/model"
)

type (
	ReviewGetRepository struct {
		DB *gorm.DB
	}
)

func (r *ReviewGetRepository) FindAll() ([]*model.Review, error) {
	var rows []*model.Review
	if err := r.DB.Find(&rows).Error; err != nil {
		return nil, err
	}
	return rows, nil
}
