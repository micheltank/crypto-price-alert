package repository

import (
	"github.com/micheltank/crypto-price-alert/price-alert-service/internal/domain"
)

type Alert struct {
	Id        int64
	Email     string
	Price     float64
	Coin      string
	Direction domain.Direction
}

func (alert Alert) ToDomain() domain.Alert {
	return domain.NewAlertFromRepository(alert.Id, alert.Email, alert.Price, alert.Coin, alert.Direction)
}
