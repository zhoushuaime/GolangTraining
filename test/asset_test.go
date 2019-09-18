package test

import (
	"fmt"
	"strconv"
	"testing"
)

// assetType2String ...
func assetType2String(s interface{}) string  {
	switch s.(type) {
	case string:
		fmt.Println("string")
		return s.(string)
	case int:
		fmt.Println("int")
		return strconv.Itoa(s.(int))
	case int64:
		fmt.Println("int64")
		return strconv.FormatInt(s.(int64),10)
	case float64:
		fmt.Println("float64")
		return strconv.FormatFloat(s.(float64),'f',0,64)
	case float32:
		fmt.Println("float32")
		return strconv.FormatFloat(float64(s.(float32)),'f',0,64)
	default:

	}
	return ""
}

// TestAssert ...
func TestAssert(t *testing.T)  {
	var x interface{}
	x = 1234567890123456789
	res := assetType2String(x)
	fmt.Println("res:",res)
}