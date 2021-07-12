package domain

import (
	"fmt"
	"github.com/leekchan/accounting"
)

type EmailNotification struct {
	notificationId string
	email          string
	template       EmailNotificationTemplate
	fields         EmailNotificationFields
}

func NewEmailNotification(notificationId string, email string, template EmailNotificationTemplate, fields EmailNotificationFields) EmailNotification {
	return EmailNotification{
		notificationId: notificationId,
		email:          email,
		template:       template,
		fields:         fields,
	}
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

func (e *EmailNotification) BuildContent() (string, string) {
	var subject, body string
	ac := accounting.Accounting{Symbol: fmt.Sprintf("%s ", e.GetFields().GetCurrency()), Precision: 2}

	switch e.GetTemplate() {
	case TemplateAbove:
		{
			subject = "The price went up!"
			body = fmt.Sprintf(`Your alert for price above %s
			The price is now %s`, ac.FormatMoneyFloat64(e.GetFields().GetTriggerPrice()), ac.FormatMoneyFloat64(e.GetFields().GetCurrentPrice()))
		}
	case TemplateBelow:
		{
			subject = "The price has gone down!"
			body = fmt.Sprintf(`Your alert for price below %s
			The price is now %s`, ac.FormatMoneyFloat64(e.GetFields().GetTriggerPrice()), ac.FormatMoneyFloat64(e.GetFields().GetCurrentPrice()))
		}
	}
	return subject, body
}