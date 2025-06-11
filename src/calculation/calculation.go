package calculation

import (
	"encoding/json"
	"fmt"
	"time"
)

// Task 结构体定义了一个任务的基本信息
// Name: 任务名称
// StartTime: 任务开始时间
// EndTime: 任务结束时间
// Next: 指向下一个任务的指针，用于构建链表
type Task struct {
	Name      string
	StartTime time.Time
	EndTime   time.Time
	Next      *Task
}

// TaskList 结构体表示任务链表
// Head: 指向链表头节点的指针
type TaskList struct {
	Head *Task
}

// TimeStandardParser 将时间字符串解析为 time.Time 类型
// 参数 TimeToParser: 要解析的时间字符串，格式为 "HH:MM"
// 返回: 解析后的 time.Time 对象
func TimeStandardParser(TimeToParser string) time.Time {
	afterParserTime, _ := time.Parse("15:04", TimeToParser)
	return afterParserTime
}

// AddTask 向任务链表中添加新任务
// 参数 t: 任务链表的指针
// 参数 TaskName: 任务名称
// 参数 TaskStartTime: 任务开始时间（字符串格式 "HH:MM"）
// 参数 TaskEndTime: 任务结束时间（字符串格式 "HH:MM"）
func (t *TaskList) AddTask(name string, startTime string, endTime string) {
	parsedTaskStartTime := TimeStandardParser(startTime)
	parsedTaskEndTime := TimeStandardParser(endTime)

	newTask := &Task{Name: name, StartTime: parsedTaskStartTime, EndTime: parsedTaskEndTime}

	if t.Head == nil {
		t.Head = newTask
		return
	}

	curTask := t.Head
	for curTask.Next != nil {
		curTask = curTask.Next
	}

	curTask.Next = newTask
}

