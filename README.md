# as aosp

## 简介

此as工程可以快速的导入aosp framework(包含java/native部分)、 aosp 系统app、 国内某些厂商扩展的fwk代码，比这个 https://www.jianshu.com/p/2ba5d6bd461e 方案还快，并且“联想”也很方便。

下面以cpp代码为例，演示代码的跳转以及提示、补全。

### 跳转
![](./res-readme/code_jump.gif)

### 提示、补全
![](./res-readme/code_completion.gif)


## 警告

关于此工程中包含miui、flyme、oppo、vivo等配置，只是一个空的gradle配置；该工程中不涉及到任何这几家公司的代码。

请勿小事化大！切勿胡说八道！



## 配置介绍

### settings.gradle

settings.gradle位于根目录下，用于项目的配置，常见的是配置子工程。一个子工程只有在setting.gradle中配置了，才能够被识别，构建的时候才会被包含进去。
如果需要新增一个aosp的模块，需要在这里配置。

使用者可以根据自己的需求修改settings.gradle的子模块，比如Settings开发不需要native部分、SystemUI部分等，可以注释使其不被识别，加快加载as速度，降低内存。


### 根目录build.gradle
根目录build.gradle文件配置了很多扩展的gradle脚本，可根据实际情况自行添加到对应的脚本或者新增脚本。
```bash
apply from: 'scripts/config.gradle'
apply from: 'scripts/aosp.gradle'
apply from: 'scripts/ext.gradle'
apply from: 'scripts/miui.gradle'
apply from: 'scripts/flyme.gradle'
apply from: 'scripts/oppo.gradle'
apply from: 'scripts/vivo.gradle'
```


#### config.gradle

config.gradle脚本除了配置基本的android sdk；还有一个很重要的功能，就是获取所有的模块名称（也就是settings.gradle配置的子模块）。

也就是说 allModules不需要手动维护，在settings.gradle里新增一个module，脚本会自动识别到project name并添加到allModules数组。
这个allModules数组的作用是方便每个module快速的互相依赖，如：
```bash
rootProject.ext.allModules.each { dependence -> compileOnly project(dependence.value) }
```
以上写法会循环依赖，真正的app gradle工程不能这样做的。这里这样子做是因为我们只是为了方便as阅读代码或者改代码，真正编译的时候还是用ninja、make。

#### aosp.gradle

- aospDir: 设置源码所在的目录，如: aospDir = "/home/solo/code/flyme"。
- aosp: 一个大数组，维护很多模块需要的路径。 
  - root: 等同于aospDir所设置的android源码根目录。
  - framework: 配置framework.jar的源码路径
  - services: 配置services.jar的源码路径
  - frameworkRes: 配置frameworkRes.apk的源码路径
  - SystemUI: 配置SystemUI.apk的源码路径
  - SystemUIPluginLib: 配置SystemUIPluginLib.jar的源码路径
  - Settings: 配置Settings.apk的源码路径
  - SettingsLib: 配置SettingsLib.aar的源码路径
  - SettingsProvider: 配置SettingsProvider.apk的源码路径
  - carFramework: 配置carFramework的源码路径
  - carServices: 配置carServices.apk的源码路径

> 以上的 framework、services、frameworkRes、SystemUI、SystemUIPluginLib、Settings、SettingsLib、SettingsProvider、carFramework、carServices
> 通过manifest、res、assets、jni、src来分别配置AndroidManifest.xml、资源文件目录、assets目录、jni代码目录、java\kt源码目录等。
> 
> 不需要的可以写空或者随便写一个不存在的文件、目录。

> 温馨提示
> 
> 在这个文件中只需要把 aospDir 设置成自己android源码的根目录就可以。
> 
> 里面具体模块的源码路径基本上都添加了（但确实不是100%添加），如果因为使用aosp版本不一致或者别的原因可以根据自己需要再添加。

#### ext.gradle

扩展framework源码所在的目录，一般用不得。

如果有需要可以改extDir对应的目录即可。


#### miui.gradle

miui代码所在的路径，主要是配置了framework、services、frameworkRes。

可以根据自己的需要修改miuiDir对应的目录即可。

#### flyme.gradle

flyme代码所在的路径，主要是配置了framework、services、frameworkRes、SystemUI、SystemUIPluginLib。

可以根据自己的需要修改flymeDir对应的目录即可。

#### oppo.gradle

oppo代码所在的路径，主要是配置了framework、services、frameworkRes。

可以根据自己的需要修改oppoDir对应的目录即可。

#### vivo.gradle

vivo代码所在的路径，主要是配置了framework、services、frameworkRes。

可以根据自己的需要修改vivoDir对应的目录即可。


### native

通过根目录下的settings.gradle可以看到有如下的配置：

```bash
/*************** aosp native ***************/
include ':aosp-native'
//include ':android-runtime'
//include ':android-services'
//include ':inputflinger'
//include ':surfaceflinger'
/*************** aosp native ***************/
```

- android-runtime: 对应的是frameworks/base/core/jni/Android.bp写的libandroid_runtime。也就是frameworks base core jni。
- android-services: 对应的是 frameworks/base/libs/services、frameworks/base/services/core/jni、frameworks/base/services/incremental。也就是libservic、libservices.core、libservices.core-gnss、service.incremental的和。
- inputflinger: 对应的是frameworks/native/services/inputflinger。
- surfaceflinger: 对应的是frameworks/native/services/surfaceflinger。
- aosp-native: 是把以上四个模块的整合到一起了，这么做是因为都放在一个模块里跳转方便，占用的内存也最少。

在每个cmake文件里都设置了这么一些变量，主要是用来控制是否加载相应的代码目录。

这里基本上只是把常用到的都打开了，如果需要把全部打开，改成true即可。
（如果都打开as占用内存会很大）
```bash
set(AOSP_SYSTEM_COMMON true)
set(AOSP_OUT false)
set(AOSP_AV false)
set(AOSP_ART false)
set(AOSP_BIONIC false)
set(AOSP_SYSTEM false)
set(AOSP_EXTERNAL false)
set(AOSP_PACKAGES false)
set(AOSP_HARDWARE false)
```

## 编译
此功能无法编译framework.jar或者services.jar，请使用aosp推荐的编译方式。
若要编译demo app调试，可以查看settings.gradle注释，根据提示注释一些module；并在config.gradle中把enable_boot_jar、build_app设置为true。

如果需要还得copy framework.jar和services.jar到system_libs中
```bash
cp out/target/common/obj/JAVA_LIBRARIES/framework-minus-apex_intermediates/classes.jar ../system_libs/framework-minus-apex.jar
cp out/target/common/obj/JAVA_LIBRARIES/services_intermediates/classes.jar ../system_libs/services.jar
```
