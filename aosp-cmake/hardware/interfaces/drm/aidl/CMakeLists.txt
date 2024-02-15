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

project(android.hardware.drm-V2-ndk)

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

#################### android.hardware.drm-V2-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk-source/gen/android/hardware/drm/CryptoSchemes.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk-source/gen/android/hardware/drm/DecryptArgs.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk-source/gen/android/hardware/drm/DestinationBuffer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk-source/gen/android/hardware/drm/DrmMetric.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk-source/gen/android/hardware/drm/DrmMetricGroup.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk-source/gen/android/hardware/drm/DrmMetricNamedValue.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk-source/gen/android/hardware/drm/DrmMetricValue.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk-source/gen/android/hardware/drm/EventType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk-source/gen/android/hardware/drm/HdcpLevel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk-source/gen/android/hardware/drm/HdcpLevels.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk-source/gen/android/hardware/drm/ICryptoPlugin.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk-source/gen/android/hardware/drm/IDrmFactory.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk-source/gen/android/hardware/drm/IDrmPlugin.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk-source/gen/android/hardware/drm/IDrmPluginListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk-source/gen/android/hardware/drm/KeyRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk-source/gen/android/hardware/drm/KeyRequestType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk-source/gen/android/hardware/drm/KeySetId.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk-source/gen/android/hardware/drm/KeyStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk-source/gen/android/hardware/drm/KeyStatusType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk-source/gen/android/hardware/drm/KeyType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk-source/gen/android/hardware/drm/KeyValue.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk-source/gen/android/hardware/drm/LogMessage.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk-source/gen/android/hardware/drm/LogPriority.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk-source/gen/android/hardware/drm/Mode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk-source/gen/android/hardware/drm/NumberOfSessions.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk-source/gen/android/hardware/drm/OfflineLicenseState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk-source/gen/android/hardware/drm/OpaqueData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk-source/gen/android/hardware/drm/Pattern.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk-source/gen/android/hardware/drm/ProvideProvisionResponseResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk-source/gen/android/hardware/drm/ProvisionRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk-source/gen/android/hardware/drm/SecureStop.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk-source/gen/android/hardware/drm/SecureStopId.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk-source/gen/android/hardware/drm/SecurityLevel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk-source/gen/android/hardware/drm/SharedBuffer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk-source/gen/android/hardware/drm/Status.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk-source/gen/android/hardware/drm/SubSample.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk-source/gen/android/hardware/drm/SupportedContentType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk-source/gen/android/hardware/drm/Uuid.cpp
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
    add_library(android.hardware.drm-V2-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.drm-V2-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.drm-V2-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.drm-V2-ndk
            android.hardware.common-V2-ndk
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
#################### android.hardware.drm-V2-ndk ####################

#################### android.hardware.drm-V1-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk_platform-source/gen/android/hardware/drm/CryptoSchemes.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk_platform-source/gen/android/hardware/drm/DecryptArgs.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk_platform-source/gen/android/hardware/drm/DestinationBuffer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk_platform-source/gen/android/hardware/drm/DrmMetric.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk_platform-source/gen/android/hardware/drm/DrmMetricGroup.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk_platform-source/gen/android/hardware/drm/DrmMetricNamedValue.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk_platform-source/gen/android/hardware/drm/DrmMetricValue.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk_platform-source/gen/android/hardware/drm/EventType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk_platform-source/gen/android/hardware/drm/HdcpLevel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk_platform-source/gen/android/hardware/drm/HdcpLevels.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk_platform-source/gen/android/hardware/drm/ICryptoPlugin.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk_platform-source/gen/android/hardware/drm/IDrmFactory.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk_platform-source/gen/android/hardware/drm/IDrmPlugin.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk_platform-source/gen/android/hardware/drm/IDrmPluginListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk_platform-source/gen/android/hardware/drm/KeyRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk_platform-source/gen/android/hardware/drm/KeyRequestType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk_platform-source/gen/android/hardware/drm/KeySetId.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk_platform-source/gen/android/hardware/drm/KeyStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk_platform-source/gen/android/hardware/drm/KeyStatusType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk_platform-source/gen/android/hardware/drm/KeyType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk_platform-source/gen/android/hardware/drm/KeyValue.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk_platform-source/gen/android/hardware/drm/LogMessage.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk_platform-source/gen/android/hardware/drm/LogPriority.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk_platform-source/gen/android/hardware/drm/Mode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk_platform-source/gen/android/hardware/drm/NumberOfSessions.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk_platform-source/gen/android/hardware/drm/OfflineLicenseState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk_platform-source/gen/android/hardware/drm/OpaqueData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk_platform-source/gen/android/hardware/drm/Pattern.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk_platform-source/gen/android/hardware/drm/ProvideProvisionResponseResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk_platform-source/gen/android/hardware/drm/ProvisionRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk_platform-source/gen/android/hardware/drm/SecureStop.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk_platform-source/gen/android/hardware/drm/SecureStopId.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk_platform-source/gen/android/hardware/drm/SecurityLevel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk_platform-source/gen/android/hardware/drm/SharedBuffer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk_platform-source/gen/android/hardware/drm/Status.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk_platform-source/gen/android/hardware/drm/SubSample.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk_platform-source/gen/android/hardware/drm/SupportedContentType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk_platform-source/gen/android/hardware/drm/Uuid.cpp
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
    add_library(android.hardware.drm-V1-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.drm-V1-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.drm-V1-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.drm-V1-ndk_platform
            android.hardware.common-V2-ndk_platform
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
#################### android.hardware.drm-V1-ndk_platform ####################

#################### android.hardware.drm-V1-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk-source/gen/android/hardware/drm/CryptoSchemes.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk-source/gen/android/hardware/drm/DecryptArgs.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk-source/gen/android/hardware/drm/DestinationBuffer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk-source/gen/android/hardware/drm/DrmMetric.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk-source/gen/android/hardware/drm/DrmMetricGroup.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk-source/gen/android/hardware/drm/DrmMetricNamedValue.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk-source/gen/android/hardware/drm/DrmMetricValue.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk-source/gen/android/hardware/drm/EventType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk-source/gen/android/hardware/drm/HdcpLevel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk-source/gen/android/hardware/drm/HdcpLevels.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk-source/gen/android/hardware/drm/ICryptoPlugin.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk-source/gen/android/hardware/drm/IDrmFactory.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk-source/gen/android/hardware/drm/IDrmPlugin.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk-source/gen/android/hardware/drm/IDrmPluginListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk-source/gen/android/hardware/drm/KeyRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk-source/gen/android/hardware/drm/KeyRequestType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk-source/gen/android/hardware/drm/KeySetId.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk-source/gen/android/hardware/drm/KeyStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk-source/gen/android/hardware/drm/KeyStatusType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk-source/gen/android/hardware/drm/KeyType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk-source/gen/android/hardware/drm/KeyValue.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk-source/gen/android/hardware/drm/LogMessage.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk-source/gen/android/hardware/drm/LogPriority.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk-source/gen/android/hardware/drm/Mode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk-source/gen/android/hardware/drm/NumberOfSessions.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk-source/gen/android/hardware/drm/OfflineLicenseState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk-source/gen/android/hardware/drm/OpaqueData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk-source/gen/android/hardware/drm/Pattern.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk-source/gen/android/hardware/drm/ProvideProvisionResponseResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk-source/gen/android/hardware/drm/ProvisionRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk-source/gen/android/hardware/drm/SecureStop.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk-source/gen/android/hardware/drm/SecureStopId.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk-source/gen/android/hardware/drm/SecurityLevel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk-source/gen/android/hardware/drm/SharedBuffer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk-source/gen/android/hardware/drm/Status.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk-source/gen/android/hardware/drm/SubSample.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk-source/gen/android/hardware/drm/SupportedContentType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk-source/gen/android/hardware/drm/Uuid.cpp
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
    add_library(android.hardware.drm-V1-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.drm-V1-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.drm-V1-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.drm-V1-ndk
            android.hardware.common-V2-ndk
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
#################### android.hardware.drm-V1-ndk ####################

#################### android.hardware.drm-V2-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk_platform-source/gen/android/hardware/drm/CryptoSchemes.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk_platform-source/gen/android/hardware/drm/DecryptArgs.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk_platform-source/gen/android/hardware/drm/DestinationBuffer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk_platform-source/gen/android/hardware/drm/DrmMetric.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk_platform-source/gen/android/hardware/drm/DrmMetricGroup.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk_platform-source/gen/android/hardware/drm/DrmMetricNamedValue.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk_platform-source/gen/android/hardware/drm/DrmMetricValue.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk_platform-source/gen/android/hardware/drm/EventType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk_platform-source/gen/android/hardware/drm/HdcpLevel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk_platform-source/gen/android/hardware/drm/HdcpLevels.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk_platform-source/gen/android/hardware/drm/ICryptoPlugin.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk_platform-source/gen/android/hardware/drm/IDrmFactory.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk_platform-source/gen/android/hardware/drm/IDrmPlugin.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk_platform-source/gen/android/hardware/drm/IDrmPluginListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk_platform-source/gen/android/hardware/drm/KeyRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk_platform-source/gen/android/hardware/drm/KeyRequestType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk_platform-source/gen/android/hardware/drm/KeySetId.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk_platform-source/gen/android/hardware/drm/KeyStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk_platform-source/gen/android/hardware/drm/KeyStatusType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk_platform-source/gen/android/hardware/drm/KeyType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk_platform-source/gen/android/hardware/drm/KeyValue.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk_platform-source/gen/android/hardware/drm/LogMessage.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk_platform-source/gen/android/hardware/drm/LogPriority.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk_platform-source/gen/android/hardware/drm/Mode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk_platform-source/gen/android/hardware/drm/NumberOfSessions.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk_platform-source/gen/android/hardware/drm/OfflineLicenseState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk_platform-source/gen/android/hardware/drm/OpaqueData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk_platform-source/gen/android/hardware/drm/Pattern.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk_platform-source/gen/android/hardware/drm/ProvideProvisionResponseResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk_platform-source/gen/android/hardware/drm/ProvisionRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk_platform-source/gen/android/hardware/drm/SecureStop.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk_platform-source/gen/android/hardware/drm/SecureStopId.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk_platform-source/gen/android/hardware/drm/SecurityLevel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk_platform-source/gen/android/hardware/drm/SharedBuffer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk_platform-source/gen/android/hardware/drm/Status.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk_platform-source/gen/android/hardware/drm/SubSample.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk_platform-source/gen/android/hardware/drm/SupportedContentType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk_platform-source/gen/android/hardware/drm/Uuid.cpp
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
    add_library(android.hardware.drm-V2-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.drm-V2-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/drm/aidl/android.hardware.drm-V2-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.drm-V2-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.drm-V2-ndk_platform
            android.hardware.common-V2-ndk_platform
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
#################### android.hardware.drm-V2-ndk_platform ####################
