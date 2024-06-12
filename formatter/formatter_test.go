/**
 * Created by Goland
 * @file   debug.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/6/12 11:40
 * @desc   file.go
 */

package formatter

import (
	"github.com/x-module/helper/internal"
	"testing"
)

func TestComma(t *testing.T) {
	assert := internal.NewAssert(t, "TestComma")

	assert.Equal("", CommaNumber("", ""))
	assert.Equal("", CommaNumber("aa", ""))
	assert.Equal("", CommaNumber("aa.a", ""))
	assert.Equal("123", CommaNumber("123", ""))
	assert.Equal("12,345", CommaNumber("12345", ""))
	assert.Equal("12,345.6789", CommaNumber("12345.6789", ""))
	assert.Equal("123,456,789,000", CommaNumber("123456789000", ""))

	assert.Equal("12,345", CommaNumber(12345, ""))
	assert.Equal("$12,345", CommaNumber(12345, "$"))
	assert.Equal("¥12,345", CommaNumber(12345, "¥"))
	assert.Equal("12,345.6789", CommaNumber(12345.6789, ""))
	assert.Equal("12,345.6789", CommaNumber(+12345.6789, ""))
	// assert.Equal("12,345,678.9", CommaNumber(12345678.9, ""))
	assert.Equal("123,456,789,000", CommaNumber(123456789000, ""))
}
