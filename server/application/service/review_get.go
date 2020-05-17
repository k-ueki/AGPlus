package service

import (
	"github.com/k-ueki/AGPlus/server/adaptor/repository"
	"github.com/k-ueki/AGPlus/server/domain/model"
)

type (
	ReviewGetService struct {
		repository.ReviewGetRepository
	}
)

func (s *ReviewGetService) List() (*model.Review, error) {
	return nil, nil
}
