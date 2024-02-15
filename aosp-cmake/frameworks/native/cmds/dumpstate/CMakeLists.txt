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

project(dumpstate)

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

#################### libdumpstateutil ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/cmds/dumpstate/DumpstateInternal.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/dumpstate/DumpstateUtil.cpp
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
    add_library(libdumpstateutil SHARED ${SOURCE_FILES})

    target_include_directories(libdumpstateutil
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/cmds/dumpstate
    )

    # Add c/c++ flags
    target_compile_options(libdumpstateutil PRIVATE
            -Wall
            -Werror
            -Wno-unused-variable
    )

    target_link_libraries(libdumpstateutil
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            liblog
            libm
    )
endif ()
#################### libdumpstateutil ####################

#################### dumpstate ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/cmds/dumpstate/DumpstateService.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/dumpstate/DumpPool.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/dumpstate/TaskQueue.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/dumpstate/dumpstate.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/dumpstate/main.cpp
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
    add_executable(dumpstate ${SOURCE_FILES})

    target_include_directories(dumpstate
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/cmds/dumpstate
            ${ANDROID_ROOT}/frameworks/native/cmds/dumpstate/binder
            ${OUT_INTERMEDIATES_ROOT}/frameworks/native/cmds/dumpstate/libdumpstateaidl/android_${OUT_ARCH_CPU}_shared/gen/aidl
    )

    # Add c/c++ flags
    target_compile_options(dumpstate PRIVATE
            -Wall
            -Werror
            -Wno-unused-variable
    )

    target_link_libraries(dumpstate
            android.hardware.dumpstate-V1-ndk
            android.hardware.dumpstate-1.0
            android.hardware.dumpstate-1.1
            libbase
            libbinder
            libbinder_ndk
            libbinderdebug
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcrypto
            libcutils
            libdebuggerd_client
            libdl
            libdumpstateaidl
            libdumpstateutil
            libdumpsys
            libdumputils
            libhardware_legacy
            libhidlbase
            libincidentcompanion
            liblog
            libm
            libserviceutils
            libutils
            libvintf
            libziparchive
            packagemanager_aidl-cpp
    )
endif ()
#################### dumpstate ####################

#################### libdumpstateaidl ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/cmds/dumpstate/libdumpstateaidl/android_${OUT_ARCH_CPU}_shared/gen/aidl/android/os/IDumpstateListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/cmds/dumpstate/libdumpstateaidl/android_${OUT_ARCH_CPU}_shared/gen/aidl/android/os/IDumpstate.cpp
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
    add_library(libdumpstateaidl SHARED ${SOURCE_FILES})

    target_include_directories(libdumpstateaidl
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/cmds/dumpstate/binder
            ${ANDROID_ROOT}/frameworks/native/cmds/dumpstate
            ${OUT_INTERMEDIATES_ROOT}/frameworks/native/cmds/dumpstate/libdumpstateaidl/android_${OUT_ARCH_CPU}_shared/gen/aidl
    )

    # Add c/c++ flags
    target_compile_options(libdumpstateaidl PRIVATE
            -Wall
            -Werror
            -Wno-unused-variable
    )

    target_link_libraries(libdumpstateaidl
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
#################### libdumpstateaidl ####################
