package analyzer

import (
	"fmt"  // 导入 fmt 包，用于打印输出
	"sort" // 导入 sort 包，用于排序

	"github.com/IooCo/go-ai/phase1/loganalyze/model" // 导入项目中的日志模型
)

// Metrics 是用于存储分析结果的结构体
type Metrics struct {
	Total    int                 // 总日志条数
	ByLevel  map[string]int      // 各级别（level）的日志计数，例如 info、error
	ByAction map[string]int      // 各 action（动作）的日志计数，例如 login、chat
	BySource map[string]int      // 各 source（来源）的日志计数，例如 user_behavior、system
	UserSet  map[string]struct{} // 记录活跃的用户 ID 的集合，struct{} 是零内存占用的空结构体
	Errors   int                 // 错误日志的条数（level == "error"）
}

// Analyze 函数对一组日志条目进行统计和分析
// 参数 entries：LogEntry 指针切片，表示所有需要分析的日志
// 返回值：返回包含所有统计结果的 Metrics 结构体指针
func Analyze(entries []*model.LogEntry) *Metrics {
	// 创建 Metrics 结构体，并初始化所有 map 字段，避免后续使用时发生空指针错误
	m := &Metrics{
		ByLevel:  make(map[string]int),      // 初始化 level 计数字典
		ByAction: make(map[string]int),      // 初始化 action 计数字典
		BySource: make(map[string]int),      // 初始化 source 计数字典
		UserSet:  make(map[string]struct{}), // 初始化用户集合
	}

	// 遍历所有日志条目
	for _, e := range entries {
		m.Total++ // 每遍历一条日志就加 1

		// 对 level 字段进行计数
		if e.Level != "" { // 判断 level 是否为空
			m.ByLevel[e.Level]++ // 对应级别的次数加 1
		}

		// 对 action 字段进行计数
		if e.Action != "" { // 判断 action 是否为空
			m.ByAction[e.Action]++ // 对应动作的次数加 1
		}

		// 对 source 字段进行计数
		if e.Source != "" { // 判断 source 是否为空
			m.BySource[e.Source]++ // 对应来源的次数加 1
		}

		// 统计活跃用户数，只记录 user_id 存在的日志
		if e.UserID != "" { // 判断用户 ID 是否为空
			m.UserSet[e.UserID] = struct{}{} // 用 map 模拟集合，只存 key，不关心 value
		}

		// 统计 error 日志的条数，区分大小写及常见拼写
		if e.Level == "error" || e.Level == "ERROR" {
			m.Errors++ // 错误计数器加 1
		}
	}

	return m // 返回结果
}

// Report 函数将分析结果以人类可读的格式输出到控制台
// 参数 m：Metrics 分析结果结构体指针
func Report(m *Metrics) {
	// 打印报告标题
	fmt.Println("========== 日志分析报告 ==========")
	// 打印总日志条数
	fmt.Printf("总条数: %d\n", m.Total)
	// 打印活跃用户数，取 UserSet 的 key 的数量
	fmt.Printf("活跃用户数: %d\n", len(m.UserSet))
	// 打印错误日志条数
	fmt.Printf("错误条数: %d\n", m.Errors)
	// 打印错误率（错误日志数/总日志数），注意先判断分母不为 0
	if m.Total > 0 {
		fmt.Printf("错误率: %.2f%%\n", float64(m.Errors)/float64(m.Total)*100)
	}

	// 按 level 分类统计结果输出
	if len(m.ByLevel) > 0 { // 有数据才输出
		fmt.Println("\n--- 按级别 ---")
		printMap(m.ByLevel) // 调用 printMap 打印每个级别及其数量
	}
	// 按 action 分类统计结果输出
	if len(m.ByAction) > 0 { // 有数据才输出
		fmt.Println("\n--- 按动作 ---")
		printMap(m.ByAction)
	}
	// 按 source 分类统计结果输出
	if len(m.BySource) > 0 { // 有数据才输出
		fmt.Println("\n--- 按来源 ---")
		printMap(m.BySource)
	}
}

// printMap 用于将 map[string]int 类型的数据按 key 排序后打印出来
// 参数 m：待打印 map
func printMap(m map[string]int) {
	var keys []string  // 创建一个字符串切片用于保存所有 key
	for k := range m { // 遍历 map 的所有 key
		keys = append(keys, k) // 将 key 加入 keys 列表
	}
	sort.Strings(keys)       // 对 key 切片进行字母序排序
	for _, k := range keys { // 按排序后的顺序输出每个 key 及对应的值
		fmt.Printf("  %s: %d\n", k, m[k])
	}
}
