package service

import (
	"github.com/k-ueki/AGPlus/server/adaptor/repository"
	"github.com/k-ueki/AGPlus/server/domain/model"
	"github.com/k-ueki/AGPlus/server/domain/query"
)

type (
	ClassGetService struct {
		repository.ClassGetRepository
	}
)

func (s *ClassGetService) List(query *query.ListPaginationQuery) ([]*model.Class, error) {
	classes, err := s.ClassGetRepository.FindAll(query)
	if err != nil {
		return nil, err
	}
	return classes, nil
}

func (s *ClassGetService) Show(id int) (*model.Class, error) {
	class, err := s.ClassGetRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return class, nil
}
