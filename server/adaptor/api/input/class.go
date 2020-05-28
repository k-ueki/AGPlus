package input

type (
	GetClass struct {
		ID int `form:"id" binding:"exists"`
	}

	ListClass struct {
		PerPage int `form:"perPage"`
		Page    int `form:"page"`
	}
)

const showClassListDefaultPerPage = 20

func (lc ListClass) GetLimit() int {
	if lc.PerPage == 0 {
		return showClassListDefaultPerPage
	}
	return lc.PerPage
}

// offSetを返す(sqlで使う想定)
func (lc ListClass) GetOffSet() int {
	if lc.Page == 1 || lc.Page == 0 {
		return 0
	}
	return lc.GetLimit() * (lc.Page - 1)
}
