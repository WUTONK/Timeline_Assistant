package io

import (
	"fmt"
	"time"
)

// TaskStatus 表示任务的状态
type TaskStatus struct {
	TaskName  string    // 任务名称
	StartTime time.Time // 开始时间
	EndTime   time.Time // 结束时间
	Status    string    // 状态（开始/结束）
}

// Timeline 表示时间轴
type Timeline struct {
	Events map[time.Time]*TaskStatus // 时间点对应的事件
	Tasks  map[string]*TaskStatus    // 任务名称对应的状态
}

// NewTimeline 创建一个新的时间轴
func NewTimeline() *Timeline {
	return &Timeline{
		Events: make(map[time.Time]*TaskStatus),
		Tasks:  make(map[string]*TaskStatus),
	}
}

// AddTaskEvent 添加任务事件
func (t *Timeline) AddTaskEvent(taskName string, eventTime time.Time, status string) {
	event := &TaskStatus{
		TaskName:  taskName,
		Status:    status,
		StartTime: eventTime,
	}

	if status == "start" {
		t.Tasks[taskName] = event
	} else if status == "end" {
		if task, exists := t.Tasks[taskName]; exists {
			task.EndTime = eventTime
		}
	}

	t.Events[eventTime] = event
}

// GetTaskDuration 获取任务持续时间
func (t *Timeline) GetTaskDuration(taskName string) time.Duration {
	if task, exists := t.Tasks[taskName]; exists && !task.EndTime.IsZero() {
		return task.EndTime.Sub(task.StartTime)
	}
	return 0
}

// PrintTimeline 打印时间轴
func (t *Timeline) PrintTimeline() {
	fmt.Println("时间轴统计：")
	for time, event := range t.Events {
		fmt.Printf("%s | %s: %s\n", time.Format("15:04"), event.TaskName, event.Status)
	}
}

// PrintTaskSummary 打印任务统计
func (t *Timeline) PrintTaskSummary() {
	fmt.Println("\n任务统计：")
	for taskName, task := range t.Tasks {
		if !task.EndTime.IsZero() {
			duration := task.EndTime.Sub(task.StartTime)
			fmt.Printf("<%s> %dh%dm\n", taskName, int(duration.Hours()), int(duration.Minutes())%60)
		}
	}
}

func Initialization_Schedule() {
	timeline := NewTimeline()

	// 示例使用
	now := time.Now()
	timeline.AddTaskEvent("打搅", now, "start")
	timeline.AddTaskEvent("睡觉", now.Add(10*time.Minute), "start")
	timeline.AddTaskEvent("打搅", now.Add(30*time.Minute), "end")
	timeline.AddTaskEvent("睡觉", now.Add(30*time.Minute), "end")

	timeline.PrintTimeline()
	timeline.PrintTaskSummary()
}
