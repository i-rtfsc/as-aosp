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

project(debuglog)

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

#################### logkit ####################
set(FILES
        ${ANDROID_ROOT}/system/logging/logkit/main.cpp
        ${ANDROID_ROOT}/system/logging/logkit/logkit.cpp
        ${ANDROID_ROOT}/system/logging/logkit/common.cpp
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
    add_executable(logkit ${SOURCE_FILES})

    target_include_directories(logkit
            PUBLIC
            ${ANDROID_ROOT}/system/logging/logkit
            ${ANDROID_ROOT}/system/logging/logkit/binder
            ${OUT_INTERMEDIATES_ROOT}/system/logging/logkit/liblogreportaidl/android_${OUT_ARCH_CPU}_shared/gen/aidl
    )

    # Add c/c++ flags
    target_compile_options(logkit PRIVATE
            -Wall
            -Werror
            -Wno-unused-variable
            -Wextra
    )

    target_link_libraries(logkit
            libbase
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            liblog
            liblogreportaidl
            libm
            libprocessgroup
            libsysutils
            libutils
    )
endif ()
#################### logkit ####################

#################### logcat_plus ####################
set(FILES
        ${ANDROID_ROOT}/system/logging/logkit/logcat_plus.cpp
        ${ANDROID_ROOT}/system/logging/logkit/common.cpp
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
    add_executable(logcat_plus ${SOURCE_FILES})

    target_include_directories(logcat_plus
            PUBLIC
            ${ANDROID_ROOT}/system/logging/logkit
            ${ANDROID_ROOT}/system/logging/logkit/binder
            ${OUT_INTERMEDIATES_ROOT}/system/logging/logkit/liblogreportaidl/android_${OUT_ARCH_CPU}_shared/gen/aidl
    )

    # Add c/c++ flags
    target_compile_options(logcat_plus PRIVATE
            -Wall
            -Werror
            -Wno-unused-variable
            -Wextra
    )

    target_link_libraries(logcat_plus
            libbase
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            liblog
            liblogreportaidl
            libm
            libprocessgroup
            libsysutils
            libutils
    )
endif ()
#################### logcat_plus ####################

#################### klog ####################
set(FILES
        ${ANDROID_ROOT}/system/logging/logkit/klog.cpp
        ${ANDROID_ROOT}/system/logging/logkit/common.cpp
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
    add_executable(klog ${SOURCE_FILES})

    target_include_directories(klog
            PUBLIC
            ${ANDROID_ROOT}/system/logging/logkit
            ${ANDROID_ROOT}/system/logging/logkit/binder
            ${OUT_INTERMEDIATES_ROOT}/system/logging/logkit/liblogreportaidl/android_${OUT_ARCH_CPU}_shared/gen/aidl
    )

    # Add c/c++ flags
    target_compile_options(klog PRIVATE
            -Wall
            -Werror
            -Wno-unused-variable
            -Wextra
    )

    target_link_libraries(klog
            libbase
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            liblog
            liblogreportaidl
            libm
            libprocessgroup
            libsysutils
            libutils
    )
endif ()
#################### klog ####################

#################### debuglog ####################
set(FILES
        ${ANDROID_ROOT}/system/logging/logkit/debuglog.cpp
        ${ANDROID_ROOT}/system/logging/logkit/common.cpp
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
    add_executable(debuglog ${SOURCE_FILES})

    target_include_directories(debuglog
            PUBLIC
            ${ANDROID_ROOT}/system/logging/logkit
            ${ANDROID_ROOT}/system/logging/logkit/binder
            ${OUT_INTERMEDIATES_ROOT}/system/logging/logkit/liblogreportaidl/android_${OUT_ARCH_CPU}_shared/gen/aidl
    )

    # Add c/c++ flags
    target_compile_options(debuglog PRIVATE
            -Wall
            -Werror
            -Wno-unused-variable
            -Wextra
    )

    target_link_libraries(debuglog
            libbase
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            liblog
            liblogreportaidl
            libm
            libprocessgroup
            libsysutils
            libutils
    )
endif ()
#################### debuglog ####################

#################### liblogreportaidl ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/logging/logkit/liblogreportaidl/android_${OUT_ARCH_CPU}_shared/gen/aidl/android/os/ILogReportListener.cpp
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
    add_library(liblogreportaidl SHARED ${SOURCE_FILES})

    target_include_directories(liblogreportaidl
            PUBLIC
            ${ANDROID_ROOT}/system/logging/logkit/binder
            ${ANDROID_ROOT}/system/logging/logkit
            ${OUT_INTERMEDIATES_ROOT}/system/logging/logkit/liblogreportaidl/android_${OUT_ARCH_CPU}_shared/gen/aidl
    )

    # Add c/c++ flags
    target_compile_options(liblogreportaidl PRIVATE
            -Wall
            -Werror
            -Wno-unused-variable
            -Wextra
    )

    target_link_libraries(liblogreportaidl
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libutils
    )
endif ()
#################### liblogreportaidl ####################
