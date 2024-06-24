/**
 * Created by Goland
 * @file   log.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/6/13 16:10
 * @desc   log.go
 */

package xlog

import (
	"github.com/sirupsen/logrus"
)

// LogInter 日志接口
type LogInter interface {
	// Debug 调试日志
	Debug(args ...any)
	// Info 信息日志
	Info(args ...any)
	// Warn 警告日志
	Warn(args ...any)
	// Error 错误日志
	Error(args ...any)
	// Fatal 致命错误日志
	Fatal(args ...any)
	// WithField 添加字段
	WithField(key string, value any) *logrus.Entry
}

//
// // Interface logger interface
// type Interface interface {
// 	LogMode(LogLevel) Interface
// 	Info(context.Context, string, ...interface{})
// 	Warn(context.Context, string, ...interface{})
// 	Error(context.Context, string, ...interface{})
// 	Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error)
// }
