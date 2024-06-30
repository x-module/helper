/**
 * Created by Goland
 * @file   random.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/6/12 14:37
 * @desc   random.go
 */

package sliceutil

import (
	"math"
	"math/rand"
	"slices"
)

// ContainSubSlice 检查切片是否包含给定的子切片。
func ContainSubSlice[T comparable](slice, subSlice []T) bool {
	for _, v := range subSlice {
		if !slices.Contains(slice, v) {
			return false
		}
	}
	return true
}

// Chunk Chunk创建一个元素切片，这些元素被分成大小相同的组。
func Chunk[T any](slice []T, size int) [][]T {
	result := [][]T{}
	if len(slice) == 0 || size <= 0 {
		return result
	}
	for _, item := range slice {
		l := len(result)
		if l == 0 || len(result[l-1]) == size {
			result = append(result, []T{})
			l++
		}
		result[l-1] = append(result[l-1], item)
	}
	return result
}

// Difference 创建一个slice，其元素在slice中，但不在comparedSlice中。
func Difference[T comparable](slice, comparedSlice []T) []T {
	var result []T
	for _, v := range slice {
		if !slices.Contains(comparedSlice, v) {
			result = append(result, v)
		}
	}
	return result
}

// DifferenceBy 它接受为slice的每个元素调用的iteratee和值来生成比较它们的标准。
func DifferenceBy[T comparable](slice []T, comparedSlice []T, iteratee func(index int, item T) T) []T {
	originSliceAfterMap := Every(slice, iteratee)
	comparedSliceAfterMap := Every(comparedSlice, iteratee)
	result := make([]T, 0)
	for i, v := range originSliceAfterMap {
		if !slices.Contains(comparedSliceAfterMap, v) {
			result = append(result, slice[i])
		}
	}
	return result
}

// DifferenceWith 接受比较器，调用比较器将slice的元素与值进行比较。结果值的顺序和引用由第一个切片决定。 比较器通过两个参数调用:(arrVal, othVal)。
func DifferenceWith[T any](slice []T, comparedSlice []T, comparator func(item1, item2 T) bool) []T {
	result := make([]T, 0)
	getIndex := func(arr []T, item T, comparison func(v1, v2 T) bool) int {
		index := -1
		for i, v := range arr {
			if comparison(item, v) {
				index = i
				break
			}
		}
		return index
	}
	for i, v := range slice {
		index := getIndex(comparedSlice, v, comparator)
		if index == -1 {
			result = append(result, slice[i])
		}
	}
	return result
}

// None 如果片中的所有值都不符合标准，则返回true。
func None[T any](slice []T, predicate func(index int, item T) bool) bool {
	l := 0
	for i, v := range slice {
		if !predicate(i, v) {
			l++
		}
	}
	return l == len(slice)
}

// Some 如果列表中的任何值通过谓词函数，则返回true。
func Some[T any](slice []T, predicate func(index int, item T) bool) bool {
	for i, v := range slice {
		if predicate(i, v) {
			return true
		}
	}
	return false
}

// Filter 遍历slice的元素，返回传递谓词函数的所有元素的slice。
func Filter[T any](slice []T, predicate func(index int, item T) bool) []T {
	result := make([]T, 0)
	for i, v := range slice {
		if predicate(i, v) {
			result = append(result, v)
		}
	}
	return result
}

// Count 返回给定项在切片中出现的次数。
func Count[T comparable](slice []T, item T) int {
	count := 0
	for _, v := range slice {
		if item == v {
			count++
		}
	}
	return count
}

// CountBy 用谓词函数遍历slice的元素，返回所有匹配元素的个数。
func CountBy[T any](slice []T, predicate func(index int, item T) bool) int {
	count := 0
	for i, v := range slice {
		if predicate(i, v) {
			count++
		}
	}
	return count
}

