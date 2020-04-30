package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/k-ueki/AGPlus/server/domain/model"
)

type (
	ClassGetRepository struct {
		DB *gorm.DB
	}
)

func (r *ClassGetRepository) FindAll() ([]*model.Class, error) {
	var rows []*model.Class
	if err := r.DB.Find(&rows).Error; err != nil {
		return nil, err
	}
	return rows, nil
}
