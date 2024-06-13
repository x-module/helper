/**
 * Created by Goland
 * @file   reflect.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/6/13 12:16
 * @desc   reflect.go
 */

package reflect

import (
	"reflect"
	"strings"
)

// FindTag 查找struct 的tag信息
func FindTag(obj any, field, tag string) string {
	dataStructType := reflect.Indirect(reflect.ValueOf(obj)).Type()
	for i := 0; i < dataStructType.NumField(); i++ {
		fd := dataStructType.Field(i)
		if fd.Name == field {
			bb := fd.Tag
			sqlTag := bb.Get(tag)

			if sqlTag == "-" || bb == "-" {
				return ""
			}

			sqlTags := strings.Split(sqlTag, ",")
			sqlFieldName := fd.Name // default
			if len(sqlTags[0]) > 0 {
				sqlFieldName = sqlTags[0]
			}
			return sqlFieldName
		}
	}
	return ""
}
