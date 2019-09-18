package main

import (
	"github.com/nsqio/go-nsq"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

var producer *nsq.Producer
var err error
func main() {
	nsqd := "127.0.0.1:4150"

	producer, err = nsq.NewProducer(nsqd, nsq.NewConfig())
	exit := make(chan bool)
	ch := make(chan os.Signal,1)
	signal.Notify(ch,
		os.Interrupt,
		os.Kill,
		syscall.SIGINT,
	)
	for i := 0; i < 10; i++ {
		producer.DeferredPublish("delay", time.Second*20, []byte(`"data":"stackoverflow "`+strconv.Itoa(i)))
		if err != nil {
			panic(err)
		}
	}
	go func() {
		select {
		case <-ch:
			exit <- true
		}

	}()

	<-exit
}
