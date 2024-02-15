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

project(android.system.suspend.control-cpp-analyzer)

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

#################### android.system.suspend-V2-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/aidl/android.system.suspend-V2-ndk-source/gen/android/system/suspend/ISystemSuspend.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/aidl/android.system.suspend-V2-ndk-source/gen/android/system/suspend/IWakeLock.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/aidl/android.system.suspend-V2-ndk-source/gen/android/system/suspend/WakeLockType.cpp
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
    add_library(android.system.suspend-V2-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.system.suspend-V2-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/aidl/android.system.suspend-V2-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.system.suspend-V2-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.system.suspend-V2-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### android.system.suspend-V2-ndk ####################

#################### android.system.suspend-V2-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/aidl/android.system.suspend-V2-ndk_platform-source/gen/android/system/suspend/ISystemSuspend.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/aidl/android.system.suspend-V2-ndk_platform-source/gen/android/system/suspend/IWakeLock.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/aidl/android.system.suspend-V2-ndk_platform-source/gen/android/system/suspend/WakeLockType.cpp
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
    add_library(android.system.suspend-V2-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.system.suspend-V2-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/aidl/android.system.suspend-V2-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.system.suspend-V2-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.system.suspend-V2-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### android.system.suspend-V2-ndk_platform ####################

#################### android.system.suspend.control-V2-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/aidl/android.system.suspend.control-V2-ndk_platform-source/gen/android/system/suspend/ISuspendControlService.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/aidl/android.system.suspend.control-V2-ndk_platform-source/gen/android/system/suspend/ISuspendCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/aidl/android.system.suspend.control-V2-ndk_platform-source/gen/android/system/suspend/IWakelockCallback.cpp
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
    add_library(android.system.suspend.control-V2-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.system.suspend.control-V2-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/aidl/android.system.suspend.control-V2-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.system.suspend.control-V2-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.system.suspend.control-V2-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### android.system.suspend.control-V2-ndk_platform ####################

#################### android.system.suspend.control.internal-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/aidl/android.system.suspend.control.internal-cpp-source/gen/android/system/suspend/internal/ISuspendControlServiceInternal.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/aidl/android.system.suspend.control.internal-cpp-source/gen/android/system/suspend/internal/SuspendInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/aidl/android.system.suspend.control.internal-cpp-source/gen/android/system/suspend/internal/WakeLockInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/aidl/android.system.suspend.control.internal-cpp-source/gen/android/system/suspend/internal/WakeupInfo.cpp
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
    add_library(android.system.suspend.control.internal-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.system.suspend.control.internal-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/aidl/android.system.suspend.control.internal-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.system.suspend.control.internal-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.system.suspend.control.internal-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libm
            libutils
    )
endif ()
#################### android.system.suspend.control.internal-cpp ####################

#################### android.system.suspend.control.internal-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/aidl/android.system.suspend.control.internal-ndk_platform-source/gen/android/system/suspend/internal/ISuspendControlServiceInternal.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/aidl/android.system.suspend.control.internal-ndk_platform-source/gen/android/system/suspend/internal/SuspendInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/aidl/android.system.suspend.control.internal-ndk_platform-source/gen/android/system/suspend/internal/WakeLockInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/aidl/android.system.suspend.control.internal-ndk_platform-source/gen/android/system/suspend/internal/WakeupInfo.cpp
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
    add_library(android.system.suspend.control.internal-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.system.suspend.control.internal-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/aidl/android.system.suspend.control.internal-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.system.suspend.control.internal-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.system.suspend.control.internal-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### android.system.suspend.control.internal-ndk_platform ####################

#################### android.system.suspend.control.internal-cpp-analyzer ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/aidl/android.system.suspend.control.internal-cpp-analyzer-source/gen/android/system/suspend/internal/ISuspendControlServiceInternal.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/aidl/android.system.suspend.control.internal-cpp-analyzer-source/gen/android/system/suspend/internal/SuspendInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/aidl/android.system.suspend.control.internal-cpp-analyzer-source/gen/android/system/suspend/internal/WakeLockInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/aidl/android.system.suspend.control.internal-cpp-analyzer-source/gen/android/system/suspend/internal/WakeupInfo.cpp
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
    add_executable(android.system.suspend.control.internal-cpp-analyzer ${SOURCE_FILES})

    target_include_directories(android.system.suspend.control.internal-cpp-analyzer
            PUBLIC
            ${ANDROID_ROOT}/system/hardware/interfaces/suspend/aidl
            ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/aidl/android.system.suspend.control.internal-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/aidl/android.system.suspend.control.internal-cpp-analyzer-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.system.suspend.control.internal-cpp-analyzer PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.system.suspend.control.internal-cpp-analyzer
            aidl-analyzer-main
            android.system.suspend.control.internal-cpp
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
#################### android.system.suspend.control.internal-cpp-analyzer ####################

#################### android.system.suspend.control.internal-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/aidl/android.system.suspend.control.internal-ndk-source/gen/android/system/suspend/internal/ISuspendControlServiceInternal.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/aidl/android.system.suspend.control.internal-ndk-source/gen/android/system/suspend/internal/SuspendInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/aidl/android.system.suspend.control.internal-ndk-source/gen/android/system/suspend/internal/WakeLockInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/aidl/android.system.suspend.control.internal-ndk-source/gen/android/system/suspend/internal/WakeupInfo.cpp
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
    add_library(android.system.suspend.control.internal-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.system.suspend.control.internal-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/aidl/android.system.suspend.control.internal-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.system.suspend.control.internal-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.system.suspend.control.internal-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### android.system.suspend.control.internal-ndk ####################

#################### android.system.suspend-V1-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/aidl/android.system.suspend-V1-ndk-source/gen/android/system/suspend/ISystemSuspend.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/aidl/android.system.suspend-V1-ndk-source/gen/android/system/suspend/IWakeLock.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/aidl/android.system.suspend-V1-ndk-source/gen/android/system/suspend/WakeLockType.cpp
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
    add_library(android.system.suspend-V1-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.system.suspend-V1-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/aidl/android.system.suspend-V1-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.system.suspend-V1-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.system.suspend-V1-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### android.system.suspend-V1-ndk ####################

#################### android.system.suspend.control-V2-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/aidl/android.system.suspend.control-V2-ndk-source/gen/android/system/suspend/ISuspendControlService.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/aidl/android.system.suspend.control-V2-ndk-source/gen/android/system/suspend/ISuspendCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/aidl/android.system.suspend.control-V2-ndk-source/gen/android/system/suspend/IWakelockCallback.cpp
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
    add_library(android.system.suspend.control-V2-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.system.suspend.control-V2-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/aidl/android.system.suspend.control-V2-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.system.suspend.control-V2-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.system.suspend.control-V2-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### android.system.suspend.control-V2-ndk ####################

#################### android.system.suspend-V1-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/aidl/android.system.suspend-V1-ndk_platform-source/gen/android/system/suspend/ISystemSuspend.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/aidl/android.system.suspend-V1-ndk_platform-source/gen/android/system/suspend/IWakeLock.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/aidl/android.system.suspend-V1-ndk_platform-source/gen/android/system/suspend/WakeLockType.cpp
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
    add_library(android.system.suspend-V1-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.system.suspend-V1-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/aidl/android.system.suspend-V1-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.system.suspend-V1-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.system.suspend-V1-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### android.system.suspend-V1-ndk_platform ####################

#################### android.system.suspend.control-V1-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/aidl/android.system.suspend.control-V1-ndk_platform-source/gen/android/system/suspend/ISuspendCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/aidl/android.system.suspend.control-V1-ndk_platform-source/gen/android/system/suspend/ISuspendControlService.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/aidl/android.system.suspend.control-V1-ndk_platform-source/gen/android/system/suspend/IWakelockCallback.cpp
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
    add_library(android.system.suspend.control-V1-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.system.suspend.control-V1-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/aidl/android.system.suspend.control-V1-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.system.suspend.control-V1-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.system.suspend.control-V1-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### android.system.suspend.control-V1-ndk_platform ####################

#################### android.system.suspend.control-cpp-analyzer ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/aidl/android.system.suspend.control-cpp-analyzer-source/gen/android/system/suspend/ISuspendControlService.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/aidl/android.system.suspend.control-cpp-analyzer-source/gen/android/system/suspend/ISuspendCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/aidl/android.system.suspend.control-cpp-analyzer-source/gen/android/system/suspend/IWakelockCallback.cpp
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
    add_executable(android.system.suspend.control-cpp-analyzer ${SOURCE_FILES})

    target_include_directories(android.system.suspend.control-cpp-analyzer
            PUBLIC
            ${ANDROID_ROOT}/system/hardware/interfaces/suspend/aidl
            ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/aidl/android.system.suspend.control-V2-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/aidl/android.system.suspend.control-cpp-analyzer-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.system.suspend.control-cpp-analyzer PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.system.suspend.control-cpp-analyzer
            aidl-analyzer-main
            android.system.suspend.control-V2-cpp
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
#################### android.system.suspend.control-cpp-analyzer ####################

#################### android.system.suspend.control-V1-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/aidl/android.system.suspend.control-V1-cpp-source/gen/android/system/suspend/ISuspendCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/aidl/android.system.suspend.control-V1-cpp-source/gen/android/system/suspend/ISuspendControlService.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/aidl/android.system.suspend.control-V1-cpp-source/gen/android/system/suspend/IWakelockCallback.cpp
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
    add_library(android.system.suspend.control-V1-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.system.suspend.control-V1-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/aidl/android.system.suspend.control-V1-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.system.suspend.control-V1-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.system.suspend.control-V1-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libm
            libutils
    )
endif ()
#################### android.system.suspend.control-V1-cpp ####################

#################### android.system.suspend.control-V2-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/aidl/android.system.suspend.control-V2-cpp-source/gen/android/system/suspend/ISuspendControlService.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/aidl/android.system.suspend.control-V2-cpp-source/gen/android/system/suspend/ISuspendCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/aidl/android.system.suspend.control-V2-cpp-source/gen/android/system/suspend/IWakelockCallback.cpp
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
    add_library(android.system.suspend.control-V2-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.system.suspend.control-V2-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/aidl/android.system.suspend.control-V2-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.system.suspend.control-V2-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.system.suspend.control-V2-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libm
            libutils
    )
endif ()
#################### android.system.suspend.control-V2-cpp ####################

#################### android.system.suspend.control-V1-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/aidl/android.system.suspend.control-V1-ndk-source/gen/android/system/suspend/ISuspendCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/aidl/android.system.suspend.control-V1-ndk-source/gen/android/system/suspend/ISuspendControlService.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/aidl/android.system.suspend.control-V1-ndk-source/gen/android/system/suspend/IWakelockCallback.cpp
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
    add_library(android.system.suspend.control-V1-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.system.suspend.control-V1-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/suspend/aidl/android.system.suspend.control-V1-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.system.suspend.control-V1-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.system.suspend.control-V1-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### android.system.suspend.control-V1-ndk ####################
