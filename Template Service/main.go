package main

import (
    "template-service/routes"
    "shared/config"
    "log"
)

func main() {
    config.LoadConfig()

    router := routes.SetupRouter()

    if err := router.Run(":8083"); err != nil {
        log.Fatalf("Failed to run server: %v", err)
    }
}
