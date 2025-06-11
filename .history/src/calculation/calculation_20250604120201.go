// package calculation

package main

import (
	"time"
)

type Task struct {
	Task_Name       string
	Task_Start_Time time.Time
}

func PrintTimeList(StartTime string, TimeInterval int, count int, TimeFormat string) {
	CurrentTime, _ := time.Parse(TimeFormat, StartTime)

}

func main() {
	// // 打印从00:00开始，每10分钟一个时间点，打印6个时间点
	// printTimeSequence("00:00", 10, 10)

	// // 也可以打印其他时间序列
	// // 比如从14:30开始，每15分钟一个时间点，打印4个时间点
	// fmt.Println("\n另一个时间序列：")
	// printTimeSequence("14:30", 15, 4)
}
