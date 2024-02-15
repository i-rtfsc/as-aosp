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

project(canhalctrl)

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

#################### canhalsend ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/automotive/can/1.0/tools/canhalsend.cpp
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
    add_executable(canhalsend ${SOURCE_FILES})

    target_include_directories(canhalsend
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/automotive/can/1.0/tools
            ${ANDROID_ROOT}/hardware/interfaces/automotive/can/1.0/tools/libcanhaltools/include
    )

    # Add c/c++ flags
    target_compile_options(canhalsend PRIVATE
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(canhalsend
            android.hardware.automotive.can-1.0
            android.hardware.automotive.can-libcanhaltools
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libhidlbase
            libm
            libutils
    )
endif ()
#################### canhalsend ####################

#################### canhaldump ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/automotive/can/1.0/tools/canhaldump.cpp
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
    add_executable(canhaldump ${SOURCE_FILES})

    target_include_directories(canhaldump
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/automotive/can/1.0/tools
            ${ANDROID_ROOT}/hardware/interfaces/automotive/can/1.0/tools/libcanhaltools/include
    )

    # Add c/c++ flags
    target_compile_options(canhaldump PRIVATE
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(canhaldump
            android.hardware.automotive.can-1.0
            android.hardware.automotive.can-hidl-utils-lib
            android.hardware.automotive.can-libcanhaltools
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libhidlbase
            libm
            libutils
    )
endif ()
#################### canhaldump ####################

#################### canhalctrl ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/automotive/can/1.0/tools/canhalctrl.cpp
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
    add_executable(canhalctrl ${SOURCE_FILES})

    target_include_directories(canhalctrl
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/automotive/can/1.0/tools
            ${ANDROID_ROOT}/hardware/interfaces/automotive/can/1.0/tools/libcanhaltools/include
    )

    # Add c/c++ flags
    target_compile_options(canhalctrl PRIVATE
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(canhalctrl
            android.hardware.automotive.can-1.0
            android.hardware.automotive.can-hidl-utils-lib
            android.hardware.automotive.can-libcanhaltools
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libhidlbase
            libm
            libutils
    )
endif ()
#################### canhalctrl ####################
