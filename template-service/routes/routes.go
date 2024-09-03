package routes

import (
    "github.com/gin-gonic/gin"
    "template-service/controllers"
    "template-service/middlewares"
    "template-service/services"
)

func SetupRouter() *gin.Engine {
    router := gin.Default()

    router.Use(middlewares.LoggingMiddleware())
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
