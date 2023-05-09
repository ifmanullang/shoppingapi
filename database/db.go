package database

import (
	"fmt"
	"gorm.io/driver/mysql"
    _ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)


var Db *gorm.DB
func InitDb() *gorm.DB { // OOP constructor
	Db = connectDB()
	return Db
}

func connectDB() (*gorm.DB) {
	
	dsn := "root:@tcp(localhost:3306)/cart_db"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})


	if err != nil {
		fmt.Println("Error...")
		return nil
	}
	return db
}