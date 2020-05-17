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

func (r *ClassGetRepository) FindAll(start, fin int) ([]*model.Class, error) {
	var rows []*model.Class
	if err := r.DB.LogMode(true).Where("id BETWEEN ? AND ?", start, fin).Find(&rows).Error; err != nil {
		return nil, err
	}
	return rows, nil
}

func (r *ClassGetRepository) FindByID(id int) (*model.Class, error) {
	var row model.Class
	if err := r.DB.First(&row, id).Error; err != nil {
		return nil, err
	}
	return &row, nil
}
