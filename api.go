package log

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
)

func Trace(format string, args ...interface{}) {
	Logger.Logf(logrus.TraceLevel, format, args...)
}

func Dubug(format string, args ...interface{}) {
	Logger.Logf(logrus.DebugLevel, format, args...)
}

func Info(format string, args ...interface{}) {
	Logger.Logf(logrus.InfoLevel, format, args...)
}

func Warn(format string, args ...interface{}) {
	Logger.Logf(logrus.WarnLevel, format, args...)
}

func Error(format string, args ...interface{}) {
	Logger.Logf(logrus.ErrorLevel, format, args...)
}

func Fatal(format string, args ...interface{}) {
	Logger.Logf(logrus.FatalLevel, format, args...)
}

func Panic(format string, args ...interface{}) {
	Logger.Logf(logrus.PanicLevel, format, args...)
}

// SetLogRotateByTime 设置按时间更新文件
// 并非整点rotate，如果要整点rotate，依葫芦画瓢自己实现一下
func SetLogRotateByTime(t string) {
	var d time.Duration
	if t == "hour" {
		d = time.Hour
	} else if t == "day" {
		d = time.Hour * 24
	}

	// [feature] 如何让机器按小时/天rotate
	go func() {
		for {
			select {
			case <-time.After(d):
				LumberjackLogger.Rotate()
			}
		}
	}()
}

// RotateEveryLaunch 每次启动会更新一次
func RotateEveryLaunch() {
	// [feature] 如何按照启动时间rotate
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP)
	go func() {
		for {
			<-c
			LumberjackLogger.Rotate()
		}
	}()
}
