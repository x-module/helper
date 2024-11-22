## 加密包支持数据加密和解密，获取 md5，hash 值。支持 base64, md5, hmac, aes, des, rsa。

```go
import "github.com/x-module/helper/cryptor"
```

#### 函数列表:

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

