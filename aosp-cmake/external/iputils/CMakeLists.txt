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

project(ping)

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

#################### tracepath ####################
set(FILES
        ${ANDROID_ROOT}/external/iputils/tracepath.c
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
    add_executable(tracepath ${SOURCE_FILES})

    target_include_directories(tracepath
            PUBLIC
            ${ANDROID_ROOT}/external/iputils
    )

    # Add c/c++ flags
    target_compile_options(tracepath PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(tracepath
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### tracepath ####################

#################### tracepath6 ####################
set(FILES
        ${ANDROID_ROOT}/external/iputils/tracepath6.c
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
    add_executable(tracepath6 ${SOURCE_FILES})

    target_include_directories(tracepath6
            PUBLIC
            ${ANDROID_ROOT}/external/iputils
    )

    # Add c/c++ flags
    target_compile_options(tracepath6 PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(tracepath6
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### tracepath6 ####################

#################### ping6 ####################
set(FILES
        ${ANDROID_ROOT}/external/iputils/ping6.c
        ${ANDROID_ROOT}/external/iputils/ping_common.c
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
    add_executable(ping6 ${SOURCE_FILES})

    target_include_directories(ping6
            PUBLIC
            ${ANDROID_ROOT}/external/iputils
    )

    # Add c/c++ flags
    target_compile_options(ping6 PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(ping6
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcrypto
            libdl
            libm
    )
endif ()
#################### ping6 ####################

#################### traceroute6 ####################
set(FILES
        ${ANDROID_ROOT}/external/iputils/traceroute6.c
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
    add_executable(traceroute6 ${SOURCE_FILES})

    target_include_directories(traceroute6
            PUBLIC
            ${ANDROID_ROOT}/external/iputils
    )

    # Add c/c++ flags
    target_compile_options(traceroute6 PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(traceroute6
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### traceroute6 ####################

#################### arping ####################
set(FILES
        ${ANDROID_ROOT}/external/iputils/arping.c
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
    add_executable(arping ${SOURCE_FILES})

    target_include_directories(arping
            PUBLIC
            ${ANDROID_ROOT}/external/iputils
    )

    # Add c/c++ flags
    target_compile_options(arping PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(arping
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### arping ####################

#################### ping ####################
set(FILES
        ${ANDROID_ROOT}/external/iputils/ping.c
        ${ANDROID_ROOT}/external/iputils/ping_common.c
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
    add_executable(ping ${SOURCE_FILES})

    target_include_directories(ping
            PUBLIC
            ${ANDROID_ROOT}/external/iputils
    )

    # Add c/c++ flags
    target_compile_options(ping PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(ping
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### ping ####################
