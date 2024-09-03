package routes

import (
    "github.com/gin-gonic/gin"
    "gitlab.com/roneeSoft/integrator/api-gateway/controllers"
    "gitlab.com/roneeSoft/integrator/shared/middlewares"
)
func SetupRouter() *gin.Engine {
    router := gin.Default()
//
    router.Use(middlewares.LoggingMiddleware())
    router.Use(middlewares.AuthMiddleware())

    integrationController := controllers.IntegrationController{}
    healthController := controllers.HealthController{}

    router.GET("/integrations/:id", integrationController.GetIntegration)
    router.POST("/integrations", integrationController.CreateIntegration)
    router.GET("/health", healthController.HealthCheck)

    return router
}
