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

project(libtiff_directory)

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

#################### libtiff_directory ####################
set(FILES
        ${ANDROID_ROOT}/external/piex/src/tiff_directory/tiff_directory.cc
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
    add_library(libtiff_directory SHARED ${SOURCE_FILES})

    target_include_directories(libtiff_directory
            PUBLIC
            ${ANDROID_ROOT}/external/piex
    )

    # Add c/c++ flags
    target_compile_options(libtiff_directory PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libtiff_directory
            libbinary_parse
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libwinpthread
    )
endif ()
#################### libtiff_directory ####################

#################### libbinary_parse ####################
set(FILES
        ${ANDROID_ROOT}/external/piex/src/binary_parse/cached_paged_byte_array.cc
        ${ANDROID_ROOT}/external/piex/src/binary_parse/range_checked_byte_ptr.cc
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
    add_library(libbinary_parse SHARED ${SOURCE_FILES})

    target_include_directories(libbinary_parse
            PUBLIC
            ${ANDROID_ROOT}/external/piex
    )

    # Add c/c++ flags
    target_compile_options(libbinary_parse PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libbinary_parse
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libwinpthread
    )
endif ()
#################### libbinary_parse ####################

#################### libimage_type_recognition ####################
set(FILES
        ${ANDROID_ROOT}/external/piex/src/image_type_recognition/image_type_recognition_lite.cc
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
    add_library(libimage_type_recognition SHARED ${SOURCE_FILES})

    target_include_directories(libimage_type_recognition
            PUBLIC
            ${ANDROID_ROOT}/external/piex
    )

    # Add c/c++ flags
    target_compile_options(libimage_type_recognition PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libimage_type_recognition
            libbinary_parse
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libwinpthread
    )
endif ()
#################### libimage_type_recognition ####################

#################### libpiex ####################
set(FILES
        ${ANDROID_ROOT}/external/piex/src/piex.cc
        ${ANDROID_ROOT}/external/piex/src/piex_cr3.cc
        ${ANDROID_ROOT}/external/piex/src/tiff_parser.cc
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
    add_library(libpiex SHARED ${SOURCE_FILES})

    target_include_directories(libpiex
            PUBLIC
            ${ANDROID_ROOT}/external/piex
    )

    # Add c/c++ flags
    target_compile_options(libpiex PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libpiex
            libbinary_parse
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libimage_type_recognition
            libm
            libtiff_directory
            libwinpthread
    )
endif ()
#################### libpiex ####################
