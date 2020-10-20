package helper

import "sort"

// IsExistSliceString ...
func IsExistSliceString(src []string, dst string) bool {

	for _, v := range src {
		if v == dst {
			return true
		}
	}
	return false
}

// IsExistSliceMap ...
func IsExistSliceMap(src []map[string]string, dst string) string {
	for _, v := range src {
		if _, ok := v[dst]; ok {
			return v[dst]
		}
	}

	return ""

}

// IsExistSliceInt ...
func IsExistSliceInt(src []int, dst int) bool {
	sort.Ints(src)
	return SearchByBinary(src, dst, len(src)-1, 0) != -1

}

// SearchByBinary 。。。
func SearchByBinary(src []int, dst int, high, low int) int {

	if low > high {
		return -1
	}

	mid := (low + high) / 2

	if dst > src[mid] {
		return SearchByBinary(src, dst, high, mid+1)
	} else if dst < src[mid] {
		return SearchByBinary(src, dst, mid-1, low)
	}

	return mid

}
