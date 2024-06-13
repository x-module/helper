
### 操作切片的方法集合。

```go
import "github.com/x-module/utils/slice"
```

#### 函数列表:
-   **<big>AppendIfAbsent</big>** : 当前切片中不包含值时，将该值追加到切片中。
-   **<big>Contain</big>** : 判断slice是否包含value。
-   **<big>ContainSubSlice</big>** : 判断slice是否包含subslice。
-   **<big>Chunk</big>** : 按照size参数均分slice。
-   **<big>Compact</big>** : 去除slice中的假值（false values are false, nil, 0, ""）。
-   **<big>Concat</big>** : 合并多个slices到一个slice中。
-   **<big>Count</big>** : 返回切片中指定元素的个数。
-   **<big>CountBy</big>** : 遍历切片，对每个元素执行函数predicate. 返回符合函数返回值为true的元素的个数。
-   **<big>Difference</big>** : 创建一个切片，其元素不包含在另一个给定切片中。
-   **<big>DifferenceBy</big>** : 将两个slice中的每个元素调用iteratee函数，并比较它们的返回值，如果不相等返回在slice中对应的值。
-   **<big>DifferenceWith</big>** : 接受比较器函数，该比较器被调用以将切片的元素与值进行比较。 结果值的顺序和引用由第一个切片确定。
-   **<big>DeleteAt</big>** : 删除切片中指定开始索引到结束索引的元素。
-   **<big>Drop</big>** : 创建一个切片，当n > 0时从开头删除n个元素，或者当n < 0时从结尾删除n个元素。
-   **<big>Equal</big>** : 检查两个切片是否相等，相等条件：切片长度相同，元素顺序和值都相同。
-   **<big>EqualWith</big>** : 检查两个切片是否相等，相等条件：对两个切片的元素调用比较函数comparator，返回true。
-   **<big>Every</big>** : 如果切片中的所有值都通过谓词函数，则返回true。
-   **<big>Filter</big>** : 返回切片中通过predicate函数真值测试的所有元素。
-   **<big>Find</big>** : 遍历切片的元素，返回第一个通过predicate函数真值测试的元素。
-   **<big>FindLast</big>** : 从头到尾遍历slice的元素，返回最后一个通过predicate函数真值测试的元素。
-   **<big>Flatten</big>** : 将多维切片展平一层。
-   **<big>FlattenDeep</big>** : 将多维切片递归展平到一层。
-   **<big>ForEach</big>** : 遍历切片的元素并为每个元素调用iteratee函数。
-   **<big>GroupBy</big>** : 迭代切片的元素，每个元素将按条件分组，返回两个切片。
-   **<big>GroupWith</big>** : 创建一个map，key是iteratee遍历slice中的每个元素返回的结果。值是切片元素。
-   **<big>IntSlice<sup>deprecated</sup></big>** : 将接口切片转换为int切片。
-   **<big>InterfaceSlice<sup>deprecated</sup></big>** : 将值转换为interface切片。
-   **<big>Intersection</big>** : 返回多个切片的交集。
-   **<big>InsertAt</big>** : 将元素插入到索引处的切片中。
-   **<big>IndexOf</big>** : 返回在切片中找到值的第一个匹配项的索引，如果找不到值，则返回-1。
-   **<big>LastIndexOf</big>** : 返回在切片中找到最后一个值的索引，如果找不到该值，则返回-1。
-   **<big>Map</big>** : 对slice中的每个元素执行map函数以创建一个新切片。
-   **<big>Merge</big>** : 合并多个切片（不会消除重复元素)。
-   **<big>Reverse</big>** : 反转切片中的元素顺序。
-   **<big>Reduce</big>** : 将切片中的元素依次运行iteratee函数，返回运行结果。
-   **<big>Replace</big>** : 返回切片的副本，其中前n个不重叠的old替换为new。
-   **<big>ReplaceAll</big>** : 返回切片的副本，将其中old全部替换为new。
-   **<big>Repeat</big>** : 创建一个切片，包含n个传入的item。
-   **<big>Shuffle</big>** : 随机打乱切片中的元素顺序。
-   **<big>Sort</big>** : 对任何有序类型（数字或字符串）的切片进行排序，使用快速排序算法。
-   **<big>SortBy</big>** : 按照less函数确定的升序规则对切片进行排序。排序不保证稳定性。
-   **<big>SortByField<sup>deprecated</sup></big>** : 按字段对结构切片进行排序。slice元素应为struct，字段类型应为int、uint、string或bool。
-   **<big>Some</big>** : 如果列表中的任何值通过谓词函数，则返回true。
-   **<big>StringSlice<sup>deprecated</sup></big>** : 将接口切片转换为字符串切片。
-   **<big>SymmetricDifference</big>** : 返回一个切片，其中的元素存在于参数切片中，但不同时存储在于参数切片中（交集取反）。
-   **<big>ToSlice</big>** : 将可变参数转为切片。
-   **<big>ToSlicePointer</big>** : 将可变参数转为指针切片。
-   **<big>Unique</big>** : 删除切片中的重复元素。
-   **<big>UniqueBy</big>** : 对切片的每个元素调用iteratee函数，然后删除重复元素。
-   **<big>UnionBy</big>** : 对切片的每个元素调用函数后，合并多个切片。
-   **<big>UpdateAt</big>** : 更新索引处的切片元素。
-   **<big>Without</big>** : 创建一个不包括所有给定值的切片。
-   **<big>KeyBy</big>** :将切片每个元素调用函数后转为map。
