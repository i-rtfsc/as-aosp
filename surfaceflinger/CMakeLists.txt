cmake_minimum_required(VERSION 3.5)
project(surfaceflinger)

set(ANDROID_ROOT ${BUILD_NATIVE_ROOT})

file(GLOB SOURCE_FILES
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/*.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/**/*.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/Scheduler/*.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/Scheduler/**/*.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/CompositionEngine/src/*.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/CompositionEngine/src/**/*.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/CompositionEngine/mock/*.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/CompositionEngine/tests/*.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/CompositionEngine/tests/**/*.cpp
)

include_directories(
        "${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/CompositionEngine/include"
        "${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests"
        "${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/fuzzer"
        "${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/CompositionEngine/include"
        "${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/FrameTimeline"
        "${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/Scheduler/include"
        "${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/Scheduler"
        "${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/layerproto/include"
        "${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/TimeStats"
        "${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/TimeStats/timestatsatomsproto/include"
        "${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/TimeStats/timestatsproto/include"
        "${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/layerproto/include"
        "${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/layerproto"
        "${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/sysprop"
        "${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/Tracing/tools"
)


file (GLOB_RECURSE HEADERS
        "${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/CompositionEngine/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/CompositionEngine/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/FrameTimeline/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/fuzzer/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/Scheduler/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/Scheduler/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/layerproto/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/layerproto/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/TimeStats/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/TimeStats/timestatsproto/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/TimeStats/timestatsatomsproto/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/sysprop/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/Tracing/tools/**/*.h"
)


add_executable(
        surfaceflinger
        ${SOURCE_FILES}
        ${HEADERS}
)