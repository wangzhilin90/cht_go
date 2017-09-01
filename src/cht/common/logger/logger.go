package logger

import (
	log "github.com/sirupsen/logrus"
	"os"
)

var Logger *log.Logger

func init() {
	file, _ := os.Create("/var/cht/test.log")
	Logger = &log.Logger{
		Out:       file,
		Formatter: &log.TextFormatter{},
		Level:     log.DebugLevel,
	}
}
