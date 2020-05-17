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

func (s *FacultyGetService) ListFaculty() ([]*entity.Faculty, error) {
	faculties, err := s.FacultyGetRepository.FindFaculties()
	if err != nil {
		return nil, err
	}
	return faculties, nil
}

func (s *FacultyGetService) ListDepartment() ([]*entity.Department, error) {
	departments, err := s.FacultyGetRepository.FindDepartments()
	if err != nil {
		return nil, err
	}
	return departments, nil
}
