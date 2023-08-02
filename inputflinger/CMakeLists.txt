cmake_minimum_required(VERSION 3.5)
project(inputflinger)

set(ANDROID_ROOT ${BUILD_NATIVE_ROOT})

file(GLOB SOURCE_FILES
        ${ANDROID_ROOT}/frameworks/base/libs/input/*.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/input/**/*.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/**/*.cpp
)

include_directories(
        "${ANDROID_ROOT}/frameworks/base/libs/input"
        "${ANDROID_ROOT}/frameworks/native/services/inputflinger"
        "${ANDROID_ROOT}/frameworks/native/services/inputflinger/include"
        "${ANDROID_ROOT}/frameworks/native/services/inputflinger/dispatcher/include"
        "${ANDROID_ROOT}/frameworks/native/services/inputflinger/benchmarks"
        "${ANDROID_ROOT}/frameworks/native/services/inputflinger/reporter"
        "${ANDROID_ROOT}/frameworks/native/services/inputflinger/host"
        "${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader"
        "${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/controller"
        "${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/include"
        "${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper"
        "${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper/accumulator"

)


file (GLOB_RECURSE HEADERS
        "${ANDROID_ROOT}/frameworks/base/libs/input/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/services/inputflinger/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/services/inputflinger/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/services/inputflinger/benchmarks/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/services/inputflinger/dispatcher/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/services/inputflinger/reporter/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/services/inputflinger/host/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/controller/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper/accumulator/**/*.h"
)


add_executable(
        inputflinger
        ${SOURCE_FILES}
        ${HEADERS}
)