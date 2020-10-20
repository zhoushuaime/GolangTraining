package helper

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// TimeDetail ...
type TimeDetail struct {
	TimeStr string
	Format  string
	Extra   string //额外参数指精确到时分秒 " 00:00:00" 和" 23:59:59"
	Zone    *time.Location
}

const (
	TimeStart = "00:00:00"
	TimeEnd   = "23:59:59"
)

// TimeFormat20060102150405 其他格式参考time包，如 time.RFC3339...
const TimeFormat20060102150405 = "2006-01-02 15:04:05"

const (
	SecondChinese = "秒"
	MinuteChinese = "分钟"
	HourChinese   = "小时"
	DayChinese    = "天"
)

const (
	DurationSecond = 1
	DurationMinute = 60
	DurationHour   = 60 * 60
	DurationDay    = 60 * 60 * 24
)

// supportTime ...
var supportTime = []struct {
	Duration int
	Unit     string
}{
	{Duration: DurationDay, Unit: DayChinese},
	{Duration: DurationHour, Unit: HourChinese},
	{Duration: DurationMinute, Unit: MinuteChinese},
	{Duration: DurationSecond, Unit: SecondChinese},
}

// calcSeconds ...
func CalcSeconds(duration int) (result string) {
	if duration <= 0 {
		return ""
	}

	for _, v := range supportTime {
		if r := duration / v.Duration; r > 0 {
			result += fmt.Sprintf("%v%v", r, v.Unit)
			duration %= v.Duration
		}
	}

	return result

}

// FormatTime 会格式化为 UTC 时间...
func FormatTime(t ...TimeDetail) ([]time.Time, error) {
	result := make([]time.Time, 0)
	for _, v := range t {
		if v.Extra != "" {
			if v.Extra != TimeStart && v.Extra != TimeEnd {
				return nil, fmt.Errorf("unsupport extra params:%v", v.Extra)
			}
			if !strings.HasSuffix(v.TimeStr, v.Extra) {
				v.TimeStr = fmt.Sprintf("%v %v", v.TimeStr, v.Extra) // 注意加空格
			}
		}

		res, err := time.ParseInLocation(v.Format, v.TimeStr, v.Zone)
		if err != nil {
			return nil, err
		}
		result = append(result, res)
	}
	return result, nil
}

// FormatStartEndTime 格式化开始事件和结束事件，适合时间范围查询...
func FormatStartEndTime(startStr, endStr string, zone *time.Location) (start, end time.Time, err error) {

	startDate := strings.Split(startStr, "-")
	if len(startDate) < 0 {
		err = errors.New("split start date error,invalid format")
		return
	}
	endDate := strings.Split(endStr, "-")
	if len(startDate) < 0 {
		err = errors.New("split end date error,invalid format")
		return
	}
	year, err1 := strconv.Atoi(startDate[0])
	month, err2 := strconv.Atoi(startDate[1])
	day, err3 := strconv.Atoi(startDate[2])
	if err1 != nil || err2 != nil || err3 != nil {
		err = errors.New("invalid start date format")
		return
	}

	start = time.Date(year, time.Month(month), day, 0, 0, 0, 0, zone)
	year, err1 = strconv.Atoi(endDate[0])
	month, err2 = strconv.Atoi(endDate[1])
	day, err3 = strconv.Atoi(endDate[2])
	if err1 != nil || err2 != nil || err3 != nil {
		err = errors.New("invalid end date format")
		return
	}

	end = time.Date(year, time.Month(month), day, 23, 59, 59, 99999, zone)

	return

}

// SplitDaysInTowDate 计算两个日期间有多少个天数,并且按天拆分，。。。
// "2020-12-23"，"2020-12-25" 这里计算出的结果为2020-12-23,2020-12-24,2020-12-25
// "2020-12-23" "2020-12-23" 这里计算出的结果为2020-12-23
func SplitDaysInTowDate(startStr, endStr string) ([]string, error) {
	if startStr == "" || endStr == "" {
		return nil, errors.New("<CalcDaysInTowDate> input date string is empty")
	}
	if strings.HasSuffix(startStr, TimeStart) || strings.HasSuffix(startStr, TimeEnd) {
		startStr = strings.TrimSuffix(startStr, TimeStart)
		startStr = strings.TrimSuffix(startStr, TimeEnd)
	}
	if strings.HasSuffix(endStr, TimeStart) || strings.HasSuffix(endStr, TimeEnd) {
		endStr = strings.TrimSuffix(endStr, TimeStart)
		endStr = strings.TrimSuffix(endStr, TimeEnd)
	}
	if startStr == endStr {
		return []string{startStr}, nil
	}

	t1, err := time.Parse("2006-01-02", startStr)
	if err != nil {
		return nil, err
	}

	t2, err := time.Parse("2006-01-02", endStr)
	if err != nil {
		return nil, err
	}

	dayCount := int(t2.Sub(t1).Hours()/24) + 1
	var days []string
	for i := 0; i < dayCount; i++ {
		days = append(days, t1.AddDate(0, 0, i).Format("2006-01-02"))
	}
	return days, nil

}

// CalcCostTime ...
//func CalcCostTime(start time.Time) int64 {
//	end := time.Now()
//	return (end.Sub(start).Nanoseconds()) / 1e6
//}

// CalcCostTime ...
func CalcCostTime() func() int64 {
	now := time.Now()
	return func() int64 {
		return int64(time.Since(now).Nanoseconds() / 1e6)
	}
}

// CalcCostTimeByName ...
func CalcCostTimeByName(name string) func() {
	//now := time.Now()
	return func() {
		//fmt.Printf("call -> %s cost time:%v\n", name, time.Since(now).Milliseconds())
	}
}
