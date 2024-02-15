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

project(libnetdutils)

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

#################### libnetdutils ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/libs/net/common/netd/libnetdutils/DumpWriter.cpp
        ${ANDROID_ROOT}/frameworks/libs/net/common/netd/libnetdutils/Fd.cpp
        ${ANDROID_ROOT}/frameworks/libs/net/common/netd/libnetdutils/InternetAddresses.cpp
        ${ANDROID_ROOT}/frameworks/libs/net/common/netd/libnetdutils/Log.cpp
        ${ANDROID_ROOT}/frameworks/libs/net/common/netd/libnetdutils/Netfilter.cpp
        ${ANDROID_ROOT}/frameworks/libs/net/common/netd/libnetdutils/Netlink.cpp
        ${ANDROID_ROOT}/frameworks/libs/net/common/netd/libnetdutils/NetlinkListener.cpp
        ${ANDROID_ROOT}/frameworks/libs/net/common/netd/libnetdutils/Slice.cpp
        ${ANDROID_ROOT}/frameworks/libs/net/common/netd/libnetdutils/Socket.cpp
        ${ANDROID_ROOT}/frameworks/libs/net/common/netd/libnetdutils/SocketOption.cpp
        ${ANDROID_ROOT}/frameworks/libs/net/common/netd/libnetdutils/Status.cpp
        ${ANDROID_ROOT}/frameworks/libs/net/common/netd/libnetdutils/Syscalls.cpp
        ${ANDROID_ROOT}/frameworks/libs/net/common/netd/libnetdutils/UniqueFd.cpp
        ${ANDROID_ROOT}/frameworks/libs/net/common/netd/libnetdutils/UniqueFile.cpp
        ${ANDROID_ROOT}/frameworks/libs/net/common/netd/libnetdutils/Utils.cpp
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
    add_library(libnetdutils SHARED ${SOURCE_FILES})

    target_include_directories(libnetdutils
            PUBLIC
            ${ANDROID_ROOT}/frameworks/libs/net/common/netd/libnetdutils/include
            ${ANDROID_ROOT}/frameworks/libs/net/common/netd/libnetdutils
    )

    # Add c/c++ flags
    target_compile_options(libnetdutils PRIVATE
            -Wall
            -Werror
            -Werror=implicit-fallthrough
            -Werror=sometimes-uninitialized
            -Werror=conditional-uninitialized
            -Wall
            -Werror
    )

    target_link_libraries(libnetdutils
            libbase
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
#################### libnetdutils ####################
