# as aosp

## 简介
此as工程可以快速的导入aosp framework(包含java/native部分)、 aosp 系统app、 国内某些厂商扩展的fwk代码，比这个 https://www.jianshu.com/p/2ba5d6bd461e 方案还快，并且“联想”也很方便。

## 警告
关于此工程中包含miui、flyme、oppo、vivo等配置，只是一个空的gradle配置；该工程中不涉及到任何这几家公司的代码。

请勿小事化大！尊重事实，切勿胡说八道！


## 配置介绍

### settings.gradle
settings.gradle 是子项目（也可以说是module）的配置文件。
如果需要新增一个aosp的模块，需要在这里配置。

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

### config.gradle
allModules不需要手动维护，在settings.gradle里新增一个module，脚本会自动识别到project name并添加到allModules数组。
这个allModules数组的作用是方便每个module快速的互相依赖，如：
```bash
rootProject.ext.allModules.each { dependence -> compileOnly project(dependence.value) }
```
以上写法会循环依赖，真正的app gradle工程不能这样做的。这里这样子做是因为我们只是为了方便as阅读代码或者改代码，真正编译的时候还是用ninja。

### aosp.gradle
配置 aosp 的模块，比如framework、services、SystemUI、Settings等等。
以及指定源码目录等。

ext.gradle、miui.gradle、flyme.gradle、oppo.gradle、vivo.gradle同理。


## 编译
此功能无法编译framework.jar或者services.jar，请使用aosp推荐的编译方式。
若要编译demo app调试，可以查看settings.gradle注释，根据提示注释一些module；并在config.gradle中把enable_boot_jar、build_app设置为true。

如果需要还得copy framework.jar和services.jar到system_libs中
```bash
cp out/target/common/obj/JAVA_LIBRARIES/framework-minus-apex_intermediates/classes.jar ../system_libs/framework-minus-apex.jar
cp out/target/common/obj/JAVA_LIBRARIES/services_intermediates/classes.jar ../system_libs/services.jar
```
