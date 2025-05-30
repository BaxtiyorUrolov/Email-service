package broker

import (
	"log"

	"github.com/nats-io/nats.go"
)

type NatsConsumer struct {
	conn    *nats.Conn
	subject string
}

func NewNatsConsumer(natsURL, subject string) (*NatsConsumer, error) {
	nc, err := nats.Connect(natsURL)
	if err != nil {
		return nil, err
	}

	return &NatsConsumer{
		conn:    nc,
		subject: subject,
	}, nil
}

func (n *NatsConsumer) ConsumeMessages(handler func([]byte)) error {
	_, err := n.conn.Subscribe(n.subject, func(m *nats.Msg) {
		handler(m.Data)
	})
	if err != nil {
		return err
	}

	log.Println("📥 NATS listener ishga tushdi...")
	return nil
}

func (n *NatsConsumer) Close() {
	if n.conn != nil {
		n.conn.Close()
	}
}
