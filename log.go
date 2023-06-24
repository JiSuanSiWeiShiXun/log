package log

/*
using logrus and lumberjack create a log package
*/

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	Logger           *logrus.Logger // 可以自己获取这个Logger更新里面的设置
	LumberjackLogger *lumberjack.Logger
	once             sync.Once
	Role             = "original"
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
			CallerPrettyfier: func(fr *runtime.Frame) (function string, file string) { // 简化文件名
				function = fr.Function
				file = fmt.Sprintf("%s:%d", filepath.Base(fr.File), fr.Line)
				return
			},
		})

		// log rotate
		LumberjackLogger = &lumberjack.Logger{
			Filename:   fmt.Sprintf("logs/%v.log", Role),
			MaxSize:    100,   //512M一个文件
			MaxBackups: 5,     //最大备份个数
			MaxAge:     7,     //最大保留天数
			Compress:   false, //归档压缩
			LocalTime:  true,  //rotate的文件后缀名使用的时区，默认为utc
		}
		mo := io.MultiWriter(LumberjackLogger, os.Stdout)
		Logger.SetOutput(mo)
	})
}

func fileInfo(skip int) string {
	// 找到封装的日志入口的上一级调用栈
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		file = "<???>"
		line = 0
	} else {
		// slash := strings.LastIndex(file, "/")
		// if slash >= 0 {
		// 	file = file[slash+1:]
		// }
		file = filepath.Base(file)
	}
	return fmt.Sprintf("%s:%d", file, line)
}
