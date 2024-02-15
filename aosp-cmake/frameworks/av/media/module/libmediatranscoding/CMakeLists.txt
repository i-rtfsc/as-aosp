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

project(mediatranscoding_aidl_interface-cpp-analyzer)

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

#################### mediatranscoding_aidl_interface-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/module/libmediatranscoding/mediatranscoding_aidl_interface-ndk-source/gen/android/media/IMediaTranscodingService.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/module/libmediatranscoding/mediatranscoding_aidl_interface-ndk-source/gen/android/media/ITranscodingClient.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/module/libmediatranscoding/mediatranscoding_aidl_interface-ndk-source/gen/android/media/ITranscodingClientCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/module/libmediatranscoding/mediatranscoding_aidl_interface-ndk-source/gen/android/media/TranscodingErrorCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/module/libmediatranscoding/mediatranscoding_aidl_interface-ndk-source/gen/android/media/TranscodingSessionPriority.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/module/libmediatranscoding/mediatranscoding_aidl_interface-ndk-source/gen/android/media/TranscodingSessionStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/module/libmediatranscoding/mediatranscoding_aidl_interface-ndk-source/gen/android/media/TranscodingType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/module/libmediatranscoding/mediatranscoding_aidl_interface-ndk-source/gen/android/media/TranscodingVideoCodecType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/module/libmediatranscoding/mediatranscoding_aidl_interface-ndk-source/gen/android/media/TranscodingVideoTrackFormat.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/module/libmediatranscoding/mediatranscoding_aidl_interface-ndk-source/gen/android/media/TranscodingSessionParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/module/libmediatranscoding/mediatranscoding_aidl_interface-ndk-source/gen/android/media/TranscodingRequestParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/module/libmediatranscoding/mediatranscoding_aidl_interface-ndk-source/gen/android/media/TranscodingResultParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/module/libmediatranscoding/mediatranscoding_aidl_interface-ndk-source/gen/android/media/TranscodingTestConfig.cpp
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
    add_library(mediatranscoding_aidl_interface-ndk SHARED ${SOURCE_FILES})

    target_include_directories(mediatranscoding_aidl_interface-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/module/libmediatranscoding/mediatranscoding_aidl_interface-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(mediatranscoding_aidl_interface-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(mediatranscoding_aidl_interface-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### mediatranscoding_aidl_interface-ndk ####################

#################### mediatranscoding_aidl_interface-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/module/libmediatranscoding/mediatranscoding_aidl_interface-ndk_platform-source/gen/android/media/IMediaTranscodingService.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/module/libmediatranscoding/mediatranscoding_aidl_interface-ndk_platform-source/gen/android/media/ITranscodingClient.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/module/libmediatranscoding/mediatranscoding_aidl_interface-ndk_platform-source/gen/android/media/ITranscodingClientCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/module/libmediatranscoding/mediatranscoding_aidl_interface-ndk_platform-source/gen/android/media/TranscodingErrorCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/module/libmediatranscoding/mediatranscoding_aidl_interface-ndk_platform-source/gen/android/media/TranscodingSessionPriority.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/module/libmediatranscoding/mediatranscoding_aidl_interface-ndk_platform-source/gen/android/media/TranscodingSessionStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/module/libmediatranscoding/mediatranscoding_aidl_interface-ndk_platform-source/gen/android/media/TranscodingType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/module/libmediatranscoding/mediatranscoding_aidl_interface-ndk_platform-source/gen/android/media/TranscodingVideoCodecType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/module/libmediatranscoding/mediatranscoding_aidl_interface-ndk_platform-source/gen/android/media/TranscodingVideoTrackFormat.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/module/libmediatranscoding/mediatranscoding_aidl_interface-ndk_platform-source/gen/android/media/TranscodingSessionParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/module/libmediatranscoding/mediatranscoding_aidl_interface-ndk_platform-source/gen/android/media/TranscodingRequestParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/module/libmediatranscoding/mediatranscoding_aidl_interface-ndk_platform-source/gen/android/media/TranscodingResultParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/module/libmediatranscoding/mediatranscoding_aidl_interface-ndk_platform-source/gen/android/media/TranscodingTestConfig.cpp
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
    add_library(mediatranscoding_aidl_interface-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(mediatranscoding_aidl_interface-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/module/libmediatranscoding/mediatranscoding_aidl_interface-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(mediatranscoding_aidl_interface-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(mediatranscoding_aidl_interface-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### mediatranscoding_aidl_interface-ndk_platform ####################

#################### libmediatranscoding ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/module/libmediatranscoding/TranscoderWrapper.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/libmediatranscoding/TranscodingClientManager.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/libmediatranscoding/TranscodingLogger.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/libmediatranscoding/TranscodingResourcePolicy.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/libmediatranscoding/TranscodingSessionController.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/libmediatranscoding/TranscodingThermalPolicy.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/libmediatranscoding/TranscodingUidPolicy.cpp
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
    add_library(libmediatranscoding SHARED ${SOURCE_FILES})

    target_include_directories(libmediatranscoding
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/module/libmediatranscoding/include
            ${ANDROID_ROOT}/frameworks/av/media/module/libmediatranscoding
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/module/libmediatranscoding/mediatranscoding_aidl_interface-ndk-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/module/libmediatranscoding/statslog_media.h/gen
            ${ANDROID_ROOT}/frameworks/av/media/module/libmediatranscoding/transcoder/include
    )

    # Add c/c++ flags
    target_compile_options(libmediatranscoding PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libmediatranscoding
            libandroid
            libbinder_ndk
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            liblog
            libm
            libmediandk
            libmediatranscoder
            libstatslog_media
            libstatssocket
            libutils
            mediatranscoding_aidl_interface-ndk
            resourceobserver_aidl_interface-V1-ndk
    )
endif ()
#################### libmediatranscoding ####################

#################### mediatranscoding_aidl_interface-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/module/libmediatranscoding/mediatranscoding_aidl_interface-cpp-source/gen/android/media/IMediaTranscodingService.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/module/libmediatranscoding/mediatranscoding_aidl_interface-cpp-source/gen/android/media/ITranscodingClient.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/module/libmediatranscoding/mediatranscoding_aidl_interface-cpp-source/gen/android/media/ITranscodingClientCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/module/libmediatranscoding/mediatranscoding_aidl_interface-cpp-source/gen/android/media/TranscodingErrorCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/module/libmediatranscoding/mediatranscoding_aidl_interface-cpp-source/gen/android/media/TranscodingSessionPriority.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/module/libmediatranscoding/mediatranscoding_aidl_interface-cpp-source/gen/android/media/TranscodingSessionStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/module/libmediatranscoding/mediatranscoding_aidl_interface-cpp-source/gen/android/media/TranscodingType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/module/libmediatranscoding/mediatranscoding_aidl_interface-cpp-source/gen/android/media/TranscodingVideoCodecType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/module/libmediatranscoding/mediatranscoding_aidl_interface-cpp-source/gen/android/media/TranscodingVideoTrackFormat.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/module/libmediatranscoding/mediatranscoding_aidl_interface-cpp-source/gen/android/media/TranscodingSessionParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/module/libmediatranscoding/mediatranscoding_aidl_interface-cpp-source/gen/android/media/TranscodingRequestParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/module/libmediatranscoding/mediatranscoding_aidl_interface-cpp-source/gen/android/media/TranscodingResultParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/module/libmediatranscoding/mediatranscoding_aidl_interface-cpp-source/gen/android/media/TranscodingTestConfig.cpp
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
    add_library(mediatranscoding_aidl_interface-cpp SHARED ${SOURCE_FILES})

    target_include_directories(mediatranscoding_aidl_interface-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/module/libmediatranscoding/mediatranscoding_aidl_interface-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(mediatranscoding_aidl_interface-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(mediatranscoding_aidl_interface-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libm
            libutils
    )
endif ()
#################### mediatranscoding_aidl_interface-cpp ####################

#################### mediatranscoding_aidl_interface-cpp-analyzer ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/module/libmediatranscoding/mediatranscoding_aidl_interface-cpp-analyzer-source/gen/android/media/IMediaTranscodingService.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/module/libmediatranscoding/mediatranscoding_aidl_interface-cpp-analyzer-source/gen/android/media/ITranscodingClient.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/module/libmediatranscoding/mediatranscoding_aidl_interface-cpp-analyzer-source/gen/android/media/ITranscodingClientCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/module/libmediatranscoding/mediatranscoding_aidl_interface-cpp-analyzer-source/gen/android/media/TranscodingErrorCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/module/libmediatranscoding/mediatranscoding_aidl_interface-cpp-analyzer-source/gen/android/media/TranscodingSessionPriority.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/module/libmediatranscoding/mediatranscoding_aidl_interface-cpp-analyzer-source/gen/android/media/TranscodingSessionStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/module/libmediatranscoding/mediatranscoding_aidl_interface-cpp-analyzer-source/gen/android/media/TranscodingType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/module/libmediatranscoding/mediatranscoding_aidl_interface-cpp-analyzer-source/gen/android/media/TranscodingVideoCodecType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/module/libmediatranscoding/mediatranscoding_aidl_interface-cpp-analyzer-source/gen/android/media/TranscodingVideoTrackFormat.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/module/libmediatranscoding/mediatranscoding_aidl_interface-cpp-analyzer-source/gen/android/media/TranscodingSessionParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/module/libmediatranscoding/mediatranscoding_aidl_interface-cpp-analyzer-source/gen/android/media/TranscodingRequestParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/module/libmediatranscoding/mediatranscoding_aidl_interface-cpp-analyzer-source/gen/android/media/TranscodingResultParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/module/libmediatranscoding/mediatranscoding_aidl_interface-cpp-analyzer-source/gen/android/media/TranscodingTestConfig.cpp
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
    add_executable(mediatranscoding_aidl_interface-cpp-analyzer ${SOURCE_FILES})

    target_include_directories(mediatranscoding_aidl_interface-cpp-analyzer
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/module/libmediatranscoding
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/module/libmediatranscoding/mediatranscoding_aidl_interface-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/module/libmediatranscoding/mediatranscoding_aidl_interface-cpp-analyzer-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(mediatranscoding_aidl_interface-cpp-analyzer PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(mediatranscoding_aidl_interface-cpp-analyzer
            aidl-analyzer-main
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libutils
            mediatranscoding_aidl_interface-cpp
    )
endif ()
#################### mediatranscoding_aidl_interface-cpp-analyzer ####################

#################### libstatslog_media ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/module/libmediatranscoding/statslog_media.cpp/gen/statslog_media.cpp
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
    add_library(libstatslog_media SHARED ${SOURCE_FILES})

    target_include_directories(libstatslog_media
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/module/libmediatranscoding
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/module/libmediatranscoding/statslog_media.cpp/gen
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/module/libmediatranscoding/statslog_media.h/gen
    )

    # Add c/c++ flags
    target_compile_options(libstatslog_media PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libstatslog_media
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            liblog
            libm
            libstatssocket
            libutils
    )
endif ()
#################### libstatslog_media ####################
