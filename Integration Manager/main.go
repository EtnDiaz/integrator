package main

import (
    "shared/config"
    "integration-manager/routes"
    "log"
)

func main() {
    config.LoadConfig()

    router := routes.SetupRouter()

    // Запуск сервера
    if err := router.Run(":8082"); err != nil {
        log.Fatalf("Failed to run server: %v", err)
    }
}
