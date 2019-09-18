package goroutine_job

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	input := []int{1, 2, 3, 4}
	chs := make([]chan string, len(input))
	chLimit := make(chan bool, 1)
	timeout := 2
	fmt.Println("execute job...")
	startTime := time.Now()
	for i, sleepTime := range input {
		chs[i] = make(chan string, 1)
		chLimit <- true
		go limitFunc(chLimit, strconv.Itoa(i), chs[i], sleepTime, timeout)
	}

	for _, ch := range chs {
		fmt.Println(<-ch)
	}
	endTime := time.Now()
	fmt.Printf("execute all job finished. Process time %s. Number of task is %v", endTime.Sub(startTime), len(input))

}

func limitFunc(chLimit chan bool, taskId string, ch chan string, sleepTime, timeout int) {
	runTask(taskId, ch, sleepTime, timeout)
	<-chLimit
}

func runTask(taskId string, ch chan string, sleepTime, timeout int) {
	chRun := make(chan string)
	go run(taskId, sleepTime, chRun)
	select {
	case re := <-chRun:
		ch <- re
	case <-time.After(time.Duration(timeout) * time.Second):
		re := fmt.Sprintf("task id %v , timeout", taskId)
		ch <- re
	}

}

func run(taskId string, sleeptime int, ch chan string) {
	time.Sleep(time.Duration(sleeptime) * time.Second)
	ch <- fmt.Sprintf("task id %v , sleep %d second", taskId, sleeptime)
	return
}
