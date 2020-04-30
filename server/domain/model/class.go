package model

type (
	Class struct {
		ID          int    `db:"id"`
		Name        string `db:"name"`
		Semester    string `db:"semester"`
		Credits     int    `db:"credits"`
		Teacher     string `db:"teacher"`
		Description string `db:"description"`
		Year        int    `db:"year"`
		DayAt       string `db:"dat_at"`
		Campus      string `db:"campus"`
	}
)

func (c *Class) TableName() string {
	return "class"
}
