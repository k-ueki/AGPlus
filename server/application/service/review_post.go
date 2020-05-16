package service

import (
	"github.com/k-ueki/AGPlus/server/adaptor/api/input"
	"github.com/k-ueki/AGPlus/server/adaptor/repository"
)

type (
	ReviewPostService struct {
		repository.ReviewPostRepository
	}
)

func (s *ReviewPostService) Store(classID int, param *input.ReviewClassRequest) error {
	return nil
}
