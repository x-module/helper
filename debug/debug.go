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
	"runtime"
)

// DumpPrint 调试输出
func DumpPrint(params any) {
	fmt.Println("----------------------- debug start ------------------------")
	json.MarshalToString(params)
	fmt.Println("----------------------- debug end --------------------------")
}

// DumpStacks 打印堆栈信息.
func DumpStacks() {
	buf := make([]byte, 16384)
	buf = buf[:runtime.Stack(buf, true)]
	fmt.Printf("=== BEGIN stack dump ===\n%s\n=== END stack dump ===\n\n", buf)
}

// GetCallFile 获取调用方法的文件路径.
func GetCallFile() string {
	_, file, _, _ := runtime.Caller(1)
	return file
}

// GetCallLine 获取调用方法的行号.
func GetCallLine() int {
	// Skip this function, and fetch the PC and file for its parent
	_, _, line, _ := runtime.Caller(1)
	return line
}
