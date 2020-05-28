package model

//go:generate go-enum -f=$GOFILE --marshal

/*
ENUM(Aoyama=1, Sagamihara)
*/
type Campus int

/*
ENUM(Faculty=1,Department)
*/
type FacultyType int
