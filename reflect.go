package utils

import (
	"reflect"
	"strings"
)

// IsSameStruct 判断是否同一个结构体
func IsSameStruct(v1, v2 interface{}) bool {
	var (
		f1 reflect.Type
		f2 reflect.Type
		ok bool
	)

	if f1, ok = v1.(reflect.Type); !ok {
		f1 = reflect.TypeOf(v1)
	}

	if f2, ok = v2.(reflect.Type); !ok {
		f2 = reflect.TypeOf(v2)
	}

	return f1.PkgPath() == f2.PkgPath() && f1.Name() == f2.Name()
}

// ConvertToTypes 把变量转换成反射类型
func ConvertToTypes(args ...interface{}) []reflect.Type {
	types := make([]reflect.Type, 0)
	for _, arg := range args {
		types = append(types, reflect.TypeOf(arg))
	}
	return types
}

// IsInstanceIn InstanceIn 判断变量是否是某些类型
func IsInstanceIn(v interface{}, types ...reflect.Type) bool {
	for _, e := range types {
		if IsSameStruct(e, v) {
			return true
		}
	}
	return false
}

// EachStructField 遍历结构体的字段
func EachStructField(s interface{}, handler func(reflect.StructField, reflect.Value)) {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	for i := 0; i < t.NumField(); i++ {
		handler(t.Field(i), v.Field(i))
	}
}

// GetTypeKey 获取类型唯一字符串
func GetTypeKey(p reflect.Type) (key string) {
	if p.Kind() == reflect.Ptr {
		p = p.Elem()
		key = "*"
	}

	pkgPath := p.PkgPath()

	if pkgPath != "" {
		key += pkgPath + "."
	}

	return key + p.Name()
}
// ContainsKind checks if a specified kind in the slice of kinds.
func ContainsKind(kinds []reflect.Kind, kind reflect.Kind) bool {
	for i := 0; i < len(kinds); i++ {
		if kind == kinds[i] {
			return true
		}
	}

	return false
}

// IsNil checks if a specified object is nil or not, without Failing.
func IsNil(object interface{}) bool {
	if object == nil {
		return true
	}

	value := reflect.ValueOf(object)
	kind := value.Kind()
	isNilableKind := ContainsKind(
		[]reflect.Kind{
			reflect.Chan, reflect.Func,
			reflect.Interface, reflect.Map,
			reflect.Ptr, reflect.Slice},
		kind)

	if isNilableKind && value.IsNil() {
		return true
	}

	return false
}

// WithoutNil 尽量不要 nil
func WithoutNil(args ...interface{}) interface{} {
	for _, arg := range args {
		switch argValue := arg.(type) {
		case func() interface{}:
			arg = argValue()
		}
		if !IsNil(arg) {
			return arg
		}
	}
	return nil
}

// ParseStructTag 解析结构体的tag
func ParseStructTag(rawTag reflect.StructTag) map[string][]string {
	results := make(map[string][]string, 0)
	for _, tagString := range strings.Split(string(rawTag), " ") {
		tag := strings.Split(tagString, ":")
		if len(tag) > 1 {
			results[tag[0]] = strings.Split(strings.ReplaceAll(tag[1], `"`, ""), ",")
		} else {
			results[tag[0]] = nil
		}
	}
	return results
}
