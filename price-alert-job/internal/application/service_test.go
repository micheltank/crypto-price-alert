package application_test

import (
	"github.com/golang/mock/gomock"
	"github.com/micheltank/crypto-price-alert/price-alert-job/internal/application"
	mock_domain "github.com/micheltank/crypto-price-alert/price-alert-job/internal/application/mock"
	"github.com/micheltank/crypto-price-alert/price-alert-job/internal/domain"
	. "github.com/onsi/gomega"
	"testing"
)

func TestAlertService(t *testing.T) {
	ctrl := gomock.NewController(t)
	priceApi := mock_domain.NewMockPriceApi(ctrl)
	alertApi := mock_domain.NewMockAlertApi(ctrl)
	notificationBroker := mock_domain.NewMockNotificationBroker(ctrl)

	priceApi.EXPECT().
		GetPrice(gomock.Any()).
		DoAndReturn(func(coin string) (domain.Price, error) {
			return domain.NewPrice(5000), nil
		}).
		MaxTimes(3)

	alertApi.EXPECT().
		GetAlerts(gomock.Any(), gomock.Any()).
		DoAndReturn(func(coin string, price float64) (domain.Alerts, error) {
			var alerts domain.Alerts
			if coin == "BTC" {
				alerts = append(alerts, domain.NewAlert(1, "john@gmail.com", 4000, coin, domain.DirectionAbove))
				alerts = append(alerts, domain.NewAlert(2, "john@gmail.com", 6000, coin, domain.DirectionBelow))
			} else {
				alerts = append(alerts, domain.NewAlert(3, "maria@gmail.com", 6000, coin, domain.DirectionBelow))
			}
			return alerts, nil
		}).
		MaxTimes(3)

	notificationBroker.EXPECT().
		SendEmail(gomock.Any()).
		DoAndReturn(func(notification domain.EmailNotification) error {
			return nil
		}).
		MaxTimes(4)

	t.Run("Regular execution", func(t *testing.T) {
		g := NewGomegaWithT(t)

		service := application.NewService(priceApi, alertApi, notificationBroker)
		processedItems, err := service.Execute()
		g.Expect(err).Should(
			Not(HaveOccurred()))
		g.Expect(processedItems).Should(
			Equal(4))
	})
}
