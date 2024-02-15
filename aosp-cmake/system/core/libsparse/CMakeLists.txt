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

project(simg2img)

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

#################### img2simg ####################
set(FILES
        ${ANDROID_ROOT}/system/core/libsparse/img2simg.cpp
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
    add_executable(img2simg ${SOURCE_FILES})

    target_include_directories(img2simg
            PUBLIC
            ${ANDROID_ROOT}/system/core/libsparse
            ${ANDROID_ROOT}/system/core/libsparse/include
    )

    # Add c/c++ flags
    target_compile_options(img2simg PRIVATE
            -Werror
    )

    target_link_libraries(img2simg
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
            libsparse
            libz
    )
endif ()
#################### img2simg ####################

#################### sparse_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/system/core/libsparse/sparse_fuzzer.cpp
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
    add_executable(sparse_fuzzer ${SOURCE_FILES})

    target_include_directories(sparse_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/system/core/libsparse
            ${ANDROID_ROOT}/system/core/libsparse/include
    )

    # Add c/c++ flags
    target_compile_options(sparse_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(sparse_fuzzer
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            liblog
            libm
            libsparse
            libz
    )
endif ()
#################### sparse_fuzzer ####################

#################### libsparse ####################
set(FILES
        ${ANDROID_ROOT}/system/core/libsparse/backed_block.cpp
        ${ANDROID_ROOT}/system/core/libsparse/output_file.cpp
        ${ANDROID_ROOT}/system/core/libsparse/sparse.cpp
        ${ANDROID_ROOT}/system/core/libsparse/sparse_crc32.cpp
        ${ANDROID_ROOT}/system/core/libsparse/sparse_err.cpp
        ${ANDROID_ROOT}/system/core/libsparse/sparse_read.cpp
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
    add_library(libsparse SHARED ${SOURCE_FILES})

    target_include_directories(libsparse
            PUBLIC
            ${ANDROID_ROOT}/system/core/libsparse/include
            ${ANDROID_ROOT}/system/core/libsparse
    )

    # Add c/c++ flags
    target_compile_options(libsparse PRIVATE
            -Werror
    )

    target_link_libraries(libsparse
            libbase
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libwinpthread
            libz
    )
endif ()
#################### libsparse ####################

#################### simg2img ####################
set(FILES
        ${ANDROID_ROOT}/system/core/libsparse/simg2img.cpp
        ${ANDROID_ROOT}/system/core/libsparse/sparse_crc32.cpp
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
    add_executable(simg2img ${SOURCE_FILES})

    target_include_directories(simg2img
            PUBLIC
            ${ANDROID_ROOT}/system/core/libsparse
            ${ANDROID_ROOT}/system/core/libsparse/include
    )

    # Add c/c++ flags
    target_compile_options(simg2img PRIVATE
            -Werror
    )

    target_link_libraries(simg2img
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
            libsparse
            libz
    )
endif ()
#################### simg2img ####################
