package httpClient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// MyError ...
type MyError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Get ...
func Get(url string, out interface{}) error {
	res := make(chan *http.Response)
	httpError := make(chan *error)
	go func() { // 开启 goroutine发送 GET 请求
		resp, err := http.Get(url)
		if err != nil {
			httpError <- &err
		}
		res <- resp
	}()
	for {
		select {
		case r := <-res: // 正确响应
			err := parseResp(r, out)
			if r != nil {
				defer r.Body.Close()
				r.Request.Close = true
			}
			return err
		case err := <-httpError: // 错误响应
			return *err
		case <-time.After(10 * time.Second): // 10秒无响应，则请求超时，退出
			return errors.New("http time out!")
		}
	}

}

// GetHttp ...
func GetHttp(url string, out interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	resp.Request.Close = true
	return parseResp(resp, out)
}

func parseResp(resp *http.Response, out interface{}) error {
	if resp == nil {
		return errors.New("resp is nil")
	}

	switch resp.StatusCode {
	case 200:
		if out == nil {
			return nil
		}
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		err = json.Unmarshal(b, &out)
		if err != nil {
			return err
		}

		return nil
	case 201, 202, 204:
		return nil
	case 400, 401, 404, 500:
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		e := &MyError{}
		err = json.Unmarshal(b, e)
		if err != nil {
			return err
		}

		return errors.New(e.Code)
	default:
		return fmt.Errorf("不支持的状态码[%d]", resp.StatusCode)
	}
}

func PostJson(url string, params []byte, out interface{}) error {
	client := http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(params))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp != nil {
		defer resp.Body.Close()
	}
	resp.Request.Close = true

	return parseResp(resp, out)

}