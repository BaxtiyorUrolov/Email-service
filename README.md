# 📧 Email Service

Bu loyiha orqali siz RabbitMQ yoki Kafka brokerlari yordamida email yuborish xizmati yaratishingiz mumkin.

## 🛠 Texnologiyalar

* Go (Golang)
* RabbitMQ
* Kafka
* Gmail SMTP

---

## 📁 Loyihaning Tuzilishi

```
Email-service/
├── cmd/
│   ├── email-service/        # Consumer servis (email yuboradi)
│   └── producer/             # Email so'rovi jo'natuvchi (publisher)
├── internal/
│   ├── broker/               # RabbitMQ va Kafka brokerlar
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
BROKER_TYPE=rabbit # yoki kafka
EMAIL_FROM=youremail@gmail.com
EMAIL_PASS=yourapppassword
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

---

## 📦 Brokerlar

### Kafka

Agar Docker orqali Kafka o'rnatgan bo‘lsangiz:

```bash
docker pull bitnami/kafka
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
