package parser

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/IooCo/go-ai/phase1/loganalyze/model"
)

// ParseJSON 从 io.Reader 逐行解析 JSON 日志
// ParseJSON 从一个 io.Reader（比如文件）中按行读取 JSON 日志，并解析成 LogEntry 切片。
// 本函数适用于每行一条 JSON（ndjson）格式的日志文件。
// 参数 r：任何实现了 io.Reader 接口的对象（比如文件、网络连接）
// 返回值：1. LogEntry 指针切片，存放所有成功解析的日志条目；2. 错误信息
func ParseJSON(r io.Reader) ([]*model.LogEntry, error) {
	// 声明一个用于存放所有解析出来的日志条目的切片
	var entries []*model.LogEntry
	// 创建一个 scanner，用于按行读取 r 里的内容
	scanner := bufio.NewScanner(r)

	// 用 for 循环，不断读取每一行日志内容
	for scanner.Scan() {
		// 取出当前行的原始字节切片
		line := scanner.Bytes()
		// 如果当前行为空（长度为0），就跳过
		if len(line) == 0 {
			continue
		}

		// 声明一个 map，用于存放解析后的 JSON（类似于字典存储）
		var raw map[string]any
		// 将当前行（JSON 格式）反序列化到 raw 这个 map 里
		if err := json.Unmarshal(line, &raw); err != nil {
			fmt.Println("解析失败:", err)
			fmt.Println("当前行:", string(line))
			continue // 如果解析失败，忽略该行，继续处理下一行
		}

		// 把解析出来的 map 转化成我们项目中的 LogEntry 结构体
		entry := rawToEntry(raw)
		// 如果转换成功（entry 不为 nil），就加入到结果切片中
		if entry != nil {
			entries = append(entries, entry)
		}
	}

	// 检查是否在扫描过程中出错（比如文件读到一半损坏）
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("read error: %w", err) // 返回错误
	}

	// 返回所有解析得到的日志条目
	return entries, nil
}

// ParseJSONFile 专门用于从文件路径读取并解析 JSON 日志。
// 参数 path: 日志文件的路径
func ParseJSONFile(path string) ([]*model.LogEntry, error) {
	// 打开指定路径的文件
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open file: %w", err) // 打不开文件就返回错误
	}
	defer f.Close() // 函数结束时自动关闭文件

	// 调用 ParseJSON 对打开的文件进行解析
	return ParseJSON(f)
}

// rawToEntry 用于把解析后的 map 映射成 LogEntry 结构体，自动兼容常见字段名。
// 参数 raw: 解析自一行 JSON 的 key-value 映射
func rawToEntry(raw map[string]any) *model.LogEntry {
	// 新建一个 LogEntry 结构体，并把 Metadata 字段设置成原始 map
	entry := &model.LogEntry{
		Metadata: raw,
	}

	// 解析时间戳字段。依次尝试各种常见字段名
	for _, key := range []string{"timestamp", "time", "@timestamp", "ts"} {
		if v, ok := raw[key]; ok {
			entry.Timestamp = parseTime(v) // 用 parseTime 进行解析
			break                          // 找到就不再继续查找
		}
	}
	// 如果上述字段都没有，或者解析失败，就赋值为当前时间
	if entry.Timestamp.IsZero() {
		entry.Timestamp = time.Now()
	}

	// 解析日志级别（level），尝试多种常见写法
	for _, key := range []string{"level", "lvl", "severity"} {
		if v, ok := raw[key]; ok {
			entry.Level = toString(v) // 转成字符串
			break
		}
	}

	// 解析日志的消息内容（message），兼容多种字段名
	for _, key := range []string{"message", "msg", "log"} {
		if v, ok := raw[key]; ok {
			entry.Message = toString(v)
			break
		}
	}

	// 解析教育/AI场景中常见的自定义字段，如果有就转换保存
	if v, ok := raw["user_id"]; ok {
		entry.UserID = toString(v)
	}
	if v, ok := raw["action"]; ok {
		entry.Action = toString(v)
	}
	if v, ok := raw["source"]; ok {
		entry.Source = toString(v)
	}
	if v, ok := raw["duration_ms"]; ok {
		entry.Duration = toInt64(v)
	}

	// 返回转换好的结构体指针
	return entry
}

// parseTime 负责解析多种时间格式。支持字符串/数字表示的时间。
// 参数 v: 任意类型的时间字段
func parseTime(v any) time.Time {
	switch t := v.(type) {
	case string:
		// 依次尝试常用时间格式
		for _, layout := range []string{
			time.RFC3339,                // 例如 2025-02-15T10:01:00Z
			time.RFC3339Nano,            // 更精确
			"2006-01-02 15:04:05",       // 经典格式
			"2006-01-02T15:04:05Z07:00", // 另一种常见格式
		} {
			if parsed, err := time.Parse(layout, t); err == nil {
				return parsed // 成功就返回
			}
		}
	case float64:
		// 数字格式，区分毫秒和秒级时间戳
		if t > 1e12 {
			return time.UnixMilli(int64(t)) // 毫秒
		}
		return time.Unix(int64(t), 0) // 秒
	}
	// 如果没有解析成功，返回零值（1970）
	return time.Time{}
}

// toString 负责把任意类型安全地转为字符串
// nil 返回空字符串，其他类型直接 fmt.Sprint
func toString(v any) string {
	if v == nil {
		return ""
	}
	switch t := v.(type) {
	case string:
		return t
	default:
		return fmt.Sprint(v) // 其他类型转字符串
	}
}

// toInt64 负责把任意类型转换为 int64 数字，常用于 duration 等数值型字段
func toInt64(v any) int64 {
	switch t := v.(type) {
	case float64:
		return int64(t)
	case int:
		return int64(t)
	case int64:
		return t
	default:
		return 0 // 不能转换时返回0
	}
}
