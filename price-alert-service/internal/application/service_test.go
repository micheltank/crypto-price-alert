package application_test

import (
	"github.com/golang/mock/gomock"
	"github.com/micheltank/crypto-price-alert/price-alert-service/internal/application"
	mock_domain "github.com/micheltank/crypto-price-alert/price-alert-service/internal/application/mock"
	"github.com/micheltank/crypto-price-alert/price-alert-service/internal/domain"
	. "github.com/onsi/gomega"
	"testing"
	"time"
)

func TestAlertService(t *testing.T) {
	ctrl := gomock.NewController(t)
	repository := mock_domain.NewMockRepository(ctrl)

	repository.EXPECT().
		Create(gomock.Any()).
		DoAndReturn(func(alert domain.Alert) (domain.Alert, error) {
			time.Sleep(1 * time.Second)
			return domain.NewAlertFromRepository(1, alert.GetEmail(), alert.GetPrice(), alert.GetCoin(), alert.GetDirection()), nil
		}).
		AnyTimes()

	t.Run("Regular alert creation", func(t *testing.T) {
		g := NewGomegaWithT(t)

		service := application.NewService(repository)
		expectedEmail := "john@gmail.com"
		command := domain.CreateAlertCommand{
			Email:     expectedEmail,
			Price:     100,
			Direction: domain.DirectionAbove,
			Coin:      "BTC",
		}
		alert, err := service.CreateAlert(command)
		g.Expect(err).Should(
			Not(HaveOccurred()))
		g.Expect(alert).Should(
			Not(BeNil()))
		g.Expect(alert.GetEmail()).Should(
			Equal(expectedEmail))
	})
	t.Run("Invalid alert creation", func(t *testing.T) {
		g := NewGomegaWithT(t)

		service := application.NewService(repository)
		command := domain.CreateAlertCommand{
			Email:     "invalid.email",
			Price:     100,
			Direction: domain.DirectionAbove,
			Coin:      "BTC",
		}
		_, err := service.CreateAlert(command)
		domainErr := domain.NewError("failed to create alert", "error.validation", "invalid email")
		g.Expect(err).Should(Equal(domainErr))
	})
}
