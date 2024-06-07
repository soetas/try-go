package utils

import (
	"log"
	"os"
)

var logger *log.Logger

func Warn(message string) {
	if file, err := os.OpenFile("./log/warn.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666); err == nil {
		if logger == nil {
			logger = log.New(file, "[WARN]: ", log.Ldate|log.Lshortfile)
		}
		logger.Println(message)
	}
}
