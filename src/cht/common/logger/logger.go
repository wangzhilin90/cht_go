package logger

import (
	"fmt"
	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
)

var Logger *log.Logger

func init() {
	baseLogPath := "/var/cht/go_backend_service.log"
	dir, _ := path.Split(baseLogPath)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		fmt.Errorf("MkdirAll failed:%v", err)
	}

	writer, err := rotatelogs.New(
		baseLogPath+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(baseLogPath),
		rotatelogs.WithMaxAge(time.Hour*24*7*30),
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if err != nil {
		fmt.Errorf("rotatelogs new failed:%v", err)
	}

	lfhook := lfshook.NewHook(lfshook.WriterMap{
		log.DebugLevel: writer,
		log.InfoLevel:  writer,
		log.WarnLevel:  writer,
		log.ErrorLevel: writer,
		log.PanicLevel: writer,
	})
	Logger = &log.Logger{
		Out:       writer,
		Formatter: &log.TextFormatter{},
		Level:     log.DebugLevel,
	}

	log.AddHook(lfhook)
}
