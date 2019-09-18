package main

import (
	"encoding/base64"
	"fmt"
	"github.com/skip2/go-qrcode"
	"strings"
)

/*
@Time : 2018/11/2 10:26 
@Author : zhoushuai
@File : go-qrcode
@Software: GoLand
*/

func main() {
	//err := qrcode.WriteFile("hi", qrcode.Medium, 256, "qr.png")
	//if err != nil {
	//	fmt.Println("qrcode encode error:", err)
	//	return
	//}
	var p []byte
	p, err := qrcode.Encode("https://baidu.com", qrcode.Medium, 256)
	if err != nil {
		fmt.Println("qrcode encode error:", err)
		return
	}

	fmt.Println("ok")
	dst := make([]byte, 5000)
	base64.StdEncoding.Encode(dst, p)
	res := strings.TrimRight(string(dst), "\x00")
	fmt.Println("res:", res)

}
