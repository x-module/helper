
## 特性

-   👏 全面、高效、可复用
-   💪 400+常用 go 工具函数，支持 string、slice、datetime、net、crypt...
-   💅 只依赖 go 标准库
-   🌍 所有导出函数单元测试覆盖率 100%

## 安装

### Note:

```go
go get github.com/x-module/helper  
```

## 用法

以包的结构组织代码的，使用时需要导入相应的包名。例如：如果使用字符串相关函数，需要导入 strutil 包:

```go
import "github.com/x-module/helper/strutil"
```

## 例子

此处以字符串工具函数 Reverse（逆序字符串）为例，需要导入 strutil 包:

```go
package main

import (
    "fmt"
    "github.com/x-module/helper/strutil"
)

func main() {
    s := "hello"
    rs := strutil.Reverse(s)
    fmt.Println(rs) //olleh
}
```



## 模块文档
### 1. catch 捕获程序异常
```go
import "github.com/x-module/helper/cache"
```
#### Function list:
- **<big>Recover</big>** :  捕获程序崩溃异常

### 2. cmd 执行系统命令
```go
import "github.com/x-module/helper/cmd"
```
#### Function list:
- **<big>Run</big>** :  运行shell命令，返回标准输出、标准错误的内容。
- **<big>Call</big>** :  运行shell命令并将结果打印到stdout和stderr。

### 3. convertor 转换相关
```go
import "github.com/x-module/helper/convertor"
```
#### Function list

- **<big>ToBool</big>** : 字符串转布尔类型，使用 strconv.ParseBool。
- **<big>ToBytes</big>** : interface 转字节切片。
- **<big>ToChar</big>** : 字符串转字符切片。
- **<big>ToChannel</big>** : 将切片转为只读 channel。
- **<big>ToFloat</big>** : 将 interface 转成 float64 类型，如果参数无法转换，会返回 0.0 和 error。
- **<big>ToInt</big>** : 将 interface 转成 int64 类型，如果参数无法转换，会返回 0 和 error。
- **<big>ToJson</big>** : 将 interface 转成 json 字符串，如果参数无法转换，会返回""和 error。
- **<big>ToMap</big>** : 将切片转为 map。
- **<big>ToPointer</big>** : 返回传入值的指针。
- **<big>ToString</big>** : 将值转换为字符串，对于数字、字符串、[]byte，将转换为字符串。 对于其他类型（切片、映射、数组、结构）将调用 json.Marshal。
- **<big>StructToMap</big>** : 将 struct 转成 map，只会转换 struct 中可导出的字段。
- **<big>MapToSlice</big>** : map 中 key 和 value 执行函数 iteratee 后，转为切片。
- **<big>EncodeByte</big>** : 将传入的 data 编码成字节切片。
- **<big>DecodeByte</big>** : 解码字节切片到目标对象，目标对象需要传入一个指针实例。
- **<big>TransInterfaceToStruct</big>** : 将 interface 类型转换为 struct 类型。
- **<big>FormatFileSize</big>** : 格式化文件大小，将字节数转换为可读的文件大小。
- **<big>ByteToHex</big>** : byte转16进制字符串。
- **<big>HexToBye</big>** : 16进制字符串转[]byte。
- **<big>StrToInt</big>** :  string转int
- **<big>StrToInt8</big>** :  string转int8
- **<big>StrToInt16</big>** :  string转int16
- **<big>StrToInt32</big>** :  string转int32
- **<big>StrToInt64</big>** :  string转int64
- **<big>dec2Bin</big>** :  将十进制转换为二进制字符串.
- **<big>bin2Dec</big>** :  将二进制字符串转换为十进制.
- **<big>hex2Bin</big>** :  将十六进制字符串转换为二进制字符串.
- **<big>bin2Hex</big>** :  将二进制字符串转换为十六进制字符串.
- **<big>dec2Hex</big>** :  将十进制转换为十六进制.
- **<big>hex2Dec</big>** :  将十六进制转换为十进制.
- **<big>bitSize</big>** :  表示结果的位宽（包括符号位），0 表示最大位宽
- **<big>hex2Byte</big>** :  16进制字符串转字节切片.
- **<big>dec2Oct</big>** :  将十进制转换为八进制.
- **<big>oct2Dec</big>** :  将八进制转换为十进制.
- **<big>Img2Base64</big>** :  将图片字节转换为base64字符串.imgType为图片扩展名.


