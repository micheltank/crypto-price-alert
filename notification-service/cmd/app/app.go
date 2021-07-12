package app

import (
	"github.com/micheltank/crypto-price-alert/notification-service/internal/application"
	infraConfig "github.com/micheltank/crypto-price-alert/notification-service/internal/infra/config"
	infra "github.com/micheltank/crypto-price-alert/notification-service/internal/infra/gomail-email-service"
	"github.com/pkg/errors"
)

type App struct {
	emailListener Listener
	stop          bool
}

func NewApp(config infraConfig.Environment) (*App, error) {
	// di
	emailService := infra.NewGomailEmailService(config)
	emailNotificationService := application.NewEmailNotificationService(emailService)
	emailListener, err := NewEmailListener(emailNotificationService)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create email listener")
	}
	return &App{
		emailListener: emailListener,
		stop:          false,
	}, nil
}

func (app *App) Run() <-chan error {
	out := make(chan error)
	go func() {
		err := app.emailListener.Listen()
		if err != nil {
			out <- errors.Wrap(err, "failed to execute email listener")
		}
	}()
	return out
}

func (app *App) Shutdown() {
	app.emailListener.Shutdown()
}
