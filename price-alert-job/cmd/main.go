package main

import (
	"github.com/micheltank/crypto-price-alert/price-alert-job/cmd/app"
	infraConfig "github.com/micheltank/crypto-price-alert/price-alert-job/internal/infra/config"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)

	appConfig := infraConfig.Env
	err := run(appConfig)
	if err != nil {
		logrus.WithError(err).
			Fatal("failed running application")
		return
	}
}

func run(appConfig infraConfig.Environment) error {
	logrus.Info("Starting application")
	app, err := app.NewApp(appConfig)
	if err != nil {
		return errors.Wrap(err, "failed to initialize app")
	}
	appErr := app.Run()
	logrus.Info("Running application")

	quit := notifyShutdown()
	defer app.Shutdown()

	select {
	case err := <-appErr:
		return errors.Wrap(err, "failed while running app")
	case <-quit:
		logrus.Info("Gracefully shutdown")
		return nil
	}
}

func notifyShutdown() chan os.Signal {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	return quit
}
