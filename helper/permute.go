package helper

// Permutation 数组全排列
func Permutation(arr []interface{}) [][]interface{} {
	res := make([][]interface{}, 0)
	l := len(arr)
	if len(arr) == 0 {
		return res
	}

	var arrange func(pos int)
	arrange = func(pos int) {
		if pos == l-1 {
			x := make([]interface{}, l)
			copy(x, arr)
			res = append(res, x)
			return
		}

		for i := pos; i < l; i++ {
			swap(arr, i, pos)
			arrange(pos + 1)
			swap(arr, i, pos)
		}
	}
	arrange(0)
	return res

}

func swap(arr []interface{}, i, j int) {
	if i != j && i < len(arr) && j < len(arr) {
		temp := arr[i]
		arr[i] = arr[j]
		arr[j] = temp
	}

}
