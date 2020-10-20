package tool

import (
	"errors"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

// ReadFile ...
func ReadFile(filename string) (string, error) {

	if filename == "" {
		return "", errors.New("filename is empty")
	}

	res, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

var (
	randObj   = rand.New(rand.NewSource(time.Now().UnixNano()))
	randMutex = sync.Mutex{}
)

// RandNum ...
func RandNum() int64 {
	randMutex.Lock()
	defer randMutex.Unlock()
	return randObj.Int63()
}

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

// GetCurrentDir 相对目录获取...
func GetCurrentDir() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	return strings.ReplaceAll(dir, "\\", "/")

}
