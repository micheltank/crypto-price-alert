package app

import (
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/micheltank/crypto-price-alert/notification-service/internal/application"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type EmailListener struct {
	consumer            *kafka.Consumer
	notificationService application.NotificationService
	inShutdown          bool
}

const Topic = "send-email"

func NewEmailListener(service application.NotificationService) (*EmailListener, error) {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
		"client.id":         "notification-notificationService-consumer",
		"group.id":          "notification-notificationService-group",
	}
	consumer, err := kafka.NewConsumer(configMap)
	if err != nil {
		return nil, err
	}
	s := EmailListener{
		consumer:            consumer,
		notificationService: service,
		inShutdown:          false,
	}
	return &s, nil
}

func (e *EmailListener) Listen() error {
	err := e.consumer.Subscribe(Topic, nil)
	if err != nil {
		return errors.Wrap(err, "failed to subscribe topic")
	}
	for {
		logrus.Debug("Reading message")
		msg, err := e.consumer.ReadMessage(-1)
		if err != nil {
			logrus.WithError(err).Error("failed to read email message")
			continue
		}
		var notification EmailNotification
		err = json.Unmarshal(msg.Value, &notification)
		if err != nil {
			logrus.WithError(err).Error("failed to unmarshall email notification")
			continue
		}
		err = e.notificationService.Send(ConvertEmailNotificationToDomain(notification))
		if err != nil {
			logrus.WithError(err).Error("failed to send email")
			continue
		}
		if e.inShutdown {
			break
		}
	}
	return nil
}

func (e *EmailListener) Shutdown() {
	// TODO: implement wait
	e.inShutdown = true
}
