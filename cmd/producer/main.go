package main

import (
	"context"
	"email-service/internal/config"
	"encoding/json"
	"flag"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/streadway/amqp"
)

type EmailPayload struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

func main() {
	// Broker flag (rabbit yoki kafka)
	broker := flag.String("broker", "rabbit", "Email broker: rabbit or kafka")
	flag.Parse()

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	payload := EmailPayload{
		To:      "urolov0208@gmail.com",
		Subject: "Salom",
		Body: `<h2 style="color: #4CAF50;">Assalomu alaykum!</h2>
<p>Men <b>Baxtiyor Urolov</b>, sizga <b>` + *broker + `</b> orqali yuborilgan avtomatik emailni sinov tariqasida yuboryapman.</p>
<p>Service muvaffaqiyatli ishladi ✅</p>`,
	}

	switch *broker {
	case "rabbit":
		sendViaRabbit(cfg.RabbitMQURL, payload)
	case "kafka":
		sendViaKafka(payload)
	default:
		log.Fatal("Noto‘g‘ri broker: rabbit yoki kafka bo‘lishi kerak")
	}
}

// ✅ RabbitMQ orqali yuborish
func sendViaRabbit(url string, payload EmailPayload) {
	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	q, _ := ch.QueueDeclare("email_queue", true, false, false, false, nil)

	body, _ := json.Marshal(payload)

	err = ch.Publish("", q.Name, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        body,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("📤 RabbitMQ orqali email yuborildi.")
}

// ✅ Kafka orqali yuborish
func sendViaKafka(payload EmailPayload) {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "email-topic",
		Balancer: &kafka.LeastBytes{},
	})
	defer writer.Close()

	body, _ := json.Marshal(payload)

	err := writer.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte("email"),
		Value: body,
		Time:  time.Now(),
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("📤 Kafka orqali email yuborildi.")
}
