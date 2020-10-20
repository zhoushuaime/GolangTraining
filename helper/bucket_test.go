package helper

import (
	"fmt"
	"testing"
)

type Str struct {
	Input    int
	Expected int
	Data     interface{}
}

// SetValue 接口实现...
func (s *Str) SetValue(value interface{}) {
	s.Data = value
}

//GetValue
func (s *Str) GetValue() []interface{} {
	str := s.Data.([]string)
	res := make([]interface{}, 0)
	for _, v := range str {
		res = append(res, v)
	}
	return res
}

// TestBucketSlice ...
func TestBucketSlice(t *testing.T) {
	data := &Str{}
	data.Data = []string{"1", "2", "3", "4", "5", "6", "7"}

	bucket := NewBucket()
	bucketCount := bucket.CalcBucketCount(data, 5) //计算下真实能分几个桶，这里除以10，显然只能分1个桶

	fmt.Println("count:",bucketCount)
	res := bucket.BucketSlice(data, bucketCount)
	for _, v := range res {
		fmt.Println("v:", v)
	}
}

// TestCalcBucketCount ...
func TestCalcBucketCount(t *testing.T) {
	bucket := NewBucket()
	testCases := []Str{
		{
			Data:     []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13"},
			Input:    10, // 输入的除数
			Expected: 2,  // 期望值
		},

		{
			Data:     []string{"2", "3", "5", "7", "19", "D"}, // 5/3=1  5%3=2
			Input:    3,
			Expected: 2,
		},
		{
			Data:     []string{"a", "c", "d", "e", "F", "D", "3", "4"},
			Input:    3,
			Expected: 3,
		},

		{
			Data:     []string{"hello", "world", "hi"},
			Input:    2,
			Expected: 2,
		},
	}

	for i, c := range testCases {
		actual := bucket.CalcBucketCount(&c, c.Input)
		if c.Expected != actual {
			t.Errorf("test failed,not match with expected.expected:%v,got:%v,index:%v", c.Expected, actual, i)
		}

	}

}
