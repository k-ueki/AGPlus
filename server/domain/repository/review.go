package repository

import "github.com/k-ueki/AGPlus/server/domain/model"

type (
	ReviewPostRepository interface {
		Store(p *model.Review) error
		Delete(id int) error
	}

	ReviewGetRepository interface {
		FindAll() ([]*model.Review, error)
	}
)
