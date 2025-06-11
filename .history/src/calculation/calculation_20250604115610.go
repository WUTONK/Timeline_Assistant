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

func PrintTimeList(startTime string, TimeInterval int, TimeRound int, TimeFormat string) {
	// 使用标准的时间格式 "15:04" 来解析输入的时间字符串
	Start, _ := time.Parse("15:04", startTime)

	for i := 0; i < TimeRound; i++ {
		CurrentTime := Start.Add(time.Duration(i*TimeInterval) * time.Minute)
		// 使用传入的 TimeFormat 来格式化输出时间
		fmt.Printf("%v|\n", CurrentTime.Format(TimeFormat))
	}
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
	// // 打印从00:00开始，每10分钟一个时间点，打印6个时间点
	// printTimeSequence("00:00", 10, 10)

	// // 也可以打印其他时间序列
	// // 比如从14:30开始，每15分钟一个时间点，打印4个时间点
	// fmt.Println("\n另一个时间序列：")
	// printTimeSequence("14:30", 15, 4)

	// 使用标准的时间格式 "15:04" 作为输出格式
	PrintTimeList("00:20", 10, 30, "15:04")
}
