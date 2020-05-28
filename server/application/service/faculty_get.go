package service

import (
	"github.com/k-ueki/AGPlus/server/adaptor/api/input"
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

func (s *FacultyGetService) ListDepartmentsByParam(param *input.ListDepartmentByCampusID) ([]*entity.Department, error) {
	if param.CampusID != 0 {
		departments, err := s.ListDepartmentsByCampusID(param.CampusID)
		if err != nil {
			return nil, err
		}
		return departments, nil
	}

	if param.FacultyID != 0 {
		departments, err := s.ListDepartmentsByFacultyID(param.FacultyID)
		if err != nil {
			return nil, err
		}
		return departments, nil
	}

	departments, err := s.ListDepartments()
	if err != nil {
		return nil, err
	}
	return departments, nil
}

func (s *FacultyGetService) ListDepartments() ([]*entity.Department, error) {
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

func (s *FacultyGetService) ListDepartmentsByFacultyID(facultyID int) ([]*entity.Department, error) {
	departments, err := s.FacultyGetRepository.FindDepartmentsByFacultyID(facultyID)
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
