package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"strconv"
)

func main() {
	provider()
}

func provider() {
	config := nsq.NewConfig()
	nsqProducer, _ := nsq.NewProducer("127.0.0.1:4150", config)
	err := nsqProducer.Ping()
	if err != nil {
		fmt.Println("nsq error ", err)
		return
	}
	defer nsqProducer.Stop()
	topicName := "test"
	count := 4
	for i := 0; i < count; i++ {
		msg:="this is a test "+strconv.Itoa(i)
		err = nsqProducer.Publish(topicName, []byte(msg))
		if err != nil {
			fmt.Println("error:",err)
		}
		fmt.Println(i," send")
	}

}


