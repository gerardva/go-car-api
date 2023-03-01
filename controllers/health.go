package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthController struct{}

func (h HealthController) Check(c *gin.Context) {
	c.String(http.StatusOK, "Check OK")
}