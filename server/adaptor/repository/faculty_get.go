package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/k-ueki/AGPlus/server/domain/entity"
	"github.com/k-ueki/AGPlus/server/domain/model"
	"github.com/k-ueki/AGPlus/server/domain/repository"
)

type (
	FacultyGetRepositoryImpl struct {
		DB *gorm.DB
	}
)

func NewFacultyGetRepository(db *gorm.DB) repository.FacultyGetRepository {
	return &FacultyGetRepositoryImpl{DB: db}
}

func (r *FacultyGetRepositoryImpl) FindFaculties() ([]*entity.Faculty, error) {
	var rows []*entity.Faculty
	if err := r.DB.Where("type = ?", model.FacultyTypeFaculty).Find(&rows).Error; err != nil {
		return nil, err
	}
	return rows, nil
}

func (r *FacultyGetRepositoryImpl) FindFacultiesByCampusID(campusID int) ([]*entity.Faculty, error) {
	var rows []*entity.Faculty
	if err := r.DB.LogMode(true).Where("type = ? AND campus = ?", model.FacultyTypeFaculty, campusID).Find(&rows).Error; err != nil {
		return nil, err
	}
	return rows, nil
}

func (r *FacultyGetRepositoryImpl) FindFacultyByID(id int) (*entity.Faculty, error) {
	var row entity.Faculty
	if err := r.DB.Where("type = ? AND id = ?", model.FacultyTypeFaculty, id).Find(&row).Error; err != nil {
		return nil, err
	}
	return &row, nil
}

func (r *FacultyGetRepositoryImpl) FindDepartments() ([]*entity.Department, error) {
	var rows []*entity.Department
	if err := r.DB.Where("type = ?", model.FacultyTypeDepartment).Find(&rows).Error; err != nil {
		return nil, err
	}
	return rows, nil
}

func (r *FacultyGetRepositoryImpl) FindDepartmentsByCampusID(campusID int) ([]*entity.Department, error) {
	var rows []*entity.Department
	if err := r.DB.Where("type = ? AND campus = ?", model.FacultyTypeDepartment, campusID).Find(&rows).Error; err != nil {
		return nil, err
	}
	return rows, nil
}

func (r *FacultyGetRepositoryImpl) FindDepartmentsByFacultyID(facultyID int) ([]*entity.Department, error) {
	var rows []*entity.Department
	if err := r.DB.Where("type = ? AND faculty_id = ?", model.FacultyTypeDepartment, facultyID).Find(&rows).Error; err != nil {
		return nil, err
	}
	return rows, nil
}

func (r *FacultyGetRepositoryImpl) FindDepartmentByID(id int) (*entity.Department, error) {
	var row entity.Department
	if err := r.DB.Where("type = ? AND id = ?", model.FacultyTypeDepartment, id).Find(&row).Error; err != nil {
		return nil, err
	}
	return &row, nil
}
