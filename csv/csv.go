package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
)

var reg = "^[1-9][0-9]*$"
//var reg = "^[0-9]*$"

func main() {

	csvFile, err := os.Open("./uids.csv")
	if err != nil {
		fmt.Println("open error...,err=", err)
		return
	}
	defer csvFile.Close()
	csvReader := csv.NewReader(csvFile)
	raws, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println("read file error:", err)
		return
	}

	var count int
	for _, raw := range raws {
		if ok, _ := regexp.MatchString(reg, raw[0]); !ok {
			continue
		}
		fmt.Println(raw[0])
		count++
	}

	fmt.Println("count:", count)

}
