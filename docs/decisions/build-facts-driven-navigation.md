# Build Facts 驱动导航决策

> 状态：Decision；替代“打开目录后主要依赖文件名猜测和按请求扫描”的 prototype 主路径。

## 1. 决策

AOSP Workspace 的最终导航主路径必须使用 Build Facts：优先本地可见 `out`，其次用户从
Linux 完整 AOSP 构建结果中复制的 Metadata Pack，再其次本地 Android.bp/Android.mk 静态
facts，最后才是 source-only fallback。

```text
Local exact out / copied Metadata Pack
        ↓
Soong metadata + module catalog + generated/source facts
        ↓
IDE native project model / Java index / clangd compdb
        ↓
AIDL/JNI shared cross-language edges
```

这不是对旧 Gradle 脚本的妥协复刻，而是提取其真正有效的部分：提前把 source roots、
classpath、jars、srcjars、generated files 和依赖关系转换为 IDE 能理解的模型。

## 2. out/pack 的宽松使用原则

Android release、manifest digest、product、variant、architecture 和 metadata schema 用于
描述精确性和增强能力，不作为 Java/framework 基础导航的硬门槛。

最新兼容 AOSP out 通常可以覆盖大部分 framework Java、模块 source roots、jars/srcjars、
AIDL 和基础 JNI 导航。即使 product/variant 不完全一致，也可以：

- 对 framework Java source roots 作为 baseline 使用；
- 对当前 Workspace 可见源码优先；
- 对不匹配 generated/compdb 标记 `compatible`/`approximate`；
- 不因缺少精确 product/variant 而拒绝整个 pack。

错误版本 metadata 可能导致错误跳转，因此 UI 必须显示匹配等级，但“无法 exact 匹配”不
应阻塞高价值 source navigation。

## 3. Profile 与工程隔离

一个本机可以同时有多个独立 profile：

```text
AOSP  -> /work/aosp-view  -> profile aosp
MIUI  -> /work/miui-view  -> profile miui
OPPO  -> /work/oppo-view  -> profile oppo
QCOM  -> /work/qcom-view  -> profile qcom
```

每个 profile 独立保存：

- Workspace View Root；
- canonical root；
- local out/metadata 或 pack；
- path mappings；
- selected modules；
- build facts match；
- index snapshot/cache；
- compdb overlay 和 diagnostics。

不同 profile 不共享 module graph、AIDL snapshot、JNI snapshot 或 compdb overlay。公共
`engines/aosp-index` 只共享 executable，不共享项目状态。

## 4. 厂商树策略

MIUI、OPPO、QCOM 等不要求官方 pack 精确覆盖所有私有 module。产品目标是：

- AOSP 基线覆盖公共 framework/Java/AIDL/C++；
- 厂商本地 out/BP facts 覆盖私有差异；
- 同名/冲突模块显示 diagnostics；
- 无法确认的 vendor generated/variant 降级为 source-only 或 approximate；
- 不为了追求 100% vendor 特例而破坏公共 AOSP 可用性。

## 5. 用户交付方式

公益项目不需要提供动态下载服务器。支持三种输入：

1. 用户本机完整 AOSP 的 `out`；
2. 用户在 Linux 构建机导出的 Metadata Pack，通过 U 盘、网盘、挂载目录或内网复制；
3. 用户自建私有 pack 制品仓库，协议格式与公共 pack 相同。

插件只读取用户已经放在本地的目录，不自动联网下载，也不把下载服务作为基础依赖。

Pack 默认必须包含 Java/AIDL 导航所需的 generated source archives。AIDL、R、sysprop、
logtags、proto 等生成 Java 如果只保留路径 metadata 而不包含对应 srcjar/entry，不能
视为完整 Core Pack；可选拆分只适用于 Native headers、超大非 Java generated 文件和
未被 selected module 引用的额外产物。
