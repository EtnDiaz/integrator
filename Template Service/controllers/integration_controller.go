package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "integration-manager/services"
    "integration-manager/utils"
)

type IntegrationController struct {
    Service services.IntegrationService
}

func (ctrl IntegrationController) CreateIntegration(c *gin.Context) {
    var integration services.Integration
    if err := c.ShouldBindJSON(&integration); err != nil {
        c.JSON(http.StatusBadRequest, utils.Response{Message: "Invalid request", Data: nil})
        return
    }

    err := ctrl.Service.Create(integration)
    if err != nil {
        c.JSON(http.StatusInternalServerError, utils.Response{Message: "Failed to create integration", Data: nil})
        return
    }

    c.JSON(http.StatusCreated, utils.Response{Message: "Integration created successfully", Data: integration})
}

func (ctrl IntegrationController) GetIntegration(c *gin.Context) {
    id := c.Param("id")
    integration, err := ctrl.Service.Get(id)
    if err != nil {
        c.JSON(http.StatusNotFound, utils.Response{Message: "Integration not found", Data: nil})
        return
    }

    c.JSON(http.StatusOK, utils.Response{Message: "Integration fetched successfully", Data: integration})
}

func (ctrl IntegrationController) DeleteIntegration(c *gin.Context) {
    id := c.Param("id")
    err := ctrl.Service.Delete(id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, utils.Response{Message: "Failed to delete integration", Data: nil})
        return
    }

    c.JSON(http.StatusOK, utils.Response{Message: "Integration deleted successfully", Data: nil})
}