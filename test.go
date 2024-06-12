/**
 * Created by Goland
 * @file   test.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/6/12 11:27
 * @desc   test.go
 */

package main

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
)

func main() {
	type Student struct {
		Name   string
		Age    int
		Height float32
	}

	// 四空格缩进的格式化输出
	c := jsoniter.Config{IndentionStep: 4}.Froze()
	if s, err := c.MarshalToString(map[string]any{
		"age":  23,
		"name": "44",
		"data": map[string]any{
			"test": "test",
		},
	}); err == nil {
		fmt.Println(s)
		// Output:
		// {
		//     "Name": "Allen",
		//     "Age": 18,
		//     "Height": 180.43
		// }
	}
}
