package goroutine_job

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

func main() {
	
	wg := sync.WaitGroup{}
	for i := 0; i < 15; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			if err := goroutine(i); err != nil {
				fmt.Println(err)
				return
			}
		}(i)
	}
	wg.Wait()
}

func goroutine(i int) error {
	timout := time.Duration(i) * time.Second
	result := make(chan interface{})
	go func() {
		time.Sleep(time.Second * 2)
		result <- map[string]interface{}{"code": i}
	}()

	select {
	case res := <-result:
		fmt.Println("result:", res)
	case <-time.After(timout):
		return errors.New("timeout")
	}

	return nil

}
