package main

import (
	"fmt"
	"sync"
)

var count int
var mutexLock sync.Mutex
var wg sync.WaitGroup

func main() {
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			mutexLock.Lock()
			count++
			fmt.Println(count)
			mutexLock.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("finished")
}
