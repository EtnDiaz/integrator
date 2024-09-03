package routes

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/roneeSoft/integrator/internal/api-gateway/controllers"
	"gitlab.com/roneeSoft/integrator/pkg/shared/middlewares"
)

func SetupRouter(logger *middlewares.Logger) *gin.Engine {
    router := gin.Default()

    // Pass the logger instance to the LoggingMiddleware
    router.Use(middlewares.LoggingMiddleware(logger))
    router.Use(middlewares.AuthMiddleware())

	integrationController := controllers.IntegrationController{}
	healthController := controllers.HealthController{}

	router.GET("/integrations/:id", integrationController.GetIntegration)
	router.POST("/integrations", integrationController.CreateIntegration)
	router.GET("/health", healthController.HealthCheck)

	return router
}
