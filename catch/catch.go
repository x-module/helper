/**
 * Created by Goland
 * @file   catch.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/6/13 08:19
 * @desc   catch.go
 */

package catch

import "runtime/debug"

// Recover 捕获异常
func Recover(logger runtime.Logger) {
	if err := recover(); err != nil {
		logger.WithField("err", err).WithField("trace", string(debug.Stack())).Error(global.SystemErr.String())
	}
}
