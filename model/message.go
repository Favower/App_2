// models/message.go
package models

import "time"

type Message struct {
    ID        uint      `gorm:"primaryKey"`
    Content   string    `gorm:"type:text"`
    Status    string    `gorm:"type:varchar(100)"`
    CreatedAt time.Time
}
