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

project(libcpuinfo)

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

#################### libcpuinfo ####################
set(FILES
        ${ANDROID_ROOT}/external/cpuinfo/src/init.c
        ${ANDROID_ROOT}/external/cpuinfo/src/api.c
        ${ANDROID_ROOT}/external/cpuinfo/src/cache.c
        ${ANDROID_ROOT}/external/cpuinfo/src/linux/processors.c
        ${ANDROID_ROOT}/external/cpuinfo/src/linux/smallfile.c
        ${ANDROID_ROOT}/external/cpuinfo/src/linux/multiline.c
        ${ANDROID_ROOT}/external/cpuinfo/src/linux/cpulist.c
        ${ANDROID_ROOT}/external/cpuinfo/src/arm/uarch.c
        ${ANDROID_ROOT}/external/cpuinfo/src/arm/cache.c
        ${ANDROID_ROOT}/external/cpuinfo/src/arm/linux/init.c
        ${ANDROID_ROOT}/external/cpuinfo/src/arm/linux/cpuinfo.c
        ${ANDROID_ROOT}/external/cpuinfo/src/arm/linux/clusters.c
        ${ANDROID_ROOT}/external/cpuinfo/src/arm/linux/chipset.c
        ${ANDROID_ROOT}/external/cpuinfo/src/arm/linux/midr.c
        ${ANDROID_ROOT}/external/cpuinfo/src/arm/linux/hwcap.c
        ${ANDROID_ROOT}/external/cpuinfo/src/arm/android/properties.c
        ${ANDROID_ROOT}/external/cpuinfo/src/arm/linux/aarch64-isa.c
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
    add_library(libcpuinfo SHARED ${SOURCE_FILES})

    target_include_directories(libcpuinfo
            PUBLIC
            ${ANDROID_ROOT}/external/cpuinfo/include
            ${ANDROID_ROOT}/external/cpuinfo/src
            ${ANDROID_ROOT}/external/cpuinfo
            ${ANDROID_ROOT}/external/cpuinfo/deps/clog/include
    )

    # Add c/c++ flags
    target_compile_options(libcpuinfo PRIVATE
            -Werror
            -Wno-unused-function
            -Wno-unused-parameter
    )

    target_link_libraries(libcpuinfo
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclog
            libdl
            libm
    )
endif ()
#################### libcpuinfo ####################
