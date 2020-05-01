package repository

import (
	"bytes"
	"encoding/json"
	"regexp"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"

	"github.com/jinzhu/gorm"
	"github.com/k-ueki/AGPlus/server/domain/model"
)

type (
	api struct {
		DB *gorm.DB
	}

	args struct {
		id int
	}
)

func (a *api) assertJSON(actual []*model.Class, data interface{}, t *testing.T) {
	act, err := json.Marshal(actual)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when marshaling expected json data", err)
	}

	expected, err := json.Marshal(data)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when marshaling expected json data", err)
	}

	if bytes.Compare(expected, act) != 0 {
		t.Errorf("the expected json: %s is different from actual %s", expected, act)
	}
}

func connectNewTestMock(t *testing.T) (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gdb, err := gorm.Open("mysql", db)
	if err != nil {
		return nil, nil, err
	}
	return gdb, mock, nil
}

func TestShouldGetAllClasses(t *testing.T) {
	db, mock, err := connectNewTestMock(t)
	if err != nil {
		t.Error("failed to connect DB : ", err)
		return
	}
	api := api{DB: db}
	repository := ClassGetRepository{DB: db}

	rows := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "hoge").AddRow(2, "piyo")

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `class`")).WillReturnRows(rows)

	resp, err := repository.FindAll()
	if err != nil {
		t.Fatal("failed to FindAll : ", err)
	}

	data := []*model.Class{
		{
			ID:   1,
			Name: "hoge",
		},
		{
			ID:   2,
			Name: "piyo",
		},
	}
	api.assertJSON(resp, data, t)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
