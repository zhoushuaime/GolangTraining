package main

import "fmt"

// print 0 2 4 6 8
func main() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Printf("%d ",x)
		case ch <-i:

		}
	}
}
