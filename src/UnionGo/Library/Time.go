package Library

import (
	"time"
)
const TIME_LAYOUT_OFTEN = "2006-01-02 15:04:05"
// 解析常用的日期时间格式：2014-01-11 16:18:00，东八区
func TimeParseOften(value string) (time.Time, error) {
	local, _ := time.LoadLocation("Local")
	return time.ParseInLocation(TIME_LAYOUT_OFTEN, value, local)
}

//返回当前时区的当前时间
func TimeNowString() (timea string) {
	timea=time.Now().Format(TIME_LAYOUT_OFTEN)
	return timea
}
