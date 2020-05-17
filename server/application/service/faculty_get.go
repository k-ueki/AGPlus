package service

import (
	"github.com/k-ueki/AGPlus/server/adaptor/repository"
	"github.com/k-ueki/AGPlus/server/domain/entity"
)

type (
	FacultyGetService struct {
		repository.FacultyGetRepository
	}
)

func (s *FacultyGetService) List() ([]*entity.Faculty, error) {
	faculties, err := s.FacultyGetRepository.FindFaculties()
	if err != nil {
		return nil, err
	}
	return faculties, nil
}
