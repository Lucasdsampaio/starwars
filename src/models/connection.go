package models

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDataBase() {
	dsn := fmt.Sprintf("host=%s port=5432 user=%s dbname=%s password=%s", os.Getenv("HOST"), os.Getenv("DATABASE_USER"), os.Getenv("DATABASE"), os.Getenv("PASSWORD"))
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Planet{})
	DB = database
}
