// package calculation

package main

import (
	"fmt"
	"time"
)

type Task struct {
	Task_Name       string
	Task_Start_Time time.Time
}

func main() {
	// 方法1：使用time.Date（推荐）
	customTime1 := time.Date(2024, 3, 20, 14, 30, 0, 0, time.Local)
	fmt.Printf("自定义时间1: %s\n", customTime1.Format("15:04"))

	// 方法2：使用time.Parse
	customTime2, _ := time.Parse("15:04", "14:30")
	fmt.Printf("自定义时间2: %s\n", customTime2.Format("15:04"))

	// 方法3：只关注时分，忽略日期
	layout := "15:04"
	timeStr := "14:30"
	customTime3, _ := time.Parse(layout, timeStr)
	fmt.Printf("自定义时间3: %s\n", customTime3.Format("15:04"))
}
