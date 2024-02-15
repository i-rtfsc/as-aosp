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

project(android.hardware.radio-library.compat)

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

#################### android.hardware.radio-library.compat ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/radio/aidl/compat/libradiocompat/CallbackManager.cpp
        ${ANDROID_ROOT}/hardware/interfaces/radio/aidl/compat/libradiocompat/DriverContext.cpp
        ${ANDROID_ROOT}/hardware/interfaces/radio/aidl/compat/libradiocompat/RadioCompatBase.cpp
        ${ANDROID_ROOT}/hardware/interfaces/radio/aidl/compat/libradiocompat/RadioIndication.cpp
        ${ANDROID_ROOT}/hardware/interfaces/radio/aidl/compat/libradiocompat/RadioResponse.cpp
        ${ANDROID_ROOT}/hardware/interfaces/radio/aidl/compat/libradiocompat/commonStructs.cpp
        ${ANDROID_ROOT}/hardware/interfaces/radio/aidl/compat/libradiocompat/config/RadioConfig.cpp
        ${ANDROID_ROOT}/hardware/interfaces/radio/aidl/compat/libradiocompat/config/RadioConfigIndication.cpp
        ${ANDROID_ROOT}/hardware/interfaces/radio/aidl/compat/libradiocompat/config/RadioConfigResponse.cpp
        ${ANDROID_ROOT}/hardware/interfaces/radio/aidl/compat/libradiocompat/config/structs.cpp
        ${ANDROID_ROOT}/hardware/interfaces/radio/aidl/compat/libradiocompat/data/RadioIndication-data.cpp
        ${ANDROID_ROOT}/hardware/interfaces/radio/aidl/compat/libradiocompat/data/RadioResponse-data.cpp
        ${ANDROID_ROOT}/hardware/interfaces/radio/aidl/compat/libradiocompat/data/RadioData.cpp
        ${ANDROID_ROOT}/hardware/interfaces/radio/aidl/compat/libradiocompat/data/structs.cpp
        ${ANDROID_ROOT}/hardware/interfaces/radio/aidl/compat/libradiocompat/ims/RadioIndication-ims.cpp
        ${ANDROID_ROOT}/hardware/interfaces/radio/aidl/compat/libradiocompat/ims/RadioResponse-ims.cpp
        ${ANDROID_ROOT}/hardware/interfaces/radio/aidl/compat/libradiocompat/ims/RadioIms.cpp
        ${ANDROID_ROOT}/hardware/interfaces/radio/aidl/compat/libradiocompat/ims/media/RadioImsMediaSession.cpp
        ${ANDROID_ROOT}/hardware/interfaces/radio/aidl/compat/libradiocompat/ims/media/RadioImsMedia.cpp
        ${ANDROID_ROOT}/hardware/interfaces/radio/aidl/compat/libradiocompat/messaging/RadioIndication-messaging.cpp
        ${ANDROID_ROOT}/hardware/interfaces/radio/aidl/compat/libradiocompat/messaging/RadioMessaging.cpp
        ${ANDROID_ROOT}/hardware/interfaces/radio/aidl/compat/libradiocompat/messaging/RadioResponse-messaging.cpp
        ${ANDROID_ROOT}/hardware/interfaces/radio/aidl/compat/libradiocompat/messaging/structs.cpp
        ${ANDROID_ROOT}/hardware/interfaces/radio/aidl/compat/libradiocompat/modem/RadioIndication-modem.cpp
        ${ANDROID_ROOT}/hardware/interfaces/radio/aidl/compat/libradiocompat/modem/RadioResponse-modem.cpp
        ${ANDROID_ROOT}/hardware/interfaces/radio/aidl/compat/libradiocompat/modem/RadioModem.cpp
        ${ANDROID_ROOT}/hardware/interfaces/radio/aidl/compat/libradiocompat/modem/structs.cpp
        ${ANDROID_ROOT}/hardware/interfaces/radio/aidl/compat/libradiocompat/network/RadioIndication-network.cpp
        ${ANDROID_ROOT}/hardware/interfaces/radio/aidl/compat/libradiocompat/network/RadioNetwork.cpp
        ${ANDROID_ROOT}/hardware/interfaces/radio/aidl/compat/libradiocompat/network/RadioResponse-network.cpp
        ${ANDROID_ROOT}/hardware/interfaces/radio/aidl/compat/libradiocompat/network/structs.cpp
        ${ANDROID_ROOT}/hardware/interfaces/radio/aidl/compat/libradiocompat/network/utils.cpp
        ${ANDROID_ROOT}/hardware/interfaces/radio/aidl/compat/libradiocompat/sap/Sap.cpp
        ${ANDROID_ROOT}/hardware/interfaces/radio/aidl/compat/libradiocompat/sap/SapCallback.cpp
        ${ANDROID_ROOT}/hardware/interfaces/radio/aidl/compat/libradiocompat/sap/structs.cpp
        ${ANDROID_ROOT}/hardware/interfaces/radio/aidl/compat/libradiocompat/sim/RadioIndication-sim.cpp
        ${ANDROID_ROOT}/hardware/interfaces/radio/aidl/compat/libradiocompat/sim/RadioResponse-sim.cpp
        ${ANDROID_ROOT}/hardware/interfaces/radio/aidl/compat/libradiocompat/sim/RadioSim.cpp
        ${ANDROID_ROOT}/hardware/interfaces/radio/aidl/compat/libradiocompat/sim/structs.cpp
        ${ANDROID_ROOT}/hardware/interfaces/radio/aidl/compat/libradiocompat/voice/RadioIndication-voice.cpp
        ${ANDROID_ROOT}/hardware/interfaces/radio/aidl/compat/libradiocompat/voice/RadioResponse-voice.cpp
        ${ANDROID_ROOT}/hardware/interfaces/radio/aidl/compat/libradiocompat/voice/RadioVoice.cpp
        ${ANDROID_ROOT}/hardware/interfaces/radio/aidl/compat/libradiocompat/voice/structs.cpp
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
    add_library(android.hardware.radio-library.compat SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio-library.compat
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/radio/aidl/compat/libradiocompat/include
            ${ANDROID_ROOT}/hardware/interfaces/radio/aidl/compat/libradiocompat
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio-library.compat PRIVATE
            -Werror
            -Wall
            -Wextra
    )

    target_link_libraries(android.hardware.radio-library.compat
            android.hardware.radio.config-V2-ndk
            android.hardware.radio.config-1.0
            android.hardware.radio.config-1.1
            android.hardware.radio.config-1.2
            android.hardware.radio.config-1.3
            android.hardware.radio.data-V2-ndk
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
            libclang_rt.ubsan_minimal
            libdl
            libhidlbase
            libm
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.radio-library.compat ####################
