## 数据转换处理的相关函数

```go
import "github.com/x-module/utils/strutil"
```

#### 函数列表:

-   **<big>ToBool</big>** : 字符串转布尔类型，使用 strconv.ParseBool。
-   **<big>ToBytes</big>** : interface 转字节切片。
-   **<big>ToChar</big>** : 字符串转字符切片。
-   **<big>ToChannel</big>** : 将切片转为只读 channel。
-   **<big>ToFloat</big>** : 将 interface 转成 float64 类型，如果参数无法转换，会返回 0.0 和 error。
-   **<big>ToInt</big>** : 将 interface 转成 int64 类型，如果参数无法转换，会返回 0 和 error。
-   **<big>ToJson</big>** : 将 interface 转成 json 字符串，如果参数无法转换，会返回""和 error。
-   **<big>ToMap</big>** : 将切片转为 map。
-   **<big>ToPointer</big>** : 返回传入值的指针。
-   **<big>ToString</big>** : 将值转换为字符串，对于数字、字符串、[]byte，将转换为字符串。 对于其他类型（切片、映射、数组、结构）将调用 json.Marshal。
-   **<big>StructToMap</big>** : 将 struct 转成 map，只会转换 struct 中可导出的字段。
-   **<big>MapToSlice</big>** : map 中 key 和 value 执行函数 iteratee 后，转为切片。
-   **<big>EncodeByte</big>** : 将传入的 data 编码成字节切片。
-   **<big>DecodeByte</big>** : 解码字节切片到目标对象，目标对象需要传入一个指针实例。
-   **<big>TransInterfaceToStruct</big>** : 将 interface 类型转换为 struct 类型。
-   **<big>FormatFileSize</big>** : 格式化文件大小，将字节数转换为可读的文件大小。
-   **<big> ByteToHex </big>** : byte转16进制字符串。
-   **<big> HexToBye  </big>** : 16进制字符串转[]byte。
- 