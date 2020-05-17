package service

import (
	"github.com/k-ueki/AGPlus/server/adaptor/repository"
	"github.com/k-ueki/AGPlus/server/domain/model"
)

type (
	FacultyGetService struct {
		repository.FacultyGetRepository
	}
)

func (s *FacultyGetService) List() ([]*model.Faculty, error) {
	return nil, nil
}
