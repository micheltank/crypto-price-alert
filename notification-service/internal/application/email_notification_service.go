package application

import (
	"github.com/micheltank/crypto-price-alert/notification-service/internal/domain"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type EmailNotificationService struct {
	service domain.EmailService
}

func NewEmailNotificationService(service domain.EmailService) *EmailNotificationService {
	s := EmailNotificationService{
		service: service,
	}
	return &s
}

func (s *EmailNotificationService) Send(notification domain.EmailNotification) error {
	logrus.Debugf("Sending email to %s with template %s", notification.GetEmail(), notification.GetTemplate())

	subject, body := notification.BuildContent()
	err := s.service.Send(notification.GetEmail(), subject, body)
	if err != nil {
		return errors.Wrap(err, "failed to send email notification")
	}

	logrus.Debugf("Email sent to %s", notification.GetEmail())
	return nil
}

