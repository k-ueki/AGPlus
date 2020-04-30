package common

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var table = "agplus"

func NewDB() (*gorm.DB, error) {
	CONNECT := "root:@tcp(localhost:3306)/" + table
	db, err := gorm.Open("mysql", CONNECT)
	if err != nil {
		return nil, err
	}
	return db, nil
}
