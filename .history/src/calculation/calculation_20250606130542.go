// package calculation

package main

import (
	"fmt"
	"time"
)

type Task struct {
	Task_Name       string
	Task_Start_Time time.Time
	Task_End_Time   time.Time
}

func PrintTimeList(StartTime string, TimeInterval int, count int, TimeFormat string) {
	CurrentTime, _ := time.Parse(TimeFormat, StartTime)

	for i := range count {
		CurrentTime = CurrentTime.Add(time.Duration(i*TimeInterval) * time.Minute)
		fmt.Printf("%s|\n", CurrentTime.Format(TimeFormat))
	}
}

func PrintTimeList_new(StartTime string, TimeInterval int, count int, TimeFormat string) {
	// 要检测任务结束/开始时间是否命中跳点。命中后再输出
	// 如果有多个任务开始/结束，要依次显示出来
	CurrentTime, _ := time.Parse(TimeFormat, StartTime)

	CurrentTimeCopyOne := CurrentTime
	CurrentTimeCopyTwo := CurrentTime
	// 怎么知道已经过了一轮：如果时间小于起始时间后，再次等于或大于起始时间，那么我们就可以判断已经过了一轮

	i := 0
	TimeHasCycle := false;
	TimeTwoSumMinute = CurrentTimeCopyTwo.Hour() * CurrentTimeCopyTwo.Minute()
	for !TimeHasCycle {
		i++
		CurrentTimeCopyOne = CurrentTimeCopyOne.Add(time.Duration(i*TimeInterval)*time.Minute)
		
		TimeOneSumMinute = CurrentTimeCopyOne.Hour() * CurrentTime.Minute()
		if TimeOneSumMinute < TimeTow
		
	}

	for i := range count {
		CurrentTime = CurrentTime.Add(time.Duration(i*TimeInterval) * time.Minute)
		fmt.Printf("%s|\n", CurrentTime.Format(TimeFormat))
	}
}

func main() {

	// // 打印从00:00开始，每10分钟一个时间点，打印6个时间点
	// printTimeSequence("00:00", 10, 10)

	// // 也可以打印其他时间序列
	// // 比如从14:30开始，每15分钟一个时间点，打印4个时间点
	// fmt.Println("\n另一个时间序列：")
	// printTimeSequence("14:30", 15, 4)
	PrintTimeList("00:00", 10, 10, "15:04")
}
