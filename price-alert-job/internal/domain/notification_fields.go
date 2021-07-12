package domain

type EmailNotificationFields struct {
	triggerPrice float64
	currentPrice float64
	currency     string
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