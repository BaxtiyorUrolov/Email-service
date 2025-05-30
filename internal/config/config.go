package config

import (
	"context"
	"github.com/sethvargo/go-envconfig"
)

// Config holds the application configuration
type Config struct {
	RabbitMQURL string `env:"RABBIT_URL, default=amqp://guest:guest@rabbitmq:5672/"`
	NatsURL     string `env:"NATS_URL, default=natsURL"`
	BrokerType  string `env:"BROKER_TYPE, default=kafka"`
	EmailFrom   string `env:"EMAIL_FROM, default=EMAIL_FROM"`
	EmailPass   string `env:"EMAIL_PASS, default=EMAIL_PASS"`
}

// Load loads environment variables into the Config struct using go-envconfig
func Load() (*Config, error) {
	var cfg Config
	// Use ProcessWith with OsLookuper as an option
	if err := envconfig.ProcessWith(context.TODO(), &envconfig.Config{
		Target:   &cfg,
		Lookuper: envconfig.OsLookuper(),
	}); err != nil {
		return nil, err
	}
	return &cfg, nil
}
