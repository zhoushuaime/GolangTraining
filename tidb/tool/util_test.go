package tool

import (
	"fmt"
	"strings"
	"testing"
)

// TestReadFile ...
func TestReadFile(t *testing.T) {
	fileName := "../sql/sql1.sql"
	res, err := ReadFile(fileName)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("res:%v", strings.Split(res, "\n"))
}

// TestPermutation ...
func TestPermutation(t *testing.T) {

	res := []interface{}{"A", "B", "C"}
	out := Permutation(res)
	for _, v := range out {
		fmt.Println(v)
	}

	fmt.Println(strings.Repeat("=", 20))
}

func TestGetCurrentDir(t *testing.T) {

	res := GetCurrentDir()
	t.Log(res)
}
