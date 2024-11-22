##  函数包控制函数执行流程，包含部分函数式编程。

```go
import "github.com/x-module/helper/function"
```

#### 函数列表:

-   **<big>methodExists</big>** 检查val结构体中是否存在methodName方法.
-   **<big>getMethod</big>** 获取val结构体的methodName方法.
-   **<big>getFuncNames</big>** 获取变量的所有函数名.
-   **<big>GetFieldValue</big>** 获取(字典/结构体的)字段值;fieldName为字段名,大小写敏感.
-   **<big>GetVariateType</big>** 获取变量类型.
-   **<big>VerifyFunc</big>** 验证是否函数,并且参数个数、类型是否正确.
-   **<big>CallFunc</big>** 动态调用函数.