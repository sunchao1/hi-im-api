# hi-im-api 文档

> **hi-im-api** 是 [hi-im](https://github.com/) 生态的 **L2 内部契约库**（Go module），对应必嗨 `lib/comm` + `lib/mesg` + seqsvr IDL；**不独立部署**。  
> **许可证**：Apache License 2.0（见仓库根目录 `LICENSE`）

---

## 阅读顺序

| 顺序 | 文档 | 内容 |
|------|------|------|
| 1 | [技术设计文档.md](技术设计文档.md) | 定位、包结构、48B 头、CMD、Redis key、gRPC proto、semver、与 hubclient 边界 |
| 2 | [M1-实施清单.md](M1-实施清单.md) | 生态 M2 文件级任务（从 beehive 迁移 + 单测） |

---

## 生态位置

```text
hi-im-api（本仓库）       ← 契约：cmd / header / proto / redis key
hi-im-hubclient          ← 传输：连 Hub（依赖本仓库拼包）
hi-im-gateway / msgsvr … ← 业务（import api + hubclient）
```

全栈方案：beehive-im 仓库 `doc/hi-im-档C技术方案设计.md` §4.5～§4.7。

---

## 模块路径（建议）

```text
github.com/<org>/hi-im-api
```

消费者 `go get` 后：

```go
import (
    "github.com/<org>/hi-im-api/pkg/im/cmd"
    "github.com/<org>/hi-im-api/pkg/im/header"
    imv1 "github.com/<org>/hi-im-api/gen/go/im/v1"
    seqv1 "github.com/<org>/hi-im-api/gen/go/seq/v1"
)
```
