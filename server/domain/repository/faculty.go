package repository

import (
	"github.com/k-ueki/AGPlus/server/domain/entity"
)

type (
	FacultyGetRepository interface {
		FindFaculties() ([]*entity.Faculty, error)
		FindFacultiesByCampusID(campusID int) ([]*entity.Faculty, error)
		FindFacultyByID(id int) (*entity.Faculty, error)
		FindDepartments() ([]*entity.Department, error)
		FindDepartmentsByCampusID(campusID int) ([]*entity.Department, error)
		FindDepartmentsByFacultyID(facultyID int) ([]*entity.Department, error)
		FindDepartmentByID(id int) (*entity.Department, error)
	}
)
