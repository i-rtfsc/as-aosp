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

project(libcuttlefish_utils_product)

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

#################### libcuttlefish_utils_product ####################
set(FILES
        ${ANDROID_ROOT}/device/google/cuttlefish/common/libs/utils/inotify.cpp
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
    add_library(libcuttlefish_utils_product SHARED ${SOURCE_FILES})

    target_include_directories(libcuttlefish_utils_product
            PUBLIC
            ${ANDROID_ROOT}/device/google/cuttlefish/common/libs/utils
    )

    # Add c/c++ flags
    target_compile_options(libcuttlefish_utils_product PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(libcuttlefish_utils_product
            cuttlefish_common_headers
            libbase
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libdl
            liblog
            libm
            libc
            libdl
            libm
    )
endif ()
#################### libcuttlefish_utils_product ####################

#################### libcuttlefish_utils ####################
set(FILES
        ${ANDROID_ROOT}/device/google/cuttlefish/common/libs/utils/archive.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/common/libs/utils/base64.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/common/libs/utils/environment.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/common/libs/utils/files.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/common/libs/utils/flag_parser.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/common/libs/utils/flags_validator.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/common/libs/utils/inotify.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/common/libs/utils/json.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/common/libs/utils/network.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/common/libs/utils/proc_file_utils.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/common/libs/utils/scope_guard.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/common/libs/utils/shared_fd_flag.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/common/libs/utils/socket2socket_proxy.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/common/libs/utils/subprocess.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/common/libs/utils/tcp_socket.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/common/libs/utils/tee_logging.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/common/libs/utils/unix_sockets.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/common/libs/utils/users.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/common/libs/utils/vsock_connection.cpp
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
    add_library(libcuttlefish_utils SHARED ${SOURCE_FILES})

    target_include_directories(libcuttlefish_utils
            PUBLIC
            ${ANDROID_ROOT}/device/google/cuttlefish/common/libs/utils
    )

    # Add c/c++ flags
    target_compile_options(libcuttlefish_utils PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(libcuttlefish_utils
            cuttlefish_common_headers
            libbase
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto
            libcuttlefish_fs
            libdl
            libjsoncpp
            libm
    )
endif ()
#################### libcuttlefish_utils ####################

#################### libvsock_utils ####################
set(FILES
        ${ANDROID_ROOT}/device/google/cuttlefish/common/libs/utils/vsock_connection.cpp
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
    add_library(libvsock_utils SHARED ${SOURCE_FILES})

    target_include_directories(libvsock_utils
            PUBLIC
            ${ANDROID_ROOT}/device/google/cuttlefish/common/libs/utils
    )

    # Add c/c++ flags
    target_compile_options(libvsock_utils PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(libvsock_utils
            cuttlefish_common_headers
            libbase
            libc++
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libcuttlefish_fs
            libdl
            libjsoncpp
            liblog
            libm
            libc
            libdl
            libm
    )
endif ()
#################### libvsock_utils ####################
