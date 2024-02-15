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

project(libwebsockets)

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

#################### libwebsockets ####################
set(FILES
        ${ANDROID_ROOT}/external/libwebsockets/lib/core-net/adopt.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/core-net/client/client.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/core-net/client/conmon.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/core-net/client/connect.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/core-net/client/connect2.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/core-net/client/connect3.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/core-net/client/connect4.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/core-net/client/sort-dns.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/core-net/close.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/core-net/dummy-callback.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/core-net/network.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/core-net/output.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/core-net/pollfd.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/core-net/service.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/core-net/sorted-usec-list.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/core-net/state.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/core-net/vhost.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/core-net/wsi-timeout.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/core-net/wsi.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/core/alloc.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/core/buflist.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/core/context.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/core/libwebsockets.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/core/logs.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/core/lws_dll2.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/core/lws_map.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/core/vfs.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/event-libs/poll/poll.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/misc/base64-decode.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/misc/cache-ttl/file.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/misc/cache-ttl/heap.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/misc/cache-ttl/lws-cache-ttl.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/misc/dir.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/misc/lejp.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/misc/lws-ring.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/misc/lwsac/cached-file.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/misc/lwsac/lwsac.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/misc/prng.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/misc/sha-1.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/plat/unix/unix-caps.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/plat/unix/unix-fds.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/plat/unix/unix-file.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/plat/unix/unix-init.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/plat/unix/unix-misc.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/plat/unix/unix-pipe.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/plat/unix/unix-service.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/plat/unix/unix-sockets.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/roles/h1/ops-h1.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/roles/h2/hpack.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/roles/h2/http2.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/roles/h2/ops-h2.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/roles/http/client/client-http.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/roles/http/cookie.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/roles/http/date.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/roles/http/header.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/roles/http/parsers.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/roles/http/server/lejp-conf.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/roles/http/server/lws-spa.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/roles/http/server/server.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/roles/listen/ops-listen.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/roles/pipe/ops-pipe.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/roles/raw-file/ops-raw-file.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/roles/raw-skt/ops-raw-skt.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/roles/ws/client-parser-ws.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/roles/ws/client-ws.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/roles/ws/ops-ws.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/roles/ws/server-ws.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/system/smd/smd.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/system/system.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/tls/openssl/openssl-client.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/tls/openssl/openssl-server.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/tls/openssl/openssl-session.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/tls/openssl/openssl-ssl.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/tls/openssl/openssl-tls.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/tls/openssl/openssl-x509.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/tls/tls-client.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/tls/tls-network.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/tls/tls-server.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/tls/tls-sessions.c
        ${ANDROID_ROOT}/external/libwebsockets/lib/tls/tls.c
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
    add_library(libwebsockets SHARED ${SOURCE_FILES})

    target_include_directories(libwebsockets
            PUBLIC
            ${ANDROID_ROOT}/external/libwebsockets/include
            ${ANDROID_ROOT}/external/libwebsockets/plugins
            ${ANDROID_ROOT}/external/libwebsockets/lib/core
            ${ANDROID_ROOT}/external/libwebsockets/lib/core-net
            ${ANDROID_ROOT}/external/libwebsockets/lib/event-libs
            ${ANDROID_ROOT}/external/libwebsockets/lib/abstract
            ${ANDROID_ROOT}/external/libwebsockets/lib/tls
            ${ANDROID_ROOT}/external/libwebsockets/lib/roles
            ${ANDROID_ROOT}/external/libwebsockets/lib/event-libs/libuv
            ${ANDROID_ROOT}/external/libwebsockets/lib/event-libs/poll
            ${ANDROID_ROOT}/external/libwebsockets/lib/event-libs/libevent
            ${ANDROID_ROOT}/external/libwebsockets/lib/event-libs/glib
            ${ANDROID_ROOT}/external/libwebsockets/lib/event-libs/libev
            ${ANDROID_ROOT}/external/libwebsockets/lib/jose/jwe
            ${ANDROID_ROOT}/external/libwebsockets/lib/jose/jws
            ${ANDROID_ROOT}/external/libwebsockets/lib/jose
            ${ANDROID_ROOT}/external/libwebsockets/lib/misc
            ${ANDROID_ROOT}/external/libwebsockets/lib/roles/http
            ${ANDROID_ROOT}/external/libwebsockets/lib/roles/http/compression
            ${ANDROID_ROOT}/external/libwebsockets/lib/roles/h1
            ${ANDROID_ROOT}/external/libwebsockets/lib/roles/h2
            ${ANDROID_ROOT}/external/libwebsockets/lib/roles/ws
            ${ANDROID_ROOT}/external/libwebsockets/lib/roles/cgi
            ${ANDROID_ROOT}/external/libwebsockets/lib/roles/dbus
            ${ANDROID_ROOT}/external/libwebsockets/lib/roles/raw-proxy
            ${ANDROID_ROOT}/external/libwebsockets/lib/system/async-dns
            ${ANDROID_ROOT}/external/libwebsockets/lib/system/smd
            ${ANDROID_ROOT}/external/libwebsockets/lib/system/metrics
            ${ANDROID_ROOT}/external/libwebsockets/lib/roles/mqtt
            ${ANDROID_ROOT}/external/libwebsockets/lib/plat/unix
            ${ANDROID_ROOT}/external/libwebsockets/lib
            ${ANDROID_ROOT}/external/libwebsockets
    )

    # Add c/c++ flags
    target_compile_options(libwebsockets PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libwebsockets
            libc
            libc++_static
            libc++demangle
            libcap
            libclang_rt.builtins
            libcrypto
            libdl
            libm
            libssl
    )
endif ()
#################### libwebsockets ####################
