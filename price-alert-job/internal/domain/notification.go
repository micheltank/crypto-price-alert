package domain

import "github.com/google/uuid"

type EmailNotification struct {
	notificationId string
	email          string
	template       EmailNotificationTemplate
	fields         EmailNotificationFields
}

type EmailNotificationTemplate string

const (
	TemplateAbove EmailNotificationTemplate = "PRICE_CHANGE_ABOVE"
	TemplateBelow EmailNotificationTemplate = "PRICE_CHANGE_BELOW"
)

func (e EmailNotification) GetNotificationId() string {
	return e.notificationId
}

func (e EmailNotification) GetEmail() string {
	return e.email
}

func (e EmailNotification) GetTemplate() EmailNotificationTemplate {
	return e.template
}

func (e EmailNotification) GetFields() EmailNotificationFields {
	return e.fields
}

func NewEmailNotification(email string, template EmailNotificationTemplate, fields EmailNotificationFields) EmailNotification {
	return EmailNotification{
		notificationId: uuid.New().String(),
		email:          email,
		template:       template,
		fields:         fields,
	}
}

func NewEmailNotificationFields(triggerPrice, currentPrice float64, currency string) EmailNotificationFields {
	return EmailNotificationFields{
		triggerPrice: triggerPrice,
		currentPrice: currentPrice,
		currency:     currency,
	}
}
