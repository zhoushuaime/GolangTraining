package main

import "fmt"

// partSort 部分排序...
func partSort(a []int, left, right int) int {
	var key = a[left]
	for right > left {
		for right > left && a[right] >= key {
			right --
		}

		a[left] = a[right]
		for right > left && a[left] <= key {
			left ++
		}

		a[right] = a[left]
	}
	a[left] = key
	return left
}

// QuickSort 递归排序...
func QuickSort(a []int, left, right int) {
	if left < right {
		p := partSort(a, left, right)
		QuickSort(a, left, p-1)  // 对基准元素对左边元素进行排序
		QuickSort(a, p+1, right) // 对基准元素对右边元素进行排序

	}
}
func main() {
	a := []int{10, 5, 3, 1, 7, 2, 8}
	QuickSort(a, 0, len(a)-1)
	fmt.Println("sort after:", a)
}
