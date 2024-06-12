## MAP处理的相关函数

```go
import "github.com/x-module/utils/maputil"
```

#### 函数列表:
-   **<big>ForEach</big>** : 对 map 中的每对 key 和 value 执行 iteratee 函数。
-   **<big>Filter</big>** : 迭代 map 中的每对 key 和 value，返回 map，其中的 key 和 value 符合 predicate 函数。
-   **<big>Intersect</big>** : 多个 map 的交集操作。
-   **<big>Keys</big>** : 返回 map 中所有 key 组成的切片。
-   **<big>Merge</big>** : 合并多个 map, 相同的 key 会被之后的 key 覆盖。
-   **<big>Minus</big>** : 返回一个 map，其中的 key 存在于 mapA，不存在于 mapB。
-   **<big>Values</big>** : 返回 map 中所有 values 组成的切片
-   **<big>IsDisjoint</big>** : 验证两个 map 是否具有不同的 key。
-   **<big>Union</big>** : 多个 map 的并集操作。