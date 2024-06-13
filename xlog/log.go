/**
 * Created by Goland
 * @file   log.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/6/13 16:10
 * @desc   log.go
 */

package xlog

// LogInter 日志接口
type LogInter interface {
	// WithField 添加字段
	WithField(key string, value any) LogInter
	// Debug 调试日志
	Debug(format string, args ...any)
	// Info 信息日志
	Info(format string, args ...any)
	// Warn 警告日志
	Warn(format string, args ...any)
	// Error 错误日志
	Error(format string, args ...any)
	// Fatal 致命错误日志
	Fatal(format string, args ...any)
}
