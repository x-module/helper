##  包含文件基本操作。

```go
import "github.com/x-module/utils/fileutil"
```

#### 函数列表：

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