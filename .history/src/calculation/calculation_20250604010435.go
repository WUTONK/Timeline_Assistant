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
	// 方法1：使用Format方法
	now := time.Now()
	formattedTime := now.Format("15:04") // 输出类似 "14:30"
	fmt.Printf("当前时间: %s\n", formattedTime)

	// 方法2：自定义格式
	customTime := fmt.Sprintf("%02d(h):%02d", now.Hour(), now.Minute())
	fmt.Printf("自定义格式: %s\n", customTime)

	// 方法3：使用time.Duration格式化
	duration := 90 * time.Minute
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	durationStr := fmt.Sprintf("%02d(h):%02d", hours, minutes)
	fmt.Printf("持续时间: %s\n", durationStr)
}
