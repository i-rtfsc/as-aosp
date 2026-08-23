# AIDL 语言与解析器设计

> 状态：Target；首版实现 source-only 语义，generated mapping 和 Soong variant 在后续切片接入。

## 1. 责任与非目标

本 bounded context 负责 AIDL 文本的 token、AST、类型 symbol、import resolver 和 source range。
它不负责：

- 复刻完整 AIDL compiler；
- `aidl_interface` 的稳定 API、hash、frozen API 检查；
- Java/C++/NDK/Rust backend 代码生成；
- Soong product/variant 求值；
- IDE PSI/LSP 生命周期。

## 2. 首版语法范围

支持跳转所需的：

- package；
- import；
- interface；
- parcelable；
- enum；
- union；
- method、参数、返回类型；
- parcelable 字段；
- 常量名称和源码 range；
- 常见方向修饰符 `in`、`out`、`inout`；
- 注释、字符串和 annotation 的跳过。

遇到未知语法时保留已解析的前缀，并返回可定位 diagnostic，不让整个 workspace 失效。

## 3. AST 与 symbol

```text
Document
  package
  imports[]
  declarations[]

Declaration
  kind: interface | parcelable | enum | union
  canonicalName
  source range
  members[]

Symbol
  canonicalName
  kind
  RepoPath
  source range
  origin: parser | generated-metadata
  confidence: exact | inferred
```

symbol key 使用 `package + name`，不使用本机绝对路径。

## 4. Resolver 规则

1. 完整类型名优先按 canonical name 查找；
2. 未限定类型名先查当前 package；
3. 再查当前文档显式 import；
4. 再查 include roots 下可见文档；
5. 多个不同文件命中时返回 `ambiguous`，不按扫描顺序猜测；
6. 缺失类型返回 `not-found`；
7. Java generated `IFoo.Stub`、`IFoo.Proxy` 的映射属于后续 generated peer 层。

## 5. 诊断

| Code | 含义 |
|---|---|
| `AIDL_PACKAGE_MISSING` | 文档缺少 package |
| `AIDL_IMPORT_MISSING` | import 后缺少类型名 |
| `AIDL_DECLARATION_MISSING_NAME` | 声明缺少名称 |
| `AIDL_UNEXPECTED_TOKEN` | parser 在当前位置无法继续 |
| `AIDL_IMPORT_NOT_FOUND` | resolver 找不到 import |
| `AIDL_TYPE_NOT_FOUND` | resolver 找不到类型 |
| `AIDL_TYPE_AMBIGUOUS` | 多个候选类型匹配 |

## 6. 文件系统边界

parser 接收 `(RepoPath, content)`，不直接读取文件。文件发现、SourceMount 和 PathMapper 由上层提供。
这样同一 parser 可以被 CLI、JetBrains PSI adapter 和 VS Code LSP 使用。

## 7. Acceptance

- interface、parcelable、enum、union fixture 能建立 symbol；
- package/import/type definition 能返回 source range；
- 同 package 和显式 import 都能解析；
- 缺失和歧义返回稳定 diagnostics；
- Windows/Linux 风格路径只影响文件访问，不改变 RepoPath symbol key；
- source-only 结果标记 `parser`/`inferred`，不宣称 generated source 已存在。
