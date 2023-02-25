package log

/*
using logrus and lumberjack create a log package
*/

import (
	"fmt"
	"io"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	Logger *logrus.Logger
	once   sync.Once
	Role   = "original"
)

func init() {
	once.Do(func() {
		Logger = logrus.New()
		Logger.SetLevel(logrus.DebugLevel) // 设置日志级别
		Logger.SetReportCaller(true)       // 文件名 行号 函数名
		// 显示短文件名
		Logger.SetFormatter(&logrus.TextFormatter{
			ForceColors:     true,                  // 显示终端色彩
			FullTimestamp:   true,                  // 设置ForceColors后，必须设置这个，不然不显示时间
			TimestampFormat: "2006-01-02 15:04:05", // 显示时间格式
		})

		// log rotate
		l := &lumberjack.Logger{
			Filename:   fmt.Sprintf("log/%v.log", Role),
			MaxSize:    100,   //512M一个文件
			MaxBackups: 5,     //最大备份个数
			MaxAge:     7,     //最大保留天数
			Compress:   false, //归档压缩
			LocalTime:  true,  //rotate的文件后缀名使用的时区，默认为utc
		}
		mo := io.MultiWriter(l, os.Stdout)
		Logger.SetOutput(mo)

		// [feature] 如何让机器按小时/天rotate
		go func() {
			for {
				select {
				case <-time.After(time.Minute):
					l.Rotate()
				}
			}
		}()

		// [feature] 如何按照启动时间rotate
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGHUP)
		go func() {
			for {
				<-c
				l.Rotate()
			}
		}()
	})
}
