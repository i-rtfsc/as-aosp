# AOSP IDE 范围与非目标

> 状态：Decision。本文限定首个可交付版本的产品边界。

## 1. In scope

- RepoPath、Workspace View、Source Mount、overlay 和外部构建路径映射；
- source-only、metadata-partial、build-backed 和 stale 状态；
- Soong Java/CC metadata、module-info、srcjar 和 compdb 读取；
- AIDL package/import/type 导航和 generated peer 映射；
- 常见 JNI 静态导出与动态注册的双向导航；
- JetBrains 与 VS Code 客户端；
- Linux、macOS、Windows；
- 部分仓库、软链接、junction、bind/network mount 和 mixed view；
- 明确诊断、缓存和增量刷新。

## 2. 首版 Out of scope

- 完整复刻 Soong 所有 mutator、product variable 和 variant 求值；
- 自行实现完整 Java/Kotlin/C++ 类型系统；
- 用一个插件包同时安装到 JetBrains 和 VS Code；
- macOS/Windows 本机运行 Linux-only `ide_query.sh`；
- 自动修改 Android SDK、AOSP 源码、out、挂载点或软链接；
- 未经确认自动扩大到完整 AOSP 索引；
- 保证运行时反射、复杂宏和动态加载 JNI 100% 可解析；
- 在 `dev-ide` 保留旧 Gradle/CMake 兼容实现；
- 创建新的 Git 仓库承载新方案。

## 3. 能力声明边界

- `source-only` 结果是源码级或启发式结果，不证明当前产品构建图包含该文件。
- `build-backed` 证明结果来自构建产物，不证明 AOSP 完整构建成功。
- AOSP compdb 可能对同一源码的多个 variant 去重，不证明所有 variant 参数均被保存。
- fixture 测试通过不等于真实网络挂载、真实 Android Studio 或大型厂商树性能已验证。
