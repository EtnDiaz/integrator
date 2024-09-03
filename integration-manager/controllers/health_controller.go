package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type HealthController struct {}

func (ctrl HealthController) HealthCheck(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"status": "healthy"})
}
