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

project(mediametrics)

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

#################### mediametrics ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/services/mediametrics/main_mediametrics.cpp
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
    add_executable(mediametrics ${SOURCE_FILES})

    target_include_directories(mediametrics
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/services/mediametrics
            ${ANDROID_ROOT}/frameworks/av/services/mediametrics/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediametrics/stats_media_metrics.h/gen
    )

    # Add c/c++ flags
    target_compile_options(mediametrics PRIVATE
            -Wall
            -Werror
            -Werror=implicit-fallthrough
            -Werror=sometimes-uninitialized
            -Werror=conditional-uninitialized
            -Wextra
    )

    target_link_libraries(mediametrics
            libaudioutils_headers
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            liblog
            libm
            libmediametrics_headers
            libmediametricsservice
            libmediautils
            libutils
            mediametricsservice-aidl-cpp
            packagemanager_aidl-cpp
    )
endif ()
#################### mediametrics ####################

#################### libmediametricsservice ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/services/mediametrics/AudioAnalytics.cpp
        ${ANDROID_ROOT}/frameworks/av/services/mediametrics/AudioPowerUsage.cpp
        ${ANDROID_ROOT}/frameworks/av/services/mediametrics/AudioTypes.cpp
        ${ANDROID_ROOT}/frameworks/av/services/mediametrics/cleaner.cpp
        ${ANDROID_ROOT}/frameworks/av/services/mediametrics/iface_statsd.cpp
        ${ANDROID_ROOT}/frameworks/av/services/mediametrics/MediaDrmStatsdHelper.cpp
        ${ANDROID_ROOT}/frameworks/av/services/mediametrics/MediaMetricsService.cpp
        ${ANDROID_ROOT}/frameworks/av/services/mediametrics/statsd_audiopolicy.cpp
        ${ANDROID_ROOT}/frameworks/av/services/mediametrics/statsd_audiorecord.cpp
        ${ANDROID_ROOT}/frameworks/av/services/mediametrics/statsd_audiothread.cpp
        ${ANDROID_ROOT}/frameworks/av/services/mediametrics/statsd_audiotrack.cpp
        ${ANDROID_ROOT}/frameworks/av/services/mediametrics/statsd_codec.cpp
        ${ANDROID_ROOT}/frameworks/av/services/mediametrics/statsd_drm.cpp
        ${ANDROID_ROOT}/frameworks/av/services/mediametrics/statsd_extractor.cpp
        ${ANDROID_ROOT}/frameworks/av/services/mediametrics/statsd_mediaparser.cpp
        ${ANDROID_ROOT}/frameworks/av/services/mediametrics/statsd_nuplayer.cpp
        ${ANDROID_ROOT}/frameworks/av/services/mediametrics/statsd_recorder.cpp
        ${ANDROID_ROOT}/frameworks/av/services/mediametrics/StringUtils.cpp
        ${ANDROID_ROOT}/frameworks/av/services/mediametrics/ValidateId.cpp
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
    add_library(libmediametricsservice SHARED ${SOURCE_FILES})

    target_include_directories(libmediametricsservice
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/services/mediametrics/include
            ${ANDROID_ROOT}/frameworks/av/services/mediametrics/include/mediametricsservice
            ${ANDROID_ROOT}/frameworks/av/services/mediametrics
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediametrics/stats_media_metrics.h/gen
    )

    # Add c/c++ flags
    target_compile_options(libmediametricsservice PRIVATE
            -Wall
            -Werror
            -Werror=implicit-fallthrough
            -Werror=sometimes-uninitialized
            -Werror=conditional-uninitialized
            -Wextra
    )

    target_link_libraries(libmediametricsservice
            libaaudio_headers
            libaudioutils_headers
            libbase
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            liblog
            libm
            libmedia_helper
            libmediametrics
            libmediautils
            libmemunreachable
            libplatformprotos
            libprotobuf-cpp-lite
            libstagefright_foundation
            libstats_media_metrics
            libstatspull
            libstatssocket
            libutils
            mediametricsservice-aidl-cpp
            packagemanager_aidl-cpp
    )
endif ()
#################### libmediametricsservice ####################

#################### libstats_media_metrics ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediametrics/stats_media_metrics.cpp/gen/stats_media_metrics.cpp
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
    add_library(libstats_media_metrics SHARED ${SOURCE_FILES})

    target_include_directories(libstats_media_metrics
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/services/mediametrics
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediametrics/stats_media_metrics.cpp/gen
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediametrics/stats_media_metrics.h/gen
    )

    target_link_libraries(libstats_media_metrics
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libm
            libstatspull
            libstatssocket
    )
endif ()
#################### libstats_media_metrics ####################
