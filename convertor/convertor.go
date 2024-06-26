/**
 * Created by Goland
 * @file   convertor.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/6/12 11:16
 * @desc   convertor.go
 */

package convertor

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/x-module/helper/json"
	"math"
	"reflect"
	"regexp"
	"strconv"
)

// func init() {
// 	extra.RegisterFuzzyDecoders()
// }

// ToBytes 将任意类型转换为字节切片
func ToBytes(value any) ([]byte, error) {
	v := reflect.ValueOf(value)

	switch value.(type) {
	case int, int8, int16, int32, int64:
		number := v.Int()
		buf := bytes.NewBuffer([]byte{})
		buf.Reset()
		err := binary.Write(buf, binary.BigEndian, number)
		return buf.Bytes(), err
	case uint, uint8, uint16, uint32, uint64:
		number := v.Uint()
		buf := bytes.NewBuffer([]byte{})
		buf.Reset()
		err := binary.Write(buf, binary.BigEndian, number)
		return buf.Bytes(), err
	case float32:
		number := float32(v.Float())
		bits := math.Float32bits(number)
		bytes := make([]byte, 4)
		binary.BigEndian.PutUint32(bytes, bits)
		return bytes, nil
	case float64:
		number := v.Float()
		bits := math.Float64bits(number)
		bytes := make([]byte, 8)
		binary.BigEndian.PutUint64(bytes, bits)
		return bytes, nil
	case bool:
		return strconv.AppendBool([]byte{}, v.Bool()), nil
	case string:
		return []byte(v.String()), nil
	case []byte:
		return v.Bytes(), nil
	default:
		newValue, err := json.Marshal(value)
		return newValue, err
	}
}

// ToChar 转换字符串为字符切片
func ToChar(s string) []string {
	c := make([]string, 0)
	if len(s) == 0 {
		c = append(c, "")
	}
	for _, v := range s {
		c = append(c, string(v))
	}
	return c
}

// ToJson 转换为json字符串
func ToJson(value any) (string, error) {
	result, err := json.Marshal(value)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

// ToJsonString 转换为json字符串
func ToJsonString(value any) string {
	result, err := json.Marshal(value)
	if err != nil {
		return ""
	}
	return string(result)
}

// ToFormatJsonString 格式化输出json字符串
func ToFormatJsonString(value any) string {
	result, err := json.Marshal(value, json.FormatString())
	if err != nil {
		return ""
	}
	return string(result)
}

// ToFloat 将任意类型转换为浮点数
func ToFloat(value any) (float64, error) {
	v := reflect.ValueOf(value)

	result := 0.0
	err := fmt.Errorf("ToInt: unvalid interface type %T", value)
	switch value.(type) {
	case int, int8, int16, int32, int64:
		result = float64(v.Int())
		return result, nil
	case uint, uint8, uint16, uint32, uint64:
		result = float64(v.Uint())
		return result, nil
	case float32, float64:
		result = v.Float()
		return result, nil
	case string:
		result, err = strconv.ParseFloat(v.String(), 64)
		if err != nil {
			result = 0.0
		}
		return result, err
	default:
		return result, err
	}
}

// ToInt 将任意类型转换为整数
func ToInt(value any) (int64, error) {
	v := reflect.ValueOf(value)

	var result int64
	err := fmt.Errorf("ToInt: invalid value type %T", value)
	switch value.(type) {
	case int, int8, int16, int32, int64:
		result = v.Int()
		return result, nil
	case uint, uint8, uint16, uint32, uint64:
		result = int64(v.Uint())
		return result, nil
	case float32, float64:
		result = int64(v.Float())
		return result, nil
	case string:
		result, err = strconv.ParseInt(v.String(), 0, 64)
		if err != nil {
			result = 0
		}
		return result, err
	default:
		return result, err
	}
}

// ToMap 将结构体切片转换为基于iteree函数的映射
func ToMap[T any, K comparable, V any](array []T, iteratee func(T) (K, V)) map[K]V {
	result := make(map[K]V, len(array))
	for _, item := range array {
		k, v := iteratee(item)
		result[k] = v
	}
	return result
}

// StructToMap 将结构转换为映射，仅转换导出的结构字段映射键指定与结构字段标签' json '值相同。
func StructToMap(value any) (map[string]any, error) {
	v := reflect.ValueOf(value)
	t := reflect.TypeOf(value)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return nil, fmt.Errorf("data type %T not support, shuld be struct or pointer to struct", value)
	}

	result := make(map[string]any)

	fieldNum := t.NumField()
	pattern := `^[A-Z]`
	regex := regexp.MustCompile(pattern)
	for i := 0; i < fieldNum; i++ {
		name := t.Field(i).Name
		tag := t.Field(i).Tag.Get("json")
		if regex.MatchString(name) && tag != "" {
			// result[name] = v.Field(i).Interface()
			result[tag] = v.Field(i).Interface()
		}
	}

	return result, nil
}

// MapToSlice 基于迭代函数将map转换为slice
func MapToSlice[T any, K comparable, V any](aMap map[K]V, iteratee func(K, V) T) []T {
	result := make([]T, 0, len(aMap))
	for k, v := range aMap {
		result = append(result, iteratee(k, v))
	}
	return result
}

