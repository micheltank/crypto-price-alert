package application

import "github.com/micheltank/crypto-price-alert/notification-service/internal/domain"

type NotificationService interface {
	Send(emailNotification domain.EmailNotification) error
}
