package database

import (
	"example.com/task_manager/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	//gorm.Open(..., &gorm.Config{}) opens the database with default GORM configuration and returns a *gorm.DB and an error.
	db, err := gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	//db.AutoMigrate(&models.User{}, &models.Task{})
	db.AutoMigrate(&models.User{}, &models.Task{})
	DB = db

}
