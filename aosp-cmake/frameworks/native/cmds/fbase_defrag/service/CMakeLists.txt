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

project(libfbase_defrag_service)

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

#################### libfbase_defrag_service ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/cmds/fbase_defrag/service/service.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/fbase_defrag/service/FbaseDefragNativeService.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/cmds/fbase_defrag/service/libfbase_defrag_service/android_${OUT_ARCH_CPU}_static/gen/aidl/vendor/meizu/fbase_defrag/IFbaseDefrag.cpp
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
    add_library(libfbase_defrag_service SHARED ${SOURCE_FILES})

    target_include_directories(libfbase_defrag_service
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/cmds/fbase_defrag/service
            ${OUT_INTERMEDIATES_ROOT}/frameworks/native/cmds/fbase_defrag/service/libfbase_defrag_service/android_${OUT_ARCH_CPU}_static/gen/aidl
            ${OUT_INTERMEDIATES_ROOT}/frameworks/native/cmds/fbase_defrag/service/libfbase_defrag_binder/android_${OUT_ARCH_CPU}_static/gen/aidl
    )

    # Add c/c++ flags
    target_compile_options(libfbase_defrag_service PRIVATE
            -Werror
    )

    target_link_libraries(libfbase_defrag_service
            libbase
            libbinder
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libfbase_defrag
            libfbase_defrag_binder
            liblog
            libm
            libutils
    )
endif ()
#################### libfbase_defrag_service ####################

#################### libfbase_defrag_binder ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/cmds/fbase_defrag/service/libfbase_defrag_binder/android_${OUT_ARCH_CPU}_static/gen/aidl/vendor/meizu/fbase_defrag/IFbaseDefrag.cpp
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
    add_library(libfbase_defrag_binder SHARED ${SOURCE_FILES})

    target_include_directories(libfbase_defrag_binder
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/cmds/fbase_defrag/service
            ${OUT_INTERMEDIATES_ROOT}/frameworks/native/cmds/fbase_defrag/service/libfbase_defrag_binder/android_${OUT_ARCH_CPU}_static/gen/aidl
    )

    # Add c/c++ flags
    target_compile_options(libfbase_defrag_binder PRIVATE
            -Werror
    )

    target_link_libraries(libfbase_defrag_binder
            libbinder
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libincremental_aidl-cpp
            libm
            libutils
    )
endif ()
#################### libfbase_defrag_binder ####################
