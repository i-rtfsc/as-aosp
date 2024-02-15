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

project(gpuservice)

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

#################### libgpuservice ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/services/gpuservice/GpuService.cpp
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
    add_library(libgpuservice SHARED ${SOURCE_FILES})

    target_include_directories(libgpuservice
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/services/gpuservice/include
            ${ANDROID_ROOT}/frameworks/native/services/gpuservice
            ${ANDROID_ROOT}/frameworks/native/services/gpuservice/gpustats/include
            ${ANDROID_ROOT}/frameworks/native/services/gpuservice/gpumem/include
            ${ANDROID_ROOT}/frameworks/native/services/gpuservice/gpuwork/include
            ${ANDROID_ROOT}/frameworks/native/services/gpuservice/gpuwork/bpfprogs/include
            ${ANDROID_ROOT}/frameworks/native/services/gpuservice/tracing/include
    )

    # Add c/c++ flags
    target_compile_options(libgpuservice PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libgpuservice
            libbase
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libgfxstats
            libgpumem
            libgpumemtracer
            libgpuwork
            libgraphicsenv
            liblog
            libm
            libserviceutils
            libutils
            libvkjson
    )
endif ()
#################### libgpuservice ####################

#################### gpuservice ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/services/gpuservice/main_gpuservice.cpp
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
    add_executable(gpuservice ${SOURCE_FILES})

    target_include_directories(gpuservice
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/services/gpuservice
            ${ANDROID_ROOT}/frameworks/native/services/gpuservice/include
    )

    # Add c/c++ flags
    target_compile_options(gpuservice PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(gpuservice
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libgpuservice
            liblog
            libm
            libutils
    )
endif ()
#################### gpuservice ####################
