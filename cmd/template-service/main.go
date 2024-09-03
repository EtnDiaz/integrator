package main

import (
	"gitlab.com/roneeSoft/integrator/internal/template-service/routes"
	"gitlab.com/roneeSoft/integrator/pkg/shared/config"
	"log"
)

func main() {
	config.LoadConfig()

	router := routes.SetupRouter()

	if err := router.Run(":8083"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
