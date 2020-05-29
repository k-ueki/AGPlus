package repository

import (
	"github.com/k-ueki/AGPlus/server/domain/entity"
)

type (
	ReviewPostRepository interface {
		Store(p *entity.Review) error
		Delete(id int) error
	}

	ReviewGetRepository interface {
		FindAll() ([]*entity.Review, error)
		FindByClassID(classID int) ([]*entity.Review, error)
		IsExist(userID, classID int) (bool, error)
	}
)
