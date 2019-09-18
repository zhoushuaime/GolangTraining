package other

import "fmt"

const MAX = 50

var fibs [MAX]int // 通过预定义数组 fibs 保存已经计算过的斐波那契序号对应的数值
// []int{0,0,0,0,}
func main() {

  res := fibonacci(5)
  fmt.Println(res)

}
// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144...
func fibonacci(n int) int {
	if n == 1 {
		return 0
	}

	if n == 2 {
		return 1
	}
	index := n - 1
	if fibs[n] != 0 {
		return fibs[index]
	}
	num := fibonacci(n-1)+ fibonacci(n-2)
	fibs[index] = num
	return fibs[index]
}
