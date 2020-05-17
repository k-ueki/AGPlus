package input

type (
	GetFaculty struct {
		ID int `query:"id"`
	}

	ListFacultyByCampusID struct {
		CampusID int `query:"campusID"`
	}

	ListDepartmentByCampusID struct {
		CampusID int `query:"campusID"`
	}
)
