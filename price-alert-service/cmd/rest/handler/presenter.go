package handler

import (
	"github.com/micheltank/crypto-price-alert/price-alert-service/internal/domain"
)

type CreateAlertRequest struct {
	Email     string           `json:"email" binding:"required"`
	Price     float64          `json:"price" binding:"required"`
	Coin      string           `json:"coin" binding:"required"`
	Direction domain.Direction `json:"direction" binding:"required"`
}

func (r *CreateAlertRequest) toCommand() domain.CreateAlertCommand {
	return domain.CreateAlertCommand{
		Email:     r.Email,
		Price:     r.Price,
		Coin:      r.Coin,
		Direction: r.Direction,
	}
}

type AlertResponse struct {
	Id        int64            `json:"id"`
	Email     string           `json:"email"`
	Price     float64          `json:"price"`
	Coin      string           `json:"coin"`
	Direction domain.Direction `json:"direction"`
}

func NewAlertResponse(alert domain.Alert) AlertResponse {
	return AlertResponse{
		Id:        alert.GetId(),
		Email:     alert.GetEmail(),
		Price:     alert.GetPrice(),
		Coin:      alert.GetCoin(),
		Direction: alert.GetDirection(),
	}
}

type AlertsResponse []AlertResponse

func NewGetAlertsResponse(alerts domain.Alerts) AlertsResponse {
	var alertsResponse AlertsResponse
	for i := 0; i < len(alerts); i++ {
		alertsResponse = append(alertsResponse, NewAlertResponse(alerts[i]))
	}
	if alertsResponse == nil {
		return AlertsResponse{}
	}
	return alertsResponse
}
