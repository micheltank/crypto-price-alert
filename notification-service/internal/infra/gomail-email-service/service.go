package infra

import (
	"crypto/tls"
	"fmt"
	"github.com/micheltank/crypto-price-alert/notification-service/internal/domain"
	infraConfig "github.com/micheltank/crypto-price-alert/notification-service/internal/infra/config"
	"github.com/pkg/errors"
	gomail "gopkg.in/mail.v2"
)

type GomailEmailService struct {
	emailSender   string
	emailPassword string
}

func NewGomailEmailService(config infraConfig.Environment) domain.EmailService {
	return &GomailEmailService{
		emailSender:   config.EmailSender,
		emailPassword: config.EmailPassword,
	}
}

func (g *GomailEmailService) Send(to, subject, body string) error {
	m := gomail.NewMessage()

	m.SetHeader("From", fmt.Sprintf("Crypto Price Alert <%s>", g.emailSender))
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	d := gomail.NewDialer("smtp.gmail.com", 587, g.emailSender, g.emailPassword)

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		return errors.Wrap(err, "failed to send email")
	}

	return nil
}
