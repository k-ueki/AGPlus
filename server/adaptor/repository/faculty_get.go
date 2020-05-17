package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/k-ueki/AGPlus/server/domain/entity"
	"github.com/k-ueki/AGPlus/server/domain/model"
)

type (
	FacultyGetRepository struct {
		DB *gorm.DB
	}
)

func (r *FacultyGetRepository) FindFaculties() ([]*entity.Faculty, error) {
	var rows []*entity.Faculty
	if err := r.DB.Where("type = ?", model.FacultyTypeFaculty).Find(&rows).Error; err != nil {
		return nil, err
	}
	return rows, nil
}

func (r *FacultyGetRepository) FindFacultiesByCampusID(campusID int) ([]*entity.Faculty, error) {
	var rows []*entity.Faculty
	if err := r.DB.Where("type = ? AND campus = ?", model.FacultyTypeFaculty, campusID).Find(&rows).Error; err != nil {
		return nil, err
	}
	return rows, nil
}

func (r *FacultyGetRepository) FindFacultyByID(id int) (*entity.Faculty, error) {
	var row entity.Faculty
	if err := r.DB.Where("type = ? AND id = ?", model.FacultyTypeFaculty, id).Find(&row).Error; err != nil {
		return nil, err
	}
	return &row, nil
}

func (r *FacultyGetRepository) FindDepartments() ([]*entity.Department, error) {
	var rows []*entity.Department
	if err := r.DB.Where("type = ?", model.FacultyTypeDepartment).Find(&rows).Error; err != nil {
		return nil, err
	}
	return rows, nil
}

func (r *FacultyGetRepository) FindDepartmentsByCampusID(campusID int) ([]*entity.Department, error) {
	var rows []*entity.Department
	if err := r.DB.Where("type = ? AND campus = ?", model.FacultyTypeDepartment, campusID).Find(&rows).Error; err != nil {
		return nil, err
	}
	return rows, nil
}

func (r *FacultyGetRepository) FindDepartmentsByFacultyID(facultyID int) ([]*entity.Department, error) {
	var rows []*entity.Department
	if err := r.DB.Where("type = ? AND faculty_id = ?", model.FacultyTypeDepartment, facultyID).Find(&rows).Error; err != nil {
		return nil, err
	}
	return rows, nil
}

func (r *FacultyGetRepository) FindDepartmentByID(id int) (*entity.Department, error) {
	var row entity.Department
	if err := r.DB.Where("type = ? AND id = ?", model.FacultyTypeDepartment, id).Find(&row).Error; err != nil {
		return nil, err
	}
	return &row, nil
}
