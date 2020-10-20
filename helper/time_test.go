package helper

import (
	"fmt"
	"testing"
	"time"
)

// TestFormatTime ....
func TestFormatTime(t *testing.T) {
	t1 := TimeDetail{
		TimeStr: "2020-03-30",
		Format:  TimeFormat20060102150405,
		Extra:   TimeStart,
		Zone:    time.Local,
	}

	t2 := TimeDetail{
		TimeStr: "2020-03-30",
		Format:  TimeFormat20060102150405,
		Extra:   TimeEnd,
		Zone:    time.Local,
	}

	res, err := FormatTime(t1, t2)
	if err != nil {
		t.Error(err)
		return
	}
	for _, v := range res {
		t.Logf("result:%v\n", v)
	}

}

// TestMinusInTowTime ...
func TestMinusInTowTime(t *testing.T) {

	type MyTime struct {
		Start    string
		End      string
		Expected int
	}
	testCases := []MyTime{

		{
			Start:    "2020-02-01",
			End:      "2020-02-01",
			Expected: 1,
		},

		{
			Start:    "2020-02-03",
			End:      "2020-02-05",
			Expected: 3,
		},

		{
			Start:    "2020-02-01",
			End:      "2020-02-29",
			Expected: 29,
		},
	}

	for i, c := range testCases {
		actual, err := SplitDaysInTowDate(c.Start, c.End)
		if err != nil {
			t.Error(err)
			return
		}
		t.Logf("result:%v\n", actual)

		if len(actual) != c.Expected {
			t.Errorf("not match with expected number,expected:%v,got:%v,index:%v", c.Expected, actual, i)
		}
	}

}

// TestCalcTimeBySeconds ...
func TestCalcTimeBySeconds(t *testing.T) {
	type Time struct {
		Duration int
		Expected string
	}
	testCases := []Time{
		{
			Duration: 301,
			Expected: "5分钟1秒",
		},
		{
			Duration: 310,
			Expected: "5分钟10秒",
		},
		{
			Duration: 60*60 + 11,
			Expected: "1小时11秒",
		},
		{
			Duration: 60 * 60 * 5,
			Expected: "5小时",
		},
		{
			Duration: 60 * 60 * 5+4,
			Expected: "5小时4秒",
		},
		{
			Duration: 60*60*24*3,
			Expected: "3天",
		},
		{
			Duration: 60*60*24 + 19,
			Expected: "1天19秒",
		},
		{
			Duration: 60*60*24 + 120,
			Expected: "1天2分钟",
		},
		{
			Duration: 60*60*24 + 121,
			Expected: "1天2分钟1秒",
		},
	}
	for _, v := range testCases {
		actual := CalcSeconds(v.Duration)
		fmt.Println("actual:",actual)
		if actual != v.Expected {
			t.Errorf("not match with expected.expected[%v],got[%v]", v.Expected, actual)
		}
	}

}
