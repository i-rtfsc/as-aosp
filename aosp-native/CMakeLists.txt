cmake_minimum_required(VERSION 3.5)
project(aosp-native)

set(CMAKE_CXX_STANDARD 20)

#set(ANDROID_ROOT "/home/solo/code/aosp")
set(ANDROID_ROOT ${BUILD_NATIVE_ROOT})

# ---------- AndroidRuntime ----------
# frameworks/base/core/jni/Android.bp
# libandroid_runtime.so -> cc_library_shared
# frameworks/base/libs/androidfw/Android.bp
# libandroidfw -> cc_library
# frameworks/base/libs/hostgraphics/Android.bp
# libhostgraphics -> cc_library_host_static
# frameworks/base/libs/hwui/Android.bp
# libhwui -> cc_library
# frameworks/base/libs/input/Android.bp
# libinputservice -> cc_library_shared
# frameworks/base/libs/storage/Android.bp
# libstorage -> cc_library_static
# frameworks/base/media/jni/Android.bp
# libmedia_jni.so -> cc_library_shared
# frameworks/base/media/native/midi/Android.bp
# libamidi.so -> cc_library_shared
# frameworks/base/native/android/Android.bp
# libandroid.so -> cc_library_shared
# frameworks/base/cmds/uinput/jni/Android.bp
# libuinputcommand_jni.so -> cc_library_shared
add_subdirectory(AndroidRuntime)
# ---------- AndroidRuntime ----------

# ---------- AndroidServices ----------
# frameworks/base/libs/services/Android.bp
# libservices.so -> cc_library_shared
# frameworks/base/services/core/jni/Android.bp
# libservices.core -> cc_library_static
# frameworks/base/services/incremental/Android.bp
# service.incremental -> cc_library
# frameworks/base/services/Android.bp
# libandroid_servers.so -> cc_library_shared
add_subdirectory(AndroidServices)
# ---------- AndroidServices ----------

# ---------- BootAnimation ----------
# frameworks/base/cmds/bootanimation/Android.bp
# bootanimation -> cc_binary
add_subdirectory(BootAnimation)
# ---------- BootAnimation ----------

# ---------- AppProcess ----------
# frameworks/base/cmds/app_process/Android.bp
# app_process -> cc_binary
add_subdirectory(AppProcess)
# ---------- AppProcess ----------

# ---------- ATrace ----------
# frameworks/native/cmds/atrace/Android.bp
# atrace -> cc_binary
#add_subdirectory(ATrace)
# ---------- ATrace ----------

# ---------- DumpState ----------
# frameworks/native/cmds/dumpstate/Android.bp
# dumpstate -> cc_binary
add_subdirectory(DumpState)
# ---------- DumpState ----------

# ---------- DumpSys ----------
# frameworks/native/cmds/dumpsys/Android.bp
# dumpsys -> cc_binary
add_subdirectory(DumpSys)
# ---------- DumpSys ----------


# ---------- Installd ----------
# frameworks/native/cmds/installd/Android.bp
# installd -> cc_binary
add_subdirectory(Installd)
# ---------- Installd ----------

# ---------- ServiceManager ----------
# frameworks/native/cmds/servicemanager/Android.bp
# servicemanager -> cc_binary
# frameworks/native/cmds/service/Android.bp
# service/vndservice -> cc_binary
add_subdirectory(ServiceManager)
# ---------- ServiceManager ----------

# ---------- InputFlinger ----------
# frameworks/native/services/inputflinger/reader/Android.bp
# libinputreader.so -> cc_library_shared
# frameworks/native/services/inputflinger/dispatcher/Android.bp
# libinputdispatcher -> cc_library_static
# frameworks/native/services/inputflinger/Android.bp
# libinputflinger.so -> cc_library_shared
add_subdirectory(InputFlinger)
# ---------- InputFlinger ----------

# ---------- SensorService ----------
# frameworks/native/services/sensorservice/Android.bp
# sensorservice -> cc_binary
#add_subdirectory(SensorService)
# ---------- SensorService ----------

# ---------- SurfaceFlinger ----------
# frameworks/native/services/surfaceflinger/Android.bp
# surfaceflinger -> cc_binary
add_subdirectory(SurfaceFlinger)
# ---------- SurfaceFlinger ----------

# ---------- VibratorService ----------
# frameworks/native/services/vibratorservice/Android.bp
# libvibratorservice -> cc_library_shared
#add_subdirectory(VibratorService)
# ---------- VibratorService ----------

# ---------- AudioServer ----------
# frameworks/av/media/audioserver/Android.bp
# audioserver -> cc_binary
#add_subdirectory(AudioServer)
# ---------- AudioServer ----------

# ---------- AudioFlinger ----------
# frameworks/av/services/audioflinger/Android.bp
# libaudioflinger -> cc_library_shared
#add_subdirectory(AudioFlinger)
# ---------- AudioFlinger ----------

# ---------- AudioPolicy ----------
# frameworks/av/services/audiopolicy/service/Android.bp
# libaudiopolicyservice -> cc_library_shared
#add_subdirectory(AudioPolicy)
# ---------- AudioPolicy ----------

# ---------- TensorFlow ----------
# external/tensorflow
#add_subdirectory(TensorFlow)
# ---------- TensorFlow ----------

# ---------- NeuralNetworks ----------
# packages/modules/NeuralNetworks
#add_subdirectory(NeuralNetworks)
# ---------- NeuralNetworks ----------

