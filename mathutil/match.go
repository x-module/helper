/**
 * Created by Goland
 * @file   match.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/6/12 13:56
 * @desc   match.go
 */

package mathutil

import (
	"golang.org/x/exp/constraints"
)

// Average 计算平均值
func Average[T constraints.Integer | constraints.Float](numbers ...T) T {
	var sum T
	n := T(len(numbers))
	for _, v := range numbers {
		sum += v
	}
	return sum / n
}
