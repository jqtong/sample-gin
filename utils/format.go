package utils

import (
	"fmt"
	"nucarf.com/store_service/api/conf/initialize"
	"strconv"
	"strings"
	"time"
)

// PrettyNum transfer numbers to pretty format
func PrettyNum(num int) string {

	if num > 10000 {
		return fmt.Sprintf("%.2f万", float64(num)/10000)
	}

	return strconv.Itoa(num)
}

//ParseTime2Str format time to str
func ParseTime2Str(strTime time.Time) string {
	theTime := strTime.Format("2006-01-02 15:04:05")
	if theTime == "0001-01-01 00:00:00" {
		theTime = "0000-00-00 00:00:00"
	}
	return theTime
}

// ParseStr2Time 时间字符串转时间格式
func ParseStr2Time(str string, format string) time.Time {
	str = strings.Replace(str, "T", " ", 1)
	str = strings.Replace(str, "Z", " ", 1)
	str = strings.Replace(str, "+08:00", " ", 1)
	str = strings.Replace(str, "+0800", " ", 1)
	str = strings.TrimSpace(str)
	var timeLayout string
	switch format {
	case "Y-m-d":
		timeLayout = "2006-01-02"
	case "Y-m-d H:i:s":
		timeLayout = "2006-01-02 15:04:05"
	case "Y-m-d H":
		timeLayout = "2006-01-02 15"
	case "Y-m-d H:i":
		timeLayout = "2006-01-02 15:04"
	default:
		timeLayout = "2006-01-02 15:04:05"
	}
	loc, _ := time.LoadLocation("Asia/Shanghai")
	theTime, _ := time.ParseInLocation(timeLayout, str, loc)
	return theTime
}

// FormatTime 格式化时间
func FormatTime(str string, oldFormat, newFormat string) string {
	str = strings.Replace(str, "T", " ", 1)
	str = strings.Replace(str, "Z", " ", 1)
	str = strings.Replace(str, "+08:00", " ", 1)
	str = strings.Replace(str, "+0800", " ", 1)
	str = strings.TrimSpace(str)
	if str == "0001-01-01 00:00:00" || str == "0000-00-00 00:00:00" {
		return ""
	}
	oldTimeLayout := getTimeLayout(oldFormat)
	newTimeLayout := getTimeLayout(newFormat)
	if str == "" {
		return time.Now().Format(newTimeLayout)
	}

	loc, _ := time.LoadLocation("Asia/Shanghai")
	theTime, _ := time.ParseInLocation(oldTimeLayout, str, loc)

	return theTime.Format(newTimeLayout)
}

func getTimeLayout(format string) string {
	switch format {
	case "Y-m-d":
		return "2006-01-02"
	case "Y-m-d H:i:s":
		return "2006-01-02 15:04:05"
	case "Y-m-d H":
		return "2006-01-02 15"
	case "Y-m-d H:i":
		return "2006-01-02 15:04"
	case "Y-m-dTH":
		return "2006-01-02T15"
	case "H:i:s":
		return "15:04:05"
	case "H:i":
		return "15:04"
	default:
		return "2006-01-02 15:04:05"
	}
}

//GetTwoFloat 获取两位小数
func GetTwoFloat(f float64, precision int) float64 {
	rest, _ := strconv.ParseFloat(fmt.Sprintf("%."+strconv.Itoa(precision)+"f", f), 64)
	return rest
}

//TimeNotNil 时间非空
func TimeNotNil(time time.Time) bool {
	return ParseTime2Str(time) != "0000-00-00 00:00:00"
}

//ParsetoString parse float to string
func ParsetoString(str interface{}) (res string) {
	switch str.(type) {
	case float64:
		res = fmt.Sprintf("%v", str.(float64))
	case string:
		res = str.(string)
	}
	return
}

//FormatYearMonthToYearMonthDay 将 start=2021-08T00 时间 转化为2021-08-01T00 end=2021-08T00 时间 转化为2021-08-31T00
func FormatYearMonthToYearMonthDay(start, end string) (string, string) {
	startFormat := start
	endFormat := end
	loc, _ := time.LoadLocation(initialize.ServerConf.Timezone)

	if len(start[:strings.Index(start, "T")]) == 7 {
		startTimeNow, _ := time.ParseInLocation("2006-01 15", strings.Replace(start, "T", " ", 1), loc)
		currentYear, currentMonth, _ := startTimeNow.Date()

		firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, loc)
		startFormat = firstOfMonth.Format("2006-01-02") + start[7:]
	}

	if len(end[:strings.Index(end, "T")]) == 7 {

		endTimeNow, _ := time.ParseInLocation("2006-01 15", strings.Replace(end, "T", " ", 1), loc)
		currentYear, currentMonth, _ := endTimeNow.Date()

		firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, loc)
		lastOfMonth := firstOfMonth.AddDate(0, 1, -1)
		endFormat = lastOfMonth.Format("2006-01-02") + end[7:]
	}

	return startFormat, endFormat
}
