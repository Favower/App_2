// config/config.go
package config

import (
    "os"
    "log"
    "github.com/joho/godotenv"
)

type Config struct {
    DBUser     string
    DBPassword string
    DBName     string
    DBHost     string
    DBPort     string
    KafkaURL   string
}

func LoadConfig() Config {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    return Config{
        DBUser:     os.Getenv("DB_USER"),
        DBPassword: os.Getenv("DB_PASSWORD"),
        DBName:     os.Getenv("DB_NAME"),
        DBHost:     os.Getenv("DB_HOST"),
        DBPort:     os.Getenv("DB_PORT"),
        KafkaURL:   os.Getenv("KAFKA_URL"),
    }
}
