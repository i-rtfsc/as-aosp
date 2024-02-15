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

project(libril-modem-lib)

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

#################### libril-modem-lib ####################
set(FILES
        ${ANDROID_ROOT}/device/google/cuttlefish/guest/hals/ril/reference-libril/RefRadioSim.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/guest/hals/ril/reference-libril/RefRadioModem.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/guest/hals/ril/reference-libril/RefRadioIms.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/guest/hals/ril/reference-libril/RefImsMedia.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/guest/hals/ril/reference-libril/RefImsMediaSession.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/guest/hals/ril/reference-libril/RefRadioNetwork.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/guest/hals/ril/reference-libril/ril.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/guest/hals/ril/reference-libril/RilSapSocket.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/guest/hals/ril/reference-libril/ril_config.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/guest/hals/ril/reference-libril/ril_event.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/guest/hals/ril/reference-libril/ril_service.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/guest/hals/ril/reference-libril/sap_service.cpp
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
    add_library(libril-modem-lib SHARED ${SOURCE_FILES})

    target_include_directories(libril-modem-lib
            PUBLIC
            ${ANDROID_ROOT}/device/google/cuttlefish/guest/hals/ril/reference-libril
    )

    # Add c/c++ flags
    target_compile_options(libril-modem-lib PRIVATE
            -Wextra
            -Wno-unused-parameter
    )

    target_link_libraries(libril-modem-lib
            android.hardware.radio-library.compat
            android.hardware.radio.config-V2-ndk
            android.hardware.radio.config-1.0
            android.hardware.radio.config-1.1
            android.hardware.radio.config-1.2
            android.hardware.radio.config-1.3
            android.hardware.radio.data-V2-ndk
            android.hardware.radio.deprecated-1.0
            android.hardware.radio.ims-V1-ndk
            android.hardware.radio.ims.media-V1-ndk
            android.hardware.radio.messaging-V2-ndk
            android.hardware.radio.modem-V2-ndk
            android.hardware.radio.network-V2-ndk
            android.hardware.radio.sap-V1-ndk
            android.hardware.radio.sim-V2-ndk
            android.hardware.radio.voice-V2-ndk
            android.hardware.radio-1.0
            android.hardware.radio-1.1
            android.hardware.radio-1.2
            android.hardware.radio-1.3
            android.hardware.radio-1.4
            android.hardware.radio-1.5
            android.hardware.radio-1.6
            libbase
            libbinder_ndk
            libc++
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libcutils
            libdl
            libhardware_legacy
            libhidlbase
            liblog
            libm
            libprotobuf-c-nano-enable_malloc
            librilutils
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### libril-modem-lib ####################
