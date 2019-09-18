package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	data, err := ioutil.ReadFile("count_line.go")
	if err != nil {
		fmt.Println("read error:", err)
		return
	}
	//fmt.Println(string(data))
	//return
	_ = data
	var r io.Reader
	r, _ = os.Open("count_line.go")
	c, err := CountLine2(r)
	if err != nil {
		fmt.Println("read file count error:", err)
		return
	}
	fmt.Println("count:", c)
}

// CountLine1
func CountLine1(r io.Reader) (int, error) {
	var (
		br  = bufio.NewReader(r)
		c   int
		err error
	)
	for {
		_, err = br.ReadString('\n')
		c++
		if err != nil {
			break
		}
	}

	if err != io.EOF {
		return 0, err
	}
	return c, nil

}

// CountLine2 ...
func CountLine2(r io.Reader) (int, error) {
	sc := bufio.NewScanner(r)
	lines := 0
	for sc.Scan() {
		lines++
		fmt.Println(sc.Text())
	}
	return lines, sc.Err()
}
