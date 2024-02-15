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

project(libaudiopolicycomponents)

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

#################### libaudiopolicycomponents ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/services/audiopolicy/common/managerdefinitions/src/AudioCollections.cpp
        ${ANDROID_ROOT}/frameworks/av/services/audiopolicy/common/managerdefinitions/src/AudioInputDescriptor.cpp
        ${ANDROID_ROOT}/frameworks/av/services/audiopolicy/common/managerdefinitions/src/AudioOutputDescriptor.cpp
        ${ANDROID_ROOT}/frameworks/av/services/audiopolicy/common/managerdefinitions/src/AudioPatch.cpp
        ${ANDROID_ROOT}/frameworks/av/services/audiopolicy/common/managerdefinitions/src/AudioPolicyConfig.cpp
        ${ANDROID_ROOT}/frameworks/av/services/audiopolicy/common/managerdefinitions/src/AudioPolicyMix.cpp
        ${ANDROID_ROOT}/frameworks/av/services/audiopolicy/common/managerdefinitions/src/AudioProfileVectorHelper.cpp
        ${ANDROID_ROOT}/frameworks/av/services/audiopolicy/common/managerdefinitions/src/AudioRoute.cpp
        ${ANDROID_ROOT}/frameworks/av/services/audiopolicy/common/managerdefinitions/src/ClientDescriptor.cpp
        ${ANDROID_ROOT}/frameworks/av/services/audiopolicy/common/managerdefinitions/src/DeviceDescriptor.cpp
        ${ANDROID_ROOT}/frameworks/av/services/audiopolicy/common/managerdefinitions/src/EffectDescriptor.cpp
        ${ANDROID_ROOT}/frameworks/av/services/audiopolicy/common/managerdefinitions/src/HwModule.cpp
        ${ANDROID_ROOT}/frameworks/av/services/audiopolicy/common/managerdefinitions/src/IOProfile.cpp
        ${ANDROID_ROOT}/frameworks/av/services/audiopolicy/common/managerdefinitions/src/PolicyAudioPort.cpp
        ${ANDROID_ROOT}/frameworks/av/services/audiopolicy/common/managerdefinitions/src/PreferredMixerAttributesInfo.cpp
        ${ANDROID_ROOT}/frameworks/av/services/audiopolicy/common/managerdefinitions/src/Serializer.cpp
        ${ANDROID_ROOT}/frameworks/av/services/audiopolicy/common/managerdefinitions/src/SoundTriggerSession.cpp
        ${ANDROID_ROOT}/frameworks/av/services/audiopolicy/common/managerdefinitions/src/TypeConverter.cpp
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
    add_library(libaudiopolicycomponents SHARED ${SOURCE_FILES})

    target_include_directories(libaudiopolicycomponents
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/services/audiopolicy/common/managerdefinitions/include
            ${ANDROID_ROOT}/frameworks/av/services/audiopolicy/common/managerdefinitions
    )

    # Add c/c++ flags
    target_compile_options(libaudiopolicycomponents PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libaudiopolicycomponents
            libaudiofoundation
            libaudiopolicycommon
            libaudiopolicymanager_interface_headers
            libaudioutils
            libbase
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libhidlbase
            liblog
            libm
            libmedia
            libmedia_helper
            libutils
            libxml2
    )
endif ()
#################### libaudiopolicycomponents ####################
