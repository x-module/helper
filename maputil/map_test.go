/**
 * Created by Goland
 * @file   map.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/6/11 19:56
 * @desc   map.go
 */

package maputil

import (
	"github.com/x-module/helper/internal"
	"testing"
)

func TestMerge(t *testing.T) {
	assert := internal.NewAssert(t, "TestMerge")

	m1 := map[int]string{
		1: "a",
		2: "b",
	}
	m2 := map[int]string{
		1: "1",
		3: "2",
	}

	expected := map[int]string{
		1: "1",
		2: "b",
		3: "2",
	}
	result := Merge(m1, m2)
	assert.Equal(expected, result)
}

func TestForEach(t *testing.T) {
	assert := internal.NewAssert(t, "TestForEach")
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
	}
	var sum int
	ForEach(m, func(_ string, value int) {
		sum += value
	})
	assert.Equal(10, sum)
}

func TestFilter(t *testing.T) {
	assert := internal.NewAssert(t, "TestFilter")
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
	}
	isEven := func(_ string, value int) bool {
		return value%2 == 0
	}
	result := Filter(m, isEven)
	assert.Equal(map[string]int{
		"b": 2,
		"d": 4,
	}, result)
}

func TestIntersect(t *testing.T) {
	assert := internal.NewAssert(t, "TestIntersect")
	m1 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	m2 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 6,
		"d": 7,
	}

	m3 := map[string]int{
		"a": 1,
		"b": 9,
		"e": 9,
	}
	assert.Equal(map[string]int{"a": 1, "b": 2, "c": 3}, Intersect(m1))
	assert.Equal(map[string]int{"a": 1, "b": 2}, Intersect(m1, m2))
	assert.Equal(map[string]int{"a": 1}, Intersect(m1, m2, m3))
}
func TestUnion(t *testing.T) {
	assert := internal.NewAssert(t, "TestIntersect")
	m1 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	m2 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 6,
		"d": 7,
	}

	m3 := map[string]int{
		"a": 1,
		"b": 9,
		"e": 9,
	}
	assert.Equal(map[string]int{"a": 1, "b": 2, "c": 3}, Union(m1))
	assert.Equal(map[string]int{"a": 1, "b": 2, "c": 6, "d": 7}, Union(m1, m2))
	assert.Equal(map[string]int{"a": 1, "b": 9, "c": 6, "d": 7, "e": 9}, Union(m1, m2, m3))
}
