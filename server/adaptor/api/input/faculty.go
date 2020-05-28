package input

import "errors"

type (
	GetFaculty struct {
		ID int `query:"id"`
	}

	ListFacultyByCampusID struct {
		CampusID int
	}

	ListDepartmentByCampusID struct {
		CampusID  int `form:"campusId"`
		FacultyID int `form:"facultyId"`
	}
)

func (l *ListDepartmentByCampusID) Validate() error {
	if l.CampusID != 0 && l.FacultyID != 0 {
		return errors.New("failed to validate")
	}
	return nil
}
