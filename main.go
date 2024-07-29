// main.go
package main

import (
    "go-microservice/config"
    "go-microservice/handlers"
    "go-microservice/services"
    "github.com/gin-gonic/gin"
)

func main() {
    cfg := config.LoadConfig()
    
    services.InitPostgres(cfg)
    services.InitKafka(cfg)

    r := gin.Default()
    r.POST("/messages", handlers.CreateMessage)
    r.GET("/statistics", handlers.GetStatistics)

    go services.ConsumeMessages()

    r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
