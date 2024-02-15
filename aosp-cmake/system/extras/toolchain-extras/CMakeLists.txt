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

project(libprofile-clang-extras_continuous)

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

#################### libprofile-clang-extras_continuous ####################
set(FILES
        ${ANDROID_ROOT}/system/extras/toolchain-extras/profile-clang-extras.cpp
        ${ANDROID_ROOT}/system/extras/toolchain-extras/profile-clang-openat.cpp
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
    add_library(libprofile-clang-extras_continuous SHARED ${SOURCE_FILES})

    target_include_directories(libprofile-clang-extras_continuous
            PUBLIC
            ${ANDROID_ROOT}/system/extras/toolchain-extras
    )

    # Add c/c++ flags
    target_compile_options(libprofile-clang-extras_continuous PRIVATE
            -Werror
    )

    target_link_libraries(libprofile-clang-extras_continuous
            libc_headers
            libclang_rt.builtins
    )
endif ()
#################### libprofile-clang-extras_continuous ####################

#################### libprofile-clang-extras_cfi_support ####################
set(FILES
        ${ANDROID_ROOT}/system/extras/toolchain-extras/profile-clang-extras.cpp
        ${ANDROID_ROOT}/system/extras/toolchain-extras/profile-clang-openat.cpp
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
    add_library(libprofile-clang-extras_cfi_support SHARED ${SOURCE_FILES})

    target_include_directories(libprofile-clang-extras_cfi_support
            PUBLIC
            ${ANDROID_ROOT}/system/extras/toolchain-extras
    )

    # Add c/c++ flags
    target_compile_options(libprofile-clang-extras_cfi_support PRIVATE
            -Werror
    )

    target_link_libraries(libprofile-clang-extras_cfi_support
            libc_headers
            libclang_rt.builtins
    )
endif ()
#################### libprofile-clang-extras_cfi_support ####################

#################### libprofile-clang-extras_ndk ####################
set(FILES
        ${ANDROID_ROOT}/system/extras/toolchain-extras/profile-clang-extras.cpp
        ${ANDROID_ROOT}/system/extras/toolchain-extras/profile-clang-openat.cpp
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
    add_library(libprofile-clang-extras_ndk SHARED ${SOURCE_FILES})

    target_include_directories(libprofile-clang-extras_ndk
            PUBLIC
            ${ANDROID_ROOT}/system/extras/toolchain-extras
    )

    # Add c/c++ flags
    target_compile_options(libprofile-clang-extras_ndk PRIVATE
            -Werror
    )

    target_link_libraries(libprofile-clang-extras_ndk
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libprofile-clang-extras_ndk ####################

#################### libprofile-clang-extras_cfi_support_continuous ####################
set(FILES
        ${ANDROID_ROOT}/system/extras/toolchain-extras/profile-clang-extras.cpp
        ${ANDROID_ROOT}/system/extras/toolchain-extras/profile-clang-openat.cpp
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
    add_library(libprofile-clang-extras_cfi_support_continuous SHARED ${SOURCE_FILES})

    target_include_directories(libprofile-clang-extras_cfi_support_continuous
            PUBLIC
            ${ANDROID_ROOT}/system/extras/toolchain-extras
    )

    # Add c/c++ flags
    target_compile_options(libprofile-clang-extras_cfi_support_continuous PRIVATE
            -Werror
    )

    target_link_libraries(libprofile-clang-extras_cfi_support_continuous
            libc_headers
            libclang_rt.builtins
    )
endif ()
#################### libprofile-clang-extras_cfi_support_continuous ####################

#################### libprofile-extras_ndk ####################
set(FILES
        ${ANDROID_ROOT}/system/extras/toolchain-extras/profile-extras.cpp
        ${ANDROID_ROOT}/system/extras/toolchain-extras/profile-globals.c
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
    add_library(libprofile-extras_ndk SHARED ${SOURCE_FILES})

    target_include_directories(libprofile-extras_ndk
            PUBLIC
            ${ANDROID_ROOT}/system/extras/toolchain-extras
    )

    # Add c/c++ flags
    target_compile_options(libprofile-extras_ndk PRIVATE
            -Werror
    )

    target_link_libraries(libprofile-extras_ndk
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libprofile-extras_ndk ####################

#################### libprofile-extras ####################
set(FILES
        ${ANDROID_ROOT}/system/extras/toolchain-extras/profile-extras.cpp
        ${ANDROID_ROOT}/system/extras/toolchain-extras/profile-globals.c
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
    add_library(libprofile-extras SHARED ${SOURCE_FILES})

    target_include_directories(libprofile-extras
            PUBLIC
            ${ANDROID_ROOT}/system/extras/toolchain-extras
    )

    # Add c/c++ flags
    target_compile_options(libprofile-extras PRIVATE
            -Werror
    )

    target_link_libraries(libprofile-extras
            libc_headers
            libclang_rt.builtins
    )
endif ()
#################### libprofile-extras ####################

#################### libprofile-clang-extras ####################
set(FILES
        ${ANDROID_ROOT}/system/extras/toolchain-extras/profile-clang-extras.cpp
        ${ANDROID_ROOT}/system/extras/toolchain-extras/profile-clang-openat.cpp
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
    add_library(libprofile-clang-extras SHARED ${SOURCE_FILES})

    target_include_directories(libprofile-clang-extras
            PUBLIC
            ${ANDROID_ROOT}/system/extras/toolchain-extras
    )

    # Add c/c++ flags
    target_compile_options(libprofile-clang-extras PRIVATE
            -Werror
    )

    target_link_libraries(libprofile-clang-extras
            libc_headers
            libclang_rt.builtins
    )
endif ()
#################### libprofile-clang-extras ####################

#################### libprofile-clang-extras_ndk_continuous ####################
set(FILES
        ${ANDROID_ROOT}/system/extras/toolchain-extras/profile-clang-extras.cpp
        ${ANDROID_ROOT}/system/extras/toolchain-extras/profile-clang-openat.cpp
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
    add_library(libprofile-clang-extras_ndk_continuous SHARED ${SOURCE_FILES})

    target_include_directories(libprofile-clang-extras_ndk_continuous
            PUBLIC
            ${ANDROID_ROOT}/system/extras/toolchain-extras
    )

    # Add c/c++ flags
    target_compile_options(libprofile-clang-extras_ndk_continuous PRIVATE
            -Werror
    )

    target_link_libraries(libprofile-clang-extras_ndk_continuous
            libc_headers
            libclang_rt.builtins
    )
endif ()
#################### libprofile-clang-extras_ndk_continuous ####################
