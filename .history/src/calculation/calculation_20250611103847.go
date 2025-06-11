package calculation

import (
	"fmt"
	"time"
)

// Task 结构体定义了一个任务的基本信息
// Task_Name: 任务名称
// Task_Start_Time: 任务开始时间
// Task_End_Time: 任务结束时间
// Next_Task: 指向下一个任务的指针，用于构建链表
type Task struct {
	Task_Name       string
	Task_Start_Time time.Time
	Task_End_Time   time.Time
	Next_Task       *Task
}

// TaskList 结构体表示任务链表
// TaskHead: 指向链表头节点的指针
type TaskList struct {
	TaskHead *Task
}

// TimeStandardParser 将时间字符串解析为 time.Time 类型
// 参数 TimeToParser: 要解析的时间字符串，格式为 "HH:MM"
// 返回: 解析后的 time.Time 对象
func TimeStandardParser(TimeToParser string) time.Time {
	AfterPeaserTime, _ := time.Parse("15:04", TimeToParser)
	return AfterPeaserTime
}

// AddTask 向任务链表中添加新任务
// 参数 t: 任务链表的指针
// 参数 TaskName: 任务名称
// 参数 TaskStartTime: 任务开始时间（字符串格式 "HH:MM"）
// 参数 TaskEndTime: 任务结束时间（字符串格式 "HH:MM"）
func AddTask(t *TaskList, TaskName string, TaskStartTime string, TaskEndTime string) {
	AfterPeaserTaskStartTime := TimeStandardParser(TaskStartTime)
	AfterPeaserTaskEndTime := TimeStandardParser(TaskEndTime)

	NewTask := &Task{Task_Name: TaskName, Task_Start_Time: AfterPeaserTaskStartTime, Task_End_Time: AfterPeaserTaskEndTime}

	if t.TaskHead == nil {
		t.TaskHead = NewTask
		return
	}

	CurrencyTask := t.TaskHead
	for CurrencyTask.Next_Task != nil {
		CurrencyTask = CurrencyTask.Next_Task
	}

	CurrencyTask.Next_Task = NewTask
}

// DeleteLasrTask 删除任务链表中的最后一个任务
// 参数 t: 任务链表
// 注意：函数名中的 "Lasr" 应该是 "Last" 的拼写错误
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

// TraversalTask 遍历任务链表（待实现）
func TraversalTask() {

}

// PrintTimeList 打印时间序列
// 参数 StartTime: 开始时间（字符串格式）
// 参数 TimeInterval: 时间间隔（分钟）
// 参数 count: 要打印的时间点数量
// 参数 TimeFormat: 时间格式
func PrintTimeList(StartTime string, TimeInterval int, count int, TimeFormat string) {
	CurrentTime, _ := time.Parse(TimeFormat, StartTime)

	for i := range count {
		CurrentTime = CurrentTime.Add(time.Duration(i*TimeInterval) * time.Minute)
		fmt.Printf("%s|\n", CurrentTime.Format(TimeFormat))
	}
}

