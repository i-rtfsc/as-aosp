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

project(pppd)

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

#################### pppol2tp-android ####################
set(FILES
        ${ANDROID_ROOT}/external/ppp/pppd/plugins/pppol2tp-android/pppol2tp-android.c
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
    add_library(pppol2tp-android SHARED ${SOURCE_FILES})

    target_include_directories(pppol2tp-android
            PUBLIC
            ${ANDROID_ROOT}/external/ppp/pppd/include
            ${ANDROID_ROOT}/external/ppp/pppd
    )

    # Add c/c++ flags
    target_compile_options(pppol2tp-android PRIVATE
            -Wno-unused-parameter
            -Werror
    )

    target_link_libraries(pppol2tp-android
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### pppol2tp-android ####################

#################### pppopptp-android ####################
set(FILES
        ${ANDROID_ROOT}/external/ppp/pppd/plugins/pppopptp-android/pppopptp-android.c
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
    add_library(pppopptp-android SHARED ${SOURCE_FILES})

    target_include_directories(pppopptp-android
            PUBLIC
            ${ANDROID_ROOT}/external/ppp/pppd/include
            ${ANDROID_ROOT}/external/ppp/pppd
    )

    # Add c/c++ flags
    target_compile_options(pppopptp-android PRIVATE
            -Wno-unused-parameter
            -Werror
    )

    target_link_libraries(pppopptp-android
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### pppopptp-android ####################

#################### libpppd ####################
set(FILES
        ${ANDROID_ROOT}/external/ppp/pppd/auth.c
        ${ANDROID_ROOT}/external/ppp/pppd/ccp.c
        ${ANDROID_ROOT}/external/ppp/pppd/chap-md5.c
        ${ANDROID_ROOT}/external/ppp/pppd/chap-new.c
        ${ANDROID_ROOT}/external/ppp/pppd/chap_ms.c
        ${ANDROID_ROOT}/external/ppp/pppd/demand.c
        ${ANDROID_ROOT}/external/ppp/pppd/eap.c
        ${ANDROID_ROOT}/external/ppp/pppd/ecp.c
        ${ANDROID_ROOT}/external/ppp/pppd/eui64.c
        ${ANDROID_ROOT}/external/ppp/pppd/fsm.c
        ${ANDROID_ROOT}/external/ppp/pppd/ipcp.c
        ${ANDROID_ROOT}/external/ppp/pppd/ipv6cp.c
        ${ANDROID_ROOT}/external/ppp/pppd/lcp.c
        ${ANDROID_ROOT}/external/ppp/pppd/magic.c
        ${ANDROID_ROOT}/external/ppp/pppd/main.c
        ${ANDROID_ROOT}/external/ppp/pppd/options.c
        ${ANDROID_ROOT}/external/ppp/pppd/pppcrypt.c
        ${ANDROID_ROOT}/external/ppp/pppd/pppox.c
        ${ANDROID_ROOT}/external/ppp/pppd/session.c
        ${ANDROID_ROOT}/external/ppp/pppd/sys-linux.c
        ${ANDROID_ROOT}/external/ppp/pppd/tty.c
        ${ANDROID_ROOT}/external/ppp/pppd/upap.c
        ${ANDROID_ROOT}/external/ppp/pppd/utils.c
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
    add_library(libpppd SHARED ${SOURCE_FILES})

    target_include_directories(libpppd
            PUBLIC
            ${ANDROID_ROOT}/external/ppp/pppd
            ${ANDROID_ROOT}/external/ppp/pppd/include
    )

    # Add c/c++ flags
    target_compile_options(libpppd PRIVATE
            -Wno-unused-parameter
            -Werror
    )

    target_link_libraries(libpppd
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcrypto
            libcutils
            libdl
            liblog
            libm
    )
endif ()
#################### libpppd ####################

#################### pppd ####################
set(FILES
        ${ANDROID_ROOT}/external/ppp/pppd/auth.c
        ${ANDROID_ROOT}/external/ppp/pppd/ccp.c
        ${ANDROID_ROOT}/external/ppp/pppd/chap-md5.c
        ${ANDROID_ROOT}/external/ppp/pppd/chap-new.c
        ${ANDROID_ROOT}/external/ppp/pppd/chap_ms.c
        ${ANDROID_ROOT}/external/ppp/pppd/demand.c
        ${ANDROID_ROOT}/external/ppp/pppd/eap.c
        ${ANDROID_ROOT}/external/ppp/pppd/ecp.c
        ${ANDROID_ROOT}/external/ppp/pppd/eui64.c
        ${ANDROID_ROOT}/external/ppp/pppd/fsm.c
        ${ANDROID_ROOT}/external/ppp/pppd/ipcp.c
        ${ANDROID_ROOT}/external/ppp/pppd/ipv6cp.c
        ${ANDROID_ROOT}/external/ppp/pppd/lcp.c
        ${ANDROID_ROOT}/external/ppp/pppd/magic.c
        ${ANDROID_ROOT}/external/ppp/pppd/main.c
        ${ANDROID_ROOT}/external/ppp/pppd/options.c
        ${ANDROID_ROOT}/external/ppp/pppd/pppcrypt.c
        ${ANDROID_ROOT}/external/ppp/pppd/pppox.c
        ${ANDROID_ROOT}/external/ppp/pppd/session.c
        ${ANDROID_ROOT}/external/ppp/pppd/sys-linux.c
        ${ANDROID_ROOT}/external/ppp/pppd/tty.c
        ${ANDROID_ROOT}/external/ppp/pppd/upap.c
        ${ANDROID_ROOT}/external/ppp/pppd/utils.c
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
    add_executable(pppd ${SOURCE_FILES})

    target_include_directories(pppd
            PUBLIC
            ${ANDROID_ROOT}/external/ppp/pppd/include
            ${ANDROID_ROOT}/external/ppp/pppd
    )

    # Add c/c++ flags
    target_compile_options(pppd PRIVATE
            -Wno-unused-parameter
            -Werror
    )

    target_link_libraries(pppd
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcrypto
            libcutils
            libdl
            liblog
            libm
    )
endif ()
#################### pppd ####################
