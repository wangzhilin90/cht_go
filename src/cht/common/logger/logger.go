package logger

import (
	cf "cht/common/config"
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
	baseLogPath := func() string {
		if cf.BConf.LogPath != "" {
			return cf.BConf.LogPath
		} else {
			Logger.Fatalf("log file path is null")
			return ""
		}
	}()
	fmt.Println("log file path is:", baseLogPath)
	dir, _ := path.Split(baseLogPath)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		fmt.Errorf("MkdirAll failed:%v", err)
	}

	writer, err := rotatelogs.New(
		baseLogPath+"%Y%m%d",
		rotatelogs.WithLinkName(baseLogPath),      //软连接
		rotatelogs.WithMaxAge(time.Hour*24*30*7),  //文件最大保存时间
		rotatelogs.WithRotationTime(time.Hour*24), //文件切割时间
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
