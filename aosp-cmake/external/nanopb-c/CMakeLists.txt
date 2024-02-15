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

project(libprotobuf-c-nano-32bit)

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

#################### libprotobuf-c-nano-32bit ####################
set(FILES
        ${ANDROID_ROOT}/external/nanopb-c/pb_common.c
        ${ANDROID_ROOT}/external/nanopb-c/pb_decode.c
        ${ANDROID_ROOT}/external/nanopb-c/pb_encode.c
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
    add_library(libprotobuf-c-nano-32bit SHARED ${SOURCE_FILES})

    target_include_directories(libprotobuf-c-nano-32bit
            PUBLIC
            ${ANDROID_ROOT}/external/nanopb-c
    )

    # Add c/c++ flags
    target_compile_options(libprotobuf-c-nano-32bit PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libprotobuf-c-nano-32bit
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libprotobuf-c-nano-32bit ####################

#################### libprotobuf-c-nano ####################
set(FILES
        ${ANDROID_ROOT}/external/nanopb-c/pb_common.c
        ${ANDROID_ROOT}/external/nanopb-c/pb_decode.c
        ${ANDROID_ROOT}/external/nanopb-c/pb_encode.c
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
    add_library(libprotobuf-c-nano SHARED ${SOURCE_FILES})

    target_include_directories(libprotobuf-c-nano
            PUBLIC
            ${ANDROID_ROOT}/external/nanopb-c
    )

    # Add c/c++ flags
    target_compile_options(libprotobuf-c-nano PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libprotobuf-c-nano
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libprotobuf-c-nano ####################

#################### libprotobuf-c-nano-16bit ####################
set(FILES
        ${ANDROID_ROOT}/external/nanopb-c/pb_common.c
        ${ANDROID_ROOT}/external/nanopb-c/pb_decode.c
        ${ANDROID_ROOT}/external/nanopb-c/pb_encode.c
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
    add_library(libprotobuf-c-nano-16bit SHARED ${SOURCE_FILES})

    target_include_directories(libprotobuf-c-nano-16bit
            PUBLIC
            ${ANDROID_ROOT}/external/nanopb-c
    )

    # Add c/c++ flags
    target_compile_options(libprotobuf-c-nano-16bit PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libprotobuf-c-nano-16bit
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libprotobuf-c-nano-16bit ####################

#################### libprotobuf-c-nano-enable_malloc-32bit ####################
set(FILES
        ${ANDROID_ROOT}/external/nanopb-c/pb_common.c
        ${ANDROID_ROOT}/external/nanopb-c/pb_decode.c
        ${ANDROID_ROOT}/external/nanopb-c/pb_encode.c
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
    add_library(libprotobuf-c-nano-enable_malloc-32bit SHARED ${SOURCE_FILES})

    target_include_directories(libprotobuf-c-nano-enable_malloc-32bit
            PUBLIC
            ${ANDROID_ROOT}/external/nanopb-c
    )

    # Add c/c++ flags
    target_compile_options(libprotobuf-c-nano-enable_malloc-32bit PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libprotobuf-c-nano-enable_malloc-32bit
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libprotobuf-c-nano-enable_malloc-32bit ####################

#################### libprotobuf-c-nano-enable_malloc ####################
set(FILES
        ${ANDROID_ROOT}/external/nanopb-c/pb_common.c
        ${ANDROID_ROOT}/external/nanopb-c/pb_decode.c
        ${ANDROID_ROOT}/external/nanopb-c/pb_encode.c
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
    add_library(libprotobuf-c-nano-enable_malloc SHARED ${SOURCE_FILES})

    target_include_directories(libprotobuf-c-nano-enable_malloc
            PUBLIC
            ${ANDROID_ROOT}/external/nanopb-c
    )

    # Add c/c++ flags
    target_compile_options(libprotobuf-c-nano-enable_malloc PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libprotobuf-c-nano-enable_malloc
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libprotobuf-c-nano-enable_malloc ####################

#################### libprotobuf-c-nano-enable_malloc-16bit ####################
set(FILES
        ${ANDROID_ROOT}/external/nanopb-c/pb_common.c
        ${ANDROID_ROOT}/external/nanopb-c/pb_decode.c
        ${ANDROID_ROOT}/external/nanopb-c/pb_encode.c
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
    add_library(libprotobuf-c-nano-enable_malloc-16bit SHARED ${SOURCE_FILES})

    target_include_directories(libprotobuf-c-nano-enable_malloc-16bit
            PUBLIC
            ${ANDROID_ROOT}/external/nanopb-c
    )

    # Add c/c++ flags
    target_compile_options(libprotobuf-c-nano-enable_malloc-16bit PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libprotobuf-c-nano-enable_malloc-16bit
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libprotobuf-c-nano-enable_malloc-16bit ####################
