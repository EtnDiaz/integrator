package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "gitlab.com/roneeSoft/integrator/shared/utils"
)

type IntegrationController struct {}

func (ctrl IntegrationController) GetIntegration(c *gin.Context) {
    id := c.Param("id")
    c.JSON(http.StatusOK, utils.Response{Message: "Integration fetched successfully", Data: id})
}

func (ctrl IntegrationController) CreateIntegration(c *gin.Context) {
    c.JSON(http.StatusCreated, utils.Response{Message: "Integration created successfully"})
}
