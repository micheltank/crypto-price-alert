package infra

import (
	"encoding/json"
	"github.com/micheltank/crypto-price-alert/price-alert-job/internal/domain"
	infra "github.com/micheltank/crypto-price-alert/price-alert-job/internal/infra/kafka"
	"github.com/pkg/errors"
)

type NotificationBroker struct {
	kafkaProducer infra.IKafkaProducer
}

const SendEmailTopic = "send-email"

func NewNotificationBroker() (NotificationBroker, error) {
	kafkaProducer, err := infra.NewKafkaProducer()
	if err != nil {
		return NotificationBroker{}, errors.Wrap(err, "failed to create kafka producer")
	}
	return NotificationBroker{
		kafkaProducer: kafkaProducer,
	}, nil
}

func (n NotificationBroker) SendEmail(notification domain.EmailNotification) error {
	msg, err := json.Marshal(ConvertEmailNotificationFromDomain(notification))
	if err != nil {
		return errors.Wrap(err, "failed to marshal notification")
	}
	err = n.kafkaProducer.Publish(msg, SendEmailTopic, nil)
	if err != nil {
		return errors.Wrap(err, "failed to send email")
	}
	return nil
}
