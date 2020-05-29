package service

import (
	"github.com/jinzhu/gorm"
	"github.com/k-ueki/AGPlus/server/adaptor/api/input"
	impl "github.com/k-ueki/AGPlus/server/adaptor/repository"
	"github.com/k-ueki/AGPlus/server/domain/model"
	"github.com/k-ueki/AGPlus/server/domain/repository"
	"github.com/pkg/errors"
)

type (
	ReviewPostService interface {
		Store(classID int, param *input.ReviewClassRequest) error
		Delete(id int) error
	}

	ReviewPostServiceImpl struct {
		repository.ReviewPostRepository
	}
)

func NewReviewPostService(db *gorm.DB) ReviewPostService {
	return &ReviewPostServiceImpl{impl.NewReviewPostRepository(db)}
}

func (s *ReviewPostServiceImpl) Store(classID int, param *input.ReviewClassRequest) error {
	if err := s.ReviewPostRepository.Store(&model.Review{
		ClassID:         classID,
		Understanding:   param.Understanding,
		Motivation:      param.Motivation,
		Attendance:      param.Attendance,
		TestsDifficulty: param.TestsDifficulty,
		Easiness:        param.Easiness,
	}); err != nil {
		return errors.Wrap(err, "failed to store")
	}
	return nil
}

func (s *ReviewPostServiceImpl) Delete(id int) error {
	if err := s.ReviewPostRepository.Delete(id); err != nil {
		return errors.Wrap(err, "failed to delete")
	}
	return nil
}
