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

project(tv_tuner_resource_manager_aidl_interface-ndk)

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

#################### tv_tuner_resource_manager_aidl_interface-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/java/android/media/tv/tunerresourcemanager/tv_tuner_resource_manager_aidl_interface-ndk-source/gen/android/media/tv/tunerresourcemanager/CasSessionRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/java/android/media/tv/tunerresourcemanager/tv_tuner_resource_manager_aidl_interface-ndk-source/gen/android/media/tv/tunerresourcemanager/IResourcesReclaimListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/java/android/media/tv/tunerresourcemanager/tv_tuner_resource_manager_aidl_interface-ndk-source/gen/android/media/tv/tunerresourcemanager/ITunerResourceManager.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/java/android/media/tv/tunerresourcemanager/tv_tuner_resource_manager_aidl_interface-ndk-source/gen/android/media/tv/tunerresourcemanager/ResourceClientProfile.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/java/android/media/tv/tunerresourcemanager/tv_tuner_resource_manager_aidl_interface-ndk-source/gen/android/media/tv/tunerresourcemanager/TunerCiCamRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/java/android/media/tv/tunerresourcemanager/tv_tuner_resource_manager_aidl_interface-ndk-source/gen/android/media/tv/tunerresourcemanager/TunerDemuxInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/java/android/media/tv/tunerresourcemanager/tv_tuner_resource_manager_aidl_interface-ndk-source/gen/android/media/tv/tunerresourcemanager/TunerDemuxRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/java/android/media/tv/tunerresourcemanager/tv_tuner_resource_manager_aidl_interface-ndk-source/gen/android/media/tv/tunerresourcemanager/TunerDescramblerRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/java/android/media/tv/tunerresourcemanager/tv_tuner_resource_manager_aidl_interface-ndk-source/gen/android/media/tv/tunerresourcemanager/TunerFrontendInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/java/android/media/tv/tunerresourcemanager/tv_tuner_resource_manager_aidl_interface-ndk-source/gen/android/media/tv/tunerresourcemanager/TunerFrontendRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/java/android/media/tv/tunerresourcemanager/tv_tuner_resource_manager_aidl_interface-ndk-source/gen/android/media/tv/tunerresourcemanager/TunerLnbRequest.cpp
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
    add_library(tv_tuner_resource_manager_aidl_interface-ndk SHARED ${SOURCE_FILES})

    target_include_directories(tv_tuner_resource_manager_aidl_interface-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/java/android/media/tv/tunerresourcemanager/tv_tuner_resource_manager_aidl_interface-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(tv_tuner_resource_manager_aidl_interface-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(tv_tuner_resource_manager_aidl_interface-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### tv_tuner_resource_manager_aidl_interface-ndk ####################

#################### tv_tuner_resource_manager_aidl_interface-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/java/android/media/tv/tunerresourcemanager/tv_tuner_resource_manager_aidl_interface-ndk_platform-source/gen/android/media/tv/tunerresourcemanager/CasSessionRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/java/android/media/tv/tunerresourcemanager/tv_tuner_resource_manager_aidl_interface-ndk_platform-source/gen/android/media/tv/tunerresourcemanager/IResourcesReclaimListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/java/android/media/tv/tunerresourcemanager/tv_tuner_resource_manager_aidl_interface-ndk_platform-source/gen/android/media/tv/tunerresourcemanager/ITunerResourceManager.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/java/android/media/tv/tunerresourcemanager/tv_tuner_resource_manager_aidl_interface-ndk_platform-source/gen/android/media/tv/tunerresourcemanager/ResourceClientProfile.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/java/android/media/tv/tunerresourcemanager/tv_tuner_resource_manager_aidl_interface-ndk_platform-source/gen/android/media/tv/tunerresourcemanager/TunerCiCamRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/java/android/media/tv/tunerresourcemanager/tv_tuner_resource_manager_aidl_interface-ndk_platform-source/gen/android/media/tv/tunerresourcemanager/TunerDemuxInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/java/android/media/tv/tunerresourcemanager/tv_tuner_resource_manager_aidl_interface-ndk_platform-source/gen/android/media/tv/tunerresourcemanager/TunerDemuxRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/java/android/media/tv/tunerresourcemanager/tv_tuner_resource_manager_aidl_interface-ndk_platform-source/gen/android/media/tv/tunerresourcemanager/TunerDescramblerRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/java/android/media/tv/tunerresourcemanager/tv_tuner_resource_manager_aidl_interface-ndk_platform-source/gen/android/media/tv/tunerresourcemanager/TunerFrontendInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/java/android/media/tv/tunerresourcemanager/tv_tuner_resource_manager_aidl_interface-ndk_platform-source/gen/android/media/tv/tunerresourcemanager/TunerFrontendRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/java/android/media/tv/tunerresourcemanager/tv_tuner_resource_manager_aidl_interface-ndk_platform-source/gen/android/media/tv/tunerresourcemanager/TunerLnbRequest.cpp
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
    add_library(tv_tuner_resource_manager_aidl_interface-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(tv_tuner_resource_manager_aidl_interface-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/java/android/media/tv/tunerresourcemanager/tv_tuner_resource_manager_aidl_interface-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(tv_tuner_resource_manager_aidl_interface-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(tv_tuner_resource_manager_aidl_interface-ndk_platform
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
#################### tv_tuner_resource_manager_aidl_interface-ndk_platform ####################
