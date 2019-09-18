package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"image/png"
	"os"
)

/*
@Time : 2018/11/2 10:34 
@Author : zhoushuai
@File : barcode
@Software: GoLand
*/

func main() {
	qrCode, _ := qr.Encode("hi", qr.M, qr.Auto)
	qrCode, _ = barcode.Scale(qrCode, 256, 256)
	file, _ := os.Create("qr2.png")
	png.Encode(file, qrCode)
	defer file.Close()
	buf := bytes.NewBuffer(nil)
	err := png.Encode(buf, qrCode)
	if err != nil {
		fmt.Println("error:",err)
		return
	}
	dst := make([]byte, 50000)
	base64.StdEncoding.Encode(dst, buf.Bytes())
	fmt.Println("dst:",string(dst))
}
