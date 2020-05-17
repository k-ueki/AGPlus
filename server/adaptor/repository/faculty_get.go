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

func (r *FacultyGetRepository) FindDepartments() ([]*entity.Faculty, error) {
	var rows []*entity.Faculty
	if err := r.DB.Where("type = ?", model.FacultyTypeDepartment).Find(&rows).Error; err != nil {
		return nil, err
	}
	return rows, nil
}
