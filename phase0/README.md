# phase0

## 项目结构
| 目录名 | 类型         | 说明                |
|--------|--------------|--------------------|
| go1    | AI 生成      | 人工智能直接生成代码 |
| go2    | AI 重构      | 基于 go1 重构优化版  |
| hello  | 独立项目     |     手工编写        |

## 快速运行
### 环境要求
- Go 1.21+

### 运行命令
```bash
# 运行 go1
cd go1 && go run main.go

# 运行 go2
cd go2 && go run main.go

# 运行 hello
cd hello && go run main.go
