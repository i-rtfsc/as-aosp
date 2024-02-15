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

project(libbrillo)

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

#################### libbrillo ####################
set(FILES
        ${ANDROID_ROOT}/external/libbrillo/brillo/backoff_entry.cc
        ${ANDROID_ROOT}/external/libbrillo/brillo/data_encoding.cc
        ${ANDROID_ROOT}/external/libbrillo/brillo/errors/error.cc
        ${ANDROID_ROOT}/external/libbrillo/brillo/errors/error_codes.cc
        ${ANDROID_ROOT}/external/libbrillo/brillo/flag_helper.cc
        ${ANDROID_ROOT}/external/libbrillo/brillo/key_value_store.cc
        ${ANDROID_ROOT}/external/libbrillo/brillo/message_loops/base_message_loop.cc
        ${ANDROID_ROOT}/external/libbrillo/brillo/message_loops/message_loop.cc
        ${ANDROID_ROOT}/external/libbrillo/brillo/message_loops/message_loop_utils.cc
        ${ANDROID_ROOT}/external/libbrillo/brillo/mime_utils.cc
        ${ANDROID_ROOT}/external/libbrillo/brillo/osrelease_reader.cc
        ${ANDROID_ROOT}/external/libbrillo/brillo/process.cc
        ${ANDROID_ROOT}/external/libbrillo/brillo/process_information.cc
        ${ANDROID_ROOT}/external/libbrillo/brillo/secure_blob.cc
        ${ANDROID_ROOT}/external/libbrillo/brillo/strings/string_utils.cc
        ${ANDROID_ROOT}/external/libbrillo/brillo/syslog_logging.cc
        ${ANDROID_ROOT}/external/libbrillo/brillo/type_name_undecorate.cc
        ${ANDROID_ROOT}/external/libbrillo/brillo/url_utils.cc
        ${ANDROID_ROOT}/external/libbrillo/brillo/userdb_utils.cc
        ${ANDROID_ROOT}/external/libbrillo/brillo/value_conversion.cc
        ${ANDROID_ROOT}/external/libbrillo/brillo/asynchronous_signal_handler.cc
        ${ANDROID_ROOT}/external/libbrillo/brillo/daemons/daemon.cc
        ${ANDROID_ROOT}/external/libbrillo/brillo/file_utils.cc
        ${ANDROID_ROOT}/external/libbrillo/brillo/process_reaper.cc
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
    add_library(libbrillo SHARED ${SOURCE_FILES})

    target_include_directories(libbrillo
            PUBLIC
            ${ANDROID_ROOT}/external/libbrillo
    )

    # Add c/c++ flags
    target_compile_options(libbrillo PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(libbrillo
            libc
            libc++
            libc++_static
            libc++demangle
            libchrome
            libclang_rt.builtins
            libdl
            libgtest_prod_headers
            libm
            libmodpb64
    )
endif ()
#################### libbrillo ####################

#################### libbrillo-http ####################
set(FILES
        ${ANDROID_ROOT}/external/libbrillo/brillo/http/curl_api.cc
        ${ANDROID_ROOT}/external/libbrillo/brillo/http/http_connection_curl.cc
        ${ANDROID_ROOT}/external/libbrillo/brillo/http/http_form_data.cc
        ${ANDROID_ROOT}/external/libbrillo/brillo/http/http_request.cc
        ${ANDROID_ROOT}/external/libbrillo/brillo/http/http_transport.cc
        ${ANDROID_ROOT}/external/libbrillo/brillo/http/http_transport_curl.cc
        ${ANDROID_ROOT}/external/libbrillo/brillo/http/http_utils.cc
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
    add_library(libbrillo-http SHARED ${SOURCE_FILES})

    target_include_directories(libbrillo-http
            PUBLIC
            ${ANDROID_ROOT}/external/libbrillo
    )

    # Add c/c++ flags
    target_compile_options(libbrillo-http PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(libbrillo-http
            libbrillo
            libbrillo-stream
            libc
            libc++
            libc++demangle
            libchrome
            libclang_rt.builtins
            libcurl
            libdl
            libgtest_prod_headers
            libm
    )
endif ()
#################### libbrillo-http ####################

#################### libbrillo-binder ####################
set(FILES
        ${ANDROID_ROOT}/external/libbrillo/brillo/binder_watcher.cc
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
    add_library(libbrillo-binder SHARED ${SOURCE_FILES})

    target_include_directories(libbrillo-binder
            PUBLIC
            ${ANDROID_ROOT}/external/libbrillo
    )

    # Add c/c++ flags
    target_compile_options(libbrillo-binder PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(libbrillo-binder
            libbinder
            libbrillo
            libc
            libc++
            libc++demangle
            libchrome
            libclang_rt.builtins
            libdl
            libgtest_prod_headers
            libm
            libutils
    )
endif ()
#################### libbrillo-binder ####################

#################### libbrillo-policy ####################
set(FILES
        ${ANDROID_ROOT}/external/libbrillo/policy/device_policy.cc
        ${ANDROID_ROOT}/external/libbrillo/policy/libpolicy.cc
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
    add_library(libbrillo-policy SHARED ${SOURCE_FILES})

    target_include_directories(libbrillo-policy
            PUBLIC
            ${ANDROID_ROOT}/external/libbrillo
    )

    # Add c/c++ flags
    target_compile_options(libbrillo-policy PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(libbrillo-policy
            libc
            libc++
            libc++demangle
            libchrome
            libclang_rt.builtins
            libdl
            libgtest_prod_headers
            libm
    )
endif ()
#################### libbrillo-policy ####################

#################### libbrillo-minijail ####################
set(FILES
        ${ANDROID_ROOT}/external/libbrillo/brillo/minijail/minijail.cc
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
    add_library(libbrillo-minijail SHARED ${SOURCE_FILES})

    target_include_directories(libbrillo-minijail
            PUBLIC
            ${ANDROID_ROOT}/external/libbrillo
    )

    # Add c/c++ flags
    target_compile_options(libbrillo-minijail PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(libbrillo-minijail
            libbrillo
            libc
            libc++
            libc++demangle
            libchrome
            libclang_rt.builtins
            libdl
            libgtest_prod_headers
            libm
            libminijail
    )
endif ()
#################### libbrillo-minijail ####################

#################### libbrillo-stream ####################
set(FILES
        ${ANDROID_ROOT}/external/libbrillo/brillo/streams/file_stream.cc
        ${ANDROID_ROOT}/external/libbrillo/brillo/streams/input_stream_set.cc
        ${ANDROID_ROOT}/external/libbrillo/brillo/streams/memory_containers.cc
        ${ANDROID_ROOT}/external/libbrillo/brillo/streams/memory_stream.cc
        ${ANDROID_ROOT}/external/libbrillo/brillo/streams/openssl_stream_bio.cc
        ${ANDROID_ROOT}/external/libbrillo/brillo/streams/stream.cc
        ${ANDROID_ROOT}/external/libbrillo/brillo/streams/stream_errors.cc
        ${ANDROID_ROOT}/external/libbrillo/brillo/streams/stream_utils.cc
        ${ANDROID_ROOT}/external/libbrillo/brillo/streams/tls_stream.cc
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
    add_library(libbrillo-stream SHARED ${SOURCE_FILES})

    target_include_directories(libbrillo-stream
            PUBLIC
            ${ANDROID_ROOT}/external/libbrillo
    )

    # Add c/c++ flags
    target_compile_options(libbrillo-stream PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(libbrillo-stream
            libbrillo
            libc
            libc++
            libc++_static
            libc++demangle
            libchrome
            libclang_rt.builtins
            libcrypto
            libdl
            libgtest_prod_headers
            libm
            libssl
    )
endif ()
#################### libbrillo-stream ####################
