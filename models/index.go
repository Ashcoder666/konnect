package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"localhost",
		"54320",
		"postgres",
		"my_password",
		"konnect",
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	err = database.AutoMigrate(&User{}, &Tempotps{}, &Partner{})

	if err != nil {
		panic(err)
	}

	DB = database
}
