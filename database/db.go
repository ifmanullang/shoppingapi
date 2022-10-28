package database

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// github.com/denisenkom/go-mssqldb

var Db *gorm.DB
func InitDb() *gorm.DB { // OOP constructor
	Db = connectDB()
	return Db
}

func connectDB() (*gorm.DB) {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err !=nil {
		fmt.Println("Error...")
		return nil
	}
	return db
}