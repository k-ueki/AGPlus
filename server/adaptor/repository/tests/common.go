package tests

import (
	"bytes"
	"encoding/json"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

type (
	api struct {
		DB *gorm.DB
	}

	args struct {
		id int
	}
)

func (a *api) assertJSON(actual, data interface{}, t *testing.T) {
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
