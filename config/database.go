package config

import (
	"github.com/IkhsanDzul/crud-golang/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=ikhsan24 dbname=crud-go port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
		// println("Failed to connect to the database")
	} else {
		println("Database connected successfully")
	}
}

func Migrate() {
	err := DB.AutoMigrate(&models.Product{}, &models.User{})
	if err != nil {
		panic(err)
	}
	println("Database migrated successfully")
}