package routes

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/roneeSoft/integrator/internal/api-gateway/controllers"
	"gitlab.com/roneeSoft/integrator/pkg/shared/middlewares"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(middlewares.LoggingMiddleware())
	router.Use(middlewares.AuthMiddleware())

	integrationService := services.NewIntegrationService()
	integrationController := controllers.IntegrationController{Service: integrationService}
	healthController := controllers.HealthController{}

	router.POST("/integrations", integrationController.CreateIntegration)
	router.GET("/integrations/:id", integrationController.GetIntegration)
	router.DELETE("/integrations/:id", integrationController.DeleteIntegration)
	router.GET("/health", healthController.HealthCheck)

	return router
}
