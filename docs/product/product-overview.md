# AOSP Workspace 产品概览

> 文档职责：说明产品目标、用户、核心价值和交付形态。实现细节见 `../architecture/` 和 `../design/`。

## 1. 产品结论

AOSP Workspace 是面向 AOSP 源码阅读与导航的跨编辑器工具。用户安装插件后，配置源码视图、可选完整
AOSP 根目录、可选 `out` 和构建变体，即可获得 Java/Kotlin、AIDL、C/C++、生成代码和 JNI 的
导航能力。

产品不要求用户为 IDE 单独维护 Gradle 模块清单，也不要求用户安装 Python 或执行历史生成脚本。

## 2. 目标用户

- Android Framework、System Server、SystemUI、Car 和 Mainline 模块开发者；
- 只同步或挂载部分 AOSP project 的开发者；
- 构建在 Linux 服务器、编辑在 macOS/Windows 的开发者；
- 使用软链接、junction、bind mount 或 overlay 组织轻量工作区的开发者；
- 同时使用 Android Studio/IntelliJ IDEA、VS Code 和 clangd 的开发者。

## 3. 核心价值

1. **轻量工作区**：完整 AOSP 可以存在于本机或服务器，但 IDE 只索引用户选择的视图。
2. **构建感知**：有 `out` 时消费 Soong metadata、srcjar 和 compdb；没有 `out` 时仍提供降级导航。
3. **跨语言导航**：AIDL 与生成语言互跳，Java/Kotlin 与 JNI C/C++ 互跳。
4. **跨平台复用**：共享索引引擎服务多个 IDE 客户端和操作系统。
5. **可诊断**：每个结果说明来源、置信度、缺失路径和降级原因。

## 4. 产品形态

| 产物 | 语言 | 作用 |
|---|---|---|
| Shared Index Engine | Go | 路径、拓扑、metadata、AIDL、JNI 和查询 |
| JetBrains Plugin | Kotlin | Android Studio/IDEA 项目模型、PSI、UI 和导航 |
| VS Code Extension | TypeScript | LSP 客户端、命令、clangd 和状态 UI |
| Remote Execution Bridge | Go，可选 | Linux 构建机上的 metadata 刷新和按需构建 |

## 5. 成功标准

- 用户不打开完整 AOSP，也能稳定索引目录 A 中的选定源码；
- 相同 profile 在 Linux、macOS、Windows 得到一致 RepoPath 身份；
- 无 `out` 不阻塞基本导航；有 `out` 时能力自动增强；
- 完整 AOSP 根目录不会因配置存在而被自动加入 IDE content root；
- 插件错误不会修改 AOSP checkout、挂载点、软链接或构建产物。
