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

project(libtelediagaidl)

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

#################### libtelediagaidl ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/meizu-telephony/packages/modules/MzNetworkDiagnostics/service/libtelediagaidl/android_${OUT_ARCH_CPU}_shared/gen/aidl/com/meizu/internal/telephony/diagnostics/ITelephonyBpfReader.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/meizu-telephony/packages/modules/MzNetworkDiagnostics/service/libtelediagaidl/android_${OUT_ARCH_CPU}_shared/gen/aidl/com/meizu/internal/telephony/diagnostics/ITelephonyBpfCallback.cpp
        ${ANDROID_ROOT}/vendor/meizu-telephony/packages/modules/MzNetworkDiagnostics/service/aidl/src/IcmpProbeResult.cpp
        ${ANDROID_ROOT}/vendor/meizu-telephony/packages/modules/MzNetworkDiagnostics/service/aidl/src/TcpProbeResult.cpp
        ${ANDROID_ROOT}/vendor/meizu-telephony/packages/modules/MzNetworkDiagnostics/service/aidl/src/TcpRetransmitSkbResponse.cpp
        ${ANDROID_ROOT}/vendor/meizu-telephony/packages/modules/MzNetworkDiagnostics/service/aidl/src/UdpProbeResult.cpp
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
    add_library(libtelediagaidl SHARED ${SOURCE_FILES})

    target_include_directories(libtelediagaidl
            PUBLIC
            ${ANDROID_ROOT}/vendor/meizu-telephony/packages/modules/MzNetworkDiagnostics/service/aidl/include
            ${ANDROID_ROOT}/vendor/meizu-telephony/packages/modules/MzNetworkDiagnostics/service/aidl
            ${ANDROID_ROOT}/vendor/meizu-telephony/packages/modules/MzNetworkDiagnostics/service
            ${OUT_INTERMEDIATES_ROOT}/vendor/meizu-telephony/packages/modules/MzNetworkDiagnostics/service/libtelediagaidl/android_${OUT_ARCH_CPU}_shared/gen/aidl
    )

    target_link_libraries(libtelediagaidl
            libbase
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            liblog
            libm
            libutils
    )
endif ()
#################### libtelediagaidl ####################

#################### libtelediag ####################
set(FILES
        ${ANDROID_ROOT}/vendor/meizu-telephony/packages/modules/MzNetworkDiagnostics/service/native/TeleDiagReaderService.cpp
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
    add_library(libtelediag SHARED ${SOURCE_FILES})

    target_include_directories(libtelediag
            PUBLIC
            ${ANDROID_ROOT}/vendor/meizu-telephony/packages/modules/MzNetworkDiagnostics/service/native/include
            ${ANDROID_ROOT}/vendor/meizu-telephony/packages/modules/MzNetworkDiagnostics/service
            ${ANDROID_ROOT}/vendor/meizu-telephony/packages/modules/MzNetworkDiagnostics/service/aidl/include
            ${OUT_INTERMEDIATES_ROOT}/vendor/meizu-telephony/packages/modules/MzNetworkDiagnostics/service/libtelediagaidl/android_${OUT_ARCH_CPU}_shared/gen/aidl
    )

    target_link_libraries(libtelediag
            bpf_commdiag_headers
            bpf_headers
            bpf_prog_headers
            libbase
            libbinder
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
            libtelediagaidl
            libutils
    )
endif ()
#################### libtelediag ####################
