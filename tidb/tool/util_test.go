package tool

import (
	"fmt"
	"strings"
	"sync"
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
	wg := sync.WaitGroup{}
	wg.Add(2)

	for i := 0; i < 2; i++ {
		go func() {
			defer wg.Done()

		}()
	}
	wg.Wait()
}

func TestGetCurrentDir(t *testing.T) {

	res := GetCurrentDir()
	t.Log(res)
}
