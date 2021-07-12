package domain

type Repository interface {
	Create(alert Alert) (Alert, error)
	List(email string) (Alerts, error)
	GetAlertsAtPrice(coin string, price float64) (Alerts, error)
}
