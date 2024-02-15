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

project(android.hardware.drm@1.0-service-lazy)

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

#################### android.hardware.drm-1.0-helper ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/drm/1.0/default/SharedLibrary.cpp
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
    add_library(android.hardware.drm-1.0-helper SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.drm-1.0-helper
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/drm/1.0/default/include
            ${ANDROID_ROOT}/hardware/interfaces/drm/1.0/default
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.drm-1.0-helper PRIVATE
            -Wall
            -Werror
            -Werror
            -Wextra
            -Wall
    )

    target_link_libraries(android.hardware.drm-1.0-helper
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            liblog
            libm
            libutils_headers
    )
endif ()
#################### android.hardware.drm-1.0-helper ####################

#################### android.hardware.drm-1.0-service ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/drm/1.0/default/service.cpp
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
    add_executable(android.hardware.drm-1.0-service ${SOURCE_FILES})

    target_include_directories(android.hardware.drm-1.0-service
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/drm/1.0/default
            ${ANDROID_ROOT}/hardware/interfaces/drm/1.0/default/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.drm-1.0-service PRIVATE
            -Werror
    )

    target_link_libraries(android.hardware.drm-1.0-service
            android.hardware.drm-1.0-helper
            android.hardware.drm-1.0
            android.hidl.memory-1.0
            libbinder
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libhardware
            libhidlbase
            liblog
            libm
            libutils
            media_plugin_headers
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.drm-1.0-service ####################

#################### android.hardware.drm-1.0-service-lazy ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/drm/1.0/default/serviceLazy.cpp
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
    add_executable(android.hardware.drm-1.0-service-lazy ${SOURCE_FILES})

    target_include_directories(android.hardware.drm-1.0-service-lazy
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/drm/1.0/default
            ${ANDROID_ROOT}/hardware/interfaces/drm/1.0/default/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.drm-1.0-service-lazy PRIVATE
            -Werror
    )

    target_link_libraries(android.hardware.drm-1.0-service-lazy
            android.hardware.drm-1.0-helper
            android.hardware.drm-1.0
            android.hidl.memory-1.0
            libbinder
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libhardware
            libhidlbase
            liblog
            libm
            libutils
            media_plugin_headers
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.drm-1.0-service-lazy ####################

#################### android.hardware.drm-1.0-impl ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/drm/1.0/default/DrmFactory.cpp
        ${ANDROID_ROOT}/hardware/interfaces/drm/1.0/default/DrmPlugin.cpp
        ${ANDROID_ROOT}/hardware/interfaces/drm/1.0/default/CryptoFactory.cpp
        ${ANDROID_ROOT}/hardware/interfaces/drm/1.0/default/CryptoPlugin.cpp
        ${ANDROID_ROOT}/hardware/interfaces/drm/1.0/default/LegacyPluginPath.cpp
        ${ANDROID_ROOT}/hardware/interfaces/drm/1.0/default/TypeConvert.cpp
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
    add_library(android.hardware.drm-1.0-impl SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.drm-1.0-impl
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/drm/1.0/default
            ${ANDROID_ROOT}/hardware/interfaces/drm/1.0/default/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.drm-1.0-impl PRIVATE
            -Werror
    )

    target_link_libraries(android.hardware.drm-1.0-impl
            android.hardware.drm-1.0-helper
            android.hardware.drm-1.0
            android.hidl.memory-1.0
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libhidlbase
            libhidlmemory
            liblog
            libm
            libstagefright_foundation
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.drm-1.0-impl ####################
