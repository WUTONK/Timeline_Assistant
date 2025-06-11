// package calculation

package main

import (
	"fmt"
	"time"
)

// 单链表结构
type Task struct {
	Task_Name       string
	Task_Start_Time time.Time
	Task_End_Time   time.Time
	Next_Task       *Task
}

type TaskList struct {
	TaskHead *Task
}

func TimeStandardParser(TimeToParser string) time.Time {
	AfterPeaserTime, _ := time.Parse("11:45", TimeToParser)
	return AfterPeaserTime
}

func AddTask(t TaskList, TaskName string, TaskStartTime string, TaskEndTime string) {
	AfterPeaserTaskStartTime := TimeStandardParser(TaskStartTime)
	AfterPeaserTaskEndTime := TimeStandardParser(TaskEndTime)

	NewTask := &Task{Task_Name: TaskName, Task_Start_Time: AfterPeaserTaskStartTime, Task_End_Time: AfterPeaserTaskEndTime}

	if t.TaskHead == nil {
		t.TaskHead = NewTask
	}

}

func DeleteLasrTask(t TaskList) {
	if t.TaskHead == nil {
		fmt.Println("队列中没有任务")
		return
	}

	if t.TaskHead.Next_Task == nil {
		t.TaskHead = nil
		fmt.Println("已删除头任务")
		return
	}

	CurrentTask := t.TaskHead
	FirstTask := &Task{}

	for {
		if CurrentTask.Next_Task == nil {
			FirstTask.Next_Task = nil
			return
		}

		FirstTask = CurrentTask
		CurrentTask = CurrentTask.Next_Task
	}

}

func TraversalTask() {

}

func PrintTimeList(StartTime string, TimeInterval int, count int, TimeFormat string) {
	CurrentTime, _ := time.Parse(TimeFormat, StartTime)

	for i := range count {
		CurrentTime = CurrentTime.Add(time.Duration(i*TimeInterval) * time.Minute)
		fmt.Printf("%s|\n", CurrentTime.Format(TimeFormat))
	}
}

func isStartAndEndTimeHit(TaskStartTime time.Time, TaskEndTime time.Time, TimeLineStartTime string, TimeInterval int, TimeFormat string) bool {

	// 是否命中开始/结束时间
	isStartHit := false
	isEndHit := false
	// 要检测任务结束/开始时间是否命中跳点。命中后再输出
	CurrentTime, _ := time.Parse(TimeFormat, TimeLineStartTime)
	CurrentTimeCopy := CurrentTime
	// 怎么知道已经过了一轮：累计加时大于等于1440min(24小时)
	TimeAddSumMinute := 0

	for i := 0; ; {
		// 加一次并计算总分钟
		i++
		CurrentTimeCopy = CurrentTime.Add(time.Duration(i*TimeInterval) * time.Minute)
		TimeAddSumMinute += i * TimeInterval

		// 检测是否命中开始/结束时间
		if TaskStartTime == CurrentTimeCopy {
			isStartHit = true
		}
		if TaskEndTime == CurrentTimeCopy {
			isEndHit = true
		}

		// 检测是否经过一轮（24h）
		if TimeAddSumMinute >= 1440 {
			if isStartHit && isEndHit {
				return true
			} else {
				fmt.Println("开始/结束未命中，无法显示")
				return false
			}
		}
	}

}

func DisplayTaskTimeLine(t TaskList, StartTime string, TimeInterval int, count int, TimeFormat string) {
	tasks := make([]*Task, 0, 10)

	for CurrentTask := t.TaskHead; CurrentTask != nil; {
		tasks = append(tasks, CurrentTask)
		CurrentTask = CurrentTask.Next_Task
	}

	for i, task := range tasks {
		fmt.Printf("任务 %d: %s\n", i, task.Task_Name)
	}
}

func main() {
	TaskLise := TaskList{TaskHead: nil}

	AddTask(TaskLise, "坠机", "00:10", "00:30")
	AddTask(TaskLise, "MAN", "00:10", "00:40")

	DisplayTaskTimeLine(TaskLise, "0", 0, 0, "0")
	// // 打印从00:00开始，每10分钟一个时间点，打印6个时间点
	// printTimeSequence("00:00", 10, 10)

	// // 也可以打印其他时间序列
	// // 比如从14:30开始，每15分钟一个时间点，打印4个时间点
	// fmt.Println("\n另一个时间序列：")
	// printTimeSequence("14:30", 15, 4)
	// PrintTimeList("00:00", 10, 10, "15:04")
}
