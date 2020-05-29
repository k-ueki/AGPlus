package service

import (
	"github.com/jinzhu/gorm"
	impl "github.com/k-ueki/AGPlus/server/adaptor/repository"
	"github.com/k-ueki/AGPlus/server/domain/model"
	"github.com/k-ueki/AGPlus/server/domain/repository"
)

type (
	ReviewGetService interface {
		List() ([]*model.Review, error)
	}

	ReviewGetServiceImpl struct {
		repository.ReviewGetRepository
	}
)

func NewReviewGetService(db *gorm.DB) ReviewGetService {
	return &ReviewGetServiceImpl{impl.NewReviewGetRepository(db)}
}

func (s *ReviewGetServiceImpl) List() ([]*model.Review, error) {
	reviews, err := s.ReviewGetRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return reviews, nil
}
