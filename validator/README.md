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
