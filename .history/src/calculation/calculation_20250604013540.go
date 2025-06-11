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

// 打印时间序列
func printTimeSequence(startTime string, intervalMinutes int, count int) {
	// 解析开始时间
	start, _ := time.Parse("15:04", startTime)

	// 打印时间序列
	for i := 0; i < count; i++ {
		currentTime := start.Add(time.Duration(i*intervalMinutes) * time.Minute)
		fmt.Printf("%s｜\n", currentTime.Format("15:04"))
	}
}

func main() {
	// 创建两个时间点
	time1, _ := time.Parse("15:04", "14:30")
	time2, _ := time.Parse("15:04", "15:30")

	// 比较时间
	fmt.Println("时间比较示例：")
	fmt.Printf("time1: %s\n", time1.Format("15:04"))
	fmt.Printf("time2: %s\n", time2.Format("15:04"))

	// 使用比较运算符
	fmt.Printf("time1 < time2: %v\n", time1.Before(time2)) // 或者 time1 < time2
	fmt.Printf("time1 > time2: %v\n", time1.After(time2))  // 或者 time1 > time2
	fmt.Printf("time1 == time2: %v\n", time1.Equal(time2)) // 或者 time1 == time2

	// 计算时间差
	diff := time2.Sub(time1)
	fmt.Printf("时间差: %v\n", diff) // 输出: 1h0m0s
}
