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
	routes.UserRoutes(H)

	models.ConnectDatabase()
	fmt.Println("Successfully connected to the database")
	fmt.Println(os.Getenv("PORT"))
	H.Run()
}
