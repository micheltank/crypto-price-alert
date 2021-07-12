package domain

type AlertApi interface {
	GetAlerts(coin string, price float64) (Alerts, error)
}
