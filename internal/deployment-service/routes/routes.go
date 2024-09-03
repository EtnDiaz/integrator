package routes

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/roneeSoft/integrator/internal/deployment-service/controllers"
	"gitlab.com/roneeSoft/integrator/pkg/shared/middlewares"
)


func SetupRouter(logger *middlewares.Logger) *gin.Engine {
    router := gin.Default()

    // Pass the logger instance to the LoggingMiddleware
    router.Use(middlewares.LoggingMiddleware(logger))
    router.Use(middlewares.AuthMiddleware())

	
	deploymentController := controllers.DeploymentController{}
	healthController := controllers.HealthController{}

	router.POST("/deploy", deploymentController.Deploy)
	router.GET("/health", healthController.HealthCheck)

	return router
}
