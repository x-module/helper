/**
 * Created by Goland
 * @file   random.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/6/12 14:37
 * @desc   random.go
 */

package random

import (
	crand "crypto/rand"
	"fmt"
	"github.com/google/uuid"
	"io"
	"math/rand"
	"strings"
	"time"
)

const (
	NUMERAL      = "0123456789"
	LowerLetters = "abcdefghijklmnopqrstuvwxyz"
	UpperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	LETTERS      = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// RandInt 在min和max之间生成随机整数，可能是min，而不是max。
func RandInt(min, max int) int {
	if min == max {
		return min
	}
	if max < min {
		min, max = max, min
	}
	return rand.Intn(max-min) + min
}

// RandBytes 生成随机字节片
func RandBytes(length int) []byte {
	if length < 1 {
		return []byte{}
	}
	b := make([]byte, length)
	if _, err := io.ReadFull(crand.Reader, b); err != nil {
		return nil
	}
	return b
}

// RandString 生成指定长度的随机字符串。
func RandString(length int) string {
	return random(LETTERS, length)
}

// RandUpper 生成一个随机大写字符串。
func RandUpper(length int) string {
	return random(UpperLetters, length)
}

// RandLower 生成一个随机小写字符串。
func RandLower(length int) string {
	return random(LowerLetters, length)
}

// RandNumeral 生成指定长度的随机数字字符串。
func RandNumeral(length int) string {
	return random(NUMERAL, length)
}

// RandNumeralOrLetter 生成一个随机数字或字母字符串。
func RandNumeralOrLetter(length int) string {
	return random(NUMERAL+LETTERS, length)
}

// random 生成一个基于给定字符串范围的随机字符串。
func random(s string, length int) string {
	b := make([]byte, length)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range b {
		b[i] = s[r.Int63()%int64(len(s))]
	}
	return string(b)
}

// UUIdV4 根据RFC 4122生成版本4的随机UUID。
func UUIdV4() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(crand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	uuid[8] = uuid[8]&^0xc0 | 0x80
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}

// GetUUIdByTime 根据时间生成UUID true去除“-”，false不去除
func GetUUIdByTime(flag bool) (string, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	if flag {
		return strings.Replace(id.String(), "-", "", -1), nil
	}
	return id.String(), err
}

// IdUUIdByRand V4 基于随机数 true去除“-”，false不去除
func IdUUIdByRand(flag bool) string {
	u4 := uuid.New()
	if flag {
		return strings.Replace(u4.String(), "-", "", -1)
	}
	return u4.String()
}
