package log

import (
	"testing"

	"github.com/JiSuanSiWeiShiXun/log"
	"github.com/sirupsen/logrus"
)

func TestLog(t *testing.T) {
	log.Logger.SetLevel(logrus.TraceLevel)
	log.Trace("trace")
	log.Debug("debug")
	log.Info("info")
	log.Logger.Infof("info by Logger")
	log.Warn("warn")
	log.Error("error")
	log.Fatal("fatal")
	// log.Panic("panic")
	t.Log("i'm still alive")
}
