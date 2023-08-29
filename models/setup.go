package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Utility function to create a connection to database and migrate the model's schema.
func ConnectDatabase() {
	db, err := gorm.Open(sqlite.Open("employee.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = db.AutoMigrate(&Employee{})
	if err != nil {
		return
	}

	DB = db
}
