/**
 * Created by Goland
 * @file   catch.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/6/13 08:19
 * @desc   catch.go
 */

package catch

import (
	"github.com/x-module/helper/xlog"
	"runtime/debug"
)

const (
	SystemErr = "system error"
)

// Recover 捕获异常
func Recover(logger xlog.LogInter) {
	if err := recover(); err != nil {
		logger.WithField("err", err).WithField("trace", string(debug.Stack())).Error(SystemErr)
	}
}
