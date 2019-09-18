package main

import (
	"bytes"
	"fmt"
	"hash/crc32"
)

func String(s string) int {
	v := int(crc32.ChecksumIEEE([]byte(s)))
	if v >= 0 {
		return v
	}

	if -v >= 0 {
		return -v
	}

	return 0
}

func Strings(strs []string) string  {
	var buf bytes.Buffer

	for _ ,v := range strs {
		buf.WriteString(fmt.Sprintf("%s-",v))
	}

	return fmt.Sprintf("%d",string(buf.String()))
}

