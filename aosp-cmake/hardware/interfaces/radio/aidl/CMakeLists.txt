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

project(android.hardware.radio-cpp-analyzer)

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

#################### android.hardware.radio.sim-V1-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V1-ndk_platform-source/gen/android/hardware/radio/sim/AppStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V1-ndk_platform-source/gen/android/hardware/radio/sim/CardPowerState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V1-ndk_platform-source/gen/android/hardware/radio/sim/CardStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V1-ndk_platform-source/gen/android/hardware/radio/sim/Carrier.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V1-ndk_platform-source/gen/android/hardware/radio/sim/CarrierRestrictions.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V1-ndk_platform-source/gen/android/hardware/radio/sim/CdmaSubscriptionSource.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V1-ndk_platform-source/gen/android/hardware/radio/sim/IRadioSim.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V1-ndk_platform-source/gen/android/hardware/radio/sim/IRadioSimIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V1-ndk_platform-source/gen/android/hardware/radio/sim/IRadioSimResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V1-ndk_platform-source/gen/android/hardware/radio/sim/IccIo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V1-ndk_platform-source/gen/android/hardware/radio/sim/IccIoResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V1-ndk_platform-source/gen/android/hardware/radio/sim/ImsiEncryptionInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V1-ndk_platform-source/gen/android/hardware/radio/sim/PbReceivedStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V1-ndk_platform-source/gen/android/hardware/radio/sim/PersoSubstate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V1-ndk_platform-source/gen/android/hardware/radio/sim/PhonebookCapacity.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V1-ndk_platform-source/gen/android/hardware/radio/sim/PhonebookRecordInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V1-ndk_platform-source/gen/android/hardware/radio/sim/PinState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V1-ndk_platform-source/gen/android/hardware/radio/sim/SelectUiccSub.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V1-ndk_platform-source/gen/android/hardware/radio/sim/SimApdu.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V1-ndk_platform-source/gen/android/hardware/radio/sim/SimLockMultiSimPolicy.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V1-ndk_platform-source/gen/android/hardware/radio/sim/SimRefreshResult.cpp
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
    add_library(android.hardware.radio.sim-V1-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.sim-V1-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V1-ndk_platform-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V1-ndk_platform-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V1-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.sim-V1-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.sim-V1-ndk_platform
            android.hardware.radio-V1-ndk_platform
            android.hardware.radio.config-V1-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.radio.sim-V1-ndk_platform ####################

#################### android.hardware.radio.sap-cpp-analyzer ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sap-cpp-analyzer-source/gen/android/hardware/radio/sap/ISap.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sap-cpp-analyzer-source/gen/android/hardware/radio/sap/ISapCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sap-cpp-analyzer-source/gen/android/hardware/radio/sap/SapApduType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sap-cpp-analyzer-source/gen/android/hardware/radio/sap/SapConnectRsp.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sap-cpp-analyzer-source/gen/android/hardware/radio/sap/SapDisconnectType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sap-cpp-analyzer-source/gen/android/hardware/radio/sap/SapResultCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sap-cpp-analyzer-source/gen/android/hardware/radio/sap/SapStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sap-cpp-analyzer-source/gen/android/hardware/radio/sap/SapTransferProtocol.cpp
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
    add_executable(android.hardware.radio.sap-cpp-analyzer ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.sap-cpp-analyzer
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/radio/aidl
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sap-V2-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sap-cpp-analyzer-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.sap-cpp-analyzer PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.sap-cpp-analyzer
            aidl-analyzer-main
            android.hardware.radio.sap-V2-cpp
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
            libutils
    )
endif ()
#################### android.hardware.radio.sap-cpp-analyzer ####################

#################### android.hardware.radio-V1-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V1-ndk-source/gen/android/hardware/radio/AccessNetwork.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V1-ndk-source/gen/android/hardware/radio/RadioAccessFamily.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V1-ndk-source/gen/android/hardware/radio/RadioConst.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V1-ndk-source/gen/android/hardware/radio/RadioError.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V1-ndk-source/gen/android/hardware/radio/RadioIndicationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V1-ndk-source/gen/android/hardware/radio/RadioResponseInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V1-ndk-source/gen/android/hardware/radio/RadioResponseInfoModem.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V1-ndk-source/gen/android/hardware/radio/RadioResponseType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V1-ndk-source/gen/android/hardware/radio/RadioTechnology.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V1-ndk-source/gen/android/hardware/radio/RadioTechnologyFamily.cpp
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
    add_library(android.hardware.radio-V1-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio-V1-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V1-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio-V1-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio-V1-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.radio-V1-ndk ####################

#################### android.hardware.radio.sap-V1-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sap-V1-ndk-source/gen/android/hardware/radio/sap/ISap.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sap-V1-ndk-source/gen/android/hardware/radio/sap/ISapCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sap-V1-ndk-source/gen/android/hardware/radio/sap/SapApduType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sap-V1-ndk-source/gen/android/hardware/radio/sap/SapConnectRsp.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sap-V1-ndk-source/gen/android/hardware/radio/sap/SapDisconnectType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sap-V1-ndk-source/gen/android/hardware/radio/sap/SapResultCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sap-V1-ndk-source/gen/android/hardware/radio/sap/SapStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sap-V1-ndk-source/gen/android/hardware/radio/sap/SapTransferProtocol.cpp
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
    add_library(android.hardware.radio.sap-V1-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.sap-V1-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sap-V1-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.sap-V1-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.sap-V1-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.radio.sap-V1-ndk ####################

#################### android.hardware.radio.modem-V1-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V1-ndk-source/gen/android/hardware/radio/modem/ActivityStatsInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V1-ndk-source/gen/android/hardware/radio/modem/ActivityStatsTechSpecificInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V1-ndk-source/gen/android/hardware/radio/modem/DeviceStateType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V1-ndk-source/gen/android/hardware/radio/modem/HardwareConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V1-ndk-source/gen/android/hardware/radio/modem/HardwareConfigModem.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V1-ndk-source/gen/android/hardware/radio/modem/HardwareConfigSim.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V1-ndk-source/gen/android/hardware/radio/modem/IRadioModem.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V1-ndk-source/gen/android/hardware/radio/modem/IRadioModemIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V1-ndk-source/gen/android/hardware/radio/modem/IRadioModemResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V1-ndk-source/gen/android/hardware/radio/modem/NvItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V1-ndk-source/gen/android/hardware/radio/modem/NvWriteItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V1-ndk-source/gen/android/hardware/radio/modem/RadioCapability.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V1-ndk-source/gen/android/hardware/radio/modem/RadioState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V1-ndk-source/gen/android/hardware/radio/modem/ResetNvType.cpp
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
    add_library(android.hardware.radio.modem-V1-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.modem-V1-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V1-ndk-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V1-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.modem-V1-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.modem-V1-ndk
            android.hardware.radio-V1-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.radio.modem-V1-ndk ####################

#################### android.hardware.radio-V3-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V3-cpp-source/gen/android/hardware/radio/AccessNetwork.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V3-cpp-source/gen/android/hardware/radio/RadioAccessFamily.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V3-cpp-source/gen/android/hardware/radio/RadioConst.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V3-cpp-source/gen/android/hardware/radio/RadioError.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V3-cpp-source/gen/android/hardware/radio/RadioIndicationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V3-cpp-source/gen/android/hardware/radio/RadioResponseInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V3-cpp-source/gen/android/hardware/radio/RadioResponseInfoModem.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V3-cpp-source/gen/android/hardware/radio/RadioResponseType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V3-cpp-source/gen/android/hardware/radio/RadioTechnology.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V3-cpp-source/gen/android/hardware/radio/RadioTechnologyFamily.cpp
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
    add_library(android.hardware.radio-V3-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio-V3-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V3-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio-V3-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio-V3-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libm
            libutils
    )
endif ()
#################### android.hardware.radio-V3-cpp ####################

#################### android.hardware.radio.modem-V3-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V3-cpp-source/gen/android/hardware/radio/modem/ActivityStatsInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V3-cpp-source/gen/android/hardware/radio/modem/ActivityStatsTechSpecificInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V3-cpp-source/gen/android/hardware/radio/modem/DeviceStateType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V3-cpp-source/gen/android/hardware/radio/modem/HardwareConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V3-cpp-source/gen/android/hardware/radio/modem/HardwareConfigModem.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V3-cpp-source/gen/android/hardware/radio/modem/HardwareConfigSim.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V3-cpp-source/gen/android/hardware/radio/modem/IRadioModem.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V3-cpp-source/gen/android/hardware/radio/modem/IRadioModemIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V3-cpp-source/gen/android/hardware/radio/modem/IRadioModemResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V3-cpp-source/gen/android/hardware/radio/modem/ImeiInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V3-cpp-source/gen/android/hardware/radio/modem/NvItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V3-cpp-source/gen/android/hardware/radio/modem/NvWriteItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V3-cpp-source/gen/android/hardware/radio/modem/RadioCapability.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V3-cpp-source/gen/android/hardware/radio/modem/RadioState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V3-cpp-source/gen/android/hardware/radio/modem/ResetNvType.cpp
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
    add_library(android.hardware.radio.modem-V3-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.modem-V3-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V3-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.modem-V3-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.modem-V3-cpp
            android.hardware.radio-V2-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libm
            libutils
    )
endif ()
#################### android.hardware.radio.modem-V3-cpp ####################

#################### android.hardware.radio.voice-V3-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-ndk-source/gen/android/hardware/radio/voice/AudioQuality.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-ndk-source/gen/android/hardware/radio/voice/Call.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-ndk-source/gen/android/hardware/radio/voice/CallForwardInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-ndk-source/gen/android/hardware/radio/voice/CdmaCallWaiting.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-ndk-source/gen/android/hardware/radio/voice/CdmaDisplayInfoRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-ndk-source/gen/android/hardware/radio/voice/CdmaInformationRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-ndk-source/gen/android/hardware/radio/voice/CdmaLineControlInfoRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-ndk-source/gen/android/hardware/radio/voice/CdmaNumberInfoRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-ndk-source/gen/android/hardware/radio/voice/CdmaOtaProvisionStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-ndk-source/gen/android/hardware/radio/voice/CdmaRedirectingNumberInfoRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-ndk-source/gen/android/hardware/radio/voice/CdmaSignalInfoRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-ndk-source/gen/android/hardware/radio/voice/CdmaT53AudioControlInfoRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-ndk-source/gen/android/hardware/radio/voice/CdmaT53ClirInfoRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-ndk-source/gen/android/hardware/radio/voice/CfData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-ndk-source/gen/android/hardware/radio/voice/ClipStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-ndk-source/gen/android/hardware/radio/voice/Dial.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-ndk-source/gen/android/hardware/radio/voice/EmergencyCallRouting.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-ndk-source/gen/android/hardware/radio/voice/EmergencyNumber.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-ndk-source/gen/android/hardware/radio/voice/EmergencyServiceCategory.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-ndk-source/gen/android/hardware/radio/voice/IRadioVoice.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-ndk-source/gen/android/hardware/radio/voice/IRadioVoiceIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-ndk-source/gen/android/hardware/radio/voice/IRadioVoiceResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-ndk-source/gen/android/hardware/radio/voice/LastCallFailCause.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-ndk-source/gen/android/hardware/radio/voice/LastCallFailCauseInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-ndk-source/gen/android/hardware/radio/voice/SrvccState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-ndk-source/gen/android/hardware/radio/voice/SsInfoData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-ndk-source/gen/android/hardware/radio/voice/StkCcUnsolSsResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-ndk-source/gen/android/hardware/radio/voice/TtyMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-ndk-source/gen/android/hardware/radio/voice/UssdModeType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-ndk-source/gen/android/hardware/radio/voice/UusInfo.cpp
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
    add_library(android.hardware.radio.voice-V3-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.voice-V3-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-ndk-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.voice-V3-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.voice-V3-ndk
            android.hardware.radio-V2-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.radio.voice-V3-ndk ####################

#################### android.hardware.radio.data-V2-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-cpp-source/gen/android/hardware/radio/data/ApnAuthType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-cpp-source/gen/android/hardware/radio/data/ApnTypes.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-cpp-source/gen/android/hardware/radio/data/DataCallFailCause.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-cpp-source/gen/android/hardware/radio/data/DataProfileInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-cpp-source/gen/android/hardware/radio/data/DataRequestReason.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-cpp-source/gen/android/hardware/radio/data/DataThrottlingAction.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-cpp-source/gen/android/hardware/radio/data/EpsQos.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-cpp-source/gen/android/hardware/radio/data/IRadioData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-cpp-source/gen/android/hardware/radio/data/IRadioDataIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-cpp-source/gen/android/hardware/radio/data/IRadioDataResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-cpp-source/gen/android/hardware/radio/data/KeepaliveRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-cpp-source/gen/android/hardware/radio/data/KeepaliveStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-cpp-source/gen/android/hardware/radio/data/LinkAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-cpp-source/gen/android/hardware/radio/data/NrQos.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-cpp-source/gen/android/hardware/radio/data/OsAppId.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-cpp-source/gen/android/hardware/radio/data/PcoDataInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-cpp-source/gen/android/hardware/radio/data/PdpProtocolType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-cpp-source/gen/android/hardware/radio/data/PortRange.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-cpp-source/gen/android/hardware/radio/data/Qos.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-cpp-source/gen/android/hardware/radio/data/QosBandwidth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-cpp-source/gen/android/hardware/radio/data/QosFilter.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-cpp-source/gen/android/hardware/radio/data/QosFilterIpsecSpi.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-cpp-source/gen/android/hardware/radio/data/QosFilterIpv6FlowLabel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-cpp-source/gen/android/hardware/radio/data/QosFilterTypeOfService.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-cpp-source/gen/android/hardware/radio/data/QosSession.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-cpp-source/gen/android/hardware/radio/data/RouteSelectionDescriptor.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-cpp-source/gen/android/hardware/radio/data/SetupDataCallResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-cpp-source/gen/android/hardware/radio/data/SliceInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-cpp-source/gen/android/hardware/radio/data/SlicingConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-cpp-source/gen/android/hardware/radio/data/TrafficDescriptor.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-cpp-source/gen/android/hardware/radio/data/UrspRule.cpp
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
    add_library(android.hardware.radio.data-V2-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.data-V2-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.data-V2-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.data-V2-cpp
            android.hardware.radio-V2-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libm
            libutils
    )
endif ()
#################### android.hardware.radio.data-V2-cpp ####################

#################### android.hardware.radio.config-V1-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V1-ndk_platform-source/gen/android/hardware/radio/config/IRadioConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V1-ndk_platform-source/gen/android/hardware/radio/config/IRadioConfigIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V1-ndk_platform-source/gen/android/hardware/radio/config/IRadioConfigResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V1-ndk_platform-source/gen/android/hardware/radio/config/PhoneCapability.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V1-ndk_platform-source/gen/android/hardware/radio/config/SimPortInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V1-ndk_platform-source/gen/android/hardware/radio/config/SimSlotStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V1-ndk_platform-source/gen/android/hardware/radio/config/SlotPortMapping.cpp
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
    add_library(android.hardware.radio.config-V1-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.config-V1-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V1-ndk_platform-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V1-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.config-V1-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.config-V1-ndk_platform
            android.hardware.radio-V1-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.radio.config-V1-ndk_platform ####################

#################### android.hardware.radio.voice-V2-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-ndk_platform-source/gen/android/hardware/radio/voice/AudioQuality.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-ndk_platform-source/gen/android/hardware/radio/voice/Call.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-ndk_platform-source/gen/android/hardware/radio/voice/CallForwardInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-ndk_platform-source/gen/android/hardware/radio/voice/CdmaCallWaiting.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-ndk_platform-source/gen/android/hardware/radio/voice/CdmaDisplayInfoRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-ndk_platform-source/gen/android/hardware/radio/voice/CdmaInformationRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-ndk_platform-source/gen/android/hardware/radio/voice/CdmaLineControlInfoRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-ndk_platform-source/gen/android/hardware/radio/voice/CdmaNumberInfoRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-ndk_platform-source/gen/android/hardware/radio/voice/CdmaOtaProvisionStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-ndk_platform-source/gen/android/hardware/radio/voice/CdmaRedirectingNumberInfoRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-ndk_platform-source/gen/android/hardware/radio/voice/CdmaSignalInfoRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-ndk_platform-source/gen/android/hardware/radio/voice/CdmaT53AudioControlInfoRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-ndk_platform-source/gen/android/hardware/radio/voice/CdmaT53ClirInfoRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-ndk_platform-source/gen/android/hardware/radio/voice/CfData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-ndk_platform-source/gen/android/hardware/radio/voice/ClipStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-ndk_platform-source/gen/android/hardware/radio/voice/Dial.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-ndk_platform-source/gen/android/hardware/radio/voice/EmergencyCallRouting.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-ndk_platform-source/gen/android/hardware/radio/voice/EmergencyNumber.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-ndk_platform-source/gen/android/hardware/radio/voice/EmergencyServiceCategory.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-ndk_platform-source/gen/android/hardware/radio/voice/IRadioVoice.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-ndk_platform-source/gen/android/hardware/radio/voice/IRadioVoiceIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-ndk_platform-source/gen/android/hardware/radio/voice/IRadioVoiceResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-ndk_platform-source/gen/android/hardware/radio/voice/LastCallFailCause.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-ndk_platform-source/gen/android/hardware/radio/voice/LastCallFailCauseInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-ndk_platform-source/gen/android/hardware/radio/voice/SrvccState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-ndk_platform-source/gen/android/hardware/radio/voice/SsInfoData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-ndk_platform-source/gen/android/hardware/radio/voice/StkCcUnsolSsResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-ndk_platform-source/gen/android/hardware/radio/voice/TtyMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-ndk_platform-source/gen/android/hardware/radio/voice/UssdModeType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-ndk_platform-source/gen/android/hardware/radio/voice/UusInfo.cpp
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
    add_library(android.hardware.radio.voice-V2-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.voice-V2-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-ndk_platform-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.voice-V2-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.voice-V2-ndk_platform
            android.hardware.radio-V2-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.radio.voice-V2-ndk_platform ####################

#################### android.hardware.radio.voice-V1-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-cpp-source/gen/android/hardware/radio/voice/AudioQuality.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-cpp-source/gen/android/hardware/radio/voice/Call.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-cpp-source/gen/android/hardware/radio/voice/CallForwardInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-cpp-source/gen/android/hardware/radio/voice/CdmaCallWaiting.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-cpp-source/gen/android/hardware/radio/voice/CdmaDisplayInfoRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-cpp-source/gen/android/hardware/radio/voice/CdmaInformationRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-cpp-source/gen/android/hardware/radio/voice/CdmaLineControlInfoRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-cpp-source/gen/android/hardware/radio/voice/CdmaNumberInfoRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-cpp-source/gen/android/hardware/radio/voice/CdmaOtaProvisionStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-cpp-source/gen/android/hardware/radio/voice/CdmaRedirectingNumberInfoRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-cpp-source/gen/android/hardware/radio/voice/CdmaSignalInfoRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-cpp-source/gen/android/hardware/radio/voice/CdmaT53AudioControlInfoRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-cpp-source/gen/android/hardware/radio/voice/CdmaT53ClirInfoRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-cpp-source/gen/android/hardware/radio/voice/CfData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-cpp-source/gen/android/hardware/radio/voice/ClipStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-cpp-source/gen/android/hardware/radio/voice/Dial.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-cpp-source/gen/android/hardware/radio/voice/EmergencyCallRouting.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-cpp-source/gen/android/hardware/radio/voice/EmergencyNumber.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-cpp-source/gen/android/hardware/radio/voice/EmergencyServiceCategory.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-cpp-source/gen/android/hardware/radio/voice/IRadioVoice.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-cpp-source/gen/android/hardware/radio/voice/IRadioVoiceIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-cpp-source/gen/android/hardware/radio/voice/IRadioVoiceResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-cpp-source/gen/android/hardware/radio/voice/LastCallFailCause.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-cpp-source/gen/android/hardware/radio/voice/LastCallFailCauseInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-cpp-source/gen/android/hardware/radio/voice/SrvccState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-cpp-source/gen/android/hardware/radio/voice/SsInfoData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-cpp-source/gen/android/hardware/radio/voice/StkCcUnsolSsResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-cpp-source/gen/android/hardware/radio/voice/TtyMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-cpp-source/gen/android/hardware/radio/voice/UssdModeType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-cpp-source/gen/android/hardware/radio/voice/UusInfo.cpp
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
    add_library(android.hardware.radio.voice-V1-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.voice-V1-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V1-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.voice-V1-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.voice-V1-cpp
            android.hardware.radio-V1-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libm
            libutils
    )
endif ()
#################### android.hardware.radio.voice-V1-cpp ####################

#################### android.hardware.radio-V2-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-ndk_platform-source/gen/android/hardware/radio/AccessNetwork.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-ndk_platform-source/gen/android/hardware/radio/RadioAccessFamily.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-ndk_platform-source/gen/android/hardware/radio/RadioConst.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-ndk_platform-source/gen/android/hardware/radio/RadioError.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-ndk_platform-source/gen/android/hardware/radio/RadioIndicationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-ndk_platform-source/gen/android/hardware/radio/RadioResponseInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-ndk_platform-source/gen/android/hardware/radio/RadioResponseInfoModem.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-ndk_platform-source/gen/android/hardware/radio/RadioResponseType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-ndk_platform-source/gen/android/hardware/radio/RadioTechnology.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-ndk_platform-source/gen/android/hardware/radio/RadioTechnologyFamily.cpp
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
    add_library(android.hardware.radio-V2-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio-V2-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio-V2-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio-V2-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.radio-V2-ndk_platform ####################

#################### android.hardware.radio.ims-V1-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims-V1-ndk-source/gen/android/hardware/radio/ims/ConnectionFailureInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims-V1-ndk-source/gen/android/hardware/radio/ims/EpsFallbackReason.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims-V1-ndk-source/gen/android/hardware/radio/ims/IRadioIms.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims-V1-ndk-source/gen/android/hardware/radio/ims/IRadioImsIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims-V1-ndk-source/gen/android/hardware/radio/ims/IRadioImsResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims-V1-ndk-source/gen/android/hardware/radio/ims/ImsCall.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims-V1-ndk-source/gen/android/hardware/radio/ims/ImsDeregistrationReason.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims-V1-ndk-source/gen/android/hardware/radio/ims/ImsRegistration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims-V1-ndk-source/gen/android/hardware/radio/ims/ImsRegistrationState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims-V1-ndk-source/gen/android/hardware/radio/ims/ImsStreamDirection.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims-V1-ndk-source/gen/android/hardware/radio/ims/ImsStreamType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims-V1-ndk-source/gen/android/hardware/radio/ims/ImsTrafficType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims-V1-ndk-source/gen/android/hardware/radio/ims/SrvccCall.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims-V1-ndk-source/gen/android/hardware/radio/ims/SuggestedAction.cpp
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
    add_library(android.hardware.radio.ims-V1-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.ims-V1-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-ndk-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims-V1-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.ims-V1-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.ims-V1-ndk
            android.hardware.radio-V2-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.radio.ims-V1-ndk ####################

#################### android.hardware.radio.sap-V2-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sap-V2-cpp-source/gen/android/hardware/radio/sap/ISap.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sap-V2-cpp-source/gen/android/hardware/radio/sap/ISapCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sap-V2-cpp-source/gen/android/hardware/radio/sap/SapApduType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sap-V2-cpp-source/gen/android/hardware/radio/sap/SapConnectRsp.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sap-V2-cpp-source/gen/android/hardware/radio/sap/SapDisconnectType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sap-V2-cpp-source/gen/android/hardware/radio/sap/SapResultCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sap-V2-cpp-source/gen/android/hardware/radio/sap/SapStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sap-V2-cpp-source/gen/android/hardware/radio/sap/SapTransferProtocol.cpp
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
    add_library(android.hardware.radio.sap-V2-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.sap-V2-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sap-V2-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.sap-V2-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.sap-V2-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libm
            libutils
    )
endif ()
#################### android.hardware.radio.sap-V2-cpp ####################

#################### android.hardware.radio-V2-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-cpp-source/gen/android/hardware/radio/AccessNetwork.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-cpp-source/gen/android/hardware/radio/RadioAccessFamily.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-cpp-source/gen/android/hardware/radio/RadioConst.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-cpp-source/gen/android/hardware/radio/RadioError.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-cpp-source/gen/android/hardware/radio/RadioIndicationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-cpp-source/gen/android/hardware/radio/RadioResponseInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-cpp-source/gen/android/hardware/radio/RadioResponseInfoModem.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-cpp-source/gen/android/hardware/radio/RadioResponseType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-cpp-source/gen/android/hardware/radio/RadioTechnology.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-cpp-source/gen/android/hardware/radio/RadioTechnologyFamily.cpp
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
    add_library(android.hardware.radio-V2-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio-V2-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio-V2-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio-V2-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libm
            libutils
    )
endif ()
#################### android.hardware.radio-V2-cpp ####################

#################### android.hardware.radio.data-V3-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-ndk_platform-source/gen/android/hardware/radio/data/ApnAuthType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-ndk_platform-source/gen/android/hardware/radio/data/ApnTypes.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-ndk_platform-source/gen/android/hardware/radio/data/DataCallFailCause.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-ndk_platform-source/gen/android/hardware/radio/data/DataProfileInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-ndk_platform-source/gen/android/hardware/radio/data/DataRequestReason.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-ndk_platform-source/gen/android/hardware/radio/data/DataThrottlingAction.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-ndk_platform-source/gen/android/hardware/radio/data/EpsQos.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-ndk_platform-source/gen/android/hardware/radio/data/IRadioData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-ndk_platform-source/gen/android/hardware/radio/data/IRadioDataIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-ndk_platform-source/gen/android/hardware/radio/data/IRadioDataResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-ndk_platform-source/gen/android/hardware/radio/data/KeepaliveRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-ndk_platform-source/gen/android/hardware/radio/data/KeepaliveStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-ndk_platform-source/gen/android/hardware/radio/data/LinkAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-ndk_platform-source/gen/android/hardware/radio/data/NrQos.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-ndk_platform-source/gen/android/hardware/radio/data/OsAppId.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-ndk_platform-source/gen/android/hardware/radio/data/PcoDataInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-ndk_platform-source/gen/android/hardware/radio/data/PdpProtocolType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-ndk_platform-source/gen/android/hardware/radio/data/PortRange.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-ndk_platform-source/gen/android/hardware/radio/data/Qos.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-ndk_platform-source/gen/android/hardware/radio/data/QosBandwidth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-ndk_platform-source/gen/android/hardware/radio/data/QosFilter.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-ndk_platform-source/gen/android/hardware/radio/data/QosFilterIpsecSpi.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-ndk_platform-source/gen/android/hardware/radio/data/QosFilterIpv6FlowLabel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-ndk_platform-source/gen/android/hardware/radio/data/QosFilterTypeOfService.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-ndk_platform-source/gen/android/hardware/radio/data/QosSession.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-ndk_platform-source/gen/android/hardware/radio/data/RouteSelectionDescriptor.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-ndk_platform-source/gen/android/hardware/radio/data/SetupDataCallResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-ndk_platform-source/gen/android/hardware/radio/data/SliceInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-ndk_platform-source/gen/android/hardware/radio/data/SlicingConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-ndk_platform-source/gen/android/hardware/radio/data/TrafficDescriptor.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-ndk_platform-source/gen/android/hardware/radio/data/UrspRule.cpp
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
    add_library(android.hardware.radio.data-V3-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.data-V3-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-ndk_platform-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.data-V3-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.data-V3-ndk_platform
            android.hardware.radio-V2-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.radio.data-V3-ndk_platform ####################

#################### android.hardware.radio.modem-V2-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V2-cpp-source/gen/android/hardware/radio/modem/ActivityStatsInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V2-cpp-source/gen/android/hardware/radio/modem/ActivityStatsTechSpecificInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V2-cpp-source/gen/android/hardware/radio/modem/DeviceStateType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V2-cpp-source/gen/android/hardware/radio/modem/HardwareConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V2-cpp-source/gen/android/hardware/radio/modem/HardwareConfigModem.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V2-cpp-source/gen/android/hardware/radio/modem/HardwareConfigSim.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V2-cpp-source/gen/android/hardware/radio/modem/IRadioModem.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V2-cpp-source/gen/android/hardware/radio/modem/IRadioModemIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V2-cpp-source/gen/android/hardware/radio/modem/IRadioModemResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V2-cpp-source/gen/android/hardware/radio/modem/ImeiInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V2-cpp-source/gen/android/hardware/radio/modem/NvItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V2-cpp-source/gen/android/hardware/radio/modem/NvWriteItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V2-cpp-source/gen/android/hardware/radio/modem/RadioCapability.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V2-cpp-source/gen/android/hardware/radio/modem/RadioState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V2-cpp-source/gen/android/hardware/radio/modem/ResetNvType.cpp
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
    add_library(android.hardware.radio.modem-V2-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.modem-V2-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V2-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.modem-V2-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.modem-V2-cpp
            android.hardware.radio-V2-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libm
            libutils
    )
endif ()
#################### android.hardware.radio.modem-V2-cpp ####################

#################### android.hardware.radio.data-V1-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-ndk-source/gen/android/hardware/radio/data/ApnAuthType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-ndk-source/gen/android/hardware/radio/data/ApnTypes.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-ndk-source/gen/android/hardware/radio/data/DataCallFailCause.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-ndk-source/gen/android/hardware/radio/data/DataProfileInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-ndk-source/gen/android/hardware/radio/data/DataRequestReason.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-ndk-source/gen/android/hardware/radio/data/DataThrottlingAction.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-ndk-source/gen/android/hardware/radio/data/EpsQos.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-ndk-source/gen/android/hardware/radio/data/IRadioData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-ndk-source/gen/android/hardware/radio/data/IRadioDataIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-ndk-source/gen/android/hardware/radio/data/IRadioDataResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-ndk-source/gen/android/hardware/radio/data/KeepaliveRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-ndk-source/gen/android/hardware/radio/data/KeepaliveStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-ndk-source/gen/android/hardware/radio/data/LinkAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-ndk-source/gen/android/hardware/radio/data/NrQos.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-ndk-source/gen/android/hardware/radio/data/OsAppId.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-ndk-source/gen/android/hardware/radio/data/PcoDataInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-ndk-source/gen/android/hardware/radio/data/PdpProtocolType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-ndk-source/gen/android/hardware/radio/data/PortRange.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-ndk-source/gen/android/hardware/radio/data/Qos.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-ndk-source/gen/android/hardware/radio/data/QosBandwidth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-ndk-source/gen/android/hardware/radio/data/QosFilter.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-ndk-source/gen/android/hardware/radio/data/QosFilterIpsecSpi.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-ndk-source/gen/android/hardware/radio/data/QosFilterIpv6FlowLabel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-ndk-source/gen/android/hardware/radio/data/QosFilterTypeOfService.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-ndk-source/gen/android/hardware/radio/data/QosSession.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-ndk-source/gen/android/hardware/radio/data/RouteSelectionDescriptor.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-ndk-source/gen/android/hardware/radio/data/SetupDataCallResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-ndk-source/gen/android/hardware/radio/data/SliceInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-ndk-source/gen/android/hardware/radio/data/SlicingConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-ndk-source/gen/android/hardware/radio/data/TrafficDescriptor.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-ndk-source/gen/android/hardware/radio/data/UrspRule.cpp
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
    add_library(android.hardware.radio.data-V1-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.data-V1-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V1-ndk-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.data-V1-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.data-V1-ndk
            android.hardware.radio-V1-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.radio.data-V1-ndk ####################

#################### android.hardware.radio.voice-V2-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-ndk-source/gen/android/hardware/radio/voice/AudioQuality.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-ndk-source/gen/android/hardware/radio/voice/Call.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-ndk-source/gen/android/hardware/radio/voice/CallForwardInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-ndk-source/gen/android/hardware/radio/voice/CdmaCallWaiting.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-ndk-source/gen/android/hardware/radio/voice/CdmaDisplayInfoRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-ndk-source/gen/android/hardware/radio/voice/CdmaInformationRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-ndk-source/gen/android/hardware/radio/voice/CdmaLineControlInfoRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-ndk-source/gen/android/hardware/radio/voice/CdmaNumberInfoRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-ndk-source/gen/android/hardware/radio/voice/CdmaOtaProvisionStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-ndk-source/gen/android/hardware/radio/voice/CdmaRedirectingNumberInfoRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-ndk-source/gen/android/hardware/radio/voice/CdmaSignalInfoRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-ndk-source/gen/android/hardware/radio/voice/CdmaT53AudioControlInfoRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-ndk-source/gen/android/hardware/radio/voice/CdmaT53ClirInfoRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-ndk-source/gen/android/hardware/radio/voice/CfData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-ndk-source/gen/android/hardware/radio/voice/ClipStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-ndk-source/gen/android/hardware/radio/voice/Dial.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-ndk-source/gen/android/hardware/radio/voice/EmergencyCallRouting.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-ndk-source/gen/android/hardware/radio/voice/EmergencyNumber.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-ndk-source/gen/android/hardware/radio/voice/EmergencyServiceCategory.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-ndk-source/gen/android/hardware/radio/voice/IRadioVoice.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-ndk-source/gen/android/hardware/radio/voice/IRadioVoiceIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-ndk-source/gen/android/hardware/radio/voice/IRadioVoiceResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-ndk-source/gen/android/hardware/radio/voice/LastCallFailCause.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-ndk-source/gen/android/hardware/radio/voice/LastCallFailCauseInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-ndk-source/gen/android/hardware/radio/voice/SrvccState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-ndk-source/gen/android/hardware/radio/voice/SsInfoData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-ndk-source/gen/android/hardware/radio/voice/StkCcUnsolSsResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-ndk-source/gen/android/hardware/radio/voice/TtyMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-ndk-source/gen/android/hardware/radio/voice/UssdModeType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-ndk-source/gen/android/hardware/radio/voice/UusInfo.cpp
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
    add_library(android.hardware.radio.voice-V2-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.voice-V2-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-ndk-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.voice-V2-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.voice-V2-ndk
            android.hardware.radio-V2-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.radio.voice-V2-ndk ####################

#################### android.hardware.radio.data-V3-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-cpp-source/gen/android/hardware/radio/data/ApnAuthType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-cpp-source/gen/android/hardware/radio/data/ApnTypes.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-cpp-source/gen/android/hardware/radio/data/DataCallFailCause.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-cpp-source/gen/android/hardware/radio/data/DataProfileInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-cpp-source/gen/android/hardware/radio/data/DataRequestReason.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-cpp-source/gen/android/hardware/radio/data/DataThrottlingAction.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-cpp-source/gen/android/hardware/radio/data/EpsQos.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-cpp-source/gen/android/hardware/radio/data/IRadioData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-cpp-source/gen/android/hardware/radio/data/IRadioDataIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-cpp-source/gen/android/hardware/radio/data/IRadioDataResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-cpp-source/gen/android/hardware/radio/data/KeepaliveRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-cpp-source/gen/android/hardware/radio/data/KeepaliveStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-cpp-source/gen/android/hardware/radio/data/LinkAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-cpp-source/gen/android/hardware/radio/data/NrQos.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-cpp-source/gen/android/hardware/radio/data/OsAppId.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-cpp-source/gen/android/hardware/radio/data/PcoDataInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-cpp-source/gen/android/hardware/radio/data/PdpProtocolType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-cpp-source/gen/android/hardware/radio/data/PortRange.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-cpp-source/gen/android/hardware/radio/data/Qos.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-cpp-source/gen/android/hardware/radio/data/QosBandwidth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-cpp-source/gen/android/hardware/radio/data/QosFilter.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-cpp-source/gen/android/hardware/radio/data/QosFilterIpsecSpi.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-cpp-source/gen/android/hardware/radio/data/QosFilterIpv6FlowLabel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-cpp-source/gen/android/hardware/radio/data/QosFilterTypeOfService.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-cpp-source/gen/android/hardware/radio/data/QosSession.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-cpp-source/gen/android/hardware/radio/data/RouteSelectionDescriptor.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-cpp-source/gen/android/hardware/radio/data/SetupDataCallResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-cpp-source/gen/android/hardware/radio/data/SliceInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-cpp-source/gen/android/hardware/radio/data/SlicingConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-cpp-source/gen/android/hardware/radio/data/TrafficDescriptor.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-cpp-source/gen/android/hardware/radio/data/UrspRule.cpp
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
    add_library(android.hardware.radio.data-V3-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.data-V3-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.data-V3-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.data-V3-cpp
            android.hardware.radio-V2-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libm
            libutils
    )
endif ()
#################### android.hardware.radio.data-V3-cpp ####################

#################### android.hardware.radio-V3-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V3-ndk-source/gen/android/hardware/radio/AccessNetwork.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V3-ndk-source/gen/android/hardware/radio/RadioAccessFamily.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V3-ndk-source/gen/android/hardware/radio/RadioConst.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V3-ndk-source/gen/android/hardware/radio/RadioError.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V3-ndk-source/gen/android/hardware/radio/RadioIndicationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V3-ndk-source/gen/android/hardware/radio/RadioResponseInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V3-ndk-source/gen/android/hardware/radio/RadioResponseInfoModem.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V3-ndk-source/gen/android/hardware/radio/RadioResponseType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V3-ndk-source/gen/android/hardware/radio/RadioTechnology.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V3-ndk-source/gen/android/hardware/radio/RadioTechnologyFamily.cpp
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
    add_library(android.hardware.radio-V3-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio-V3-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V3-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio-V3-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio-V3-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.radio-V3-ndk ####################

#################### android.hardware.radio.ims-V2-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims-V2-ndk-source/gen/android/hardware/radio/ims/ConnectionFailureInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims-V2-ndk-source/gen/android/hardware/radio/ims/EpsFallbackReason.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims-V2-ndk-source/gen/android/hardware/radio/ims/IRadioIms.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims-V2-ndk-source/gen/android/hardware/radio/ims/IRadioImsIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims-V2-ndk-source/gen/android/hardware/radio/ims/IRadioImsResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims-V2-ndk-source/gen/android/hardware/radio/ims/ImsCall.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims-V2-ndk-source/gen/android/hardware/radio/ims/ImsDeregistrationReason.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims-V2-ndk-source/gen/android/hardware/radio/ims/ImsRegistration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims-V2-ndk-source/gen/android/hardware/radio/ims/ImsRegistrationState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims-V2-ndk-source/gen/android/hardware/radio/ims/ImsStreamDirection.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims-V2-ndk-source/gen/android/hardware/radio/ims/ImsStreamType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims-V2-ndk-source/gen/android/hardware/radio/ims/ImsTrafficType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims-V2-ndk-source/gen/android/hardware/radio/ims/SrvccCall.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims-V2-ndk-source/gen/android/hardware/radio/ims/SuggestedAction.cpp
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
    add_library(android.hardware.radio.ims-V2-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.ims-V2-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-ndk-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims-V2-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.ims-V2-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.ims-V2-ndk
            android.hardware.radio-V2-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.radio.ims-V2-ndk ####################

#################### android.hardware.radio.modem-V3-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V3-ndk-source/gen/android/hardware/radio/modem/ActivityStatsInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V3-ndk-source/gen/android/hardware/radio/modem/ActivityStatsTechSpecificInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V3-ndk-source/gen/android/hardware/radio/modem/DeviceStateType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V3-ndk-source/gen/android/hardware/radio/modem/HardwareConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V3-ndk-source/gen/android/hardware/radio/modem/HardwareConfigModem.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V3-ndk-source/gen/android/hardware/radio/modem/HardwareConfigSim.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V3-ndk-source/gen/android/hardware/radio/modem/IRadioModem.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V3-ndk-source/gen/android/hardware/radio/modem/IRadioModemIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V3-ndk-source/gen/android/hardware/radio/modem/IRadioModemResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V3-ndk-source/gen/android/hardware/radio/modem/ImeiInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V3-ndk-source/gen/android/hardware/radio/modem/NvItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V3-ndk-source/gen/android/hardware/radio/modem/NvWriteItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V3-ndk-source/gen/android/hardware/radio/modem/RadioCapability.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V3-ndk-source/gen/android/hardware/radio/modem/RadioState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V3-ndk-source/gen/android/hardware/radio/modem/ResetNvType.cpp
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
    add_library(android.hardware.radio.modem-V3-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.modem-V3-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-ndk-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V3-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.modem-V3-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.modem-V3-ndk
            android.hardware.radio-V2-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.radio.modem-V3-ndk ####################

#################### android.hardware.radio.sap-V1-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sap-V1-cpp-source/gen/android/hardware/radio/sap/ISap.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sap-V1-cpp-source/gen/android/hardware/radio/sap/ISapCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sap-V1-cpp-source/gen/android/hardware/radio/sap/SapApduType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sap-V1-cpp-source/gen/android/hardware/radio/sap/SapConnectRsp.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sap-V1-cpp-source/gen/android/hardware/radio/sap/SapDisconnectType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sap-V1-cpp-source/gen/android/hardware/radio/sap/SapResultCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sap-V1-cpp-source/gen/android/hardware/radio/sap/SapStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sap-V1-cpp-source/gen/android/hardware/radio/sap/SapTransferProtocol.cpp
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
    add_library(android.hardware.radio.sap-V1-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.sap-V1-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sap-V1-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.sap-V1-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.sap-V1-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libm
            libutils
    )
endif ()
#################### android.hardware.radio.sap-V1-cpp ####################

#################### android.hardware.radio-V1-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V1-cpp-source/gen/android/hardware/radio/AccessNetwork.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V1-cpp-source/gen/android/hardware/radio/RadioAccessFamily.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V1-cpp-source/gen/android/hardware/radio/RadioConst.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V1-cpp-source/gen/android/hardware/radio/RadioError.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V1-cpp-source/gen/android/hardware/radio/RadioIndicationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V1-cpp-source/gen/android/hardware/radio/RadioResponseInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V1-cpp-source/gen/android/hardware/radio/RadioResponseInfoModem.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V1-cpp-source/gen/android/hardware/radio/RadioResponseType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V1-cpp-source/gen/android/hardware/radio/RadioTechnology.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V1-cpp-source/gen/android/hardware/radio/RadioTechnologyFamily.cpp
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
    add_library(android.hardware.radio-V1-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio-V1-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V1-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio-V1-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio-V1-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libm
            libutils
    )
endif ()
#################### android.hardware.radio-V1-cpp ####################

#################### android.hardware.radio.data-V2-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk_platform-source/gen/android/hardware/radio/data/ApnAuthType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk_platform-source/gen/android/hardware/radio/data/ApnTypes.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk_platform-source/gen/android/hardware/radio/data/DataCallFailCause.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk_platform-source/gen/android/hardware/radio/data/DataProfileInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk_platform-source/gen/android/hardware/radio/data/DataRequestReason.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk_platform-source/gen/android/hardware/radio/data/DataThrottlingAction.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk_platform-source/gen/android/hardware/radio/data/EpsQos.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk_platform-source/gen/android/hardware/radio/data/IRadioData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk_platform-source/gen/android/hardware/radio/data/IRadioDataIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk_platform-source/gen/android/hardware/radio/data/IRadioDataResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk_platform-source/gen/android/hardware/radio/data/KeepaliveRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk_platform-source/gen/android/hardware/radio/data/KeepaliveStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk_platform-source/gen/android/hardware/radio/data/LinkAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk_platform-source/gen/android/hardware/radio/data/NrQos.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk_platform-source/gen/android/hardware/radio/data/OsAppId.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk_platform-source/gen/android/hardware/radio/data/PcoDataInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk_platform-source/gen/android/hardware/radio/data/PdpProtocolType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk_platform-source/gen/android/hardware/radio/data/PortRange.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk_platform-source/gen/android/hardware/radio/data/Qos.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk_platform-source/gen/android/hardware/radio/data/QosBandwidth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk_platform-source/gen/android/hardware/radio/data/QosFilter.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk_platform-source/gen/android/hardware/radio/data/QosFilterIpsecSpi.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk_platform-source/gen/android/hardware/radio/data/QosFilterIpv6FlowLabel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk_platform-source/gen/android/hardware/radio/data/QosFilterTypeOfService.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk_platform-source/gen/android/hardware/radio/data/QosSession.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk_platform-source/gen/android/hardware/radio/data/RouteSelectionDescriptor.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk_platform-source/gen/android/hardware/radio/data/SetupDataCallResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk_platform-source/gen/android/hardware/radio/data/SliceInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk_platform-source/gen/android/hardware/radio/data/SlicingConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk_platform-source/gen/android/hardware/radio/data/TrafficDescriptor.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk_platform-source/gen/android/hardware/radio/data/UrspRule.cpp
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
    add_library(android.hardware.radio.data-V2-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.data-V2-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-ndk_platform-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.data-V2-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.data-V2-ndk_platform
            android.hardware.radio-V2-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.radio.data-V2-ndk_platform ####################

#################### android.hardware.radio.network-V1-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk_platform-source/gen/android/hardware/radio/network/AccessTechnologySpecificInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk_platform-source/gen/android/hardware/radio/network/BarringInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk_platform-source/gen/android/hardware/radio/network/BarringTypeSpecificInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk_platform-source/gen/android/hardware/radio/network/Cdma2000RegistrationInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk_platform-source/gen/android/hardware/radio/network/CdmaRoamingType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk_platform-source/gen/android/hardware/radio/network/CdmaSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk_platform-source/gen/android/hardware/radio/network/CellConnectionStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk_platform-source/gen/android/hardware/radio/network/CellIdentity.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk_platform-source/gen/android/hardware/radio/network/CellIdentityCdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk_platform-source/gen/android/hardware/radio/network/CellIdentityGsm.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk_platform-source/gen/android/hardware/radio/network/CellIdentityLte.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk_platform-source/gen/android/hardware/radio/network/CellIdentityNr.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk_platform-source/gen/android/hardware/radio/network/CellIdentityTdscdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk_platform-source/gen/android/hardware/radio/network/CellIdentityWcdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk_platform-source/gen/android/hardware/radio/network/CellInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk_platform-source/gen/android/hardware/radio/network/CellInfoCdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk_platform-source/gen/android/hardware/radio/network/CellInfoGsm.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk_platform-source/gen/android/hardware/radio/network/CellInfoLte.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk_platform-source/gen/android/hardware/radio/network/CellInfoNr.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk_platform-source/gen/android/hardware/radio/network/CellInfoRatSpecificInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk_platform-source/gen/android/hardware/radio/network/CellInfoTdscdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk_platform-source/gen/android/hardware/radio/network/CellInfoWcdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk_platform-source/gen/android/hardware/radio/network/ClosedSubscriberGroupInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk_platform-source/gen/android/hardware/radio/network/Domain.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk_platform-source/gen/android/hardware/radio/network/EutranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk_platform-source/gen/android/hardware/radio/network/EutranRegistrationInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk_platform-source/gen/android/hardware/radio/network/EvdoSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk_platform-source/gen/android/hardware/radio/network/GeranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk_platform-source/gen/android/hardware/radio/network/GsmSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk_platform-source/gen/android/hardware/radio/network/IRadioNetwork.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk_platform-source/gen/android/hardware/radio/network/IRadioNetworkIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk_platform-source/gen/android/hardware/radio/network/IRadioNetworkResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk_platform-source/gen/android/hardware/radio/network/IndicationFilter.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk_platform-source/gen/android/hardware/radio/network/LceDataInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk_platform-source/gen/android/hardware/radio/network/LinkCapacityEstimate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk_platform-source/gen/android/hardware/radio/network/LteSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk_platform-source/gen/android/hardware/radio/network/LteVopsInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk_platform-source/gen/android/hardware/radio/network/NetworkScanRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk_platform-source/gen/android/hardware/radio/network/NetworkScanResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk_platform-source/gen/android/hardware/radio/network/NgranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk_platform-source/gen/android/hardware/radio/network/NrDualConnectivityState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk_platform-source/gen/android/hardware/radio/network/NrIndicators.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk_platform-source/gen/android/hardware/radio/network/NrSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk_platform-source/gen/android/hardware/radio/network/NrVopsInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk_platform-source/gen/android/hardware/radio/network/OperatorInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk_platform-source/gen/android/hardware/radio/network/PhoneRestrictedState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk_platform-source/gen/android/hardware/radio/network/PhysicalChannelConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk_platform-source/gen/android/hardware/radio/network/PhysicalChannelConfigBand.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk_platform-source/gen/android/hardware/radio/network/RadioAccessSpecifier.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk_platform-source/gen/android/hardware/radio/network/RadioAccessSpecifierBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk_platform-source/gen/android/hardware/radio/network/RadioBandMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk_platform-source/gen/android/hardware/radio/network/RegState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk_platform-source/gen/android/hardware/radio/network/RegStateResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk_platform-source/gen/android/hardware/radio/network/RegistrationFailCause.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk_platform-source/gen/android/hardware/radio/network/SignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk_platform-source/gen/android/hardware/radio/network/SignalThresholdInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk_platform-source/gen/android/hardware/radio/network/SuppSvcNotification.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk_platform-source/gen/android/hardware/radio/network/TdscdmaSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk_platform-source/gen/android/hardware/radio/network/UsageSetting.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk_platform-source/gen/android/hardware/radio/network/UtranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk_platform-source/gen/android/hardware/radio/network/WcdmaSignalStrength.cpp
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
    add_library(android.hardware.radio.network-V1-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.network-V1-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V1-ndk_platform-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.network-V1-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.network-V1-ndk_platform
            android.hardware.radio-V1-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.radio.network-V1-ndk_platform ####################

#################### android.hardware.radio.modem-V1-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V1-cpp-source/gen/android/hardware/radio/modem/ActivityStatsInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V1-cpp-source/gen/android/hardware/radio/modem/ActivityStatsTechSpecificInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V1-cpp-source/gen/android/hardware/radio/modem/DeviceStateType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V1-cpp-source/gen/android/hardware/radio/modem/HardwareConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V1-cpp-source/gen/android/hardware/radio/modem/HardwareConfigModem.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V1-cpp-source/gen/android/hardware/radio/modem/HardwareConfigSim.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V1-cpp-source/gen/android/hardware/radio/modem/IRadioModem.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V1-cpp-source/gen/android/hardware/radio/modem/IRadioModemIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V1-cpp-source/gen/android/hardware/radio/modem/IRadioModemResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V1-cpp-source/gen/android/hardware/radio/modem/NvItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V1-cpp-source/gen/android/hardware/radio/modem/NvWriteItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V1-cpp-source/gen/android/hardware/radio/modem/RadioCapability.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V1-cpp-source/gen/android/hardware/radio/modem/RadioState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V1-cpp-source/gen/android/hardware/radio/modem/ResetNvType.cpp
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
    add_library(android.hardware.radio.modem-V1-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.modem-V1-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V1-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V1-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.modem-V1-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.modem-V1-cpp
            android.hardware.radio-V1-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libm
            libutils
    )
endif ()
#################### android.hardware.radio.modem-V1-cpp ####################

#################### android.hardware.radio.data-V2-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk-source/gen/android/hardware/radio/data/ApnAuthType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk-source/gen/android/hardware/radio/data/ApnTypes.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk-source/gen/android/hardware/radio/data/DataCallFailCause.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk-source/gen/android/hardware/radio/data/DataProfileInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk-source/gen/android/hardware/radio/data/DataRequestReason.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk-source/gen/android/hardware/radio/data/DataThrottlingAction.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk-source/gen/android/hardware/radio/data/EpsQos.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk-source/gen/android/hardware/radio/data/IRadioData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk-source/gen/android/hardware/radio/data/IRadioDataIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk-source/gen/android/hardware/radio/data/IRadioDataResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk-source/gen/android/hardware/radio/data/KeepaliveRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk-source/gen/android/hardware/radio/data/KeepaliveStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk-source/gen/android/hardware/radio/data/LinkAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk-source/gen/android/hardware/radio/data/NrQos.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk-source/gen/android/hardware/radio/data/OsAppId.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk-source/gen/android/hardware/radio/data/PcoDataInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk-source/gen/android/hardware/radio/data/PdpProtocolType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk-source/gen/android/hardware/radio/data/PortRange.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk-source/gen/android/hardware/radio/data/Qos.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk-source/gen/android/hardware/radio/data/QosBandwidth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk-source/gen/android/hardware/radio/data/QosFilter.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk-source/gen/android/hardware/radio/data/QosFilterIpsecSpi.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk-source/gen/android/hardware/radio/data/QosFilterIpv6FlowLabel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk-source/gen/android/hardware/radio/data/QosFilterTypeOfService.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk-source/gen/android/hardware/radio/data/QosSession.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk-source/gen/android/hardware/radio/data/RouteSelectionDescriptor.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk-source/gen/android/hardware/radio/data/SetupDataCallResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk-source/gen/android/hardware/radio/data/SliceInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk-source/gen/android/hardware/radio/data/SlicingConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk-source/gen/android/hardware/radio/data/TrafficDescriptor.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk-source/gen/android/hardware/radio/data/UrspRule.cpp
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
    add_library(android.hardware.radio.data-V2-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.data-V2-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-ndk-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.data-V2-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.data-V2-ndk
            android.hardware.radio-V2-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.radio.data-V2-ndk ####################

#################### android.hardware.radio.messaging-V1-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V1-ndk_platform-source/gen/android/hardware/radio/messaging/CdmaBroadcastSmsConfigInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V1-ndk_platform-source/gen/android/hardware/radio/messaging/CdmaSmsAck.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V1-ndk_platform-source/gen/android/hardware/radio/messaging/CdmaSmsAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V1-ndk_platform-source/gen/android/hardware/radio/messaging/CdmaSmsMessage.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V1-ndk_platform-source/gen/android/hardware/radio/messaging/CdmaSmsSubaddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V1-ndk_platform-source/gen/android/hardware/radio/messaging/CdmaSmsWriteArgs.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V1-ndk_platform-source/gen/android/hardware/radio/messaging/GsmBroadcastSmsConfigInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V1-ndk_platform-source/gen/android/hardware/radio/messaging/GsmSmsMessage.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V1-ndk_platform-source/gen/android/hardware/radio/messaging/IRadioMessaging.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V1-ndk_platform-source/gen/android/hardware/radio/messaging/IRadioMessagingIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V1-ndk_platform-source/gen/android/hardware/radio/messaging/IRadioMessagingResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V1-ndk_platform-source/gen/android/hardware/radio/messaging/ImsSmsMessage.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V1-ndk_platform-source/gen/android/hardware/radio/messaging/SendSmsResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V1-ndk_platform-source/gen/android/hardware/radio/messaging/SmsAcknowledgeFailCause.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V1-ndk_platform-source/gen/android/hardware/radio/messaging/SmsWriteArgs.cpp
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
    add_library(android.hardware.radio.messaging-V1-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.messaging-V1-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V1-ndk_platform-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V1-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.messaging-V1-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.messaging-V1-ndk_platform
            android.hardware.radio-V1-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.radio.messaging-V1-ndk_platform ####################

#################### android.hardware.radio.voice-V1-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-ndk-source/gen/android/hardware/radio/voice/AudioQuality.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-ndk-source/gen/android/hardware/radio/voice/Call.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-ndk-source/gen/android/hardware/radio/voice/CallForwardInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-ndk-source/gen/android/hardware/radio/voice/CdmaCallWaiting.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-ndk-source/gen/android/hardware/radio/voice/CdmaDisplayInfoRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-ndk-source/gen/android/hardware/radio/voice/CdmaInformationRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-ndk-source/gen/android/hardware/radio/voice/CdmaLineControlInfoRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-ndk-source/gen/android/hardware/radio/voice/CdmaNumberInfoRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-ndk-source/gen/android/hardware/radio/voice/CdmaOtaProvisionStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-ndk-source/gen/android/hardware/radio/voice/CdmaRedirectingNumberInfoRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-ndk-source/gen/android/hardware/radio/voice/CdmaSignalInfoRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-ndk-source/gen/android/hardware/radio/voice/CdmaT53AudioControlInfoRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-ndk-source/gen/android/hardware/radio/voice/CdmaT53ClirInfoRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-ndk-source/gen/android/hardware/radio/voice/CfData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-ndk-source/gen/android/hardware/radio/voice/ClipStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-ndk-source/gen/android/hardware/radio/voice/Dial.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-ndk-source/gen/android/hardware/radio/voice/EmergencyCallRouting.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-ndk-source/gen/android/hardware/radio/voice/EmergencyNumber.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-ndk-source/gen/android/hardware/radio/voice/EmergencyServiceCategory.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-ndk-source/gen/android/hardware/radio/voice/IRadioVoice.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-ndk-source/gen/android/hardware/radio/voice/IRadioVoiceIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-ndk-source/gen/android/hardware/radio/voice/IRadioVoiceResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-ndk-source/gen/android/hardware/radio/voice/LastCallFailCause.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-ndk-source/gen/android/hardware/radio/voice/LastCallFailCauseInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-ndk-source/gen/android/hardware/radio/voice/SrvccState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-ndk-source/gen/android/hardware/radio/voice/SsInfoData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-ndk-source/gen/android/hardware/radio/voice/StkCcUnsolSsResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-ndk-source/gen/android/hardware/radio/voice/TtyMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-ndk-source/gen/android/hardware/radio/voice/UssdModeType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-ndk-source/gen/android/hardware/radio/voice/UusInfo.cpp
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
    add_library(android.hardware.radio.voice-V1-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.voice-V1-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V1-ndk-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.voice-V1-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.voice-V1-ndk
            android.hardware.radio-V1-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.radio.voice-V1-ndk ####################

#################### android.hardware.radio.ims.media-V1-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V1-ndk_platform-source/gen/android/hardware/radio/ims/media/AmrMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V1-ndk_platform-source/gen/android/hardware/radio/ims/media/AmrParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V1-ndk_platform-source/gen/android/hardware/radio/ims/media/AnbrMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V1-ndk_platform-source/gen/android/hardware/radio/ims/media/CallQuality.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V1-ndk_platform-source/gen/android/hardware/radio/ims/media/CodecMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V1-ndk_platform-source/gen/android/hardware/radio/ims/media/CodecParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V1-ndk_platform-source/gen/android/hardware/radio/ims/media/CodecSpecificParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V1-ndk_platform-source/gen/android/hardware/radio/ims/media/CodecType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V1-ndk_platform-source/gen/android/hardware/radio/ims/media/DtmfParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V1-ndk_platform-source/gen/android/hardware/radio/ims/media/EvsBandwidth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V1-ndk_platform-source/gen/android/hardware/radio/ims/media/EvsMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V1-ndk_platform-source/gen/android/hardware/radio/ims/media/EvsParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V1-ndk_platform-source/gen/android/hardware/radio/ims/media/IImsMedia.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V1-ndk_platform-source/gen/android/hardware/radio/ims/media/IImsMediaListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V1-ndk_platform-source/gen/android/hardware/radio/ims/media/IImsMediaSession.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V1-ndk_platform-source/gen/android/hardware/radio/ims/media/IImsMediaSessionListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V1-ndk_platform-source/gen/android/hardware/radio/ims/media/LocalEndPoint.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V1-ndk_platform-source/gen/android/hardware/radio/ims/media/MediaDirection.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V1-ndk_platform-source/gen/android/hardware/radio/ims/media/MediaQualityStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V1-ndk_platform-source/gen/android/hardware/radio/ims/media/MediaQualityThreshold.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V1-ndk_platform-source/gen/android/hardware/radio/ims/media/RtcpConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V1-ndk_platform-source/gen/android/hardware/radio/ims/media/RtcpXrReportBlockType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V1-ndk_platform-source/gen/android/hardware/radio/ims/media/RtpAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V1-ndk_platform-source/gen/android/hardware/radio/ims/media/RtpConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V1-ndk_platform-source/gen/android/hardware/radio/ims/media/RtpError.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V1-ndk_platform-source/gen/android/hardware/radio/ims/media/RtpHeaderExtension.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V1-ndk_platform-source/gen/android/hardware/radio/ims/media/RtpSessionParams.cpp
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
    add_library(android.hardware.radio.ims.media-V1-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.ims.media-V1-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-ndk_platform-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk_platform-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V1-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.ims.media-V1-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.ims.media-V1-ndk_platform
            android.hardware.radio-V2-ndk_platform
            android.hardware.radio.data-V2-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.radio.ims.media-V1-ndk_platform ####################

#################### android.hardware.radio.sap-V2-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sap-V2-ndk_platform-source/gen/android/hardware/radio/sap/ISap.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sap-V2-ndk_platform-source/gen/android/hardware/radio/sap/ISapCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sap-V2-ndk_platform-source/gen/android/hardware/radio/sap/SapApduType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sap-V2-ndk_platform-source/gen/android/hardware/radio/sap/SapConnectRsp.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sap-V2-ndk_platform-source/gen/android/hardware/radio/sap/SapDisconnectType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sap-V2-ndk_platform-source/gen/android/hardware/radio/sap/SapResultCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sap-V2-ndk_platform-source/gen/android/hardware/radio/sap/SapStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sap-V2-ndk_platform-source/gen/android/hardware/radio/sap/SapTransferProtocol.cpp
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
    add_library(android.hardware.radio.sap-V2-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.sap-V2-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sap-V2-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.sap-V2-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.sap-V2-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.radio.sap-V2-ndk_platform ####################

#################### android.hardware.radio.voice-V3-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-cpp-source/gen/android/hardware/radio/voice/AudioQuality.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-cpp-source/gen/android/hardware/radio/voice/Call.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-cpp-source/gen/android/hardware/radio/voice/CallForwardInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-cpp-source/gen/android/hardware/radio/voice/CdmaCallWaiting.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-cpp-source/gen/android/hardware/radio/voice/CdmaDisplayInfoRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-cpp-source/gen/android/hardware/radio/voice/CdmaInformationRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-cpp-source/gen/android/hardware/radio/voice/CdmaLineControlInfoRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-cpp-source/gen/android/hardware/radio/voice/CdmaNumberInfoRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-cpp-source/gen/android/hardware/radio/voice/CdmaOtaProvisionStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-cpp-source/gen/android/hardware/radio/voice/CdmaRedirectingNumberInfoRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-cpp-source/gen/android/hardware/radio/voice/CdmaSignalInfoRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-cpp-source/gen/android/hardware/radio/voice/CdmaT53AudioControlInfoRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-cpp-source/gen/android/hardware/radio/voice/CdmaT53ClirInfoRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-cpp-source/gen/android/hardware/radio/voice/CfData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-cpp-source/gen/android/hardware/radio/voice/ClipStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-cpp-source/gen/android/hardware/radio/voice/Dial.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-cpp-source/gen/android/hardware/radio/voice/EmergencyCallRouting.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-cpp-source/gen/android/hardware/radio/voice/EmergencyNumber.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-cpp-source/gen/android/hardware/radio/voice/EmergencyServiceCategory.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-cpp-source/gen/android/hardware/radio/voice/IRadioVoice.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-cpp-source/gen/android/hardware/radio/voice/IRadioVoiceIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-cpp-source/gen/android/hardware/radio/voice/IRadioVoiceResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-cpp-source/gen/android/hardware/radio/voice/LastCallFailCause.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-cpp-source/gen/android/hardware/radio/voice/LastCallFailCauseInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-cpp-source/gen/android/hardware/radio/voice/SrvccState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-cpp-source/gen/android/hardware/radio/voice/SsInfoData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-cpp-source/gen/android/hardware/radio/voice/StkCcUnsolSsResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-cpp-source/gen/android/hardware/radio/voice/TtyMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-cpp-source/gen/android/hardware/radio/voice/UssdModeType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-cpp-source/gen/android/hardware/radio/voice/UusInfo.cpp
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
    add_library(android.hardware.radio.voice-V3-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.voice-V3-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.voice-V3-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.voice-V3-cpp
            android.hardware.radio-V2-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libm
            libutils
    )
endif ()
#################### android.hardware.radio.voice-V3-cpp ####################

#################### android.hardware.radio-V2-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-ndk-source/gen/android/hardware/radio/AccessNetwork.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-ndk-source/gen/android/hardware/radio/RadioAccessFamily.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-ndk-source/gen/android/hardware/radio/RadioConst.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-ndk-source/gen/android/hardware/radio/RadioError.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-ndk-source/gen/android/hardware/radio/RadioIndicationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-ndk-source/gen/android/hardware/radio/RadioResponseInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-ndk-source/gen/android/hardware/radio/RadioResponseInfoModem.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-ndk-source/gen/android/hardware/radio/RadioResponseType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-ndk-source/gen/android/hardware/radio/RadioTechnology.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-ndk-source/gen/android/hardware/radio/RadioTechnologyFamily.cpp
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
    add_library(android.hardware.radio-V2-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio-V2-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio-V2-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio-V2-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.radio-V2-ndk ####################

#################### android.hardware.radio.sap-V2-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sap-V2-ndk-source/gen/android/hardware/radio/sap/ISap.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sap-V2-ndk-source/gen/android/hardware/radio/sap/ISapCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sap-V2-ndk-source/gen/android/hardware/radio/sap/SapApduType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sap-V2-ndk-source/gen/android/hardware/radio/sap/SapConnectRsp.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sap-V2-ndk-source/gen/android/hardware/radio/sap/SapDisconnectType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sap-V2-ndk-source/gen/android/hardware/radio/sap/SapResultCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sap-V2-ndk-source/gen/android/hardware/radio/sap/SapStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sap-V2-ndk-source/gen/android/hardware/radio/sap/SapTransferProtocol.cpp
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
    add_library(android.hardware.radio.sap-V2-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.sap-V2-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sap-V2-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.sap-V2-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.sap-V2-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.radio.sap-V2-ndk ####################

#################### android.hardware.radio.modem-V2-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V2-ndk-source/gen/android/hardware/radio/modem/ActivityStatsInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V2-ndk-source/gen/android/hardware/radio/modem/ActivityStatsTechSpecificInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V2-ndk-source/gen/android/hardware/radio/modem/DeviceStateType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V2-ndk-source/gen/android/hardware/radio/modem/HardwareConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V2-ndk-source/gen/android/hardware/radio/modem/HardwareConfigModem.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V2-ndk-source/gen/android/hardware/radio/modem/HardwareConfigSim.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V2-ndk-source/gen/android/hardware/radio/modem/IRadioModem.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V2-ndk-source/gen/android/hardware/radio/modem/IRadioModemIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V2-ndk-source/gen/android/hardware/radio/modem/IRadioModemResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V2-ndk-source/gen/android/hardware/radio/modem/ImeiInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V2-ndk-source/gen/android/hardware/radio/modem/NvItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V2-ndk-source/gen/android/hardware/radio/modem/NvWriteItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V2-ndk-source/gen/android/hardware/radio/modem/RadioCapability.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V2-ndk-source/gen/android/hardware/radio/modem/RadioState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V2-ndk-source/gen/android/hardware/radio/modem/ResetNvType.cpp
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
    add_library(android.hardware.radio.modem-V2-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.modem-V2-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-ndk-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V2-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.modem-V2-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.modem-V2-ndk
            android.hardware.radio-V2-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.radio.modem-V2-ndk ####################

#################### android.hardware.radio.data-V3-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-ndk-source/gen/android/hardware/radio/data/ApnAuthType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-ndk-source/gen/android/hardware/radio/data/ApnTypes.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-ndk-source/gen/android/hardware/radio/data/DataCallFailCause.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-ndk-source/gen/android/hardware/radio/data/DataProfileInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-ndk-source/gen/android/hardware/radio/data/DataRequestReason.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-ndk-source/gen/android/hardware/radio/data/DataThrottlingAction.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-ndk-source/gen/android/hardware/radio/data/EpsQos.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-ndk-source/gen/android/hardware/radio/data/IRadioData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-ndk-source/gen/android/hardware/radio/data/IRadioDataIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-ndk-source/gen/android/hardware/radio/data/IRadioDataResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-ndk-source/gen/android/hardware/radio/data/KeepaliveRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-ndk-source/gen/android/hardware/radio/data/KeepaliveStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-ndk-source/gen/android/hardware/radio/data/LinkAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-ndk-source/gen/android/hardware/radio/data/NrQos.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-ndk-source/gen/android/hardware/radio/data/OsAppId.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-ndk-source/gen/android/hardware/radio/data/PcoDataInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-ndk-source/gen/android/hardware/radio/data/PdpProtocolType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-ndk-source/gen/android/hardware/radio/data/PortRange.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-ndk-source/gen/android/hardware/radio/data/Qos.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-ndk-source/gen/android/hardware/radio/data/QosBandwidth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-ndk-source/gen/android/hardware/radio/data/QosFilter.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-ndk-source/gen/android/hardware/radio/data/QosFilterIpsecSpi.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-ndk-source/gen/android/hardware/radio/data/QosFilterIpv6FlowLabel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-ndk-source/gen/android/hardware/radio/data/QosFilterTypeOfService.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-ndk-source/gen/android/hardware/radio/data/QosSession.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-ndk-source/gen/android/hardware/radio/data/RouteSelectionDescriptor.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-ndk-source/gen/android/hardware/radio/data/SetupDataCallResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-ndk-source/gen/android/hardware/radio/data/SliceInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-ndk-source/gen/android/hardware/radio/data/SlicingConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-ndk-source/gen/android/hardware/radio/data/TrafficDescriptor.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-ndk-source/gen/android/hardware/radio/data/UrspRule.cpp
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
    add_library(android.hardware.radio.data-V3-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.data-V3-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-ndk-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V3-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.data-V3-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.data-V3-ndk
            android.hardware.radio-V2-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.radio.data-V3-ndk ####################

#################### android.hardware.radio.ims-V2-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims-V2-ndk_platform-source/gen/android/hardware/radio/ims/ConnectionFailureInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims-V2-ndk_platform-source/gen/android/hardware/radio/ims/EpsFallbackReason.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims-V2-ndk_platform-source/gen/android/hardware/radio/ims/IRadioIms.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims-V2-ndk_platform-source/gen/android/hardware/radio/ims/IRadioImsIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims-V2-ndk_platform-source/gen/android/hardware/radio/ims/IRadioImsResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims-V2-ndk_platform-source/gen/android/hardware/radio/ims/ImsCall.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims-V2-ndk_platform-source/gen/android/hardware/radio/ims/ImsDeregistrationReason.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims-V2-ndk_platform-source/gen/android/hardware/radio/ims/ImsRegistration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims-V2-ndk_platform-source/gen/android/hardware/radio/ims/ImsRegistrationState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims-V2-ndk_platform-source/gen/android/hardware/radio/ims/ImsStreamDirection.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims-V2-ndk_platform-source/gen/android/hardware/radio/ims/ImsStreamType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims-V2-ndk_platform-source/gen/android/hardware/radio/ims/ImsTrafficType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims-V2-ndk_platform-source/gen/android/hardware/radio/ims/SrvccCall.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims-V2-ndk_platform-source/gen/android/hardware/radio/ims/SuggestedAction.cpp
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
    add_library(android.hardware.radio.ims-V2-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.ims-V2-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-ndk_platform-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims-V2-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.ims-V2-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.ims-V2-ndk_platform
            android.hardware.radio-V2-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.radio.ims-V2-ndk_platform ####################

#################### android.hardware.radio-V3-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V3-ndk_platform-source/gen/android/hardware/radio/AccessNetwork.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V3-ndk_platform-source/gen/android/hardware/radio/RadioAccessFamily.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V3-ndk_platform-source/gen/android/hardware/radio/RadioConst.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V3-ndk_platform-source/gen/android/hardware/radio/RadioError.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V3-ndk_platform-source/gen/android/hardware/radio/RadioIndicationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V3-ndk_platform-source/gen/android/hardware/radio/RadioResponseInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V3-ndk_platform-source/gen/android/hardware/radio/RadioResponseInfoModem.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V3-ndk_platform-source/gen/android/hardware/radio/RadioResponseType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V3-ndk_platform-source/gen/android/hardware/radio/RadioTechnology.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V3-ndk_platform-source/gen/android/hardware/radio/RadioTechnologyFamily.cpp
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
    add_library(android.hardware.radio-V3-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio-V3-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V3-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio-V3-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio-V3-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.radio-V3-ndk_platform ####################

#################### android.hardware.radio.data-V1-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-cpp-source/gen/android/hardware/radio/data/ApnAuthType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-cpp-source/gen/android/hardware/radio/data/ApnTypes.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-cpp-source/gen/android/hardware/radio/data/DataCallFailCause.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-cpp-source/gen/android/hardware/radio/data/DataProfileInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-cpp-source/gen/android/hardware/radio/data/DataRequestReason.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-cpp-source/gen/android/hardware/radio/data/DataThrottlingAction.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-cpp-source/gen/android/hardware/radio/data/EpsQos.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-cpp-source/gen/android/hardware/radio/data/IRadioData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-cpp-source/gen/android/hardware/radio/data/IRadioDataIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-cpp-source/gen/android/hardware/radio/data/IRadioDataResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-cpp-source/gen/android/hardware/radio/data/KeepaliveRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-cpp-source/gen/android/hardware/radio/data/KeepaliveStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-cpp-source/gen/android/hardware/radio/data/LinkAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-cpp-source/gen/android/hardware/radio/data/NrQos.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-cpp-source/gen/android/hardware/radio/data/OsAppId.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-cpp-source/gen/android/hardware/radio/data/PcoDataInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-cpp-source/gen/android/hardware/radio/data/PdpProtocolType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-cpp-source/gen/android/hardware/radio/data/PortRange.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-cpp-source/gen/android/hardware/radio/data/Qos.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-cpp-source/gen/android/hardware/radio/data/QosBandwidth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-cpp-source/gen/android/hardware/radio/data/QosFilter.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-cpp-source/gen/android/hardware/radio/data/QosFilterIpsecSpi.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-cpp-source/gen/android/hardware/radio/data/QosFilterIpv6FlowLabel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-cpp-source/gen/android/hardware/radio/data/QosFilterTypeOfService.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-cpp-source/gen/android/hardware/radio/data/QosSession.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-cpp-source/gen/android/hardware/radio/data/RouteSelectionDescriptor.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-cpp-source/gen/android/hardware/radio/data/SetupDataCallResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-cpp-source/gen/android/hardware/radio/data/SliceInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-cpp-source/gen/android/hardware/radio/data/SlicingConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-cpp-source/gen/android/hardware/radio/data/TrafficDescriptor.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-cpp-source/gen/android/hardware/radio/data/UrspRule.cpp
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
    add_library(android.hardware.radio.data-V1-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.data-V1-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V1-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.data-V1-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.data-V1-cpp
            android.hardware.radio-V1-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libm
            libutils
    )
endif ()
#################### android.hardware.radio.data-V1-cpp ####################

#################### android.hardware.radio.voice-V2-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-cpp-source/gen/android/hardware/radio/voice/AudioQuality.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-cpp-source/gen/android/hardware/radio/voice/Call.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-cpp-source/gen/android/hardware/radio/voice/CallForwardInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-cpp-source/gen/android/hardware/radio/voice/CdmaCallWaiting.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-cpp-source/gen/android/hardware/radio/voice/CdmaDisplayInfoRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-cpp-source/gen/android/hardware/radio/voice/CdmaInformationRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-cpp-source/gen/android/hardware/radio/voice/CdmaLineControlInfoRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-cpp-source/gen/android/hardware/radio/voice/CdmaNumberInfoRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-cpp-source/gen/android/hardware/radio/voice/CdmaOtaProvisionStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-cpp-source/gen/android/hardware/radio/voice/CdmaRedirectingNumberInfoRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-cpp-source/gen/android/hardware/radio/voice/CdmaSignalInfoRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-cpp-source/gen/android/hardware/radio/voice/CdmaT53AudioControlInfoRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-cpp-source/gen/android/hardware/radio/voice/CdmaT53ClirInfoRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-cpp-source/gen/android/hardware/radio/voice/CfData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-cpp-source/gen/android/hardware/radio/voice/ClipStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-cpp-source/gen/android/hardware/radio/voice/Dial.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-cpp-source/gen/android/hardware/radio/voice/EmergencyCallRouting.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-cpp-source/gen/android/hardware/radio/voice/EmergencyNumber.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-cpp-source/gen/android/hardware/radio/voice/EmergencyServiceCategory.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-cpp-source/gen/android/hardware/radio/voice/IRadioVoice.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-cpp-source/gen/android/hardware/radio/voice/IRadioVoiceIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-cpp-source/gen/android/hardware/radio/voice/IRadioVoiceResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-cpp-source/gen/android/hardware/radio/voice/LastCallFailCause.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-cpp-source/gen/android/hardware/radio/voice/LastCallFailCauseInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-cpp-source/gen/android/hardware/radio/voice/SrvccState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-cpp-source/gen/android/hardware/radio/voice/SsInfoData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-cpp-source/gen/android/hardware/radio/voice/StkCcUnsolSsResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-cpp-source/gen/android/hardware/radio/voice/TtyMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-cpp-source/gen/android/hardware/radio/voice/UssdModeType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-cpp-source/gen/android/hardware/radio/voice/UusInfo.cpp
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
    add_library(android.hardware.radio.voice-V2-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.voice-V2-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V2-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.voice-V2-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.voice-V2-cpp
            android.hardware.radio-V2-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libm
            libutils
    )
endif ()
#################### android.hardware.radio.voice-V2-cpp ####################

#################### android.hardware.radio.modem-V1-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V1-ndk_platform-source/gen/android/hardware/radio/modem/ActivityStatsInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V1-ndk_platform-source/gen/android/hardware/radio/modem/ActivityStatsTechSpecificInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V1-ndk_platform-source/gen/android/hardware/radio/modem/DeviceStateType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V1-ndk_platform-source/gen/android/hardware/radio/modem/HardwareConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V1-ndk_platform-source/gen/android/hardware/radio/modem/HardwareConfigModem.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V1-ndk_platform-source/gen/android/hardware/radio/modem/HardwareConfigSim.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V1-ndk_platform-source/gen/android/hardware/radio/modem/IRadioModem.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V1-ndk_platform-source/gen/android/hardware/radio/modem/IRadioModemIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V1-ndk_platform-source/gen/android/hardware/radio/modem/IRadioModemResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V1-ndk_platform-source/gen/android/hardware/radio/modem/NvItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V1-ndk_platform-source/gen/android/hardware/radio/modem/NvWriteItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V1-ndk_platform-source/gen/android/hardware/radio/modem/RadioCapability.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V1-ndk_platform-source/gen/android/hardware/radio/modem/RadioState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V1-ndk_platform-source/gen/android/hardware/radio/modem/ResetNvType.cpp
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
    add_library(android.hardware.radio.modem-V1-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.modem-V1-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V1-ndk_platform-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V1-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.modem-V1-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.modem-V1-ndk_platform
            android.hardware.radio-V1-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.radio.modem-V1-ndk_platform ####################

#################### android.hardware.radio.voice-V3-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-ndk_platform-source/gen/android/hardware/radio/voice/AudioQuality.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-ndk_platform-source/gen/android/hardware/radio/voice/Call.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-ndk_platform-source/gen/android/hardware/radio/voice/CallForwardInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-ndk_platform-source/gen/android/hardware/radio/voice/CdmaCallWaiting.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-ndk_platform-source/gen/android/hardware/radio/voice/CdmaDisplayInfoRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-ndk_platform-source/gen/android/hardware/radio/voice/CdmaInformationRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-ndk_platform-source/gen/android/hardware/radio/voice/CdmaLineControlInfoRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-ndk_platform-source/gen/android/hardware/radio/voice/CdmaNumberInfoRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-ndk_platform-source/gen/android/hardware/radio/voice/CdmaOtaProvisionStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-ndk_platform-source/gen/android/hardware/radio/voice/CdmaRedirectingNumberInfoRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-ndk_platform-source/gen/android/hardware/radio/voice/CdmaSignalInfoRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-ndk_platform-source/gen/android/hardware/radio/voice/CdmaT53AudioControlInfoRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-ndk_platform-source/gen/android/hardware/radio/voice/CdmaT53ClirInfoRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-ndk_platform-source/gen/android/hardware/radio/voice/CfData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-ndk_platform-source/gen/android/hardware/radio/voice/ClipStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-ndk_platform-source/gen/android/hardware/radio/voice/Dial.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-ndk_platform-source/gen/android/hardware/radio/voice/EmergencyCallRouting.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-ndk_platform-source/gen/android/hardware/radio/voice/EmergencyNumber.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-ndk_platform-source/gen/android/hardware/radio/voice/EmergencyServiceCategory.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-ndk_platform-source/gen/android/hardware/radio/voice/IRadioVoice.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-ndk_platform-source/gen/android/hardware/radio/voice/IRadioVoiceIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-ndk_platform-source/gen/android/hardware/radio/voice/IRadioVoiceResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-ndk_platform-source/gen/android/hardware/radio/voice/LastCallFailCause.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-ndk_platform-source/gen/android/hardware/radio/voice/LastCallFailCauseInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-ndk_platform-source/gen/android/hardware/radio/voice/SrvccState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-ndk_platform-source/gen/android/hardware/radio/voice/SsInfoData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-ndk_platform-source/gen/android/hardware/radio/voice/StkCcUnsolSsResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-ndk_platform-source/gen/android/hardware/radio/voice/TtyMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-ndk_platform-source/gen/android/hardware/radio/voice/UssdModeType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-ndk_platform-source/gen/android/hardware/radio/voice/UusInfo.cpp
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
    add_library(android.hardware.radio.voice-V3-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.voice-V3-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-ndk_platform-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V3-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.voice-V3-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.voice-V3-ndk_platform
            android.hardware.radio-V2-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.radio.voice-V3-ndk_platform ####################

#################### android.hardware.radio.config-V3-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V3-cpp-source/gen/android/hardware/radio/config/IRadioConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V3-cpp-source/gen/android/hardware/radio/config/IRadioConfigIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V3-cpp-source/gen/android/hardware/radio/config/IRadioConfigResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V3-cpp-source/gen/android/hardware/radio/config/MultipleEnabledProfilesMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V3-cpp-source/gen/android/hardware/radio/config/PhoneCapability.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V3-cpp-source/gen/android/hardware/radio/config/SimPortInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V3-cpp-source/gen/android/hardware/radio/config/SimSlotStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V3-cpp-source/gen/android/hardware/radio/config/SlotPortMapping.cpp
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
    add_library(android.hardware.radio.config-V3-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.config-V3-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V3-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.config-V3-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.config-V3-cpp
            android.hardware.radio-V2-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libm
            libutils
    )
endif ()
#################### android.hardware.radio.config-V3-cpp ####################

#################### android.hardware.radio.ims.media-V2-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V2-ndk-source/gen/android/hardware/radio/ims/media/AmrMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V2-ndk-source/gen/android/hardware/radio/ims/media/AmrParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V2-ndk-source/gen/android/hardware/radio/ims/media/AnbrMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V2-ndk-source/gen/android/hardware/radio/ims/media/CallQuality.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V2-ndk-source/gen/android/hardware/radio/ims/media/CodecMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V2-ndk-source/gen/android/hardware/radio/ims/media/CodecParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V2-ndk-source/gen/android/hardware/radio/ims/media/CodecSpecificParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V2-ndk-source/gen/android/hardware/radio/ims/media/CodecType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V2-ndk-source/gen/android/hardware/radio/ims/media/DtmfParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V2-ndk-source/gen/android/hardware/radio/ims/media/EvsBandwidth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V2-ndk-source/gen/android/hardware/radio/ims/media/EvsMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V2-ndk-source/gen/android/hardware/radio/ims/media/EvsParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V2-ndk-source/gen/android/hardware/radio/ims/media/IImsMedia.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V2-ndk-source/gen/android/hardware/radio/ims/media/IImsMediaListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V2-ndk-source/gen/android/hardware/radio/ims/media/IImsMediaSession.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V2-ndk-source/gen/android/hardware/radio/ims/media/IImsMediaSessionListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V2-ndk-source/gen/android/hardware/radio/ims/media/LocalEndPoint.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V2-ndk-source/gen/android/hardware/radio/ims/media/MediaDirection.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V2-ndk-source/gen/android/hardware/radio/ims/media/MediaQualityStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V2-ndk-source/gen/android/hardware/radio/ims/media/MediaQualityThreshold.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V2-ndk-source/gen/android/hardware/radio/ims/media/RtcpConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V2-ndk-source/gen/android/hardware/radio/ims/media/RtcpXrReportBlockType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V2-ndk-source/gen/android/hardware/radio/ims/media/RtpAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V2-ndk-source/gen/android/hardware/radio/ims/media/RtpConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V2-ndk-source/gen/android/hardware/radio/ims/media/RtpError.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V2-ndk-source/gen/android/hardware/radio/ims/media/RtpHeaderExtension.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V2-ndk-source/gen/android/hardware/radio/ims/media/RtpSessionParams.cpp
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
    add_library(android.hardware.radio.ims.media-V2-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.ims.media-V2-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-ndk-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V2-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.ims.media-V2-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.ims.media-V2-ndk
            android.hardware.radio-V2-ndk
            android.hardware.radio.data-V2-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.radio.ims.media-V2-ndk ####################

#################### android.hardware.radio.network-V3-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk_platform-source/gen/android/hardware/radio/network/AccessTechnologySpecificInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk_platform-source/gen/android/hardware/radio/network/BarringInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk_platform-source/gen/android/hardware/radio/network/BarringTypeSpecificInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk_platform-source/gen/android/hardware/radio/network/Cdma2000RegistrationInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk_platform-source/gen/android/hardware/radio/network/CdmaRoamingType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk_platform-source/gen/android/hardware/radio/network/CdmaSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk_platform-source/gen/android/hardware/radio/network/CellConnectionStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk_platform-source/gen/android/hardware/radio/network/CellIdentity.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk_platform-source/gen/android/hardware/radio/network/CellIdentityCdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk_platform-source/gen/android/hardware/radio/network/CellIdentityGsm.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk_platform-source/gen/android/hardware/radio/network/CellIdentityLte.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk_platform-source/gen/android/hardware/radio/network/CellIdentityNr.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk_platform-source/gen/android/hardware/radio/network/CellIdentityTdscdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk_platform-source/gen/android/hardware/radio/network/CellIdentityWcdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk_platform-source/gen/android/hardware/radio/network/CellInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk_platform-source/gen/android/hardware/radio/network/CellInfoCdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk_platform-source/gen/android/hardware/radio/network/CellInfoGsm.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk_platform-source/gen/android/hardware/radio/network/CellInfoLte.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk_platform-source/gen/android/hardware/radio/network/CellInfoNr.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk_platform-source/gen/android/hardware/radio/network/CellInfoRatSpecificInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk_platform-source/gen/android/hardware/radio/network/CellInfoTdscdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk_platform-source/gen/android/hardware/radio/network/CellInfoWcdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk_platform-source/gen/android/hardware/radio/network/ClosedSubscriberGroupInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk_platform-source/gen/android/hardware/radio/network/Domain.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk_platform-source/gen/android/hardware/radio/network/EmergencyMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk_platform-source/gen/android/hardware/radio/network/EmergencyNetworkScanTrigger.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk_platform-source/gen/android/hardware/radio/network/EmergencyRegResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk_platform-source/gen/android/hardware/radio/network/EmergencyScanType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk_platform-source/gen/android/hardware/radio/network/EutranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk_platform-source/gen/android/hardware/radio/network/EutranRegistrationInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk_platform-source/gen/android/hardware/radio/network/EvdoSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk_platform-source/gen/android/hardware/radio/network/GeranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk_platform-source/gen/android/hardware/radio/network/GsmSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk_platform-source/gen/android/hardware/radio/network/IRadioNetwork.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk_platform-source/gen/android/hardware/radio/network/IRadioNetworkIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk_platform-source/gen/android/hardware/radio/network/IRadioNetworkResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk_platform-source/gen/android/hardware/radio/network/IndicationFilter.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk_platform-source/gen/android/hardware/radio/network/LceDataInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk_platform-source/gen/android/hardware/radio/network/LinkCapacityEstimate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk_platform-source/gen/android/hardware/radio/network/LteSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk_platform-source/gen/android/hardware/radio/network/LteVopsInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk_platform-source/gen/android/hardware/radio/network/NetworkScanRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk_platform-source/gen/android/hardware/radio/network/NetworkScanResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk_platform-source/gen/android/hardware/radio/network/NgranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk_platform-source/gen/android/hardware/radio/network/NrDualConnectivityState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk_platform-source/gen/android/hardware/radio/network/NrIndicators.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk_platform-source/gen/android/hardware/radio/network/NrSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk_platform-source/gen/android/hardware/radio/network/NrVopsInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk_platform-source/gen/android/hardware/radio/network/OperatorInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk_platform-source/gen/android/hardware/radio/network/PhoneRestrictedState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk_platform-source/gen/android/hardware/radio/network/PhysicalChannelConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk_platform-source/gen/android/hardware/radio/network/PhysicalChannelConfigBand.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk_platform-source/gen/android/hardware/radio/network/RadioAccessSpecifier.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk_platform-source/gen/android/hardware/radio/network/RadioAccessSpecifierBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk_platform-source/gen/android/hardware/radio/network/RadioBandMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk_platform-source/gen/android/hardware/radio/network/RegState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk_platform-source/gen/android/hardware/radio/network/RegStateResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk_platform-source/gen/android/hardware/radio/network/RegistrationFailCause.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk_platform-source/gen/android/hardware/radio/network/SignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk_platform-source/gen/android/hardware/radio/network/SignalThresholdInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk_platform-source/gen/android/hardware/radio/network/SuppSvcNotification.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk_platform-source/gen/android/hardware/radio/network/TdscdmaSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk_platform-source/gen/android/hardware/radio/network/UsageSetting.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk_platform-source/gen/android/hardware/radio/network/UtranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk_platform-source/gen/android/hardware/radio/network/WcdmaSignalStrength.cpp
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
    add_library(android.hardware.radio.network-V3-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.network-V3-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-ndk_platform-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.network-V3-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.network-V3-ndk_platform
            android.hardware.radio-V2-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.radio.network-V3-ndk_platform ####################

#################### android.hardware.radio.config-V1-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V1-ndk-source/gen/android/hardware/radio/config/IRadioConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V1-ndk-source/gen/android/hardware/radio/config/IRadioConfigIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V1-ndk-source/gen/android/hardware/radio/config/IRadioConfigResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V1-ndk-source/gen/android/hardware/radio/config/PhoneCapability.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V1-ndk-source/gen/android/hardware/radio/config/SimPortInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V1-ndk-source/gen/android/hardware/radio/config/SimSlotStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V1-ndk-source/gen/android/hardware/radio/config/SlotPortMapping.cpp
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
    add_library(android.hardware.radio.config-V1-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.config-V1-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V1-ndk-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V1-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.config-V1-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.config-V1-ndk
            android.hardware.radio-V1-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.radio.config-V1-ndk ####################

#################### android.hardware.radio.sim-V3-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-ndk-source/gen/android/hardware/radio/sim/AppStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-ndk-source/gen/android/hardware/radio/sim/CardPowerState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-ndk-source/gen/android/hardware/radio/sim/CardStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-ndk-source/gen/android/hardware/radio/sim/Carrier.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-ndk-source/gen/android/hardware/radio/sim/CarrierRestrictions.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-ndk-source/gen/android/hardware/radio/sim/CdmaSubscriptionSource.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-ndk-source/gen/android/hardware/radio/sim/IRadioSim.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-ndk-source/gen/android/hardware/radio/sim/IRadioSimIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-ndk-source/gen/android/hardware/radio/sim/IRadioSimResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-ndk-source/gen/android/hardware/radio/sim/IccIo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-ndk-source/gen/android/hardware/radio/sim/IccIoResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-ndk-source/gen/android/hardware/radio/sim/ImsiEncryptionInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-ndk-source/gen/android/hardware/radio/sim/PbReceivedStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-ndk-source/gen/android/hardware/radio/sim/PersoSubstate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-ndk-source/gen/android/hardware/radio/sim/PhonebookCapacity.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-ndk-source/gen/android/hardware/radio/sim/PhonebookRecordInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-ndk-source/gen/android/hardware/radio/sim/PinState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-ndk-source/gen/android/hardware/radio/sim/SelectUiccSub.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-ndk-source/gen/android/hardware/radio/sim/SessionInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-ndk-source/gen/android/hardware/radio/sim/SimApdu.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-ndk-source/gen/android/hardware/radio/sim/SimLockMultiSimPolicy.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-ndk-source/gen/android/hardware/radio/sim/SimRefreshResult.cpp
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
    add_library(android.hardware.radio.sim-V3-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.sim-V3-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-ndk-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V2-ndk-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.sim-V3-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.sim-V3-ndk
            android.hardware.radio-V2-ndk
            android.hardware.radio.config-V2-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.radio.sim-V3-ndk ####################

#################### android.hardware.radio.network-V1-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk-source/gen/android/hardware/radio/network/AccessTechnologySpecificInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk-source/gen/android/hardware/radio/network/BarringInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk-source/gen/android/hardware/radio/network/BarringTypeSpecificInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk-source/gen/android/hardware/radio/network/Cdma2000RegistrationInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk-source/gen/android/hardware/radio/network/CdmaRoamingType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk-source/gen/android/hardware/radio/network/CdmaSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk-source/gen/android/hardware/radio/network/CellConnectionStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk-source/gen/android/hardware/radio/network/CellIdentity.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk-source/gen/android/hardware/radio/network/CellIdentityCdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk-source/gen/android/hardware/radio/network/CellIdentityGsm.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk-source/gen/android/hardware/radio/network/CellIdentityLte.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk-source/gen/android/hardware/radio/network/CellIdentityNr.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk-source/gen/android/hardware/radio/network/CellIdentityTdscdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk-source/gen/android/hardware/radio/network/CellIdentityWcdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk-source/gen/android/hardware/radio/network/CellInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk-source/gen/android/hardware/radio/network/CellInfoCdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk-source/gen/android/hardware/radio/network/CellInfoGsm.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk-source/gen/android/hardware/radio/network/CellInfoLte.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk-source/gen/android/hardware/radio/network/CellInfoNr.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk-source/gen/android/hardware/radio/network/CellInfoRatSpecificInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk-source/gen/android/hardware/radio/network/CellInfoTdscdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk-source/gen/android/hardware/radio/network/CellInfoWcdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk-source/gen/android/hardware/radio/network/ClosedSubscriberGroupInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk-source/gen/android/hardware/radio/network/Domain.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk-source/gen/android/hardware/radio/network/EutranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk-source/gen/android/hardware/radio/network/EutranRegistrationInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk-source/gen/android/hardware/radio/network/EvdoSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk-source/gen/android/hardware/radio/network/GeranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk-source/gen/android/hardware/radio/network/GsmSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk-source/gen/android/hardware/radio/network/IRadioNetwork.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk-source/gen/android/hardware/radio/network/IRadioNetworkIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk-source/gen/android/hardware/radio/network/IRadioNetworkResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk-source/gen/android/hardware/radio/network/IndicationFilter.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk-source/gen/android/hardware/radio/network/LceDataInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk-source/gen/android/hardware/radio/network/LinkCapacityEstimate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk-source/gen/android/hardware/radio/network/LteSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk-source/gen/android/hardware/radio/network/LteVopsInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk-source/gen/android/hardware/radio/network/NetworkScanRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk-source/gen/android/hardware/radio/network/NetworkScanResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk-source/gen/android/hardware/radio/network/NgranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk-source/gen/android/hardware/radio/network/NrDualConnectivityState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk-source/gen/android/hardware/radio/network/NrIndicators.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk-source/gen/android/hardware/radio/network/NrSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk-source/gen/android/hardware/radio/network/NrVopsInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk-source/gen/android/hardware/radio/network/OperatorInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk-source/gen/android/hardware/radio/network/PhoneRestrictedState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk-source/gen/android/hardware/radio/network/PhysicalChannelConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk-source/gen/android/hardware/radio/network/PhysicalChannelConfigBand.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk-source/gen/android/hardware/radio/network/RadioAccessSpecifier.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk-source/gen/android/hardware/radio/network/RadioAccessSpecifierBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk-source/gen/android/hardware/radio/network/RadioBandMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk-source/gen/android/hardware/radio/network/RegState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk-source/gen/android/hardware/radio/network/RegStateResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk-source/gen/android/hardware/radio/network/RegistrationFailCause.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk-source/gen/android/hardware/radio/network/SignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk-source/gen/android/hardware/radio/network/SignalThresholdInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk-source/gen/android/hardware/radio/network/SuppSvcNotification.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk-source/gen/android/hardware/radio/network/TdscdmaSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk-source/gen/android/hardware/radio/network/UsageSetting.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk-source/gen/android/hardware/radio/network/UtranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk-source/gen/android/hardware/radio/network/WcdmaSignalStrength.cpp
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
    add_library(android.hardware.radio.network-V1-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.network-V1-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V1-ndk-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.network-V1-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.network-V1-ndk
            android.hardware.radio-V1-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.radio.network-V1-ndk ####################

#################### android.hardware.radio.messaging-V3-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V3-ndk_platform-source/gen/android/hardware/radio/messaging/CdmaBroadcastSmsConfigInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V3-ndk_platform-source/gen/android/hardware/radio/messaging/CdmaSmsAck.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V3-ndk_platform-source/gen/android/hardware/radio/messaging/CdmaSmsAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V3-ndk_platform-source/gen/android/hardware/radio/messaging/CdmaSmsMessage.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V3-ndk_platform-source/gen/android/hardware/radio/messaging/CdmaSmsSubaddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V3-ndk_platform-source/gen/android/hardware/radio/messaging/CdmaSmsWriteArgs.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V3-ndk_platform-source/gen/android/hardware/radio/messaging/GsmBroadcastSmsConfigInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V3-ndk_platform-source/gen/android/hardware/radio/messaging/GsmSmsMessage.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V3-ndk_platform-source/gen/android/hardware/radio/messaging/IRadioMessaging.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V3-ndk_platform-source/gen/android/hardware/radio/messaging/IRadioMessagingIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V3-ndk_platform-source/gen/android/hardware/radio/messaging/IRadioMessagingResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V3-ndk_platform-source/gen/android/hardware/radio/messaging/ImsSmsMessage.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V3-ndk_platform-source/gen/android/hardware/radio/messaging/SendSmsResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V3-ndk_platform-source/gen/android/hardware/radio/messaging/SmsAcknowledgeFailCause.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V3-ndk_platform-source/gen/android/hardware/radio/messaging/SmsWriteArgs.cpp
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
    add_library(android.hardware.radio.messaging-V3-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.messaging-V3-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-ndk_platform-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V3-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.messaging-V3-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.messaging-V3-ndk_platform
            android.hardware.radio-V2-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.radio.messaging-V3-ndk_platform ####################

#################### android.hardware.radio.messaging-V2-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V2-ndk-source/gen/android/hardware/radio/messaging/CdmaBroadcastSmsConfigInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V2-ndk-source/gen/android/hardware/radio/messaging/CdmaSmsAck.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V2-ndk-source/gen/android/hardware/radio/messaging/CdmaSmsAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V2-ndk-source/gen/android/hardware/radio/messaging/CdmaSmsMessage.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V2-ndk-source/gen/android/hardware/radio/messaging/CdmaSmsSubaddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V2-ndk-source/gen/android/hardware/radio/messaging/CdmaSmsWriteArgs.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V2-ndk-source/gen/android/hardware/radio/messaging/GsmBroadcastSmsConfigInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V2-ndk-source/gen/android/hardware/radio/messaging/GsmSmsMessage.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V2-ndk-source/gen/android/hardware/radio/messaging/IRadioMessaging.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V2-ndk-source/gen/android/hardware/radio/messaging/IRadioMessagingIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V2-ndk-source/gen/android/hardware/radio/messaging/IRadioMessagingResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V2-ndk-source/gen/android/hardware/radio/messaging/ImsSmsMessage.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V2-ndk-source/gen/android/hardware/radio/messaging/SendSmsResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V2-ndk-source/gen/android/hardware/radio/messaging/SmsAcknowledgeFailCause.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V2-ndk-source/gen/android/hardware/radio/messaging/SmsWriteArgs.cpp
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
    add_library(android.hardware.radio.messaging-V2-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.messaging-V2-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-ndk-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V2-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.messaging-V2-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.messaging-V2-ndk
            android.hardware.radio-V2-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.radio.messaging-V2-ndk ####################

#################### android.hardware.radio.network-V3-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-cpp-source/gen/android/hardware/radio/network/AccessTechnologySpecificInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-cpp-source/gen/android/hardware/radio/network/BarringInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-cpp-source/gen/android/hardware/radio/network/BarringTypeSpecificInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-cpp-source/gen/android/hardware/radio/network/Cdma2000RegistrationInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-cpp-source/gen/android/hardware/radio/network/CdmaRoamingType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-cpp-source/gen/android/hardware/radio/network/CdmaSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-cpp-source/gen/android/hardware/radio/network/CellConnectionStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-cpp-source/gen/android/hardware/radio/network/CellIdentity.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-cpp-source/gen/android/hardware/radio/network/CellIdentityCdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-cpp-source/gen/android/hardware/radio/network/CellIdentityGsm.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-cpp-source/gen/android/hardware/radio/network/CellIdentityLte.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-cpp-source/gen/android/hardware/radio/network/CellIdentityNr.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-cpp-source/gen/android/hardware/radio/network/CellIdentityTdscdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-cpp-source/gen/android/hardware/radio/network/CellIdentityWcdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-cpp-source/gen/android/hardware/radio/network/CellInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-cpp-source/gen/android/hardware/radio/network/CellInfoCdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-cpp-source/gen/android/hardware/radio/network/CellInfoGsm.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-cpp-source/gen/android/hardware/radio/network/CellInfoLte.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-cpp-source/gen/android/hardware/radio/network/CellInfoNr.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-cpp-source/gen/android/hardware/radio/network/CellInfoRatSpecificInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-cpp-source/gen/android/hardware/radio/network/CellInfoTdscdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-cpp-source/gen/android/hardware/radio/network/CellInfoWcdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-cpp-source/gen/android/hardware/radio/network/ClosedSubscriberGroupInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-cpp-source/gen/android/hardware/radio/network/Domain.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-cpp-source/gen/android/hardware/radio/network/EmergencyMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-cpp-source/gen/android/hardware/radio/network/EmergencyNetworkScanTrigger.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-cpp-source/gen/android/hardware/radio/network/EmergencyRegResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-cpp-source/gen/android/hardware/radio/network/EmergencyScanType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-cpp-source/gen/android/hardware/radio/network/EutranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-cpp-source/gen/android/hardware/radio/network/EutranRegistrationInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-cpp-source/gen/android/hardware/radio/network/EvdoSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-cpp-source/gen/android/hardware/radio/network/GeranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-cpp-source/gen/android/hardware/radio/network/GsmSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-cpp-source/gen/android/hardware/radio/network/IRadioNetwork.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-cpp-source/gen/android/hardware/radio/network/IRadioNetworkIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-cpp-source/gen/android/hardware/radio/network/IRadioNetworkResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-cpp-source/gen/android/hardware/radio/network/IndicationFilter.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-cpp-source/gen/android/hardware/radio/network/LceDataInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-cpp-source/gen/android/hardware/radio/network/LinkCapacityEstimate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-cpp-source/gen/android/hardware/radio/network/LteSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-cpp-source/gen/android/hardware/radio/network/LteVopsInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-cpp-source/gen/android/hardware/radio/network/NetworkScanRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-cpp-source/gen/android/hardware/radio/network/NetworkScanResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-cpp-source/gen/android/hardware/radio/network/NgranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-cpp-source/gen/android/hardware/radio/network/NrDualConnectivityState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-cpp-source/gen/android/hardware/radio/network/NrIndicators.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-cpp-source/gen/android/hardware/radio/network/NrSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-cpp-source/gen/android/hardware/radio/network/NrVopsInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-cpp-source/gen/android/hardware/radio/network/OperatorInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-cpp-source/gen/android/hardware/radio/network/PhoneRestrictedState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-cpp-source/gen/android/hardware/radio/network/PhysicalChannelConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-cpp-source/gen/android/hardware/radio/network/PhysicalChannelConfigBand.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-cpp-source/gen/android/hardware/radio/network/RadioAccessSpecifier.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-cpp-source/gen/android/hardware/radio/network/RadioAccessSpecifierBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-cpp-source/gen/android/hardware/radio/network/RadioBandMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-cpp-source/gen/android/hardware/radio/network/RegState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-cpp-source/gen/android/hardware/radio/network/RegStateResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-cpp-source/gen/android/hardware/radio/network/RegistrationFailCause.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-cpp-source/gen/android/hardware/radio/network/SignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-cpp-source/gen/android/hardware/radio/network/SignalThresholdInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-cpp-source/gen/android/hardware/radio/network/SuppSvcNotification.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-cpp-source/gen/android/hardware/radio/network/TdscdmaSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-cpp-source/gen/android/hardware/radio/network/UsageSetting.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-cpp-source/gen/android/hardware/radio/network/UtranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-cpp-source/gen/android/hardware/radio/network/WcdmaSignalStrength.cpp
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
    add_library(android.hardware.radio.network-V3-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.network-V3-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.network-V3-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.network-V3-cpp
            android.hardware.radio-V2-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libm
            libutils
    )
endif ()
#################### android.hardware.radio.network-V3-cpp ####################

#################### android.hardware.radio.sim-V1-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V1-cpp-source/gen/android/hardware/radio/sim/AppStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V1-cpp-source/gen/android/hardware/radio/sim/CardPowerState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V1-cpp-source/gen/android/hardware/radio/sim/CardStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V1-cpp-source/gen/android/hardware/radio/sim/Carrier.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V1-cpp-source/gen/android/hardware/radio/sim/CarrierRestrictions.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V1-cpp-source/gen/android/hardware/radio/sim/CdmaSubscriptionSource.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V1-cpp-source/gen/android/hardware/radio/sim/IRadioSim.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V1-cpp-source/gen/android/hardware/radio/sim/IRadioSimIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V1-cpp-source/gen/android/hardware/radio/sim/IRadioSimResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V1-cpp-source/gen/android/hardware/radio/sim/IccIo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V1-cpp-source/gen/android/hardware/radio/sim/IccIoResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V1-cpp-source/gen/android/hardware/radio/sim/ImsiEncryptionInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V1-cpp-source/gen/android/hardware/radio/sim/PbReceivedStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V1-cpp-source/gen/android/hardware/radio/sim/PersoSubstate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V1-cpp-source/gen/android/hardware/radio/sim/PhonebookCapacity.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V1-cpp-source/gen/android/hardware/radio/sim/PhonebookRecordInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V1-cpp-source/gen/android/hardware/radio/sim/PinState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V1-cpp-source/gen/android/hardware/radio/sim/SelectUiccSub.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V1-cpp-source/gen/android/hardware/radio/sim/SimApdu.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V1-cpp-source/gen/android/hardware/radio/sim/SimLockMultiSimPolicy.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V1-cpp-source/gen/android/hardware/radio/sim/SimRefreshResult.cpp
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
    add_library(android.hardware.radio.sim-V1-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.sim-V1-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V1-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V1-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V1-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.sim-V1-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.sim-V1-cpp
            android.hardware.radio-V1-cpp
            android.hardware.radio.config-V1-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libm
            libutils
    )
endif ()
#################### android.hardware.radio.sim-V1-cpp ####################

#################### android.hardware.radio.config-V2-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V2-cpp-source/gen/android/hardware/radio/config/IRadioConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V2-cpp-source/gen/android/hardware/radio/config/IRadioConfigIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V2-cpp-source/gen/android/hardware/radio/config/IRadioConfigResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V2-cpp-source/gen/android/hardware/radio/config/MultipleEnabledProfilesMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V2-cpp-source/gen/android/hardware/radio/config/PhoneCapability.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V2-cpp-source/gen/android/hardware/radio/config/SimPortInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V2-cpp-source/gen/android/hardware/radio/config/SimSlotStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V2-cpp-source/gen/android/hardware/radio/config/SlotPortMapping.cpp
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
    add_library(android.hardware.radio.config-V2-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.config-V2-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V2-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.config-V2-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.config-V2-cpp
            android.hardware.radio-V2-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libm
            libutils
    )
endif ()
#################### android.hardware.radio.config-V2-cpp ####################

#################### android.hardware.radio.sim-V2-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-ndk_platform-source/gen/android/hardware/radio/sim/AppStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-ndk_platform-source/gen/android/hardware/radio/sim/CardPowerState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-ndk_platform-source/gen/android/hardware/radio/sim/CardStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-ndk_platform-source/gen/android/hardware/radio/sim/Carrier.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-ndk_platform-source/gen/android/hardware/radio/sim/CarrierRestrictions.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-ndk_platform-source/gen/android/hardware/radio/sim/CdmaSubscriptionSource.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-ndk_platform-source/gen/android/hardware/radio/sim/IRadioSim.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-ndk_platform-source/gen/android/hardware/radio/sim/IRadioSimIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-ndk_platform-source/gen/android/hardware/radio/sim/IRadioSimResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-ndk_platform-source/gen/android/hardware/radio/sim/IccIo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-ndk_platform-source/gen/android/hardware/radio/sim/IccIoResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-ndk_platform-source/gen/android/hardware/radio/sim/ImsiEncryptionInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-ndk_platform-source/gen/android/hardware/radio/sim/PbReceivedStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-ndk_platform-source/gen/android/hardware/radio/sim/PersoSubstate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-ndk_platform-source/gen/android/hardware/radio/sim/PhonebookCapacity.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-ndk_platform-source/gen/android/hardware/radio/sim/PhonebookRecordInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-ndk_platform-source/gen/android/hardware/radio/sim/PinState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-ndk_platform-source/gen/android/hardware/radio/sim/SelectUiccSub.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-ndk_platform-source/gen/android/hardware/radio/sim/SessionInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-ndk_platform-source/gen/android/hardware/radio/sim/SimApdu.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-ndk_platform-source/gen/android/hardware/radio/sim/SimLockMultiSimPolicy.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-ndk_platform-source/gen/android/hardware/radio/sim/SimRefreshResult.cpp
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
    add_library(android.hardware.radio.sim-V2-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.sim-V2-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-ndk_platform-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V2-ndk_platform-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.sim-V2-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.sim-V2-ndk_platform
            android.hardware.radio-V2-ndk_platform
            android.hardware.radio.config-V2-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.radio.sim-V2-ndk_platform ####################

#################### android.hardware.radio-V1-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V1-ndk_platform-source/gen/android/hardware/radio/AccessNetwork.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V1-ndk_platform-source/gen/android/hardware/radio/RadioAccessFamily.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V1-ndk_platform-source/gen/android/hardware/radio/RadioConst.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V1-ndk_platform-source/gen/android/hardware/radio/RadioError.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V1-ndk_platform-source/gen/android/hardware/radio/RadioIndicationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V1-ndk_platform-source/gen/android/hardware/radio/RadioResponseInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V1-ndk_platform-source/gen/android/hardware/radio/RadioResponseInfoModem.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V1-ndk_platform-source/gen/android/hardware/radio/RadioResponseType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V1-ndk_platform-source/gen/android/hardware/radio/RadioTechnology.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V1-ndk_platform-source/gen/android/hardware/radio/RadioTechnologyFamily.cpp
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
    add_library(android.hardware.radio-V1-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio-V1-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V1-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio-V1-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio-V1-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.radio-V1-ndk_platform ####################

#################### android.hardware.radio.sim-V2-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-ndk-source/gen/android/hardware/radio/sim/AppStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-ndk-source/gen/android/hardware/radio/sim/CardPowerState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-ndk-source/gen/android/hardware/radio/sim/CardStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-ndk-source/gen/android/hardware/radio/sim/Carrier.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-ndk-source/gen/android/hardware/radio/sim/CarrierRestrictions.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-ndk-source/gen/android/hardware/radio/sim/CdmaSubscriptionSource.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-ndk-source/gen/android/hardware/radio/sim/IRadioSim.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-ndk-source/gen/android/hardware/radio/sim/IRadioSimIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-ndk-source/gen/android/hardware/radio/sim/IRadioSimResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-ndk-source/gen/android/hardware/radio/sim/IccIo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-ndk-source/gen/android/hardware/radio/sim/IccIoResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-ndk-source/gen/android/hardware/radio/sim/ImsiEncryptionInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-ndk-source/gen/android/hardware/radio/sim/PbReceivedStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-ndk-source/gen/android/hardware/radio/sim/PersoSubstate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-ndk-source/gen/android/hardware/radio/sim/PhonebookCapacity.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-ndk-source/gen/android/hardware/radio/sim/PhonebookRecordInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-ndk-source/gen/android/hardware/radio/sim/PinState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-ndk-source/gen/android/hardware/radio/sim/SelectUiccSub.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-ndk-source/gen/android/hardware/radio/sim/SessionInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-ndk-source/gen/android/hardware/radio/sim/SimApdu.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-ndk-source/gen/android/hardware/radio/sim/SimLockMultiSimPolicy.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-ndk-source/gen/android/hardware/radio/sim/SimRefreshResult.cpp
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
    add_library(android.hardware.radio.sim-V2-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.sim-V2-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-ndk-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V2-ndk-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.sim-V2-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.sim-V2-ndk
            android.hardware.radio-V2-ndk
            android.hardware.radio.config-V2-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.radio.sim-V2-ndk ####################

#################### android.hardware.radio.modem-V3-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V3-ndk_platform-source/gen/android/hardware/radio/modem/ActivityStatsInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V3-ndk_platform-source/gen/android/hardware/radio/modem/ActivityStatsTechSpecificInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V3-ndk_platform-source/gen/android/hardware/radio/modem/DeviceStateType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V3-ndk_platform-source/gen/android/hardware/radio/modem/HardwareConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V3-ndk_platform-source/gen/android/hardware/radio/modem/HardwareConfigModem.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V3-ndk_platform-source/gen/android/hardware/radio/modem/HardwareConfigSim.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V3-ndk_platform-source/gen/android/hardware/radio/modem/IRadioModem.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V3-ndk_platform-source/gen/android/hardware/radio/modem/IRadioModemIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V3-ndk_platform-source/gen/android/hardware/radio/modem/IRadioModemResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V3-ndk_platform-source/gen/android/hardware/radio/modem/ImeiInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V3-ndk_platform-source/gen/android/hardware/radio/modem/NvItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V3-ndk_platform-source/gen/android/hardware/radio/modem/NvWriteItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V3-ndk_platform-source/gen/android/hardware/radio/modem/RadioCapability.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V3-ndk_platform-source/gen/android/hardware/radio/modem/RadioState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V3-ndk_platform-source/gen/android/hardware/radio/modem/ResetNvType.cpp
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
    add_library(android.hardware.radio.modem-V3-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.modem-V3-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-ndk_platform-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V3-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.modem-V3-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.modem-V3-ndk_platform
            android.hardware.radio-V2-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.radio.modem-V3-ndk_platform ####################

#################### android.hardware.radio.config-V2-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V2-ndk_platform-source/gen/android/hardware/radio/config/IRadioConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V2-ndk_platform-source/gen/android/hardware/radio/config/IRadioConfigIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V2-ndk_platform-source/gen/android/hardware/radio/config/IRadioConfigResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V2-ndk_platform-source/gen/android/hardware/radio/config/MultipleEnabledProfilesMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V2-ndk_platform-source/gen/android/hardware/radio/config/PhoneCapability.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V2-ndk_platform-source/gen/android/hardware/radio/config/SimPortInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V2-ndk_platform-source/gen/android/hardware/radio/config/SimSlotStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V2-ndk_platform-source/gen/android/hardware/radio/config/SlotPortMapping.cpp
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
    add_library(android.hardware.radio.config-V2-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.config-V2-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-ndk_platform-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V2-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.config-V2-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.config-V2-ndk_platform
            android.hardware.radio-V2-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.radio.config-V2-ndk_platform ####################

#################### android.hardware.radio.voice-V1-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-ndk_platform-source/gen/android/hardware/radio/voice/AudioQuality.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-ndk_platform-source/gen/android/hardware/radio/voice/Call.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-ndk_platform-source/gen/android/hardware/radio/voice/CallForwardInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-ndk_platform-source/gen/android/hardware/radio/voice/CdmaCallWaiting.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-ndk_platform-source/gen/android/hardware/radio/voice/CdmaDisplayInfoRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-ndk_platform-source/gen/android/hardware/radio/voice/CdmaInformationRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-ndk_platform-source/gen/android/hardware/radio/voice/CdmaLineControlInfoRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-ndk_platform-source/gen/android/hardware/radio/voice/CdmaNumberInfoRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-ndk_platform-source/gen/android/hardware/radio/voice/CdmaOtaProvisionStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-ndk_platform-source/gen/android/hardware/radio/voice/CdmaRedirectingNumberInfoRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-ndk_platform-source/gen/android/hardware/radio/voice/CdmaSignalInfoRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-ndk_platform-source/gen/android/hardware/radio/voice/CdmaT53AudioControlInfoRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-ndk_platform-source/gen/android/hardware/radio/voice/CdmaT53ClirInfoRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-ndk_platform-source/gen/android/hardware/radio/voice/CfData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-ndk_platform-source/gen/android/hardware/radio/voice/ClipStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-ndk_platform-source/gen/android/hardware/radio/voice/Dial.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-ndk_platform-source/gen/android/hardware/radio/voice/EmergencyCallRouting.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-ndk_platform-source/gen/android/hardware/radio/voice/EmergencyNumber.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-ndk_platform-source/gen/android/hardware/radio/voice/EmergencyServiceCategory.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-ndk_platform-source/gen/android/hardware/radio/voice/IRadioVoice.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-ndk_platform-source/gen/android/hardware/radio/voice/IRadioVoiceIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-ndk_platform-source/gen/android/hardware/radio/voice/IRadioVoiceResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-ndk_platform-source/gen/android/hardware/radio/voice/LastCallFailCause.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-ndk_platform-source/gen/android/hardware/radio/voice/LastCallFailCauseInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-ndk_platform-source/gen/android/hardware/radio/voice/SrvccState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-ndk_platform-source/gen/android/hardware/radio/voice/SsInfoData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-ndk_platform-source/gen/android/hardware/radio/voice/StkCcUnsolSsResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-ndk_platform-source/gen/android/hardware/radio/voice/TtyMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-ndk_platform-source/gen/android/hardware/radio/voice/UssdModeType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-ndk_platform-source/gen/android/hardware/radio/voice/UusInfo.cpp
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
    add_library(android.hardware.radio.voice-V1-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.voice-V1-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V1-ndk_platform-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.voice-V1-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.voice-V1-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.voice-V1-ndk_platform
            android.hardware.radio-V1-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.radio.voice-V1-ndk_platform ####################

#################### android.hardware.radio.messaging-V1-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V1-cpp-source/gen/android/hardware/radio/messaging/CdmaBroadcastSmsConfigInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V1-cpp-source/gen/android/hardware/radio/messaging/CdmaSmsAck.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V1-cpp-source/gen/android/hardware/radio/messaging/CdmaSmsAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V1-cpp-source/gen/android/hardware/radio/messaging/CdmaSmsMessage.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V1-cpp-source/gen/android/hardware/radio/messaging/CdmaSmsSubaddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V1-cpp-source/gen/android/hardware/radio/messaging/CdmaSmsWriteArgs.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V1-cpp-source/gen/android/hardware/radio/messaging/GsmBroadcastSmsConfigInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V1-cpp-source/gen/android/hardware/radio/messaging/GsmSmsMessage.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V1-cpp-source/gen/android/hardware/radio/messaging/IRadioMessaging.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V1-cpp-source/gen/android/hardware/radio/messaging/IRadioMessagingIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V1-cpp-source/gen/android/hardware/radio/messaging/IRadioMessagingResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V1-cpp-source/gen/android/hardware/radio/messaging/ImsSmsMessage.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V1-cpp-source/gen/android/hardware/radio/messaging/SendSmsResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V1-cpp-source/gen/android/hardware/radio/messaging/SmsAcknowledgeFailCause.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V1-cpp-source/gen/android/hardware/radio/messaging/SmsWriteArgs.cpp
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
    add_library(android.hardware.radio.messaging-V1-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.messaging-V1-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V1-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V1-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.messaging-V1-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.messaging-V1-cpp
            android.hardware.radio-V1-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libm
            libutils
    )
endif ()
#################### android.hardware.radio.messaging-V1-cpp ####################

#################### android.hardware.radio.network-V2-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-cpp-source/gen/android/hardware/radio/network/AccessTechnologySpecificInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-cpp-source/gen/android/hardware/radio/network/BarringInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-cpp-source/gen/android/hardware/radio/network/BarringTypeSpecificInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-cpp-source/gen/android/hardware/radio/network/Cdma2000RegistrationInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-cpp-source/gen/android/hardware/radio/network/CdmaRoamingType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-cpp-source/gen/android/hardware/radio/network/CdmaSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-cpp-source/gen/android/hardware/radio/network/CellConnectionStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-cpp-source/gen/android/hardware/radio/network/CellIdentity.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-cpp-source/gen/android/hardware/radio/network/CellIdentityCdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-cpp-source/gen/android/hardware/radio/network/CellIdentityGsm.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-cpp-source/gen/android/hardware/radio/network/CellIdentityLte.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-cpp-source/gen/android/hardware/radio/network/CellIdentityNr.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-cpp-source/gen/android/hardware/radio/network/CellIdentityTdscdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-cpp-source/gen/android/hardware/radio/network/CellIdentityWcdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-cpp-source/gen/android/hardware/radio/network/CellInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-cpp-source/gen/android/hardware/radio/network/CellInfoCdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-cpp-source/gen/android/hardware/radio/network/CellInfoGsm.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-cpp-source/gen/android/hardware/radio/network/CellInfoLte.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-cpp-source/gen/android/hardware/radio/network/CellInfoNr.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-cpp-source/gen/android/hardware/radio/network/CellInfoRatSpecificInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-cpp-source/gen/android/hardware/radio/network/CellInfoTdscdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-cpp-source/gen/android/hardware/radio/network/CellInfoWcdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-cpp-source/gen/android/hardware/radio/network/ClosedSubscriberGroupInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-cpp-source/gen/android/hardware/radio/network/Domain.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-cpp-source/gen/android/hardware/radio/network/EmergencyMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-cpp-source/gen/android/hardware/radio/network/EmergencyNetworkScanTrigger.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-cpp-source/gen/android/hardware/radio/network/EmergencyRegResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-cpp-source/gen/android/hardware/radio/network/EmergencyScanType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-cpp-source/gen/android/hardware/radio/network/EutranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-cpp-source/gen/android/hardware/radio/network/EutranRegistrationInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-cpp-source/gen/android/hardware/radio/network/EvdoSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-cpp-source/gen/android/hardware/radio/network/GeranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-cpp-source/gen/android/hardware/radio/network/GsmSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-cpp-source/gen/android/hardware/radio/network/IRadioNetwork.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-cpp-source/gen/android/hardware/radio/network/IRadioNetworkIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-cpp-source/gen/android/hardware/radio/network/IRadioNetworkResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-cpp-source/gen/android/hardware/radio/network/IndicationFilter.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-cpp-source/gen/android/hardware/radio/network/LceDataInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-cpp-source/gen/android/hardware/radio/network/LinkCapacityEstimate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-cpp-source/gen/android/hardware/radio/network/LteSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-cpp-source/gen/android/hardware/radio/network/LteVopsInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-cpp-source/gen/android/hardware/radio/network/NetworkScanRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-cpp-source/gen/android/hardware/radio/network/NetworkScanResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-cpp-source/gen/android/hardware/radio/network/NgranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-cpp-source/gen/android/hardware/radio/network/NrDualConnectivityState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-cpp-source/gen/android/hardware/radio/network/NrIndicators.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-cpp-source/gen/android/hardware/radio/network/NrSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-cpp-source/gen/android/hardware/radio/network/NrVopsInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-cpp-source/gen/android/hardware/radio/network/OperatorInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-cpp-source/gen/android/hardware/radio/network/PhoneRestrictedState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-cpp-source/gen/android/hardware/radio/network/PhysicalChannelConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-cpp-source/gen/android/hardware/radio/network/PhysicalChannelConfigBand.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-cpp-source/gen/android/hardware/radio/network/RadioAccessSpecifier.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-cpp-source/gen/android/hardware/radio/network/RadioAccessSpecifierBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-cpp-source/gen/android/hardware/radio/network/RadioBandMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-cpp-source/gen/android/hardware/radio/network/RegState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-cpp-source/gen/android/hardware/radio/network/RegStateResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-cpp-source/gen/android/hardware/radio/network/RegistrationFailCause.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-cpp-source/gen/android/hardware/radio/network/SignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-cpp-source/gen/android/hardware/radio/network/SignalThresholdInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-cpp-source/gen/android/hardware/radio/network/SuppSvcNotification.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-cpp-source/gen/android/hardware/radio/network/TdscdmaSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-cpp-source/gen/android/hardware/radio/network/UsageSetting.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-cpp-source/gen/android/hardware/radio/network/UtranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-cpp-source/gen/android/hardware/radio/network/WcdmaSignalStrength.cpp
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
    add_library(android.hardware.radio.network-V2-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.network-V2-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.network-V2-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.network-V2-cpp
            android.hardware.radio-V2-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libm
            libutils
    )
endif ()
#################### android.hardware.radio.network-V2-cpp ####################

#################### android.hardware.radio.messaging-V3-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V3-ndk-source/gen/android/hardware/radio/messaging/CdmaBroadcastSmsConfigInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V3-ndk-source/gen/android/hardware/radio/messaging/CdmaSmsAck.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V3-ndk-source/gen/android/hardware/radio/messaging/CdmaSmsAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V3-ndk-source/gen/android/hardware/radio/messaging/CdmaSmsMessage.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V3-ndk-source/gen/android/hardware/radio/messaging/CdmaSmsSubaddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V3-ndk-source/gen/android/hardware/radio/messaging/CdmaSmsWriteArgs.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V3-ndk-source/gen/android/hardware/radio/messaging/GsmBroadcastSmsConfigInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V3-ndk-source/gen/android/hardware/radio/messaging/GsmSmsMessage.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V3-ndk-source/gen/android/hardware/radio/messaging/IRadioMessaging.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V3-ndk-source/gen/android/hardware/radio/messaging/IRadioMessagingIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V3-ndk-source/gen/android/hardware/radio/messaging/IRadioMessagingResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V3-ndk-source/gen/android/hardware/radio/messaging/ImsSmsMessage.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V3-ndk-source/gen/android/hardware/radio/messaging/SendSmsResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V3-ndk-source/gen/android/hardware/radio/messaging/SmsAcknowledgeFailCause.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V3-ndk-source/gen/android/hardware/radio/messaging/SmsWriteArgs.cpp
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
    add_library(android.hardware.radio.messaging-V3-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.messaging-V3-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-ndk-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V3-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.messaging-V3-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.messaging-V3-ndk
            android.hardware.radio-V2-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.radio.messaging-V3-ndk ####################

#################### android.hardware.radio.config-V1-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V1-cpp-source/gen/android/hardware/radio/config/IRadioConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V1-cpp-source/gen/android/hardware/radio/config/IRadioConfigIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V1-cpp-source/gen/android/hardware/radio/config/IRadioConfigResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V1-cpp-source/gen/android/hardware/radio/config/PhoneCapability.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V1-cpp-source/gen/android/hardware/radio/config/SimPortInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V1-cpp-source/gen/android/hardware/radio/config/SimSlotStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V1-cpp-source/gen/android/hardware/radio/config/SlotPortMapping.cpp
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
    add_library(android.hardware.radio.config-V1-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.config-V1-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V1-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V1-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.config-V1-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.config-V1-cpp
            android.hardware.radio-V1-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libm
            libutils
    )
endif ()
#################### android.hardware.radio.config-V1-cpp ####################

#################### android.hardware.radio.sim-V3-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-ndk_platform-source/gen/android/hardware/radio/sim/AppStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-ndk_platform-source/gen/android/hardware/radio/sim/CardPowerState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-ndk_platform-source/gen/android/hardware/radio/sim/CardStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-ndk_platform-source/gen/android/hardware/radio/sim/Carrier.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-ndk_platform-source/gen/android/hardware/radio/sim/CarrierRestrictions.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-ndk_platform-source/gen/android/hardware/radio/sim/CdmaSubscriptionSource.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-ndk_platform-source/gen/android/hardware/radio/sim/IRadioSim.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-ndk_platform-source/gen/android/hardware/radio/sim/IRadioSimIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-ndk_platform-source/gen/android/hardware/radio/sim/IRadioSimResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-ndk_platform-source/gen/android/hardware/radio/sim/IccIo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-ndk_platform-source/gen/android/hardware/radio/sim/IccIoResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-ndk_platform-source/gen/android/hardware/radio/sim/ImsiEncryptionInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-ndk_platform-source/gen/android/hardware/radio/sim/PbReceivedStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-ndk_platform-source/gen/android/hardware/radio/sim/PersoSubstate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-ndk_platform-source/gen/android/hardware/radio/sim/PhonebookCapacity.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-ndk_platform-source/gen/android/hardware/radio/sim/PhonebookRecordInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-ndk_platform-source/gen/android/hardware/radio/sim/PinState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-ndk_platform-source/gen/android/hardware/radio/sim/SelectUiccSub.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-ndk_platform-source/gen/android/hardware/radio/sim/SessionInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-ndk_platform-source/gen/android/hardware/radio/sim/SimApdu.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-ndk_platform-source/gen/android/hardware/radio/sim/SimLockMultiSimPolicy.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-ndk_platform-source/gen/android/hardware/radio/sim/SimRefreshResult.cpp
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
    add_library(android.hardware.radio.sim-V3-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.sim-V3-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-ndk_platform-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V2-ndk_platform-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.sim-V3-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.sim-V3-ndk_platform
            android.hardware.radio-V2-ndk_platform
            android.hardware.radio.config-V2-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.radio.sim-V3-ndk_platform ####################

#################### android.hardware.radio.config-V3-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V3-ndk-source/gen/android/hardware/radio/config/IRadioConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V3-ndk-source/gen/android/hardware/radio/config/IRadioConfigIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V3-ndk-source/gen/android/hardware/radio/config/IRadioConfigResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V3-ndk-source/gen/android/hardware/radio/config/MultipleEnabledProfilesMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V3-ndk-source/gen/android/hardware/radio/config/PhoneCapability.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V3-ndk-source/gen/android/hardware/radio/config/SimPortInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V3-ndk-source/gen/android/hardware/radio/config/SimSlotStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V3-ndk-source/gen/android/hardware/radio/config/SlotPortMapping.cpp
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
    add_library(android.hardware.radio.config-V3-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.config-V3-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-ndk-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V3-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.config-V3-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.config-V3-ndk
            android.hardware.radio-V2-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.radio.config-V3-ndk ####################

#################### android.hardware.radio.modem-V2-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V2-ndk_platform-source/gen/android/hardware/radio/modem/ActivityStatsInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V2-ndk_platform-source/gen/android/hardware/radio/modem/ActivityStatsTechSpecificInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V2-ndk_platform-source/gen/android/hardware/radio/modem/DeviceStateType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V2-ndk_platform-source/gen/android/hardware/radio/modem/HardwareConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V2-ndk_platform-source/gen/android/hardware/radio/modem/HardwareConfigModem.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V2-ndk_platform-source/gen/android/hardware/radio/modem/HardwareConfigSim.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V2-ndk_platform-source/gen/android/hardware/radio/modem/IRadioModem.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V2-ndk_platform-source/gen/android/hardware/radio/modem/IRadioModemIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V2-ndk_platform-source/gen/android/hardware/radio/modem/IRadioModemResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V2-ndk_platform-source/gen/android/hardware/radio/modem/ImeiInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V2-ndk_platform-source/gen/android/hardware/radio/modem/NvItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V2-ndk_platform-source/gen/android/hardware/radio/modem/NvWriteItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V2-ndk_platform-source/gen/android/hardware/radio/modem/RadioCapability.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V2-ndk_platform-source/gen/android/hardware/radio/modem/RadioState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V2-ndk_platform-source/gen/android/hardware/radio/modem/ResetNvType.cpp
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
    add_library(android.hardware.radio.modem-V2-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.modem-V2-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-ndk_platform-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.modem-V2-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.modem-V2-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.modem-V2-ndk_platform
            android.hardware.radio-V2-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.radio.modem-V2-ndk_platform ####################

#################### android.hardware.radio.config-V3-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V3-ndk_platform-source/gen/android/hardware/radio/config/IRadioConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V3-ndk_platform-source/gen/android/hardware/radio/config/IRadioConfigIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V3-ndk_platform-source/gen/android/hardware/radio/config/IRadioConfigResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V3-ndk_platform-source/gen/android/hardware/radio/config/MultipleEnabledProfilesMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V3-ndk_platform-source/gen/android/hardware/radio/config/PhoneCapability.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V3-ndk_platform-source/gen/android/hardware/radio/config/SimPortInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V3-ndk_platform-source/gen/android/hardware/radio/config/SimSlotStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V3-ndk_platform-source/gen/android/hardware/radio/config/SlotPortMapping.cpp
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
    add_library(android.hardware.radio.config-V3-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.config-V3-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-ndk_platform-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V3-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.config-V3-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.config-V3-ndk_platform
            android.hardware.radio-V2-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.radio.config-V3-ndk_platform ####################

#################### android.hardware.radio-cpp-analyzer ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-cpp-analyzer-source/gen/android/hardware/radio/AccessNetwork.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-cpp-analyzer-source/gen/android/hardware/radio/RadioAccessFamily.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-cpp-analyzer-source/gen/android/hardware/radio/RadioConst.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-cpp-analyzer-source/gen/android/hardware/radio/RadioError.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-cpp-analyzer-source/gen/android/hardware/radio/RadioIndicationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-cpp-analyzer-source/gen/android/hardware/radio/RadioResponseInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-cpp-analyzer-source/gen/android/hardware/radio/RadioResponseInfoModem.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-cpp-analyzer-source/gen/android/hardware/radio/RadioResponseType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-cpp-analyzer-source/gen/android/hardware/radio/RadioTechnology.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-cpp-analyzer-source/gen/android/hardware/radio/RadioTechnologyFamily.cpp
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
    add_executable(android.hardware.radio-cpp-analyzer ${SOURCE_FILES})

    target_include_directories(android.hardware.radio-cpp-analyzer
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/radio/aidl
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V3-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-cpp-analyzer-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio-cpp-analyzer PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio-cpp-analyzer
            aidl-analyzer-main
            android.hardware.radio-V3-cpp
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
            libutils
    )
endif ()
#################### android.hardware.radio-cpp-analyzer ####################

#################### android.hardware.radio.ims-V1-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims-V1-ndk_platform-source/gen/android/hardware/radio/ims/ConnectionFailureInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims-V1-ndk_platform-source/gen/android/hardware/radio/ims/EpsFallbackReason.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims-V1-ndk_platform-source/gen/android/hardware/radio/ims/IRadioIms.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims-V1-ndk_platform-source/gen/android/hardware/radio/ims/IRadioImsIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims-V1-ndk_platform-source/gen/android/hardware/radio/ims/IRadioImsResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims-V1-ndk_platform-source/gen/android/hardware/radio/ims/ImsCall.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims-V1-ndk_platform-source/gen/android/hardware/radio/ims/ImsDeregistrationReason.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims-V1-ndk_platform-source/gen/android/hardware/radio/ims/ImsRegistration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims-V1-ndk_platform-source/gen/android/hardware/radio/ims/ImsRegistrationState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims-V1-ndk_platform-source/gen/android/hardware/radio/ims/ImsStreamDirection.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims-V1-ndk_platform-source/gen/android/hardware/radio/ims/ImsStreamType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims-V1-ndk_platform-source/gen/android/hardware/radio/ims/ImsTrafficType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims-V1-ndk_platform-source/gen/android/hardware/radio/ims/SrvccCall.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims-V1-ndk_platform-source/gen/android/hardware/radio/ims/SuggestedAction.cpp
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
    add_library(android.hardware.radio.ims-V1-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.ims-V1-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-ndk_platform-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims-V1-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.ims-V1-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.ims-V1-ndk_platform
            android.hardware.radio-V2-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.radio.ims-V1-ndk_platform ####################

#################### android.hardware.radio.sim-V1-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V1-ndk-source/gen/android/hardware/radio/sim/AppStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V1-ndk-source/gen/android/hardware/radio/sim/CardPowerState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V1-ndk-source/gen/android/hardware/radio/sim/CardStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V1-ndk-source/gen/android/hardware/radio/sim/Carrier.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V1-ndk-source/gen/android/hardware/radio/sim/CarrierRestrictions.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V1-ndk-source/gen/android/hardware/radio/sim/CdmaSubscriptionSource.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V1-ndk-source/gen/android/hardware/radio/sim/IRadioSim.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V1-ndk-source/gen/android/hardware/radio/sim/IRadioSimIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V1-ndk-source/gen/android/hardware/radio/sim/IRadioSimResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V1-ndk-source/gen/android/hardware/radio/sim/IccIo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V1-ndk-source/gen/android/hardware/radio/sim/IccIoResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V1-ndk-source/gen/android/hardware/radio/sim/ImsiEncryptionInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V1-ndk-source/gen/android/hardware/radio/sim/PbReceivedStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V1-ndk-source/gen/android/hardware/radio/sim/PersoSubstate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V1-ndk-source/gen/android/hardware/radio/sim/PhonebookCapacity.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V1-ndk-source/gen/android/hardware/radio/sim/PhonebookRecordInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V1-ndk-source/gen/android/hardware/radio/sim/PinState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V1-ndk-source/gen/android/hardware/radio/sim/SelectUiccSub.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V1-ndk-source/gen/android/hardware/radio/sim/SimApdu.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V1-ndk-source/gen/android/hardware/radio/sim/SimLockMultiSimPolicy.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V1-ndk-source/gen/android/hardware/radio/sim/SimRefreshResult.cpp
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
    add_library(android.hardware.radio.sim-V1-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.sim-V1-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V1-ndk-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V1-ndk-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V1-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.sim-V1-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.sim-V1-ndk
            android.hardware.radio-V1-ndk
            android.hardware.radio.config-V1-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.radio.sim-V1-ndk ####################

#################### android.hardware.radio.network-V3-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk-source/gen/android/hardware/radio/network/AccessTechnologySpecificInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk-source/gen/android/hardware/radio/network/BarringInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk-source/gen/android/hardware/radio/network/BarringTypeSpecificInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk-source/gen/android/hardware/radio/network/Cdma2000RegistrationInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk-source/gen/android/hardware/radio/network/CdmaRoamingType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk-source/gen/android/hardware/radio/network/CdmaSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk-source/gen/android/hardware/radio/network/CellConnectionStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk-source/gen/android/hardware/radio/network/CellIdentity.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk-source/gen/android/hardware/radio/network/CellIdentityCdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk-source/gen/android/hardware/radio/network/CellIdentityGsm.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk-source/gen/android/hardware/radio/network/CellIdentityLte.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk-source/gen/android/hardware/radio/network/CellIdentityNr.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk-source/gen/android/hardware/radio/network/CellIdentityTdscdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk-source/gen/android/hardware/radio/network/CellIdentityWcdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk-source/gen/android/hardware/radio/network/CellInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk-source/gen/android/hardware/radio/network/CellInfoCdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk-source/gen/android/hardware/radio/network/CellInfoGsm.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk-source/gen/android/hardware/radio/network/CellInfoLte.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk-source/gen/android/hardware/radio/network/CellInfoNr.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk-source/gen/android/hardware/radio/network/CellInfoRatSpecificInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk-source/gen/android/hardware/radio/network/CellInfoTdscdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk-source/gen/android/hardware/radio/network/CellInfoWcdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk-source/gen/android/hardware/radio/network/ClosedSubscriberGroupInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk-source/gen/android/hardware/radio/network/Domain.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk-source/gen/android/hardware/radio/network/EmergencyMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk-source/gen/android/hardware/radio/network/EmergencyNetworkScanTrigger.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk-source/gen/android/hardware/radio/network/EmergencyRegResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk-source/gen/android/hardware/radio/network/EmergencyScanType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk-source/gen/android/hardware/radio/network/EutranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk-source/gen/android/hardware/radio/network/EutranRegistrationInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk-source/gen/android/hardware/radio/network/EvdoSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk-source/gen/android/hardware/radio/network/GeranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk-source/gen/android/hardware/radio/network/GsmSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk-source/gen/android/hardware/radio/network/IRadioNetwork.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk-source/gen/android/hardware/radio/network/IRadioNetworkIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk-source/gen/android/hardware/radio/network/IRadioNetworkResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk-source/gen/android/hardware/radio/network/IndicationFilter.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk-source/gen/android/hardware/radio/network/LceDataInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk-source/gen/android/hardware/radio/network/LinkCapacityEstimate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk-source/gen/android/hardware/radio/network/LteSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk-source/gen/android/hardware/radio/network/LteVopsInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk-source/gen/android/hardware/radio/network/NetworkScanRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk-source/gen/android/hardware/radio/network/NetworkScanResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk-source/gen/android/hardware/radio/network/NgranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk-source/gen/android/hardware/radio/network/NrDualConnectivityState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk-source/gen/android/hardware/radio/network/NrIndicators.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk-source/gen/android/hardware/radio/network/NrSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk-source/gen/android/hardware/radio/network/NrVopsInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk-source/gen/android/hardware/radio/network/OperatorInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk-source/gen/android/hardware/radio/network/PhoneRestrictedState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk-source/gen/android/hardware/radio/network/PhysicalChannelConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk-source/gen/android/hardware/radio/network/PhysicalChannelConfigBand.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk-source/gen/android/hardware/radio/network/RadioAccessSpecifier.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk-source/gen/android/hardware/radio/network/RadioAccessSpecifierBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk-source/gen/android/hardware/radio/network/RadioBandMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk-source/gen/android/hardware/radio/network/RegState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk-source/gen/android/hardware/radio/network/RegStateResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk-source/gen/android/hardware/radio/network/RegistrationFailCause.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk-source/gen/android/hardware/radio/network/SignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk-source/gen/android/hardware/radio/network/SignalThresholdInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk-source/gen/android/hardware/radio/network/SuppSvcNotification.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk-source/gen/android/hardware/radio/network/TdscdmaSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk-source/gen/android/hardware/radio/network/UsageSetting.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk-source/gen/android/hardware/radio/network/UtranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk-source/gen/android/hardware/radio/network/WcdmaSignalStrength.cpp
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
    add_library(android.hardware.radio.network-V3-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.network-V3-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-ndk-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V3-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.network-V3-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.network-V3-ndk
            android.hardware.radio-V2-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.radio.network-V3-ndk ####################

#################### android.hardware.radio.messaging-V2-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V2-cpp-source/gen/android/hardware/radio/messaging/CdmaBroadcastSmsConfigInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V2-cpp-source/gen/android/hardware/radio/messaging/CdmaSmsAck.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V2-cpp-source/gen/android/hardware/radio/messaging/CdmaSmsAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V2-cpp-source/gen/android/hardware/radio/messaging/CdmaSmsMessage.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V2-cpp-source/gen/android/hardware/radio/messaging/CdmaSmsSubaddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V2-cpp-source/gen/android/hardware/radio/messaging/CdmaSmsWriteArgs.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V2-cpp-source/gen/android/hardware/radio/messaging/GsmBroadcastSmsConfigInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V2-cpp-source/gen/android/hardware/radio/messaging/GsmSmsMessage.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V2-cpp-source/gen/android/hardware/radio/messaging/IRadioMessaging.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V2-cpp-source/gen/android/hardware/radio/messaging/IRadioMessagingIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V2-cpp-source/gen/android/hardware/radio/messaging/IRadioMessagingResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V2-cpp-source/gen/android/hardware/radio/messaging/ImsSmsMessage.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V2-cpp-source/gen/android/hardware/radio/messaging/SendSmsResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V2-cpp-source/gen/android/hardware/radio/messaging/SmsAcknowledgeFailCause.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V2-cpp-source/gen/android/hardware/radio/messaging/SmsWriteArgs.cpp
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
    add_library(android.hardware.radio.messaging-V2-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.messaging-V2-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V2-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.messaging-V2-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.messaging-V2-cpp
            android.hardware.radio-V2-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libm
            libutils
    )
endif ()
#################### android.hardware.radio.messaging-V2-cpp ####################

#################### android.hardware.radio.network-V1-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-cpp-source/gen/android/hardware/radio/network/AccessTechnologySpecificInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-cpp-source/gen/android/hardware/radio/network/BarringInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-cpp-source/gen/android/hardware/radio/network/BarringTypeSpecificInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-cpp-source/gen/android/hardware/radio/network/Cdma2000RegistrationInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-cpp-source/gen/android/hardware/radio/network/CdmaRoamingType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-cpp-source/gen/android/hardware/radio/network/CdmaSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-cpp-source/gen/android/hardware/radio/network/CellConnectionStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-cpp-source/gen/android/hardware/radio/network/CellIdentity.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-cpp-source/gen/android/hardware/radio/network/CellIdentityCdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-cpp-source/gen/android/hardware/radio/network/CellIdentityGsm.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-cpp-source/gen/android/hardware/radio/network/CellIdentityLte.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-cpp-source/gen/android/hardware/radio/network/CellIdentityNr.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-cpp-source/gen/android/hardware/radio/network/CellIdentityTdscdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-cpp-source/gen/android/hardware/radio/network/CellIdentityWcdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-cpp-source/gen/android/hardware/radio/network/CellInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-cpp-source/gen/android/hardware/radio/network/CellInfoCdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-cpp-source/gen/android/hardware/radio/network/CellInfoGsm.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-cpp-source/gen/android/hardware/radio/network/CellInfoLte.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-cpp-source/gen/android/hardware/radio/network/CellInfoNr.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-cpp-source/gen/android/hardware/radio/network/CellInfoRatSpecificInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-cpp-source/gen/android/hardware/radio/network/CellInfoTdscdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-cpp-source/gen/android/hardware/radio/network/CellInfoWcdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-cpp-source/gen/android/hardware/radio/network/ClosedSubscriberGroupInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-cpp-source/gen/android/hardware/radio/network/Domain.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-cpp-source/gen/android/hardware/radio/network/EutranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-cpp-source/gen/android/hardware/radio/network/EutranRegistrationInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-cpp-source/gen/android/hardware/radio/network/EvdoSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-cpp-source/gen/android/hardware/radio/network/GeranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-cpp-source/gen/android/hardware/radio/network/GsmSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-cpp-source/gen/android/hardware/radio/network/IRadioNetwork.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-cpp-source/gen/android/hardware/radio/network/IRadioNetworkIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-cpp-source/gen/android/hardware/radio/network/IRadioNetworkResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-cpp-source/gen/android/hardware/radio/network/IndicationFilter.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-cpp-source/gen/android/hardware/radio/network/LceDataInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-cpp-source/gen/android/hardware/radio/network/LinkCapacityEstimate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-cpp-source/gen/android/hardware/radio/network/LteSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-cpp-source/gen/android/hardware/radio/network/LteVopsInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-cpp-source/gen/android/hardware/radio/network/NetworkScanRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-cpp-source/gen/android/hardware/radio/network/NetworkScanResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-cpp-source/gen/android/hardware/radio/network/NgranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-cpp-source/gen/android/hardware/radio/network/NrDualConnectivityState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-cpp-source/gen/android/hardware/radio/network/NrIndicators.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-cpp-source/gen/android/hardware/radio/network/NrSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-cpp-source/gen/android/hardware/radio/network/NrVopsInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-cpp-source/gen/android/hardware/radio/network/OperatorInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-cpp-source/gen/android/hardware/radio/network/PhoneRestrictedState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-cpp-source/gen/android/hardware/radio/network/PhysicalChannelConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-cpp-source/gen/android/hardware/radio/network/PhysicalChannelConfigBand.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-cpp-source/gen/android/hardware/radio/network/RadioAccessSpecifier.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-cpp-source/gen/android/hardware/radio/network/RadioAccessSpecifierBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-cpp-source/gen/android/hardware/radio/network/RadioBandMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-cpp-source/gen/android/hardware/radio/network/RegState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-cpp-source/gen/android/hardware/radio/network/RegStateResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-cpp-source/gen/android/hardware/radio/network/RegistrationFailCause.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-cpp-source/gen/android/hardware/radio/network/SignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-cpp-source/gen/android/hardware/radio/network/SignalThresholdInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-cpp-source/gen/android/hardware/radio/network/SuppSvcNotification.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-cpp-source/gen/android/hardware/radio/network/TdscdmaSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-cpp-source/gen/android/hardware/radio/network/UsageSetting.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-cpp-source/gen/android/hardware/radio/network/UtranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-cpp-source/gen/android/hardware/radio/network/WcdmaSignalStrength.cpp
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
    add_library(android.hardware.radio.network-V1-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.network-V1-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V1-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V1-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.network-V1-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.network-V1-cpp
            android.hardware.radio-V1-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libm
            libutils
    )
endif ()
#################### android.hardware.radio.network-V1-cpp ####################

#################### android.hardware.radio.sim-V3-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-cpp-source/gen/android/hardware/radio/sim/AppStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-cpp-source/gen/android/hardware/radio/sim/CardPowerState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-cpp-source/gen/android/hardware/radio/sim/CardStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-cpp-source/gen/android/hardware/radio/sim/Carrier.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-cpp-source/gen/android/hardware/radio/sim/CarrierRestrictions.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-cpp-source/gen/android/hardware/radio/sim/CdmaSubscriptionSource.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-cpp-source/gen/android/hardware/radio/sim/IRadioSim.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-cpp-source/gen/android/hardware/radio/sim/IRadioSimIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-cpp-source/gen/android/hardware/radio/sim/IRadioSimResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-cpp-source/gen/android/hardware/radio/sim/IccIo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-cpp-source/gen/android/hardware/radio/sim/IccIoResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-cpp-source/gen/android/hardware/radio/sim/ImsiEncryptionInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-cpp-source/gen/android/hardware/radio/sim/PbReceivedStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-cpp-source/gen/android/hardware/radio/sim/PersoSubstate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-cpp-source/gen/android/hardware/radio/sim/PhonebookCapacity.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-cpp-source/gen/android/hardware/radio/sim/PhonebookRecordInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-cpp-source/gen/android/hardware/radio/sim/PinState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-cpp-source/gen/android/hardware/radio/sim/SelectUiccSub.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-cpp-source/gen/android/hardware/radio/sim/SessionInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-cpp-source/gen/android/hardware/radio/sim/SimApdu.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-cpp-source/gen/android/hardware/radio/sim/SimLockMultiSimPolicy.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-cpp-source/gen/android/hardware/radio/sim/SimRefreshResult.cpp
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
    add_library(android.hardware.radio.sim-V3-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.sim-V3-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V2-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V3-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.sim-V3-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.sim-V3-cpp
            android.hardware.radio-V2-cpp
            android.hardware.radio.config-V2-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libm
            libutils
    )
endif ()
#################### android.hardware.radio.sim-V3-cpp ####################

#################### android.hardware.radio.network-V2-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk_platform-source/gen/android/hardware/radio/network/AccessTechnologySpecificInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk_platform-source/gen/android/hardware/radio/network/BarringInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk_platform-source/gen/android/hardware/radio/network/BarringTypeSpecificInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk_platform-source/gen/android/hardware/radio/network/Cdma2000RegistrationInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk_platform-source/gen/android/hardware/radio/network/CdmaRoamingType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk_platform-source/gen/android/hardware/radio/network/CdmaSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk_platform-source/gen/android/hardware/radio/network/CellConnectionStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk_platform-source/gen/android/hardware/radio/network/CellIdentity.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk_platform-source/gen/android/hardware/radio/network/CellIdentityCdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk_platform-source/gen/android/hardware/radio/network/CellIdentityGsm.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk_platform-source/gen/android/hardware/radio/network/CellIdentityLte.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk_platform-source/gen/android/hardware/radio/network/CellIdentityNr.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk_platform-source/gen/android/hardware/radio/network/CellIdentityTdscdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk_platform-source/gen/android/hardware/radio/network/CellIdentityWcdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk_platform-source/gen/android/hardware/radio/network/CellInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk_platform-source/gen/android/hardware/radio/network/CellInfoCdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk_platform-source/gen/android/hardware/radio/network/CellInfoGsm.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk_platform-source/gen/android/hardware/radio/network/CellInfoLte.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk_platform-source/gen/android/hardware/radio/network/CellInfoNr.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk_platform-source/gen/android/hardware/radio/network/CellInfoRatSpecificInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk_platform-source/gen/android/hardware/radio/network/CellInfoTdscdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk_platform-source/gen/android/hardware/radio/network/CellInfoWcdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk_platform-source/gen/android/hardware/radio/network/ClosedSubscriberGroupInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk_platform-source/gen/android/hardware/radio/network/Domain.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk_platform-source/gen/android/hardware/radio/network/EmergencyMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk_platform-source/gen/android/hardware/radio/network/EmergencyNetworkScanTrigger.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk_platform-source/gen/android/hardware/radio/network/EmergencyRegResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk_platform-source/gen/android/hardware/radio/network/EmergencyScanType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk_platform-source/gen/android/hardware/radio/network/EutranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk_platform-source/gen/android/hardware/radio/network/EutranRegistrationInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk_platform-source/gen/android/hardware/radio/network/EvdoSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk_platform-source/gen/android/hardware/radio/network/GeranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk_platform-source/gen/android/hardware/radio/network/GsmSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk_platform-source/gen/android/hardware/radio/network/IRadioNetwork.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk_platform-source/gen/android/hardware/radio/network/IRadioNetworkIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk_platform-source/gen/android/hardware/radio/network/IRadioNetworkResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk_platform-source/gen/android/hardware/radio/network/IndicationFilter.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk_platform-source/gen/android/hardware/radio/network/LceDataInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk_platform-source/gen/android/hardware/radio/network/LinkCapacityEstimate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk_platform-source/gen/android/hardware/radio/network/LteSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk_platform-source/gen/android/hardware/radio/network/LteVopsInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk_platform-source/gen/android/hardware/radio/network/NetworkScanRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk_platform-source/gen/android/hardware/radio/network/NetworkScanResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk_platform-source/gen/android/hardware/radio/network/NgranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk_platform-source/gen/android/hardware/radio/network/NrDualConnectivityState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk_platform-source/gen/android/hardware/radio/network/NrIndicators.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk_platform-source/gen/android/hardware/radio/network/NrSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk_platform-source/gen/android/hardware/radio/network/NrVopsInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk_platform-source/gen/android/hardware/radio/network/OperatorInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk_platform-source/gen/android/hardware/radio/network/PhoneRestrictedState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk_platform-source/gen/android/hardware/radio/network/PhysicalChannelConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk_platform-source/gen/android/hardware/radio/network/PhysicalChannelConfigBand.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk_platform-source/gen/android/hardware/radio/network/RadioAccessSpecifier.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk_platform-source/gen/android/hardware/radio/network/RadioAccessSpecifierBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk_platform-source/gen/android/hardware/radio/network/RadioBandMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk_platform-source/gen/android/hardware/radio/network/RegState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk_platform-source/gen/android/hardware/radio/network/RegStateResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk_platform-source/gen/android/hardware/radio/network/RegistrationFailCause.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk_platform-source/gen/android/hardware/radio/network/SignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk_platform-source/gen/android/hardware/radio/network/SignalThresholdInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk_platform-source/gen/android/hardware/radio/network/SuppSvcNotification.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk_platform-source/gen/android/hardware/radio/network/TdscdmaSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk_platform-source/gen/android/hardware/radio/network/UsageSetting.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk_platform-source/gen/android/hardware/radio/network/UtranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk_platform-source/gen/android/hardware/radio/network/WcdmaSignalStrength.cpp
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
    add_library(android.hardware.radio.network-V2-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.network-V2-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-ndk_platform-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.network-V2-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.network-V2-ndk_platform
            android.hardware.radio-V2-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.radio.network-V2-ndk_platform ####################

#################### android.hardware.radio.ims.media-V1-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V1-ndk-source/gen/android/hardware/radio/ims/media/AmrMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V1-ndk-source/gen/android/hardware/radio/ims/media/AmrParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V1-ndk-source/gen/android/hardware/radio/ims/media/AnbrMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V1-ndk-source/gen/android/hardware/radio/ims/media/CallQuality.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V1-ndk-source/gen/android/hardware/radio/ims/media/CodecMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V1-ndk-source/gen/android/hardware/radio/ims/media/CodecParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V1-ndk-source/gen/android/hardware/radio/ims/media/CodecSpecificParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V1-ndk-source/gen/android/hardware/radio/ims/media/CodecType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V1-ndk-source/gen/android/hardware/radio/ims/media/DtmfParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V1-ndk-source/gen/android/hardware/radio/ims/media/EvsBandwidth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V1-ndk-source/gen/android/hardware/radio/ims/media/EvsMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V1-ndk-source/gen/android/hardware/radio/ims/media/EvsParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V1-ndk-source/gen/android/hardware/radio/ims/media/IImsMedia.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V1-ndk-source/gen/android/hardware/radio/ims/media/IImsMediaListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V1-ndk-source/gen/android/hardware/radio/ims/media/IImsMediaSession.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V1-ndk-source/gen/android/hardware/radio/ims/media/IImsMediaSessionListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V1-ndk-source/gen/android/hardware/radio/ims/media/LocalEndPoint.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V1-ndk-source/gen/android/hardware/radio/ims/media/MediaDirection.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V1-ndk-source/gen/android/hardware/radio/ims/media/MediaQualityStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V1-ndk-source/gen/android/hardware/radio/ims/media/MediaQualityThreshold.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V1-ndk-source/gen/android/hardware/radio/ims/media/RtcpConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V1-ndk-source/gen/android/hardware/radio/ims/media/RtcpXrReportBlockType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V1-ndk-source/gen/android/hardware/radio/ims/media/RtpAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V1-ndk-source/gen/android/hardware/radio/ims/media/RtpConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V1-ndk-source/gen/android/hardware/radio/ims/media/RtpError.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V1-ndk-source/gen/android/hardware/radio/ims/media/RtpHeaderExtension.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V1-ndk-source/gen/android/hardware/radio/ims/media/RtpSessionParams.cpp
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
    add_library(android.hardware.radio.ims.media-V1-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.ims.media-V1-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-ndk-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V1-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.ims.media-V1-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.ims.media-V1-ndk
            android.hardware.radio-V2-ndk
            android.hardware.radio.data-V2-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.radio.ims.media-V1-ndk ####################

#################### android.hardware.radio.data-V1-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-ndk_platform-source/gen/android/hardware/radio/data/ApnAuthType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-ndk_platform-source/gen/android/hardware/radio/data/ApnTypes.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-ndk_platform-source/gen/android/hardware/radio/data/DataCallFailCause.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-ndk_platform-source/gen/android/hardware/radio/data/DataProfileInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-ndk_platform-source/gen/android/hardware/radio/data/DataRequestReason.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-ndk_platform-source/gen/android/hardware/radio/data/DataThrottlingAction.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-ndk_platform-source/gen/android/hardware/radio/data/EpsQos.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-ndk_platform-source/gen/android/hardware/radio/data/IRadioData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-ndk_platform-source/gen/android/hardware/radio/data/IRadioDataIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-ndk_platform-source/gen/android/hardware/radio/data/IRadioDataResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-ndk_platform-source/gen/android/hardware/radio/data/KeepaliveRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-ndk_platform-source/gen/android/hardware/radio/data/KeepaliveStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-ndk_platform-source/gen/android/hardware/radio/data/LinkAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-ndk_platform-source/gen/android/hardware/radio/data/NrQos.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-ndk_platform-source/gen/android/hardware/radio/data/OsAppId.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-ndk_platform-source/gen/android/hardware/radio/data/PcoDataInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-ndk_platform-source/gen/android/hardware/radio/data/PdpProtocolType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-ndk_platform-source/gen/android/hardware/radio/data/PortRange.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-ndk_platform-source/gen/android/hardware/radio/data/Qos.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-ndk_platform-source/gen/android/hardware/radio/data/QosBandwidth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-ndk_platform-source/gen/android/hardware/radio/data/QosFilter.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-ndk_platform-source/gen/android/hardware/radio/data/QosFilterIpsecSpi.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-ndk_platform-source/gen/android/hardware/radio/data/QosFilterIpv6FlowLabel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-ndk_platform-source/gen/android/hardware/radio/data/QosFilterTypeOfService.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-ndk_platform-source/gen/android/hardware/radio/data/QosSession.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-ndk_platform-source/gen/android/hardware/radio/data/RouteSelectionDescriptor.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-ndk_platform-source/gen/android/hardware/radio/data/SetupDataCallResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-ndk_platform-source/gen/android/hardware/radio/data/SliceInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-ndk_platform-source/gen/android/hardware/radio/data/SlicingConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-ndk_platform-source/gen/android/hardware/radio/data/TrafficDescriptor.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-ndk_platform-source/gen/android/hardware/radio/data/UrspRule.cpp
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
    add_library(android.hardware.radio.data-V1-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.data-V1-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V1-ndk_platform-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V1-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.data-V1-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.data-V1-ndk_platform
            android.hardware.radio-V1-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.radio.data-V1-ndk_platform ####################

#################### android.hardware.radio.config-V2-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V2-ndk-source/gen/android/hardware/radio/config/IRadioConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V2-ndk-source/gen/android/hardware/radio/config/IRadioConfigIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V2-ndk-source/gen/android/hardware/radio/config/IRadioConfigResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V2-ndk-source/gen/android/hardware/radio/config/MultipleEnabledProfilesMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V2-ndk-source/gen/android/hardware/radio/config/PhoneCapability.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V2-ndk-source/gen/android/hardware/radio/config/SimPortInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V2-ndk-source/gen/android/hardware/radio/config/SimSlotStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V2-ndk-source/gen/android/hardware/radio/config/SlotPortMapping.cpp
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
    add_library(android.hardware.radio.config-V2-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.config-V2-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-ndk-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V2-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.config-V2-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.config-V2-ndk
            android.hardware.radio-V2-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.radio.config-V2-ndk ####################

#################### android.hardware.radio.ims.media-V2-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V2-ndk_platform-source/gen/android/hardware/radio/ims/media/AmrMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V2-ndk_platform-source/gen/android/hardware/radio/ims/media/AmrParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V2-ndk_platform-source/gen/android/hardware/radio/ims/media/AnbrMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V2-ndk_platform-source/gen/android/hardware/radio/ims/media/CallQuality.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V2-ndk_platform-source/gen/android/hardware/radio/ims/media/CodecMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V2-ndk_platform-source/gen/android/hardware/radio/ims/media/CodecParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V2-ndk_platform-source/gen/android/hardware/radio/ims/media/CodecSpecificParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V2-ndk_platform-source/gen/android/hardware/radio/ims/media/CodecType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V2-ndk_platform-source/gen/android/hardware/radio/ims/media/DtmfParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V2-ndk_platform-source/gen/android/hardware/radio/ims/media/EvsBandwidth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V2-ndk_platform-source/gen/android/hardware/radio/ims/media/EvsMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V2-ndk_platform-source/gen/android/hardware/radio/ims/media/EvsParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V2-ndk_platform-source/gen/android/hardware/radio/ims/media/IImsMedia.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V2-ndk_platform-source/gen/android/hardware/radio/ims/media/IImsMediaListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V2-ndk_platform-source/gen/android/hardware/radio/ims/media/IImsMediaSession.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V2-ndk_platform-source/gen/android/hardware/radio/ims/media/IImsMediaSessionListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V2-ndk_platform-source/gen/android/hardware/radio/ims/media/LocalEndPoint.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V2-ndk_platform-source/gen/android/hardware/radio/ims/media/MediaDirection.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V2-ndk_platform-source/gen/android/hardware/radio/ims/media/MediaQualityStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V2-ndk_platform-source/gen/android/hardware/radio/ims/media/MediaQualityThreshold.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V2-ndk_platform-source/gen/android/hardware/radio/ims/media/RtcpConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V2-ndk_platform-source/gen/android/hardware/radio/ims/media/RtcpXrReportBlockType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V2-ndk_platform-source/gen/android/hardware/radio/ims/media/RtpAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V2-ndk_platform-source/gen/android/hardware/radio/ims/media/RtpConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V2-ndk_platform-source/gen/android/hardware/radio/ims/media/RtpError.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V2-ndk_platform-source/gen/android/hardware/radio/ims/media/RtpHeaderExtension.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V2-ndk_platform-source/gen/android/hardware/radio/ims/media/RtpSessionParams.cpp
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
    add_library(android.hardware.radio.ims.media-V2-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.ims.media-V2-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-ndk_platform-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.data-V2-ndk_platform-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.ims.media-V2-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.ims.media-V2-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.ims.media-V2-ndk_platform
            android.hardware.radio-V2-ndk_platform
            android.hardware.radio.data-V2-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.radio.ims.media-V2-ndk_platform ####################

#################### android.hardware.radio.messaging-V3-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V3-cpp-source/gen/android/hardware/radio/messaging/CdmaBroadcastSmsConfigInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V3-cpp-source/gen/android/hardware/radio/messaging/CdmaSmsAck.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V3-cpp-source/gen/android/hardware/radio/messaging/CdmaSmsAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V3-cpp-source/gen/android/hardware/radio/messaging/CdmaSmsMessage.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V3-cpp-source/gen/android/hardware/radio/messaging/CdmaSmsSubaddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V3-cpp-source/gen/android/hardware/radio/messaging/CdmaSmsWriteArgs.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V3-cpp-source/gen/android/hardware/radio/messaging/GsmBroadcastSmsConfigInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V3-cpp-source/gen/android/hardware/radio/messaging/GsmSmsMessage.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V3-cpp-source/gen/android/hardware/radio/messaging/IRadioMessaging.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V3-cpp-source/gen/android/hardware/radio/messaging/IRadioMessagingIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V3-cpp-source/gen/android/hardware/radio/messaging/IRadioMessagingResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V3-cpp-source/gen/android/hardware/radio/messaging/ImsSmsMessage.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V3-cpp-source/gen/android/hardware/radio/messaging/SendSmsResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V3-cpp-source/gen/android/hardware/radio/messaging/SmsAcknowledgeFailCause.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V3-cpp-source/gen/android/hardware/radio/messaging/SmsWriteArgs.cpp
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
    add_library(android.hardware.radio.messaging-V3-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.messaging-V3-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V3-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.messaging-V3-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.messaging-V3-cpp
            android.hardware.radio-V2-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libm
            libutils
    )
endif ()
#################### android.hardware.radio.messaging-V3-cpp ####################

#################### android.hardware.radio.sap-V1-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sap-V1-ndk_platform-source/gen/android/hardware/radio/sap/ISap.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sap-V1-ndk_platform-source/gen/android/hardware/radio/sap/ISapCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sap-V1-ndk_platform-source/gen/android/hardware/radio/sap/SapApduType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sap-V1-ndk_platform-source/gen/android/hardware/radio/sap/SapConnectRsp.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sap-V1-ndk_platform-source/gen/android/hardware/radio/sap/SapDisconnectType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sap-V1-ndk_platform-source/gen/android/hardware/radio/sap/SapResultCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sap-V1-ndk_platform-source/gen/android/hardware/radio/sap/SapStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sap-V1-ndk_platform-source/gen/android/hardware/radio/sap/SapTransferProtocol.cpp
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
    add_library(android.hardware.radio.sap-V1-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.sap-V1-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sap-V1-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.sap-V1-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.sap-V1-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.radio.sap-V1-ndk_platform ####################

#################### android.hardware.radio.network-V2-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk-source/gen/android/hardware/radio/network/AccessTechnologySpecificInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk-source/gen/android/hardware/radio/network/BarringInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk-source/gen/android/hardware/radio/network/BarringTypeSpecificInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk-source/gen/android/hardware/radio/network/Cdma2000RegistrationInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk-source/gen/android/hardware/radio/network/CdmaRoamingType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk-source/gen/android/hardware/radio/network/CdmaSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk-source/gen/android/hardware/radio/network/CellConnectionStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk-source/gen/android/hardware/radio/network/CellIdentity.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk-source/gen/android/hardware/radio/network/CellIdentityCdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk-source/gen/android/hardware/radio/network/CellIdentityGsm.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk-source/gen/android/hardware/radio/network/CellIdentityLte.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk-source/gen/android/hardware/radio/network/CellIdentityNr.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk-source/gen/android/hardware/radio/network/CellIdentityTdscdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk-source/gen/android/hardware/radio/network/CellIdentityWcdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk-source/gen/android/hardware/radio/network/CellInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk-source/gen/android/hardware/radio/network/CellInfoCdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk-source/gen/android/hardware/radio/network/CellInfoGsm.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk-source/gen/android/hardware/radio/network/CellInfoLte.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk-source/gen/android/hardware/radio/network/CellInfoNr.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk-source/gen/android/hardware/radio/network/CellInfoRatSpecificInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk-source/gen/android/hardware/radio/network/CellInfoTdscdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk-source/gen/android/hardware/radio/network/CellInfoWcdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk-source/gen/android/hardware/radio/network/ClosedSubscriberGroupInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk-source/gen/android/hardware/radio/network/Domain.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk-source/gen/android/hardware/radio/network/EmergencyMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk-source/gen/android/hardware/radio/network/EmergencyNetworkScanTrigger.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk-source/gen/android/hardware/radio/network/EmergencyRegResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk-source/gen/android/hardware/radio/network/EmergencyScanType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk-source/gen/android/hardware/radio/network/EutranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk-source/gen/android/hardware/radio/network/EutranRegistrationInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk-source/gen/android/hardware/radio/network/EvdoSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk-source/gen/android/hardware/radio/network/GeranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk-source/gen/android/hardware/radio/network/GsmSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk-source/gen/android/hardware/radio/network/IRadioNetwork.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk-source/gen/android/hardware/radio/network/IRadioNetworkIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk-source/gen/android/hardware/radio/network/IRadioNetworkResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk-source/gen/android/hardware/radio/network/IndicationFilter.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk-source/gen/android/hardware/radio/network/LceDataInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk-source/gen/android/hardware/radio/network/LinkCapacityEstimate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk-source/gen/android/hardware/radio/network/LteSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk-source/gen/android/hardware/radio/network/LteVopsInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk-source/gen/android/hardware/radio/network/NetworkScanRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk-source/gen/android/hardware/radio/network/NetworkScanResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk-source/gen/android/hardware/radio/network/NgranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk-source/gen/android/hardware/radio/network/NrDualConnectivityState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk-source/gen/android/hardware/radio/network/NrIndicators.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk-source/gen/android/hardware/radio/network/NrSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk-source/gen/android/hardware/radio/network/NrVopsInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk-source/gen/android/hardware/radio/network/OperatorInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk-source/gen/android/hardware/radio/network/PhoneRestrictedState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk-source/gen/android/hardware/radio/network/PhysicalChannelConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk-source/gen/android/hardware/radio/network/PhysicalChannelConfigBand.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk-source/gen/android/hardware/radio/network/RadioAccessSpecifier.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk-source/gen/android/hardware/radio/network/RadioAccessSpecifierBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk-source/gen/android/hardware/radio/network/RadioBandMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk-source/gen/android/hardware/radio/network/RegState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk-source/gen/android/hardware/radio/network/RegStateResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk-source/gen/android/hardware/radio/network/RegistrationFailCause.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk-source/gen/android/hardware/radio/network/SignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk-source/gen/android/hardware/radio/network/SignalThresholdInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk-source/gen/android/hardware/radio/network/SuppSvcNotification.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk-source/gen/android/hardware/radio/network/TdscdmaSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk-source/gen/android/hardware/radio/network/UsageSetting.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk-source/gen/android/hardware/radio/network/UtranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk-source/gen/android/hardware/radio/network/WcdmaSignalStrength.cpp
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
    add_library(android.hardware.radio.network-V2-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.network-V2-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-ndk-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.network-V2-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.network-V2-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.network-V2-ndk
            android.hardware.radio-V2-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.radio.network-V2-ndk ####################

#################### android.hardware.radio.messaging-V1-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V1-ndk-source/gen/android/hardware/radio/messaging/CdmaBroadcastSmsConfigInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V1-ndk-source/gen/android/hardware/radio/messaging/CdmaSmsAck.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V1-ndk-source/gen/android/hardware/radio/messaging/CdmaSmsAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V1-ndk-source/gen/android/hardware/radio/messaging/CdmaSmsMessage.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V1-ndk-source/gen/android/hardware/radio/messaging/CdmaSmsSubaddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V1-ndk-source/gen/android/hardware/radio/messaging/CdmaSmsWriteArgs.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V1-ndk-source/gen/android/hardware/radio/messaging/GsmBroadcastSmsConfigInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V1-ndk-source/gen/android/hardware/radio/messaging/GsmSmsMessage.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V1-ndk-source/gen/android/hardware/radio/messaging/IRadioMessaging.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V1-ndk-source/gen/android/hardware/radio/messaging/IRadioMessagingIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V1-ndk-source/gen/android/hardware/radio/messaging/IRadioMessagingResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V1-ndk-source/gen/android/hardware/radio/messaging/ImsSmsMessage.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V1-ndk-source/gen/android/hardware/radio/messaging/SendSmsResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V1-ndk-source/gen/android/hardware/radio/messaging/SmsAcknowledgeFailCause.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V1-ndk-source/gen/android/hardware/radio/messaging/SmsWriteArgs.cpp
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
    add_library(android.hardware.radio.messaging-V1-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.messaging-V1-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V1-ndk-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V1-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.messaging-V1-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.messaging-V1-ndk
            android.hardware.radio-V1-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.radio.messaging-V1-ndk ####################

#################### android.hardware.radio.messaging-V2-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V2-ndk_platform-source/gen/android/hardware/radio/messaging/CdmaBroadcastSmsConfigInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V2-ndk_platform-source/gen/android/hardware/radio/messaging/CdmaSmsAck.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V2-ndk_platform-source/gen/android/hardware/radio/messaging/CdmaSmsAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V2-ndk_platform-source/gen/android/hardware/radio/messaging/CdmaSmsMessage.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V2-ndk_platform-source/gen/android/hardware/radio/messaging/CdmaSmsSubaddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V2-ndk_platform-source/gen/android/hardware/radio/messaging/CdmaSmsWriteArgs.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V2-ndk_platform-source/gen/android/hardware/radio/messaging/GsmBroadcastSmsConfigInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V2-ndk_platform-source/gen/android/hardware/radio/messaging/GsmSmsMessage.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V2-ndk_platform-source/gen/android/hardware/radio/messaging/IRadioMessaging.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V2-ndk_platform-source/gen/android/hardware/radio/messaging/IRadioMessagingIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V2-ndk_platform-source/gen/android/hardware/radio/messaging/IRadioMessagingResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V2-ndk_platform-source/gen/android/hardware/radio/messaging/ImsSmsMessage.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V2-ndk_platform-source/gen/android/hardware/radio/messaging/SendSmsResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V2-ndk_platform-source/gen/android/hardware/radio/messaging/SmsAcknowledgeFailCause.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V2-ndk_platform-source/gen/android/hardware/radio/messaging/SmsWriteArgs.cpp
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
    add_library(android.hardware.radio.messaging-V2-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.messaging-V2-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-ndk_platform-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.messaging-V2-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.messaging-V2-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.messaging-V2-ndk_platform
            android.hardware.radio-V2-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.radio.messaging-V2-ndk_platform ####################

#################### android.hardware.radio.sim-V2-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-cpp-source/gen/android/hardware/radio/sim/AppStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-cpp-source/gen/android/hardware/radio/sim/CardPowerState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-cpp-source/gen/android/hardware/radio/sim/CardStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-cpp-source/gen/android/hardware/radio/sim/Carrier.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-cpp-source/gen/android/hardware/radio/sim/CarrierRestrictions.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-cpp-source/gen/android/hardware/radio/sim/CdmaSubscriptionSource.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-cpp-source/gen/android/hardware/radio/sim/IRadioSim.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-cpp-source/gen/android/hardware/radio/sim/IRadioSimIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-cpp-source/gen/android/hardware/radio/sim/IRadioSimResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-cpp-source/gen/android/hardware/radio/sim/IccIo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-cpp-source/gen/android/hardware/radio/sim/IccIoResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-cpp-source/gen/android/hardware/radio/sim/ImsiEncryptionInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-cpp-source/gen/android/hardware/radio/sim/PbReceivedStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-cpp-source/gen/android/hardware/radio/sim/PersoSubstate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-cpp-source/gen/android/hardware/radio/sim/PhonebookCapacity.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-cpp-source/gen/android/hardware/radio/sim/PhonebookRecordInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-cpp-source/gen/android/hardware/radio/sim/PinState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-cpp-source/gen/android/hardware/radio/sim/SelectUiccSub.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-cpp-source/gen/android/hardware/radio/sim/SessionInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-cpp-source/gen/android/hardware/radio/sim/SimApdu.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-cpp-source/gen/android/hardware/radio/sim/SimLockMultiSimPolicy.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-cpp-source/gen/android/hardware/radio/sim/SimRefreshResult.cpp
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
    add_library(android.hardware.radio.sim-V2-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.radio.sim-V2-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio-V2-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.config-V2-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/radio/aidl/android.hardware.radio.sim-V2-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.radio.sim-V2-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.radio.sim-V2-cpp
            android.hardware.radio-V2-cpp
            android.hardware.radio.config-V2-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libm
            libutils
    )
endif ()
#################### android.hardware.radio.sim-V2-cpp ####################
