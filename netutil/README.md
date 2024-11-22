
### 13. netutil 网络包支持获取 ip 地址，发送 http 请求。

```go
import "github.com/x-module/helper/netutil"
```

#### 函数列表:

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
