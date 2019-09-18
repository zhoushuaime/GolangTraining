package test

import (
	"encoding/json"
	"fmt"
	"github.com/nsqio/go-nsq"
	"testing"
	"time"
)

type Asset struct {
	AssetID string `json:"asset_id"`
	Name    string `json:"name"`
}

// 创建生产者
var producer *nsq.Producer
var consumer *nsq.Consumer
var err error
var addr = "127.0.0.1:4150"

func init() {
	producer, err = nsq.NewProducer(addr, nsq.NewConfig())

}

// NsqPublic ...
func NsqPublic(assetID, name string, duration int) error {
	data := struct {
		AssetID string `json:"asset_id"`
		Name    string `json:"name"`
	}{}
	if err != nil {
		return err
	}
	res, _ := json.Marshal(data)
	return producer.DeferredPublish("notice", time.Duration(duration)*time.Second, res)
}

// TestNsq ...
func TestNsq(t *testing.T) {
	assetID := "hello"
	name := "world"
	duration := 5
	err := NsqPublic(assetID, name, duration)
	if err != nil {
		t.Error(err)
	}

}

func main() {

	cfg := nsq.NewConfig()
	consumer, _ = nsq.NewConsumer("notice", "test", cfg)
	res := Asset{}
	consumer.AddHandler(&res)
	err = consumer.ConnectToNSQD(addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-consumer.StopChan
	status := consumer.Stats()
	fmt.Printf("receiced:%v,finished:%v", status.MessagesReceived, status.MessagesFinished)
}

func (a *Asset) HandleMessage(message *nsq.Message) error {
	fmt.Printf("receive:%v,message:%v\n", message.NSQDAddress, string(message.Body))
	return nil
}
