package main

import (
	"testing"
)

// TestStack ...
func TestStack(t *testing.T) {
	s := Stack{}

	data, err := s.GetTop()
	if err != nil {
		t.Errorf("err:%v",err)
	}
	t.Log("data:",data)
	s.InitStack(3)
	str := "ABCDEFG"

	for i := 0; i < len(str); i++ {
		if err := s.Push(string(str[i])); err != nil {
			t.Errorf("err:%v", err)
		}
	}
}
