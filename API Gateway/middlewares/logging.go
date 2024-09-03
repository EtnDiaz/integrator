package middlewares

import (
    "time"
    "github.com/gin-gonic/gin"
    "log"
)

func LoggingMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        startTime := time.Now()

        c.Next()

        latency := time.Since(startTime)
        log.Printf("Request: %s %s | Latency: %s | Status: %d",
            c.Request.Method,
            c.Request.URL.Path,
            latency,
            c.Writer.Status(),
        )
    }
}
