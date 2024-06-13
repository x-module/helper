/**
 * Created by Goland
 * @file   random.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/6/12 14:37
 * @desc   random.go
 */

package sliceutil

import (
	"github.com/x-module/helper/internal"
	"math"
	"testing"
)

func TestContainSubSlice(t *testing.T) {
	assert := internal.NewAssert(t, "TestContainSubSlice")
	assert.Equal(true, ContainSubSlice([]string{"a", "a", "b", "c"}, []string{"a", "a"}))
	assert.Equal(false, ContainSubSlice([]string{"a", "a", "b", "c"}, []string{"a", "d"}))
	assert.Equal(true, ContainSubSlice([]int{1, 2, 3}, []int{1, 2}))
	assert.Equal(false, ContainSubSlice([]int{1, 2, 3}, []int{0, 1}))
}

func TestChunk(t *testing.T) {
	assert := internal.NewAssert(t, "TestChunk")

	arr := []string{"a", "b", "c", "d", "e"}

	assert.Equal([][]string{}, Chunk(arr, -1))

	assert.Equal([][]string{}, Chunk(arr, 0))

	r1 := [][]string{{"a"}, {"b"}, {"c"}, {"d"}, {"e"}}
	assert.Equal(r1, Chunk(arr, 1))

	r2 := [][]string{{"a", "b"}, {"c", "d"}, {"e"}}
	assert.Equal(r2, Chunk(arr, 2))

	r3 := [][]string{{"a", "b", "c"}, {"d", "e"}}
	assert.Equal(r3, Chunk(arr, 3))

	r4 := [][]string{{"a", "b", "c", "d"}, {"e"}}
	assert.Equal(r4, Chunk(arr, 4))

	r5 := [][]string{{"a", "b", "c", "d", "e"}}
	assert.Equal(r5, Chunk(arr, 5))

	r6 := [][]string{{"a", "b", "c", "d", "e"}}
	assert.Equal(r6, Chunk(arr, 6))
}

func TestEvery(t *testing.T) {
	nums := []int{1, 2, 3, 5}
	isEven := func(i, num int) bool {
		return num%2 == 0
	}

	assert := internal.NewAssert(t, "TestEvery")
	assert.Equal(false, Foreach(nums, isEven))
}

func TestNone(t *testing.T) {
	nums := []int{1, 2, 3, 5}
	check := func(i, num int) bool {
		return num%2 == 1
	}

	assert := internal.NewAssert(t, "TestNone")
	assert.Equal(false, None(nums, check))
}

func TestSome(t *testing.T) {
	nums := []int{1, 2, 3, 5}
	hasEven := func(i, num int) bool {
		return num%2 == 0
	}

	assert := internal.NewAssert(t, "TestSome")
	assert.Equal(true, Some(nums, hasEven))
}

func TestFilter(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	isEven := func(i, num int) bool {
		return num%2 == 0
	}

	assert := internal.NewAssert(t, "TestFilter")
	assert.Equal([]int{2, 4}, Filter(nums, isEven))

	type student struct {
		name string
		age  int
	}
	students := []student{
		{"a", 10},
		{"b", 11},
		{"c", 12},
		{"d", 13},
		{"e", 14},
	}
	studentsOfAageGreat12 := []student{
		{"d", 13},
		{"e", 14},
	}
	filterFunc := func(i int, s student) bool {
		return s.age > 12
	}

	assert.Equal(studentsOfAageGreat12, Filter(students, filterFunc))
}

func TestGroupBy(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6}
	evenFunc := func(i, num int) bool {
		return (num % 2) == 0
	}
	expectedEven := []int{2, 4, 6}
	expectedOdd := []int{1, 3, 5}
	even, odd := GroupBy(nums, evenFunc)

	assert := internal.NewAssert(t, "TestGroupBy")
	assert.Equal(expectedEven, even)
	assert.Equal(expectedOdd, odd)
}

func TestGroupWith(t *testing.T) {
	nums := []float64{6.1, 4.2, 6.3}
	floor := func(num float64) float64 {
		return math.Floor(num)
	}
	expected := map[float64][]float64{
		4: {4.2},
		6: {6.1, 6.3},
	}
	actual := GroupWith(nums, floor)
	assert := internal.NewAssert(t, "TestGroupWith")
	assert.Equal(expected, actual)
}

func TestCount(t *testing.T) {
	numbers := []int{1, 2, 3, 3, 5, 6}

	assert := internal.NewAssert(t, "TestCountBy")

	assert.Equal(1, Count(numbers, 1))
	assert.Equal(2, Count(numbers, 3))
}

func TestCountBy(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6}
	evenFunc := func(i, num int) bool {
		return (num % 2) == 0
	}

	assert := internal.NewAssert(t, "TestCountBy")
	assert.Equal(3, CountBy(nums, evenFunc))
}

func TestFind(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	even := func(i, num int) bool {
		return num%2 == 0
	}
	res, ok := FindFirst(nums, even)
	if !ok {
		t.Fatal("found nothing")
	}

	assert := internal.NewAssert(t, "TestFind")
	assert.Equal(2, *res)
}

