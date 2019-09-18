package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	data := make(chan int)
	go func() {
		time.Sleep(5 * time.Second)
		data <- 1
	}()
	select {
	case <-ctx.Done():
		fmt.Println("http timeout!")
	case x := <-data:
		fmt.Println("data:", x)
	}
}
