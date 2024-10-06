package config

import (
    "fmt"
    "log"
    "os"
    "time"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "github.com/joho/godotenv"
)

var DB *gorm.DB

func ConnectDatabase() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")

    // Create the PostgreSQL DSN (Data Source Name)
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
        dbHost, dbUser, dbPassword, dbName, dbPort)

    // Retry logic for connecting to the database
    var database *gorm.DB
    maxAttempts := 10
    for attempts := 0; attempts < maxAttempts; attempts++ {
        database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
        if err == nil {
            log.Println("Successfully connected to the database!")
            DB = database
            return
        }

        log.Printf("Failed to connect to database (attempt %d/%d): %v", attempts+1, maxAttempts, err)
        time.Sleep(5 * time.Second) 
    }

    log.Fatal("Max attempts reached. Could not connect to the database:", err)
}