// DeleteLastTask 删除任务链表中的最后一个任务
// 参数 t: 任务链表
func DeleteLastTask(t TaskList) {
	if t.Head == nil {
		fmt.Println("队列中没有任务")
		return
	}

	if t.Head.Next == nil {
		t.Head = nil
		fmt.Println("已删除头任务")
		return
	}

	curTask := t.Head
	fstTask := &Task{}

	for {
		if curTask.Next == nil {
			fstTask.Next = nil
			return
		}

		fstTask = curTask
		curTask = curTask.Next
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
func PrintTimeList(startTime string, timeInterval int, count int, timeFormat string) {
	curTime, _ := time.Parse(timeFormat, startTime)

	for i := range count {
		curTime = curTime.Add(time.Duration(i*timeInterval) * time.Minute)
		fmt.Printf("%s|\n", curTime.Format(timeFormat))
	}
}

// TaskTimeLine 显示任务时间线
// 参数 t: 任务链表
// 参数 StartTime: 开始时间（未使用）
// 参数 TimeInterval: 时间间隔（未使用）
// 参数 count: 计数（未使用）
// 参数 TimeFormat: 时间格式
func (t *TaskList) TaskTimeLine(timeLineStartTime string, timeInterval int, count int, timeFormat string) []string {
	// 计算需要的时间点数量（24小时 = 1440分钟）
	// 加1是因为我们需要包含起始时间点
	timePoints := (1440 / timeInterval) + 1
	tasks := make([]*Task, 0, 10)               // 任务列表，使用容量而不是长度
	infoRightList := make([]string, timePoints) // 右侧信息列表，预分配足够的空间
	infoLeftList := make([]string, timePoints)  // 左侧信息列表，预分配足够的空间

	// 用来写入文件的字符串map
	strSliceToWrite := make([]string, 0, timePoints)

	// 检测任务的时间和时间轴尺度是否匹配，如果是则将任务存储到任务列表切片中
	for cur := t.Head; cur != nil; {
		hit := TimelineIsStartEndTimeHit(timeLineStartTime, cur.StartTime, cur.EndTime, timeInterval, timeFormat)
		if hit {
			tasks = append(tasks, cur)
		}
		cur = cur.Next
	}

	startTime, _ := time.Parse(timeFormat, timeLineStartTime)

	for i := 0; i < timePoints; i++ {
		curTime := startTime.Add(time.Duration(i*timeInterval) * time.Minute)

		// 在每一个时间点遍历一次所有任务，将开始/结束时间匹配的任务添加到对应时间点的右侧显示列表中
		for taskIndex, task := range tasks {
			if task != nil { // 空值安全检查

				curInfoRight := TimelineRightString(taskIndex, task, curTime)
				if curInfoRight != "nil" {
					infoRightList[i] += curInfoRight // 将任务信息加入右侧信息队列
				}
			} else {
				fmt.Println("任务列表中没有任务与时间轴尺度匹配")
			}
		}

		//左侧时间部分
		infoLeftList[i] = curTime.Format(timeFormat) + "|"
	}

	// 遍历拼接左右部分
	for i := range infoLeftList {
		line := fmt.Sprintf("%s%s\n", infoLeftList[i], infoRightList[i])
		fmt.Print(line)
		strSliceToWrite = append(strSliceToWrite, line)
	}

	// 添加任务汇总部分
	for i, task := range tasks {
		fmt.Println(tasks)
		fmt.Println(task)
		fmt.Println("进入循环")
		Line := task.TimelineTaskInfoString(timeInterval, i, timeFormat)
		// fmt.Print(Line)
		strSliceToWrite = append(strSliceToWrite, Line)
	}

	return strSliceToWrite
}

func Jout(v any) {
	bs, _ := json.Marshal(v)
	fmt.Println(string(bs))
}

// TimelineIsStartEndTimeHit 检查任务的开始和结束时间是否命中时间线
// 参数 TaskStartTime: 任务开始时间
// 参数 TaskEndTime: 任务结束时间
// 参数 TimeLineStartTime: 时间线开始时间
// 参数 TimeInterval: 时间间隔（分钟）
// 参数 TimeFormat: 时间格式
// 返回: 如果开始和结束时间都命中则返回 true，否则返回 false
func TimelineIsStartEndTimeHit(timeLineStartTime string, taskStartTime time.Time, taskEndTime time.Time, timeInterval int, timeFormat string) bool {

	// 是否命中开始/结束时间
	isStartHit := false
	isEndHit := false
	// 要检测任务结束/开始时间是否命中跳点。命中后再输出
	curTime, _ := time.Parse(timeFormat, timeLineStartTime)
	curTimeCopy := curTime
	// 怎么知道已经过了一轮：累计加时大于等于1440min(24小时)
	timeAddSumMinute := 0

	for i := 0; ; {
		// 加一次并计算总分钟
		i++
		curTimeCopy = curTime.Add(time.Duration(i*timeInterval) * time.Minute)
		timeAddSumMinute = i * timeInterval

		// 检测是否命中开始/结束时间
		if taskStartTime == curTimeCopy {
			isStartHit = true
		}
		if taskEndTime == curTimeCopy {
			isEndHit = true
		}

		// 检测是否经过一轮（24h）
		if timeAddSumMinute >= 1440 {
			if isStartHit && isEndHit {
				return true
			}
			fmt.Println("开始/结束未命中，无法显示")
			return false
		}
	}

}

// TimelineRightString 显示右侧信息
func TimelineRightString(taskIndex int, task *Task, curTime time.Time) string {
	status := "DontNeedDisplay"

	if curTime == task.StartTime {
		status = "开始"
	} else if curTime == task.EndTime {
		status = "结束"
	}

	if status != "DontNeedDisplay" {
		infoRight := fmt.Sprintf("<--任务%v%s%s   ", taskIndex, task.Name, status)
		return infoRight
	}
	return "nil"
}

// TimelineTaskInfoString 显示任务统计信息
func (t *Task) TimelineTaskInfoString(timeInterval int, taskIndex int, timeFormat string) string {

	// 统计任务耗时
	// 原理：循环增加时间周期，然后将周期加到开始时间，当开始时间等于结束时间，就可以知道经过了多久
	timeConsuming := t.EndTime.Sub(t.StartTime)
	baseTime := time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC)
	taskInfo := fmt.Sprintf("任务%v:%v 耗时%s\n",
		taskIndex,
		t.Name,
		time.Time.Format(
			baseTime.Add(timeConsuming),
			timeFormat,
		))

	return taskInfo
}
