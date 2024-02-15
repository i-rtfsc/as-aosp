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

project(libnos_for_recovery)

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

#################### libnos_for_recovery ####################
set(FILES
        ${ANDROID_ROOT}/external/nos/host/generic/libnos/debug.cpp
        ${ANDROID_ROOT}/external/nos/host/generic/libnos_transport/transport.c
        ${ANDROID_ROOT}/external/nos/host/generic/libnos_transport/crc16.c
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
    add_library(libnos_for_recovery SHARED ${SOURCE_FILES})

    target_include_directories(libnos_for_recovery
            PUBLIC
            ${ANDROID_ROOT}/external/nos/host/generic/nugget/include
            ${ANDROID_ROOT}/external/nos/host/generic/libnos/include
            ${ANDROID_ROOT}/external/nos/host/generic/libnos_transport/include
            ${ANDROID_ROOT}/external/nos/host/generic
    )

    # Add c/c++ flags
    target_compile_options(libnos_for_recovery PRIVATE
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(libnos_for_recovery
            libbase
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### libnos_for_recovery ####################

#################### libnos_citadel_for_recovery ####################
set(FILES
        ${ANDROID_ROOT}/external/nos/host/generic/libnos/NuggetClient.cpp
        ${ANDROID_ROOT}/external/nos/host/generic/libnos_datagram/citadel.cpp
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
    add_library(libnos_citadel_for_recovery SHARED ${SOURCE_FILES})

    target_include_directories(libnos_citadel_for_recovery
            PUBLIC
            ${ANDROID_ROOT}/external/nos/host/generic
            ${ANDROID_ROOT}/external/nos/host/generic/nugget/include
            ${ANDROID_ROOT}/external/nos/host/generic/libnos/include
            ${ANDROID_ROOT}/external/nos/host/generic/libnos_transport/include
    )

    # Add c/c++ flags
    target_compile_options(libnos_citadel_for_recovery PRIVATE
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(libnos_citadel_for_recovery
            libbase
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
            libnos_for_recovery
    )
endif ()
#################### libnos_citadel_for_recovery ####################

#################### libnos_client_citadel ####################
set(FILES
        ${ANDROID_ROOT}/external/nos/host/generic/libnos/NuggetClient.cpp
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
    add_library(libnos_client_citadel SHARED ${SOURCE_FILES})

    target_include_directories(libnos_client_citadel
            PUBLIC
            ${ANDROID_ROOT}/external/nos/host/generic
            ${ANDROID_ROOT}/external/nos/host/generic/nugget/include
            ${ANDROID_ROOT}/external/nos/host/generic/libnos/include
            ${ANDROID_ROOT}/external/nos/host/generic/libnos_transport/include
    )

    # Add c/c++ flags
    target_compile_options(libnos_client_citadel PRIVATE
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(libnos_client_citadel
            libc++
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libdl
            libm
            libnos
            libnos_datagram_citadel
            libnos_transport
            nos_headers
            libc
            libdl
            libm
    )
endif ()
#################### libnos_client_citadel ####################

#################### libnos_citadel_for_fastboot ####################
set(FILES
        ${ANDROID_ROOT}/external/nos/host/generic/libnos/NuggetClient.cpp
        ${ANDROID_ROOT}/external/nos/host/generic/libnos_datagram/citadel.cpp
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
    add_library(libnos_citadel_for_fastboot SHARED ${SOURCE_FILES})

    target_include_directories(libnos_citadel_for_fastboot
            PUBLIC
            ${ANDROID_ROOT}/external/nos/host/generic
            ${ANDROID_ROOT}/external/nos/host/generic/nugget/include
            ${ANDROID_ROOT}/external/nos/host/generic/libnos/include
            ${ANDROID_ROOT}/external/nos/host/generic/libnos_transport/include
    )

    # Add c/c++ flags
    target_compile_options(libnos_citadel_for_fastboot PRIVATE
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(libnos_citadel_for_fastboot
            libbase
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
            libnos_for_fastboot
            libc
            libdl
            libm
    )
endif ()
#################### libnos_citadel_for_fastboot ####################

#################### libnos_for_fastboot ####################
set(FILES
        ${ANDROID_ROOT}/external/nos/host/generic/libnos/debug.cpp
        ${ANDROID_ROOT}/external/nos/host/generic/libnos_transport/transport.c
        ${ANDROID_ROOT}/external/nos/host/generic/libnos_transport/crc16.c
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
    add_library(libnos_for_fastboot SHARED ${SOURCE_FILES})

    target_include_directories(libnos_for_fastboot
            PUBLIC
            ${ANDROID_ROOT}/external/nos/host/generic/nugget/include
            ${ANDROID_ROOT}/external/nos/host/generic/libnos/include
            ${ANDROID_ROOT}/external/nos/host/generic/libnos_transport/include
            ${ANDROID_ROOT}/external/nos/host/generic
    )

    # Add c/c++ flags
    target_compile_options(libnos_for_fastboot PRIVATE
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(libnos_for_fastboot
            libbase
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
            libc
            libdl
            libm
    )
endif ()
#################### libnos_for_fastboot ####################
