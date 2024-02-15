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

project(auditctl)

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

#################### replay_messages ####################
set(FILES
        ${ANDROID_ROOT}/system/logging/logd/ReplayMessages.cpp
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
    add_executable(replay_messages ${SOURCE_FILES})

    target_include_directories(replay_messages
            PUBLIC
            ${ANDROID_ROOT}/system/logging/logd
    )

    # Add c/c++ flags
    target_compile_options(replay_messages PRIVATE
            -Werror
            -Wextra
    )

    target_link_libraries(replay_messages
            libbase
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libcutils_headers
            libdl
            liblog
            liblogd
            libm
            libselinux
            libutils
            libz
            libzstd
    )
endif ()
#################### replay_messages ####################

#################### liblogd ####################
set(FILES
        ${ANDROID_ROOT}/system/logging/logd/CompressionEngine.cpp
        ${ANDROID_ROOT}/system/logging/logd/LogBufferElement.cpp
        ${ANDROID_ROOT}/system/logging/logd/LogReaderList.cpp
        ${ANDROID_ROOT}/system/logging/logd/LogReaderThread.cpp
        ${ANDROID_ROOT}/system/logging/logd/LogSize.cpp
        ${ANDROID_ROOT}/system/logging/logd/LogStatistics.cpp
        ${ANDROID_ROOT}/system/logging/logd/LogTags.cpp
        ${ANDROID_ROOT}/system/logging/logd/LogdLock.cpp
        ${ANDROID_ROOT}/system/logging/logd/PruneList.cpp
        ${ANDROID_ROOT}/system/logging/logd/SerializedFlushToState.cpp
        ${ANDROID_ROOT}/system/logging/logd/SerializedLogBuffer.cpp
        ${ANDROID_ROOT}/system/logging/logd/SerializedLogChunk.cpp
        ${ANDROID_ROOT}/system/logging/logd/SimpleLogBuffer.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/logging/logd/liblogd/android_${OUT_ARCH_CPU}_static/gen/aidl/android/os/logcat/ILogcatManagerService.cpp
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
    add_library(liblogd SHARED ${SOURCE_FILES})

    target_include_directories(liblogd
            PUBLIC
            ${ANDROID_ROOT}/system/logging/logd
            ${OUT_INTERMEDIATES_ROOT}/system/logging/logd/liblogd/android_${OUT_ARCH_CPU}_static/gen/aidl
    )

    # Add c/c++ flags
    target_compile_options(liblogd PRIVATE
            -Werror
            -Wextra
    )

    target_link_libraries(liblogd
            libbase
            libbinder
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcutils_headers
            libdl
            liblog
            libm
            libz
            libzstd
    )
endif ()
#################### liblogd ####################

#################### logd ####################
set(FILES
        ${ANDROID_ROOT}/system/logging/logd/libaudit.cpp
        ${ANDROID_ROOT}/system/logging/logd/main.cpp
        ${ANDROID_ROOT}/system/logging/logd/CommandListener.cpp
        ${ANDROID_ROOT}/system/logging/logd/LogdNativeService.cpp
        ${ANDROID_ROOT}/system/logging/logd/LogAudit.cpp
        ${ANDROID_ROOT}/system/logging/logd/LogKlog.cpp
        ${ANDROID_ROOT}/system/logging/logd/LogListener.cpp
        ${ANDROID_ROOT}/system/logging/logd/LogPermissions.cpp
        ${ANDROID_ROOT}/system/logging/logd/LogReader.cpp
        ${ANDROID_ROOT}/system/logging/logd/TrustyLog.cpp
        ${ANDROID_ROOT}/system/logging/logd/PkgIds.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/logging/logd/logd/android_${OUT_ARCH_CPU}_cfi/gen/aidl/android/os/logcat/ILogcatManagerService.cpp
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
    add_executable(logd ${SOURCE_FILES})

    target_include_directories(logd
            PUBLIC
            ${ANDROID_ROOT}/system/logging/logd
            ${OUT_INTERMEDIATES_ROOT}/system/logging/logd/logd/android_${OUT_ARCH_CPU}_cfi/gen/aidl
            ${OUT_INTERMEDIATES_ROOT}/system/logging/logd/liblogd_binder/android_${OUT_ARCH_CPU}_static_cfi/gen/aidl
    )

    # Add c/c++ flags
    target_compile_options(logd PRIVATE
            -Werror
            -Wextra
    )

    target_link_libraries(logd
            libbase
            libbinder
            libc
            libc++
            libc++demangle
            libcap
            libclang_rt.builtins
            libcutils
            libcutils_headers
            libdl
            liblog
            liblogd
            liblogd_binder
            libm
            libpackagelistparser
            libprocessgroup
            libsysutils
            libutils
            libz
            libzstd
    )
endif ()
#################### logd ####################

#################### liblogd_binder ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/logging/logd/liblogd_binder/android_${OUT_ARCH_CPU}_static/gen/aidl/android/os/ILogd.cpp
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
    add_library(liblogd_binder SHARED ${SOURCE_FILES})

    target_include_directories(liblogd_binder
            PUBLIC
            ${ANDROID_ROOT}/system/logging/logd
            ${OUT_INTERMEDIATES_ROOT}/system/logging/logd/liblogd_binder/android_${OUT_ARCH_CPU}_static/gen/aidl
    )

    # Add c/c++ flags
    target_compile_options(liblogd_binder PRIVATE
            -Werror
            -Wextra
    )

    target_link_libraries(liblogd_binder
            libbase
            libbinder
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcutils_headers
            libdl
            libincremental_aidl-cpp
            libm
            libutils
            libz
            libzstd
    )
endif ()
#################### liblogd_binder ####################

#################### auditctl ####################
set(FILES
        ${ANDROID_ROOT}/system/logging/logd/auditctl.cpp
        ${ANDROID_ROOT}/system/logging/logd/libaudit.cpp
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
    add_executable(auditctl ${SOURCE_FILES})

    target_include_directories(auditctl
            PUBLIC
            ${ANDROID_ROOT}/system/logging/logd
    )

    # Add c/c++ flags
    target_compile_options(auditctl PRIVATE
            -Werror
            -Wextra
    )

    target_link_libraries(auditctl
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### auditctl ####################
