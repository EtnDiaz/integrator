package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"
	"time"
)

const (
	DebugLevel = iota
	InfoLevel
	WarningLevel
	ErrorLevel
	FatalLevel
)

type Logger struct {
	logger      *log.Logger
	level       int
	jsonOutput  bool
	colorOutput bool
}

func NewLogger(prefix string, level int, jsonOutput, colorOutput bool) *Logger {
	return &Logger{
		logger:      log.New(os.Stdout, prefix, log.LstdFlags|log.Lshortfile),
		level:       level,
		jsonOutput:  jsonOutput,
		colorOutput: colorOutput,
	}
}

func (l *Logger) logMessage(level int, levelStr, message string) {
	if level < l.level {
		return
	}

	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "unknown"
		line = 0
	}

	logEntry := fmt.Sprintf("%s:%d %s", file, line, message)

	if l.jsonOutput {
		logData := map[string]interface{}{
			"timestamp": time.Now().Format(time.RFC3339),
			"level":     levelStr,
			"message":   message,
			"file":      file,
			"line":      line,
		}
		jsonLog, _ := json.Marshal(logData)
		l.logger.Println(string(jsonLog))
	} else if l.colorOutput {
		logEntry = l.colorize(level, fmt.Sprintf("[%s] %s", levelStr, logEntry))
		l.logger.Println(logEntry)
	} else {
		l.logger.Printf("[%s] %s", levelStr, logEntry)
	}
}

func (l *Logger) Info(message string) {
	l.logMessage(InfoLevel, "INFO", message)
}

func (l *Logger) Debug(message string) {
	l.logMessage(DebugLevel, "DEBUG", message)
}

func (l *Logger) Warning(message string) {
	l.logMessage(WarningLevel, "WARNING", message)
}

func (l *Logger) Error(message string) {
	l.logMessage(ErrorLevel, "ERROR", message)
}

func (l *Logger) Fatal(message string) {
	l.logMessage(FatalLevel, "FATAL", message)
	os.Exit(1)
}

func (l *Logger) colorize(level int, message string) string {
	switch level {
	case DebugLevel:
		return "\033[34m" + message + "\033[0m" // Blue
	case InfoLevel:
		return "\033[32m" + message + "\033[0m" // Green
	case WarningLevel:
		return "\033[33m" + message + "\033[0m" // Yellow
	case ErrorLevel:
		return "\033[31m" + message + "\033[0m" // Red
	case FatalLevel:
		return "\033[35m" + message + "\033[0m" // Magenta
	default:
		return message
	}
}
