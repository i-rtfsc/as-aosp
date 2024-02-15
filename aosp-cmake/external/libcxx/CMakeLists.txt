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

project(libc++experimental)

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

#################### libc++experimental ####################
set(FILES
        ${ANDROID_ROOT}/external/libcxx/src/experimental/memory_resource.cpp
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
    add_library(libc++experimental SHARED ${SOURCE_FILES})

    target_include_directories(libc++experimental
            PUBLIC
            ${ANDROID_ROOT}/external/libcxx/include
            ${ANDROID_ROOT}/external/libcxx
    )

    # Add c/c++ flags
    target_compile_options(libc++experimental PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libc++experimental
            libc
            libclang_rt.builtins
            libdl
            libm
            libwinpthread
    )
endif ()
#################### libc++experimental ####################

#################### libc++_static ####################
set(FILES
        ${ANDROID_ROOT}/external/libcxx/src/algorithm.cpp
        ${ANDROID_ROOT}/external/libcxx/src/any.cpp
        ${ANDROID_ROOT}/external/libcxx/src/bind.cpp
        ${ANDROID_ROOT}/external/libcxx/src/charconv.cpp
        ${ANDROID_ROOT}/external/libcxx/src/chrono.cpp
        ${ANDROID_ROOT}/external/libcxx/src/condition_variable.cpp
        ${ANDROID_ROOT}/external/libcxx/src/debug.cpp
        ${ANDROID_ROOT}/external/libcxx/src/exception.cpp
        ${ANDROID_ROOT}/external/libcxx/src/future.cpp
        ${ANDROID_ROOT}/external/libcxx/src/hash.cpp
        ${ANDROID_ROOT}/external/libcxx/src/ios.cpp
        ${ANDROID_ROOT}/external/libcxx/src/iostream.cpp
        ${ANDROID_ROOT}/external/libcxx/src/locale.cpp
        ${ANDROID_ROOT}/external/libcxx/src/memory.cpp
        ${ANDROID_ROOT}/external/libcxx/src/mutex.cpp
        ${ANDROID_ROOT}/external/libcxx/src/new.cpp
        ${ANDROID_ROOT}/external/libcxx/src/optional.cpp
        ${ANDROID_ROOT}/external/libcxx/src/random.cpp
        ${ANDROID_ROOT}/external/libcxx/src/regex.cpp
        ${ANDROID_ROOT}/external/libcxx/src/shared_mutex.cpp
        ${ANDROID_ROOT}/external/libcxx/src/stdexcept.cpp
        ${ANDROID_ROOT}/external/libcxx/src/string.cpp
        ${ANDROID_ROOT}/external/libcxx/src/strstream.cpp
        ${ANDROID_ROOT}/external/libcxx/src/system_error.cpp
        ${ANDROID_ROOT}/external/libcxx/src/thread.cpp
        ${ANDROID_ROOT}/external/libcxx/src/typeinfo.cpp
        ${ANDROID_ROOT}/external/libcxx/src/utility.cpp
        ${ANDROID_ROOT}/external/libcxx/src/valarray.cpp
        ${ANDROID_ROOT}/external/libcxx/src/variant.cpp
        ${ANDROID_ROOT}/external/libcxx/src/vector.cpp
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
    add_library(libc++_static SHARED ${SOURCE_FILES})

    target_include_directories(libc++_static
            PUBLIC
            ${ANDROID_ROOT}/external/libcxx/include
            ${ANDROID_ROOT}/external/libcxx
            ${ANDROID_ROOT}/external/libcxxabi/include
    )

    # Add c/c++ flags
    target_compile_options(libc++_static PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libc++_static
            libc
            libc++abi
            libclang_rt.builtins
            libdl
            libm
            libwinpthread
    )
endif ()
#################### libc++_static ####################

#################### libc++fs ####################
set(FILES
        ${ANDROID_ROOT}/external/libcxx/src/filesystem/directory_iterator.cpp
        ${ANDROID_ROOT}/external/libcxx/src/filesystem/operations.cpp
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
    add_library(libc++fs SHARED ${SOURCE_FILES})

    target_include_directories(libc++fs
            PUBLIC
            ${ANDROID_ROOT}/external/libcxx/include
            ${ANDROID_ROOT}/external/libcxx
    )

    # Add c/c++ flags
    target_compile_options(libc++fs PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libc++fs
            libc
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libc++fs ####################
