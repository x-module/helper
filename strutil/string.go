/**
 * Created by Goland
 * @file   string.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/6/11 17:45
 * @desc   string.go
 */

package strutil

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

// CamelCase 将字符串转换为驼峰命名
func CamelCase(s string) string {
	var builder strings.Builder
	strs := splitIntoStrings(s, false)
	fmt.Printf("%+v\n", strs)
	for i, str := range strs {
		if i == 0 {
			builder.WriteString(strings.ToLower(str))
		} else {
			builder.WriteString(Capitalize(str))
		}
	}
	return builder.String()
}

// UnderScoreCase 将字符串转换为下划线命名
func UnderScoreCase(s string) string {
	var builder strings.Builder
	strs := splitIntoStrings(s, false)
	for i, str := range strs {
		if i == 0 {
			builder.WriteString(strings.ToLower(str))
		} else {
			builder.WriteString("_")
			builder.WriteString(strings.ToLower(str))
		}
	}
	return builder.String()
}

// Capitalize 将字符串的第一个字符转换为大写，其余字符转换为小写
func Capitalize(s string) string {
	result := make([]rune, len(s))
	for i, v := range s {
		if i == 0 {
			result[i] = unicode.ToUpper(v)
		} else {
			result[i] = unicode.ToLower(v)
		}
	}
	return string(result)
}

// UpperFirst 首字母大写
func UpperFirst(s string) string {
	if len(s) == 0 {
		return ""
	}
	// unpacks the first UTF-8 encoding，  returns the rune and its width in bytes.
	r, size := utf8.DecodeRuneInString(s)
	r = unicode.ToUpper(r)
	return string(r) + s[size:]
}

// LowerFirst 首字母小写
func LowerFirst(s string) string {
	if len(s) == 0 {
		return ""
	}
	r, size := utf8.DecodeRuneInString(s)
	r = unicode.ToLower(r)
	return string(r) + s[size:]
}

// AppendEnd 在字符串末尾填充指定字符
func AppendEnd(source string, size int, padStr string) string {
	len1 := len(source)
	len2 := len(padStr)

	if len1 >= size {
		return source
	}

	fill := ""
	if len2 >= size-len1 {
		fill = padStr[0 : size-len1]
	} else {
		fill = strings.Repeat(padStr, size-len1)
	}
	return source + fill[0:size-len1]
}

// AppendStart 在字符串末头填充指定字符
func AppendStart(source string, size int, padStr string) string {
	len1 := len(source)
	len2 := len(padStr)

	if len1 >= size {
		return source
	}

	fill := ""
	if len2 >= size-len1 {
		fill = padStr[0 : size-len1]
	} else {
		fill = strings.Repeat(padStr, size-len1)
	}
	return fill[0:size-len1] + source
}

// KebabCase 将字符串转换成  aa-bb-cc 的形式
func KebabCase(s string) string {
	result := splitIntoStrings(s, false)
	return strings.Join(result, "-")
}

// UpperKebabCase 将字符串转换成  AA-BB-CC 的形式
func UpperKebabCase(s string) string {
	result := splitIntoStrings(s, true)
	return strings.Join(result, "-")
}

// SnakeCase 将字符串转换成 aa_bb_cc 的形式
func SnakeCase(s string) string {
	result := splitIntoStrings(s, false)
	return strings.Join(result, "_")
}

// UpperSnakeCase 将字符串转换成 AA_BB_CC 的形式
func UpperSnakeCase(s string) string {
	result := splitIntoStrings(s, true)
	return strings.Join(result, "_")
}

// Before 返回字符串中第一个出现的指定字符之前的所有字符
func Before(s, char string) string {
	if s == "" || char == "" {
		return s
	}
	i := strings.Index(s, char)
	return s[0:i]
}

// BeforeLast 返回字符串中最后一个出现的指定字符之前的所有字符
func BeforeLast(s, char string) string {
	if s == "" || char == "" {
		return s
	}
	i := strings.LastIndex(s, char)
	return s[0:i]
}

// After 返回字符串中第一个出现的指定字符之后的所有字符
func After(s, char string) string {
	if s == "" || char == "" {
		return s
	}
	i := strings.Index(s, char)
	return s[i+len(char):]
}

// AfterLast 返回字符串中最后一个出现的指定字符之后的所有字符
func AfterLast(s, char string) string {
	if s == "" || char == "" {
		return s
	}
	i := strings.LastIndex(s, char)
	return s[i+len(char):]
}

// IsString 判断是否为字符串
func IsString(v any) bool {
	if v == nil {
		return false
	}
	switch v.(type) {
	case string:
		return true
	default:
		return false
	}
}

// Reverse 反转字符串
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// Wrap 用给定字符串包装字符串
func Wrap(str string, wrapWith string) string {
	if str == "" || wrapWith == "" {
		return str
	}
	var sb strings.Builder
	sb.WriteString(wrapWith)
	sb.WriteString(str)
	sb.WriteString(wrapWith)
	return sb.String()
}

// Unwrap 解除给定字符串包装字符串
func Unwrap(str string, wrapToken string) string {
	if str == "" || wrapToken == "" {
		return str
	}
	firstIndex := strings.Index(str, wrapToken)
	lastIndex := strings.LastIndex(str, wrapToken)
	if firstIndex == 0 && lastIndex > 0 && lastIndex <= len(str)-1 {
		if len(wrapToken) <= lastIndex {
			str = str[len(wrapToken):lastIndex]
		}
	}
	return str
}

// Substring 截取字符串
func Substring(s string, offset int, length uint) string {
	rs := []rune(s)
	size := len(rs)
	if offset < 0 {
		offset = size + offset
		if offset < 0 {
			offset = 0
		}
	}
	if offset > size {
		return ""
	}
	if length > uint(size)-uint(offset) {
		length = uint(size - offset)
	}
	str := string(rs[offset : offset+int(length)])
	return strings.Replace(str, "\x00", "", -1)
}

// SimilarText 计算两个字符串的相似度.
func SimilarText(str1, str2 string, len1, len2 int) int {
	var sum, max int
	pos1, pos2 := 0, 0
	// Find the longest segment of the same section in two strings
	for i := 0; i < len1; i++ {
		for j := 0; j < len2; j++ {
			for l := 0; (i+l < len1) && (j+l < len2) && (str1[i+l] == str2[j+l]); l++ {
				if l+1 > max {
					max = l + 1
					pos1 = i
					pos2 = j
				}
			}
		}
	}
	if sum = max; sum > 0 {
		if pos1 > 0 && pos2 > 0 {
			sum += SimilarText(str1, str2, pos1, pos2)
		}
		if (pos1+max < len1) && (pos2+max < len2) {
			s1 := []byte(str1)
			s2 := []byte(str2)
			sum += SimilarText(string(s1[pos1+max:]), string(s2[pos2+max:]), len1-pos1-max, len2-pos2-max)
		}
	}
	return sum
}

func Replace(source string, targets []string, replace []string) string {
	if len(targets) == 0 || len(replace) == 0 {
		return source
	}
	if len(targets) == len(replace) {
		for key, target := range targets {
			source = strings.Replace(source, target, replace[key], -1)
		}
	} else {
		for _, target := range targets {
			source = strings.Replace(source, target, replace[0], -1)
		}
	}
	return source
}
