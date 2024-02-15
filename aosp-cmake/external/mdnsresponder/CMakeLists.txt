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

project(dnssd)

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

#################### mdnsd ####################
set(FILES
        ${ANDROID_ROOT}/external/mdnsresponder/mDNSCore/DNSCommon.c
        ${ANDROID_ROOT}/external/mdnsresponder/mDNSCore/DNSDigest.c
        ${ANDROID_ROOT}/external/mdnsresponder/mDNSCore/mDNS.c
        ${ANDROID_ROOT}/external/mdnsresponder/mDNSCore/uDNS.c
        ${ANDROID_ROOT}/external/mdnsresponder/mDNSPosix/mDNSPosix.c
        ${ANDROID_ROOT}/external/mdnsresponder/mDNSPosix/mDNSUNP.c
        ${ANDROID_ROOT}/external/mdnsresponder/mDNSPosix/PosixDaemon.c
        ${ANDROID_ROOT}/external/mdnsresponder/mDNSShared/dnssd_ipc.c
        ${ANDROID_ROOT}/external/mdnsresponder/mDNSShared/GenLinkedList.c
        ${ANDROID_ROOT}/external/mdnsresponder/mDNSShared/mDNSDebug.c
        ${ANDROID_ROOT}/external/mdnsresponder/mDNSShared/PlatformCommon.c
        ${ANDROID_ROOT}/external/mdnsresponder/mDNSShared/uds_daemon.c
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
    add_executable(mdnsd ${SOURCE_FILES})

    target_include_directories(mdnsd
            PUBLIC
            ${ANDROID_ROOT}/external/mdnsresponder/mDNSCore
            ${ANDROID_ROOT}/external/mdnsresponder/mDNSPosix
            ${ANDROID_ROOT}/external/mdnsresponder/mDNSShared
            ${ANDROID_ROOT}/external/mdnsresponder
    )

    # Add c/c++ flags
    target_compile_options(mdnsd PRIVATE
            -DMDNS_DEBUGMSGS=0
            -Wall
            -Wextra
            -Wno-unused
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Werror=implicit-function-declaration
            -Werror
    )

    target_link_libraries(mdnsd
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            liblog
            libm
    )
endif ()
#################### mdnsd ####################

#################### libmdnssd ####################
set(FILES
        ${ANDROID_ROOT}/external/mdnsresponder/mDNSShared/dnssd_clientlib.c
        ${ANDROID_ROOT}/external/mdnsresponder/mDNSShared/dnssd_clientstub.c
        ${ANDROID_ROOT}/external/mdnsresponder/mDNSShared/dnssd_ipc.c
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
    add_library(libmdnssd SHARED ${SOURCE_FILES})

    target_include_directories(libmdnssd
            PUBLIC
            ${ANDROID_ROOT}/external/mdnsresponder/mDNSShared
            ${ANDROID_ROOT}/external/mdnsresponder
    )

    # Add c/c++ flags
    target_compile_options(libmdnssd PRIVATE
            -DMDNS_DEBUGMSGS=0
            -Wall
            -Wextra
            -Wno-unused
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Werror=implicit-function-declaration
            -Werror
    )

    target_link_libraries(libmdnssd
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            liblog
            libm
            libwinpthread
    )
endif ()
#################### libmdnssd ####################

#################### dnssd ####################
set(FILES
        ${ANDROID_ROOT}/external/mdnsresponder/Clients/dns-sd.c
        ${ANDROID_ROOT}/external/mdnsresponder/Clients/ClientCommon.c
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
    add_executable(dnssd ${SOURCE_FILES})

    target_include_directories(dnssd
            PUBLIC
            ${ANDROID_ROOT}/external/mdnsresponder
            ${ANDROID_ROOT}/external/mdnsresponder/mDNSShared
    )

    # Add c/c++ flags
    target_compile_options(dnssd PRIVATE
            -DMDNS_DEBUGMSGS=0
            -Wall
            -Wextra
            -Wno-unused
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Werror=implicit-function-declaration
            -Werror
    )

    target_link_libraries(dnssd
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            liblog
            libmdnssd
    )
endif ()
#################### dnssd ####################
