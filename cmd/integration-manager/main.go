package main

import (
	"gitlab.com/roneeSoft/integrator/internal/integration-manager/routes"
	"gitlab.com/roneeSoft/integrator/pkg/shared/config"
	"log"
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

	if err := router.Run(":8082"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
