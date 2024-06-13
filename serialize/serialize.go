/**
 * Created by Goland
 * @file   net.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/6/12 14:07
 * @desc   serialize.go
 */

package serializ

import (
	"bytes"
	"encoding/gob"
)

// Encode 二进制到struct相互转换
func Encode(data any) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Decode 反序列化
func Decode(data []byte, to any) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	return dec.Decode(to)
}
