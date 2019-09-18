package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

type MS struct {
	lock sync.RWMutex
	M    map[string]interface{}
}

type TimeConf struct {
	Timeout time.Duration `json:"timeout"`
}

func main() {
	n, err := fmt.Fprintln(&UpperWriter{os.Stdout}, "hello","world")
	fmt.Printf("n:%v,err:%v", n, err)
}

type UpperWriter struct {
	io.Writer
}

func (upperWriter *UpperWriter) Write(p []byte) (n int, err error) {
	return upperWriter.Writer.Write(bytes.ToUpper(p))
}
