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

project(libevent)

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

#################### libevent ####################
set(FILES
        ${ANDROID_ROOT}/external/libevent/buffer.c
        ${ANDROID_ROOT}/external/libevent/bufferevent.c
        ${ANDROID_ROOT}/external/libevent/bufferevent_filter.c
        ${ANDROID_ROOT}/external/libevent/bufferevent_pair.c
        ${ANDROID_ROOT}/external/libevent/bufferevent_ratelim.c
        ${ANDROID_ROOT}/external/libevent/bufferevent_sock.c
        ${ANDROID_ROOT}/external/libevent/event.c
        ${ANDROID_ROOT}/external/libevent/evmap.c
        ${ANDROID_ROOT}/external/libevent/evthread.c
        ${ANDROID_ROOT}/external/libevent/evthread_pthread.c
        ${ANDROID_ROOT}/external/libevent/evutil.c
        ${ANDROID_ROOT}/external/libevent/evutil_rand.c
        ${ANDROID_ROOT}/external/libevent/evutil_time.c
        ${ANDROID_ROOT}/external/libevent/listener.c
        ${ANDROID_ROOT}/external/libevent/log.c
        ${ANDROID_ROOT}/external/libevent/signal.c
        ${ANDROID_ROOT}/external/libevent/strlcpy.c
        ${ANDROID_ROOT}/external/libevent/evdns.c
        ${ANDROID_ROOT}/external/libevent/event_tagging.c
        ${ANDROID_ROOT}/external/libevent/evrpc.c
        ${ANDROID_ROOT}/external/libevent/http.c
        ${ANDROID_ROOT}/external/libevent/poll.c
        ${ANDROID_ROOT}/external/libevent/select.c
        ${ANDROID_ROOT}/external/libevent/epoll.c
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
    add_library(libevent SHARED ${SOURCE_FILES})

    target_include_directories(libevent
            PUBLIC
            ${ANDROID_ROOT}/external/libevent/include
            ${ANDROID_ROOT}/external/libevent/compat
            ${ANDROID_ROOT}/external/libevent
    )

    # Add c/c++ flags
    target_compile_options(libevent PRIVATE
            -Wno-unused-parameter
            -Werror
    )

    target_link_libraries(libevent
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libevent ####################
