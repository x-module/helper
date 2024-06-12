/**
 * Created by Goland
 * @file   cryptor.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/6/12 13:18
 * @desc   cryptor.go
 */

package cryptor

import (
	"bufio"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"
)

// Base64StdEncode base64编码
func Base64StdEncode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

// Base64StdDecode base64解码
func Base64StdDecode(s string) string {
	b, _ := base64.StdEncoding.DecodeString(s)
	return string(b)
}

// Md5String md5加密
func Md5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// Md5File 文件md5加密
func Md5File(filename string) (string, error) {
	if fileInfo, err := os.Stat(filename); err != nil {
		return "", err
	} else if fileInfo.IsDir() {
		return "", nil
	}

	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()

	chunkSize := 65536
	for buf, reader := make([]byte, chunkSize), bufio.NewReader(file); ; {
		n, err := reader.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}
		hash.Write(buf[:n])
	}

	checksum := fmt.Sprintf("%x", hash.Sum(nil))
	return checksum, nil
}

// HmacMd5 返回使用md5的字符串的hmac哈希。
func HmacMd5(data, key string) string {
	h := hmac.New(md5.New, []byte(key))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum([]byte("")))
}

// HmacSha1 返回字符串使用sha1的hmac哈希值。
func HmacSha1(data, key string) string {
	h := hmac.New(sha1.New, []byte(key))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum([]byte("")))
}

// HmacSha256 返回字符串使用sha256的hmac哈希值。
func HmacSha256(data, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum([]byte("")))
}

// HmacSha512 返回字符串使用sha512的hmac哈希值。
func HmacSha512(data, key string) string {
	h := hmac.New(sha512.New, []byte(key))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum([]byte("")))
}

// Sha1 返回字符串的sha1值(SHA-1哈希算法)。
func Sha1(s string) string {
	o := sha1.New()
	o.Write([]byte(s))
	return strings.ToUpper(hex.EncodeToString(o.Sum(nil)))
}

// Sha256 返回字符串的sha256值(SHA256哈希算法)。
func Sha256(data string) string {
	sha256 := sha256.New()
	sha256.Write([]byte(data))
	return hex.EncodeToString(sha256.Sum([]byte("")))
}

// SHA-512 返回字符串的sha512值(sha512哈希算法)。
func Sha512(data string) string {
	sha512 := sha512.New()
	sha512.Write([]byte(data))
	return hex.EncodeToString(sha512.Sum([]byte("")))
}
