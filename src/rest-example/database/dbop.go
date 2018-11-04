package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type DBOperator struct {
	DB *gorm.DB
}

func NewOperator(dbUrl string) (*DBOperator, error) {
	db, error := gorm.Open("mysql", dbUrl)

	if error != nil {
		fmt.Printf("%v", error)
	}
	return &DBOperator{
		DB: db,
	}, nil
}
