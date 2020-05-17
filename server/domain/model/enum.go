package model

//go:generate go-enum -f=$GOFILE --marshal

/*
ENUM(Aoyama=1, Sagamihara)
*/
type Campus int

/*
ENUM(Campus = 1,Faculty,Department)
*/
type FacultyType int
