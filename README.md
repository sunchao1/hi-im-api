# hi-im-api

hi-im 生态 **L2 内部契约库**（Go module）：IM 报头、CMD 命令字、Protobuf 消息体、Redis 键规范、gRPC seq 定义。

**不独立部署** — 由 gateway、msgsvr、hi-im-hubclient 等通过 `go get` 引用。

| 文档 | 说明 |
|------|------|
| [doc/技术设计文档.md](doc/技术设计文档.md) | 主设计 |
| [doc/M1-实施清单.md](doc/M1-实施清单.md) | 生态 M2 任务清单 |

**必嗨等价**：`lib/comm` + `lib/mesg` + seqsvr IDL  
**生态总览**：beehive-im `doc/hi-im-档C技术方案设计.md`

**许可证**：Apache License 2.0

## Module

```text
github.com/sunchao1/hi-im-api
```

## 包结构

```text
pkg/im/header   48B 报头编解码（wire 52B，与 beehive MesgHeadHton 兼容）
pkg/im/cmd      CMD 命令字
pkg/im/const    TTL、LSND_TYPE 等常量
pkg/rediskey    Redis 键格式化
pkg/errno       错误码
pkg/httpx       HTTP JSON 响应壳
gen/go/im/v1    IM protobuf 消息
gen/go/seq/v1   SeqService gRPC
```

## 开发

```bash
# 需 protoc + protoc-gen-go + protoc-gen-go-grpc
make gen    # 生成 proto
make test   # go test ./...
```

## 引用示例

```go
import (
    "github.com/sunchao1/hi-im-api/pkg/im/cmd"
    "github.com/sunchao1/hi-im-api/pkg/im/header"
    imv1 "github.com/sunchao1/hi-im-api/gen/go/im/v1"
    seqv1 "github.com/sunchao1/hi-im-api/gen/go/seq/v1"
)
```
