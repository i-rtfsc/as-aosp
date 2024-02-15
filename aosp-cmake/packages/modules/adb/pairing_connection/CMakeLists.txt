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

project(libadb_pairing_connection_static)

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

#################### libadb_pairing_connection_static ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/adb/pairing_connection/pairing_connection.cpp
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
    add_library(libadb_pairing_connection_static SHARED ${SOURCE_FILES})

    target_include_directories(libadb_pairing_connection_static
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/adb/pairing_connection/include
            ${ANDROID_ROOT}/packages/modules/adb/pairing_connection
    )

    # Add c/c++ flags
    target_compile_options(libadb_pairing_connection_static PRIVATE
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(libadb_pairing_connection_static
            libadb_pairing_auth
            libadb_protos_static
            libadb_tls_connection_static
            libbase
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto
            libdl
            liblog
            libm
            libprotobuf-cpp-lite
            libssl
            libwinpthread
    )
endif ()
#################### libadb_pairing_connection_static ####################

#################### libadb_pairing_server_static ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/adb/pairing_connection/pairing_server.cpp
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
    add_library(libadb_pairing_server_static SHARED ${SOURCE_FILES})

    target_include_directories(libadb_pairing_server_static
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/adb/pairing_connection/include
            ${ANDROID_ROOT}/packages/modules/adb/pairing_connection
    )

    # Add c/c++ flags
    target_compile_options(libadb_pairing_server_static PRIVATE
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(libadb_pairing_server_static
            libadb_crypto_static
            libadb_pairing_auth
            libadb_pairing_connection
            libadb_protos_static
            libbase
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto
            libcrypto_utils
            libcutils
            libdl
            liblog
            libm
    )
endif ()
#################### libadb_pairing_server_static ####################

#################### libadb_pairing_connection ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/adb/pairing_connection/pairing_connection.cpp
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
    add_library(libadb_pairing_connection SHARED ${SOURCE_FILES})

    target_include_directories(libadb_pairing_connection
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/adb/pairing_connection/include
            ${ANDROID_ROOT}/packages/modules/adb/pairing_connection
    )

    # Add c/c++ flags
    target_compile_options(libadb_pairing_connection PRIVATE
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(libadb_pairing_connection
            libadb_pairing_auth
            libadb_protos
            libadb_tls_connection
            libbase
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto
            libdl
            liblog
            libm
            libprotobuf-cpp-lite
            libssl
            libwinpthread
    )
endif ()
#################### libadb_pairing_connection ####################

#################### libadb_pairing_server ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/adb/pairing_connection/pairing_server.cpp
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
    add_library(libadb_pairing_server SHARED ${SOURCE_FILES})

    target_include_directories(libadb_pairing_server
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/adb/pairing_connection/include
            ${ANDROID_ROOT}/packages/modules/adb/pairing_connection
    )

    # Add c/c++ flags
    target_compile_options(libadb_pairing_server PRIVATE
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(libadb_pairing_server
            libadb_crypto
            libadb_pairing_auth
            libadb_pairing_connection
            libadb_protos
            libbase
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto
            libcrypto_utils
            libcutils
            libdl
            liblog
            libm
    )
endif ()
#################### libadb_pairing_server ####################
