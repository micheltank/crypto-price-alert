package main

import (
	"github.com/micheltank/crypto-price-alert/price-alert-service/cmd/grpc"
	"github.com/micheltank/crypto-price-alert/price-alert-service/cmd/rest"
	"github.com/micheltank/crypto-price-alert/price-alert-service/internal/infra/config"
	"github.com/micheltank/crypto-price-alert/price-alert-service/internal/infra/db"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)

	appConfig := config.Env
	err := db.Migrate(appConfig.DbConfig)
	if err != nil {
		logrus.WithError(err).
			Fatal("failed to migrate")
		return
	}
	err = run(appConfig)
	if err != nil {
		logrus.WithError(err).
			Fatal("failed running application")
		return
	}
}

func run(appConfig config.Environment) error {
	logrus.Info("Starting application")

	// HTTP Server
	restApiServer, err := rest.NewServer(appConfig)
	if err != nil {
		return errors.Wrap(err, "failed to initialize restApiServer")
	}
	restApiErr := restApiServer.Run()
	logrus.Infof("Running http server on port %d", appConfig.Port)
	defer restApiServer.Shutdown()

	// GRPC Server
	grpcApiServer, err := grpc.NewServer(appConfig)
	if err != nil {
		return errors.Wrap(err, "failed to initialize grpcApiServer")
	}
	grpcApiErr := grpcApiServer.Run()
	logrus.Infof("Running grpc server on port %d", appConfig.GrpcPort)
	defer grpcApiServer.Shutdown()

	// Shutdown
	quit := notifyShutdown()
	select {
	case err := <-restApiErr:
		return errors.Wrap(err, "failed while running restApiServer")
	case err := <-grpcApiErr:
		return errors.Wrap(err, "failed while running restApiServer")
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