/**
 * Created by Goland
 * @file   map.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/6/11 19:56
 * @desc   map.go
 */

package maputil

import (
	"reflect"
)

// Merge 合并多个 map,后面的相同key将被覆盖
func Merge[K comparable, V any](maps ...map[K]V) map[K]V {
	result := make(map[K]V, 0)
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}

// ForEach 对map中的每个键和值对执行迭代函数
func ForEach[K comparable, V any](m map[K]V, iteratee func(key K, value V)) {
	for k, v := range m {
		iteratee(k, v)
	}
}

// Filter 过滤器遍历map，返回一个包含所有键和值对的新map，通过谓词函数
func Filter[K comparable, V any](m map[K]V, predicate func(key K, value V) bool) map[K]V {
	result := make(map[K]V)
	for k, v := range m {
		if predicate(k, v) {
			result[k] = v
		}
	}
	return result
}

// Intersect 求maps的交集， Intersect遍历映射，返回所有给定映射中的键和值对的新映射。
func Intersect[K comparable, V any](maps ...map[K]V) map[K]V {
	if len(maps) == 0 {
		return map[K]V{}
	}
	if len(maps) == 1 {
		return maps[0]
	}
	var result map[K]V
	reducer := func(m1, m2 map[K]V) map[K]V {
		m := make(map[K]V)
		for k, v1 := range m1 {
			if v2, ok := m2[k]; ok && reflect.DeepEqual(v1, v2) {
				m[k] = v1
			}
		}
		return m
	}
	reduceMaps := make([]map[K]V, 2)
	result = reducer(maps[0], maps[1])
	for i := 2; i < len(maps); i++ {
		reduceMaps[0] = result
		reduceMaps[1] = maps[i]
		result = reducer(reduceMaps[0], reduceMaps[1])
	}
	return result
}

// Minus 创建一个键在mapA中但不在mapB中的映射。
func Minus[K comparable, V any](mapA, mapB map[K]V) map[K]V {
	result := make(map[K]V)
	for k, v := range mapA {
		if _, ok := mapB[k]; !ok {
			result[k] = v
		}
	}
	return result
}

// IsDisjoint  如果两个map没有相同的键，则它们是不相交的。
func IsDisjoint[K comparable, V any](mapA, mapB map[K]V) bool {
	for k := range mapA {
		if _, ok := mapB[k]; ok {
			return false
		}
	}

	return true
}

// Union 求两个map的并集, Union遍历映射，返回所有给定映射中的键和值对的新映射。后面相同的值会覆盖之前的
func Union[K comparable, V any](maps ...map[K]V) map[K]V {
	result := make(map[K]V)
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}
