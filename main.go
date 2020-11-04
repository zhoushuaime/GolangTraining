package main

import (
	"fmt"
)

func main() {

	input := []int{1, 1, 2, 2, 3, 4, 4, 5, 5, 6, 6, 7, 7}
	res := findSingleNumber(input)
	fmt.Println(res)

}

// findSingleNumber ...
func findSingleNumber(input []int) int {

	if len(input) == 0 {
		return -1

	}
	res := 0

	for _, v := range input {
		res = res ^ v
	}

	return res
}
