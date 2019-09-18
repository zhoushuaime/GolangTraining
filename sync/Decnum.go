package main

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	number int
	sync.Mutex
}

func (sc *SafeCounter) Increment() {
	sc.Lock()
	sc.number++
	sc.Unlock()
}
func (sc *SafeCounter) Decrement() {
	sc.Lock()
	sc.number--
	sc.Unlock()
}
func (sc *SafeCounter) getNumber() int {
	sc.Lock()
	number:=sc.number
	sc.Unlock()
	return number
}
func main() {
	sc := new(SafeCounter)
	for i := 0; i < 100; i++ {
		go sc.Increment()
		go sc.Decrement()
	}
	fmt.Println(sc.getNumber())
}
