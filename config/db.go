package config

import (
	"assignment-2/app/model"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "host=localhost user=postgres password=postgre dbname=postgres port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	DB.Debug().AutoMigrate(&model.Order{}, &model.Item{})

	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	fmt.Println("Database connection established")
}

func GetDB() *gorm.DB {
	return DB
}
