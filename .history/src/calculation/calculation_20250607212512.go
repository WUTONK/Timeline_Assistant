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

func TimeSumMinuteCalculate(t time.Time) int {
	SumMinute := t.Hour()*60 + t.Minute()
	return SumMinute
}

func isStartAndEndTimeHit(TaskStartTime time.Time, TaskEndTime time.Time, StartTime string, TimeLineStartTime string, TimeInterval int, count int, TimeFormat string) {

	// 是否命中开始/结束时间
	StartTime := TaskStartTime
	EndTime := TaskEndTime
	isStartHit := false
	isEndHit := false

	// 要检测任务结束/开始时间是否命中跳点。命中后再输出
	// 如果有多个任务开始/结束，要依次显示出来
	CurrentTime, _ := time.Parse(TimeFormat, StartTime)
	CurrentTimeCopy := CurrentTime
	// 怎么知道已经过了一轮：累计加时大于等于1440min(24小时)
	TimeAddSumMinute := 0
	CurrentTimeSumMinute := TimeSumMinuteCalculate(CurrentTime)

	for i := 0; ; {

		// 加一轮并计算总分钟
		i++
		CurrentTimeCopy = CurrentTime.Add(time.Duration(i*TimeInterval) * time.Minute)
		TimeAddSumMinute += i * TimeInterval

		// 检测是否经过一轮
		if TimeAddSumMinute >= 1440 {
			break
		}

		// 检测是否经过一轮

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
