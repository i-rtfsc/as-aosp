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

project(update_engine_stable_client)

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

#################### libupdate_engine_stable-V1-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/update_engine/stable/libupdate_engine_stable-V1-ndk_platform-source/gen/android/os/IUpdateEngineStable.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/update_engine/stable/libupdate_engine_stable-V1-ndk_platform-source/gen/android/os/IUpdateEngineStableCallback.cpp
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
    add_library(libupdate_engine_stable-V1-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(libupdate_engine_stable-V1-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/system/update_engine/stable/libupdate_engine_stable-V1-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(libupdate_engine_stable-V1-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(libupdate_engine_stable-V1-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libupdate_engine_stable-V1-ndk_platform ####################

#################### libupdate_engine_stable-V1-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/update_engine/stable/libupdate_engine_stable-V1-ndk-source/gen/android/os/IUpdateEngineStable.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/update_engine/stable/libupdate_engine_stable-V1-ndk-source/gen/android/os/IUpdateEngineStableCallback.cpp
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
    add_library(libupdate_engine_stable-V1-ndk SHARED ${SOURCE_FILES})

    target_include_directories(libupdate_engine_stable-V1-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/system/update_engine/stable/libupdate_engine_stable-V1-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(libupdate_engine_stable-V1-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(libupdate_engine_stable-V1-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libupdate_engine_stable-V1-ndk ####################

#################### libupdate_engine_stable-V2-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/update_engine/stable/libupdate_engine_stable-V2-cpp-source/gen/android/os/IUpdateEngineStable.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/update_engine/stable/libupdate_engine_stable-V2-cpp-source/gen/android/os/IUpdateEngineStableCallback.cpp
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
    add_library(libupdate_engine_stable-V2-cpp SHARED ${SOURCE_FILES})

    target_include_directories(libupdate_engine_stable-V2-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/system/update_engine/stable/libupdate_engine_stable-V2-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(libupdate_engine_stable-V2-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(libupdate_engine_stable-V2-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libm
            libutils
    )
endif ()
#################### libupdate_engine_stable-V2-cpp ####################

#################### libupdate_engine_stable-V1-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/update_engine/stable/libupdate_engine_stable-V1-cpp-source/gen/android/os/IUpdateEngineStable.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/update_engine/stable/libupdate_engine_stable-V1-cpp-source/gen/android/os/IUpdateEngineStableCallback.cpp
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
    add_library(libupdate_engine_stable-V1-cpp SHARED ${SOURCE_FILES})

    target_include_directories(libupdate_engine_stable-V1-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/system/update_engine/stable/libupdate_engine_stable-V1-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(libupdate_engine_stable-V1-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(libupdate_engine_stable-V1-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libm
            libutils
    )
endif ()
#################### libupdate_engine_stable-V1-cpp ####################

#################### libupdate_engine_stable-cpp-analyzer ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/update_engine/stable/libupdate_engine_stable-cpp-analyzer-source/gen/android/os/IUpdateEngineStable.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/update_engine/stable/libupdate_engine_stable-cpp-analyzer-source/gen/android/os/IUpdateEngineStableCallback.cpp
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
    add_executable(libupdate_engine_stable-cpp-analyzer ${SOURCE_FILES})

    target_include_directories(libupdate_engine_stable-cpp-analyzer
            PUBLIC
            ${ANDROID_ROOT}/system/update_engine/stable
            ${OUT_INTERMEDIATES_ROOT}/system/update_engine/stable/libupdate_engine_stable-V2-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/system/update_engine/stable/libupdate_engine_stable-cpp-analyzer-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(libupdate_engine_stable-cpp-analyzer PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(libupdate_engine_stable-cpp-analyzer
            aidl-analyzer-main
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libupdate_engine_stable-V2-cpp
            libutils
    )
endif ()
#################### libupdate_engine_stable-cpp-analyzer ####################

#################### libupdate_engine_stable-V2-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/update_engine/stable/libupdate_engine_stable-V2-ndk-source/gen/android/os/IUpdateEngineStable.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/update_engine/stable/libupdate_engine_stable-V2-ndk-source/gen/android/os/IUpdateEngineStableCallback.cpp
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
    add_library(libupdate_engine_stable-V2-ndk SHARED ${SOURCE_FILES})

    target_include_directories(libupdate_engine_stable-V2-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/system/update_engine/stable/libupdate_engine_stable-V2-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(libupdate_engine_stable-V2-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(libupdate_engine_stable-V2-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libupdate_engine_stable-V2-ndk ####################

#################### update_engine_stable_client ####################
set(FILES
        ${ANDROID_ROOT}/system/update_engine/stable/update_engine_stable_client.cc
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
    add_executable(update_engine_stable_client ${SOURCE_FILES})

    target_include_directories(update_engine_stable_client
            PUBLIC
            ${ANDROID_ROOT}/system/update_engine/stable
            ${OUT_INTERMEDIATES_ROOT}/system/update_engine/stable/libupdate_engine_stable-V1-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(update_engine_stable_client PRIVATE
            -Werror
    )

    target_link_libraries(update_engine_stable_client
            libbase
            libbinder_ndk
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libdl
            libgflags
            liblog
            libm
            libupdate_engine_headers
            libupdate_engine_stable-V1-ndk
            libc
            libdl
            libm
    )
endif ()
#################### update_engine_stable_client ####################

#################### libupdate_engine_stable-V2-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/update_engine/stable/libupdate_engine_stable-V2-ndk_platform-source/gen/android/os/IUpdateEngineStable.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/update_engine/stable/libupdate_engine_stable-V2-ndk_platform-source/gen/android/os/IUpdateEngineStableCallback.cpp
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
    add_library(libupdate_engine_stable-V2-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(libupdate_engine_stable-V2-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/system/update_engine/stable/libupdate_engine_stable-V2-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(libupdate_engine_stable-V2-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(libupdate_engine_stable-V2-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libupdate_engine_stable-V2-ndk_platform ####################
