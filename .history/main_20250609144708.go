package main

import (
	time_io "Timeline_Assistant/src/io"
)

func main() {
	time_io.Initialization_Schedule()
	// 这里你可以使用 calculation 包中的功能
	// 例如：
	// taskList := &calculation.TaskList{TaskHead: nil}
	// calculation.AddTask(taskList, "任务名称", "00:00", "01:00")
}
