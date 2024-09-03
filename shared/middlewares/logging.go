package utils

import (
    "log"
    "os"
)

type Logger struct {
    logger *log.Logger
}

func NewLogger(prefix string) *Logger {
    return &Logger{
        logger: log.New(os.Stdout, prefix, log.LstdFlags|log.Lshortfile),
    }
}

func (l *Logger) Info(message string) {
    l.logger.Println("[INFO] " + message)
}

func (l *Logger) Warning(message string) {
    l.logger.Println("[WARNING] " + message)
}

func (l *Logger) Error(message string) {
    l.logger.Println("[ERROR] " + message)
}

func (l *Logger) Fatal(message string) {
    l.logger.Fatalln("[FATAL] " + message)
}
