package repository

import (
	"github.com/k-ueki/AGPlus/server/domain/entity"
	"github.com/k-ueki/AGPlus/server/domain/model"
)

type (
	ReviewPostRepository interface {
		Store(p *model.Review) error
		Delete(id int) error
	}

	ReviewGetRepository interface {
		FindAll() ([]*entity.Review, error)
		FindByClassID(classID int) ([]*entity.Review, error)
		IsExist(userID, classID int) (bool, error)
	}
)
