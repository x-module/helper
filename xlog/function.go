/**
 * Created by goland.
 * @file   log.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2023/2/1 19:10
 * @desc   log.go
 */

package xlog

import (
	"github.com/sirupsen/logrus"
)

// WithField 添加字段
func WithField(key string, value any) *logrus.Entry {
	return Logger.WithField(key, value)
}

// Debug 调试日志
func Debug(args ...any) {
	Logger.Debug(args...)
}

// Info 信息日志
func Info(args ...any) {
	Logger.Info(args...)
}

// Warn 警告日志
func Warn(args ...any) {
	Logger.Warn(args...)
}

// Warning 警告日志
func Warning(args ...any) {
	Logger.Warn(args...)
}

// Error 错误日志
func Error(args ...any) {
	Logger.Error(args...)
}

// Fatal 致命错误日志
func Fatal(args ...any) {
	Logger.Fatal(args...)
}

// Debugf 调试日志
func Debugf(format string, args ...interface{}) {
	Logger.Debugf(format, args)
}

// Infof 信息日志
func Infof(format string, args ...interface{}) {
	Logger.Infof(format, args)
}

// Warnf 警告日志
func Warnf(format string, args ...interface{}) {
	Logger.Warnf(format, args)
}

// Warningf 警告日志
func Warningf(format string, args ...interface{}) {
	Logger.Warnf(format, args)
}

// Errorf 错误日志
func Errorf(format string, args ...interface{}) {
	Logger.Errorf(format, args)
}

// Fatalf 致命错误日志
func Fatalf(format string, args ...interface{}) {
	Logger.Fatalf(format, args)
}
