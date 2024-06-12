/**
 * Created by Goland
 * @file   debug.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/6/12 11:40
 * @desc   debug.go
 */

package debug

import "github.com/x-module/helper/json"

func Display(params any) {
	json.MarshalToString(params)
}
