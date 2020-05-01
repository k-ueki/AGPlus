package repository

import "github.com/k-ueki/AGPlus/server/domain/model"

type (
	ClassGetRepository interface {
		FindAll() ([]*model.Class, error)
		FindByID(id int) (*model.Class, error)
	}
)
