package test

import (
	"fmt"
	"testing"
	"time"
)

func TestNumber(t *testing.T) {
	n := 75
	base := 2
	result := []int{}
	for n != 0 {
		res := n % base
		n = n / base
		result = append(result, res)
		fmt.Print(res)
	}
	fmt.Println("\n=================result================")

	len := len(result)

	for i := len - 1; i >= 0; i-- {
		fmt.Print(result[i])
	}
	const run = time.Second
	for begin := time.Now(); time.Since(begin) <= run; {
		t.Logf("%v\n",begin)

	}
}
