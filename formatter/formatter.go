/**
 * Created by Goland
 * @file   debug.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/6/12 11:40
 * @desc   file.go
 */

package formatter

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

// CommaNumber 从右开始每隔3个数字加上逗号。前面加符号char。
// 如果value是无效的数字字符串，例如"aa"，则返回空字符串
// Comma("12345", "$") => "$12,345", Comma(12345, "$") => "$12,345"
func CommaNumber[T constraints.Float | constraints.Integer | string](value T, symbol string) string {
	s, err := numberToString(value)
	if err != nil {
		return ""
	}

	dotIndex := strings.Index(s, ".")
	if dotIndex != -1 {
		return symbol + commaString(s[:dotIndex]) + s[dotIndex:]
	}

	return symbol + commaString(s)
}

func commaString(s string) string {
	if len(s) <= 3 {
		return s
	}
	return commaString(s[:len(s)-3]) + "," + commaString(s[len(s)-3:])
}

func numberToString(value any) (string, error) {
	switch reflect.TypeOf(value).Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return fmt.Sprintf("%v", value), nil

		// todo: need to handle 12345678.9 => 1.23456789e+07
	case reflect.Float32, reflect.Float64:
		return fmt.Sprintf("%v", value), nil

	case reflect.String:
		{
			sv := fmt.Sprintf("%v", value)
			if strings.Contains(sv, ".") {
				_, err := strconv.ParseFloat(sv, 64)
				if err != nil {
					return "", err
				}
				return sv, nil
			} else {
				_, err := strconv.ParseInt(sv, 10, 64)
				if err != nil {
					return "", nil
				}
				return sv, nil
			}
		}
	default:
		return "", nil
	}
}

// FormatSize 格式化显示大小
func FormatSize(raw float64) string {
	if raw < 1024 {
		return fmt.Sprintf("%.1fB", raw)
	}
	if raw < 1024*1024 {
		return fmt.Sprintf("%.1fK", raw/1024.0)
	}
	if raw < 1024*1024*1024 {
		return fmt.Sprintf("%.1fM", raw/1024.0/1024.0)
	}
	if raw < 1024*1024*1024*1024 {
		return fmt.Sprintf("%.1fG", raw/1024.0/1024.0/1024.0)
	}
	if raw < 1024*1024*1024*1024*1024 {
		return fmt.Sprintf("%.1fT", raw/1024.0/1024.0/1024.0/1024.0)
	}
	if raw < 1024*1024*1024*1024*1024*1024 {
		return fmt.Sprintf("%.1fP", raw/1024.0/1024.0/1024.0/1024.0/1024.0)
	}
	return "TooLarge"
}
