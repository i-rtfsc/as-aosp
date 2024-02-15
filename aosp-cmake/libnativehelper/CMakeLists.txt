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

project(libnativehelper_compat_libc++)

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

#################### libnativehelper_compat_libc++ ####################
set(FILES
        ${ANDROID_ROOT}/libnativehelper/ExpandableString.c
        ${ANDROID_ROOT}/libnativehelper/JNIHelp.c
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
    add_library(libnativehelper_compat_libc++ SHARED ${SOURCE_FILES})

    target_include_directories(libnativehelper_compat_libc++
            PUBLIC
            ${ANDROID_ROOT}/libnativehelper/header_only_include
            ${ANDROID_ROOT}/libnativehelper/include
            ${ANDROID_ROOT}/libnativehelper/include_platform_header_only
            ${ANDROID_ROOT}/libnativehelper
            ${ANDROID_ROOT}/libnativehelper/include_jni
    )

    # Add c/c++ flags
    target_compile_options(libnativehelper_compat_libc++ PRIVATE
            -Werror
    )

    target_link_libraries(libnativehelper_compat_libc++
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            liblog
            libm
    )
endif ()
#################### libnativehelper_compat_libc++ ####################

#################### libnativehelper ####################
set(FILES
        ${ANDROID_ROOT}/libnativehelper/DlHelp.c
        ${ANDROID_ROOT}/libnativehelper/ExpandableString.c
        ${ANDROID_ROOT}/libnativehelper/JNIHelp.c
        ${ANDROID_ROOT}/libnativehelper/JNIPlatformHelp.c
        ${ANDROID_ROOT}/libnativehelper/JniConstants.c
        ${ANDROID_ROOT}/libnativehelper/JniInvocation.c
        ${ANDROID_ROOT}/libnativehelper/file_descriptor_jni.c
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
    add_library(libnativehelper SHARED ${SOURCE_FILES})

    target_include_directories(libnativehelper
            PUBLIC
            ${ANDROID_ROOT}/libnativehelper/header_only_include
            ${ANDROID_ROOT}/libnativehelper/include
            ${ANDROID_ROOT}/libnativehelper/include_jni
            ${ANDROID_ROOT}/libnativehelper/include_platform
            ${ANDROID_ROOT}/libnativehelper/include_platform_header_only
            ${ANDROID_ROOT}/libnativehelper
    )

    # Add c/c++ flags
    target_compile_options(libnativehelper PRIVATE
            -Werror
    )

    target_link_libraries(libnativehelper
            libc
            libclang_rt.builtins
            libdl
            liblog
            libm
    )
endif ()
#################### libnativehelper ####################

#################### libnativehelper_lazy ####################
set(FILES
        ${ANDROID_ROOT}/libnativehelper/libnativehelper_lazy.c
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
    add_library(libnativehelper_lazy SHARED ${SOURCE_FILES})

    target_include_directories(libnativehelper_lazy
            PUBLIC
            ${ANDROID_ROOT}/libnativehelper/header_only_include
            ${ANDROID_ROOT}/libnativehelper/include
            ${ANDROID_ROOT}/libnativehelper/include_jni
            ${ANDROID_ROOT}/libnativehelper/include_platform
            ${ANDROID_ROOT}/libnativehelper/include_platform_header_only
            ${ANDROID_ROOT}/libnativehelper
    )

    # Add c/c++ flags
    target_compile_options(libnativehelper_lazy PRIVATE
            -Werror
    )

    target_link_libraries(libnativehelper_lazy
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            liblog
            libm
    )
endif ()
#################### libnativehelper_lazy ####################
