package main

import "fmt"


/**
归并排序  迭代实现
 */
func mergeSor(arr []int) []int {
	n := len(arr)
	if n < 2 {
		return arr
	}
	mid := n / 2
	left := mergeSor(arr[:mid])  // 归并左边
	right := mergeSor(arr[mid:]) // 归并右边
	fmt.Println("arr[:mid]", arr[:mid])
	fmt.Println("arr[mid:]", arr[mid:])
	//return nil
	return merge(left, right)
}

func merge(left []int, right []int) []int {
	tmp := make([]int, 0)
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			tmp = append(tmp, left[i])
			i++
		} else {
			tmp = append(tmp, right[j])
			j++

		}
	}
	// 将左右边末尾的添加进临时list中
	tmp = append(tmp, left[i:]...)
	tmp = append(tmp, right[j:]...)
	fmt.Println("tmp:", tmp)
	return tmp
}
func main() {
	arr := []int{6, 4, 98, 32, 12, 43}
	var max int = arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	fmt.Println("max number:", max)

}

func t(a int) int {
	a ++
	if a > 3 {
		return a

	}
	x := t(a)
	fmt.Println("a:", a)
	return x
}

func add(a int) int {
	a++
	return a
}

func m(a []int) []int {
	if len(a) < 2 {
		fmt.Println("return:", a)
		return a
	}
	mid := len(a) / 2
	fmt.Println("a[:mid]:", a[:mid])
	fmt.Println("a[mid:]:", a[mid:])
	l := m(a[:mid])
	r := m(a[mid:])
	fmt.Println("l:", l)
	fmt.Println("r:", r)

	return l
}
func Merge(a []int) []int {
	if len(a) < 2 {
		return a
	}
	mid := len(a) / 2
	fmt.Println("a", a)
	fmt.Println("a[:mid]:", a[:mid])
	left := Merge(a[:mid])
	right := Merge(a[mid:])
	fmt.Println("a[mid:]:", a[mid:])
	fmt.Println("left:", left)
	fmt.Println("right:", right)
	return SortTwoArrgs(left, right)
}
func SortTwoArrgs(a []int, b []int) []int {
	c := make([]int, 0)
	left, right := 0, 0
	for left < len(a) && right < len(b) {
		if a[left] > b[right] {
			c = append(c, b[right])
			right++

		} else {
			c = append(c, a[left])
			left++
		}
	}
	// 剩余元素添加进数组
	if left < right {
		c = append(c, a[left:]...)
	} else {
		c = append(c, b[right:]...)
	}

	fmt.Println("sort finished...", c)
	return c
}
