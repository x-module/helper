/**
 * Created by Goland
 * @file   error.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/6/13 19:18
 * @desc   error.go
 */

package xerror

import "fmt"

// Unwrap if err is nil then it returns a valid value
// If err is not nil, Unwrap panics with err.
// Play: https://go.dev/play/p/w84d7Mb3Afk
func Unwrap[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}
	return val
}

func PanicErr(err error, log string) {
	if err != nil {
		panic(fmt.Sprintf("error:%s,message:%s", err.Error(), log))
	}
}

func IgnoreErr(err error, log ...string) {
	if err != nil {
		panic(fmt.Sprintf("error:%s,message:%s", err.Error(), log))
	}
}
