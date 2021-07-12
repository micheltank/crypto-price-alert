package application

import (
	"github.com/micheltank/crypto-price-alert/price-alert-service/internal/domain"
	wErrors "github.com/pkg/errors"
)

type IService interface {
	GetAlerts(email string) (domain.Alerts, error)
	CreateAlert(command domain.CreateAlertCommand) (domain.Alert, error)
	GetAlertsAtPrice(coin string, price float64) (domain.Alerts, error)
}

type Service struct {
	repository domain.Repository
}

func NewService(repository domain.Repository) *Service {
	s := Service{
		repository: repository,
	}
	return &s
}

func (s *Service) GetAlerts(email string) (domain.Alerts, error) {
	keys, err := s.repository.List(email)
	if err != nil {
		return nil, err
	}
	return keys, nil
}

func (s *Service) CreateAlert(command domain.CreateAlertCommand) (domain.Alert, error) {
	alert, err := domain.NewAlert(command.Email, command.Price, command.Coin, command.Direction)
	if err != nil {
		return domain.Alert{}, domain.NewError("failed to create alert", "error.validation", err.Error())
	}
	alert, err = s.repository.Create(alert)
	if err != nil {
		return domain.Alert{}, wErrors.Wrap(err, "failed to create alert on database")
	}
	return alert, nil
}

func (s *Service) GetAlertsAtPrice(coin string, price float64) (domain.Alerts, error) {
	keys, err := s.repository.GetAlertsAtPrice(coin, price)
	if err != nil {
		return nil, err
	}
	return keys, nil
}