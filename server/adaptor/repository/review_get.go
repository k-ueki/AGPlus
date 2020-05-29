package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/k-ueki/AGPlus/server/domain/model"
	"github.com/k-ueki/AGPlus/server/domain/repository"
)

type (
	ReviewGetRepositoryImpl struct {
		DB *gorm.DB
	}
)

func NewReviewGetRepository(db *gorm.DB) repository.ReviewGetRepository {
	return &ReviewGetRepositoryImpl{DB: db}
}

func (r *ReviewGetRepositoryImpl) FindAll() ([]*model.Review, error) {
	var rows []*model.Review
	if err := r.DB.Find(&rows).Error; err != nil {
		return nil, err
	}
	return rows, nil
}
