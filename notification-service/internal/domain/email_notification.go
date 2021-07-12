package domain

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

func NewEmailNotification(notificationId string, email string, template EmailNotificationTemplate, fields EmailNotificationFields) EmailNotification {
	return EmailNotification{
		notificationId: notificationId,
		email:          email,
		template:       template,
		fields:         fields,
	}
}