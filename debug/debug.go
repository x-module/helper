/**
 * Created by Goland
 * @file   debug.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/6/12 11:40
 * @desc   debug.go
 */

package debug

import (
	"fmt"
	"github.com/x-module/helper/json"
)

// Display 调试输出
func Display(params any) {
	fmt.Println("----------------------- debug start ------------------------")
	json.MarshalToString(params)
	fmt.Println("----------------------- debug end --------------------------")
}
