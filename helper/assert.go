package helper

import (
	"strconv"
)

// assertType2String ...float直接用fmt.Sprint()会有问题
func assertType2String(s interface{}) string {
	switch s.(type) {
	case string:
		return s.(string)
	case int:
		return strconv.Itoa(s.(int))
	case int64:
		return strconv.FormatInt(s.(int64), 10)
	case float64:
		return strconv.FormatFloat(s.(float64), 'f', 0, 64)
	case float32:
		return strconv.FormatFloat(float64(s.(float32)), 'f', 0, 64)
	default:

	}
	return ""
}

// assertType2Int ...
func assertType2Int(s interface{}) (result int, err error) {
	switch s.(type) {
	case int:
		result = s.(int)
	case int64:
		result = int(s.(int64))
	case float64:
		result = int(s.(float64))
	case float32:
		result = int(s.(float32))
	case string:
		result, err = strconv.Atoi(s.(string))
	}
	return result, err
}

// assertType2Int64 ...
func assertType2Int64(s interface{}) (result int64, err error) {
	switch s.(type) {
	case int:
		result = int64(s.(int))
	case int64:
		result = s.(int64)
	case float64:
		result = int64(s.(float64))
	case float32:
		result = int64(s.(float32))
	case string:
		result, err = strconv.ParseInt(s.(string), 10, 64)
	}
	return result, err
}
