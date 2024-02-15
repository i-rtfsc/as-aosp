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

project(libclearkeycasplugin)

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

#################### libclearkeycasplugin ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/drm/mediacas/plugins/clearkey/ClearKeyCasPlugin.cpp
        ${ANDROID_ROOT}/frameworks/av/drm/mediacas/plugins/clearkey/ClearKeyFetcher.cpp
        ${ANDROID_ROOT}/frameworks/av/drm/mediacas/plugins/clearkey/ClearKeyLicenseFetcher.cpp
        ${ANDROID_ROOT}/frameworks/av/drm/mediacas/plugins/clearkey/ClearKeySessionLibrary.cpp
        ${ANDROID_ROOT}/frameworks/av/drm/mediacas/plugins/clearkey/ecm.cpp
        ${ANDROID_ROOT}/frameworks/av/drm/mediacas/plugins/clearkey/ecm_generator.cpp
        ${ANDROID_ROOT}/frameworks/av/drm/mediacas/plugins/clearkey/JsonAssetLoader.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/drm/mediacas/plugins/clearkey/libclearkeycasplugin/android_vendor.34_${OUT_ARCH_CPU}_shared/gen/proto/frameworks/av/drm/mediacas/plugins/clearkey/protos/license_protos.pb.cc
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
    add_library(libclearkeycasplugin SHARED ${SOURCE_FILES})

    target_include_directories(libclearkeycasplugin
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/drm/mediacas/plugins/clearkey
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/drm/mediacas/plugins/clearkey/libclearkeycasplugin/android_vendor.34_${OUT_ARCH_CPU}_shared/gen/proto/frameworks/av/drm/mediacas/plugins/clearkey
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/drm/mediacas/plugins/clearkey/libclearkeycasplugin/android_vendor.34_${OUT_ARCH_CPU}_shared/gen/proto
    )

    # Add c/c++ flags
    target_compile_options(libclearkeycasplugin PRIVATE
            -Werror
    )

    target_link_libraries(libclearkeycasplugin
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libcrypto
            libdl
            libjsmn
            liblog
            libm
            libprotobuf-cpp-full
            libprotobuf-cpp-lite
            libstagefright_foundation
            libutils
            media_plugin_headers
            libc
            libdl
            libm
    )
endif ()
#################### libclearkeycasplugin ####################
