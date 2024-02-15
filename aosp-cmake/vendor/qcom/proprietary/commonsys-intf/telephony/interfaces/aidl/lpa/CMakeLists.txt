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

project(vendor.qti.hardware.radio.lpa-V2-ndk_platform)

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

#################### vendor.qti.hardware.radio.lpa-V2-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V2-ndk_platform-source/gen/vendor/qti/hardware/radio/lpa/IUimLpa.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V2-ndk_platform-source/gen/vendor/qti/hardware/radio/lpa/IUimLpaIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V2-ndk_platform-source/gen/vendor/qti/hardware/radio/lpa/IUimLpaResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V2-ndk_platform-source/gen/vendor/qti/hardware/radio/lpa/UimLpaAddProfileFailureCause.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V2-ndk_platform-source/gen/vendor/qti/hardware/radio/lpa/UimLpaAddProfileProgressInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V2-ndk_platform-source/gen/vendor/qti/hardware/radio/lpa/UimLpaAddProfileStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V2-ndk_platform-source/gen/vendor/qti/hardware/radio/lpa/UimLpaEuiccResetMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V2-ndk_platform-source/gen/vendor/qti/hardware/radio/lpa/UimLpaHttpCustomHeader.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V2-ndk_platform-source/gen/vendor/qti/hardware/radio/lpa/UimLpaHttpTransactionInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V2-ndk_platform-source/gen/vendor/qti/hardware/radio/lpa/UimLpaIconType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V2-ndk_platform-source/gen/vendor/qti/hardware/radio/lpa/UimLpaProfileClassType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V2-ndk_platform-source/gen/vendor/qti/hardware/radio/lpa/UimLpaProfileInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V2-ndk_platform-source/gen/vendor/qti/hardware/radio/lpa/UimLpaProfilePolicyMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V2-ndk_platform-source/gen/vendor/qti/hardware/radio/lpa/UimLpaProfileState.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V2-ndk_platform-source/gen/vendor/qti/hardware/radio/lpa/UimLpaResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V2-ndk_platform-source/gen/vendor/qti/hardware/radio/lpa/UimLpaSrvAddrOp.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V2-ndk_platform-source/gen/vendor/qti/hardware/radio/lpa/UimLpaSrvAddrOpReq.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V2-ndk_platform-source/gen/vendor/qti/hardware/radio/lpa/UimLpaSrvAddrOpResp.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V2-ndk_platform-source/gen/vendor/qti/hardware/radio/lpa/UimLpaUserConsentType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V2-ndk_platform-source/gen/vendor/qti/hardware/radio/lpa/UimLpaUserEventId.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V2-ndk_platform-source/gen/vendor/qti/hardware/radio/lpa/UimLpaUserReq.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V2-ndk_platform-source/gen/vendor/qti/hardware/radio/lpa/UimLpaUserResp.cpp
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
    add_library(vendor.qti.hardware.radio.lpa-V2-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(vendor.qti.hardware.radio.lpa-V2-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V2-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.qti.hardware.radio.lpa-V2-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.qti.hardware.radio.lpa-V2-ndk_platform
            libbinder_ndk
            libbinder_ndk
            libc++
            libc++
            libc++_static
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
            libc
            libdl
            libm
    )
endif ()
#################### vendor.qti.hardware.radio.lpa-V2-ndk_platform ####################

#################### vendor.qti.hardware.radio.lpa-V2-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V2-ndk-source/gen/vendor/qti/hardware/radio/lpa/IUimLpa.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V2-ndk-source/gen/vendor/qti/hardware/radio/lpa/IUimLpaIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V2-ndk-source/gen/vendor/qti/hardware/radio/lpa/IUimLpaResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V2-ndk-source/gen/vendor/qti/hardware/radio/lpa/UimLpaAddProfileFailureCause.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V2-ndk-source/gen/vendor/qti/hardware/radio/lpa/UimLpaAddProfileProgressInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V2-ndk-source/gen/vendor/qti/hardware/radio/lpa/UimLpaAddProfileStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V2-ndk-source/gen/vendor/qti/hardware/radio/lpa/UimLpaEuiccResetMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V2-ndk-source/gen/vendor/qti/hardware/radio/lpa/UimLpaHttpCustomHeader.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V2-ndk-source/gen/vendor/qti/hardware/radio/lpa/UimLpaHttpTransactionInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V2-ndk-source/gen/vendor/qti/hardware/radio/lpa/UimLpaIconType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V2-ndk-source/gen/vendor/qti/hardware/radio/lpa/UimLpaProfileClassType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V2-ndk-source/gen/vendor/qti/hardware/radio/lpa/UimLpaProfileInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V2-ndk-source/gen/vendor/qti/hardware/radio/lpa/UimLpaProfilePolicyMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V2-ndk-source/gen/vendor/qti/hardware/radio/lpa/UimLpaProfileState.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V2-ndk-source/gen/vendor/qti/hardware/radio/lpa/UimLpaResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V2-ndk-source/gen/vendor/qti/hardware/radio/lpa/UimLpaSrvAddrOp.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V2-ndk-source/gen/vendor/qti/hardware/radio/lpa/UimLpaSrvAddrOpReq.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V2-ndk-source/gen/vendor/qti/hardware/radio/lpa/UimLpaSrvAddrOpResp.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V2-ndk-source/gen/vendor/qti/hardware/radio/lpa/UimLpaUserConsentType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V2-ndk-source/gen/vendor/qti/hardware/radio/lpa/UimLpaUserEventId.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V2-ndk-source/gen/vendor/qti/hardware/radio/lpa/UimLpaUserReq.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V2-ndk-source/gen/vendor/qti/hardware/radio/lpa/UimLpaUserResp.cpp
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
    add_library(vendor.qti.hardware.radio.lpa-V2-ndk SHARED ${SOURCE_FILES})

    target_include_directories(vendor.qti.hardware.radio.lpa-V2-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V2-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.qti.hardware.radio.lpa-V2-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.qti.hardware.radio.lpa-V2-ndk
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
#################### vendor.qti.hardware.radio.lpa-V2-ndk ####################

#################### vendor.qti.hardware.radio.lpa-V1-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V1-ndk_platform-source/gen/vendor/qti/hardware/radio/lpa/IUimLpa.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V1-ndk_platform-source/gen/vendor/qti/hardware/radio/lpa/IUimLpaIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V1-ndk_platform-source/gen/vendor/qti/hardware/radio/lpa/IUimLpaResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V1-ndk_platform-source/gen/vendor/qti/hardware/radio/lpa/UimLpaAddProfileFailureCause.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V1-ndk_platform-source/gen/vendor/qti/hardware/radio/lpa/UimLpaAddProfileProgressInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V1-ndk_platform-source/gen/vendor/qti/hardware/radio/lpa/UimLpaAddProfileStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V1-ndk_platform-source/gen/vendor/qti/hardware/radio/lpa/UimLpaEuiccResetMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V1-ndk_platform-source/gen/vendor/qti/hardware/radio/lpa/UimLpaHttpCustomHeader.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V1-ndk_platform-source/gen/vendor/qti/hardware/radio/lpa/UimLpaHttpTransactionInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V1-ndk_platform-source/gen/vendor/qti/hardware/radio/lpa/UimLpaIconType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V1-ndk_platform-source/gen/vendor/qti/hardware/radio/lpa/UimLpaProfileClassType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V1-ndk_platform-source/gen/vendor/qti/hardware/radio/lpa/UimLpaProfileInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V1-ndk_platform-source/gen/vendor/qti/hardware/radio/lpa/UimLpaProfilePolicyMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V1-ndk_platform-source/gen/vendor/qti/hardware/radio/lpa/UimLpaProfileState.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V1-ndk_platform-source/gen/vendor/qti/hardware/radio/lpa/UimLpaResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V1-ndk_platform-source/gen/vendor/qti/hardware/radio/lpa/UimLpaSrvAddrOp.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V1-ndk_platform-source/gen/vendor/qti/hardware/radio/lpa/UimLpaSrvAddrOpReq.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V1-ndk_platform-source/gen/vendor/qti/hardware/radio/lpa/UimLpaSrvAddrOpResp.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V1-ndk_platform-source/gen/vendor/qti/hardware/radio/lpa/UimLpaUserConsentType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V1-ndk_platform-source/gen/vendor/qti/hardware/radio/lpa/UimLpaUserEventId.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V1-ndk_platform-source/gen/vendor/qti/hardware/radio/lpa/UimLpaUserReq.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V1-ndk_platform-source/gen/vendor/qti/hardware/radio/lpa/UimLpaUserResp.cpp
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
    add_library(vendor.qti.hardware.radio.lpa-V1-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(vendor.qti.hardware.radio.lpa-V1-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V1-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.qti.hardware.radio.lpa-V1-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.qti.hardware.radio.lpa-V1-ndk_platform
            libbinder_ndk
            libbinder_ndk
            libc++
            libc++
            libc++_static
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
            libc
            libdl
            libm
    )
endif ()
#################### vendor.qti.hardware.radio.lpa-V1-ndk_platform ####################

#################### vendor.qti.hardware.radio.lpa-V1-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V1-ndk-source/gen/vendor/qti/hardware/radio/lpa/IUimLpa.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V1-ndk-source/gen/vendor/qti/hardware/radio/lpa/IUimLpaIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V1-ndk-source/gen/vendor/qti/hardware/radio/lpa/IUimLpaResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V1-ndk-source/gen/vendor/qti/hardware/radio/lpa/UimLpaAddProfileFailureCause.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V1-ndk-source/gen/vendor/qti/hardware/radio/lpa/UimLpaAddProfileProgressInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V1-ndk-source/gen/vendor/qti/hardware/radio/lpa/UimLpaAddProfileStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V1-ndk-source/gen/vendor/qti/hardware/radio/lpa/UimLpaEuiccResetMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V1-ndk-source/gen/vendor/qti/hardware/radio/lpa/UimLpaHttpCustomHeader.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V1-ndk-source/gen/vendor/qti/hardware/radio/lpa/UimLpaHttpTransactionInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V1-ndk-source/gen/vendor/qti/hardware/radio/lpa/UimLpaIconType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V1-ndk-source/gen/vendor/qti/hardware/radio/lpa/UimLpaProfileClassType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V1-ndk-source/gen/vendor/qti/hardware/radio/lpa/UimLpaProfileInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V1-ndk-source/gen/vendor/qti/hardware/radio/lpa/UimLpaProfilePolicyMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V1-ndk-source/gen/vendor/qti/hardware/radio/lpa/UimLpaProfileState.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V1-ndk-source/gen/vendor/qti/hardware/radio/lpa/UimLpaResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V1-ndk-source/gen/vendor/qti/hardware/radio/lpa/UimLpaSrvAddrOp.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V1-ndk-source/gen/vendor/qti/hardware/radio/lpa/UimLpaSrvAddrOpReq.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V1-ndk-source/gen/vendor/qti/hardware/radio/lpa/UimLpaSrvAddrOpResp.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V1-ndk-source/gen/vendor/qti/hardware/radio/lpa/UimLpaUserConsentType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V1-ndk-source/gen/vendor/qti/hardware/radio/lpa/UimLpaUserEventId.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V1-ndk-source/gen/vendor/qti/hardware/radio/lpa/UimLpaUserReq.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V1-ndk-source/gen/vendor/qti/hardware/radio/lpa/UimLpaUserResp.cpp
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
    add_library(vendor.qti.hardware.radio.lpa-V1-ndk SHARED ${SOURCE_FILES})

    target_include_directories(vendor.qti.hardware.radio.lpa-V1-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/lpa/vendor.qti.hardware.radio.lpa-V1-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.qti.hardware.radio.lpa-V1-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.qti.hardware.radio.lpa-V1-ndk
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
#################### vendor.qti.hardware.radio.lpa-V1-ndk ####################
