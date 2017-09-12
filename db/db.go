package db

import (
	_ "../model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

func connect() *gorm.DB {
	options := "host=localhost" + " user=postgres"  + " dbname=postgres"  + " sslmode=disable password=postgres" 
	db, err := gorm.Open("postgres", options)
	if err != nil {
		log.Println("ERROR:", err)
		panic("failed to connect database")

	}
	return db
}


//CreateTable can create new table from struct interface
func CreateTable(model interface{}) error {
	db := connect()
	defer db.Close()
	db.CreateTable(model)
	return nil
}