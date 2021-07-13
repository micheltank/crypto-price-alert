package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/micheltank/crypto-price-alert/price-alert-service/cmd/rest/presenter"
	"github.com/micheltank/crypto-price-alert/price-alert-service/internal/application"
	"github.com/micheltank/crypto-price-alert/price-alert-service/internal/domain"
	"github.com/sirupsen/logrus"
	"net/http"
)

func MakeAlertsHandler(routerGroup *gin.RouterGroup, service application.IService) {
	routerGroup.GET("/alerts", func(c *gin.Context) {
		v1GetAlerts(c, service)
	})
	routerGroup.POST("/alerts", func(c *gin.Context) {
		v1CreateAlert(c, service)
	})
}

// v1CreateAlert godoc
// @Summary Create an alert
// @Description Create an alert
// @ID create-alert
// @Tags Alerts
// @Produce  json
// @Success 201 {object} presenter.AlertResponse
// @Error 500 {object} presenter.ApiError
// @Router /alerts [post]
func v1CreateAlert(c *gin.Context, service application.IService) {
	var request presenter.CreateAlertRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	command := request.ToCommand()
	alert, err := service.CreateAlert(command)
	if e, ok := err.(*domain.ErrorDomain); ok && e != nil {
		logrus.WithError(err).Error("failed to execute v1CreateAlert")
		c.JSON(http.StatusBadRequest, presenter.NewApiError(*e))
		return
	}
	if err != nil {
		logrus.WithError(err).Error("failed to execute v1CreateAlert")
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusCreated, presenter.NewAlertResponse(alert))
}

// v1GetAlerts godoc
// @Summary Create an alert by email
// @Description get an alert by email
// @ID create-alert
// @Tags Alerts
// @Param email query string true "Email"
// @Produce  json
// @Success 200 {object} presenter.AlertResponse
// @Error 500 {object} presenter.ApiError
// @Router /alerts [get]
func v1GetAlerts(c *gin.Context, service application.IService) {
	email, found := c.GetQuery("email")
	if !found {
		c.Status(http.StatusBadRequest)
	}
	alerts, err := service.GetAlerts(email)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, presenter.NewGetAlertsResponse(alerts))
}
