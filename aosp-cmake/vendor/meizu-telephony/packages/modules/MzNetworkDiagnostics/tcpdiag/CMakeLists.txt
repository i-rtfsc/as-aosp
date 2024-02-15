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

project(TcpDiagCli)

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

#################### libtcpdiag ####################
set(FILES
        ${ANDROID_ROOT}/vendor/meizu-telephony/packages/modules/MzNetworkDiagnostics/tcpdiag/TcpDiag.cpp
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
    add_library(libtcpdiag SHARED ${SOURCE_FILES})

    target_include_directories(libtcpdiag
            PUBLIC
            ${ANDROID_ROOT}/vendor/meizu-telephony/packages/modules/MzNetworkDiagnostics/tcpdiag/include
            ${ANDROID_ROOT}/vendor/meizu-telephony/packages/modules/MzNetworkDiagnostics/tcpdiag
    )

    target_link_libraries(libtcpdiag
            bpf_commdiag_headers
            bpf_headers
            bpf_prog_headers
            libbase
            libbpf_android
            libbpf_bcc
            libbpf_minimal
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            liblog
            libm
            libnetdutils
            libutils
    )
endif ()
#################### libtcpdiag ####################

#################### TcpDiagDbgCli ####################
set(FILES
        ${ANDROID_ROOT}/vendor/meizu-telephony/packages/modules/MzNetworkDiagnostics/tcpdiag/TcpDiagDbgCli.cpp
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
    add_executable(TcpDiagDbgCli ${SOURCE_FILES})

    target_include_directories(TcpDiagDbgCli
            PUBLIC
            ${ANDROID_ROOT}/vendor/meizu-telephony/packages/modules/MzNetworkDiagnostics/tcpdiag
    )

    # Add c/c++ flags
    target_compile_options(TcpDiagDbgCli PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(TcpDiagDbgCli
            bpf_commdiag_headers
            bpf_headers
            bpf_prog_headers
            libbase
            libbpf_android
            libbpf_bcc
            libbpf_minimal
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            liblog
            libm
            libnetdutils
            libutils
    )
endif ()
#################### TcpDiagDbgCli ####################

#################### TcpDiagCli ####################
set(FILES
        ${ANDROID_ROOT}/vendor/meizu-telephony/packages/modules/MzNetworkDiagnostics/tcpdiag/TcpDiagCli.cpp
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
    add_executable(TcpDiagCli ${SOURCE_FILES})

    target_include_directories(TcpDiagCli
            PUBLIC
            ${ANDROID_ROOT}/vendor/meizu-telephony/packages/modules/MzNetworkDiagnostics/tcpdiag
    )

    # Add c/c++ flags
    target_compile_options(TcpDiagCli PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(TcpDiagCli
            bpf_commdiag_headers
            bpf_headers
            bpf_prog_headers
            libbase
            libbpf_android
            libbpf_bcc
            libbpf_minimal
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            liblog
            libm
            libnetdutils
            libutils
    )
endif ()
#################### TcpDiagCli ####################
