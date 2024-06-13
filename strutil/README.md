## 字符串处理的相关函数

```go
import "github.com/x-module/utils/strutil"
```

#### 函数列表:

- **<big>After</big>** : 返回源字符串中指定字符串首次出现时的位置之后的子字符串。
- **<big>AfterLast</big>** : 返回源字符串中指定字符串最后一次出现时的位置之后的子字符串。
- **<big>Before</big>** : 返回源字符串中指定字符串第一次出现时的位置之前的子字符串。
- **<big>BeforeLast</big>** : 返回源字符串中指定字符串最后一次出现时的位置之前的子字符串。
- **<big>CamelCase</big>** : 将字符串转换为 CamelCase 驼峰式字符串, 非字母和数字会被忽略。
- **<big>Capitalize</big>** : 将字符串的第一个字符转换为大写。
- **<big>IsString</big>** : 判断传入参数的数据类型是否为字符串。
- **<big>KebabCase</big>** : 将字符串转换为 kebab-case 形式字符串, 非字母和数字会被忽略。
- **<big>UpperKebabCase</big>** : 将字符串转换为大写 KEBAB-CASE 形式字符串, 非字母和数字会被忽略。
- **<big>LowerFirst</big>** : 将字符串的第一个字符转换为小写形式。
- **<big>UpperFirst</big>** : 将字符串的第一个字符转换为大写形式。
- **<big>AppendEnd</big>** : 如果字符串短于限制大小，则在右侧用给定字符填充字符串。 如果填充字符超出大小，它们将被截断。
- **<big>AppendStart</big>** : 如果字符串短于限制大小，则在左侧用给定字符填充字符串。 如果填充字符超出大小，它们将被截断。
- **<big>Reverse</big>** : 返回字符顺序与给定字符串相反的字符串。
- **<big>SnakeCase</big>** : 将字符串转换为 snake_case 形式, 非字母和数字会被忽略。
- **<big>UpperSnakeCase</big>** : 将字符串转换为大写 SNAKE_CASE 形式, 非字母和数字会被忽略。
- **<big>Substring</big>** : 根据指定的位置和长度截取子字符串。
- **<big>SimilarText</big>** : 计算两个字符串的相似度。
