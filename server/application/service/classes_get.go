package service

import (
	"github.com/k-ueki/AGPlus/server/adaptor/repository"
	"github.com/k-ueki/AGPlus/server/domain/model"
)

type (
	ClassGetService struct {
		repository.ClassGetRepository
	}
)

func (s *ClassGetService) List() ([]*model.Class, error) {
	classes, err := s.ClassGetRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return classes, nil
}
