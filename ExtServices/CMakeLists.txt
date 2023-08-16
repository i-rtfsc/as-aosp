cmake_minimum_required(VERSION 3.5)
project(ExtServices)

# set(CMAKE_CXX_STANDARD 20)

#set(ANDROID_ROOT "/home/solo/code/aosp")
set(ANDROID_ROOT ${BUILD_NATIVE_ROOT})

set(AOSP_FWK false)
set(AOSP_SYSTEM_COMMON true)
set(AOSP_AV false)
set(AOSP_BIONIC false)
set(AOSP_SYSTEM false)
set(AOSP_EXTERNAL false)
set(AOSP_HARDWARE false)

file(GLOB SOURCE_FILES
        ${ANDROID_ROOT}/packages/modules/ExtServices/jni/*.cpp
        ${ANDROID_ROOT}/packages/modules/ExtServices/native/*.cpp
        ${ANDROID_ROOT}/packages/modules/ExtServices/native/pHash/*.cpp
        ${ANDROID_ROOT}/packages/modules/ExtServices/native/pHash/fuzzer/*.cpp
        ${ANDROID_ROOT}/packages/modules/ExtServices/native/tests/*.cpp
)

include_directories(
        "${ANDROID_ROOT}/packages/modules/ExtServices/jni"
        "${ANDROID_ROOT}/packages/modules/ExtServices/native"
        "${ANDROID_ROOT}/packages/modules/ExtServices/native/pHash/fuzzer"
        "${ANDROID_ROOT}/packages/modules/ExtServices/native/tests"
)

file (GLOB_RECURSE HEADERS
        "${ANDROID_ROOT}/packages/modules/ExtServices/jni/**/*.h"
        "${ANDROID_ROOT}/packages/modules/ExtServices/native/**/*.h"
        "${ANDROID_ROOT}/packages/modules/ExtServices/native/pHash/fuzzer/**/*.h"
        "${ANDROID_ROOT}/packages/modules/ExtServices/native/tests/**/*.h"
)


# add aosp framework
if (${AOSP_FKW})
    message("enable aosp framework")

    # add aosp framework header dir
    include_directories(
        "${ANDROID_ROOT}/frameworks/base/native/android/include_platform"
        "${ANDROID_ROOT}/frameworks/native/include"
        "${ANDROID_ROOT}/frameworks/native/libs/arect/include"
        "${ANDROID_ROOT}/frameworks/native/libs/nativebase/include"
        "${ANDROID_ROOT}/frameworks/native/libs/nativewindow/include"
        "${ANDROID_ROOT}/frameworks/native/libs/nativewindow/include-private"
        "${ANDROID_ROOT}/frameworks/native/opengl/include"
    )

    # add aosp framework header file
    file (GLOB_RECURSE HEADERS_FKW
        "${ANDROID_ROOT}/frameworks/base/native/android/include_platform/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/libs/arect/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/libs/nativebase/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/libs/nativewindow/include-private/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/libs/nativewindow/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/opengl/include/**/*.h"
    )
    # append aosp framework header
    list (APPEND HEADERS ${HEADERS_FKW})
endif ()


# add aosp system common
if (${AOSP_SYSTEM_COMMON})
    message("enable aosp system common")

    # add aosp system common header dir
    include_directories(
        "${ANDROID_ROOT}/libnativehelper/header_only_include"
        "${ANDROID_ROOT}/libnativehelper/include"
        "${ANDROID_ROOT}/libnativehelper/include_jni"
        "${ANDROID_ROOT}/system/core/include"
        "${ANDROID_ROOT}/system/core/property_service/libpropertyinfoparser/include"
        "${ANDROID_ROOT}/system/libbase/include"
        "${ANDROID_ROOT}/system/logging/liblog/include"
)

    # add aosp system common header file
    file (GLOB_RECURSE HEADERS_SYSTEM_COMMON
        "${ANDROID_ROOT}/libnativehelper/header_only_include/**/*.h"
        "${ANDROID_ROOT}/libnativehelper/include/**/*.h"
        "${ANDROID_ROOT}/libnativehelper/include_jni/**/*.h"
        "${ANDROID_ROOT}/system/core/include/**/*.h"
        "${ANDROID_ROOT}/system/core/property_service/libpropertyinfoparser/include/**/*.h"
        "${ANDROID_ROOT}/system/libbase/include/**/*.h"
        "${ANDROID_ROOT}/system/logging/liblog/include/**/*.h"
)
    # append aosp system common header
    list (APPEND HEADERS ${HEADERS_SYSTEM_COMMON})
endif ()


