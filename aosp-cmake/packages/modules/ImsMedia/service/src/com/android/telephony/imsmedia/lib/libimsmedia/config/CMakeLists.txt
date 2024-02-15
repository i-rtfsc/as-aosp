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

project(libimsmedia_config)

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

#################### libimsmedia_config ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/config/src/AmrParams.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/config/src/AudioConfig.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/config/src/CallQuality.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/config/src/EvsParams.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/config/src/MediaQualityStatus.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/config/src/MediaQualityThreshold.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/config/src/RtcpConfig.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/config/src/RtpConfig.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/config/src/RtpHeaderExtension.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/config/src/TextConfig.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/config/src/VideoConfig.cpp
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
    add_library(libimsmedia_config SHARED ${SOURCE_FILES})

    target_include_directories(libimsmedia_config
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/config
            ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/config/include
    )

    # Add c/c++ flags
    target_compile_options(libimsmedia_config PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(libimsmedia_config
            libbinder
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libimsmedia_config_headers
            liblog
            libm
            libutils
    )
endif ()
#################### libimsmedia_config ####################
