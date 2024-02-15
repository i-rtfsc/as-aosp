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

project(libmegaaudio_jni)

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

#################### libmegaaudio_jni ####################
set(FILES
        ${ANDROID_ROOT}/cts/apps/CtsVerifier/jni/megaaudio/common/Globals.cpp
        ${ANDROID_ROOT}/cts/apps/CtsVerifier/jni/megaaudio/common/OboeStream.cpp
        ${ANDROID_ROOT}/cts/apps/CtsVerifier/jni/megaaudio/common/StreamBase.cpp
        ${ANDROID_ROOT}/cts/apps/CtsVerifier/jni/megaaudio/player/BlipAudioSource.cpp
        ${ANDROID_ROOT}/cts/apps/CtsVerifier/jni/megaaudio/player/JavaSourceProxy.cpp
        ${ANDROID_ROOT}/cts/apps/CtsVerifier/jni/megaaudio/player/NativeAudioSource.cpp
        ${ANDROID_ROOT}/cts/apps/CtsVerifier/jni/megaaudio/player/NoiseAudioSource.cpp
        ${ANDROID_ROOT}/cts/apps/CtsVerifier/jni/megaaudio/player/OboePlayer.cpp
        ${ANDROID_ROOT}/cts/apps/CtsVerifier/jni/megaaudio/player/SinAudioSource.cpp
        ${ANDROID_ROOT}/cts/apps/CtsVerifier/jni/megaaudio/player/SilenceAudioSource.cpp
        ${ANDROID_ROOT}/cts/apps/CtsVerifier/jni/megaaudio/player/WaveTableSource.cpp
        ${ANDROID_ROOT}/cts/apps/CtsVerifier/jni/megaaudio/recorder/AppCallbackAudioSink.cpp
        ${ANDROID_ROOT}/cts/apps/CtsVerifier/jni/megaaudio/recorder/DefaultAudioSink.cpp
        ${ANDROID_ROOT}/cts/apps/CtsVerifier/jni/megaaudio/recorder/NativeAudioSink.cpp
        ${ANDROID_ROOT}/cts/apps/CtsVerifier/jni/megaaudio/recorder/OboeRecorder.cpp
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
    add_library(libmegaaudio_jni SHARED ${SOURCE_FILES})

    target_include_directories(libmegaaudio_jni
            PUBLIC
            ${ANDROID_ROOT}/cts/apps/CtsVerifier/jni/megaaudio/common
            ${ANDROID_ROOT}/cts/apps/CtsVerifier/jni/megaaudio
    )

    # Add c/c++ flags
    target_compile_options(libmegaaudio_jni PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(libmegaaudio_jni
            jni_headers
            libOpenSLES
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libgtest
            libgtest_main
            liblog
            libm
            oboe
    )
endif ()
#################### libmegaaudio_jni ####################
