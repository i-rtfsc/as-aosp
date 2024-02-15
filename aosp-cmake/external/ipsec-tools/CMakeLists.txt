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

project(racoon)

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

#################### racoon ####################
set(FILES
        ${ANDROID_ROOT}/external/ipsec-tools/src/racoon/algorithm.c
        ${ANDROID_ROOT}/external/ipsec-tools/src/racoon/crypto_openssl.c
        ${ANDROID_ROOT}/external/ipsec-tools/src/racoon/genlist.c
        ${ANDROID_ROOT}/external/ipsec-tools/src/racoon/handler.c
        ${ANDROID_ROOT}/external/ipsec-tools/src/racoon/isakmp.c
        ${ANDROID_ROOT}/external/ipsec-tools/src/racoon/isakmp_agg.c
        ${ANDROID_ROOT}/external/ipsec-tools/src/racoon/isakmp_base.c
        ${ANDROID_ROOT}/external/ipsec-tools/src/racoon/isakmp_cfg.c
        ${ANDROID_ROOT}/external/ipsec-tools/src/racoon/isakmp_frag.c
        ${ANDROID_ROOT}/external/ipsec-tools/src/racoon/isakmp_ident.c
        ${ANDROID_ROOT}/external/ipsec-tools/src/racoon/isakmp_inf.c
        ${ANDROID_ROOT}/external/ipsec-tools/src/racoon/isakmp_newg.c
        ${ANDROID_ROOT}/external/ipsec-tools/src/racoon/isakmp_quick.c
        ${ANDROID_ROOT}/external/ipsec-tools/src/racoon/isakmp_unity.c
        ${ANDROID_ROOT}/external/ipsec-tools/src/racoon/isakmp_xauth.c
        ${ANDROID_ROOT}/external/ipsec-tools/src/racoon/ipsec_doi.c
        ${ANDROID_ROOT}/external/ipsec-tools/src/racoon/nattraversal.c
        ${ANDROID_ROOT}/external/ipsec-tools/src/racoon/oakley.c
        ${ANDROID_ROOT}/external/ipsec-tools/src/racoon/pfkey.c
        ${ANDROID_ROOT}/external/ipsec-tools/src/racoon/policy.c
        ${ANDROID_ROOT}/external/ipsec-tools/src/racoon/proposal.c
        ${ANDROID_ROOT}/external/ipsec-tools/src/racoon/remoteconf.c
        ${ANDROID_ROOT}/external/ipsec-tools/src/racoon/schedule.c
        ${ANDROID_ROOT}/external/ipsec-tools/src/racoon/sockmisc.c
        ${ANDROID_ROOT}/external/ipsec-tools/src/racoon/str2val.c
        ${ANDROID_ROOT}/external/ipsec-tools/src/racoon/strnames.c
        ${ANDROID_ROOT}/external/ipsec-tools/src/racoon/vendorid.c
        ${ANDROID_ROOT}/external/ipsec-tools/src/racoon/vmbuf.c
        ${ANDROID_ROOT}/external/ipsec-tools/main.c
        ${ANDROID_ROOT}/external/ipsec-tools/setup.c
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
    add_executable(racoon ${SOURCE_FILES})

    target_include_directories(racoon
            PUBLIC
            ${ANDROID_ROOT}/external/ipsec-tools/src/include-glibc
            ${ANDROID_ROOT}/external/ipsec-tools/src/racoon
            ${ANDROID_ROOT}/external/ipsec-tools/src/racoon/missing
            ${ANDROID_ROOT}/external/ipsec-tools
            ${ANDROID_ROOT}/external/ipsec-tools/src/libipsec
    )

    # Add c/c++ flags
    target_compile_options(racoon PRIVATE
            -Wno-unused-parameter
            -Werror
            -Wno-unused-variable
            -Wno-unused-but-set-variable
            -Wno-unused-function
            -Wno-unused-label
            -Wno-unused-value
    )

    target_link_libraries(racoon
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcrypto
            libcutils
            libdl
            libipsec
            libkeystore-engine
            liblog
            libm
            libnetd_client
    )
endif ()
#################### racoon ####################

#################### libipsec ####################
set(FILES
        ${ANDROID_ROOT}/external/ipsec-tools/src/libipsec/pfkey.c
        ${ANDROID_ROOT}/external/ipsec-tools/src/libipsec/ipsec_strerror.c
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
    add_library(libipsec SHARED ${SOURCE_FILES})

    target_include_directories(libipsec
            PUBLIC
            ${ANDROID_ROOT}/external/ipsec-tools/src/libipsec
            ${ANDROID_ROOT}/external/ipsec-tools/src/include-glibc
            ${ANDROID_ROOT}/external/ipsec-tools
    )

    # Add c/c++ flags
    target_compile_options(libipsec PRIVATE
            -Wno-unused-parameter
            -Werror
    )

    target_link_libraries(libipsec
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libipsec ####################
