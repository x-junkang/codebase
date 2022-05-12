package main

import (
	"fmt"
	"time"
)

func main() {
	// 获取每天的零点时间戳, 一个小时的时间戳是3600
	timeStr := "2022-04-01 10:22:59"
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr, time.Local)
	timeUnix := t.Unix()
	fmt.Println(timeUnix, timeStr)
}
