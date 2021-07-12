package domain

import (
	"github.com/pkg/errors"
	"regexp"
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

func (a *Alert) SetPrice(price float64) {
	a.price = price
}

func (a *Alert) SetCoin(coin string) {
	a.coin = coin
}

func (a *Alert) GetCoin() string {
	return a.coin
}

func (a *Alert) GetDirection() Direction {
	return a.direction
}

func (a *Alert) SetDirection(direction Direction) {
	a.direction = direction
}

func (a *Alert) ShouldAlertAtTheGivenPrice(currentPrice float64) bool {
	if a.direction == DirectionAbove && a.price > currentPrice {
		return true
	}
	if a.direction == DirectionBelow && a.price < currentPrice {
		return true
	}
	return false
}

func (a *Alert) validate() error {
	pattern := "^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
	match, _ := regexp.MatchString(pattern, a.email)
	if !match {
		return errors.New("invalid email")
	}
	if a.direction == DirectionAbove && a.price <= 0 {
		return errors.New("above price cannot be equal or less than zero")
	}
	if a.direction == DirectionBelow && a.price <= 0 {
		return errors.New("below price cannot be equal or less than zero")
	}
	return nil
}
