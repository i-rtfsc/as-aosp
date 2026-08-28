# Rust Project 集成设计

> 状态：Decision；Rust Analyzer 使用本机归一化的 Soong rust-project。

## 1. 输入

Soong 在设置 `SOONG_GEN_RUST_PROJECT=1` 后生成：

```text
OUT_DIR/soong/rust-project.json
```

该文件包含 crate、依赖、cfg、edition、proc macro、sysroot 和 OUT_DIR。Portable Out
保持原始路径复制它，但 IDE 客户端不能直接假设其中路径已经适合本机。

## 2. 路径归一化

engine 读取原始 JSON，并在 IDE cache 中生成本机副本：

- 源码相对路径映射到 Workspace View Root；
- `out/...` 路径映射到本机 Out Root；
- sysroot 映射到 Workspace View Root 下的 prebuilt toolchain；
- source include/exclude dirs 使用同样规则；
- 已声明 path mapping 的外部绝对路径允许映射；
- 无法映射的绝对路径产生 diagnostic，不能静默保留构建机路径。

输出只写：

```text
<ide-cache>/rust/rust-project.json
```

不修改源码目录或 Out Root。

## 3. RPC 与客户端

RPC：

```text
aosp/rust/project
```

参数包含 Workspace View Root、Out Root、cacheRoot 和 pathMappings，返回本机 project
路径、crate 数和 diagnostics。

VS Code 在安装 rust-analyzer 时，将返回路径加入 workspace-scoped
`rust-analyzer.linkedProjects`。JetBrains 只有在 Rust 插件提供稳定 project API 时才接入；
没有 Rust 插件时只显示 capability，不修改 Project Model。

## 4. CLion CMake

`OUT_DIR/development/ide/clion/**/CMakeLists.txt` 包含构建机绝对 AOSP 根和 Linux toolchain。
它适合服务器原地打开，不适合作为跨机器 Portable Out 的默认 native 投影。

跨机器 C/C++ 导航继续以 `compile_commands.json` overlay 为权威。CMake metadata 仅用于
诊断、模块补充和未来显式 CLion import，不自动加载未经路径归一化的文件。

## 5. 验收

- Rust project 的 source、out、sysroot 和 proc macro 路径均映射到本机；
- VS Code 配置只写当前 workspace；
- Out Root 变化后旧 project cache 失效；
- 缺少 rust-project 时 Java/C++ 导航不受影响；
- 原始 CLion CMake 中的构建机路径不会写入 IDE 配置。
