package email

import (
	"net/smtp"
)

type EmailSender struct {
	From     string
	Password string
	SMTPHost string
	SMTPPort string
}

func NewSender(from, password, host, port string) *EmailSender {
	return &EmailSender{from, password, host, port}
}

func (s *EmailSender) Send(to, subject, body string) error {
	msg := "MIME-Version: 1.0\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\"\r\n" +
		"From: Baxtiyor Urolov <" + s.From + ">\r\n" +
		"To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n\r\n" +
		body

	auth := smtp.PlainAuth("", s.From, s.Password, s.SMTPHost)

	return smtp.SendMail(s.SMTPHost+":"+s.SMTPPort, auth, s.From, []string{to}, []byte(msg))
}
