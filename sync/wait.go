package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var c1 = make(chan int)

type WgMutex struct {
	sync.WaitGroup
	sync.Mutex
}

func main() {
	start := time.Now()
	runtime.GOMAXPROCS(runtime.NumCPU())
	//wg := sync.WaitGroup{}
	//wg := new(WgMutex)
	//mutex := new(sync.Mutex)
	max, num := 100, 4
	//for i := 0; i < num; i++ {
	//	wg.Add(1)
	//	go number(i, (max/num)*(i+1), wg)
	//
	//}
	//wg.Wait()
	//for i := 0; i < 500000; i++ {
	//	fmt.Println(i)
	//}
	for i := 0; i < num; i++ {
		go number2((max / num) * (i + 1))

	}
	for i := 0; i < num; i++ {
		fmt.Println(<-c1)
	}
	fmt.Println(time.Since(start))
}

func number(n int, m int, wg *WgMutex) {
	//wg.Lock()
	fmt.Printf("%d %d", n, m)
	//for i := 0; i < 10; i++ {
	//	fmt.Printf("%d ", i)
	//}
	//wg.Unlock()
	fmt.Println()
	wg.Done()

}

func no(n int) {
	fmt.Printf("%d ", n)
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}
func number2(m int) {

	c1 <- m
}
