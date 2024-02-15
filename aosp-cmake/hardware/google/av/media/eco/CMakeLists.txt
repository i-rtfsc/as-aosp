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

project(libmedia_ecoservice)

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

#################### libmedia_ecoservice ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/av/media/eco/libmedia_ecoservice/android_${OUT_ARCH_CPU}_shared/gen/aidl/android/media/eco/IECOService.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/av/media/eco/libmedia_ecoservice/android_${OUT_ARCH_CPU}_shared/gen/aidl/android/media/eco/IECOSession.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/av/media/eco/libmedia_ecoservice/android_${OUT_ARCH_CPU}_shared/gen/aidl/android/media/eco/IECOServiceStatsProvider.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/av/media/eco/libmedia_ecoservice/android_${OUT_ARCH_CPU}_shared/gen/aidl/android/media/eco/IECOServiceInfoListener.cpp
        ${ANDROID_ROOT}/hardware/google/av/media/eco/ECOData.cpp
        ${ANDROID_ROOT}/hardware/google/av/media/eco/ECODebug.cpp
        ${ANDROID_ROOT}/hardware/google/av/media/eco/ECOService.cpp
        ${ANDROID_ROOT}/hardware/google/av/media/eco/ECOSession.cpp
        ${ANDROID_ROOT}/hardware/google/av/media/eco/ECOUtils.cpp
        ${ANDROID_ROOT}/hardware/google/av/media/eco/ECOServiceStatsProvider.cpp
        ${ANDROID_ROOT}/hardware/google/av/media/eco/ECOC2Utils.cpp
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
    add_library(libmedia_ecoservice SHARED ${SOURCE_FILES})

    target_include_directories(libmedia_ecoservice
            PUBLIC
            ${ANDROID_ROOT}/hardware/google/av/media/eco/include
            ${ANDROID_ROOT}/hardware/google/av/media/eco
            ${OUT_INTERMEDIATES_ROOT}/hardware/google/av/media/eco/libmedia_ecoservice/android_${OUT_ARCH_CPU}_shared/gen/aidl
    )

    # Add c/c++ flags
    target_compile_options(libmedia_ecoservice PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libmedia_ecoservice
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcodec2_hidl-1.0
            libcutils
            libdl
            liblog
            libm
            libutils
    )
endif ()
#################### libmedia_ecoservice ####################
