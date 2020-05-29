package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/k-ueki/AGPlus/server/domain/entity"
	"github.com/k-ueki/AGPlus/server/domain/model"
	"github.com/k-ueki/AGPlus/server/domain/repository"
	"github.com/pkg/errors"
)

type (
	ReviewPostRepositoryImpl struct {
		DB *gorm.DB
	}
)

func NewReviewPostRepository(db *gorm.DB) repository.ReviewPostRepository {
	return &ReviewPostRepositoryImpl{DB: db}
}

func (r *ReviewPostRepositoryImpl) Store(p *model.Review) error {
	if err := r.DB.Save(p).Error; err != nil {
		return errors.Wrap(err, "failed to insert review")
	}
	return nil
}

func (r *ReviewPostRepositoryImpl) Delete(id int) error {
	return errors.Wrapf(r.DB.Delete(&entity.Review{ID: id}).Error, "failed to delete review(id=%d)", id)
}
