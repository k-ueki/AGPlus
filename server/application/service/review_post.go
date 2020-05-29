package service

import (
	"github.com/jinzhu/gorm"
	"github.com/k-ueki/AGPlus/server/adaptor/api/input"
	impl "github.com/k-ueki/AGPlus/server/adaptor/repository"
	"github.com/k-ueki/AGPlus/server/domain/entity"
	"github.com/k-ueki/AGPlus/server/domain/repository"
	"github.com/pkg/errors"
)

type (
	ReviewPostService interface {
		Store(classID int, user *entity.User, param *input.ReviewClassRequest) error
		Delete(id int) error
	}

	ReviewPostServiceImpl struct {
		repository.ReviewPostRepository
		repository.ReviewGetRepository
	}
)

func NewReviewPostService(db *gorm.DB) ReviewPostService {
	return &ReviewPostServiceImpl{
		impl.NewReviewPostRepository(db),
		impl.NewReviewGetRepository(db),
	}
}

func (s *ReviewPostServiceImpl) Store(classID int, user *entity.User, param *input.ReviewClassRequest) error {
	isExistReview, err := s.ReviewGetRepository.IsExist(user.ID, classID)
	if err != nil {
		return err
	}
	if isExistReview {
		return errors.New("already posted")
	}

	if err := s.ReviewPostRepository.Store(&entity.Review{
		ClassID:         classID,
		UserID:          user.ID,
		Understanding:   param.Understanding,
		Motivation:      param.Motivation,
		Attendance:      param.Attendance,
		TestsDifficulty: param.TestsDifficulty,
		Easiness:        param.Easiness,
		Comment:         param.Comment,
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
