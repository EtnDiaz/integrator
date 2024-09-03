// controllers/template_controller.go
package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "gitlab.com/roneeSoft/integrator/internal/services"
    "gitlab.com/roneeSoft/integrator/pkg/shared/utils"
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
func (ctrl *TemplateController) GetTemplate(c *gin.Context) {
    id := c.Param("id")

    template, err := ctrl.Service.GetByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, utils.Response{Message: "Template not found", Data: err.Error()})
        return
    }

    c.JSON(http.StatusOK, utils.Response{Message: "Template retrieved successfully", Data: template})
}
func (ctrl *TemplateController) UpdateTemplate(c *gin.Context) {
    id := c.Param("id")

    var template services.Template
    if err := c.ShouldBindJSON(&template); err != nil {
        c.JSON(http.StatusBadRequest, utils.Response{Message: "Invalid request", Data: nil})
        return
    }

    // Pass both the ID and the template to the Update method
    if err := ctrl.Service.Update(id, template); err != nil {
        c.JSON(http.StatusInternalServerError, utils.Response{Message: "Failed to update template", Data: err.Error()})
        return
    }

    c.JSON(http.StatusOK, utils.Response{Message: "Template updated successfully", Data: template})
}


func (ctrl *TemplateController) DeleteTemplate(c *gin.Context) {
    id := c.Param("id")

    if err := ctrl.Service.Delete(id); err != nil {
        c.JSON(http.StatusInternalServerError, utils.Response{Message: "Failed to delete template", Data: err.Error()})
        return
    }

    c.JSON(http.StatusOK, utils.Response{Message: "Template deleted successfully", Data: nil})
}
