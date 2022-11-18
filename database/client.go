package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"jwt-golang/models"
	"log"
)

// Instance we are defining an instance of the database. This variable will be used across the entire application to communicate with the database.
var Instance *gorm.DB
var dbError error

// Connect function takes in the MySQL connection string (which we are going to pass from the main method shortly) and tries to connect to the database using GORM.
func Connect(connectionString string) {
	Instance, dbError = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database")
}

func Migrate() {
	Instance.AutoMigrate(&models.User{})
	log.Println("Database Migration Completed")
}
