# ShortLink Service

一个使用 **Go + Gin + MySQL + Redis** 实现的极简短链服务，提供生成短链和根据短码跳转原始链接的能力，并内置 Swagger 文档。

## 功能特性
- `POST /api/shorten`：接收长链接，返回 `http://localhost:8080/{code}` 形式的短链。
- `GET /{code}`：根据短码重定向到原始链接。
- Redis 做热点缓存，MySQL 持久化存储。
- 使用 `swag` 自动生成 Swagger 文档，可在运行时通过 `/swagger/index.html` 访问。

## 环境要求
- Go 1.20+（`go.mod` 指定 1.25.4，确保 Go 版本兼容）
- MySQL 8+
- Redis 7+
- `swag` CLI（安装：`go install github.com/swaggo/swag/cmd/swag@latest`）

## 快速开始
```bash
# 1. 安装依赖
go mod tidy

# 2. 启动 MySQL 与 Redis（可选：使用自带 docker-compose）
docker-compose up -d

# 3. 生成 Swagger（会写入 docs/）
make swagger

# 4. 启动服务
make run
# 或 make dev（等价于先生成 swagger 再运行服务）
```

默认配置位于 `internal/config/config.yaml`，包含服务端口、MySQL DSN（`shortlink` 数据库）和 Redis 连接信息，可按需调整。

## 常用 Make 目标
| 命令        | 说明                                   |
|-------------|----------------------------------------|
| `make tidy` | `go mod tidy`                          |
| `make build`| 构建二进制到 `bin/shortlink`           |
| `make run`  | 直接运行 `cmd/main.go`                 |
| `make swagger` | 运行 `swag init` 生成文档到 `docs/` |
| `make dev`  | 顺序执行 `make swagger` 与 `make run` |

## Swagger 文档
- 生成：`make swagger`（内部执行 `swag init --dir ./cmd,./internal/... --generalInfo main.go --parseInternal`）
- 访问：启动服务后打开 `http://localhost:8080/swagger/index.html`

## 目录结构
```
cmd/              # 程序入口 main.go
internal/api/     # HTTP handler & 路由
internal/service/ # 业务逻辑
internal/repo/    # MySQL / Redis 访问
internal/model/   # 数据模型
internal/config/  # 配置加载
docs/             # 自动生成的 Swagger 文件
```

欢迎根据自己的需求扩展更多短链统计、权限控制等功能。若重新定义字段或新增接口，别忘了补充注释并重新运行 `make swagger` 更新文档。***
