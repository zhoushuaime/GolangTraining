package main

import (
	"fmt"
	"math/big"
	"reflect"
	"runtime"
	"strconv"
	"strings"
)

type slice struct {
	//make([]string, 0, 10)
	runtime.Error
}

func HandleData() {

	value := "de0b6b3a7640000"

	if strings.HasPrefix(value, "0x") {
		value = strings.TrimPrefix(value, "0x")
	} else if strings.HasPrefix(value, "0X") {
		value = strings.TrimPrefix(value, "0X")
	}

	fmt.Println("value:", value[len("0x"):])

	//v2 := "1bc16d674ec80000"
	// 收取手续费
	val := new(big.Int)
	_, ok := val.SetString(value, 16)
	if !ok {
		fmt.Println("not 16")
		return
	}

	fmt.Println(val.String())
	val1 := new(big.Int)
	val1.SetString(val.Text(16), 16)
	fmt.Println("val1:", val1.String())
	fmt.Println("val:", val.Text(16))

}

type name = string
type M map[string]interface{}

func (m M) Get(key string) interface{} {
	return m[key]
}

// HandlerAmount ... token 1:1兑换
func HandlerAmount(value string) (int64, string) {
	val := new(big.Int)
	val.SetString(value, 16)

	base := new(big.Int)
	b := strconv.Itoa(1e18) // 1后面18个0
	base.SetString(b, 10)
	val.Div(val, base)
	return val.Int64(), val.String()
}
func main() {
	//HandleData()
	fmt.Println(s())

}
func s() bool {
	s1 := []string{}
	//s2 := make([]string, 0)
	var s3 []string
	return reflect.DeepEqual(s1, s3)
}
