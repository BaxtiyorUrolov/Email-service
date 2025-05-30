package main

import (
	"email-service/internal/broker"
	"email-service/internal/config"
	"email-service/internal/email"
	"encoding/json"
	"fmt"
	"log"
)

type EmailPayload struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	sender := email.NewSender(
		cfg.EmailFrom,
		cfg.EmailPass,
		"smtp.gmail.com",
		"587",
	)

	if cfg.BrokerType == "rabbit" {
		runRabbitMQ(cfg.RabbitMQURL, sender)
	} else if cfg.BrokerType == "kafka" {
		runKafka(sender)
	} else if cfg.BrokerType == "nats" {
		runNats(cfg.NatsURL, "email-subject", sender) // NATS URL va subject
	} else {
		log.Fatalf("Noto‘g‘ri broker turi: %s", cfg.BrokerType)
	}
}

func runRabbitMQ(rabbitURL string, sender *email.EmailSender) {
	consumer, err := broker.NewConsumer(rabbitURL, "email_queue")
	if err != nil {
		log.Fatal(err)
	}
	defer consumer.Close()

	err = consumer.ConsumeMessages(func(body []byte) {
		var payload EmailPayload
		if err := json.Unmarshal(body, &payload); err != nil {
			log.Println("invalid payload:", err)
			return
		}
		if err := sender.Send(payload.To, payload.Subject, payload.Body); err != nil {
			log.Println("email send failed:", err)
		} else {
			fmt.Println("Email sent to", payload.To)
		}
	})

	if err != nil {
		log.Fatal(err)
	}
	select {}
}

func runKafka(sender *email.EmailSender) {
	broker.StartKafkaListener(sender)
}

func runNats(natsURL, subject string, sender *email.EmailSender) {
	consumer, err := broker.NewNatsConsumer(natsURL, subject)
	if err != nil {
		log.Fatal(err)
	}
	defer consumer.Close()

	err = consumer.ConsumeMessages(func(body []byte) {
		var payload EmailPayload
		if err := json.Unmarshal(body, &payload); err != nil {
			log.Println("❌ JSON xato:", err)
			return
		}
		if err := sender.Send(payload.To, payload.Subject, payload.Body); err != nil {
			log.Println("❌ Email yuborilmadi:", err)
		} else {
			fmt.Println("✅ Email sent to", payload.To)
		}
	})

	if err != nil {
		log.Fatal(err)
	}
	select {} // Dastur to‘xtamasligi uchun
}
