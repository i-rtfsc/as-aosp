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

project(libsqlite3_android)

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

#################### libsqlite3_android ####################
set(FILES
        ${ANDROID_ROOT}/external/sqlite/android/PhoneNumberUtils.cpp
        ${ANDROID_ROOT}/external/sqlite/android/OldPhoneNumberUtils.cpp
        ${ANDROID_ROOT}/external/sqlite/android/sqlite3_android.cpp
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
    add_library(libsqlite3_android SHARED ${SOURCE_FILES})

    target_include_directories(libsqlite3_android
            PUBLIC
            ${ANDROID_ROOT}/external/sqlite/android
    )

    # Add c/c++ flags
    target_compile_options(libsqlite3_android PRIVATE
            -Wall
            -Werror
            -Wno-unused-function
            -Wno-unused-variable
    )

    target_link_libraries(libsqlite3_android
            libandroidicu
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libicui18n
            libicuuc
            liblog
            libm
    )
endif ()
#################### libsqlite3_android ####################

#################### libsqlite3_android_noicu ####################
set(FILES
        ${ANDROID_ROOT}/external/sqlite/android/PhoneNumberUtils.cpp
        ${ANDROID_ROOT}/external/sqlite/android/OldPhoneNumberUtils.cpp
        ${ANDROID_ROOT}/external/sqlite/android/sqlite3_android.cpp
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
    add_library(libsqlite3_android_noicu SHARED ${SOURCE_FILES})

    target_include_directories(libsqlite3_android_noicu
            PUBLIC
            ${ANDROID_ROOT}/external/sqlite/android
    )

    # Add c/c++ flags
    target_compile_options(libsqlite3_android_noicu PRIVATE
            -Wall
            -Werror
            -Wno-unused-function
            -Wno-unused-variable
    )

    target_link_libraries(libsqlite3_android_noicu
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            liblog
            libm
    )
endif ()
#################### libsqlite3_android_noicu ####################
