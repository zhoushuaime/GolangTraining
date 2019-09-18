package main

import "fmt"

// QuickSort2 ...
func QuickSort2(a []int, left, right int) {

	if left > right {
		return
	}
	start, end := left, right
	key := a[left] // 基准数

	//for start < end {
	for start != end {
		for end > start && a[end] >= key { // 先从右往左找
			end --
		}

		for end > start && a[start] <= key { // 再从左往右找
			start ++
		}

		// 当start和end没有相遇时，交换两个数的位置

		if start < end {
			var temp int
			temp = a[end]
			a[end] = a[start]
			a[start] = temp
		}

	}

	// 最终将基准数归位
	a[left] = a[start]
	a[start] = key
	QuickSort2(a, left, start-1)
	//QuickSort2(a, start+1, right)
	QuickSort2(a, end+1, right)
}

func main() {
	var a = []int{6, 1, 3, 7, 9, 3, 4, 5, 10, 8}
	QuickSort2(a, 0, len(a)-1)
	fmt.Println(a)

}
