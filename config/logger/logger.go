package logger

import (
	"log"
	"os"
)

var logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)

func Log(args ...any) {
	logger.Println(args...)
}