func TestFindLast(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	even := func(i, num int) bool {
		return num%2 == 0
	}
	res, ok := FindLast(nums, even)
	if !ok {
		t.Fatal("found nothing")
	}

	assert := internal.NewAssert(t, "TestFindLast")
	assert.Equal(4, *res)
}

func TestFindFoundNothing(t *testing.T) {
	nums := []int{1, 1, 1, 1, 1, 1}
	findFunc := func(i, num int) bool {
		return num > 1
	}
	_, ok := FindFirst(nums, findFunc)
	// if ok {
	// 	t.Fatal("found something")
	// }
	assert := internal.NewAssert(t, "TestFindFoundNothing")
	assert.Equal(false, ok)
}

func TestMap(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	multiplyTwo := func(i, num int) int {
		return num * 2
	}

	assert := internal.NewAssert(t, "TestMap")
	assert.Equal([]int{2, 4, 6, 8}, ToMap(nums, multiplyTwo))

	type student struct {
		name string
		age  int
	}
	students := []student{
		{"a", 1},
		{"b", 2},
		{"c", 3},
	}
	studentsOfAdd10Aage := []student{
		{"a", 11},
		{"b", 12},
		{"c", 13},
	}
	mapFunc := func(i int, s student) student {
		s.age += 10
		return s
	}

	assert.Equal(studentsOfAdd10Aage, ToMap(students, mapFunc))
}

func TestDrop(t *testing.T) {
	assert := internal.NewAssert(t, "TestInterfaceSlice")

	assert.Equal([]int{}, Drop([]int{}, 0))
	assert.Equal([]int{}, Drop([]int{}, 1))
	assert.Equal([]int{}, Drop([]int{}, -1))

	assert.Equal([]int{1, 2, 3, 4, 5}, Drop([]int{1, 2, 3, 4, 5}, 0))
	assert.Equal([]int{2, 3, 4, 5}, Drop([]int{1, 2, 3, 4, 5}, 1))
	assert.Equal([]int{}, Drop([]int{1, 2, 3, 4, 5}, 5))
	assert.Equal([]int{}, Drop([]int{1, 2, 3, 4, 5}, 6))

	assert.Equal([]int{1, 2, 3, 4}, Drop([]int{1, 2, 3, 4, 5}, -1))
	assert.Equal([]int{}, Drop([]int{1, 2, 3, 4, 5}, -6))
	assert.Equal([]int{}, Drop([]int{1, 2, 3, 4, 5}, -6))
}

func TestUnique(t *testing.T) {
	assert := internal.NewAssert(t, "TestUnique")

	assert.Equal([]int{1, 2, 3}, Unique([]int{1, 2, 2, 3}))
	assert.Equal([]string{"a", "b", "c"}, Unique([]string{"a", "a", "b", "c"}))
}

func TestUnion(t *testing.T) {
	assert := internal.NewAssert(t, "TestUnion")

	s1 := []int{1, 3, 4, 6}
	s2 := []int{1, 2, 5, 6}
	s3 := []int{0, 4, 5, 7}

	assert.Equal([]int{1, 3, 4, 6, 2, 5, 0, 7}, Union(s1, s2, s3))
	assert.Equal([]int{1, 3, 4, 6, 2, 5}, Union(s1, s2))
	assert.Equal([]int{1, 3, 4, 6}, Union(s1))
}

func TestUnionBy(t *testing.T) {
	assert := internal.NewAssert(t, "TestUnionBy")

	testFunc := func(i int) int {
		return i / 2
	}

	result := UnionBy(testFunc, []int{0, 1, 2, 3, 4, 5}, []int{0, 2, 10})
	assert.Equal(result, []int{0, 2, 4, 10})
}

func TestMerge(t *testing.T) {
	assert := internal.NewAssert(t, "TestMerge")

	s1 := []int{1, 2, 3, 4}
	s2 := []int{2, 3, 4, 5}
	s3 := []int{4, 5, 6}

	assert.Equal([]int{1, 2, 3, 4, 2, 3, 4, 5, 4, 5, 6}, Merge(s1, s2, s3))
	assert.Equal([]int{1, 2, 3, 4, 2, 3, 4, 5}, Merge(s1, s2))
	assert.Equal([]int{2, 3, 4, 5, 4, 5, 6}, Merge(s2, s3))
}

func TestIntersection(t *testing.T) {
	s1 := []int{1, 2, 2, 3}
	s2 := []int{1, 2, 3, 4}
	s3 := []int{0, 2, 3, 5, 6}
	s4 := []int{0, 5, 6}

	expected := [][]int{
		{2, 3},
		{1, 2, 3},
		{1, 2, 3},
		{},
	}
	res := []any{
		Intersection(s1, s2, s3),
		Intersection(s1, s2),
		Intersection(s1),
		Intersection(s1, s4),
	}

	assert := internal.NewAssert(t, "TestIntersection")

	for i := 0; i < len(res); i++ {
		assert.Equal(expected[i], res[i])
	}
}

