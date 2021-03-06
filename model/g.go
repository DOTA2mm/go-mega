package model

import (
	"log"

	"github.com/dota2mm/go-mega/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// SetDB func
func SetDB(database *gorm.DB) {
	db = database
}

// ConnectToDB func
func ConnectToDB() *gorm.DB {
	connectingStr := config.GetMysqlConnectingString()
	log.Println("Connect to db...")
	log.Printf("DB configuration: %s", connectingStr)
	db, err := gorm.Open("mysql", connectingStr)
	if err != nil {
		panic("Failed to connect database")
	}
	db.Set("gorm:table_options", "CHARSET=utf8mb4")
	db.SingularTable(true)
	return db
}
