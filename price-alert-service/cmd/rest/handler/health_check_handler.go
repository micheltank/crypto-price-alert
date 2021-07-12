package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func MakeHealthCheckHandler(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/health-check", func(c *gin.Context) {
		c.Writer.WriteHeader(http.StatusOK)
	})
}