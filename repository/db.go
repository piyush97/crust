package repository

import (
	"log"
	"os"

	"gorm.io/driver/mysql"

	"github.com/piyush97/crust/model"
	"gorm.io/gorm"
)

/**
 * Database Connection
 * @return *gorm.DB
 * @return error
 */
func DB() *gorm.DB {

	username := os.Getenv("DB_USERNAME")                                                                              // Get the username from the environmental variable.
	password := os.Getenv("DB_PASSWORD")                                                                              // Get the password from the environmental variable.
	dbname := os.Getenv("DB_NAME")                                                                                    // Get the database name from the environmental variable.
	dsn := username + ":" + password + "@tcp(127.0.0.1:3306)/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local" // Create the DSN.

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) // Open the database connection.

	if err != nil {
		log.Fatal("Error connecting to Database") // Log the error if there is one.
		return nil
	}

	db.AutoMigrate(&model.User{}, &model.Product{}, &model.Order{}) // Auto migrate the database.
	return db                                                       // Return the database connection.
}
