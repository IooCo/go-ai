# LogAnalyze

面向 AI+教育 场景的 JSON 日志分析器，使用 Go 编写，通过 CLI 对日志文件进行解析和统计。

## 功能

- 解析 ndjson 格式（每行一条 JSON）的日志文件
- 统计总条数、活跃用户数、错误数、错误率
- 按级别（level）、动作（action）、来源（source）聚合
- 兼容常见字段名：`timestamp`/`time`、`level`/`lvl`、`message`/`msg` 等

## 环境要求

- Go 1.22+

## 安装与运行

```bash
# 克隆或进入项目目录
cd loganalyze

# 方式一：直接运行
go run . -f sample.json

# 方式二：编译后运行
go build -o loganalyze
./loganalyze -f sample.json
```

## 使用方法

```bash
loganalyze -f <日志文件路径>
```

| 参数 | 说明 |
|------|------|
| `-f` | JSON 日志文件路径（必填） |

## 输入格式

支持 **ndjson**（每行一条 JSON 对象）。推荐字段：

| 字段 | 说明 |
|------|------|
| `timestamp` / `time` | 时间戳 |
| `level` | 日志级别（info/error 等） |
| `source` | 来源（user_behavior/ai_interaction/system） |
| `user_id` | 用户 ID |
| `action` | 动作（login/chat/answer 等） |
| `message` / `msg` | 日志消息 |

示例（`sample.json`）：

```json
{"timestamp":"2025-02-15T10:00:00Z","level":"info","source":"user_behavior","user_id":"u001","action":"login","message":"用户登录"}
{"timestamp":"2025-02-15T10:01:00Z","level":"info","source":"ai_interaction","user_id":"u001","action":"chat","duration_ms":120,"message":"AI对话"}
{"timestamp":"2025-02-15T10:03:00Z","level":"error","source":"system","message":"接口超时"}
```

## 输出示例

```
========== 日志分析报告 ==========
总条数: 4
活跃用户数: 2
错误条数: 1
错误率: 25.00%

--- 按级别 ---
  error: 1
  info: 3

--- 按动作 ---
  answer: 1
  chat: 1
  login: 1

--- 按来源 ---
  ai_interaction: 1
  system: 1
  user_behavior: 2
```

## 项目结构

```
loganalyze/
├── main.go           # CLI 入口
├── model/            # 数据模型
│   └── log_entry.go
├── parser/           # JSON 解析
│   └── json.go
├── analyzer/         # 统计分析
│   └── metrics.go
├── sample.json       # 示例日志
├── go.mod
└── README.md
```

## AI使用说明

```
该项目通过人工进行项目的初始设定，再经由ai优化，人工审核后经ai编写，由于当前功力不够，所以未作过多人工修改，手动查看了一些人类看来待完善的问题让ai进行了完善。
ai编写了readme文件的大部分内容，只有ai使用说明是人工亲自写的。