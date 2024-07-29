// services/postgres.go
package services

import (
    "go-microservice/config"
    "go-microservice/models"
    "github.com/jinzhu/gorm"
    _ "github.com/lib/pq"
    "log"
)

var DB *gorm.DB

func InitPostgres(cfg config.Config) {
    var err error
    dsn := "host=" + cfg.DBHost + " user=" + cfg.DBUser + " password=" + cfg.DBPassword + " dbname=" + cfg.DBName + " port=" + cfg.DBPort + " sslmode=disable"
    DB, err = gorm.Open("postgres", dsn)
    if err != nil {
        log.Fatal(err)
    }

    DB.AutoMigrate(&models.Message{})
}

func GetDB() *gorm.DB {
    return DB
}
