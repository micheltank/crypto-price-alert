package infra

import "github.com/micheltank/crypto-price-alert/price-alert-job/internal/domain"

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

func ConvertEmailNotificationFromDomain(notification domain.EmailNotification) EmailNotification {
	return EmailNotification{
		NotificationId: notification.GetNotificationId(),
		Email:          notification.GetEmail(),
		Template:       string(notification.GetTemplate()),
		Fields: EmailNotificationFields{
			TriggerPrice: notification.GetFields().GetTriggerPrice(),
			CurrentPrice: notification.GetFields().GetCurrentPrice(),
			Currency:     notification.GetFields().GetCurrency(),
		},
	}
}
