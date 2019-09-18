package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

type ConsumerHandle struct {
	Consumer *nsq.Consumer
	MsgGood  int
}

const BaseURL = "127.0.0.1"

var consumers = make([]*nsq.Consumer, 0)
var mux = &sync.Mutex{}

func main() {
	topicName := "test"
	count := 4
	for i := 0; i < count; i++ {
		go readMsg(topicName, i)
	}
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, syscall.SIGINT)
	fmt.Println("server is exiting")
	exit := make(chan bool)
	go func() {
		select {
		case <-ch:
			for _, v := range consumers {
				v.StopChan <- 1
				v.Stop()
			}
			exit <- true
		}
	}()
	<-exit
	fmt.Println("stop server")
}
func (h *ConsumerHandle) HandleMessage(message *nsq.Message) error {
	msg := string(message.Body) + "  " + strconv.Itoa(h.MsgGood)
	fmt.Println(msg)
	return nil
}

func readMsg(topicName string, count int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	config := nsq.NewConfig()
	config.MaxInFlight = 10000
	config.MaxBackoffDuration = time.Minute
	consumer, _ := nsq.NewConsumer(topicName, "ch"+strconv.Itoa(count), config)
	consumerHandle := &ConsumerHandle{Consumer: consumer, MsgGood: count}
	consumer.AddHandler(consumerHandle)
	err := consumer.ConnectToNSQLookupd(BaseURL + ":4161")
	if err != nil {
		fmt.Println(err)
		return
	}
	mux.Lock()
	consumers = append(consumers, consumer)
	mux.Unlock()
	<-consumer.StopChan
	fmt.Println("stop")
}
