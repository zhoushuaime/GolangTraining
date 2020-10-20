package client

import (
	"fmt"
	"strings"
	"sync"
	"testing"
	"zhoushuai.com/GolangTraining/tidb/tool"
)

// TestGetTransDetail ...
func TestGetTransDetail(t *testing.T) {

	cli := NewTransClient(1)
	res, err := cli.GetTransDetail()
	if err != nil {
		t.Fatalf("test failed,err:%v", err)
	}
	t.Logf("result:%v", res)

}

// TestMultiClientExecTrans ...
func TestMultiClientExecTrans(t *testing.T) {

	wg := sync.WaitGroup{}
	wg.Add(2) // two client
	result := make([]TransDetail, 0)
	for i := 0; i < 2; i++ {
		go func(id int) {
			defer wg.Done()
			id++
			cli := NewTransClient(id)
			res, err := cli.GetTransDetail()
			if err != nil {
				t.Fatalf("test failed,err:%v", err)
			}

			result = append(result, res...)
		}(i)
	}
	wg.Wait()

	for _, v := range result {
		t.Logf("%+v", v)
	}
	separator := strings.Repeat("=", 20)
	fmt.Println(fmt.Sprintf("%v finisehd %v", separator, separator))

}

// TestGetAllTransDetail ...
func TestGetAllTransDetail(t *testing.T) {
	cli := NewTransClient()
	res, err := cli.GetAllTransDetail()
	if err != nil {
		t.Fatalf("test failed,err:%v", err)
	}

	for _, v := range res {
		t.Logf("result:%+v", v)
	}

	t.Log(strings.Repeat("=", 50))

	// 全排列数据
	input := make([]interface{}, 0)
	for i := 0; i < len(res); i++ {
		input = append(input, res[i])
	}

	out := tool.Permutation(input)

	for k, v := range out {
		t.Logf("k:%v,%+v", k, v)
	}

}
