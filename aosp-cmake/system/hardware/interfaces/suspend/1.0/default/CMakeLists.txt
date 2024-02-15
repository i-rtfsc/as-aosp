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

project(android.system.suspend-service)

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

#################### android.system.suspend-service ####################
set(FILES
        ${ANDROID_ROOT}/system/hardware/interfaces/suspend/1.0/default/main.cpp
        ${ANDROID_ROOT}/system/hardware/interfaces/suspend/1.0/default/SuspendControlService.cpp
        ${ANDROID_ROOT}/system/hardware/interfaces/suspend/1.0/default/SystemSuspend.cpp
        ${ANDROID_ROOT}/system/hardware/interfaces/suspend/1.0/default/SystemSuspendHidl.cpp
        ${ANDROID_ROOT}/system/hardware/interfaces/suspend/1.0/default/SystemSuspendAidl.cpp
        ${ANDROID_ROOT}/system/hardware/interfaces/suspend/1.0/default/WakeLockEntryList.cpp
        ${ANDROID_ROOT}/system/hardware/interfaces/suspend/1.0/default/WakeupList.cpp
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
    add_executable(android.system.suspend-service ${SOURCE_FILES})

    target_include_directories(android.system.suspend-service
            PUBLIC
            ${ANDROID_ROOT}/system/hardware/interfaces/suspend/1.0/default
            ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/1.0/default/libSuspendProperties/android_${OUT_ARCH_CPU}_static/gen/sysprop/include
            ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/1.0/default/libSuspendProperties/android_${OUT_ARCH_CPU}_shared/gen/sysprop/include
    )

    # Add c/c++ flags
    target_compile_options(android.system.suspend-service PRIVATE
            -Wall
            -Werror
            -Wall
            -Werror
    )

    target_link_libraries(android.system.suspend-service
            android.system.suspend-V1-ndk
            android.system.suspend.control-V1-cpp
            android.system.suspend.control.internal-cpp
            android.system.suspend-1.0
            libSuspendProperties
            libbase
            libbinder
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libhidlbase
            liblog
            libm
            libutils
    )
endif ()
#################### android.system.suspend-service ####################

#################### libSuspendProperties ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/1.0/default/libSuspendProperties/android_${OUT_ARCH_CPU}_static/gen/sysprop/SuspendProperties.sysprop.cpp
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
    add_library(libSuspendProperties SHARED ${SOURCE_FILES})

    target_include_directories(libSuspendProperties
            PUBLIC
            ${ANDROID_ROOT}/system/hardware/interfaces/suspend/1.0/default
            ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/1.0/default/libSuspendProperties/android_${OUT_ARCH_CPU}_static/gen/sysprop/include
    )

    # Add c/c++ flags
    target_compile_options(libSuspendProperties PRIVATE
            -Werror
    )

    target_link_libraries(libSuspendProperties
            libbase_headers
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            liblog
            libm
    )
endif ()
#################### libSuspendProperties ####################
