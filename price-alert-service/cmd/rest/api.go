package rest

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micheltank/crypto-price-alert/price-alert-service/cmd/rest/handler"
	"github.com/micheltank/crypto-price-alert/price-alert-service/internal/application"
	"github.com/micheltank/crypto-price-alert/price-alert-service/internal/infra/config"
	"github.com/micheltank/crypto-price-alert/price-alert-service/internal/infra/repository"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"net/http"
	"sync"
	"time"
)

type Api struct {
	httpServer *http.Server
	Db         *sql.DB
}

// NewServer godoc
// @title Price Alert Service API
// @version 1.0
func NewServer(config config.Environment) (*Api, error) {
	dataSourceName := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", config.DbConfig.User, config.DbConfig.Password, config.DbConfig.Host, config.DbConfig.Port, config.DbConfig.Name)
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open postgres connection")
	}

	router := gin.Default()
	base := router.Group("/api")

	v1 := base.Group("/v1")

	// di
	alertRepository := repository.NewAlertPostgreSql(db)
	alertService := application.NewService(alertRepository)

	// handlers
	handler.MakeHealthCheckHandler(base)
	handler.MakeAlertsHandler(v1, alertService)

	// documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	httpServer := &http.Server{Addr: fmt.Sprintf(":%d", config.Port), Handler: router}

	return &Api{
		httpServer: httpServer,
		Db:         db,
	}, nil
}

func (api *Api) Run() <-chan error {
	out := make(chan error)
	go func() {
		if err := api.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			out <- errors.Wrap(err, "failed to listen and serve api")
		}
	}()
	return out
}

func (api *Api) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := api.httpServer.Shutdown(ctx)
		if err != nil {
			logrus.
				WithError(err).
				Error("Server forced to shutdown")
		}
	}()
	err := api.Db.Close()
	if err != nil {
		logrus.WithError(err).Error("Failed to close db")
	}
}
