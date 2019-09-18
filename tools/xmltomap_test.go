package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"testing"
	"time"
)

// OrderInfo ...
type OrderInfo struct {
	OrderID     string    `xml:"order_id" json:"order_id"`
	CreatedAt   time.Time `xml:"created_at" json:"created_at"`
	CallBackURI string    `xml:"call_back_uri" json:"call_back_uri"`
}

// TestXmlToMap ...
func TestXmlToMap(t *testing.T) {
	order := &OrderInfo{
		OrderID:     "123",
		CreatedAt:   time.Now(),
		CallBackURI: "https://",
	}
	XMLToMap(order)
}

// XMLToMap ...
func XMLToMap(req interface{}) {
	data, err := xml.Marshal(req)
	if err != nil {
		return
	}
	m := XmlToMap(data)
	str := bytes.NewBufferString("")
	for k, v := range m {
		str.WriteString(fmt.Sprintf("<%s>%s</%s>", k, v, k))
	}
	xmlStr := fmt.Sprintf("<xml>%s</xml>", str.String())
	fmt.Println(xmlStr)
}
