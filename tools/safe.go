package main

import (
	"fmt"
	"sync"
	"unsafe"
)

func main1() {
	for i := 0; i < 10; i++ {
		temp := i
		p := uintptr(unsafe.Pointer(&temp))
		fmt.Printf("address of temp:%x\n", p)
	}
}

func main2() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		temp := 1
		go func() {
			defer wg.Done()
			p := uintptr(unsafe.Pointer(&temp))
			fmt.Printf("address of temp:%x\n", p)

		}()
	}
	wg.Wait()
}

func main() {
	for i := 0; i < 10; i++ {
		temp := i
		fmt.Println(&temp)
	}
}
