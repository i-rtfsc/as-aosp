# Copyright (c) 2024 anqi.huang@outlook.com
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use self file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

cmake_minimum_required(VERSION 3.5)

project(libmediautils)

# Check if the major CMakeLists.txt file has correctly set
# the ANDROID_ROOT directory.
if (IS_DIRECTORY ${ANDROID_ROOT})
    message(STATUS "${PROJECT_NAME}, has been set android root")
else ()
    if (IS_DIRECTORY ${BUILD_NATIVE_ROOT})
        message(WARNING "${PROJECT_NAME}, set ANDROID_ROOT via BUILD_NATIVE_ROOT=${BUILD_NATIVE_ROOT}")
        set(ANDROID_ROOT ${BUILD_NATIVE_ROOT})
        set(OUT_INTERMEDIATES_ROOT ${ANDROID_ROOT}/out/soong/.intermediates)
    else ()
        message(FATAL_ERROR "${PROJECT_NAME}, ANDROID_ROOT was NULL")
    endif ()
endif ()

#################### libmediautils ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/utils/AImageReaderUtils.cpp
        ${ANDROID_ROOT}/frameworks/av/media/utils/BatteryNotifier.cpp
        ${ANDROID_ROOT}/frameworks/av/media/utils/ISchedulingPolicyService.cpp
        ${ANDROID_ROOT}/frameworks/av/media/utils/Library.cpp
        ${ANDROID_ROOT}/frameworks/av/media/utils/MediaUtilsDelayed.cpp
        ${ANDROID_ROOT}/frameworks/av/media/utils/MethodStatistics.cpp
        ${ANDROID_ROOT}/frameworks/av/media/utils/Process.cpp
        ${ANDROID_ROOT}/frameworks/av/media/utils/ProcessInfo.cpp
        ${ANDROID_ROOT}/frameworks/av/media/utils/SchedulingPolicyService.cpp
        ${ANDROID_ROOT}/frameworks/av/media/utils/ServiceUtilities.cpp
        ${ANDROID_ROOT}/frameworks/av/media/utils/ThreadSnapshot.cpp
        ${ANDROID_ROOT}/frameworks/av/media/utils/TimeCheck.cpp
        ${ANDROID_ROOT}/frameworks/av/media/utils/TimerThread.cpp
        ${ANDROID_ROOT}/frameworks/av/media/utils/LimitProcessMemory.cpp
        ${ANDROID_ROOT}/frameworks/av/media/utils/MemoryLeakTrackUtil.cpp
)
set(SOURCE_FILES "")
foreach (file ${FILES})
    if (EXISTS ${file})
        list(APPEND SOURCE_FILES ${file})
    endif ()
endforeach ()

if (NOT SOURCE_FILES)
    message(WARNING "${PROJECT_NAME}, source files not exist")
else ()
    add_library(libmediautils SHARED ${SOURCE_FILES})

    target_include_directories(libmediautils
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/utils/include
            ${ANDROID_ROOT}/frameworks/av/media/utils
    )

    # Add c/c++ flags
    target_compile_options(libmediautils PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(libmediautils
            android.hardware.graphics.bufferqueue-1.0
            android.hidl.token-1.0-utils
            bionic_libc_platform_headers
            libaudioclient_aidl_conversion
            libaudioutils
            libbase
            libbatterystats_aidl
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libc_malloc_debug_backtrace
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libhidlbase
            liblog
            libm
            libpermission
            libprocessinfoservice_aidl
            libutils
            packagemanager_aidl-cpp
    )
endif ()
#################### libmediautils ####################

#################### libmediautils_delayed ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/utils/MediaUtilsDelayedLibrary.cpp
)
set(SOURCE_FILES "")
foreach (file ${FILES})
    if (EXISTS ${file})
        list(APPEND SOURCE_FILES ${file})
    endif ()
endforeach ()

if (NOT SOURCE_FILES)
    message(WARNING "${PROJECT_NAME}, source files not exist")
else ()
    add_library(libmediautils_delayed SHARED ${SOURCE_FILES})

    target_include_directories(libmediautils_delayed
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/utils
    )

    # Add c/c++ flags
    target_compile_options(libmediautils_delayed PRIVATE
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(libmediautils_delayed
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            liblog
            libm
            libutils
            libutilscallstack
    )
endif ()
#################### libmediautils_delayed ####################

#################### libmediautils_vendor ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/utils/MemoryLeakTrackUtil.cpp
)
set(SOURCE_FILES "")
foreach (file ${FILES})
    if (EXISTS ${file})
        list(APPEND SOURCE_FILES ${file})
    endif ()
endforeach ()

if (NOT SOURCE_FILES)
    message(WARNING "${PROJECT_NAME}, source files not exist")
else ()
    add_library(libmediautils_vendor SHARED ${SOURCE_FILES})

    target_include_directories(libmediautils_vendor
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/utils/include
            ${ANDROID_ROOT}/frameworks/av/media/utils
    )

    # Add c/c++ flags
    target_compile_options(libmediautils_vendor PRIVATE
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(libmediautils_vendor
            bionic_libc_platform_headers
            libc
            libc++
            libc++_static
            libc++demangle
            libc_malloc_debug_backtrace
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            liblog
            libm
            libutils
    )
endif ()
#################### libmediautils_vendor ####################
