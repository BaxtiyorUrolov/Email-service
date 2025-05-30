package broker

import (
	"context"
	"email-service/internal/email"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"log"
)

type EmailMessage struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

func StartKafkaListener(sender *email.EmailSender) {
	brokers := []string{"172.17.0.3:9092"}
	topic := "email-topic"

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		Topic:   topic,
		GroupID: "email-consumer-group",
	})
	log.Println("📥 Kafka listener ishga tushdi...")

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Println("❌ Kafka o‘qishda xato:", err)
			continue
		}

		var emailMsg EmailMessage
		err = json.Unmarshal(msg.Value, &emailMsg)
		if err != nil {
			log.Println("❌ JSON xato:", err)
			continue
		}

		err = sender.Send(emailMsg.To, emailMsg.Subject, emailMsg.Body)
		if err != nil {
			log.Println("❌ Email yuborilmadi:", err)
		} else {
			log.Printf("✅ Email yuborildi: %s", emailMsg.To)
		}
	}
}
