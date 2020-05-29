package service

import (
	"github.com/jinzhu/gorm"
	"github.com/k-ueki/AGPlus/server/adaptor/api/input"
	impl "github.com/k-ueki/AGPlus/server/adaptor/repository"
	"github.com/k-ueki/AGPlus/server/domain/entity"
	"github.com/k-ueki/AGPlus/server/domain/repository"
)

type (
	FacultyGetService interface {
		ListFacultyByParam(param *input.ListFacultyByCampusID) ([]*entity.Faculty, error)
		ListFaculty() ([]*entity.Faculty, error)
		ListFacultyByCampusID(campusID int) ([]*entity.Faculty, error)
		ShowFacultyByID(id int) (*entity.Faculty, error)
		ListDepartmentsByParam(param *input.ListDepartmentByCampusID) ([]*entity.Department, error)
		ListDepartments() ([]*entity.Department, error)
		ListDepartmentsByCampusID(campusID int) ([]*entity.Department, error)
		ListDepartmentsByFacultyID(facultyID int) ([]*entity.Department, error)
		ShowDepartmentByID(id int) (*entity.Department, error)
	}

	FacultyGetServiceImpl struct {
		repository.FacultyGetRepository
	}
)

func NewFacultyGetServiceImpl(db *gorm.DB) FacultyGetService {
	return &FacultyGetServiceImpl{impl.NewFacultyGetRepository(db)}
}

func (s *FacultyGetServiceImpl) ListFacultyByParam(param *input.ListFacultyByCampusID) ([]*entity.Faculty, error) {
	if param.CampusID != 0 {
		faculties, err := s.ListFacultyByCampusID(param.CampusID)
		if err != nil {
			return nil, err
		}
		return faculties, nil
	}

	faculties, err := s.ListFaculty()
	if err != nil {
		return nil, err
	}
	return faculties, nil
}

func (s *FacultyGetServiceImpl) ListFaculty() ([]*entity.Faculty, error) {
	faculties, err := s.FacultyGetRepository.FindFaculties()
	if err != nil {
		return nil, err
	}
	return faculties, nil
}

func (s *FacultyGetServiceImpl) ListFacultyByCampusID(campusID int) ([]*entity.Faculty, error) {
	faculties, err := s.FacultyGetRepository.FindFacultiesByCampusID(campusID)
	if err != nil {
		return nil, err
	}
	return faculties, nil
}

func (s *FacultyGetServiceImpl) ShowFacultyByID(id int) (*entity.Faculty, error) {
	faculty, err := s.FacultyGetRepository.FindFacultyByID(id)
	if err != nil {
		return nil, err
	}
	return faculty, nil
}

func (s *FacultyGetServiceImpl) ListDepartmentsByParam(param *input.ListDepartmentByCampusID) ([]*entity.Department, error) {
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

func (s *FacultyGetServiceImpl) ListDepartments() ([]*entity.Department, error) {
	departments, err := s.FacultyGetRepository.FindDepartments()
	if err != nil {
		return nil, err
	}
	return departments, nil
}

func (s *FacultyGetServiceImpl) ListDepartmentsByCampusID(campusID int) ([]*entity.Department, error) {
	departments, err := s.FacultyGetRepository.FindDepartmentsByCampusID(campusID)
	if err != nil {
		return nil, err
	}
	return departments, nil
}

func (s *FacultyGetServiceImpl) ListDepartmentsByFacultyID(facultyID int) ([]*entity.Department, error) {
	departments, err := s.FacultyGetRepository.FindDepartmentsByFacultyID(facultyID)
	if err != nil {
		return nil, err
	}
	return departments, nil
}

func (s *FacultyGetServiceImpl) ShowDepartmentByID(id int) (*entity.Department, error) {
	department, err := s.FacultyGetRepository.FindDepartmentByID(id)
	if err != nil {
		return nil, err
	}
	return department, nil
}
