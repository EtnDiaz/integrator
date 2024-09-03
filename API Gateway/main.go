package main

import (
    "shared/config"
    "api-gateway/routes"
    "log"
)

func main() {
    config.LoadConfig()

    router := routes.SetupRouter()

    if err := router.Run(":8080"); err != nil {
        log.Fatalf("Failed to run server: %v", err)
    }
}
