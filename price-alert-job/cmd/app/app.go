package app

import (
	"github.com/micheltank/crypto-price-alert/price-alert-job/internal/application"
	infraAlert "github.com/micheltank/crypto-price-alert/price-alert-job/internal/infra/alert-service"
	infraConfig "github.com/micheltank/crypto-price-alert/price-alert-job/internal/infra/config"
	infraCryptoCompare "github.com/micheltank/crypto-price-alert/price-alert-job/internal/infra/crypto-compare"
	infra "github.com/micheltank/crypto-price-alert/price-alert-job/internal/infra/notification-broker"
	"github.com/pkg/errors"
	"time"
)

type App struct {
	service     application.IService
	periodicity time.Duration
	inShutdown  bool
}

func NewApp(config infraConfig.Environment) (*App, error) {
	// di
	priceApi := infraCryptoCompare.NewCryptoComparePriceApi(config.CryptoCompareApiHost, config.CryptoCompareApiKey)
	alertApi, err := infraAlert.NewAlertApi(config.PriceAlertServiceHost)
	notificationBroker, err := infra.NewNotificationBroker()
	if err != nil {
		return nil, errors.Wrap(err, "failed to create alert api")
	}
	alertService := application.NewService(priceApi, alertApi, notificationBroker)

	periodicity, err := time.ParseDuration(config.Periodicity)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse periodicity from config")
	}
	return &App{
		service:     alertService,
		periodicity: periodicity,
		inShutdown:  false,
	}, nil
}

func (app *App) Run() <-chan error {
	out := make(chan error)
	go func() {
		for {
			_, err := app.service.Execute()
			if err != nil {
				out <- errors.Wrap(err, "failed to execute job")
			}
			if app.inShutdown {
				break
			}
			time.Sleep(app.periodicity)
		}
	}()
	return out
}

func (app *App) Shutdown() {
	// TODO: implement wait
	app.inShutdown = true
}
