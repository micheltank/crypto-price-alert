package handler

import (
	"github.com/gin-gonic/gin"
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

func v1CreateAlert(c *gin.Context, service application.IService) {
	var request CreateAlertRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	command := request.toCommand()
	alert, err := service.CreateAlert(command)
	if e, ok := err.(*domain.ErrorDomain); ok && e != nil {
		logrus.WithError(err).Error("failed to execute v1CreateAlert")
		c.JSON(http.StatusBadRequest, NewApiError(*e))
		return
	}
	if err != nil {
		logrus.WithError(err).Error("failed to execute v1CreateAlert")
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusCreated, NewAlertResponse(alert))
}

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
	c.JSON(http.StatusOK, NewGetAlertsResponse(alerts))
}
