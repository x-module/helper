### 操作切片的方法集合。

```go
import "github.com/x-module/utils/slice"
```

#### 函数列表:
- **<big>ContainSubSlice</big>** : 检查切片是否包含给定的子切片。
- **<big>Chunk Chunk</big>** :  创建一个元素切片，这些元素被分成大小相同的组。
- **<big>Difference </big>** : 创建一个slice，其元素在slice中，但不在comparedSlice中。
- **<big>DifferenceBy</big>** :  它接受为slice的每个元素调用的iteratee和值来生成比较它们的标准。
- **<big>DifferenceWith </big>** : 接受比较器，调用比较器将slice的元素与值进行比较。结果值的顺序和引用由第一个切片决定。 比较器通过两个参数调用:(arrVal, othVal)。
-  **<big>AppendIfAbsent</big>** : 当前切片中不包含值时，将该值追加到切片中。
- **<big>None </big>** : 如果片中的所有值都不符合标准，则返回true。
-  **<big>Contain</big>** : 判断slice是否包含value。
- **<big>Some </big>** : 如果列表中的任何值通过谓词函数，则返回true。
-  **<big>ContainSubSlice</big>** : 判断slice是否包含subslice。
- **<big>Filter </big>** : 遍历slice的元素，返回传递谓词函数的所有元素的slice。
-  **<big>Chunk</big>** : 按照size参数均分slice。
- **<big>Count </big>** : 返回给定项在切片中出现的次数。
-  **<big>Compact</big>** : 去除slice中的假值（false values are false, nil, 0, ""）。
- **<big>CountBy </big>** : 用谓词函数遍历slice的元素，返回所有匹配元素的个数。
-  **<big>Concat</big>** : 合并多个slices到一个slice中。
- **<big>GroupBy </big>** : 迭代片的元素，每个元素将按标准分组，返回两个片。
-  **<big>Count</big>** : 返回切片中指定元素的个数。
- **<big>GroupWith</big>** :  返回由slice thru迭代器中每个元素运行结果生成的键组成的映射。
-  **<big>CountBy</big>** : 遍历切片，对每个元素执行函数predicate. 返回符合函数返回值为true的元素的个数。
- **<big>FindFirst</big>** :  遍历slice的元素，返回第一个通过谓词函数真值测试的元素。 如果return T为nil，则没有匹配谓词函数的项。
-  **<big>Difference</big>** : 创建一个切片，其元素不包含在另一个给定切片中。
- **<big>FindLast </big>** : 遍历slice的元素，返回最后一个通过谓词函数真值测试的元素。 如果return T为nil，则没有匹配谓词函数的项。
-  **<big>DifferenceBy</big>** : 将两个slice中的每个元素调用iteratee函数，并比较它们的返回值，如果不相等返回在slice中对应的值。
- **<big>Foreach </big>** : 通过运行slice thru迭代函数的每个元素来创建一个值片。
-  **<big>DifferenceWith</big>** : 接受比较器函数，该比较器被调用以将切片的元素与值进行比较。 结果值的顺序和引用由第一个切片确定。
- **<big>Replace </big>** : 返回切片的副本，其中旧的前n个不重叠的实例替换为new。
-  **<big>DeleteAt</big>** : 删除切片中指定开始索引到结束索引的元素。
- **<big>ReplaceAll</big>** :  返回片的副本，其中所有不重叠的old实例替换为new。
-  **<big>Drop</big>** : 创建一个切片，当n > 0时从开头删除n个元素，或者当n < 0时从结尾删除n个元素。
- **<big>Repeat </big>** : 创建一个长度为n的slice，其元素参数为item。
-  **<big>Equal</big>** : 检查两个切片是否相等，相等条件：切片长度相同，元素顺序和值都相同。
- **<big>Delete </big>** : 删除从开始索引到结束索引- 1的切片元素。
-  **<big>EqualWith</big>** : 检查两个切片是否相等，相等条件：对两个切片的元素调用比较函数comparator，返回true。
- **<big>Drop </big>** : 创建一个切片，当n > 0时从开始删除n个元素，或者当n < 0时从结束删除n个元素。
-  **<big>Every</big>** : 如果切片中的所有值都通过谓词函数，则返回true。
- **<big>Unique</big>** :  唯一删除重复元素的切片。
-  **<big>Filter</big>** : 返回切片中通过predicate函数真值测试的所有元素。
- **<big>UniqueBy</big>** :  对slice的每一项调用iteratee函数判断重复项。
-  **<big>Find</big>** : 遍历切片的元素，返回第一个通过predicate函数真值测试的元素。
- **<big>Merge </big>** : 创建一个删除所有假值的切片。值false、nil、0和""为false。
-  **<big>FindLast</big>** : 从头到尾遍历slice的元素，返回最后一个通过predicate函数真值测试的元素。
- **<big>Union </big>** : 从所有给定的切片中按顺序创建一个唯一元素的切片。
-  **<big>Flatten</big>** : 将多维切片展平一层。
- **<big>UnionBy</big>** :  类似于Union，但它接受intersection，每个slice的每个元素都会被调用。
-  **<big>FlattenDeep</big>** : 将多维切片递归展平到一层。
- **<big>Intersection</big>** :  创建一个包含所有切片的唯一元素的切片。
-  **<big>ForEach</big>** : 遍历切片的元素并为每个元素调用iteratee函数。
- **<big>Reverse</big>** :  返回元素顺序的slice。
-  **<big>GroupBy</big>** : 迭代切片的元素，每个元素将按条件分组，返回两个切片。
- **<big>Shuffle </big>** : 洗牌切片。
-  **<big>GroupWith</big>** : 创建一个map，key是iteratee遍历slice中的每个元素返回的结果。值是切片元素。
- **<big>Without </big>** : 创建一个不包含所有给定项的slice。
-  **<big>IntSlice<sup>deprecated</sup></big>** : 将接口切片转换为int切片。
- **<big>ToSlicePointer </big>** : 返回一个指向变量参数转换切片的指针。
-  **<big>InterfaceSlice<sup>deprecated</sup></big>** : 将值转换为interface切片。
- **<big>AppendIfAbsent</big>** :  只添加不存在的项。
-  **<big>Intersection</big>** : 返回多个切片的交集。
- **<big>ToMap</big>** : 根据回调函数将切片转换为映射。 

