/**
 * Created by goland.
 * @file   log.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2023/2/1 19:10
 * @desc   log.go
 */

package xlog

// Debug 调试日志
func Debug(format string, args ...any) {
	Logger.Debug(format, args...)
}

// Info 信息日志
func Info(format string, args ...any) {
	Logger.Info(format, args...)
}

// Warn 警告日志
func Warn(format string, args ...any) {
	Logger.Warn(format, args...)
}

// Warning 警告日志
func Warning(format string, args ...any) {
	Logger.Warn(format, args...)
}

// Error 错误日志
func Error(format string, args ...any) {
	Logger.Error(format, args...)
}

// Fatal 致命错误日志
func Fatal(format string, args ...any) {
	Logger.Fatal(format, args...)
}
