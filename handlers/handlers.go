// handlers/handlers.go
package handlers

import (
    "net/http"
    "go-microservice/models"
    "go-microservice/services"
    "github.com/gin-gonic/gin"
    "log"
)

func CreateMessage(c *gin.Context) {
    var msg models.Message
    if err := c.ShouldBindJSON(&msg); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    db := services.GetDB()
    msg.Status = "pending"
    db.Create(&msg)

    if err := services.SendMessage(msg.Content); err != nil {
        log.Println("Failed to send message to Kafka", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send message to Kafka"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Message created successfully"})
}

func GetStatistics(c *gin.Context) {
    var count int64
    db := services.GetDB()
    db.Model(&models.Message{}).Where("status = ?", "processed").Count(&count)

    c.JSON(http.StatusOK, gin.H{"processed_messages": count})
}
