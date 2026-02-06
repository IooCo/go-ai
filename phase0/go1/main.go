// 包声明：每个 Go 源文件必须以 package 开头
// main 包是程序的入口，包含 main 函数的包才能被编译为可执行文件
package main

// 导入标准库或第三方包
// 可以用括号写多行导入，或 import "fmt" 单行
import (
	"fmt"
	"strings"
)

// ========== 常量与类型 ==========

// const 定义常量，编译时确定，不可修改
const greeting = "你好，Go！"

// 自定义结构体：表示一个「任务」
type Task struct {
	ID     int
	Title  string
	Done   bool
}

// String 方法：为 Task 实现 fmt.Stringer 接口
// 这样在 fmt 包打印时就会自动调用此方法
func (t Task) String() string {
	status := "未完成"
	if t.Done {
		status = "已完成"
	}
	return fmt.Sprintf("[%d] %s (%s)", t.ID, t.Title, status)
}

// ========== 业务函数 ==========

// 创建新任务，返回 Task 指针
func newTask(id int, title string) *Task {
	return &Task{ID: id, Title: strings.TrimSpace(title), Done: false}
}

// 完成指定 ID 的任务，返回是否找到并更新
func completeTask(tasks []*Task, id int) bool {
	for _, t := range tasks {
		if t.ID == id {
			t.Done = true
			return true
		}
	}
	return false
}

// 打印所有任务
func printTasks(tasks []*Task) {
	fmt.Println("--- 任务列表 ---")
	for _, t := range tasks {
		fmt.Println(t) // 会调用 Task.String()
	}
	fmt.Println()
}

// ========== 程序入口 ==========

func main() {
	// 1. 变量声明
	// 短变量声明 := 由编译器推断类型，仅能在函数内使用
	msg := greeting
	fmt.Println(msg)

	// 2. 切片（动态数组）：[]*Task 表示「指向 Task 的指针」的切片
	tasks := []*Task{
		newTask(1, "学习 Go 基础"),
		newTask(2, "编写第一个程序"),
		newTask(3, "理解接口与方法"),
	}
	printTasks(tasks)

	// 3. 完成第 2 个任务
	ok := completeTask(tasks, 2)
	if ok {
		fmt.Println("已将任务 2 标记为已完成。")
	}
	printTasks(tasks)

	// 4. 统计：遍历并计数
	doneCount := 0
	for _, t := range tasks {
		if t.Done {
			doneCount++
		}
	}
	fmt.Printf("进度：%d/%d 已完成\n", doneCount, len(tasks))
}
