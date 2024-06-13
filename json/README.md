##  json操作
```go
import "github.com/x-module/helper/json"
```

#### 函数列表:
-   **<big>RegisterFuzzyDecoders</big>** 注册模糊解码器
-   **<big>FormatString</big>** 格式化输出,缩进4个空格
-   **<big>SortMapKeys</big>** 按照key排序
-   **<big>DisallowUnknownFields</big>** 禁止未知字段
-   **<big>OnlyTaggedField</big>** 只解析有tag的字段
-   **<big>TagKey</big>** 设置tag key
-   **<big>CaseSensitive</big>** 区分大小写
-   **<big>getConfig</big>** 获取配置
-   **<big>MarshalToString</big>** 序列化为字符串
-   **<big>Marshal</big>** 序列化为字节
-   **<big>Unmarshal</big>** 反序列化