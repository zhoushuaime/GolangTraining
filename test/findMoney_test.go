package test

import (
	"testing"
)

// TestFindMoney ...
func TestFindMoney(t *testing.T) {
	arr := []int{5, 2, 3}
	aim := 20
	res := find2(aim, arr)
	t.Log("res:", res)
	t.Log("ARR len:", len(arr))
}

/**
经典的动态规划问题
 */

/**
“给定数组arr，arr中所有的值都为正数且不重复。每个值代表一种面值的货币，每种面值的货币可以使用任意张，再给定一个整数aim代表要找的钱数，求组成aim的最少货币数。”
 */

func find(aim int, arr []int) int {

	if aim == 0 || len(arr) == 0 {
		return aim
	}

	for _, v := range arr {

		if aim%v == 0 {

			return aim / v
		}

	}
	return -1
}

/**
“给定数组arr，arr中所有的值都为正数。每个值仅代表一张钱的面值，再给定一个整数aim代表要找的钱数，求组成aim的最少货币数。”
 */

func find2(aim int, arr []int) int {
	for i, v := range arr {
		if aim == v {
			return 1
		}
		if aim == arr[i]+arr[i+1] {
			return 2
		}
	}
	return -1
}