func TestReverse(t *testing.T) {
	assert := internal.NewAssert(t, "TestReverse")

	s1 := []int{1, 2, 3, 4, 5}
	Reverse(s1)
	assert.Equal([]int{5, 4, 3, 2, 1}, s1)

	s2 := []string{"a", "b", "c", "d", "e"}
	Reverse(s2)
	assert.Equal([]string{"e", "d", "c", "b", "a"}, s2)
}

func TestDifference(t *testing.T) {
	assert := internal.NewAssert(t, "TestDifference")

	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{4, 5, 6}
	assert.Equal([]int{1, 2, 3}, Difference(s1, s2))
}

func TestDifferenceWith(t *testing.T) {
	assert := internal.NewAssert(t, "TestDifferenceWith")

	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{4, 5, 6, 7, 8}
	isDouble := func(v1, v2 int) bool {
		return v2 == 2*v1
	}
	assert.Equal([]int{1, 5}, DifferenceWith(s1, s2, isDouble))
}

func TestDifferenceBy(t *testing.T) {
	assert := internal.NewAssert(t, "TestDifferenceBy")

	s1 := []int{1, 2, 3, 4, 5} // after add one: 2 3 4 5 6
	s2 := []int{3, 4, 5}       // after add one: 4 5 6
	addOne := func(i int, v int) int {
		return v + 1
	}
	assert.Equal([]int{1, 2}, DifferenceBy(s1, s2, addOne))
}

func TestWithout(t *testing.T) {
	assert := internal.NewAssert(t, "TestWithout")
	assert.Equal([]int{3, 4, 5}, Without([]int{1, 2, 3, 4, 5}, 1, 2))
	assert.Equal([]int{1, 2, 3, 4, 5}, Without([]int{1, 2, 3, 4, 5}))
}

func TestShuffle(t *testing.T) {
	assert := internal.NewAssert(t, "TestShuffle")

	s := []int{1, 2, 3, 4, 5}
	res := Shuffle(s)
	t.Log("Shuffle result: ", res)

	assert.Equal(5, len(res))
}

func TestToSlicePointer(t *testing.T) {
	assert := internal.NewAssert(t, "TestToSlicePointer")

	str1 := "a"
	str2 := "b"
	assert.Equal([]*string{&str1}, ToSlicePointer(str1))
	assert.Equal([]*string{&str1, &str2}, ToSlicePointer(str1, str2))
}

func TestAppendIfAbsent(t *testing.T) {
	assert := internal.NewAssert(t, "TestAppendIfAbsent")

	str1 := []string{"a", "b"}
	assert.Equal([]string{"a", "b"}, AppendIfAbsent(str1, "a"))
	assert.Equal([]string{"a", "b", "c"}, AppendIfAbsent(str1, "c"))
}

func TestReplace(t *testing.T) {
	assert := internal.NewAssert(t, "TestReplace")

	strs := []string{"a", "b", "a", "c", "d", "a"}

	assert.Equal([]string{"a", "b", "a", "c", "d", "a"}, Replace(strs, "a", "x", 0))
	assert.Equal([]string{"x", "b", "a", "c", "d", "a"}, Replace(strs, "a", "x", 1))
	assert.Equal([]string{"x", "b", "x", "c", "d", "a"}, Replace(strs, "a", "x", 2))
	assert.Equal([]string{"x", "b", "x", "c", "d", "x"}, Replace(strs, "a", "x", 3))
	assert.Equal([]string{"x", "b", "x", "c", "d", "x"}, Replace(strs, "a", "x", 4))

	assert.Equal([]string{"x", "b", "x", "c", "d", "x"}, Replace(strs, "a", "x", -1))
	assert.Equal([]string{"x", "b", "x", "c", "d", "x"}, Replace(strs, "a", "x", -2))

	assert.Equal([]string{"a", "b", "a", "c", "d", "a"}, Replace(strs, "x", "y", 1))
	assert.Equal([]string{"a", "b", "a", "c", "d", "a"}, Replace(strs, "x", "y", -1))
}

func TestReplaceAll(t *testing.T) {
	assert := internal.NewAssert(t, "TestReplaceAll")

	strs := []string{"a", "b", "a", "c", "d", "a"}

	assert.Equal([]string{"x", "b", "x", "c", "d", "x"}, ReplaceAll(strs, "a", "x"))
	assert.Equal([]string{"a", "b", "a", "c", "d", "a"}, ReplaceAll(strs, "e", "x"))
}

func TestKeyBy(t *testing.T) {
	assert := internal.NewAssert(t, "TestKeyBy")

	result := ToMap([]string{"a", "ab", "abc"}, func(_ int, str string) int {
		return len(str)
	})

	assert.Equal(result, map[int]string{1: "a", 2: "ab", 3: "abc"})
}

func TestRepeat(t *testing.T) {
	assert := internal.NewAssert(t, "TestRepeat")

	result := Repeat("a", 6)

	assert.Equal(result, []string{"a", "a", "a", "a", "a", "a"})
}
