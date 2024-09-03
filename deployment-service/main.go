package main

import (
    "deployment-service/routes"
    "log"
    "shared/config"
)

func main() {
    config.LoadConfig()

    router := routes.SetupRouter()

    if err := router.Run(":8081"); err != nil {
        log.Fatalf("Failed to run server: %v", err)
    }
}
