package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	DB, err = gorm.Open(postgres.Open("host=localhost user=postgres dbname=crud-go port=5432 sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
}