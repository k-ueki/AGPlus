package service

import (
	"github.com/jinzhu/gorm"
	impl "github.com/k-ueki/AGPlus/server/adaptor/repository"
	"github.com/k-ueki/AGPlus/server/domain/model"
	"github.com/k-ueki/AGPlus/server/domain/query"
	repository "github.com/k-ueki/AGPlus/server/domain/repository"
)

type (
	ClassGetService interface {
		List(query *query.ListPaginationQuery) ([]*model.Class, error)
		Show(id int) (*model.Class, error)
	}

	ClassGetServiceImpl struct {
		ClassGetRepository repository.ClassGetRepository
	}
)

func NewClassGetService(db *gorm.DB) ClassGetService {
	return &ClassGetServiceImpl{impl.NewClassGetRepository(db)}
}

func (s *ClassGetServiceImpl) List(query *query.ListPaginationQuery) ([]*model.Class, error) {
	classes, err := s.ClassGetRepository.FindAll(query)
	if err != nil {
		return nil, err
	}
	return classes, nil
}

func (s *ClassGetServiceImpl) Show(id int) (*model.Class, error) {
	class, err := s.ClassGetRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return class, nil
}
