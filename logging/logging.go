package logging

import (
	"log"
	"os"
	"time"
	
	"github.com/google/uuid"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func init() {
	currentTime := time.Now()

	file, err := os.OpenFile("logs" + currentTime.Format("2006-01-02") + ".log" , os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(file, "", 0)
	WarningLogger = log.New(file, "", 0)
	ErrorLogger = log.New(file, "", 0)

	requestUuid := uuid.New().String()

	InfoLogger.SetPrefix("[INFO]: " + requestUuid + " ")
	WarningLogger.SetPrefix("[WARNING]: " + requestUuid + " ")
	ErrorLogger.SetPrefix("[ERROR]: " + requestUuid + " ")
}

func LogInfo(message string) {
	InfoLogger.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	InfoLogger.Println(message)
}

func LogWarning(message string) {
	InfoLogger.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	WarningLogger.Println(message)
}

func LogError(message string) {
	InfoLogger.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	ErrorLogger.Println(message)
}
