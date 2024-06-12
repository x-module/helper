/**
 * Created by Goland
 * @file   json.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/6/12 11:41
 * @desc   json.go
 */

package json

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/json-iterator/go/extra"
)

// RegisterFuzzyDecoders 注册模糊解码器
func RegisterFuzzyDecoders() {
	extra.RegisterFuzzyDecoders()
}

// FormatString 格式化输出,缩进4个空格
func FormatString(config ...jsoniter.Config) jsoniter.Config {
	conf := getConfig(config...)
	conf.IndentionStep = 4
	return conf
}

// SortMapKeys 按照key排序
func SortMapKeys(config ...jsoniter.Config) jsoniter.Config {
	conf := getConfig(config...)
	conf.SortMapKeys = true
	return conf
}

// DisallowUnknownFields 禁止未知字段
func DisallowUnknownFields(config ...jsoniter.Config) jsoniter.Config {
	conf := getConfig(config...)
	conf.DisallowUnknownFields = true
	return conf
}

// OnlyTaggedField 只解析有tag的字段
func OnlyTaggedField(config ...jsoniter.Config) jsoniter.Config {
	conf := getConfig(config...)
	conf.OnlyTaggedField = true
	return conf
}

// TagKey 设置tag key
func TagKey(key string, config ...jsoniter.Config) jsoniter.Config {
	conf := getConfig(config...)
	conf.TagKey = key
	return conf
}

// CaseSensitive 区分大小写
func CaseSensitive(config ...jsoniter.Config) jsoniter.Config {
	conf := getConfig(config...)
	conf.CaseSensitive = true
	return conf
}

// getConfig 获取配置

func getConfig(config ...jsoniter.Config) jsoniter.Config {
	if len(config) == 0 {
		return jsoniter.Config{}
	}
	return config[0]
}

// MarshalToString 序列化为字符串
func MarshalToString(data any) {
	c := FormatString().Froze()
	if s, err := c.MarshalToString(data); err == nil {
		fmt.Println(s)
	} else {
		fmt.Printf("%+v\n", data)
	}
}

// Marshal 序列化为字节

// todo 支持传入多个配置

func Marshal(data any, config ...jsoniter.Config) ([]byte, error) {
	conf := getConfig(config...).Froze()
	if s, err := conf.Marshal(data); err == nil {
		return s, nil
	} else {
		return nil, err
	}
}

// Unmarshal 反序列化
func Unmarshal(data []byte, v any, config ...jsoniter.Config) error {
	conf := getConfig(config...).Froze()
	return conf.Unmarshal(data, v)
}
