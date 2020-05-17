package entity

import "github.com/k-ueki/AGPlus/server/domain/model"

type (
	Faculty struct {
		ID     int
		Name   string
		Type   model.FacultyType
		Campus model.Campus
	}

	Department struct {
	}
)

func (f *Faculty) TableName() string {
	return "faculty"
}
