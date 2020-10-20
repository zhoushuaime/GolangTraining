package helper

import (
	"fmt"
	"testing"
)

// TestIsExistSliceInt ...
func TestIsExistSliceInt(t *testing.T) {
	// IsExistSliceInt ...
	slice1 := []int{4, 5}
	slice2 := []int{10, 9, 8, 7, 5, 6, 7}
	slice3 := []int{11, 12, 13, 14, 42, 43, 44, 56}
	testCases := []struct {
		Data         []int
		DestData     int
		ExpectResult bool
	}{
		{Data: slice1, DestData: 1, ExpectResult: false},
		{Data: slice2, DestData: 6, ExpectResult: true},
		{Data: slice3, DestData: 11, ExpectResult: true},
	}

	for i, c := range testCases {
		actual := IsExistSliceInt(c.Data, c.DestData)
		if actual != c.ExpectResult {
			t.Errorf("test failed,case=%v,expected=%v,actual=%v,input=%v", i+1, c.ExpectResult, actual, c.Data)
			return
		}
	}

	t.Log("all cases test ok!")

}

func TestSearchByBinary(t *testing.T) {
	slice1 := []int{4, 5}
	slice2 := []int{5, 6, 7, 12, 45, 56, 77, 467}
	slice3 := []int{9, 12, 45, 123, 1234, 1512}
	testCases := []struct {
		Data         []int
		DestData     int
		ExpectResult int
	}{
		{Data: slice1, DestData: 1, ExpectResult: -1},
		{Data: slice2, DestData: 6, ExpectResult: 1},
		{Data: slice3, DestData: 123, ExpectResult: 3},
	}

	for i, c := range testCases {
		actual := SearchByBinary(c.Data, c.DestData, len(c.Data)-1, 0)
		fmt.Println("actual:",actual)
		if actual != c.ExpectResult {
			t.Errorf("test failed,case=%v,expected=%v,actual=%v,input=%v", i+1, c.ExpectResult, actual, c.Data)
			return
		}
	}

	t.Log("all cases test ok!")
}
