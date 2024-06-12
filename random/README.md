
### 随机数生成器包，可以生成随机[]bytes, int, string。

```go
import "github.com/x-module/utils/random"
```

#### 函数列表:

-   **<big>RandBytes</big>** : 生成随机字节切片。
-   **<big>RandInt</big>** : 生成随机int, 范围[min, max)。
-   **<big>RandString</big>** : 生成给定长度的随机字符串，只包含字母(a-zA-Z)。
-   **<big>RandUpper</big>** : 生成给定长度的随机大写字母字符串(A-Z)。
-   **<big>RandLower</big>** : 生成给定长度的随机小写字母字符串(a-z)。
-   **<big>RandNumeral</big>** : 生成给定长度的随机数字字符串(0-9)。
-   **<big>RandNumeralOrLetter</big>** : 生成给定长度的随机字符串（数字+字母)。
-   **<big>UUIdV4</big>** : 生成UUID v4字符串。
-   **<big>GetUUIdByTime</big>** : 生成基于时间的UUID字符串。
-   **<big>IdUUIdByRand</big>** : 生成随机的UUID字符串。
