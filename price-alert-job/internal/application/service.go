package application

import (
	"github.com/micheltank/crypto-price-alert/price-alert-job/internal/domain"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type IService interface {
	Execute() error
}

type Service struct {
	priceApi           domain.PriceApi
	alertApi           domain.AlertApi
	notificationBroker domain.NotificationBroker
}

func NewService(priceApi domain.PriceApi, alertApi domain.AlertApi, notificationBroker domain.NotificationBroker) *Service {
	s := Service{
		priceApi:           priceApi,
		alertApi:           alertApi,
		notificationBroker: notificationBroker,
	}
	return &s
}

func (s *Service) Execute() error {
	logrus.Info("Executing job")
	for _, coin := range domain.SupportedCoins {
		price, err := s.priceApi.GetPrice(coin)
		if err != nil {
			return errors.Wrap(err, "failed to fetch price")
		}
		alerts, err := s.alertApi.GetAlerts(coin, price.USD)
		if err != nil {
			return errors.Wrap(err, "failed to get alerts")
		}
		for _, alert := range alerts {
			template := domain.TemplateBelow
			if alert.GetDirection() == domain.DirectionAbove {
				template = domain.TemplateAbove
			}
			fields := domain.NewEmailNotificationFields(alert.GetPrice(), price.USD, "USD")
			notification := domain.NewEmailNotification(alert.GetEmail(), template, fields)
			err := s.notificationBroker.SendEmail(notification)
			logrus.Debugf("Sending email notification %s to %s at price %f", notification.GetNotificationId(), alert.GetEmail(), price.USD)
			if err != nil {
				return errors.Wrap(err, "failed to get alerts")
			}
		}
	}
	return nil
}