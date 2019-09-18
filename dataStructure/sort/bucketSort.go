package main

import "fmt"

/**
桶排序
 */
func BucketSort(a []int) {
	var barrel []int = make([]int, 11)
	for _, v := range a {
		barrel[v]++
	}
	//fmt.Println(barrel)
	for i := 0; i < len(barrel); i++ {
		for j := 0; j < barrel[i]; j++ { // 出现的次数，出现几次打印几次
			fmt.Printf("%d ", i)
		}
	}
}

func main() {
	var a = []int{6, 1, 3, 7, 9, 3, 4, 5, 10, 2}
	//var a = []int{6, 1, 3, 7, 9, 3, 4, 5}
	BucketSort(a)
}
