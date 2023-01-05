package steam

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

func getEndPoint(interface_ string, version int) string {
	method := getCallerName()
	ms := strings.Split(method, ".")
	return fmt.Sprintf("%s/%s/v%d", interface_, ms[len(ms)-1], version)
}

func getCallerName() string {
	pc, _, _, _ := runtime.Caller(2)
	return runtime.FuncForPC(pc).Name()
}

// 读取 element 中 name 字段的值
// ATTENTION：如果获取的字段值为空时不能仅仅判断是否为 Nil，此时 return 的是值为 Nil，类型为原字段类型的 interface 变量。
func getValueByFiledName(element interface{}, name string) interface{} {
	value := reflect.ValueOf(element)
	var keys []string
	if strings.ContainsAny(name, ".") {
		keys = strings.Split(name, ".")
		firstKey := keys[0]
		keys = append(keys[1:])
		name = strings.Join(keys, ".")
		return getValueByFiledName(getValueByFiledName(element, firstKey), name)
	} else {
		switch value.Kind() {
		case reflect.Ptr:
			return value.Elem().FieldByName(name).Interface()
		case reflect.Struct:
			return value.FieldByName(name).Interface()
		case reflect.Map:
			return value.MapIndex(reflect.ValueOf(name)).Interface()
		case reflect.Slice, reflect.Array:
			if value.IsZero() {
				return nil
			} else {
				return extractArrayField(name, value.Interface())
			}

		default:
			panic(fmt.Sprintf("not support get value from type[%s]", value.Kind()))
		}
	}
}

// 该方法可以将数组中指定字段的值提取出来，组成新的切片
func extractArrayField(fieldName string, array interface{}) []interface{} {
	slice, len := convertToSlice(array)

	fieldValues := make([]interface{}, len)

	for i, item := range slice {
		fieldValues[i] = getValueByFiledName(item, fieldName)
	}
	return fieldValues
}

// 将 array 转为 []interface{}，并返回 array 的长度
// 将 slice 和 array 以外的类型变量传入会引起 panic
func convertToSlice(array interface{}) ([]interface{}, int) {
	v := reflect.ValueOf(array)
	switch v.Kind() {
	case reflect.Slice, reflect.Array:
		result := make([]interface{}, v.Len())
		for i := 0; i < v.Len(); i++ {
			result[i] = v.Index(i).Interface()
		}
		return result, v.Len()
	case reflect.Ptr:
		return convertToSlice(v.Elem().Interface())
	default:
		panic(fmt.Sprintf("not support convert type[%s] to slice", v.Kind()))
	}
}
