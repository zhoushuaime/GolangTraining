package helper

import (
	"github.com/google/uuid"
)

// IBucket ...
type IBucket interface {
	GetValue() []interface{}
	SetValue(value interface{})
}

// Bucket ...
type Bucket struct {
	Num   int
	Count int
}

// NewBucket ...
func NewBucket() *Bucket {

	return &Bucket{}
}

// BucketSlice 对数据进行分桶,分几个桶由用户决定,保证每个桶中个数不能相差太远，充分利用...
func (b *Bucket) BucketSlice(bucket IBucket, bucketCount int) [][]interface{} {

	src := bucket.GetValue()
	l := len(src)
	if bucketCount > l {
		bucketCount = l
	}

	remain := l % bucketCount
	var count = l / bucketCount
	if remain > 0 {
		count = count + 1 // 尾部的一个
	}
	result := make([][]interface{}, 0)
	for i := 0; i < bucketCount; i++ {
		temp := make([]interface{}, 0)
		for j := i * count; j < (i+1)*count && j < l; j++ {
			temp = append(temp, src[j])
		}
		// [0,1,2,3,4,5,6]

		result = append(result, temp)
	}

	return result
}

// CalcBucketCount defaultDivision 默认的除多少，也就是一个桶装多少个，以此来计算总的桶个数...
func (b *Bucket) CalcBucketCount(bucket IBucket, defaultDivision int) int {
	if defaultDivision <= 0 {
		panic("<CalcBucketCount> division number must be great than zero")
	}

	src := bucket.GetValue()
	l := len(src)
	b.Num = int(uuid.New().ID())
	for {
		if result := l / defaultDivision; result > 0 {
			r := l % defaultDivision
			if r != 0 && r < defaultDivision && l > defaultDivision {

				return result + 1
			}
			b.Count = result
			return result

		}

		defaultDivision--
	}
}
