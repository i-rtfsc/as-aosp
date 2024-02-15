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

project(init_ueventHandler_fuzzer)

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

#################### init_property_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/system/core/init/fuzzer/init_property_fuzzer.cpp
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
    add_executable(init_property_fuzzer ${SOURCE_FILES})

    target_include_directories(init_property_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/system/core/init/fuzzer
    )

    # Add c/c++ flags
    target_compile_options(init_property_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(init_property_fuzzer
            libbase
            libc
            libc++
            libc++demangle
            libc++fs
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            libfs_mgr
            libhidl-gen-utils
            libinit
            libinit_headers
            libkeyutils
            liblmkd_utils
            liblog
            libm
            libmodprobe
            libprocessgroup
            libpropertyinfoparser
            libprotobuf-cpp-lite
            libselinux
            libsnapshot_init
    )
endif ()
#################### init_property_fuzzer ####################

#################### init_parser_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/system/core/init/fuzzer/init_parser_fuzzer.cpp
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
    add_executable(init_parser_fuzzer ${SOURCE_FILES})

    target_include_directories(init_parser_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/system/core/init/fuzzer
    )

    # Add c/c++ flags
    target_compile_options(init_parser_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(init_parser_fuzzer
            libbase
            libc
            libc++
            libc++demangle
            libc++fs
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            libfs_mgr
            libhidl-gen-utils
            libhidlmetadata
            libinit
            libinit_headers
            libkeyutils
            liblmkd_utils
            liblog
            libm
            libmodprobe
            libprocessgroup
            libpropertyinfoparser
            libprotobuf-cpp-lite
            libselinux
            libsnapshot_init
    )
endif ()
#################### init_parser_fuzzer ####################

#################### init_ueventHandler_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/system/core/init/fuzzer/init_ueventHandler_fuzzer.cpp
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
    add_executable(init_ueventHandler_fuzzer ${SOURCE_FILES})

    target_include_directories(init_ueventHandler_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/system/core/init/fuzzer
    )

    # Add c/c++ flags
    target_compile_options(init_ueventHandler_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(init_ueventHandler_fuzzer
            libbase
            libc
            libc++
            libc++demangle
            libc++fs
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            libfs_mgr
            libhidl-gen-utils
            libinit
            libinit_headers
            libkeyutils
            liblmkd_utils
            liblog
            libm
            libmodprobe
            libprocessgroup
            libpropertyinfoparser
            libprotobuf-cpp-lite
            libselinux
            libsnapshot_init
    )
endif ()
#################### init_ueventHandler_fuzzer ####################
