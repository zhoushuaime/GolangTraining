package main

import (
	"math/rand"
	"testing"
	"time"
)


// BenchmarkChanSort ...
func BenchmarkChanSort(t *testing.B) {
	data := make([]int, 1000000)
	for i := 0; i < len(data); i++ {
		data[i] = rand.Intn(1000)
	}
	start := time.Now()
	Sort(CreateData(data))
	end := time.Now()
	t.Logf("elapsed time:%v", end.Sub(start))
	/*
	c := Sort(CreateData(data))
	for {
		if num, ok := <-c; ok {
			t.Logf("%v\n",num)
		} else {
			break
		}
	}
	*/

}

// CreateData 初始数据写入管道
func CreateData(data []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, v := range data {
			out <- v
		}
		close(out)
	}()
	return out
}

// Sort 排序后写入
func Sort(data <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		a := []int{}
		for v := range data{
			a = append(a, v)
		}
		a = MergerSort(a)

		for _, v := range a {
			out <- v
		}
		close(out)
	}()

	return out
}

// sortByMerger 归并排序
func sortByMerger(left, right []int) (result []int) {
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}

	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return result
}

// MergerSort 左右递归排序
func MergerSort(r []int) []int {
	length := len(r)
	if length <= 1 {
		return r
	}
	num := length / 2
	left := MergerSort(r[:num])
	right := MergerSort(r[num:])
	return sortByMerger(left, right)
}

/*
归并排序
 */
func Merger(in1, in2 chan int) <-chan int {
	out := make(chan int)
	go func() {
		v1, ok1 := <-in1
		v2, ok2 := <-in2
		for ok1 || ok2 {
			if !ok2 || (ok1 && v1 <= v2) {
				out <- v1
				v1, ok1 = <-in1
			} else {
				out <- v2
				v2, ok2 = <-in2
			}
		}
		close(out)

	}()
	return out
}
