package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/k-ueki/AGPlus/server/domain/model"
	"github.com/k-ueki/AGPlus/server/domain/query"
)

type (
	ClassGetRepository struct {
		DB *gorm.DB
	}
)

func (r *ClassGetRepository) FindAll(query *query.ListPaginationQuery) ([]*model.Class, error) {
	var rows []*model.Class
	if err := r.DB.LogMode(true).Offset(query.Offset).Limit(query.Limit).Find(&rows).Error; err != nil {
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
