package db

import (
	"go-template-api/model"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Conn *gorm.DB
var err error

func ConnectDb() {
	Conn, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	log.Println("----------Connect database----------")
	Conn.AutoMigrate(&model.User_Table{})
	Conn.AutoMigrate(&model.Task_Table{})
}
