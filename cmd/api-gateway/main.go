package main

import (
	"gitlab.com/roneeSoft/integrator/internal/api-gateway/routes"
	"gitlab.com/roneeSoft/integrator/pkg/shared/config"
	"log"
)

func main() {
	config.LoadConfig()

	router := routes.SetupRouter()

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
