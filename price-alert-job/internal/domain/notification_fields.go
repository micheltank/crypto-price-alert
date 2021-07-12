package domain

type EmailNotificationFields struct {
	triggerPrice float64
	currentPrice float64
	currency     string
}

func NewEmailNotificationFields(triggerPrice, currentPrice float64, currency string) EmailNotificationFields {
	return EmailNotificationFields{
		triggerPrice: triggerPrice,
		currentPrice: currentPrice,
		currency:     currency,
	}
}

func (e EmailNotificationFields) GetTriggerPrice() float64 {
	return e.triggerPrice
}

func (e EmailNotificationFields) GetCurrentPrice() float64 {
	return e.currentPrice
}

func (e EmailNotificationFields) GetCurrency() string {
	return e.currency
}