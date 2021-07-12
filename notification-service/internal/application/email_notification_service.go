package application

import (
	"fmt"
	"github.com/leekchan/accounting"
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

	subject, body := s.buildContent(notification)
	err := s.service.Send(notification.GetEmail(), subject, body)
	if err != nil {
		return errors.Wrap(err, "failed to send email notification")
	}

	logrus.Debugf("Email sent to %s", notification.GetEmail())
	return nil
}

func (s *EmailNotificationService) buildContent(notification domain.EmailNotification) (string, string) {
	var subject, body string
	ac := accounting.Accounting{Symbol: fmt.Sprintf("%s ", notification.GetFields().GetCurrency()), Precision: 2}

	switch notification.GetTemplate() {
	case domain.TemplateAbove:
		{
			subject = "The price went up!"
			body = fmt.Sprintf(`Your alert for price above %s
			The price is now %s`, ac.FormatMoneyFloat64(notification.GetFields().GetTriggerPrice()), ac.FormatMoneyFloat64(notification.GetFields().GetCurrentPrice()))
		}
	case domain.TemplateBelow:
		{
			subject = "The price has gone down!"
			body = fmt.Sprintf(`Your alert for price below %s
			The price is now %s`, ac.FormatMoneyFloat64(notification.GetFields().GetTriggerPrice()), ac.FormatMoneyFloat64(notification.GetFields().GetCurrentPrice()))
		}
	}
	return subject, body
}