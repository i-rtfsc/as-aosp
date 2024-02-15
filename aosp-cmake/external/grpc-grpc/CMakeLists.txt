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

project(libgrpc_transport_chttp2)

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

#################### libgrpc_transport_chttp2 ####################
set(FILES
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/transport/chttp2/transport/bin_decoder.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/transport/chttp2/transport/bin_encoder.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/transport/chttp2/transport/chttp2_plugin.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/transport/chttp2/transport/chttp2_transport.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/transport/chttp2/transport/flow_control.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/transport/chttp2/transport/frame_data.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/transport/chttp2/transport/frame_goaway.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/transport/chttp2/transport/frame_ping.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/transport/chttp2/transport/frame_rst_stream.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/transport/chttp2/transport/frame_settings.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/transport/chttp2/transport/frame_window_update.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/transport/chttp2/transport/hpack_encoder.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/transport/chttp2/transport/hpack_parser.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/transport/chttp2/transport/hpack_mapping.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/transport/chttp2/transport/hpack_table.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/transport/chttp2/transport/http2_settings.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/transport/chttp2/transport/huffsyms.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/transport/chttp2/transport/incoming_metadata.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/transport/chttp2/transport/parsing.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/transport/chttp2/transport/stream_lists.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/transport/chttp2/transport/stream_map.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/transport/chttp2/transport/varint.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/transport/chttp2/transport/writing.cc
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
    add_library(libgrpc_transport_chttp2 SHARED ${SOURCE_FILES})

    target_include_directories(libgrpc_transport_chttp2
            PUBLIC
            ${ANDROID_ROOT}/external/grpc-grpc/include
            ${ANDROID_ROOT}/external/grpc-grpc
    )

    # Add c/c++ flags
    target_compile_options(libgrpc_transport_chttp2 PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(libgrpc_transport_chttp2
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libgrpc_base
            libgrpc_http_filters
            libgrpc_transport_chttp2_alpn
            libm
            libz
    )
endif ()
#################### libgrpc_transport_chttp2 ####################

#################### libgrpc++_base_unsecure ####################
set(FILES
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/client/channel_cc.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/client/client_context.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/client/create_channel.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/client/create_channel_internal.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/client/create_channel_posix.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/client/credentials_cc.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/client/generic_stub.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/common/alarm.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/common/channel_arguments.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/common/channel_filter.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/common/completion_queue_cc.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/common/core_codegen.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/common/resource_quota_cc.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/common/rpc_method.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/common/version_cc.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/server/async_generic_service.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/server/channel_argument_option.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/server/create_default_thread_pool.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/server/dynamic_thread_pool.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/server/health/default_health_check_service.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/server/health/health.pb.c
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/server/health/health_check_service.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/server/health/health_check_service_server_builder_option.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/server/server_builder.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/server/server_cc.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/server/server_context.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/server/server_credentials.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/server/server_posix.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/thread_manager/thread_manager.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/util/byte_buffer_cc.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/util/status.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/util/string_ref.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/util/time_cc.cc
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
    add_library(libgrpc++_base_unsecure SHARED ${SOURCE_FILES})

    target_include_directories(libgrpc++_base_unsecure
            PUBLIC
            ${ANDROID_ROOT}/external/grpc-grpc/include
            ${ANDROID_ROOT}/external/grpc-grpc
    )

    # Add c/c++ flags
    target_compile_options(libgrpc++_base_unsecure PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
            -Wno-unused-parameter
    )

    target_link_libraries(libgrpc++_base_unsecure
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libgrpc++_codegen_base_src
            libgrpc_unsecure
            libm
            libz
    )
endif ()
#################### libgrpc++_base_unsecure ####################

#################### libgrpc_lb_policy_grpclb_secure ####################
set(FILES
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/filters/client_channel/lb_policy/grpclb/client_load_reporting_filter.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/filters/client_channel/lb_policy/grpclb/grpclb.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/filters/client_channel/lb_policy/grpclb/grpclb_channel_secure.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/filters/client_channel/lb_policy/grpclb/grpclb_client_stats.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/filters/client_channel/lb_policy/grpclb/load_balancer_api.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/filters/client_channel/lb_policy/grpclb/proto/grpc/lb/v1/google/protobuf/duration.pb.c
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/filters/client_channel/lb_policy/grpclb/proto/grpc/lb/v1/google/protobuf/timestamp.pb.c
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/filters/client_channel/lb_policy/grpclb/proto/grpc/lb/v1/load_balancer.pb.c
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
    add_library(libgrpc_lb_policy_grpclb_secure SHARED ${SOURCE_FILES})

    target_include_directories(libgrpc_lb_policy_grpclb_secure
            PUBLIC
            ${ANDROID_ROOT}/external/grpc-grpc/include
            ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/filters/client_channel/lb_policy/grpclb
            ${ANDROID_ROOT}/external/grpc-grpc
    )

    # Add c/c++ flags
    target_compile_options(libgrpc_lb_policy_grpclb_secure PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(libgrpc_lb_policy_grpclb_secure
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libgrpc_base
            libgrpc_client_channel
            libgrpc_deadline_filter
            libgrpc_resolver_fake
            libm
            libprotobuf-c-nano
            libz
    )
endif ()
#################### libgrpc_lb_policy_grpclb_secure ####################

#################### libgrpc_resolver_fake ####################
set(FILES
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/filters/client_channel/resolver/fake/fake_resolver.cc
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
    add_library(libgrpc_resolver_fake SHARED ${SOURCE_FILES})

    target_include_directories(libgrpc_resolver_fake
            PUBLIC
            ${ANDROID_ROOT}/external/grpc-grpc/include
            ${ANDROID_ROOT}/external/grpc-grpc
    )

    # Add c/c++ flags
    target_compile_options(libgrpc_resolver_fake PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(libgrpc_resolver_fake
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libgrpc_base
            libgrpc_client_channel
            libm
            libz
    )
endif ()
#################### libgrpc_resolver_fake ####################

#################### libgrpc_resolver_sockaddr ####################
set(FILES
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/filters/client_channel/resolver/sockaddr/sockaddr_resolver.cc
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
    add_library(libgrpc_resolver_sockaddr SHARED ${SOURCE_FILES})

    target_include_directories(libgrpc_resolver_sockaddr
            PUBLIC
            ${ANDROID_ROOT}/external/grpc-grpc/include
            ${ANDROID_ROOT}/external/grpc-grpc
    )

    # Add c/c++ flags
    target_compile_options(libgrpc_resolver_sockaddr PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(libgrpc_resolver_sockaddr
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libgrpc_base
            libgrpc_client_channel
            libm
            libz
    )
endif ()
#################### libgrpc_resolver_sockaddr ####################

#################### libgrpc_base ####################
set(FILES
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/surface/lame_client.cc
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
    add_library(libgrpc_base SHARED ${SOURCE_FILES})

    target_include_directories(libgrpc_base
            PUBLIC
            ${ANDROID_ROOT}/external/grpc-grpc
            ${ANDROID_ROOT}/external/grpc-grpc/include
    )

    # Add c/c++ flags
    target_compile_options(libgrpc_base PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(libgrpc_base
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libgrpc_base_c
            liblog
            libm
            libz
    )
endif ()
#################### libgrpc_base ####################

#################### libgrpc_max_age_filter ####################
set(FILES
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/filters/max_age/max_age_filter.cc
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
    add_library(libgrpc_max_age_filter SHARED ${SOURCE_FILES})

    target_include_directories(libgrpc_max_age_filter
            PUBLIC
            ${ANDROID_ROOT}/external/grpc-grpc/include
            ${ANDROID_ROOT}/external/grpc-grpc
    )

    # Add c/c++ flags
    target_compile_options(libgrpc_max_age_filter PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(libgrpc_max_age_filter
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libgrpc_base
            libm
            libz
    )
endif ()
#################### libgrpc_max_age_filter ####################

#################### libgrpc_transport_chttp2_client_insecure ####################
set(FILES
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/transport/chttp2/client/insecure/channel_create.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/transport/chttp2/client/insecure/channel_create_posix.cc
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
    add_library(libgrpc_transport_chttp2_client_insecure SHARED ${SOURCE_FILES})

    target_include_directories(libgrpc_transport_chttp2_client_insecure
            PUBLIC
            ${ANDROID_ROOT}/external/grpc-grpc/include
            ${ANDROID_ROOT}/external/grpc-grpc
    )

    # Add c/c++ flags
    target_compile_options(libgrpc_transport_chttp2_client_insecure PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(libgrpc_transport_chttp2_client_insecure
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libgrpc_base
            libgrpc_client_channel
            libgrpc_deadline_filter
            libgrpc_transport_chttp2
            libgrpc_transport_chttp2_client_connector
            libm
            libprotobuf-c-nano
            libz
    )
endif ()
#################### libgrpc_transport_chttp2_client_insecure ####################

#################### libgrpc_http_filters ####################
set(FILES
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/filters/http/client/http_client_filter.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/filters/http/http_filters_plugin.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/filters/http/message_compress/message_compress_filter.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/filters/http/server/http_server_filter.cc
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
    add_library(libgrpc_http_filters SHARED ${SOURCE_FILES})

    target_include_directories(libgrpc_http_filters
            PUBLIC
            ${ANDROID_ROOT}/external/grpc-grpc/include
            ${ANDROID_ROOT}/external/grpc-grpc
    )

    # Add c/c++ flags
    target_compile_options(libgrpc_http_filters PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(libgrpc_http_filters
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libgrpc_base
            libgrpc_client_channel
            libm
            libprotobuf-c-nano
            libz
    )
endif ()
#################### libgrpc_http_filters ####################

#################### libgrpc_workaround_cronet_compression_filter ####################
set(FILES
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/filters/workarounds/workaround_cronet_compression_filter.cc
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
    add_library(libgrpc_workaround_cronet_compression_filter SHARED ${SOURCE_FILES})

    target_include_directories(libgrpc_workaround_cronet_compression_filter
            PUBLIC
            ${ANDROID_ROOT}/external/grpc-grpc/include
            ${ANDROID_ROOT}/external/grpc-grpc
    )

    # Add c/c++ flags
    target_compile_options(libgrpc_workaround_cronet_compression_filter PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(libgrpc_workaround_cronet_compression_filter
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libgrpc_base
            libgrpc_server_backward_compatibility
            libm
            libprotobuf-c-nano
            libz
    )
endif ()
#################### libgrpc_workaround_cronet_compression_filter ####################

#################### libgrpc_transport_chttp2_alpn ####################
set(FILES
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/transport/chttp2/alpn/alpn.cc
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
    add_library(libgrpc_transport_chttp2_alpn SHARED ${SOURCE_FILES})

    target_include_directories(libgrpc_transport_chttp2_alpn
            PUBLIC
            ${ANDROID_ROOT}/external/grpc-grpc
            ${ANDROID_ROOT}/external/grpc-grpc/include
    )

    # Add c/c++ flags
    target_compile_options(libgrpc_transport_chttp2_alpn PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(libgrpc_transport_chttp2_alpn
            gpr_base
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            liblog
            libm
            libz
    )
endif ()
#################### libgrpc_transport_chttp2_alpn ####################

#################### libgrpc++_unsecure ####################
set(FILES
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/client/insecure_credentials.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/common/insecure_create_auth_context.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/server/insecure_server_credentials.cc
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
    add_library(libgrpc++_unsecure SHARED ${SOURCE_FILES})

    target_include_directories(libgrpc++_unsecure
            PUBLIC
            ${ANDROID_ROOT}/external/grpc-grpc/include
            ${ANDROID_ROOT}/external/grpc-grpc
    )

    # Add c/c++ flags
    target_compile_options(libgrpc++_unsecure PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(libgrpc++_unsecure
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libgrpc++_base_unsecure
            libgrpc_unsecure
            liblog
            libm
            libz
    )
endif ()
#################### libgrpc++_unsecure ####################

#################### libgrpc_resolver_dns_ares ####################
set(FILES
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/filters/client_channel/resolver/dns/c_ares/dns_resolver_ares.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/filters/client_channel/resolver/dns/c_ares/grpc_ares_ev_driver.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/filters/client_channel/resolver/dns/c_ares/grpc_ares_ev_driver_posix.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/filters/client_channel/resolver/dns/c_ares/grpc_ares_ev_driver_windows.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/filters/client_channel/resolver/dns/c_ares/grpc_ares_wrapper.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/filters/client_channel/resolver/dns/c_ares/grpc_ares_wrapper_fallback.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/filters/client_channel/resolver/dns/c_ares/grpc_ares_wrapper_posix.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/filters/client_channel/resolver/dns/c_ares/grpc_ares_wrapper_windows.cc
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
    add_library(libgrpc_resolver_dns_ares SHARED ${SOURCE_FILES})

    target_include_directories(libgrpc_resolver_dns_ares
            PUBLIC
            ${ANDROID_ROOT}/external/grpc-grpc/include
            ${ANDROID_ROOT}/external/grpc-grpc/third_party/cares
            ${ANDROID_ROOT}/external/grpc-grpc/third_party/cares/config_android
            ${ANDROID_ROOT}/external/grpc-grpc
    )

    # Add c/c++ flags
    target_compile_options(libgrpc_resolver_dns_ares PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(libgrpc_resolver_dns_ares
            libaddress_sorting
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libgrpc_base
            libgrpc_client_channel
            libm
            libprotobuf-c-nano
            libz
    )
endif ()
#################### libgrpc_resolver_dns_ares ####################

#################### libgrpc_transport_inproc ####################
set(FILES
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/transport/inproc/inproc_plugin.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/transport/inproc/inproc_transport.cc
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
    add_library(libgrpc_transport_inproc SHARED ${SOURCE_FILES})

    target_include_directories(libgrpc_transport_inproc
            PUBLIC
            ${ANDROID_ROOT}/external/grpc-grpc/include
            ${ANDROID_ROOT}/external/grpc-grpc
    )

    # Add c/c++ flags
    target_compile_options(libgrpc_transport_inproc PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(libgrpc_transport_inproc
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libgrpc_base
            libm
            libprotobuf-c-nano
            libz
    )
endif ()
#################### libgrpc_transport_inproc ####################

#################### libgrpc_unsecure ####################
set(FILES
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/surface/init.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/surface/init_unsecure.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/plugin_registry/grpc_unsecure_plugin_registry.cc
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
    add_library(libgrpc_unsecure SHARED ${SOURCE_FILES})

    target_include_directories(libgrpc_unsecure
            PUBLIC
            ${ANDROID_ROOT}/external/grpc-grpc/include
            ${ANDROID_ROOT}/external/grpc-grpc
            ${ANDROID_ROOT}/external/grpc-grpc/third_party/cares
            ${ANDROID_ROOT}/external/grpc-grpc/third_party/cares/config_android
    )

    # Add c/c++ flags
    target_compile_options(libgrpc_unsecure PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(libgrpc_unsecure
            census
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libgrpc_base
            libgrpc_base_c
            libgrpc_client_authority_filter
            libgrpc_client_channel
            libgrpc_deadline_filter
            libgrpc_http_filters
            libgrpc_lb_policy_grpclb
            libgrpc_lb_policy_pick_first
            libgrpc_lb_policy_round_robin
            libgrpc_max_age_filter
            libgrpc_message_size_filter
            libgrpc_resolver_dns_ares
            libgrpc_resolver_dns_native
            libgrpc_resolver_fake
            libgrpc_resolver_sockaddr
            libgrpc_server_backward_compatibility
            libgrpc_transport_chttp2
            libgrpc_transport_chttp2_client_connector
            libgrpc_transport_chttp2_client_insecure
            libgrpc_transport_chttp2_server_insecure
            libgrpc_transport_inproc
            libgrpc_workaround_cronet_compression_filter
            libm
            libprotobuf-c-nano
            libz
    )
endif ()
#################### libgrpc_unsecure ####################

#################### libgrpc_transport_chttp2_server_secure ####################
set(FILES
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/transport/chttp2/server/chttp2_server.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/transport/chttp2/server/secure/server_secure_chttp2.cc
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
    add_library(libgrpc_transport_chttp2_server_secure SHARED ${SOURCE_FILES})

    target_include_directories(libgrpc_transport_chttp2_server_secure
            PUBLIC
            ${ANDROID_ROOT}/external/grpc-grpc/include
            ${ANDROID_ROOT}/external/grpc-grpc
    )

    # Add c/c++ flags
    target_compile_options(libgrpc_transport_chttp2_server_secure PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(libgrpc_transport_chttp2_server_secure
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libgrpc_base
            libgrpc_transport_chttp2
            libm
            libprotobuf-c-nano
            libz
    )
endif ()
#################### libgrpc_transport_chttp2_server_secure ####################

#################### libgrpc_lb_policy_grpclb ####################
set(FILES
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/filters/client_channel/lb_policy/grpclb/client_load_reporting_filter.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/filters/client_channel/lb_policy/grpclb/grpclb.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/filters/client_channel/lb_policy/grpclb/grpclb_channel.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/filters/client_channel/lb_policy/grpclb/grpclb_client_stats.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/filters/client_channel/lb_policy/grpclb/load_balancer_api.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/filters/client_channel/lb_policy/grpclb/proto/grpc/lb/v1/google/protobuf/duration.pb.c
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/filters/client_channel/lb_policy/grpclb/proto/grpc/lb/v1/google/protobuf/timestamp.pb.c
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/filters/client_channel/lb_policy/grpclb/proto/grpc/lb/v1/load_balancer.pb.c
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
    add_library(libgrpc_lb_policy_grpclb SHARED ${SOURCE_FILES})

    target_include_directories(libgrpc_lb_policy_grpclb
            PUBLIC
            ${ANDROID_ROOT}/external/grpc-grpc/include
            ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/filters/client_channel/lb_policy/grpclb
            ${ANDROID_ROOT}/external/grpc-grpc
    )

    # Add c/c++ flags
    target_compile_options(libgrpc_lb_policy_grpclb PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(libgrpc_lb_policy_grpclb
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libgrpc_base
            libgrpc_client_channel
            libgrpc_deadline_filter
            libgrpc_resolver_fake
            libm
            libprotobuf-c-nano
            libz
    )
endif ()
#################### libgrpc_lb_policy_grpclb ####################

#################### libgrpc_resolver_dns_native ####################
set(FILES
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/filters/client_channel/resolver/dns/native/dns_resolver.cc
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
    add_library(libgrpc_resolver_dns_native SHARED ${SOURCE_FILES})

    target_include_directories(libgrpc_resolver_dns_native
            PUBLIC
            ${ANDROID_ROOT}/external/grpc-grpc/include
            ${ANDROID_ROOT}/external/grpc-grpc
    )

    # Add c/c++ flags
    target_compile_options(libgrpc_resolver_dns_native PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(libgrpc_resolver_dns_native
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libgrpc_base
            libgrpc_client_channel
            libm
            libz
    )
endif ()
#################### libgrpc_resolver_dns_native ####################

#################### libgrpc_deadline_filter ####################
set(FILES
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/filters/deadline/deadline_filter.cc
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
    add_library(libgrpc_deadline_filter SHARED ${SOURCE_FILES})

    target_include_directories(libgrpc_deadline_filter
            PUBLIC
            ${ANDROID_ROOT}/external/grpc-grpc/include
            ${ANDROID_ROOT}/external/grpc-grpc
    )

    # Add c/c++ flags
    target_compile_options(libgrpc_deadline_filter PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(libgrpc_deadline_filter
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libgrpc_base
            libm
            libz
    )
endif ()
#################### libgrpc_deadline_filter ####################

#################### libgrpc_transport_chttp2_client_connector ####################
set(FILES
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/transport/chttp2/client/authority.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/transport/chttp2/client/chttp2_connector.cc
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
    add_library(libgrpc_transport_chttp2_client_connector SHARED ${SOURCE_FILES})

    target_include_directories(libgrpc_transport_chttp2_client_connector
            PUBLIC
            ${ANDROID_ROOT}/external/grpc-grpc/include
            ${ANDROID_ROOT}/external/grpc-grpc
    )

    # Add c/c++ flags
    target_compile_options(libgrpc_transport_chttp2_client_connector PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(libgrpc_transport_chttp2_client_connector
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libgrpc_base
            libgrpc_client_channel
            libgrpc_deadline_filter
            libgrpc_transport_chttp2
            libm
            libprotobuf-c-nano
            libz
    )
endif ()
#################### libgrpc_transport_chttp2_client_connector ####################

#################### libgrpc_lb_policy_pick_first ####################
set(FILES
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/filters/client_channel/lb_policy/pick_first/pick_first.cc
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
    add_library(libgrpc_lb_policy_pick_first SHARED ${SOURCE_FILES})

    target_include_directories(libgrpc_lb_policy_pick_first
            PUBLIC
            ${ANDROID_ROOT}/external/grpc-grpc/include
            ${ANDROID_ROOT}/external/grpc-grpc
    )

    # Add c/c++ flags
    target_compile_options(libgrpc_lb_policy_pick_first PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(libgrpc_lb_policy_pick_first
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libgrpc_base
            libgrpc_client_channel
            libgrpc_deadline_filter
            libm
            libz
    )
endif ()
#################### libgrpc_lb_policy_pick_first ####################

#################### libgrpc_client_authority_filter ####################
set(FILES
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/filters/http/client_authority_filter.cc
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
    add_library(libgrpc_client_authority_filter SHARED ${SOURCE_FILES})

    target_include_directories(libgrpc_client_authority_filter
            PUBLIC
            ${ANDROID_ROOT}/external/grpc-grpc/include
            ${ANDROID_ROOT}/external/grpc-grpc
    )

    # Add c/c++ flags
    target_compile_options(libgrpc_client_authority_filter PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(libgrpc_client_authority_filter
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libgrpc_base
            libm
            libz
    )
endif ()
#################### libgrpc_client_authority_filter ####################

#################### libalts_frame_protector ####################
set(FILES
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/tsi/alts/crypt/aes_gcm.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/tsi/alts/crypt/gsec.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/tsi/alts/frame_protector/alts_counter.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/tsi/alts/frame_protector/alts_crypter.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/tsi/alts/frame_protector/alts_frame_protector.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/tsi/alts/frame_protector/alts_record_protocol_crypter_common.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/tsi/alts/frame_protector/alts_seal_privacy_integrity_crypter.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/tsi/alts/frame_protector/alts_unseal_privacy_integrity_crypter.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/tsi/alts/frame_protector/frame_handler.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/tsi/alts/zero_copy_frame_protector/alts_grpc_integrity_only_record_protocol.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/tsi/alts/zero_copy_frame_protector/alts_grpc_privacy_integrity_record_protocol.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/tsi/alts/zero_copy_frame_protector/alts_grpc_record_protocol_common.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/tsi/alts/zero_copy_frame_protector/alts_iovec_record_protocol.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/tsi/alts/zero_copy_frame_protector/alts_zero_copy_grpc_protector.cc
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
    add_library(libalts_frame_protector SHARED ${SOURCE_FILES})

    target_include_directories(libalts_frame_protector
            PUBLIC
            ${ANDROID_ROOT}/external/grpc-grpc/include
            ${ANDROID_ROOT}/external/grpc-grpc
    )

    # Add c/c++ flags
    target_compile_options(libalts_frame_protector PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(libalts_frame_protector
            gpr_base
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcrypto
            libdl
            libgrpc_base
            libm
            libssl
            libtsi_interface
            libz
    )
endif ()
#################### libalts_frame_protector ####################

#################### libgrpc++_base ####################
set(FILES
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/client/channel_cc.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/client/client_context.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/client/create_channel.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/client/create_channel_internal.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/client/create_channel_posix.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/client/credentials_cc.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/client/generic_stub.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/common/alarm.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/common/channel_arguments.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/common/channel_filter.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/common/completion_queue_cc.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/common/core_codegen.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/common/resource_quota_cc.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/common/rpc_method.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/common/version_cc.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/server/async_generic_service.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/server/channel_argument_option.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/server/create_default_thread_pool.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/server/dynamic_thread_pool.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/server/health/default_health_check_service.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/server/health/health.pb.c
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/server/health/health_check_service.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/server/health/health_check_service_server_builder_option.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/server/server_builder.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/server/server_cc.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/server/server_context.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/server/server_credentials.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/server/server_posix.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/thread_manager/thread_manager.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/util/byte_buffer_cc.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/util/status.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/util/string_ref.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/util/time_cc.cc
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
    add_library(libgrpc++_base SHARED ${SOURCE_FILES})

    target_include_directories(libgrpc++_base
            PUBLIC
            ${ANDROID_ROOT}/external/grpc-grpc/include
            ${ANDROID_ROOT}/external/grpc-grpc
    )

    # Add c/c++ flags
    target_compile_options(libgrpc++_base PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
            -Wno-unused-parameter
    )

    target_link_libraries(libgrpc++_base
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libgrpc
            libgrpc++_codegen_base_src
            libm
            libz
    )
endif ()
#################### libgrpc++_base ####################

#################### libgrpc_lb_policy_round_robin ####################
set(FILES
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/filters/client_channel/lb_policy/round_robin/round_robin.cc
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
    add_library(libgrpc_lb_policy_round_robin SHARED ${SOURCE_FILES})

    target_include_directories(libgrpc_lb_policy_round_robin
            PUBLIC
            ${ANDROID_ROOT}/external/grpc-grpc/include
            ${ANDROID_ROOT}/external/grpc-grpc
    )

    # Add c/c++ flags
    target_compile_options(libgrpc_lb_policy_round_robin PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(libgrpc_lb_policy_round_robin
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libgrpc_base
            libgrpc_client_channel
            libgrpc_deadline_filter
            libm
            libz
    )
endif ()
#################### libgrpc_lb_policy_round_robin ####################

#################### gpr_base ####################
set(FILES
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/gpr/alloc.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/gpr/arena.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/gpr/atm.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/gpr/cpu_iphone.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/gpr/cpu_linux.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/gpr/cpu_posix.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/gpr/cpu_windows.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/gpr/env_linux.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/gpr/env_posix.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/gpr/env_windows.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/gpr/host_port.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/gpr/log.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/gpr/log_android.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/gpr/log_linux.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/gpr/log_posix.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/gpr/log_windows.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/gpr/mpscq.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/gpr/murmur_hash.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/gpr/string.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/gpr/string_posix.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/gpr/string_util_windows.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/gpr/string_windows.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/gpr/sync.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/gpr/sync_posix.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/gpr/sync_windows.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/gpr/time.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/gpr/time_posix.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/gpr/time_precise.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/gpr/time_windows.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/gpr/tls_pthread.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/gpr/tmpfile_msys.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/gpr/tmpfile_posix.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/gpr/tmpfile_windows.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/gpr/wrap_memcpy.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/gprpp/fork.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/gprpp/thd_posix.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/gprpp/thd_windows.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/profiling/basic_timers.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/profiling/stap_timers.cc
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
    add_library(gpr_base SHARED ${SOURCE_FILES})

    target_include_directories(gpr_base
            PUBLIC
            ${ANDROID_ROOT}/external/grpc-grpc/include
            ${ANDROID_ROOT}/external/grpc-grpc
    )

    # Add c/c++ flags
    target_compile_options(gpr_base PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(gpr_base
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            liblog
            libm
            libz
    )
endif ()
#################### gpr_base ####################

#################### libalts_util ####################
set(FILES
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/security/credentials/alts/check_gcp_environment.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/security/credentials/alts/check_gcp_environment_linux.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/security/credentials/alts/check_gcp_environment_no_op.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/security/credentials/alts/check_gcp_environment_windows.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/security/credentials/alts/grpc_alts_credentials_client_options.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/security/credentials/alts/grpc_alts_credentials_options.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/security/credentials/alts/grpc_alts_credentials_server_options.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/tsi/alts/handshaker/alts_handshaker_service_api.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/tsi/alts/handshaker/altscontext.pb.c
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/tsi/alts/handshaker/alts_handshaker_service_api_util.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/tsi/alts/handshaker/handshaker.pb.c
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/tsi/alts/handshaker/transport_security_common_api.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/tsi/alts/handshaker/transport_security_common.pb.c
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
    add_library(libalts_util SHARED ${SOURCE_FILES})

    target_include_directories(libalts_util
            PUBLIC
            ${ANDROID_ROOT}/external/grpc-grpc/include
            ${ANDROID_ROOT}/external/grpc-grpc
    )

    # Add c/c++ flags
    target_compile_options(libalts_util PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(libalts_util
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libprotobuf-c-nano
            libz
    )
endif ()
#################### libalts_util ####################

#################### libtsi ####################
set(FILES
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/tsi/alts/handshaker/alts_handshaker_client.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/tsi/alts/handshaker/alts_tsi_event.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/tsi/alts/handshaker/alts_tsi_handshaker.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/tsi/alts/handshaker/alts_tsi_utils.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/tsi/alts_transport_security.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/tsi/fake_transport_security.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/tsi/local_transport_security.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/tsi/ssl/session_cache/ssl_session_boringssl.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/tsi/ssl/session_cache/ssl_session_cache.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/tsi/ssl/session_cache/ssl_session_openssl.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/tsi/ssl_transport_security.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/tsi/transport_security_grpc.cc
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
    add_library(libtsi SHARED ${SOURCE_FILES})

    target_include_directories(libtsi
            PUBLIC
            ${ANDROID_ROOT}/external/grpc-grpc/include
            ${ANDROID_ROOT}/external/grpc-grpc/third_party/nanopb
            ${ANDROID_ROOT}/external/grpc-grpc
    )

    # Add c/c++ flags
    target_compile_options(libtsi PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(libtsi
            gpr_base
            libalts_frame_protector
            libalts_util
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcrypto
            libdl
            libgrpc_base
            libgrpc_transport_chttp2_client_insecure
            libm
            libssl
            libz
    )
endif ()
#################### libtsi ####################

#################### libgrpc_server_backward_compatibility ####################
set(FILES
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/filters/workarounds/workaround_utils.cc
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
    add_library(libgrpc_server_backward_compatibility SHARED ${SOURCE_FILES})

    target_include_directories(libgrpc_server_backward_compatibility
            PUBLIC
            ${ANDROID_ROOT}/external/grpc-grpc/include
            ${ANDROID_ROOT}/external/grpc-grpc
    )

    # Add c/c++ flags
    target_compile_options(libgrpc_server_backward_compatibility PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(libgrpc_server_backward_compatibility
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libgrpc_base
            libm
            libprotobuf-c-nano
            libz
    )
endif ()
#################### libgrpc_server_backward_compatibility ####################

#################### libgrpc_trace ####################
set(FILES
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/debug/trace.cc
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
    add_library(libgrpc_trace SHARED ${SOURCE_FILES})

    target_include_directories(libgrpc_trace
            PUBLIC
            ${ANDROID_ROOT}/external/grpc-grpc
            ${ANDROID_ROOT}/external/grpc-grpc/include
    )

    # Add c/c++ flags
    target_compile_options(libgrpc_trace PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(libgrpc_trace
            gpr_base
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            liblog
            libm
            libz
    )
endif ()
#################### libgrpc_trace ####################

#################### libgrpc_base_c ####################
set(FILES
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/avl/avl.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/backoff/backoff.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/channel/channel_args.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/channel/channel_stack.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/channel/channel_stack_builder.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/channel/channel_trace.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/channel/channelz.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/channel/channelz_registry.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/channel/connected_channel.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/channel/handshaker.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/channel/handshaker_factory.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/channel/handshaker_registry.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/channel/status_util.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/compression/compression.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/compression/compression_internal.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/compression/message_compress.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/compression/stream_compression.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/compression/stream_compression_gzip.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/compression/stream_compression_identity.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/debug/stats.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/debug/stats_data.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/http/format_request.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/http/httpcli.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/http/parser.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/buffer_list.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/call_combiner.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/combiner.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/endpoint.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/endpoint_pair_posix.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/endpoint_pair_uv.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/endpoint_pair_windows.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/error.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/ev_epoll1_linux.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/ev_epollex_linux.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/ev_epollsig_linux.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/ev_poll_posix.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/ev_posix.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/ev_windows.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/exec_ctx.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/executor.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/fork_posix.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/fork_windows.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/gethostname_fallback.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/gethostname_host_name_max.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/gethostname_sysconf.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/internal_errqueue.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/iocp_windows.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/iomgr.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/iomgr_custom.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/iomgr_internal.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/iomgr_posix.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/iomgr_windows.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/is_epollexclusive_available.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/load_file.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/lockfree_event.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/network_status_tracker.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/polling_entity.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/pollset.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/pollset_custom.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/pollset_set.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/pollset_set_custom.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/pollset_set_windows.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/pollset_uv.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/pollset_windows.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/resolve_address.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/resolve_address_custom.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/resolve_address_posix.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/resolve_address_windows.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/resource_quota.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/sockaddr_utils.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/socket_factory_posix.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/socket_mutator.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/socket_utils_common_posix.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/socket_utils_linux.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/socket_utils_posix.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/socket_utils_windows.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/socket_windows.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/tcp_client.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/tcp_client_custom.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/tcp_client_posix.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/tcp_client_windows.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/tcp_custom.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/tcp_posix.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/tcp_server.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/tcp_server_custom.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/tcp_server_posix.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/tcp_server_utils_posix_common.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/tcp_server_utils_posix_ifaddrs.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/tcp_server_utils_posix_noifaddrs.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/tcp_server_windows.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/tcp_uv.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/tcp_windows.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/time_averaged_stats.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/timer.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/timer_custom.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/timer_generic.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/timer_heap.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/timer_manager.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/timer_uv.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/udp_server.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/unix_sockets_posix.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/unix_sockets_posix_noop.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/wakeup_fd_cv.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/wakeup_fd_eventfd.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/wakeup_fd_nospecial.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/wakeup_fd_pipe.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/iomgr/wakeup_fd_posix.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/json/json.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/json/json_reader.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/json/json_string.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/json/json_writer.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/slice/b64.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/slice/percent_encoding.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/slice/slice.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/slice/slice_buffer.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/slice/slice_intern.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/slice/slice_string_helpers.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/surface/api_trace.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/surface/byte_buffer.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/surface/byte_buffer_reader.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/surface/call.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/surface/call_details.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/surface/call_log_batch.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/surface/channel.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/surface/channel_init.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/surface/channel_ping.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/surface/channel_stack_type.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/surface/completion_queue.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/surface/completion_queue_factory.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/surface/event_string.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/surface/metadata_array.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/surface/server.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/surface/validate_metadata.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/surface/version.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/transport/bdp_estimator.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/transport/byte_stream.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/transport/connectivity_state.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/transport/error_utils.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/transport/metadata.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/transport/metadata_batch.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/transport/pid_controller.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/transport/service_config.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/transport/static_metadata.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/transport/status_conversion.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/transport/status_metadata.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/transport/timeout_encoding.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/transport/transport.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/transport/transport_op_string.cc
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
    add_library(libgrpc_base_c SHARED ${SOURCE_FILES})

    target_include_directories(libgrpc_base_c
            PUBLIC
            ${ANDROID_ROOT}/external/grpc-grpc
            ${ANDROID_ROOT}/external/grpc-grpc/include
    )

    # Add c/c++ flags
    target_compile_options(libgrpc_base_c PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(libgrpc_base_c
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libgrpc_trace
            liblog
            libm
            libz
    )
endif ()
#################### libgrpc_base_c ####################

#################### libtsi_interface ####################
set(FILES
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/tsi/transport_security.cc
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
    add_library(libtsi_interface SHARED ${SOURCE_FILES})

    target_include_directories(libtsi_interface
            PUBLIC
            ${ANDROID_ROOT}/external/grpc-grpc
            ${ANDROID_ROOT}/external/grpc-grpc/include
    )

    # Add c/c++ flags
    target_compile_options(libtsi_interface PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(libtsi_interface
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libgrpc_trace
            libm
            libz
    )
endif ()
#################### libtsi_interface ####################

#################### libgrpc_client_channel ####################
set(FILES
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/filters/client_channel/backup_poller.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/filters/client_channel/channel_connectivity.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/filters/client_channel/client_channel.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/filters/client_channel/client_channel_channelz.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/filters/client_channel/client_channel_factory.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/filters/client_channel/client_channel_plugin.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/filters/client_channel/connector.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/filters/client_channel/http_connect_handshaker.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/filters/client_channel/http_proxy.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/filters/client_channel/lb_policy.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/filters/client_channel/lb_policy_factory.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/filters/client_channel/lb_policy_registry.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/filters/client_channel/method_params.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/filters/client_channel/parse_address.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/filters/client_channel/proxy_mapper.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/filters/client_channel/proxy_mapper_registry.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/filters/client_channel/resolver.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/filters/client_channel/resolver_registry.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/filters/client_channel/retry_throttle.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/filters/client_channel/subchannel.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/filters/client_channel/subchannel_index.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/filters/client_channel/uri_parser.cc
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
    add_library(libgrpc_client_channel SHARED ${SOURCE_FILES})

    target_include_directories(libgrpc_client_channel
            PUBLIC
            ${ANDROID_ROOT}/external/grpc-grpc/include
            ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/filters/client_channel
            ${ANDROID_ROOT}/external/grpc-grpc
    )

    # Add c/c++ flags
    target_compile_options(libgrpc_client_channel PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(libgrpc_client_channel
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libgrpc_base
            libgrpc_deadline_filter
            libm
            libz
    )
endif ()
#################### libgrpc_client_channel ####################

#################### libgrpc++_codegen_base_src ####################
set(FILES
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/codegen/codegen_init.cc
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
    add_library(libgrpc++_codegen_base_src SHARED ${SOURCE_FILES})

    target_include_directories(libgrpc++_codegen_base_src
            PUBLIC
            ${ANDROID_ROOT}/external/grpc-grpc/include
            ${ANDROID_ROOT}/external/grpc-grpc
    )

    # Add c/c++ flags
    target_compile_options(libgrpc++_codegen_base_src PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(libgrpc++_codegen_base_src
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libz
    )
endif ()
#################### libgrpc++_codegen_base_src ####################

#################### census ####################
set(FILES
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/filters/census/grpc_context.cc
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
    add_library(census SHARED ${SOURCE_FILES})

    target_include_directories(census
            PUBLIC
            ${ANDROID_ROOT}/external/grpc-grpc/include
            ${ANDROID_ROOT}/external/grpc-grpc
    )

    # Add c/c++ flags
    target_compile_options(census PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(census
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libgrpc_base
            libm
            libz
    )
endif ()
#################### census ####################

#################### libgrpc_secure ####################
set(FILES
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/http/httpcli_security_connector.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/security/context/security_context.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/security/credentials/alts/alts_credentials.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/security/credentials/composite/composite_credentials.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/security/credentials/credentials.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/security/credentials/credentials_metadata.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/security/credentials/fake/fake_credentials.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/security/credentials/google_default/credentials_generic.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/security/credentials/google_default/google_default_credentials.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/security/credentials/iam/iam_credentials.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/security/credentials/jwt/json_token.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/security/credentials/jwt/jwt_credentials.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/security/credentials/jwt/jwt_verifier.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/security/credentials/local/local_credentials.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/security/credentials/oauth2/oauth2_credentials.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/security/credentials/plugin/plugin_credentials.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/security/credentials/ssl/ssl_credentials.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/security/security_connector/alts_security_connector.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/security/security_connector/load_system_roots_fallback.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/security/security_connector/load_system_roots_linux.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/security/security_connector/local_security_connector.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/security/security_connector/security_connector.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/security/transport/client_auth_filter.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/security/transport/secure_endpoint.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/security/transport/security_handshaker.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/security/transport/server_auth_filter.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/security/transport/target_authority_table.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/security/transport/tsi_error.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/security/util/json_util.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/surface/init_secure.cc
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
    add_library(libgrpc_secure SHARED ${SOURCE_FILES})

    target_include_directories(libgrpc_secure
            PUBLIC
            ${ANDROID_ROOT}/external/grpc-grpc/include
            ${ANDROID_ROOT}/external/grpc-grpc/third_party/nanopb
            ${ANDROID_ROOT}/external/grpc-grpc
    )

    # Add c/c++ flags
    target_compile_options(libgrpc_secure PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(libgrpc_secure
            libalts_frame_protector
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcrypto
            libdl
            libgrpc_base
            libgrpc_transport_chttp2_alpn
            libm
            libssl
            libtsi
            libz
    )
endif ()
#################### libgrpc_secure ####################

#################### libgrpc_transport_chttp2_server_insecure ####################
set(FILES
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/transport/chttp2/server/chttp2_server.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/transport/chttp2/server/insecure/server_chttp2.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/transport/chttp2/server/insecure/server_chttp2_posix.cc
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
    add_library(libgrpc_transport_chttp2_server_insecure SHARED ${SOURCE_FILES})

    target_include_directories(libgrpc_transport_chttp2_server_insecure
            PUBLIC
            ${ANDROID_ROOT}/external/grpc-grpc/include
            ${ANDROID_ROOT}/external/grpc-grpc
    )

    # Add c/c++ flags
    target_compile_options(libgrpc_transport_chttp2_server_insecure PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(libgrpc_transport_chttp2_server_insecure
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libgrpc_base
            libgrpc_transport_chttp2
            libm
            libprotobuf-c-nano
            libz
    )
endif ()
#################### libgrpc_transport_chttp2_server_insecure ####################

#################### libgrpc++ ####################
set(FILES
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/client/insecure_credentials.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/client/secure_credentials.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/common/auth_property_iterator.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/common/secure_auth_context.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/common/secure_channel_arguments.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/common/secure_create_auth_context.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/server/insecure_server_credentials.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/cpp/server/secure_server_credentials.cc
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
    add_library(libgrpc++ SHARED ${SOURCE_FILES})

    target_include_directories(libgrpc++
            PUBLIC
            ${ANDROID_ROOT}/external/grpc-grpc/include
            ${ANDROID_ROOT}/external/grpc-grpc
    )

    # Add c/c++ flags
    target_compile_options(libgrpc++ PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(libgrpc++
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcrypto
            libdl
            libgrpc
            libgrpc++_base
            liblog
            libm
            libssl
            libz
    )
endif ()
#################### libgrpc++ ####################

#################### libgrpc_transport_chttp2_client_secure ####################
set(FILES
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/transport/chttp2/client/secure/secure_channel_create.cc
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
    add_library(libgrpc_transport_chttp2_client_secure SHARED ${SOURCE_FILES})

    target_include_directories(libgrpc_transport_chttp2_client_secure
            PUBLIC
            ${ANDROID_ROOT}/external/grpc-grpc/include
            ${ANDROID_ROOT}/external/grpc-grpc
    )

    # Add c/c++ flags
    target_compile_options(libgrpc_transport_chttp2_client_secure PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(libgrpc_transport_chttp2_client_secure
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libgrpc_base
            libgrpc_client_channel
            libgrpc_deadline_filter
            libgrpc_transport_chttp2
            libgrpc_transport_chttp2_client_connector
            libm
            libprotobuf-c-nano
            libz
    )
endif ()
#################### libgrpc_transport_chttp2_client_secure ####################

#################### libgrpc ####################
set(FILES
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/lib/surface/init.cc
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/plugin_registry/grpc_plugin_registry.cc
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
    add_library(libgrpc SHARED ${SOURCE_FILES})

    target_include_directories(libgrpc
            PUBLIC
            ${ANDROID_ROOT}/external/grpc-grpc/include
            ${ANDROID_ROOT}/external/grpc-grpc
            ${ANDROID_ROOT}/external/grpc-grpc/third_party/cares
            ${ANDROID_ROOT}/external/grpc-grpc/third_party/cares/config_android
    )

    # Add c/c++ flags
    target_compile_options(libgrpc PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(libgrpc
            census
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libgrpc_base
            libgrpc_base_c
            libgrpc_client_authority_filter
            libgrpc_client_channel
            libgrpc_deadline_filter
            libgrpc_http_filters
            libgrpc_lb_policy_grpclb_secure
            libgrpc_lb_policy_pick_first
            libgrpc_lb_policy_round_robin
            libgrpc_max_age_filter
            libgrpc_message_size_filter
            libgrpc_resolver_dns_ares
            libgrpc_resolver_dns_native
            libgrpc_resolver_fake
            libgrpc_resolver_sockaddr
            libgrpc_secure
            libgrpc_server_backward_compatibility
            libgrpc_transport_chttp2
            libgrpc_transport_chttp2_client_connector
            libgrpc_transport_chttp2_client_insecure
            libgrpc_transport_chttp2_client_secure
            libgrpc_transport_chttp2_server_insecure
            libgrpc_transport_chttp2_server_secure
            libgrpc_transport_inproc
            libgrpc_workaround_cronet_compression_filter
            libm
            libprotobuf-c-nano
            libz
    )
endif ()
#################### libgrpc ####################

#################### libgrpc_message_size_filter ####################
set(FILES
        ${ANDROID_ROOT}/external/grpc-grpc/src/core/ext/filters/message_size/message_size_filter.cc
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
    add_library(libgrpc_message_size_filter SHARED ${SOURCE_FILES})

    target_include_directories(libgrpc_message_size_filter
            PUBLIC
            ${ANDROID_ROOT}/external/grpc-grpc/include
            ${ANDROID_ROOT}/external/grpc-grpc
    )

    # Add c/c++ flags
    target_compile_options(libgrpc_message_size_filter PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(libgrpc_message_size_filter
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libgrpc_base
            libm
            libz
    )
endif ()
#################### libgrpc_message_size_filter ####################