# add aosp av
if (${AOSP_AV})
    message("enable aosp av")

    # add aosp av header dir
    include_directories(
        "${ANDROID_ROOT}/frameworks/av/include"
)

    # add aosp av header file
    file (GLOB_RECURSE HEADERS_AV
        "${ANDROID_ROOT}/frameworks/av/include/**/*.h"
)
    # append aosp av header
    list (APPEND HEADERS ${HEADERS_AV})
endif ()



# add aosp bionic
if (${AOSP_BIONIC})
    message("enable aosp bionic")

    # add aosp bionic header dir
    include_directories(
        "${ANDROID_ROOT}/bionic/libc/async_safe/include"
        "${ANDROID_ROOT}/bionic/libc/include"
        "${ANDROID_ROOT}/bionic/libc/kernel/android/scsi"
        "${ANDROID_ROOT}/bionic/libc/kernel/android/uapi"
        "${ANDROID_ROOT}/bionic/libc/kernel/uapi"
        "${ANDROID_ROOT}/bionic/libc/kernel/uapi/asm-arm64"
        "${ANDROID_ROOT}/bionic/libc/system_properties/include"
)

    # add aosp bionic header file
    file (GLOB_RECURSE HEADERS_BIONIC
        "${ANDROID_ROOT}/bionic/libc/async_safe/include/**/*.h"
        "${ANDROID_ROOT}/bionic/libc/include/**/*.h"
        "${ANDROID_ROOT}/bionic/libc/kernel/android/scsi/**/*.h"
        "${ANDROID_ROOT}/bionic/libc/kernel/android/uapi/**/*.h"
        "${ANDROID_ROOT}/bionic/libc/kernel/uapi/**/*.h"
        "${ANDROID_ROOT}/bionic/libc/kernel/uapi/asm-arm64/**/*.h"
        "${ANDROID_ROOT}/bionic/libc/system_properties/include/**/*.h"
)
    # append aosp bionic header
    list (APPEND HEADERS ${HEADERS_BIONIC})
endif ()


# add aosp system
if (${AOSP_SYSTEM})
    message("enable aosp system")

    # add aosp system header dir
    include_directories(
        "${ANDROID_ROOT}/system/media/audio/include"
)

    # add aosp system header file
    file (GLOB_RECURSE HEADERS_SYSTEM
        "${ANDROID_ROOT}/system/media/audio/include/**/*.h"
)
    # append aosp system header
    list (APPEND HEADERS ${HEADERS_SYSTEM})
endif ()


# add aosp external
if (${AOSP_EXTERNAL})
    message("enable aosp external")

    # add aosp external header dir
    include_directories(
        "${ANDROID_ROOT}/external/fft2d/src/fft2d"
        "${ANDROID_ROOT}/external/fmtlib/include"
        "${ANDROID_ROOT}/external/googletest/googlemock/include"
        "${ANDROID_ROOT}/external/googletest/googletest/include"
        "${ANDROID_ROOT}/external/libcxx/include"
        "${ANDROID_ROOT}/external/libcxxabi/include"
)

    # add aosp external header file
    file (GLOB_RECURSE HEADERS_EXTERNAL
        "${ANDROID_ROOT}/external/fft2d/src/fft2d/**/*.h"
        "${ANDROID_ROOT}/external/fmtlib/include/**/*.h"
        "${ANDROID_ROOT}/external/googletest/googlemock/include/**/*.h"
        "${ANDROID_ROOT}/external/googletest/googletest/include/**/*.h"
        "${ANDROID_ROOT}/external/libcxx/include/**/*.h"
        "${ANDROID_ROOT}/external/libcxxabi/include/**/*.h"
)
    # append aosp external header
    list (APPEND HEADERS ${HEADERS_EXTERNAL})
endif ()


# add aosp hardware
if (${AOSP_HARDWARE})
    message("enable aosp hardware")

    # add aosp hardware header dir
    include_directories(
        "${ANDROID_ROOT}/hardware/libhardware/include"
        "${ANDROID_ROOT}/hardware/libhardware_legacy/include"
        "${ANDROID_ROOT}/hardware/ril/include"
)

    # add aosp hardware header file
    file (GLOB_RECURSE HEADERS_HARDWARE
        "${ANDROID_ROOT}/hardware/libhardware/include/**/*.h"
        "${ANDROID_ROOT}/hardware/libhardware_legacy/include/**/*.h"
        "${ANDROID_ROOT}/hardware/ril/include/**/*.h"
)
    # append aosp hardware header
    list (APPEND HEADERS ${HEADERS_HARDWARE})
endif ()


add_executable(
        ExtServices
        ${SOURCE_FILES}
        ${HEADERS}
)
