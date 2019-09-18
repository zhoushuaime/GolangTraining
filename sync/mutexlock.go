package main

import (
	"fmt"
	"sync"
)

var x int
var sum []int
var mutex = sync.Mutex{}
//var mutex = sync.RWMutex{}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(x int) {
			res := Gohead(1)
			sum = append(sum, res)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("\n===============================")
	amount := 0
	for _, v := range sum {
		amount += v
	}
	fmt.Println("sum:", amount)
	fmt.Printf("sum:%v", sum)

}

func Gohead(i int) int {
	mutex.Lock()
	defer mutex.Unlock()
	x += i
	fmt.Printf("%d\t", x)
	return x
}
