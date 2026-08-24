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

`Member` 与 `Parameter` 除整体声明 range 外还保存 `typeRange`。references 必须返回
具体类型 token（包括限定名、泛型和数组类型的完整 type span），不能长期用整个
method/field 声明代替。parser 无法恢复精确 token 时才降级为 `inferred`。

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

## 6. Generated peer 规则

首版只生成确定的候选路径，不运行 AIDL compiler：

| 声明 | Java | C++/NDK 候选 |
|---|---|---|
| 顶层 interface/type | `package/path/Type.java` | `package/path/Type.cpp`、`BpX.h`、`BnX.h` |
| parcelable/enum/union | `package/path/Type.java` | backend-specific type/header 候选 |
| nested declaration | 外层 `Outer.java` | 外层生成单元或 backend 输出，不能单独假设 `Nested.java` |

如果 interface 名为 `IFoo`，C++ Bp/Bn base name 通常为 `Foo`；这只是 AIDL 当前 generator
的命名规则，不代表所有厂商 backend 都完全一致。

每个 peer 返回：

```text
source AIDL RepoPath
generated candidate PathRef
language/backend
origin: naming-rule | build-metadata
confidence: inferred | exact
exists
```

source-only 只返回 `inferred`。当 out metadata 提供匹配路径且本机文件存在时升级为 `exact`；metadata
有路径但文件不可见时保留候选并产生 missing 诊断。

## 7. 文件系统边界

parser 接收 `(RepoPath, content)`，不直接读取文件。文件发现、SourceMount 和 PathMapper 由上层提供。
这样同一 parser 可以被 CLI、JetBrains PSI adapter 和 VS Code LSP 使用。

## 8. Acceptance

- interface、parcelable、enum、union fixture 能建立 symbol；
- package/import/type definition 能返回 source range；
- 同 package 和显式 import 都能解析；
- 缺失和歧义返回稳定 diagnostics；
- Windows/Linux 风格路径只影响文件访问，不改变 RepoPath symbol key；
- 直接从 AIDL 文本解析的声明标记 `parser`/`exact`；无 out 的 Java/C++ generated peer 映射标记 `inferred`，不宣称生成文件已存在。
- interface 名称 `IFoo` 能生成 Java、C++、Bp/Bn 和 NDK peer 候选；
- nested declaration 不被错误映射为独立 Java 文件；
- 匹配的 existing generated path 升级为 `build-metadata`/`exact`，缺失路径保持 `inferred` 并诊断。
