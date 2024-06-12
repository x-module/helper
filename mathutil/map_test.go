/**
 * Created by Goland
 * @file   map.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/6/11 19:56
 * @desc   map.go
 */

package mathutil

import (
	"github.com/x-module/helper/internal"
	"testing"
)

func TestMerge(t *testing.T) {
	assert := internal.NewAssert(t, "TestMerge")
	numbers := []int{
		1, 3, 6, 8, 9, 12, 34,
	}
	result := Average(numbers...)
	assert.Equal(10, result)
}
