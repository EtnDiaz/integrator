package routes

import (
    "github.com/gin-gonic/gin"
    "gitlab.com/roneeSoft/integrator/deployment-service/controllers"
    "gitlab.com/roneeSoft/integrator/shared/middlewares"
)

func SetupRouter() *gin.Engine {
    router := gin.Default()

    router.Use(middlewares.LoggingMiddleware())
    router.Use(middlewares.AuthMiddleware())

    deploymentController := controllers.DeploymentController{}
    healthController := controllers.HealthController{}

    router.POST("/deploy", deploymentController.Deploy)
    router.GET("/health", healthController.HealthCheck)

    return router
}