// EncodeByte 将数据编码为字节片。
func EncodeByte(data any) ([]byte, error) {
	buffer := bytes.NewBuffer(nil)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(data)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

// DecodeByte 解码字节片数据到目标对象
func DecodeByte(data []byte, target any) error {
	buffer := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buffer)
	return decoder.Decode(target)
}

// FormatFileSize 字节的单位转换 保留两位小数
func FormatFileSize(fileSize float64) (size string) {
	if fileSize < 1024 {
		// return strconv.FormatInt(fileSize, 10) + "B"
		return fmt.Sprintf("%.2fB", fileSize/float64(1))
	} else if fileSize < (1024 * 1024) {
		return fmt.Sprintf("%.2fKB", fileSize/float64(1024))
	} else if fileSize < (1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fMB", fileSize/float64(1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fGB", fileSize/float64(1024*1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fTB", fileSize/float64(1024*1024*1024*1024))
	} else { // if fileSize < (1024 * 1024 * 1024 * 1024 * 1024 * 1024)
		return fmt.Sprintf("%.2fEB", fileSize/float64(1024*1024*1024*1024*1024))
	}
}

// TransInterfaceToStruct 转换interface 到struct
func TransInterfaceToStruct(params any, v any) error {
	jsonData, err := json.Marshal(params)
	if err != nil {
		return err
	}
	err = jsoniter.UnmarshalFromString(string(jsonData), v)
	if err != nil {
		return err
	}
	return err
}

// ByteToHex byte转16进制字符串
func ByteToHex(data []byte) string {
	buffer := new(bytes.Buffer)
	for _, b := range data {

		s := strconv.FormatInt(int64(b&0xff), 16)
		if len(s) == 1 {
			buffer.WriteString("0")
		}
		buffer.WriteString(s)
	}

	return buffer.String()
}

// HexToBye 16进制字符串转[]byte
func HexToBye(hex string) []byte {
	length := len(hex) / 2
	slice := make([]byte, length)
	rs := []rune(hex)

	for i := 0; i < length; i++ {
		s := string(rs[i*2 : i*2+2])
		value, _ := strconv.ParseInt(s, 16, 10)
		slice[i] = byte(value & 0xFF)
	}
	return slice
}

// StrToInt string转int
func StrToInt(v string) int {
	i, err := strconv.Atoi(v)
	if err != nil {
		return 0
	}
	return i
}

// StrToInt8 string转int8
func StrToInt8(v string) int8 {
	i, err := strconv.ParseInt(v, 10, 8)
	if err != nil {
		return 0
	}
	return int8(i)
}

// StrToInt16 string转int16
func StrToInt16(v string) int16 {
	i, err := strconv.ParseInt(v, 10, 16)
	if err != nil {
		return 0
	}
	return int16(i)
}

// StrToInt32 string转int32
func StrToInt32(v string) int32 {
	i, err := strconv.ParseInt(v, 10, 32)
	if err != nil {
		return 0
	}
	return int32(i)
}

// StrToInt64 string转int64
func StrToInt64(v string) int64 {
	i, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return 0
	}
	return i
}

// dec2Bin 将十进制转换为二进制字符串.
func dec2Bin(num int64) string {
	return strconv.FormatInt(num, 2)
}

// bin2Dec 将二进制字符串转换为十进制.
func bin2Dec(str string) (int64, error) {
	i, err := strconv.ParseInt(str, 2, 0)
	if err != nil {
		return 0, err
	}
	return i, nil
}

// hex2Bin 将十六进制字符串转换为二进制字符串.
func hex2Bin(str string) (string, error) {
	i, err := strconv.ParseInt(str, 16, 0)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, 2), nil
}

// bin2Hex 将二进制字符串转换为十六进制字符串.
func bin2Hex(str string) (string, error) {
	i, err := strconv.ParseInt(str, 2, 0)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, 16), nil
}

// dec2Hex 将十进制转换为十六进制.
func dec2Hex(num int64) string {
	return strconv.FormatInt(num, 16)
}

// hex2Dec 将十六进制转换为十进制.
func hex2Dec(str string) (int64, error) {
	start := 0
	if len(str) > 2 && str[0:2] == "0x" {
		start = 2
	}

	// bitSize 表示结果的位宽（包括符号位），0 表示最大位宽
	return strconv.ParseInt(str[start:], 16, 0)
}

// hex2Byte 16进制字符串转字节切片.
func hex2Byte(str string) ([]byte, error) {
	start := 0
	if len(str) > 2 && str[0:2] == "0x" {
		start = 2
	}

	h, e := hex.DecodeString(str[start:])
	return h, e
}

// dec2Oct 将十进制转换为八进制.
func dec2Oct(num int64) string {
	return strconv.FormatInt(num, 8)
}

// oct2Dec 将八进制转换为十进制.
func oct2Dec(str string) (int64, error) {
	start := 0
	if len(str) > 1 && str[0:1] == "0" {
		start = 1
	}

	return strconv.ParseInt(str[start:], 8, 0)
}

// Img2Base64 将图片字节转换为base64字符串.imgType为图片扩展名.
func Img2Base64(content []byte, imgType string) string {
	return fmt.Sprintf("data:image/%s;base64,%s", imgType, base64.StdEncoding.EncodeToString(content))
}
