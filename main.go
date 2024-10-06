package main

import (
	"go-crud-app/config"
	"go-crud-app/models"
	"go-crud-app/routes"
	"log"
    "github.com/joho/godotenv"
    
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    } else {
        log.Println(".env file loaded successfully")
    }

    config.ConnectDatabase()

    config.DB.AutoMigrate(&models.User{})

    r := routes.SetupRouter()
    r.Run(":8081")
}
