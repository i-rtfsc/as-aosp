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

project(android.automotive.computepipe.router@1.0)

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

#################### android.automotive.computepipe.router-1.0-impl ####################
set(FILES
        ${ANDROID_ROOT}/packages/services/Car/cpp/computepipe/router/1.0/PipeClient.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/computepipe/router/1.0/PipeRunner.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/computepipe/router/1.0/PipeRegistration.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/computepipe/router/1.0/PipeQuery.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/computepipe/router/1.0/RemoteState.cpp
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
    add_library(android.automotive.computepipe.router-1.0-impl SHARED ${SOURCE_FILES})

    target_include_directories(android.automotive.computepipe.router-1.0-impl
            PUBLIC
            ${ANDROID_ROOT}/packages/services/Car/cpp/computepipe/router/1.0/include
            ${ANDROID_ROOT}/packages/services/Car/cpp/computepipe/router/1.0
    )

    # Add c/c++ flags
    target_compile_options(android.automotive.computepipe.router-1.0-impl PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
    )

    target_link_libraries(android.automotive.computepipe.router-1.0-impl
            android.automotive.computepipe.registry-V2-ndk
            android.automotive.computepipe.runner-V2-ndk
            computepipe_router_headers
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            liblog
            libm
            libutils
    )
endif ()
#################### android.automotive.computepipe.router-1.0-impl ####################

#################### android.automotive.computepipe.router-1.0 ####################
set(FILES
        ${ANDROID_ROOT}/packages/services/Car/cpp/computepipe/router/1.0/service.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/computepipe/router/1.0/RouterSvc.cpp
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
    add_executable(android.automotive.computepipe.router-1.0 ${SOURCE_FILES})

    target_include_directories(android.automotive.computepipe.router-1.0
            PUBLIC
            ${ANDROID_ROOT}/packages/services/Car/cpp/computepipe/router/1.0
            ${ANDROID_ROOT}/packages/services/Car/cpp/computepipe/router/1.0/include
    )

    # Add c/c++ flags
    target_compile_options(android.automotive.computepipe.router-1.0 PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(android.automotive.computepipe.router-1.0
            android.automotive.computepipe.registry-V2-ndk
            android.automotive.computepipe.router-1.0-impl
            computepipe_router_headers
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            liblog
            libm
            libutils
    )
endif ()
#################### android.automotive.computepipe.router-1.0 ####################
