package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DNS = "basic_go_rest_api:Secret123!@tcp(localhost:3306)/basic_go_rest_api"

func InitalMigration() {
	DB, err := gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DataBase")
	}

	DB.AutoMigrate(&User{})
}
