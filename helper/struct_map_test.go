package helper

import (
	"testing"
)

func TestConvertStruct2Map(t *testing.T) {
	m := struct {
		Name string `json:"name"`
		Did int64  `json:"did"`
	}{}
	m.Name = "test"
	m.Did = 1689770042046021633
	res, err := ConvertStruct2MapByReflect(m)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(res)
}
