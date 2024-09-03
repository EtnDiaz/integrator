package routes

import (
    "github.com/gin-gonic/gin"
    "deployment-service/controllers"
    "shared/middlewares"
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
