package test

import (
	"strings"
	"testing"
	"time"
)

/*
@Time : 2018/12/19 18:31
@Author : joshua
@File : bibao_test
@Software: GoLand
*/

// TestBiBao ...
func TestBiBao(t *testing.T) {
	timestamp := time.Now().Unix()
	t.Logf("timestamp:%v\n", timestamp)
	now := time.Now()
	t.Logf("now:%v\n", now)
}

func HandleCnEn(str, lang string) string {
	StartFlag, EndFlag := "[", "]"
	res := str
	if strings.Contains(str, StartFlag) && strings.HasSuffix(str, EndFlag) {
		// cn
		res = str[:strings.Index(str, StartFlag)]
		// en
		if lang == "en" {
			res = str[strings.Index(str, StartFlag)+1 : strings.Index(str, EndFlag)]
		}

	} else {
		if lang == "en" {
			res = "UnKnown"
		}
	}

	return res
}

func f(i int, t *testing.T) func() {

	return func() {
		i++
		t.Log(i)
	}


}
