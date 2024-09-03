package main

import (
    "api-gateway/config"
    "api-gateway/routes"
    "log"
)

func main() {
    // Загрузка конфигурации
    config.LoadConfig()

    // Настройка роутов
    router := routes.SetupRouter()

    // Запуск сервера
    if err := router.Run(":8080"); err != nil {
        log.Fatalf("Failed to run server: %v", err)
    }
}
