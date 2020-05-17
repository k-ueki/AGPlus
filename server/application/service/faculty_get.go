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

func (s *FacultyGetService) ListFacultyByCampusID(campusID int) ([]*entity.Faculty, error) {
	faculties, err := s.FacultyGetRepository.FindFacultiesByCampusID(campusID)
	if err != nil {
		return nil, err
	}
	return faculties, nil
}

func (s *FacultyGetService) ShowFacultyByID(id int) (*entity.Faculty, error) {
	faculty, err := s.FacultyGetRepository.FindFacultyByID(id)
	if err != nil {
		return nil, err
	}
	return faculty, nil
}

func (s *FacultyGetService) ListDepartment() ([]*entity.Department, error) {
	departments, err := s.FacultyGetRepository.FindDepartments()
	if err != nil {
		return nil, err
	}
	return departments, nil
}

func (s *FacultyGetService) ListDepartmentsByCampusID(campusID int) ([]*entity.Department, error) {
	departments, err := s.FacultyGetRepository.FindDepartmentsByCampusID(campusID)
	if err != nil {
		return nil, err
	}
	return departments, nil
}

func (s *FacultyGetService) ShowDepartmentByID(id int) (*entity.Department, error) {
	department, err := s.FacultyGetRepository.FindDepartmentByID(id)
	if err != nil {
		return nil, err
	}
	return department, nil
}
