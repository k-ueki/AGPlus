package entity

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
		Comment         string
	}

	Evaluation struct {
		Understanding   int
		Motivation      int
		Attendance      int
		TestsDifficulty int
		Easiness        int
	}

	ReviewInfo struct {
		ReviewUser User
		Class
		Review
	}
)

func (r *Review) TableName() string {
	return "review"
}
