cmake_minimum_required(VERSION 3.5)
project(aosp-native)

# 设置 C++ 标准
set(CMAKE_CXX_STANDARD 20)
# 设置 abi 架构
set(CMAKE_ANDROID_ARCH_ABI arm64-v8a)

# 判断 build.gradle 是否已经设置 BUILD_NATIVE_ROOT
if(IS_DIRECTORY ${BUILD_NATIVE_ROOT})
    message(STATUS "${PROJECT_NAME}, set ANDROID_ROOT via BUILD_NATIVE_ROOT=${BUILD_NATIVE_ROOT}")
    set(ANDROID_ROOT ${BUILD_NATIVE_ROOT})
else()
    # 如果 build.gradle 没有设置 BUILD_NATIVE_ROOT
    # 或者是用 CLion 打开这个工程，则需要这这里手动设置 ANDROID_ROOT
    set(ANDROID_ROOT "$ENV{HOME}/code/oppo")
    message(STATUS "${PROJECT_NAME}, ANDROID_ROOT=${ANDROID_ROOT}")
endif()

add_subdirectory(AndroidRuntime)
add_subdirectory(AndroidServices)
add_subdirectory(BootAnimation)
add_subdirectory(AppProcess)
add_subdirectory(ATrace)
add_subdirectory(DumpState)
add_subdirectory(DumpSys)
add_subdirectory(Installd)
add_subdirectory(ServiceManager)
add_subdirectory(InputFlinger)
add_subdirectory(SurfaceFlinger)
#add_subdirectory(SensorService)
#add_subdirectory(VibratorService)

##这三个模块还需要打开 AOSP_FWK_AV
#add_subdirectory(AudioServer)
#add_subdirectory(AudioFlinger)
#add_subdirectory(AudioPolicy)

#add_subdirectory(TensorFlow)
#add_subdirectory(NeuralNetworks)

