package app

import (
	"github.com/micheltank/crypto-price-alert/notification-service/internal/domain"
)

type EmailNotification struct {
	NotificationId string                  `json:"notificationId"`
	Email          string                  `json:"email"`
	Template       string                  `json:"template"`
	Fields         EmailNotificationFields `json:"fields"`
}

type EmailNotificationFields struct {
	TriggerPrice float64 `json:"triggerPrice"`
	CurrentPrice float64 `json:"currentPrice"`
	Currency     string  `json:"currency"`
}

func ConvertEmailNotificationToDomain(notification EmailNotification) domain.EmailNotification {
	return domain.NewEmailNotification(
		notification.NotificationId,
		notification.Email,
		domain.EmailNotificationTemplate(notification.Template),
		domain.NewEmailNotificationFields(
			notification.Fields.TriggerPrice,
			notification.Fields.CurrentPrice,
			notification.Fields.Currency,
		))
}