// GroupBy 迭代片的元素，每个元素将按标准分组，返回两个片。
func GroupBy[T any](slice []T, groupFn func(index int, item T) bool) ([]T, []T) {
	if len(slice) == 0 {
		return make([]T, 0), make([]T, 0)
	}
	groupB := make([]T, 0)
	groupA := make([]T, 0)
	for i, v := range slice {
		ok := groupFn(i, v)
		if ok {
			groupA = append(groupA, v)
		} else {
			groupB = append(groupB, v)
		}
	}
	return groupA, groupB
}

// GroupWith 返回由slice thru迭代器中每个元素运行结果生成的键组成的映射。
func GroupWith[T any, U comparable](slice []T, iteratee func(item T) U) map[U][]T {
	result := make(map[U][]T)
	for _, v := range slice {
		key := iteratee(v)
		if _, ok := result[key]; !ok {
			result[key] = []T{}
		}
		result[key] = append(result[key], v)
	}
	return result
}

// FindFirst 遍历slice的元素，返回第一个通过谓词函数真值测试的元素。 如果return T为nil，则没有匹配谓词函数的项。
func FindFirst[T any](slice []T, predicate func(index int, item T) bool) (*T, bool) {
	index := -1
	for i, v := range slice {
		if predicate(i, v) {
			index = i
			break
		}
	}
	if index == -1 {
		return nil, false
	}
	return &slice[index], true
}

// FindLast 遍历slice的元素，返回最后一个通过谓词函数真值测试的元素。 如果return T为nil，则没有匹配谓词函数的项。
func FindLast[T any](slice []T, predicate func(index int, item T) bool) (*T, bool) {
	index := -1
	for i := len(slice) - 1; i >= 0; i-- {
		if predicate(i, slice[i]) {
			index = i
			break
		}
	}
	if index == -1 {
		return nil, false
	}
	return &slice[index], true
}

// Every 通过运行slice thru迭代函数的每个元素来创建一个值片。
func Every[T any, U any](slice []T, iteratee func(index int, item T) U) []U {
	result := make([]U, len(slice), cap(slice))
	for i, v := range slice {
		result[i] = iteratee(i, v)
	}
	return result
}

// Foreach 通过运行slice thru迭代函数的每个元素来创建一个值片。
func Foreach[T any](slice []T, iteratee func(index int, item T)) {
	for i, v := range slice {
		iteratee(i, v)
	}
}

// Replace 返回切片的副本，其中旧的前n个不重叠的实例替换为new。
func Replace[T comparable](slice []T, old T, new T, n int) []T {
	result := make([]T, len(slice))
	copy(result, slice)
	for i := range result {
		if result[i] == old && n != 0 {
			result[i] = new
			n--
		}
	}
	return result
}

// ReplaceAll 返回片的副本，其中所有不重叠的old实例替换为new。
func ReplaceAll[T comparable](slice []T, old T, new T) []T {
	return Replace(slice, old, new, -1)
}

// Repeat 创建一个长度为n的slice，其元素参数为item。
func Repeat[T any](item T, n int) []T {
	result := make([]T, n)
	for i := range result {
		result[i] = item
	}
	return result
}

// Delete 删除从开始索引到结束索引- 1的切片元素。
func Delete[T any](slice []T, start int, end ...int) []T {
	size := len(slice)
	if start < 0 || start >= size {
		return slice
	}
	if len(end) > 0 {
		end := end[0]
		if end <= start {
			return slice
		}
		if end > size {
			end = size
		}
		slice = append(slice[:start], slice[end:]...)
		return slice
	}
	if start == size-1 {
		slice = slice[:start]
	} else {
		slice = append(slice[:start], slice[start+1:]...)
	}
	return slice
}

// Drop 创建一个切片，当n > 0时从开始删除n个元素，或者当n < 0时从结束删除n个元素。
func Drop[T any](slice []T, n int) []T {
	size := len(slice)
	if size == 0 || n == 0 {
		return slice
	}
	if math.Abs(float64(n)) >= float64(size) {
		return []T{}
	}
	if n < 0 {
		return slice[0 : size+n]
	}
	return slice[n:size]
}

