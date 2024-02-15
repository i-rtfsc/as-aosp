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

project(rmi4update)

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

#################### rmi4update ####################
set(FILES
        ${ANDROID_ROOT}/external/rmi4utils/rmi4update/main.cpp
        ${ANDROID_ROOT}/external/rmi4utils/rmi4update/rmi4update.cpp
        ${ANDROID_ROOT}/external/rmi4utils/rmi4update/updateutil.cpp
        ${ANDROID_ROOT}/external/rmi4utils/rmi4update/firmware_image.cpp
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
    add_executable(rmi4update ${SOURCE_FILES})

    target_include_directories(rmi4update
            PUBLIC
            ${ANDROID_ROOT}/external/rmi4utils/rmi4update
    )

    # Add c/c++ flags
    target_compile_options(rmi4update PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-private-field
            -Wno-unused-variable
    )

    target_link_libraries(rmi4update
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            rmidevice
    )
endif ()
#################### rmi4update ####################
