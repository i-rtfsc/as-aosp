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

project(libaudio_aidl_conversion_common_ndk)

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

#################### libaudio_aidl_conversion_common_ndk ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/audioaidlconversion/AidlConversionCppNdk.cpp
        ${ANDROID_ROOT}/frameworks/av/media/audioaidlconversion/AidlConversionNdk.cpp
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
    add_library(libaudio_aidl_conversion_common_ndk SHARED ${SOURCE_FILES})

    target_include_directories(libaudio_aidl_conversion_common_ndk
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/audioaidlconversion/include
            ${ANDROID_ROOT}/frameworks/av/media/audioaidlconversion
    )

    # Add c/c++ flags
    target_compile_options(libaudio_aidl_conversion_common_ndk PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libaudio_aidl_conversion_common_ndk
            android.hardware.audio.common-V2-ndk
            android.media.audio.common.types-V2-ndk
            framework-permission-aidl-cpp
            libaudio_aidl_conversion_common_util_ndk
            libaudio_system_headers
            libaudioaidlcommon
            libbase
            libbinder
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libhardware_headers
            liblog
            libm
            libshmemcompat
            libstagefright_foundation
            libutils
            shared-file-region-aidl-cpp
    )
endif ()
#################### libaudio_aidl_conversion_common_ndk ####################

#################### libaudio_aidl_conversion_common_ndk_cpp ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/audioaidlconversion/AidlConversionNdkCpp.cpp
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
    add_library(libaudio_aidl_conversion_common_ndk_cpp SHARED ${SOURCE_FILES})

    target_include_directories(libaudio_aidl_conversion_common_ndk_cpp
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/audioaidlconversion/include
            ${ANDROID_ROOT}/frameworks/av/media/audioaidlconversion
    )

    # Add c/c++ flags
    target_compile_options(libaudio_aidl_conversion_common_ndk_cpp PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libaudio_aidl_conversion_common_ndk_cpp
            android.media.audio.common.types-V2-cpp
            android.media.audio.common.types-V2-ndk
            framework-permission-aidl-cpp
            libaudio_system_headers
            libbase
            libbase_headers
            libbinder
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            liberror_headers
            libhardware_headers
            liblog
            libm
            libshmemcompat
            libstagefright_foundation
            libutils
            shared-file-region-aidl-cpp
    )
endif ()
#################### libaudio_aidl_conversion_common_ndk_cpp ####################

#################### libaudio_aidl_conversion_effect_ndk ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/audioaidlconversion/AidlConversionEffect.cpp
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
    add_library(libaudio_aidl_conversion_effect_ndk SHARED ${SOURCE_FILES})

    target_include_directories(libaudio_aidl_conversion_effect_ndk
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/audioaidlconversion/include
            ${ANDROID_ROOT}/frameworks/av/media/audioaidlconversion
    )

    # Add c/c++ flags
    target_compile_options(libaudio_aidl_conversion_effect_ndk PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libaudio_aidl_conversion_effect_ndk
            android.hardware.audio.common-V2-ndk
            android.hardware.audio.effect-V1-ndk
            android.media.audio.common.types-V2-ndk
            framework-permission-aidl-cpp
            libaudio_aidl_conversion_common_ndk
            libaudio_aidl_conversion_common_util_ndk
            libaudio_system_headers
            libbase
            libbinder
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libhardware_headers
            liblog
            libm
            libshmemcompat
            libstagefright_foundation
            libutils
            shared-file-region-aidl-cpp
    )
endif ()
#################### libaudio_aidl_conversion_effect_ndk ####################

#################### libaudio_aidl_conversion_common_cpp ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/audioaidlconversion/AidlConversionCppNdk.cpp
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
    add_library(libaudio_aidl_conversion_common_cpp SHARED ${SOURCE_FILES})

    target_include_directories(libaudio_aidl_conversion_common_cpp
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/audioaidlconversion/include
            ${ANDROID_ROOT}/frameworks/av/media/audioaidlconversion
    )

    # Add c/c++ flags
    target_compile_options(libaudio_aidl_conversion_common_cpp PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libaudio_aidl_conversion_common_cpp
            android.media.audio.common.types-V2-cpp
            framework-permission-aidl-cpp
            libaudio_aidl_conversion_common_util_cpp
            libaudio_system_headers
            libbase
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libhardware_headers
            liblog
            libm
            libshmemcompat
            libstagefright_foundation
            libutils
            shared-file-region-aidl-cpp
    )
endif ()
#################### libaudio_aidl_conversion_common_cpp ####################

#################### libaudio_aidl_conversion_core_ndk ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/audioaidlconversion/AidlConversionCore.cpp
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
    add_library(libaudio_aidl_conversion_core_ndk SHARED ${SOURCE_FILES})

    target_include_directories(libaudio_aidl_conversion_core_ndk
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/audioaidlconversion/include
            ${ANDROID_ROOT}/frameworks/av/media/audioaidlconversion
    )

    # Add c/c++ flags
    target_compile_options(libaudio_aidl_conversion_core_ndk PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libaudio_aidl_conversion_core_ndk
            android.hardware.audio.common-V2-ndk
            android.hardware.audio.core-V1-ndk
            android.media.audio.common.types-V2-ndk
            framework-permission-aidl-cpp
            libaudio_aidl_conversion_common_ndk
            libaudio_aidl_conversion_common_util_ndk
            libaudio_system_headers
            libbase
            libbinder
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libhardware_headers
            liblog
            libm
            libshmemcompat
            libstagefright_foundation
            libutils
            shared-file-region-aidl-cpp
    )
endif ()
#################### libaudio_aidl_conversion_core_ndk ####################
