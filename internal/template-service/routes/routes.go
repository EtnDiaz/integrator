package routes

import (
    "github.com/gin-gonic/gin"
    "gitlab.com/roneeSoft/integrator/internal/template-service/controllers"
    "gitlab.com/roneeSoft/integrator/internal/services"
    "gitlab.com/roneeSoft/integrator/pkg/shared/middlewares"
)

func SetupRouter(logger *middlewares.Logger) *gin.Engine {
    router := gin.Default()

    // Pass the logger instance to the LoggingMiddleware
    router.Use(middlewares.LoggingMiddleware(logger))
    router.Use(middlewares.AuthMiddleware())

    templateService := services.NewTemplateService()
    templateController := controllers.TemplateController{Service: templateService}
    healthController := controllers.HealthController{}

    router.POST("/templates", templateController.CreateTemplate)
    router.GET("/templates/:id", templateController.GetTemplate)
    router.PUT("/templates/:id", templateController.UpdateTemplate)
    router.DELETE("/templates/:id", templateController.DeleteTemplate)
    router.GET("/health", healthController.HealthCheck)

    return router
}
