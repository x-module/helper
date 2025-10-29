/**
 * Created by goland.
 * @file   log.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2023/2/1 19:10
 * @desc   log.go
 */

package xlog

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"github.com/x-module/helper/fileutil"
	"github.com/x-module/helper/xerror"
	"os"
	"path"
	"time"
)

// Logger 日志句柄
var Logger *Xlogger

type LogMode string

const (
	// DebugMode indicates gin mode is debug.
	DebugMode LogMode = "debug"
	// ReleaseMode indicates gin mode is release.
	ReleaseMode LogMode = "release"
	// TestMode indicates gin mode is test.
	TestMode LogMode = "test"
)

type XLogInter interface {
	// GetLoggerSource 获取日志资源
	GetLoggerSource() *logrus.Logger
}

// Xlogger 系统日志
type Xlogger struct {
	logger *logrus.Logger
}

// WithField 添加字段
func (x *Xlogger) WithField(key string, value any) *logrus.Entry {
	return x.logger.WithField(key, value)
}

// Debug 调试日志
func (x *Xlogger) Debug(format string, args ...any) {
	x.logger.Debug(fmt.Sprintf(format, args...))
}

// Info 信息日志
func (x *Xlogger) Info(format string, args ...any) {
	x.logger.Info(fmt.Sprintf(format, args...))
}

// Warn 警告日志
func (x *Xlogger) Warn(format string, args ...any) {
	x.logger.Warn(fmt.Sprintf(format, args...))
}

// Warning 警告日志
func (x *Xlogger) Warning(format string, args ...any) {
	x.logger.Warn(fmt.Sprintf(format, args...))
}

// Error 错误日志
func (x *Xlogger) Error(format string, args ...any) {
	x.logger.Error(fmt.Sprintf(format, args...))
}

// Fatal 致命错误日志
func (x *Xlogger) Fatal(format string, args ...any) {
	x.logger.Fatal(fmt.Sprintf(format, args...))
}

// GetLoggerSource 获取日志资源
func (x *Xlogger) GetLoggerSource() *logrus.Logger {
	return x.logger
}

type LogConfig struct {
	LogPath  string
	LogFile  string
	LogModel LogMode
}

// InitLogger 日志模块初始化
func InitLogger(config LogConfig) *Xlogger {
	if !fileutil.IsExist(config.LogPath) {
		err := fileutil.CreateDir(config.LogPath)
		xerror.PanicErr(err, "init system error. make log data err.path:"+config.LogFile)
	}
	fileName := path.Join(config.LogPath, config.LogFile)
	if !fileutil.IsExist(fileName) {
		fileutil.CreateFile(fileName)
		if !fileutil.IsExist(fileName) {
			panic("init system error. create log file err. log file:" + fileName)
		}
	}
	// 写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	xerror.PanicErr(err, "open log file error")
	// 实例化
	logger := logrus.New()
	// 设置日志级别
	switch config.LogModel {
	case ReleaseMode:
		logger.SetLevel(logrus.WarnLevel)
	case DebugMode:
		logger.SetLevel(logrus.DebugLevel)
	case TestMode:
		logger.SetLevel(logrus.InfoLevel)
	}

	// 设置输出
	logger.Out = src
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})

	logger.SetOutput(os.Stdout)

	logWriter, err := rotatelogs.New(
		// 分割后的文件名称
		fileName+".%Y%m%d.log",
		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(fileName),
		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(30*24*time.Hour),
		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	logger.AddHook(lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	}))
	// 是否显示文件位置
	// Logger.SetReportCaller(true)
	Logger = &Xlogger{logger: logger}
	return Logger
}
