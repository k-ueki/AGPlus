package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/k-ueki/AGPlus/server/domain/model"
	"github.com/pkg/errors"
)

type (
	ReviewPostRepository struct {
		DB *gorm.DB
	}
)

func (r *ReviewPostRepository) Store(p *model.Review) error {
	if err := r.DB.Save(p).Error; err != nil {
		return errors.Wrap(err, "failed to insert review")
	}
	return nil
}
