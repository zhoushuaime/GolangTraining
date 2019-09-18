package test

import (
	"fmt"
	"regexp"
	"testing"
)



func getInstane() *Config {
	oSingle.Do(
		func() {
			config = new(Config)
		})
	return config
}

// TestTime ...
func TestTime(t *testing.T) {

	txnhash := "testHash"
	go func() {
		var x string
		x = txnhash
		fmt.Println("x:", x)
	}()
	fmt.Println(txnhash)
}

// TestRegexpNumber ...
func TestRegexpNumber(t *testing.T) {
	p := "\\d{6}"
	str := "178232"
	m, err := regexp.MatchString(p, str)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("res:", m)

}