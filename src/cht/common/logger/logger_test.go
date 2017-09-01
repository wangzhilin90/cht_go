package logger

import (
	"testing"
)

func Testlog(t *testing.T) {
	Logger.Debug("debug test")
	Logger.Infof("info test")
	Logger.Warnln("warn test")
	Logger.Errorln("error test")
	Logger.Fatal("fatal test")
}
