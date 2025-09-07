package platform

import (
	"log"
	"os"
	"time"
)

var Logger *log.Logger

func InitLogger() {
	Logger = log.New(os.Stdout, "", log.LstdFlags)
}

func LogInfo(message string) {
	Logger.Printf("INFO: %s", message)
}

func LogError(message string) {
	Logger.Printf("ERROR: %s", message)
}

func GetCurrentTime() string {
	return time.Now().Format(time.RFC3339)
}