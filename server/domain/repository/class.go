package repository

import (
	"github.com/k-ueki/AGPlus/server/domain/model"
	"github.com/k-ueki/AGPlus/server/domain/query"
)

type (
	ClassGetRepository interface {
		FindAll(query *query.ListPaginationQuery) ([]*model.Class, error)
		FindByID(id int) (*model.Class, error)
	}
)
