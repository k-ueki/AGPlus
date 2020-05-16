package model

type (
	ClassReview struct {
		Class
		Evaluation
	}

	Evaluation struct {
		EaseOfUnderstanding int
		Motivation          int
		TestsDifficulty     int
		Easiness            int
	}
)
