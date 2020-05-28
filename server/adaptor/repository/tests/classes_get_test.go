package tests

import (
	"regexp"
	"testing"

	"github.com/k-ueki/AGPlus/server/adaptor/repository"
	"github.com/k-ueki/AGPlus/server/domain/query"

	sqlmock "github.com/DATA-DOG/go-sqlmock"

	"github.com/k-ueki/AGPlus/server/domain/model"
)

func TestShouldGetAllClasses(t *testing.T) {
	db, mock, err := connectNewTestMock(t)
	if err != nil {
		t.Error("failed to connect DB : ", err)
		return
	}
	api := api{DB: db}
	repo := repository.ClassGetRepository{DB: db}

	rows := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "hoge").AddRow(2, "piyo")

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `class`")).WillReturnRows(rows)

	resp, err := repo.FindAll(&query.ListPaginationQuery{
		Limit:  20,
		Offset: 0,
	})
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
