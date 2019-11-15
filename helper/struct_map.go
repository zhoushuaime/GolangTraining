package helper

import (
	"fmt"
	"reflect"
)

// ConvertStruct2MapByReflect ...
func ConvertStruct2MapByReflect(data interface{}) (map[string]interface{},error) {
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)

	if v.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return nil,fmt.Errorf("input data not struct type,got:%v",v.Kind().String())
	}

	num := t.NumField()
	out := make(map[string]interface{})
	for i := 0; i < num; i++ {
		out[t.Field(i).Name] = v.Field(i).Interface()
	}
	return out,nil

}
