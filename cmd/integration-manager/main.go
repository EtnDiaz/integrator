package main

import (
	"gitlab.com/roneeSoft/integrator/internal/integration-manager/routes"
	"gitlab.com/roneeSoft/integrator/pkg/shared/config"
	"log"
)

func main() {
	config.LoadConfig()

	router := routes.SetupRouter()

	if err := router.Run(":8082"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
