package model

import "time"

// LogEntry 统一的日志条目（AI+教育场景）
type LogEntry struct {
	Timestamp time.Time         `json:"timestamp"`
	Level     string            `json:"level"`
	Source    string            `json:"source"`
	UserID    string            `json:"user_id"`
	Action    string            `json:"action"`
	Duration  int64             `json:"duration_ms"`
	Message   string            `json:"message"`
	Metadata  map[string]any    `json:"metadata"`
}
