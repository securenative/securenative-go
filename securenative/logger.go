package securenative

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type LoggerInterface interface {
	Info(msg string)
	Debug(msg string)
	Warning(msg string)
	Error(msg string)
}

type SecureNativeLogger struct {
	Log      *log.Logger
	LogLevel string
}

func NewSecureNativeLogger(logLevel string) *SecureNativeLogger {
	logger := log.New(os.Stdout, fmt.Sprintf("%s: ", logLevel), log.Ldate|log.Ltime|log.Lshortfile)
	return &SecureNativeLogger{Log: logger, LogLevel: logLevel}
}

func (l *SecureNativeLogger) Info(msg string) {
	if strings.ToLower(l.LogLevel) == "info" {
		l.Log.Println(msg)
	}
}

func (l *SecureNativeLogger) Debug(msg string) {
	if strings.ToLower(l.LogLevel) == "debug" {
		l.Log.Println(msg)
	}
}

func (l *SecureNativeLogger) Warning(msg string) {
	if strings.ToLower(l.LogLevel) == "warning" {
		l.Log.Println(msg)
	}
}

func (l *SecureNativeLogger) Error(msg string) {
	if strings.ToLower(l.LogLevel) == "error" {
		l.Log.Println(msg)
	}
}
