package model

type (
	Review struct {
		ClassID         int
		Understanding   int `gorm:"understanding"`
		Motivation      int `gorm:"motivation"`
		Attendance      int `gorm:"attendance"`
		TestsDifficulty int `gorm:"tests_difficulty"`
		Easiness        int `gorm:"easiness"`
	}
)

func (r *Review) TableName() string {
	return "review"
}
