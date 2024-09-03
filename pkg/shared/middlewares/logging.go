package middlewares

import (
    "fmt"
    "time"
    "encoding/json"
    "github.com/gin-gonic/gin"
)

type Logger struct {
    jsonOutput bool
}

func (l *Logger) Error(message string) {
    fmt.Println("ERROR: " + message)
}

func (l *Logger) Warning(message string) {
    fmt.Println("WARNING: " + message)
}

func (l *Logger) Info(message string) {
    fmt.Println("INFO: " + message)
}

func LoggingMiddleware(logger *Logger) gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()

        // Process request
        c.Next()

        // Calculate the time taken to process the request
        duration := time.Since(start)

        // Get the status code, method, and path of the request
        statusCode := c.Writer.Status()
        method := c.Request.Method
        path := c.Request.URL.Path

        // Format the log message
        logMessage := map[string]interface{}{
            "method":     method,
            "path":       path,
            "statusCode": statusCode,
            "duration":   duration.Seconds(),
        }

        // Convert log message to JSON if jsonOutput is enabled
        logStr := ""
        if logger.jsonOutput {
            logBytes, _ := json.Marshal(logMessage)
            logStr = string(logBytes)
        } else {
            logStr = fmt.Sprintf("%s %s %d %0.3f", method, path, statusCode, duration.Seconds())
        }

        // Log based on the status code
        switch {
        case statusCode >= 500:
            logger.Error(logStr)
        case statusCode >= 400:
            logger.Warning(logStr)
        case statusCode >= 300:
            logger.Info(logStr)
        default:
            logger.Info(logStr)
        }
    }
}
