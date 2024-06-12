package convertor

import (
	"github.com/x-module/helper/internal"
	"slices"
	"strconv"
	"testing"
)

func TestToChar(t *testing.T) {
	assert := internal.NewAssert(t, "TestToChar")
	cases := []string{"", "abc", "1 2#3"}
	expected := [][]string{
		{""},
		{"a", "b", "c"},
		{"1", " ", "2", "#", "3"},
	}
	for i := 0; i < len(cases); i++ {
		assert.Equal(expected[i], ToChar(cases[i]))
	}
}

func TestToInt(t *testing.T) {
	assert := internal.NewAssert(t, "TestToInt")

	cases := []any{"123", "-123", 123,
		uint(123), uint8(123), uint16(123), uint32(123), uint64(123),
		float32(12.3), float64(12.3),
		"abc", false, "111111111111111111111111111111111111111"}

	expected := []int64{123, -123, 123, 123, 123, 123, 123, 123, 12, 12, 0, 0, 0}

	for i := 0; i < len(cases); i++ {
		actual, _ := ToInt(cases[i])
		assert.Equal(expected[i], actual)
	}
}

func TestToFloat(t *testing.T) {
	assert := internal.NewAssert(t, "TestToFloat")

	cases := []any{
		"", "-1", "-.11", "1.23e3", ".123e10", "abc",
		int(0), int8(1), int16(-1), int32(123), int64(123),
		uint(123), uint8(123), uint16(123), uint32(123), uint64(123),
		float64(12.3), float32(12.3),
	}
	expected := []float64{0, -1, -0.11, 1230, 0.123e10, 0,
		0, 1, -1, 123, 123, 123, 123, 123, 123, 123, 12.3, 12.300000190734863}

	for i := 0; i < len(cases); i++ {
		actual, _ := ToFloat(cases[i])
		assert.Equal(expected[i], actual)
	}
}

func TestToJson(t *testing.T) {
	assert := internal.NewAssert(t, "TestToJson")

	var aMap = map[string]int{"a": 1, "b": 2, "c": 3}
	mapJsonStr, _ := ToJson(aMap)
	assert.Equal("{\"a\":1,\"b\":2,\"c\":3}", mapJsonStr)

	type TestStruct struct {
		Name string
	}
	aStruct := TestStruct{Name: "TestStruct"}
	structJsonStr, _ := ToJson(aStruct)
	assert.Equal("{\"Name\":\"TestStruct\"}", structJsonStr)
}

func TestToMap(t *testing.T) {
	assert := internal.NewAssert(t, "TestToMap")

	type Message struct {
		name string
		code int
	}
	messages := []Message{
		{name: "Hello", code: 100},
		{name: "Hi", code: 101},
	}
	result := ToMap(messages, func(msg Message) (int, string) {
		return msg.code, msg.name
	})
	expected := map[int]string{100: "Hello", 101: "Hi"}

	assert.Equal(expected, result)
}

func TestStructToMap(t *testing.T) {
	assert := internal.NewAssert(t, "TestStructToMap")

	type People struct {
		Name string `json:"name"`
		age  int
	}
	p := People{
		"test",
		100,
	}
	pm, _ := StructToMap(p)

	expected := map[string]any{"name": "test"}
	assert.Equal(expected, pm)
}

func TestMapToSlice(t *testing.T) {
	assert := internal.NewAssert(t, "TestMapToSlice")

	aMap := map[string]int{"a": 1, "b": 2, "c": 3}
	result := MapToSlice(aMap, func(key string, value int) string {
		return key + ":" + strconv.Itoa(value)
	})

	assert.Equal(3, len(result))
	assert.Equal(true, slices.Contains(result, "a:1"))
	assert.Equal(true, slices.Contains(result, "b:2"))
	assert.Equal(true, slices.Contains(result, "c:3"))
}

func TestEncodeByte(t *testing.T) {
	assert := internal.NewAssert(t, "TestEncodeByte")

	byteData, _ := EncodeByte("abc")
	expected := []byte{6, 12, 0, 3, 97, 98, 99}
	assert.Equal(expected, byteData)
}

func TestDecodeByte(t *testing.T) {
	assert := internal.NewAssert(t, "TestDecodeByte")
	var obj string
	byteData := []byte{6, 12, 0, 3, 97, 98, 99}
	err := DecodeByte(byteData, &obj)
	assert.IsNil(err)
	assert.Equal("abc", obj)
}