// Unique 唯一删除重复元素的切片。
func Unique[T comparable](slice []T) []T {
	var result []T
	for _, v := range slice {
		if !slices.Contains(result, v) {
			result = append(result, v)
		}
	}
	return result
}

// UniqueBy 对slice的每一项调用iteratee函数判断重复项。
func UniqueBy[T comparable](slice []T, iteratee func(key int, item T) bool) []T {
	var result []T
	for k, v := range slice {
		if iteratee(k, v) {
			result = append(result, v)
		}
	}
	return result
}

// Merge 创建一个删除所有假值的切片。值false、nil、0和""为false。
func Merge[T any](slices ...[]T) []T {
	result := make([]T, 0)
	for _, v := range slices {
		result = append(result, v...)
	}
	return result
}

// Union 从所有给定的切片中按顺序创建一个唯一元素的切片。
func Union[T comparable](slices ...[]T) []T {
	var result []T
	contain := map[T]struct{}{}
	for _, slice := range slices {
		for _, item := range slice {
			if _, ok := contain[item]; !ok {
				contain[item] = struct{}{}
				result = append(result, item)
			}
		}
	}
	return result
}

// UnionBy 类似于Union，但它接受intersection，每个slice的每个元素都会被调用。
func UnionBy[T any, V comparable](predicate func(item T) V, slices ...[]T) []T {
	var result []T
	contain := map[V]struct{}{}
	for _, slice := range slices {
		for _, item := range slice {
			val := predicate(item)
			if _, ok := contain[val]; !ok {
				contain[val] = struct{}{}
				result = append(result, item)
			}
		}
	}
	return result
}

// Intersection 创建一个包含所有切片的唯一元素的切片。
func Intersection[T comparable](slices ...[]T) []T {
	if len(slices) == 0 {
		return []T{}
	}
	if len(slices) == 1 {
		return Unique(slices[0])
	}
	reducer := func(sliceA, sliceB []T) []T {
		hashMap := make(map[T]int)
		for _, v := range sliceA {
			hashMap[v] = 1
		}
		out := make([]T, 0)
		for _, val := range sliceB {
			if v, ok := hashMap[val]; v == 1 && ok {
				out = append(out, val)
				hashMap[val]++
			}
		}
		return out
	}
	result := reducer(slices[0], slices[1])
	reduceSlice := make([][]T, 2)
	for i := 2; i < len(slices); i++ {
		reduceSlice[0] = result
		reduceSlice[1] = slices[i]
		result = reducer(reduceSlice[0], reduceSlice[1])
	}
	return result
}

// Reverse 返回元素顺序的slice。
func Reverse[T any](slice []T) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// Shuffle 洗牌切片。
func Shuffle[T any](slice []T) []T {
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
	return slice
}

// Without 创建一个不包含所有给定项的slice。
func Without[T comparable](slice []T, items ...T) []T {
	if len(items) == 0 || len(slice) == 0 {
		return slice
	}
	result := make([]T, 0, len(slice))
	for _, v := range slice {
		if !slices.Contains(items, v) {
			result = append(result, v)
		}
	}
	return result
}

// ToSlicePointer 返回一个指向变量参数转换切片的指针。
func ToSlicePointer[T any](items ...T) []*T {
	result := make([]*T, len(items))
	for i := range items {
		result[i] = &items[i]
	}
	return result
}

// AppendIfAbsent 只添加不存在的项。
func AppendIfAbsent[T comparable](slice []T, item T) []T {
	if !slices.Contains(slice, item) {
		slice = append(slice, item)
	}
	return slice
}

// ToMap 根据回调函数将切片转换为映射。
func ToMap[T any, U comparable](slice []T, iteratee func(index int, item T) U) map[U]T {
	result := make(map[U]T, len(slice))
	for k, v := range slice {
		k := iteratee(k, v)
		result[k] = v
	}
	return result
}

func Compact[T comparable](slice []T) []T {
	var zero T
	result := []T{}
	for _, v := range slice {
		if v != zero {
			result = append(result, v)
		}
	}
	return result
}
