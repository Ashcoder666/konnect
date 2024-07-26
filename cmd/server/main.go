package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ashcoder666/konnect/models"
	"github.com/ashcoder666/konnect/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func main() {
	loadEnv()
	H := gin.Default()
	// err := connectDB()
	routes.AuthRoutes(H)

	models.ConnectDatabase()
	fmt.Println("Successfully connected to the database")
	fmt.Println(os.Getenv("PORT"))
	H.Run()
}

// func errHandler(err error) {
// 	log.Fatal(err)
// }

// func connectDB() error {
// 	// err := godotenv.Load()

// 	// errHandler(fmt.Errorf("env not loaded"))

// 	fmt.Println(os.Getenv("DATBASE_USERNAME"), "jjj")

// 	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:33060)/%s?charset=utf8mb4&parseTime=True&loc=Local", "root", "qwerty", "konnect")
// 	fmt.Println(dsn)
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatal("Failed to connect to MySQL:", err)

// 	}

// 	sqlDB, err := db.DB()
// 	if err != nil {
// 		log.Fatal("Failed to get SQL DB instance:", err)
// 	}
// 	defer sqlDB.Close()

// 	// Test the database connection
// 	err = sqlDB.Ping()
// 	if err != nil {
// 		log.Fatal("Error connecting to database:", err)
// 		return err
// 	}
// 	fmt.Println("Connected to MySQL!")

// 	return nil
// }
