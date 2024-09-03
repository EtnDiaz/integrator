package main

import (
	"log"

	"gitlab.com/roneeSoft/integrator/internal/template-service/routes"
	"gitlab.com/roneeSoft/integrator/pkg/shared/config"
	"gitlab.com/roneeSoft/integrator/pkg/shared/middlewares"
)

func main() {
	config.LoadConfig()

	// Create a logger instance
	logger := &middlewares.Logger{
		// Set any necessary fields for your logger, e.g., jsonOutput: true
	}

	// Pass the logger to SetupRouter
	router := routes.SetupRouter(logger)

	if err := router.Run(":8083"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
