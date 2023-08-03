cmake_minimum_required(VERSION 3.5)
project(android-services)

set(ANDROID_ROOT ${BUILD_NATIVE_ROOT})

file(GLOB SOURCE_FILES
        ${ANDROID_ROOT}/frameworks/base/libs/services/*.c
        ${ANDROID_ROOT}/frameworks/base/libs/services/*.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/services/src/content/*.c
        ${ANDROID_ROOT}/frameworks/base/libs/services/src/content/*.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/services/src/os/*.c
        ${ANDROID_ROOT}/frameworks/base/libs/services/src/os/*.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/*.c
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/*.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/BroadcastRadio/*.c
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/BroadcastRadio/*.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/gnss/*.c
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/gnss/*.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/stats/*.c
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/stats/*.cpp
        ${ANDROID_ROOT}/frameworks/base/services/incremental/*.c
        ${ANDROID_ROOT}/frameworks/base/services/incremental/*.cpp
        ${ANDROID_ROOT}/frameworks/base/services/incremental/test/*.c
        ${ANDROID_ROOT}/frameworks/base/services/incremental/test/*.cpp
)

include_directories(
        ${ANDROID_ROOT}/frameworks/base/libs/services
        ${ANDROID_ROOT}/frameworks/base/libs/services/include/android/content
        ${ANDROID_ROOT}/frameworks/base/libs/services/include/android/os
        ${ANDROID_ROOT}/frameworks/base/libs/services/src/content
        ${ANDROID_ROOT}/frameworks/base/libs/services/src/os
        ${ANDROID_ROOT}/frameworks/base/services/core/jni
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/BroadcastRadio
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/gnss
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/stats
        ${ANDROID_ROOT}/frameworks/base/services/incremental
        ${ANDROID_ROOT}/frameworks/base/services/incremental/include
        ${ANDROID_ROOT}/frameworks/base/services/incremental/test
)

file (GLOB_RECURSE HEADERS
        ${ANDROID_ROOT}/frameworks/base/libs/services/include/android/content/*.h
        ${ANDROID_ROOT}/frameworks/base/libs/services/include/android/content/*.hpp
        ${ANDROID_ROOT}/frameworks/base/libs/services/include/android/os/*.h
        ${ANDROID_ROOT}/frameworks/base/libs/services/include/android/os/*.hpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/*.h
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/*.hpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/BroadcastRadio/*.h
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/BroadcastRadio/*.hpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/gnss/*.h
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/gnss/*.hpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/stats/*.h
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/stats/*.hpp
        ${ANDROID_ROOT}/frameworks/base/services/incremental/*.h
        ${ANDROID_ROOT}/frameworks/base/services/incremental/*.hpp
        ${ANDROID_ROOT}/frameworks/base/services/incremental/include/*.h
        ${ANDROID_ROOT}/frameworks/base/services/incremental/include/*.hpp
)

add_executable(
        android-services
        ${SOURCE_FILES}
        ${HEADERS}
)
