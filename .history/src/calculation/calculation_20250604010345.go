// package calculation

package main

import "time"

// TaskEvent 表示任务事件
type TaskEvent struct {
	IsStart   bool      // 是否是开始事件
	TaskName  string    // 任务名称
	EventTime time.Time // 事件时间
}

// type Task struct{
// 	Task_Name string
// 	Task_Start_Time time.Time.
// }

func main() {
	var TaskMap []TaskEvent
	TaskMap = make([]TaskEvent, 0)

	// 使用示例
	TaskMap = append(TaskMap, TaskEvent{
		IsStart:   true,
		TaskName:  "睡觉",
		EventTime: time.Now(),
	})
}
