<!-- TOC -->
* [native 模块](#native-模块)
  * [AndroidRuntime](#androidruntime)
  * [AndroidServices](#androidservices)
  * [BootAnimation](#bootanimation)
  * [AppProcess](#appprocess)
  * [ATrace](#atrace)
  * [DumpState](#dumpstate)
  * [DumpSys](#dumpsys)
  * [Installd](#installd)
  * [ServiceManager](#servicemanager)
  * [InputFlinger](#inputflinger)
  * [SensorService](#sensorservice)
  * [SurfaceFlinger](#surfaceflinger)
  * [VibratorService](#vibratorservice)
  * [AudioServer](#audioserver)
  * [AudioFlinger](#audioflinger)
  * [AudioPolicy](#audiopolicy)
  * [TensorFlow](#tensorflow)
  * [NeuralNetworks](#neuralnetworks)
<!-- TOC -->

# native 模块

## AndroidRuntime

最初 AndroidRuntime 只是想表达 libandroid_runtime.so 模块； 但在使用过程中发现如果只包含头文件，那么只能跳转到.h，不能跳转到.cpp。
所以把 frameworks/base/core/jni/Android.bp 里 shared_libs 中的常用模块也添加进来。

- frameworks/base/core/jni/Android.bp
> libandroid_runtime.so -> cc_library_shared

- frameworks/base/libs/androidfw/Android.bp
> libandroidfw -> cc_library

- frameworks/base/libs/hostgraphics/Android.bp 
> libhostgraphics -> cc_library_host_static

- frameworks/base/libs/hwui/Android.bp
> libhwui -> cc_library

- frameworks/base/libs/input/Android.bp
> libinputservice -> cc_library_shared

- frameworks/base/libs/storage/Android.bp
> libstorage -> cc_library_static

- frameworks/base/media/jni/Android.bp
> libmedia_jni.so -> cc_library_shared

- frameworks/base/media/native/midi/Android.bp
> libamidi.so -> cc_library_shared

- frameworks/base/native/android/Android.bp
> libandroid.so -> cc_library_shared

- frameworks/base/cmds/uinput/jni/Android.bp
> libuinputcommand_jni.so -> cc_library_shared

- frameworks/native/libs/battery/Android.bp
> libbattery -> cc_library

- frameworks/native/libs/binder/Android.bp
> libbinder -> cc_library

- frameworks/native/libs/ui/Android.bp
> libui -> cc_library_shared

- frameworks/native/libs/graphicsenv/Android.bp
> libgraphicsenv -> cc_library_shared

- frameworks/native/libs/gui/Android.bp
> libgui -> cc_library_shared

- frameworks/base/libs/hwui/Android.bp
> libhwui -> cc_library

- frameworks/native/libs/permission/Android.bp
> libpermission -> cc_library

- frameworks/native/libs/sensor/Android.bp
> libsensor -> cc_library

- frameworks/native/libs/input/Android.bp
> libinput -> cc_library

- frameworks/native/libs/nativedisplay/Android.bp
> libnativedisplay -> cc_library_shared

- frameworks/native/libs/nativewindow/Android.bp
> libnativewindow -> cc_library

- system/core/libcutils/Android.bp
> libcutils -> cc_library

- system/core/libutils/Android.bp
> libutils -> cc_library


以下两个模块未添加到 AndroidRuntime ，目前觉得加到 media 比较合适。
后续打算新增 MediaServer 包含 video + audio(目前的当audio模块)
> frameworks/av/media/libaudioclient/Android.bp
> frameworks/av/media/libaudiofoundation/Android.bp


## AndroidServices

包含以下 Android.bp 等：
- frameworks/base/libs/services/Android.bp
> libservices.so -> cc_library_shared

- frameworks/base/services/core/jni/Android.bp
> libservices.core -> cc_library_static

- frameworks/base/services/incremental/Android.bp
> service.incremental -> cc_library

- frameworks/base/services/Android.bp
> libandroid_servers.so -> cc_library_shared


## BootAnimation

包含以下 Android.bp 等：

- frameworks/base/cmds/bootanimation/Android.bp
> bootanimation -> cc_binary


## AppProcess

包含以下 Android.bp 等：

- frameworks/base/cmds/app_process/Android.bp
> app_process -> cc_binary


## ATrace

包含以下 Android.bp 等：

- frameworks/native/cmds/atrace/Android.bp
> atrace -> cc_binary


## DumpState

包含以下 Android.bp 等：

- frameworks/native/cmds/dumpstate/Android.bp
> dumpstate -> cc_binary


## DumpSys

包含以下 Android.bp 等：

- frameworks/native/cmds/dumpsys/Android.bp
> dumpsys -> cc_binary


## Installd

包含以下 Android.bp 等：

- frameworks/native/cmds/installd/Android.bp
> installd -> cc_binary

## ServiceManager

包含以下 Android.bp 等：

- frameworks/native/cmds/servicemanager/Android.bp
> servicemanager -> cc_binary

- frameworks/native/cmds/service/Android.bp
> service/vndservice -> cc_binary


## InputFlinger

包含以下 Android.bp 等：

- frameworks/native/services/inputflinger/reader/Android.bp
> libinputreader.so -> cc_library_shared

- frameworks/native/services/inputflinger/dispatcher/Android.bp
> libinputdispatcher -> cc_library_static

- frameworks/native/services/inputflinger/Android.bp
> libinputflinger.so -> cc_library_shared


## SensorService

包含以下 Android.bp 等：

- frameworks/native/services/sensorservice/Android.bp
> sensorservice -> cc_binary


## SurfaceFlinger

包含以下 Android.bp 等：

- frameworks/native/services/surfaceflinger/Android.bp
> surfaceflinger -> cc_binary


## VibratorService

包含以下 Android.bp 等：

- frameworks/native/services/vibratorservice/Android.bp
> libvibratorservice -> cc_library_shared


## AudioServer

包含以下 Android.bp 等：

- frameworks/av/media/audioserver/Android.bp
> audioserver -> cc_binary


## AudioFlinger

包含以下 Android.bp 等：

- frameworks/av/services/audioflinger/Android.bp
> libaudioflinger -> cc_library_shared

## AudioPolicy

包含以下 Android.bp 等：

- frameworks/av/services/audiopolicy/service/Android.bp
> libaudiopolicyservice -> cc_library_shared

## TensorFlow

external/tensorflow

## NeuralNetworks

packages/modules/NeuralNetworks
