package input

type (
	GetClass struct {
		ID int `form:"id" binding:"exists"`
	}
)
