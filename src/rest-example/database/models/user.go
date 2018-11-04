package models

import (
	//"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	//gorm.Model
	Id   int    `gorm:"AUTO_INCREMENT;PRIMARY_KEY",json:id"`
	Name string `gorm:"type:varchar(50);unique_index",json:name"`
}

func (User) TableName() string {
	return "user"
}
