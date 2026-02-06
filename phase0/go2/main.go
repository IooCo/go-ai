package main

import (
	"fmt"
	"strings"
)

// 保持原有常量含义不变
const greeting = "你好，Go！"

// Task 表示一个待办任务
type Task struct {
	ID    int
	Title string
	Done  bool
}

// 为 Task 实现 String 方法，便于直接打印
func (t Task) String() string {
	var builder strings.Builder
	builder.WriteString("[")
	builder.WriteString(fmt.Sprint(t.ID))
	builder.WriteString("] ")
	builder.WriteString(t.Title)
	builder.WriteString(" (")
	if t.Done {
		builder.WriteString("已完成")
	} else {
		builder.WriteString("未完成")
	}
	builder.WriteString(")")
	return builder.String()
}

// TaskList 是 Task 的列表类型，为其定义一组操作方法
type TaskList struct {
	Tasks []Task
}

// NewTask 创建一个新的任务（使用值类型而不是指针）
func NewTask(id int, title string) Task {
	t := Task{
		ID:    id,
		Title: strings.TrimSpace(title),
		Done:  false,
	}
	return t
}

// Add 向列表中追加任务
func (ts *TaskList) Add(t Task) {
	ts.Tasks = append(ts.Tasks, t)
}

// Complete 将指定 ID 的任务标记为完成
func (ts *TaskList) Complete(id int) bool {
	for idx, task := range ts.Tasks {
		if task.ID == id {
			ts.Tasks[idx].Done = true
			return true
		}
	}
	return false
}

// Print 打印当前所有任务
func (ts *TaskList) Print() {
	fmt.Println("--- 任务列表 ---")
	for _, t := range ts.Tasks {
		fmt.Println(t)
	}
	fmt.Println()
}

// DoneCount 返回已完成任务数量
func (ts *TaskList) DoneCount() int {
	done := 0
	for _, t := range ts.Tasks {
		if t.Done {
			done++
		}
	}
	return done
}

// 初始化任务列表的高层封装，方便 main 函数调用
func initTasks() *TaskList {
	ts := &TaskList{}
	for i, name := range []string{"学习 Go 基础", "编写第一个程序", "理解接口与方法"} {
		ts.Add(NewTask(i+1, name))
	}
	return ts
}

// run 是程序的主流程控制函数
func run() {
	// 打印问候语
	fmt.Println(greeting)

	// 初始化任务
	tasks := initTasks()
	tasks.Print()

	// 完成第 2 个任务
	success := tasks.Complete(2)
	if success {
		fmt.Println("已将任务 2 标记为已完成。")
	}
	tasks.Print()

	// 统计并打印进度
	fmt.Printf("进度：%d/%d 已完成\n", tasks.DoneCount(), len(tasks.Tasks))
}

// 程序入口
func main() {
	run()
}
