package main

import (
	"flag" // 用于命令行参数解析
	"fmt"  // 用于格式化输出
	"os"   // 提供对操作系统功能的访问，例如退出程序

	"github.com/IooCo/go-ai/phase1/loganalyze/analyzer" // 日志分析模块
	"github.com/IooCo/go-ai/phase1/loganalyze/parser"   // 日志解析模块
)

// 程序主入口
func main() {
	// 定义命令行参数 -f，用于指定要分析的 JSON 日志文件路径
	file := flag.String("f", "", "JSON 日志文件路径")
	flag.Parse() // 解析命令行参数

	// 检查是否提供了文件路径参数。如果没有，则给出错误提示并退出程序。
	if *file == "" {
		fmt.Fprintln(os.Stderr, "用法: loganalyze -f <日志文件>")   // 告诉用户正确的用法
		fmt.Fprintln(os.Stderr, "示例: loganalyze -f app.json") // 给出示例
		flag.Usage()                                          // 打印 flag 包自动生成的参数说明
		os.Exit(1)                                            // 使用非零值退出，表示错误
	}

	// 调用 parser.ParseJSONFile 函数解析指定的 JSON 日志文件
	// entries: 返回解析后的日志条目切片
	// err: 解析过程中出现的错误
	entries, err := parser.ParseJSONFile(*file)
	if err != nil { // 解析失败，输出错误并退出
		fmt.Fprintf(os.Stderr, "解析失败: %v\n", err)
		os.Exit(1)
	}

	// 调用 analyzer.Analyze 对所有日志条目进行分析
	// m: 分析结果的数据结构（可包含统计信息、异常信息等，根据 analyzer 实现而定）
	m := analyzer.Analyze(entries)

	// 调用 analyzer.Report 输出分析结果
	analyzer.Report(m)
}
