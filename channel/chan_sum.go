package main

import (
	"fmt"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	fmt.Println(s[:len(s)/2])
	go sum(s[:len(s)/2], c)
	// go sum(s[len(s)/2:], c)
	//x, y := <-c, <-c // receive from c

	//fmt.Println(x, y, x+y)

	x := <-c
	m := make(map[string]string)
	m["key"] = "value"
	fmt.Println(x)

}
