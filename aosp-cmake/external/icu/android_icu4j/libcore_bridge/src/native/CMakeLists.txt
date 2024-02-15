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

project(libicu_jni)

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

#################### libicu_jni ####################
set(FILES
        ${ANDROID_ROOT}/external/icu/android_icu4j/libcore_bridge/src/native/IcuUtilities.cpp
        ${ANDROID_ROOT}/external/icu/android_icu4j/libcore_bridge/src/native/JniConstants.cpp
        ${ANDROID_ROOT}/external/icu/android_icu4j/libcore_bridge/src/native/MatcherState.cpp
        ${ANDROID_ROOT}/external/icu/android_icu4j/libcore_bridge/src/native/Register.cpp
        ${ANDROID_ROOT}/external/icu/android_icu4j/libcore_bridge/src/native/com_android_i18n_timezone_internal_Memory.cpp
        ${ANDROID_ROOT}/external/icu/android_icu4j/libcore_bridge/src/native/com_android_i18n_util_Log.cpp
        ${ANDROID_ROOT}/external/icu/android_icu4j/libcore_bridge/src/native/com_android_icu_charset_NativeConverter.cpp
        ${ANDROID_ROOT}/external/icu/android_icu4j/libcore_bridge/src/native/com_android_icu_text_TimeZoneNamesNative.cpp
        ${ANDROID_ROOT}/external/icu/android_icu4j/libcore_bridge/src/native/com_android_icu_util_CaseMapperNative.cpp
        ${ANDROID_ROOT}/external/icu/android_icu4j/libcore_bridge/src/native/com_android_icu_util_ICU4CMetadata.cpp
        ${ANDROID_ROOT}/external/icu/android_icu4j/libcore_bridge/src/native/com_android_icu_util_LocaleNative.cpp
        ${ANDROID_ROOT}/external/icu/android_icu4j/libcore_bridge/src/native/com_android_icu_util_regex_MatcherNative.cpp
        ${ANDROID_ROOT}/external/icu/android_icu4j/libcore_bridge/src/native/com_android_icu_util_regex_PatternNative.cpp
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
    add_library(libicu_jni SHARED ${SOURCE_FILES})

    target_include_directories(libicu_jni
            PUBLIC
            ${ANDROID_ROOT}/external/icu/android_icu4j/libcore_bridge/src/native
    )

    # Add c/c++ flags
    target_compile_options(libicu_jni PRIVATE
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(libicu_jni
            jni_headers
            libbase
            libc++
            libicui18n
            libicuuc
            liblog
            libnativehelper
    )
endif ()
#################### libicu_jni ####################