### 4. cryptor hash、加密相关
```go
import "github.com/x-module/helper/cryptor"
```
#### Function list
-   **<big>HmacMd5</big>** : 返回字符串md5 hmac值。
-   **<big>HmacSha1</big>** : 返回字符串sha1 hmac值。
-   **<big>HmacSha256</big>** : 返回字符串sha256 hmac值。
-   **<big>HmacSha512</big>** : 返回字符串sha256 hmac值。
-   **<big>Md5String</big>** : 返回字符串md5值。
-   **<big>Md5File</big>** : 返回文件md5值。
-   **<big>Sha256</big>** :返回字符串sha256哈希值。
-   **<big>Sha512</big>** : 返回字符串sha512哈希值。
-   **<big>GenerateRsaKey</big>** : 在当前目录下创建rsa私钥文件和公钥文件。
-   **<big>RsaEncrypt</big>** : 用公钥文件ras加密数据。
-   **<big>RsaDecrypt</big>** : 用私钥文件rsa解密数据。

### 5. debug 调试相关
```go
import "github.com/x-module/helper/debug"
```
#### Function list
-   **<big>DumpPrint</big>** : 调试输出
-   **<big>DumpStacks </big>** :打印堆栈信息.
-   **<big>GetCallFile</big>** : 获取调用方法的文件路径.
-   **<big>GetCallLine</big>** : 获取调用方法的行号.

### 6. fileutil 包含文件基本操作
```go
import "github.com/x-module/helper/fileutil "
```
#### Function list
-   **<big>ClearFile</big>** : 清空文件内容。
-   **<big>CreateFile</big>** : 创建文件，创建成功返回 true, 否则返回 false。
-   **<big>CreateDir</big>** : 创建嵌套目录，例如/a/, /a/b/。
-   **<big>CopyFile</big>** :拷贝文件，会覆盖原有的文件。
-   **<big>FileMode</big>** : 获取文件 mode 信息。
-   **<big>MiMeType</big>** : 获取文件 mime 类型, 参数的类型必须是 string 或者\*os.File。
-   **<big>IsExist</big>** : 判断文件或目录是否存在。
-   **<big>IsLink</big>** : 判断文件是否是符号链接。
-   **<big>IsDir</big>** : 判断参数是否是目录。
-   **<big>ListFileNames</big>** : 返回目录下所有文件名。
-   **<big>RemoveFile</big>** : 删除文件。
-   **<big>ReadFileToString</big>** : 读取文件内容并返回字符串。
-   **<big>ReadFileByLine</big>** : 按行读取文件内容，返回字符串切片包含每一行。
-   **<big>Zip</big>** : zip 压缩文件, 参数可以是文件或目录。
-   **<big>UnZip</big>** : zip 解压缩文件并保存在目录中。

### 7. formatter 数据格式化基本操作
```go
import "github.com/x-module/helper/formatter "
```
#### Function list
-   **<big>CommaNumber</big>** : 用逗号每隔 3 位分割数字/字符串，支持前缀添加符号。
-   **<big>FormatSize</big>** : 格式化显示大小。

### 7. function 其他可用方法
```go
import "github.com/x-module/helper/function "
```
#### Function list
-   **<big>methodExists</big>** 检查val结构体中是否存在methodName方法.
-   **<big>getMethod</big>** 获取val结构体的methodName方法.
-   **<big>getFuncNames</big>** 获取变量的所有函数名.
-   **<big>GetFieldValue</big>** 获取(字典/结构体的)字段值;fieldName为字段名,大小写敏感.
-   **<big>GetVariateType</big>** 获取变量类型.
-   **<big>VerifyFunc</big>** 验证是否函数,并且参数个数、类型是否正确.
-   **<big>CallFunc</big>** 动态调用函数.

