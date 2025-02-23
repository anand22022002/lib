package database

import (
	"online-library-system/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	database, err := gorm.Open(sqlite.Open("librarydb.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	database.AutoMigrate(&models.Library{}, &models.User{}, &models.BookInventory{}, &models.RequestEvent{}, &models.IssueRegistry{})
	DB = database
}
