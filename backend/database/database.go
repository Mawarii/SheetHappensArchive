package database

import (
	"fmt"
	"os"
	"sheethappens/backend/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB
var err error

func DatabaseInit() {
	host := os.Getenv("POSTGRES_HOST")
	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	port := os.Getenv("POSTGRES_PORT")
	timezone := os.Getenv("TZ")

	dbString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", host, username, password, dbName, port, timezone)
	database, err = gorm.Open(postgres.Open(dbString), &gorm.Config{})

	database.AutoMigrate(&model.User{})
	database.AutoMigrate(&model.Character{})

	if err != nil {
		panic(err)
	}
}

func DB() *gorm.DB {
	return database
}
