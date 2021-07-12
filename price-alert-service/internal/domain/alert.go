package domain

import (
	"regexp"
)

var (
	SupportedCoins = []string{"BTC", "ETH", "BNB"}
)

type Alert struct {
	id        int64
	email     string
	price     float64
	coin      string
	direction Direction
}

type Alerts []Alert

func NewAlert(email string, price float64, coin string, direction Direction) (Alert, error) {
	alert := Alert{
		email:     email,
		price:     price,
		coin:      coin,
		direction: direction,
	}
	err := alert.validate()
	if err != nil {
		return Alert{}, err
	}
	return alert, nil
}

func NewAlertFromRepository(id int64, email string, price float64, coin string, direction Direction) Alert {
	alert := Alert{
		id:        id,
		email:     email,
		price:     price,
		coin:      coin,
		direction: direction,
	}
	return alert
}

func (a *Alert) GetId() int64 {
	return a.id
}

func (a *Alert) GetEmail() string {
	return a.email
}

func (a *Alert) GetPrice() float64 {
	return a.price
}

func (a *Alert) GetCoin() string {
	return a.coin
}

func (a *Alert) GetDirection() Direction {
	return a.direction
}

func (a *Alert) validate() error {
	pattern := "^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
	match, _ := regexp.MatchString(pattern, a.email)
	if !match {
		return ErrInvalidEmail
	}
	if a.direction == DirectionAbove && a.price <= 0 {
		return ErrInvalidPriceAbove
	}
	if a.direction == DirectionBelow && a.price <= 0 {
		return ErrInvalidPriceBelow
	}
	if !a.coinSupported() {
		return ErrUnsupportedCoin
	}
	return nil
}

func (a *Alert) coinSupported() bool {
	for _, coin := range SupportedCoins {
		if coin == a.coin {
			return true
		}
	}
	return false
}
