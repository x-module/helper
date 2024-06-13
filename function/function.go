/**
 * Created by Goland
 * @file   function.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/6/13 12:00
 * @desc   function.go
 */

package function

import (
	"errors"
	"fmt"
	"reflect"
)

// methodExists 检查val结构体中是否存在methodName方法.
func methodExists(val any, methodName string) (bool, error) {
	if methodName == "" {
		return false, errors.New("[methodExists]`methodName can not be empty.")
	}
	r := reflect.ValueOf(val)
	if r.Type().Kind() != reflect.Ptr {
		r = reflect.New(reflect.TypeOf(val))
	}
	method := r.MethodByName(methodName)
	if !method.IsValid() {
		return false, fmt.Errorf("[methodExists] Method `%s` not exists in interface `%s`", methodName, r.Type())
	}
	return true, nil
}

// getMethod 获取val结构体的methodName方法.
// 注意:返回的方法中的第一个参数是接收者. 所以,调用返回的方法时,必须将接收者作为第一个参数传递.
func getMethod(val any, methodName string) any {
	if val == nil || methodName == "" {
		return nil
	}
	r := reflect.ValueOf(val)
	if r.Type().Kind() != reflect.Ptr {
		r = reflect.New(reflect.TypeOf(val))
	}
	method := r.MethodByName(methodName)
	if !method.IsValid() {
		return nil
	}
	return method.Interface()
}

// getFuncNames 获取变量的所有函数名.
func getFuncNames(val any) (res []string) {
	if val == nil {
		return
	}
	r := reflect.ValueOf(val)
	if r.Type().Kind() != reflect.Ptr {
		r = reflect.New(reflect.TypeOf(val))
	}
	typ := r.Type()
	for i := 0; i < r.NumMethod(); i++ {
		res = append(res, typ.Method(i).Name)
	}
	return
}

// GetFieldValue 获取(字典/结构体的)字段值;fieldName为字段名,大小写敏感.
func GetFieldValue(arr any, fieldName string) (res any, err error) {
	val := reflect.ValueOf(arr)
	switch val.Kind() {
	case reflect.Map:
		for _, subKey := range val.MapKeys() {
			if fmt.Sprintf("%s", subKey) == fieldName {
				res = val.MapIndex(subKey).Interface()
				break
			}
		}
	case reflect.Struct:
		field := val.FieldByName(fieldName)
		if !field.IsValid() || !field.CanInterface() {
			break
		}
		res = field.Interface()
	default:
		err = errors.New("[GetFieldValue]`arr type must be map|struct; but : " + val.Kind().String())
	}
	return
}

// GetVariateType 获取变量类型.
func GetVariateType(v any) string {
	return fmt.Sprintf("%T", v)
}

// VerifyFunc 验证是否函数,并且参数个数、类型是否正确.
// 返回有效的函数、有效的参数.
func VerifyFunc(f any, args ...any) (vf reflect.Value, vargs []reflect.Value, err error) {
	vf = reflect.ValueOf(f)
	if vf.Kind() != reflect.Func {
		return reflect.ValueOf(nil), nil, fmt.Errorf("[VerifyFunc] %v is not the function", f)
	}

	tf := vf.Type()
	num := len(args)
	if tf.NumIn() != num {
		return reflect.ValueOf(nil), nil, fmt.Errorf("[VerifyFunc] %d number of the argument is incorrect", num)
	}

	vargs = make([]reflect.Value, num)
	for i := 0; i < num; i++ {
		typ := tf.In(i).Kind()
		if (typ != reflect.Interface) && (typ != reflect.TypeOf(args[i]).Kind()) {
			return reflect.ValueOf(nil), nil, fmt.Errorf("[VerifyFunc] %d-td argument`s type is incorrect", i+1)
		}
		vargs[i] = reflect.ValueOf(args[i])
	}

	return vf, vargs, nil
}

// CallFunc 动态调用函数.
func CallFunc(f any, args ...any) ([]any, error) {
	vf, vargs, err := VerifyFunc(f, args...)
	if err != nil {
		return nil, err
	}

	ret := vf.Call(vargs)
	num := len(ret)
	results := make([]any, num)
	for i := 0; i < num; i++ {
		results[i] = ret[i].Interface()
	}

	return results, nil
}
