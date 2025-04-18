package logger

import (
	"log"
	"os"
)

var logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)

func Log(msg string, args ...any) {
	logger.Println(msg, args)
}
