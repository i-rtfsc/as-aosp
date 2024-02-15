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

project(otbr-agent)

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

#################### otbr-agent ####################
set(FILES
        ${ANDROID_ROOT}/external/ot-br-posix/src/agent/application.cpp
        ${ANDROID_ROOT}/external/ot-br-posix/src/agent/main.cpp
        ${ANDROID_ROOT}/external/ot-br-posix/src/border_agent/border_agent.cpp
        ${ANDROID_ROOT}/external/ot-br-posix/src/ncp/ncp_openthread.cpp
        ${ANDROID_ROOT}/external/ot-br-posix/src/sdp_proxy/advertising_proxy.cpp
        ${ANDROID_ROOT}/external/ot-br-posix/src/sdp_proxy/discovery_proxy.cpp
        ${ANDROID_ROOT}/external/ot-br-posix/src/utils/thread_helper.cpp
        ${ANDROID_ROOT}/external/ot-br-posix/src/common/dns_utils.cpp
        ${ANDROID_ROOT}/external/ot-br-posix/src/common/logging.cpp
        ${ANDROID_ROOT}/external/ot-br-posix/src/common/task_runner.cpp
        ${ANDROID_ROOT}/external/ot-br-posix/src/common/types.cpp
        ${ANDROID_ROOT}/external/ot-br-posix/src/common/mainloop.cpp
        ${ANDROID_ROOT}/external/ot-br-posix/src/common/mainloop_manager.cpp
        ${ANDROID_ROOT}/external/ot-br-posix/src/mdns/mdns.cpp
        ${ANDROID_ROOT}/external/ot-br-posix/src/mdns/mdns_mdnssd.cpp
        ${ANDROID_ROOT}/external/ot-br-posix/src/utils/dns_utils.cpp
        ${ANDROID_ROOT}/external/ot-br-posix/src/utils/hex.cpp
        ${ANDROID_ROOT}/external/ot-br-posix/src/utils/infra_link_selector.cpp
        ${ANDROID_ROOT}/external/ot-br-posix/src/utils/socket_utils.cpp
        ${ANDROID_ROOT}/external/ot-br-posix/src/utils/string_utils.cpp
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
    add_executable(otbr-agent ${SOURCE_FILES})

    target_include_directories(otbr-agent
            PUBLIC
            ${ANDROID_ROOT}/external/ot-br-posix/include
            ${ANDROID_ROOT}/external/ot-br-posix/src
            ${ANDROID_ROOT}/external/ot-br-posix
            ${OUT_INTERMEDIATES_ROOT}/external/ot-br-posix/otbr_version_header/gen
    )

    # Add c/c++ flags
    target_compile_options(otbr-agent PRIVATE
            -Werror
            -Wall
            -Wextra
    )

    target_link_libraries(otbr-agent
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libm
            libmdnssd
            libopenthread-cli
            libutils
            openthread-tcplp
            ot-core
    )
endif ()
#################### otbr-agent ####################
