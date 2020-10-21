package helper

import (
	"fmt"
	"strings"
	"testing"
)

// TestPermutation ...
func TestPermutation(t *testing.T) {

	res := []interface{}{"A", "B", "C"}
	out := Permutation(res)
	for _, v := range out {
		fmt.Println(v)
	}

	fmt.Println(strings.Repeat("=", 20))
}
