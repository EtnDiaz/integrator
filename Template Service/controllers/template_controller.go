// controllers/template_controller.go
package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "template-service/services"
    "shared/utils"
)

type TemplateController struct {
    Service *services.TemplateService
}

func (ctrl *TemplateController) CreateTemplate(c *gin.Context) {
    var template services.Template
    if err := c.ShouldBindJSON(&template); err != nil {
        c.JSON(http.StatusBadRequest, utils.Response{Message: "Invalid request", Data: nil})
        return
    }

    if err := ctrl.Service.Create(template); err != nil {
        c.JSON(http.StatusInternalServerError, utils.Response{Message: "Failed to create template", Data: err.Error()})
        return
    }

    c.JSON(http.StatusCreated, utils.Response{Message: "Template created successfully", Data: template})
}
