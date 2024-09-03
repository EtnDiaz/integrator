package routes

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/roneeSoft/integrator/internal/integration-manager/controllers"
	"gitlab.com/roneeSoft/integrator/pkg/shared/middlewares"
    "gitlab.com/roneeSoft/integrator/internal/services"
)



func SetupRouter(logger *middlewares.Logger) *gin.Engine {
    router := gin.Default()

    // Pass the logger instance to the LoggingMiddleware
    router.Use(middlewares.LoggingMiddleware(logger))
    router.Use(middlewares.AuthMiddleware())

	integrationService := services.NewIntegrationService()
	integrationController := controllers.IntegrationController{Service: integrationService}
	healthController := controllers.HealthController{}

	router.POST("/integrations", integrationController.CreateIntegration)
	router.GET("/integrations/:id", integrationController.GetIntegration)
	router.GET("/health", healthController.HealthCheck)

	return router
}
