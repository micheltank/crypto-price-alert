package domain

type NotificationBroker interface {
	SendEmail(notification EmailNotification) error
}
