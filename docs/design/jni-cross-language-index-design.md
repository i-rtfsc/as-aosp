# JNI 跨语言索引设计

> 状态：Target；首版实现 source-only 静态/注册表索引，复杂宏和运行时绑定保持 unresolved。

## 1. 责任与非目标

本 bounded context 负责把 Java/Kotlin native declaration 与 C/C++ JNI implementation 关联起来。
它不负责完整 C++ 类型系统、宏求值、链接器符号证明或运行时反射。

## 2. Java/Kotlin 输入

支持：

- Java `native` method；
- Kotlin `external fun`；
- optional `@FastNative`、`@CriticalNative` 作为 metadata；
- package、nested class 和方法签名。

Java identity：

```text
JniJavaSymbol
  className: android.view.InputEventReceiver
  methodName: nativeDispose
  descriptor: (J)V, when available
  RepoPath + source range
```

## 3. C/C++ 输入

支持：

1. 静态导出：`Java_android_pkg_Class_method`；
2. `JNINativeMethod` 数组；
3. `MAKE_JNI_NATIVE_METHOD` 常见宏；
4. `jniRegisterNativeMethods`；
5. `RegisterMethodsOrDie`；
6. `AndroidRuntime::registerNativeMethods`；
7. `JNIEnv::RegisterNatives`；
8. `JNI_OnLoad` 中可直接识别的注册调用。

## 4. Descriptor

解析 JNI descriptor：

```text
primitive: Z B C S I J F D V
object: Landroid/os/Bundle;
array: [I, [Ljava/lang/String;
method: (Ljava/lang/String;I)V
```

descriptor 错误只影响 overload 精确匹配，不阻止按 class/method 建立 inferred link。

## 5. Class binding

注册调用中的 class string 使用 `/`，转换为 Java FQCN：

```text
android/view/InputEventReceiver
  -> android.view.InputEventReceiver
```

字符串字面量可以得到 `exact` class binding；运行时拼接或函数返回值只能得到 `unresolved`。

## 6. 结果等级

| 等级 | 证据 |
|---|---|
| `exact` | 注册表 name/signature/function pointer，或完整静态 JNI 导出名 |
| `inferred` | 唯一 class/method 候选或不完整 descriptor |
| `unresolved` | 多候选、复杂宏、动态 class 或缺失 Java/C++ 对端 |

每条 `JniLink` 包含 Java symbol、native symbol、registration site、origin、confidence 和 diagnostics。

## 7. Acceptance

- Java native 和 Kotlin external fixture 可建立 declaration；
- 静态导出名能反解 package/class/method；
- `JNINativeMethod` 表能解析 name、signature、function pointer；
- 注册函数能绑定 class string；
- Java/C++ 双向查询返回相同 link；
- 多候选和动态 class 不被错误标记 exact。
