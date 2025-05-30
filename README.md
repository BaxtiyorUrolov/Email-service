# 📧 Email Service

Bu loyiha orqali siz RabbitMQ yoki Kafka brokerlari yordamida email yuborish xizmati yaratishingiz mumkin.

## 🛠 Texnologiyalar

* Go (Golang)
* RabbitMQ
* Kafka
* Nats
* Gmail SMTP

---

## 📁 Loyihaning Tuzilishi

```
Email-service/
├── cmd/
│   ├── email-service/        # Consumer servis (email yuboradi)
│   └── producer/             # Email so'rovi jo'natuvchi (publisher)
├── internal/
│   ├── broker/               # RabbitMQ, Kafka va Nats brokerlar
│   ├── config/               # Konfiguratsiya yuklash
│   └── email/                # Email yuborish logikasi
├── .env.example              # Muqobil .env fayl
├── go.mod / go.sum
├── README.md
```

---

## ⚙️ Sozlamalar

`.env` faylini quyidagicha yarating:

```
RABBIT_URL=amqp://guest:guest@localhost:5672/
NATS_URL=nats://localhost:4222
#rabbit, kafka or nats
BROKER_TYPE=BrokerType
EMAIL_FROM=EmailFrom
EMAIL_PASS=EmailPass
```

> **Eslatma:** `EMAIL_PASS` uchun Gmail App Password ishlatilishi kerak.

---

## 🚀 Ishga Tushurish

### 1. Email Yuborish Servisini Ishga Tushurish (Consumer)

```bash
export $(cat .env) && go run cmd/email-service/main.go
```

### 2. Email So‘rov Jo‘natuvchi (Producer)

#### Kafka orqali:

```bash
export $(cat .env) && go run cmd/producer/main.go --broker=kafka
```

#### RabbitMQ orqali:

```bash
export $(cat .env) && go run cmd/producer/main.go --broker=rabbit
```

#### Nats orqali:

```bash
export $(cat .env) && go run cmd/producer/main.go --broker=nats
```

---

## 📦 Brokerlar

### Kafka

Agar Docker orqali Kafka o'rnatgan bo‘lsangiz:

```bash
docker run -d --name zookeeper -p 2181:2181 confluentinc/cp-zookeeper:latest

docker run -d --name kafka -p 9092:9092 \
  -e KAFKA_BROKER_ID=1 \
  -e KAFKA_ZOOKEEPER_CONNECT=172.17.0.2:2181 \
  -e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://172.17.0.3:9092 \
  -e KAFKA_LISTENER_SECURITY_PROTOCOL_MAP=PLAINTEXT:PLAINTEXT \
  -e KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR=1 \
  -e KAFKA_AUTO_CREATE_TOPICS_ENABLE=true \
  confluentinc/cp-kafka:latest
```

So‘ng brokerni ishga tushiring va 9092 port ochiq ekanligiga ishonch hosil qiling.

### RabbitMQ

```bash
docker pull rabbitmq:3-management
```

So‘ng brokerni ishga tushiring:

```bash
docker run -d -p 5672:5672 -p 15672:15672 --name rabbitmq rabbitmq:3-management
```

Web UI: [http://localhost:15672](http://localhost:15672)  (login: guest / guest)

### Nats

Agar Docker orqali Nats o'rnatmoqchi bo‘lsangiz:

```bash
docker pull nats:latest
```

So‘ng brokerni ishga tushiring.
---

## ✉️ Email Namuna

Email quyidagi ko‘rinishda yuboriladi:

```html
<h2 style="color: #4CAF50;">Assalomu alaykum!</h2>
<p>Men <b>Baxtiyor Urolov</b>, sizga broker orqali yuborilgan avtomatik emailni sinov tariqasida yuboryapman.</p>
<p>Service muvaffaqiyatli ishladi ✅</p>
```

---

## 👨🏻‍💻 Muallif

Baxtiyor Urolov
🔗 GitHub: [@BaxtiyorUrolov](https://github.com/BaxtiyorUrolov)

Agar sizga loyiha yoqqan bo‘lsa ⭐ berishni unutmang!