// DisplayTaskTimeLine 显示任务时间线
// 参数 t: 任务链表
// 参数 StartTime: 开始时间（未使用）
// 参数 TimeInterval: 时间间隔（未使用）
// 参数 count: 计数（未使用）
// 参数 TimeFormat: 时间格式
func DisplayTaskTimeLine(t TaskList, TimeLineStartTime string, TimeInterval int, count int, TimeFormat string) []string {
	// 计算需要的时间点数量（24小时 = 1440分钟）
	// 加1是因为我们需要包含起始时间点
	timePoints := (1440 / TimeInterval) + 1
	Tasks := make([]*Task, 0, 10)               // 任务列表，使用容量而不是长度
	InfoRightList := make([]string, timePoints) // 右侧信息列表，预分配足够的空间
	InfoLeftList := make([]string, timePoints)  // 左侧信息列表，预分配足够的空间

	// 用来写入文件的字符串map
	StringSpliceToWrite := make([]string, 0, timePoints)

	// 检测任务的时间和时间轴尺度是否匹配，如果是则将任务存储到任务列表切片中
	for CurrentTask := t.TaskHead; CurrentTask != nil; {
		Hit := DisplayTaskTimeLine_isStartAndEndTimeHit(TimeLineStartTime, CurrentTask.Task_Start_Time, CurrentTask.Task_End_Time, TimeInterval, TimeFormat)
		if Hit {
			Tasks = append(Tasks, CurrentTask)
		}
		CurrentTask = CurrentTask.Next_Task
	}

	StartTime, _ := time.Parse(TimeFormat, TimeLineStartTime)

	for i := 0; i < timePoints; i++ {
		CurrencyTime := StartTime.Add(time.Duration(i*TimeInterval) * time.Minute)

		// 在每一个时间点遍历一次所有任务，将开始/结束时间匹配的任务添加到对应时间点的右侧显示列表中
		for TaskIndex, Task := range Tasks {
			if Task != nil { // 空值安全检查

				CurrentInfoRight := DisplayTaskTimeLine_Right(TaskIndex, Task, CurrencyTime)
				if CurrentInfoRight != "nil" {
					InfoRightList[i] += CurrentInfoRight // 将任务信息加入右侧信息队列
				}
			} else {
				fmt.Println("任务列表中没有任务与时间轴尺度匹配")
			}
		}

		//左侧时间部分
		InfoLeftList[i] = CurrencyTime.Format(TimeFormat) + "|"
	}

	// 遍历拼接左右部分
	for i := range InfoLeftList {
		Line := fmt.Sprintf("%s%s\n", InfoLeftList[i], InfoRightList[i])
		fmt.Print(Line)
		StringSpliceToWrite = append(StringSpliceToWrite, Line)
	}

	// 添加任务汇总部分
	for i, Task := range Tasks {
		Line := DisplayTaskTimeLine_DisplayTaskInfo(Task, TimeInterval, i, TimeFormat)
		fmt.Print(Line)
		StringSpliceToWrite = append(StringSpliceToWrite, Line)
	}

	return StringSpliceToWrite
}

// isStartAndEndTimeHit 检查任务的开始和结束时间是否命中时间线
// 参数 TaskStartTime: 任务开始时间
// 参数 TaskEndTime: 任务结束时间
// 参数 TimeLineStartTime: 时间线开始时间
// 参数 TimeInterval: 时间间隔（分钟）
// 参数 TimeFormat: 时间格式
// 返回: 如果开始和结束时间都命中则返回 true，否则返回 false
func DisplayTaskTimeLine_isStartAndEndTimeHit(TimeLineStartTime string, TaskStartTime time.Time, TaskEndTime time.Time, TimeInterval int, TimeFormat string) bool {

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
		TimeAddSumMinute = i * TimeInterval

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

// 显示右侧信息
func DisplayTaskTimeLine_Right(TaskIndex int, Task *Task, CurrencyTime time.Time) string {
	TaskStatus := "DontNeedDisplay"

	if CurrencyTime == Task.Task_Start_Time {
		TaskStatus = "开始"
	} else if CurrencyTime == Task.Task_End_Time {
		TaskStatus = "结束"
	}

	if TaskStatus != "DontNeedDisplay" {
		InfoRight := fmt.Sprintf("<--任务%v%s%s   ", TaskIndex, Task.Task_Name, TaskStatus)
		return InfoRight
	}
	return "nil"
}

// DisplayTaskTimeLine_DisplayTaskInfo 显示任务汇总信息
// 参数 Task: 任务对象
// 参数 TimeInterval: 时间间隔
// 参数 TaskIndex: 任务索引
// 参数 TimeFormat: 时间格式
// 返回: 格式化后的任务汇总信息字符串
func DisplayTaskTimeLine_DisplayTaskInfo(Task *Task, TimeInterval int, TaskIndex int, TimeFormat string) string {
	if Task == nil {
		return ""
	}
	return fmt.Sprintf("任务%d: %s (开始时间: %s, 结束时间: %s)\n",
		TaskIndex,
		Task.Task_Name,
		Task.Task_Start_Time.Format(TimeFormat),
		Task.Task_End_Time.Format(TimeFormat))
}
