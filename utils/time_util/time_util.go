package time_util

import "time"

var (
	monthDay = map[int]int{
		1:  31,
		2:  28,
		3:  31,
		4:  30,
		5:  31,
		6:  30,
		7:  31,
		8:  31,
		9:  30,
		10: 31,
		11: 30,
		12: 31,
		13: 29, //表示润年二月
	}
)

//是否润年
func IsLeapYear(year int) bool {
	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		return true
	}
	return false
}

//是否每个月最后一天
func IsMonthLastDay(year, month, day int) bool {
	if IsLeapYear(year) && month == 2 {
		month = 13
	}

	if monthDay[month] == day {
		return true
	}

	return false
}

//
func MonthLastDay(year, month int) int {
	if IsLeapYear(year) && month == 2 {
		month = 13
	}
	return monthDay[month]
}

//判断时间戳是否是属于今天，秒级时间戳
func IsToday(ts int64) bool {

	y, m, d := time.Unix(ts, 0).Date()
	y1, m1, d1 := time.Now().Date()

	if d != d1 || m != m1 || y != y1{
		return false
	}

	return true
}

//获取今天0点时间戳
func GetTodayZeroTs() int64 {
	y, m, d := time.Now().Date()
	ts := time.Date(y, m, d, 0, 0, 0, 0, time.Local).Unix()
	return ts
}

//获取次日0点时间戳
func GetNextDayZeroTs() int64 {
	y, m, d := time.Now().Date()
	ts := time.Date(y, m, d, 0, 0, 0, 0, time.Local).
		AddDate(0, 0, 1).Unix()
	return ts
}

// 时间字符串转为本地时间
func ParseTimeInLocal(timeStr string) (time.Time, error) {
	return time.ParseInLocation("2006-01-02 15:04:05", timeStr, time.Local)
}
