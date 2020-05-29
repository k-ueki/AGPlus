package service

import (
	"github.com/jinzhu/gorm"
	impl "github.com/k-ueki/AGPlus/server/adaptor/repository"
	"github.com/k-ueki/AGPlus/server/domain/entity"
	"github.com/k-ueki/AGPlus/server/domain/repository"
)

type (
	ReviewGetService interface {
		List() ([]*entity.Review, error)
		Show(classID int) ([]*entity.Review, error)
	}

	ReviewGetServiceImpl struct {
		repository.ReviewGetRepository
	}
)

func NewReviewGetService(db *gorm.DB) ReviewGetService {
	return &ReviewGetServiceImpl{impl.NewReviewGetRepository(db)}
}

func (s *ReviewGetServiceImpl) List() ([]*entity.Review, error) {
	reviews, err := s.ReviewGetRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return reviews, nil
}
func (s *ReviewGetServiceImpl) Show(classID int) ([]*entity.Review, error) {
	reviews, err := s.ReviewGetRepository.FindByClassID(classID)
	if err != nil {
		return nil, err
	}
	return reviews, nil
}
