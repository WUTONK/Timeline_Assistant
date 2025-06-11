package main

import (
	"Timeline_Assistant/src/calculation"
	timeIO "Timeline_Assistant/src/io"
	"fmt"
)

// main 函数：程序入口
// 创建任务链表并添加示例任务
func main() {
	taskList := &calculation.TaskList{Head: nil}

	taskList.AddTask("坠机", "00:10", "00:30")
	taskList.AddTask("MAN", "00:10", "00:40")
	taskList.AddTask("打德佬", "12:00", "13:50")
	taskList.AddTask("OFN启动", "12:00", "13:50")
	taskList.AddTask("黑联启动", "12:00", "13:50")

	content := taskList.TaskTimeLine("00:00", 10, 0, "15:04")
	fmt.Println("--------------------")
	timeIO.WriteFile("./file.txt", content)

	// // 打印从00:00开始，每10分钟一个时间点，打印6个时间点
	// printTimeSequence("00:00", 10, 10)

	// // 也可以打印其他时间序列
	// // 比如从14:30开始，每15分钟一个时间点，打印4个时间点
	// fmt.Println("\n另一个时间序列：")
	// printTimeSequence("14:30", 15, 4)
	// PrintTimeList("00:00", 10, 10, "15:04")
}
