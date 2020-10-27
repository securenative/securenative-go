package securenative_go

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

type SdKLogger struct {
	Log      *log.Logger
	LogLevel string
}

var SecureNativeLogger *SdKLogger

func InitLogger(logLevel string) *SdKLogger {
	SecureNativeLogger = &SdKLogger{
		Log:      log.New(os.Stdout, fmt.Sprintf("%s: ", logLevel), log.Ldate|log.Ltime|log.Lshortfile),
		LogLevel: logLevel,
	}
	return SecureNativeLogger
}

func GetLogger() *SdKLogger {
	if SecureNativeLogger == nil {
		InitLogger("CRITICAL")
	}
	return SecureNativeLogger
}

func (l *SdKLogger) Info(msg string) {
	if l.Log != nil && strings.ToLower(l.LogLevel) == "info" || l.Log != nil && strings.ToLower(l.LogLevel) == "debug" {
		l.Log.Println(msg)
	}
}

func (l *SdKLogger) Debug(msg string) {
	if l.Log != nil && strings.ToLower(l.LogLevel) == "debug" {
		l.Log.Println(msg)
	}
}

func (l *SdKLogger) Warning(msg string) {
	if l.Log != nil && strings.ToLower(l.LogLevel) == "warning" || l.Log != nil && strings.ToLower(l.LogLevel) == "debug" {
		l.Log.Println(msg)
	}
}

func (l *SdKLogger) Error(msg string) {
	if l.Log != nil && strings.ToLower(l.LogLevel) == "error" || l.Log != nil && strings.ToLower(l.LogLevel) == "debug" {
		l.Log.Println(msg)
	}
}

func (l *SdKLogger) Critical(msg string) {
	if l.Log != nil && strings.ToLower(l.LogLevel) == "critical" || l.Log != nil && strings.ToLower(l.LogLevel) == "debug" {
		l.Log.Println(msg)
	}
}