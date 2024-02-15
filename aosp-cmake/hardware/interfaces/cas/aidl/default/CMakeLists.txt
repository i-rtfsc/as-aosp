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

project(android.hardware.cas-service_fuzzer)

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

#################### android.hardware.cas-service.example ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/cas/aidl/default/service.cpp
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
    add_executable(android.hardware.cas-service.example ${SOURCE_FILES})

    target_include_directories(android.hardware.cas-service.example
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/cas/aidl/default
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.cas-service.example PRIVATE
            -Werror
    )

    target_link_libraries(android.hardware.cas-service.example
            android.hardware.cas-V1-ndk
            libaidlcommonsupport
            libbinder_ndk
            libc++
            libc++demangle
            libc
            libcasexampleimpl
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            liblog
            libm
            libutils
            media_plugin_headers
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.cas-service.example ####################

#################### libcasexampleimpl ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/cas/aidl/default/CasImpl.cpp
        ${ANDROID_ROOT}/hardware/interfaces/cas/aidl/default/DescramblerImpl.cpp
        ${ANDROID_ROOT}/hardware/interfaces/cas/aidl/default/MediaCasService.cpp
        ${ANDROID_ROOT}/hardware/interfaces/cas/aidl/default/SharedLibrary.cpp
        ${ANDROID_ROOT}/hardware/interfaces/cas/aidl/default/TypeConvert.cpp
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
    add_library(libcasexampleimpl SHARED ${SOURCE_FILES})

    target_include_directories(libcasexampleimpl
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/cas/aidl/default
    )

    # Add c/c++ flags
    target_compile_options(libcasexampleimpl PRIVATE
            -Werror
    )

    target_link_libraries(libcasexampleimpl
            android.hardware.cas-V1-ndk
            libaidlcommonsupport
            libbinder_ndk
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            liblog
            libm
            libstagefright_foundation_headers
            libutils
            media_plugin_headers
    )
endif ()
#################### libcasexampleimpl ####################

#################### android.hardware.cas-service.example-lazy ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/cas/aidl/default/service.cpp
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
    add_executable(android.hardware.cas-service.example-lazy ${SOURCE_FILES})

    target_include_directories(android.hardware.cas-service.example-lazy
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/cas/aidl/default
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.cas-service.example-lazy PRIVATE
            -Werror
    )

    target_link_libraries(android.hardware.cas-service.example-lazy
            android.hardware.cas-V1-ndk
            libaidlcommonsupport
            libbinder_ndk
            libc++
            libc++demangle
            libc
            libcasexampleimpl
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            liblog
            libm
            libutils
            media_plugin_headers
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.cas-service.example-lazy ####################

#################### android.hardware.cas-service_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/cas/aidl/default/fuzzer.cpp
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
    add_executable(android.hardware.cas-service_fuzzer ${SOURCE_FILES})

    target_include_directories(android.hardware.cas-service_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/cas/aidl/default
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.cas-service_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(android.hardware.cas-service_fuzzer
            android.hardware.cas-V1-ndk
            libaidlcommonsupport
            libbase
            libbinder
            libbinder_ndk
            libbinder_random_parcel
            libc++
            libc++demangle
            libc
            libcasexampleimpl
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcutils
            libdl
            liblog
            libm
            libutils
            media_plugin_headers
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.cas-service_fuzzer ####################
