package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
)

func main() {
	data := []byte(`this is a test`)
	res, err := send(data)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	// 这里res结果其实是个base64编码后的结果 ,结果为：dGhpcyBpcyBhIHRlc3Q=
	// 1.将base64换成字符串，有下面几种方式
	str := string(res[:]) // 第一种
	fmt.Println("base64 result:", str)
	// str = fmt.Sprintf("%s",res) // 第二种
	// fmt.Sprintf("%s",res)和string(res),string(res[:]) 在下面的base64解码都会出错，后面做说明,尝试将[]byte通过byte buf转string
	buf := bytes.Buffer{} // 第三种
	n, err := buf.Write(res)
	if err != nil {
		fmt.Println("buf.Write(res) error:", err)
		return
	}
	fmt.Println("buf.Write(res) success,len:", n)
	str = buf.String()

	// 2.发现上面确实能正确转成字符串，但是base64会失败
	// 想起之前marshal过一次，这一次，直接把[]byte unmarshal到string上试试
	if err := json.Unmarshal(res, &str); err != nil {
		fmt.Println("json unmarshal res error:", err)
		return
	}
	fmt.Println("json.Unmarshal result:", str)

	// 3. decode ,发现成功
	decodeResult, err := base64.URLEncoding.DecodeString(str)
	if err != nil {
		fmt.Println("base64 decode error:", err)
	}
	fmt.Println("decodeResult:", string(decodeResult))

}

func send(data interface{}) ([]byte, error) {
	// 这儿先不用断言成[]byte，尝试通过其他方式将interface{}转为[]byte，借助于json marshal转成[]byte下
	var reader io.Reader
	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	reader = bytes.NewReader(b)
	readData, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return readData, nil
}
