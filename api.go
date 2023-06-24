package log

/*
提供了package级别的日志记录函数，但是可以直接使用Logger记录
*/

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
)

func Trace(format string, args ...interface{}) {
	entry := Logger.WithFields(logrus.Fields{})
	entry.Data["file"] = fileInfo(2)
	entry.Logf(logrus.TraceLevel, format, args...)
}

func Debug(format string, args ...interface{}) {
	entry := Logger.WithFields(logrus.Fields{})
	entry.Data["file"] = fileInfo(2)
	entry.Logf(logrus.DebugLevel, format, args...)
}

func Info(format string, args ...interface{}) {
	entry := Logger.WithFields(logrus.Fields{})
	entry.Data["file"] = fileInfo(2)
	entry.Logf(logrus.InfoLevel, format, args...)
}

func Warn(format string, args ...interface{}) {
	entry := Logger.WithFields(logrus.Fields{})
	entry.Data["file"] = fileInfo(2)
	entry.Logf(logrus.WarnLevel, format, args...)
}

func Error(format string, args ...interface{}) {
	entry := Logger.WithFields(logrus.Fields{})
	entry.Data["file"] = fileInfo(2)
	entry.Logf(logrus.ErrorLevel, format, args...)
}

func Fatal(format string, args ...interface{}) {
	entry := Logger.WithFields(logrus.Fields{})
	entry.Data["file"] = fileInfo(2)
	entry.Fatalf(format, args...)
}

func Panic(format string, args ...interface{}) {
	entry := Logger.WithFields(logrus.Fields{})
	entry.Data["file"] = fileInfo(2)
	entry.Logf(logrus.PanicLevel, format, args...)
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
