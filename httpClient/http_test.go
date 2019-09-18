package httpClient

import (
	"encoding/json"
	"fmt"
	"sync"
	"testing"
)

func TestHttp(t *testing.T) {

	wg := sync.WaitGroup{}
	//client := http.Client{}
	url := "http://127.0.0.1/v1/notify"
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			resp := make(map[string]interface{})
			err := Get(url, &resp)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("i=%v,result=%v\n", i, resp)
			wg.Done()
		}(i)
	}
	wg.Wait()

}

var data = []byte(`
{
    "dst":"test"
   
}`)

func TestHttpPost(t *testing.T) {
	postdata := make(map[string]interface{})
	var lock sync.RWMutex
	_ = json.Unmarshal(data, &postdata)
	out := make(map[string]interface{})
	wg := sync.WaitGroup{}
	//client := http.Client{}
	url := "http://127.0.0.1:8080/antispam"
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			lock.Lock()
			postdata["index"] = i
			defer lock.Unlock()
			req, _ := json.Marshal(postdata)
			fmt.Printf("post_data:%+v\n", postdata)
			defer wg.Done()
			err := PostJson(url, req, &out)
			if err != nil {
				t.Error(err)
				return
			}
			t.Logf("i=%v,result:%+v\n", i, out)

		}(i)
	}
	wg.Wait()

}
