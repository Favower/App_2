// services/kafka.go
package services

import (
    "context"
    "log"
    "go-microservice/config"
    "github.com/segmentio/kafka-go"
)

var KafkaWriter *kafka.Writer
var KafkaReader *kafka.Reader

func InitKafka(cfg config.Config) {
    KafkaWriter = kafka.NewWriter(kafka.WriterConfig{
        Brokers: []string{cfg.KafkaURL},
        Topic:   "messages",
    })

    KafkaReader = kafka.NewReader(kafka.ReaderConfig{
        Brokers: []string{cfg.KafkaURL},
        Topic:   "messages",
        GroupID: "message-group",
    })
}

func SendMessage(msg string) error {
    err := KafkaWriter.WriteMessages(context.Background(),
        kafka.Message{
            Value: []byte(msg),
        },
    )
    return err
}

func ConsumeMessages() {
    for {
        msg, err := KafkaReader.ReadMessage(context.Background())
        if err != nil {
            log.Fatal(err)
        }
        log.Printf("Message: %s", string(msg.Value))
    }
}