### 8. json json操作
```go
import "github.com/x-module/helper/json "
```
#### Function list
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

### 9. maputil map可用操作
```go
import "github.com/x-module/helper/maputil "
```
#### Function list
-   **<big>ForEach</big>** : 对 map 中的每对 key 和 value 执行 iteratee 函数。
-   **<big>Filter</big>** : 迭代 map 中的每对 key 和 value，返回 map，其中的 key 和 value 符合 predicate 函数。
-   **<big>Intersect</big>** : 多个 map 的交集操作。
-   **<big>Keys</big>** : 返回 map 中所有 key 组成的切片。
-   **<big>Merge</big>** : 合并多个 map, 相同的 key 会被之后的 key 覆盖。
-   **<big>Minus</big>** : 返回一个 map，其中的 key 存在于 mapA，不存在于 mapB。
-   **<big>Values</big>** : 返回 map 中所有 values 组成的切片
-   **<big>IsDisjoint</big>** : 验证两个 map 是否具有不同的 key。
-   **<big>Union</big>** : 多个 map 的并集操作。

### 10. mathutil 常用的数学操作
```go
import "github.com/x-module/helper/mathutil "
```
#### Function list
-   **<big>Average</big>** : 计算map中的值的平均值。

### 11. netutil 网络操作
```go
import "github.com/x-module/helper/netutil "
```
#### Function list
-   **<big>ConvertMapToQueryString</big>** : 将map转换成http查询字符串。
-   **<big>EncodeUrl</big>** : 编码url query string的值(?a=1&b=[2] -> ?a=1&b=%5B2%5D)。
-   **<big>GetInternalIp</big>** : 获取内部ipv4。
-   **<big>GetIps</big>** : 获取系统ipv4地址列表。
-   **<big>GetMacAddrs</big>** : 获取系统mac地址列。
-   **<big>GetPublicIpInfo</big>** : 获取[公网ip信息](http://ip-api.com/json/).
-   **<big>GetRequestPublicIp</big>** : 获取http请求ip。
-   **<big>IsPublicIP</big>** : 判断ip是否是公共ip。
-   **<big>IsInternalIP</big>** : 判断ip是否是局域网ip。
-   **<big>StructToUrlValues</big>** : 将结构体转为url values, 仅转化结构体导出字段并且包含`json` tag。
-   **<big>RequestUtils</big>** : 网络请求相关库。

### 12. random 随机相关操作
```go
import "github.com/x-module/helper/random "
```
#### Function list
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

### 12. reflect 反射相关操作
```go
import "github.com/x-module/helper/reflect "
```
#### Function list
-   **<big>FindTag</big>** : 查找struct 的tag信息。

### 13. serialize 序列化相关操作
```go
import "github.com/x-module/helper/serialize "
```
#### Function list
-   **<big>Encode</big>** : 二进制到struct相互转换
-   **<big>Decode</big>** : 反序列化


### 14. sliceutil 数组相关操作
```go
import "github.com/x-module/helper/sliceutil "
```
#### Function list
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

### 15. strutil 字符串相关操作
```go
import "github.com/x-module/helper/strutil "
```
#### Function list
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

### 16. validator 数据验证相关操作
```go
import "github.com/x-module/helper/validator "
```
#### Function list
### 验证器包，包含常用字符串格式验证函数。

```go
import "github.com/x-module/utils/validator"
```

#### 函数列表:
-   **<big> CheckIsMobile</big>** :  手机号码检测
-   **<big> IsIdCard</big>** :  判断是否是18或15位身份证
-   **<big> IsURL </big>** : 是否为URL地址
-   **<big> IsQQ </big>** : 验证是否为QQ号
-   **<big> IsWeChat</big>** :  验证是否为微信号
-   **<big> IsWeibo</big>** :  验证是否为微博ID
-   **<big> IsPassword</big>** :  验证密码是否合法
-   **<big> IsBankCardNo</big>** :  验证是否为大陆银行卡号
-   **<big> IsTime</big>** :  验证是否为时间格式（HH:mm:ss）
-   **<big> IsDate</big>** :  验证是否为日期格式（yyyy-MM-dd）
-   **<big> IsDateTime</big>** :  验证是否为日期时间格式（yyyy-MM-dd HH:mm:ss）
-   **<big> IsIDCard</big>** :  验证身份证号(18或15位)
-   **<big> IsIDCard18</big>** :  验证18位身份证号
-   **<big> IsIDCard15 </big>** : 验证15位身份证号
-   **<big> IsIPv4</big>** :  是否为ipv4地址
-   **<big> IsIPv6 </big>** : 是否为ipv6地址
-   **<big> IsAllChinese </big>** : 验证给定的字符串全部为中文
-   **<big> IsContainChinese</big>** :  验证给定的字符串包含中文
-   **<big> IsChineseName</big>** :  验证是否为中文名
-   **<big> IsEnglishName</big>** :  验证是否为英文名
-   **<big> IsNumber</big>** :  验证是否全部为数字
-   **<big> IsPostalCode</big>** :  验证是否为邮编号码
-   **<big> IsTelephone</big>** :  验证是否为座机号码
-   **<big> IsAlpha </big>** : 检查字符串是否只包含字母(a-zA-Z)。
-   **<big> IsAllUpper</big>** :  检查字符串是否全是大写字母A-Z。
-   **<big> IsAllLower </big>** : 检查字符串是否全是小写字母a-z。
-   **<big> ContainUpper</big>** :  检查字符串是否至少包含一个大写字母A-Z。
-   **<big> ContainLower</big>** :  检查字符串是否至少包含一个小写字母a-z。
-   **<big> ContainLetter</big>** :  检查字符串是否至少包含一个字母。
-   **<big> IsJSON</big>** :  检查字符串是否为有效的JSON。
-   **<big> IsNumberStr</big>** :  检查字符串是否可以转换为数字。
-   **<big> IsFloatStr</big>** :  检查字符串是否可以转换为浮点数。
-   **<big> IsIntStr</big>** :  检查字符串是否可以转换为整数。
-   **<big> IsIp</big>** :  检查字符串是否是有效ip地址
-   **<big> IsPort </big>** : 检查字符串是否是有效端口
-   **<big> IsUrl</big>** :  检查字符串是否是url
-   **<big> IsDns</big>** :  检查字符串是否为DNS
-   **<big> IsEmail</big>** :  检查字符串是否为电子邮件地址。
-   **<big> IsChineseMobile</big>** :  检查字符串是否为中国手机号码。
-   **<big> IsChineseIdNum</big>** :  检查字符串是否为中国身份证。
-   **<big> ContainChinese</big>** :  检查字符串是否中文。
-   **<big> IsChinesePhone</big>** :  检查字符串是否为中文电话号码。
-   **<big> IsCreditCard</big>** :  检查字符串是否为信用卡。
-   **<big> IsBase64</big>** :  检查字符串是否为base64字符串。
-   **<big> IsRegexMatch </big>** : 检查字符串是否与regexp匹配。
-   **<big> IsStrongPassword</big>** :  检查字符串是否为强密码，如果len(password)小于length参数，返回false 强密码:α(低+上)+数字+特殊字符 (!@#$%^&*()?><).
-   **<big> IsWeakPassword</big>** :  检查字符串是否为弱密码
-   **<big> IsGBK</big>** :  判断是否是GBK编码
-   **<big>IsPointer</big>** 检查变量是否指针类型;
-   **<big>isBool</big>** 是否布尔值.
-   **<big>isMap</big>** 检查变量是否字典.
-   **<big>isString</big>** 变量是否字符串.
-   **<big>isByte</big>** 变量是否字节切片.
-   **<big>isBinary</big>** 字符串是否二进制.
-   **<big>isInt</big>** 变量是否整型数值.
-   **<big>isNumeric</big>** 变量是否数值(不包含复数).

## 如何贡献代码
非常感激任何的代码提交以使 lancet 的功能越来越强大。创建 pull request 时请遵守以下规则。

1. Fork lancet 仓库。
2. 创建自己的特性分支。
3. 提交变更。
4. Push 分支。
5. 创建新的 pull request。
