package repository

import (
	"github.com/k-ueki/AGPlus/server/domain/entity"
	"github.com/k-ueki/AGPlus/server/domain/query"
)

type (
	ClassGetRepository interface {
		FindAll(query *query.ListPaginationQuery) ([]*entity.Class, error)
		FindByID(id int) (*entity.Class, error)
	}
)
