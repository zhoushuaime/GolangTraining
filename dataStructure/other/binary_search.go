package other

import "fmt"

// BinarySearch ...
func BinarySearch(arr []int, low, high, dst int) (index int) {

	if low > high {
		return -1
	}

	mid := low + (high-low)/2

	if dst > arr[mid] { // 从后半部分找
		return BinarySearch(arr, mid+1, high, dst)
	} else if dst < arr[mid] {
		return BinarySearch(arr, low, mid-1, dst)
	}

	return mid
}

func main() {
	arr := []int{1, 4, 6, 7, 9, 12, 43, 45, 65, 122, 211}
	low, high := 0, len(arr)-1
	res := BinarySearch(arr, low, high, 45)
	fmt.Println("index :", res)

}
