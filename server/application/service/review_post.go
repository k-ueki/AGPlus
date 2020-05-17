package service

import (
	"github.com/k-ueki/AGPlus/server/adaptor/api/input"
	"github.com/k-ueki/AGPlus/server/adaptor/repository"
	"github.com/k-ueki/AGPlus/server/domain/model"
	"github.com/pkg/errors"
)

type (
	ReviewPostService struct {
		repository.ReviewPostRepository
	}
)

func (s *ReviewPostService) Store(classID int, param *input.ReviewClassRequest) error {
	if err := s.ReviewPostRepository.Store(&model.Review{
		ClassID:         classID,
		Understanding:   param.Understanding,
		Motivation:      param.Motivation,
		Attendance:      param.Attendance,
		TestsDifficulty: param.TestsDifficulty,
		Easiness:        param.Easiness,
	}); err != nil {
		return errors.Wrap(err, "failed to store")
	}
	return nil
}

func (s *ReviewPostService) Delete(id int) error {
	if err := s.ReviewPostRepository.Delete(id); err != nil {
		return errors.Wrap(err, "failed to delete")
	}
	return nil
}
