package entity

import "github.com/k-ueki/AGPlus/server/domain/model"

type (
	Review struct {
		ID              int
		UserID          int
		ClassID         int
		Understanding   int
		Motivation      int
		Attendance      int
		TestsDifficulty int
		Easiness        int
	}

	Evaluation struct {
		Understanding   int
		Motivation      int
		Attendance      int
		TestsDifficulty int
		Easiness        int
	}

	ReviewInfo struct {
		ReviewUser model.User
		model.Class
		Review
	}
)

func (r *Review) TableName() string {
	return "review"
}
