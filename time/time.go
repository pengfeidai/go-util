package time

import "time"

/*
 * 获取当前时间 ”YYYY-MM-DD hh:mm:ss“
 */
func GetDate() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

/*
 * Unix时间戳
 */
func GetUnix() int64 {
	return time.Now().Unix()
}

/*
 * 毫秒级时间戳
 */
func GetMilliUnix() int64 {
	return time.Now().UnixNano() / 1000000
}

/*
 * 纳秒级时间戳
 */
func GetNanoUnix() int64 {
	return time.Now().UnixNano()
}
