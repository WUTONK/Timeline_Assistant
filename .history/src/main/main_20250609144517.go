package main

import calculation

// main 函数：程序入口
// 创建任务链表并添加示例任务
func main() {
	TaskList := &TaskList{TaskHead: nil}

	AddTask(TaskList, "坠机", "00:10", "00:30")
	AddTask(TaskList, "MAN", "00:10", "00:40")
	AddTask(TaskList, "打德佬", "12:00", "13:50")
	AddTask(TaskList, "OFN启动", "12:00", "13:50")
	AddTask(TaskList, "黑联启动", "12:00", "13:50")

	DisplayTaskTimeLine(*TaskList, "00:00", 10, 0, "15:04")
	// // 打印从00:00开始，每10分钟一个时间点，打印6个时间点
	// printTimeSequence("00:00", 10, 10)

	// // 也可以打印其他时间序列
	// // 比如从14:30开始，每15分钟一个时间点，打印4个时间点
	// fmt.Println("\n另一个时间序列：")
	// printTimeSequence("14:30", 15, 4)
	// PrintTimeList("00:00", 10, 10, "15:04")
}