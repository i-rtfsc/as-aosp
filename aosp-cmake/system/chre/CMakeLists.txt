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

project(chre_aidl_hal_client)

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

#################### android.hardware.contexthub-1.1-common-impl ####################
set(FILES
        ${ANDROID_ROOT}/system/chre/host/hal_generic/V1_1/generic_context_hub_v1_1.cc
        ${ANDROID_ROOT}/system/chre/host/hal_generic/common/context_hub_settings_util.cc
        ${ANDROID_ROOT}/system/chre/host/hal_generic/common/hal_chre_socket_connection.cc
        ${ANDROID_ROOT}/system/chre/host/hal_generic/common/permissions_util.cc
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
    add_library(android.hardware.contexthub-1.1-common-impl SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.contexthub-1.1-common-impl
            PUBLIC
            ${ANDROID_ROOT}/system/chre/host/hal_generic/V1_1
            ${ANDROID_ROOT}/system/chre
            ${ANDROID_ROOT}/system/chre/host/hal_generic/common
            ${ANDROID_ROOT}/system/chre/host/common/include
            ${ANDROID_ROOT}/system/chre/platform/shared/include
            ${ANDROID_ROOT}/system/chre/util/include
            ${ANDROID_ROOT}/system/chre/external/flatbuffers/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.contexthub-1.1-common-impl PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(android.hardware.contexthub-1.1-common-impl
            android.hardware.contexthub-1.0
            android.hardware.contexthub-1.1
            android.hardware.contexthub-1.2
            android.hardware.contexthub-1.X-common-utils
            android.hardware.contexthub-1.X-shared-impl
            chre_client
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libcutils
            libdl
            libhidlbase
            liblog
            libm
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.contexthub-1.1-common-impl ####################

#################### chre_linux ####################
set(FILES
        ${ANDROID_ROOT}/system/chre/core/audio_request_manager.cc
        ${ANDROID_ROOT}/system/chre/core/ble_request_manager.cc
        ${ANDROID_ROOT}/system/chre/core/ble_request_multiplexer.cc
        ${ANDROID_ROOT}/system/chre/core/ble_request.cc
        ${ANDROID_ROOT}/system/chre/core/debug_dump_manager.cc
        ${ANDROID_ROOT}/system/chre/core/event_loop_manager.cc
        ${ANDROID_ROOT}/system/chre/core/event_loop.cc
        ${ANDROID_ROOT}/system/chre/core/event_ref_queue.cc
        ${ANDROID_ROOT}/system/chre/core/event.cc
        ${ANDROID_ROOT}/system/chre/core/gnss_manager.cc
        ${ANDROID_ROOT}/system/chre/core/host_comms_manager.cc
        ${ANDROID_ROOT}/system/chre/core/host_endpoint_manager.cc
        ${ANDROID_ROOT}/system/chre/core/init.cc
        ${ANDROID_ROOT}/system/chre/core/nanoapp.cc
        ${ANDROID_ROOT}/system/chre/core/sensor_request_manager.cc
        ${ANDROID_ROOT}/system/chre/core/sensor_request_multiplexer.cc
        ${ANDROID_ROOT}/system/chre/core/sensor_request.cc
        ${ANDROID_ROOT}/system/chre/core/sensor_type_helpers.cc
        ${ANDROID_ROOT}/system/chre/core/sensor_type.cc
        ${ANDROID_ROOT}/system/chre/core/sensor.cc
        ${ANDROID_ROOT}/system/chre/core/settings.cc
        ${ANDROID_ROOT}/system/chre/core/system_health_monitor.cc
        ${ANDROID_ROOT}/system/chre/core/timer_pool.cc
        ${ANDROID_ROOT}/system/chre/core/wifi_request_manager.cc
        ${ANDROID_ROOT}/system/chre/core/wifi_scan_request.cc
        ${ANDROID_ROOT}/system/chre/platform/linux/assert.cc
        ${ANDROID_ROOT}/system/chre/platform/linux/context.cc
        ${ANDROID_ROOT}/system/chre/platform/linux/fatal_error.cc
        ${ANDROID_ROOT}/system/chre/platform/linux/host_link.cc
        ${ANDROID_ROOT}/system/chre/platform/linux/memory_manager.cc
        ${ANDROID_ROOT}/system/chre/platform/linux/memory.cc
        ${ANDROID_ROOT}/system/chre/platform/linux/pal_audio.cc
        ${ANDROID_ROOT}/system/chre/platform/linux/pal_ble.cc
        ${ANDROID_ROOT}/system/chre/platform/linux/pal_gnss.cc
        ${ANDROID_ROOT}/system/chre/platform/linux/pal_nan.cc
        ${ANDROID_ROOT}/system/chre/platform/linux/pal_sensor.cc
        ${ANDROID_ROOT}/system/chre/platform/linux/pal_wifi.cc
        ${ANDROID_ROOT}/system/chre/platform/linux/platform_debug_dump_manager.cc
        ${ANDROID_ROOT}/system/chre/platform/linux/platform_log.cc
        ${ANDROID_ROOT}/system/chre/platform/linux/platform_nanoapp.cc
        ${ANDROID_ROOT}/system/chre/platform/linux/platform_pal.cc
        ${ANDROID_ROOT}/system/chre/platform/linux/power_control_manager.cc
        ${ANDROID_ROOT}/system/chre/platform/linux/system_time.cc
        ${ANDROID_ROOT}/system/chre/platform/linux/system_timer.cc
        ${ANDROID_ROOT}/system/chre/platform/linux/task_util/task.cc
        ${ANDROID_ROOT}/system/chre/platform/linux/task_util/task_manager.cc
        ${ANDROID_ROOT}/system/chre/platform/shared/audio_pal/platform_audio.cc
        ${ANDROID_ROOT}/system/chre/platform/shared/chre_api_audio.cc
        ${ANDROID_ROOT}/system/chre/platform/shared/chre_api_ble.cc
        ${ANDROID_ROOT}/system/chre/platform/shared/chre_api_core.cc
        ${ANDROID_ROOT}/system/chre/platform/shared/chre_api_gnss.cc
        ${ANDROID_ROOT}/system/chre/platform/shared/chre_api_re.cc
        ${ANDROID_ROOT}/system/chre/platform/shared/chre_api_sensor.cc
        ${ANDROID_ROOT}/system/chre/platform/shared/chre_api_user_settings.cc
        ${ANDROID_ROOT}/system/chre/platform/shared/chre_api_wifi.cc
        ${ANDROID_ROOT}/system/chre/platform/shared/log_buffer.cc
        ${ANDROID_ROOT}/system/chre/platform/shared/memory_manager.cc
        ${ANDROID_ROOT}/system/chre/platform/shared/pal_system_api.cc
        ${ANDROID_ROOT}/system/chre/platform/shared/platform_ble.cc
        ${ANDROID_ROOT}/system/chre/platform/shared/platform_gnss.cc
        ${ANDROID_ROOT}/system/chre/platform/shared/platform_wifi.cc
        ${ANDROID_ROOT}/system/chre/platform/shared/system_time.cc
        ${ANDROID_ROOT}/system/chre/platform/shared/tracing.cc
        ${ANDROID_ROOT}/system/chre/platform/shared/version.cc
        ${ANDROID_ROOT}/system/chre/platform/shared/sensor_pal/platform_sensor.cc
        ${ANDROID_ROOT}/system/chre/platform/shared/sensor_pal/platform_sensor_type_helpers.cc
        ${ANDROID_ROOT}/system/chre/platform/shared/sensor_pal/platform_sensor_manager.cc
        ${ANDROID_ROOT}/system/chre/util/buffer_base.cc
        ${ANDROID_ROOT}/system/chre/util/container_support.cc
        ${ANDROID_ROOT}/system/chre/util/dynamic_vector_base.cc
        ${ANDROID_ROOT}/system/chre/util/intrusive_list_base.cc
        ${ANDROID_ROOT}/system/chre/util/nanoapp/audio.cc
        ${ANDROID_ROOT}/system/chre/util/nanoapp/ble.cc
        ${ANDROID_ROOT}/system/chre/util/nanoapp/callbacks.cc
        ${ANDROID_ROOT}/system/chre/util/nanoapp/debug.cc
        ${ANDROID_ROOT}/system/chre/util/nanoapp/wifi.cc
        ${ANDROID_ROOT}/system/chre/util/pigweed/chre_channel_output.cc
        ${ANDROID_ROOT}/system/chre/util/pigweed/rpc_client.cc
        ${ANDROID_ROOT}/system/chre/util/pigweed/rpc_helper.cc
        ${ANDROID_ROOT}/system/chre/util/pigweed/rpc_server.cc
        ${ANDROID_ROOT}/system/chre/util/system/debug_dump.cc
        ${ANDROID_ROOT}/system/chre/util/system/event_callbacks.cc
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
    add_library(chre_linux SHARED ${SOURCE_FILES})

    target_include_directories(chre_linux
            PUBLIC
            ${ANDROID_ROOT}/system/chre/chre_api/include
            ${ANDROID_ROOT}/system/chre/chre_api/include/chre_api
            ${ANDROID_ROOT}/system/chre/core/include
            ${ANDROID_ROOT}/system/chre/pal/include
            ${ANDROID_ROOT}/system/chre/pal/util/include
            ${ANDROID_ROOT}/system/chre/platform/shared/audio_pal/include
            ${ANDROID_ROOT}/system/chre/platform/linux/include
            ${ANDROID_ROOT}/system/chre/platform/shared/include
            ${ANDROID_ROOT}/system/chre/platform/shared/sensor_pal/include
            ${ANDROID_ROOT}/system/chre/platform/include
            ${ANDROID_ROOT}/system/chre/util/include
            ${ANDROID_ROOT}/system/chre
    )

    # Add c/c++ flags
    target_compile_options(chre_linux PRIVATE
            -Werror
            -DCHRE_MINIMUM_LOG_LEVEL=CHRE_LOG_LEVEL_DEBUG
    )

    target_link_libraries(chre_linux
            chre_api
            libc++_static
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.builtins
            libdl
            libgmock
            libgmock
            libgtest
            libgtest
            libm
            pw_rpc_chre
            pw_rpc_chre
            libc
            libdl
            libm
    )
endif ()
#################### chre_linux ####################

#################### chre_config_util ####################
set(FILES
        ${ANDROID_ROOT}/system/chre/host/common/config_util.cc
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
    add_library(chre_config_util SHARED ${SOURCE_FILES})

    target_include_directories(chre_config_util
            PUBLIC
            ${ANDROID_ROOT}/system/chre/host/common/include
            ${ANDROID_ROOT}/system/chre
    )

    # Add c/c++ flags
    target_compile_options(chre_config_util PRIVATE
            -Werror
    )

    target_link_libraries(chre_config_util
            libc++_static
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.builtins
            libdl
            libjsoncpp
            libjsoncpp
            liblog
            liblog
            libm
            libc
            libdl
            libm
    )
endif ()
#################### chre_config_util ####################

#################### pw_rpc_nanopb_chre ####################
set(FILES
        ${ANDROID_ROOT}/external/pigweed/pw_rpc/nanopb/common.cc
        ${ANDROID_ROOT}/external/pigweed/pw_rpc/nanopb/method.cc
        ${ANDROID_ROOT}/external/pigweed/pw_rpc/nanopb/server_reader_writer.cc
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
    add_library(pw_rpc_nanopb_chre SHARED ${SOURCE_FILES})

    target_include_directories(pw_rpc_nanopb_chre
            PUBLIC
            ${ANDROID_ROOT}/system/chre
    )

    # Add c/c++ flags
    target_compile_options(pw_rpc_nanopb_chre PRIVATE
            -Werror
    )

    target_link_libraries(pw_rpc_nanopb_chre
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libprotobuf-c-nano
            pw_rpc_chre
            pw_rpc_nanopb_include_dirs
            pw_rpc_raw_chre
    )
endif ()
#################### pw_rpc_nanopb_chre ####################

#################### android.hardware.contexthub-1.2-service.generic ####################
set(FILES
        ${ANDROID_ROOT}/system/chre/host/hal_generic/V1_2/generic_context_hub_v1_2.cc
        ${ANDROID_ROOT}/system/chre/host/hal_generic/V1_2/service.cc
        ${ANDROID_ROOT}/system/chre/host/hal_generic/common/context_hub_settings_util.cc
        ${ANDROID_ROOT}/system/chre/host/hal_generic/common/hal_chre_socket_connection.cc
        ${ANDROID_ROOT}/system/chre/host/hal_generic/common/permissions_util.cc
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
    add_executable(android.hardware.contexthub-1.2-service.generic ${SOURCE_FILES})

    target_include_directories(android.hardware.contexthub-1.2-service.generic
            PUBLIC
            ${ANDROID_ROOT}/system/chre
            ${ANDROID_ROOT}/system/chre/host/hal_generic/common
            ${ANDROID_ROOT}/system/chre/host/common/include
            ${ANDROID_ROOT}/system/chre/platform/shared/include
            ${ANDROID_ROOT}/system/chre/util/include
            ${ANDROID_ROOT}/system/chre/external/flatbuffers/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.contexthub-1.2-service.generic PRIVATE
            -Wall
            -Werror
            -Wall
            -Werror
    )

    target_link_libraries(android.hardware.contexthub-1.2-service.generic
            android.hardware.contexthub-1.0
            android.hardware.contexthub-1.1
            android.hardware.contexthub-1.2
            android.hardware.contexthub-1.X-common-utils
            android.hardware.contexthub-1.X-shared-impl
            chre_client
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libcutils
            libdl
            libhidlbase
            liblog
            libm
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.contexthub-1.2-service.generic ####################

#################### chre_pal_linux ####################
set(FILES
        ${ANDROID_ROOT}/system/chre/platform/shared/pal_system_api.cc
        ${ANDROID_ROOT}/system/chre/platform/linux/assert.cc
        ${ANDROID_ROOT}/system/chre/platform/linux/fatal_error.cc
        ${ANDROID_ROOT}/system/chre/platform/linux/memory.cc
        ${ANDROID_ROOT}/system/chre/platform/linux/pal_gnss.cc
        ${ANDROID_ROOT}/system/chre/platform/linux/pal_nan.cc
        ${ANDROID_ROOT}/system/chre/platform/linux/pal_wifi.cc
        ${ANDROID_ROOT}/system/chre/platform/linux/pal_wwan.cc
        ${ANDROID_ROOT}/system/chre/platform/linux/platform_log.cc
        ${ANDROID_ROOT}/system/chre/platform/linux/system_time.cc
        ${ANDROID_ROOT}/system/chre/platform/linux/task_util/task.cc
        ${ANDROID_ROOT}/system/chre/platform/linux/task_util/task_manager.cc
        ${ANDROID_ROOT}/system/chre/util/dynamic_vector_base.cc
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
    add_library(chre_pal_linux SHARED ${SOURCE_FILES})

    target_include_directories(chre_pal_linux
            PUBLIC
            ${ANDROID_ROOT}/system/chre/platform/shared/include
            ${ANDROID_ROOT}/system/chre/platform/include
            ${ANDROID_ROOT}/system/chre/platform/linux/include
            ${ANDROID_ROOT}/system/chre/util/include
            ${ANDROID_ROOT}/system/chre
            ${ANDROID_ROOT}/system/chre/pal/include
            ${ANDROID_ROOT}/system/chre/chre_api/include/chre_api
            ${ANDROID_ROOT}/system/chre/chre_api/include
    )

    # Add c/c++ flags
    target_compile_options(chre_pal_linux PRIVATE
            -Werror
            -DCHRE_MINIMUM_LOG_LEVEL=CHRE_LOG_LEVEL_DEBUG
    )

    target_link_libraries(chre_pal_linux
            chre_pal
            libc++_static
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.builtins
            libdl
            libgmock
            libgmock
            libgtest
            libgtest
            libm
            libc
            libdl
            libm
    )
endif ()
#################### chre_pal_linux ####################

#################### chre_daemon_exynos ####################
set(FILES
        ${ANDROID_ROOT}/system/chre/host/common/daemon_base.cc
        ${ANDROID_ROOT}/system/chre/host/common/fbs_daemon_base.cc
        ${ANDROID_ROOT}/system/chre/host/common/file_stream.cc
        ${ANDROID_ROOT}/system/chre/host/common/fragmented_load_transaction.cc
        ${ANDROID_ROOT}/system/chre/host/common/host_protocol_host.cc
        ${ANDROID_ROOT}/system/chre/host/common/log_message_parser.cc
        ${ANDROID_ROOT}/system/chre/host/common/socket_server.cc
        ${ANDROID_ROOT}/system/chre/host/common/st_hal_lpma_handler.cc
        ${ANDROID_ROOT}/system/chre/platform/shared/host_protocol_common.cc
        ${ANDROID_ROOT}/system/chre/host/exynos/exynos_daemon.cc
        ${ANDROID_ROOT}/system/chre/host/exynos/main.cc
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
    add_executable(chre_daemon_exynos ${SOURCE_FILES})

    target_include_directories(chre_daemon_exynos
            PUBLIC
            ${ANDROID_ROOT}/system/chre/external/flatbuffers/include
            ${ANDROID_ROOT}/system/chre/host/common/include
            ${ANDROID_ROOT}/system/chre/platform/shared/include
            ${ANDROID_ROOT}/system/chre/util/include
            ${ANDROID_ROOT}/system/chre/host/exynos
            ${ANDROID_ROOT}/system/chre
    )

    # Add c/c++ flags
    target_compile_options(chre_daemon_exynos PRIVATE
            -Werror
    )

    target_link_libraries(chre_daemon_exynos
            android.hardware.soundtrigger-2.0
            chre_config_util
            libbase
            libbinder_ndk
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libcutils
            libdl
            libhidlbase
            libjsoncpp
            liblog
            libm
            libpower
            libutils
            pw_detokenizer
            pw_polyfill_headers
            pw_span_headers
            pw_varint
            libc
            libdl
            libm
    )
endif ()
#################### chre_daemon_exynos ####################

#################### chre_client ####################
set(FILES
        ${ANDROID_ROOT}/system/chre/host/common/file_stream.cc
        ${ANDROID_ROOT}/system/chre/host/common/fragmented_load_transaction.cc
        ${ANDROID_ROOT}/system/chre/host/common/host_protocol_host.cc
        ${ANDROID_ROOT}/system/chre/host/common/socket_client.cc
        ${ANDROID_ROOT}/system/chre/platform/shared/host_protocol_common.cc
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
    add_library(chre_client SHARED ${SOURCE_FILES})

    target_include_directories(chre_client
            PUBLIC
            ${ANDROID_ROOT}/system/chre/host/common/include
            ${ANDROID_ROOT}/system/chre/platform/shared/include
            ${ANDROID_ROOT}/system/chre/util/include
            ${ANDROID_ROOT}/system/chre
            ${ANDROID_ROOT}/system/chre/external/flatbuffers/include
    )

    # Add c/c++ flags
    target_compile_options(chre_client PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(chre_client
            chre_flatbuffers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            liblog
            libm
            libutils
    )
endif ()
#################### chre_client ####################

#################### android.hardware.contexthub-1.0-impl.generic ####################
set(FILES
        ${ANDROID_ROOT}/system/chre/host/hal_generic/V1_0/generic_context_hub_v1_0.cc
        ${ANDROID_ROOT}/system/chre/host/hal_generic/common/hal_chre_socket_connection.cc
        ${ANDROID_ROOT}/system/chre/host/hal_generic/common/permissions_util.cc
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
    add_library(android.hardware.contexthub-1.0-impl.generic SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.contexthub-1.0-impl.generic
            PUBLIC
            ${ANDROID_ROOT}/system/chre
            ${ANDROID_ROOT}/system/chre/host/hal_generic/common
            ${ANDROID_ROOT}/system/chre/host/common/include
            ${ANDROID_ROOT}/system/chre/platform/shared/include
            ${ANDROID_ROOT}/system/chre/util/include
            ${ANDROID_ROOT}/system/chre/external/flatbuffers/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.contexthub-1.0-impl.generic PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(android.hardware.contexthub-1.0-impl.generic
            android.hardware.contexthub-1.0
            android.hardware.contexthub-1.1
            android.hardware.contexthub-1.2
            android.hardware.contexthub-1.X-common-utils
            android.hardware.contexthub-1.X-shared-impl
            chre_client
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libcutils
            libdl
            libhidlbase
            liblog
            libm
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.contexthub-1.0-impl.generic ####################

#################### event_logger ####################
set(FILES
        ${ANDROID_ROOT}/system/chre/host/hal_generic/aidl/event_logger.cc
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
    add_library(event_logger SHARED ${SOURCE_FILES})

    target_include_directories(event_logger
            PUBLIC
            ${ANDROID_ROOT}/system/chre/util/include
            ${ANDROID_ROOT}/system/chre/host/common/include
            ${ANDROID_ROOT}/system/chre
            ${ANDROID_ROOT}/system/chre/external/flatbuffers/include
            ${ANDROID_ROOT}/system/chre/platform/shared/include
    )

    # Add c/c++ flags
    target_compile_options(event_logger PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(event_logger
            android.hardware.contexthub-V2-ndk
            android.hardware.contexthub-V2-ndk
            chre_flatbuffers
            chre_flatbuffers
            libc++_static
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.builtins
            libcutils
            libcutils
            libdl
            libm
            libutils
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### event_logger ####################

#################### pw_rpc_raw_chre ####################
set(FILES
        ${ANDROID_ROOT}/external/pigweed/pw_rpc/raw/method.cc
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
    add_library(pw_rpc_raw_chre SHARED ${SOURCE_FILES})

    target_include_directories(pw_rpc_raw_chre
            PUBLIC
            ${ANDROID_ROOT}/system/chre
    )

    # Add c/c++ flags
    target_compile_options(pw_rpc_raw_chre PRIVATE
            -Werror
    )

    target_link_libraries(pw_rpc_raw_chre
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            pw_rpc_chre
            pw_rpc_raw_include_dirs
    )
endif ()
#################### pw_rpc_raw_chre ####################

#################### android.hardware.contexthub-1.1-service.generic ####################
set(FILES
        ${ANDROID_ROOT}/system/chre/host/hal_generic/V1_1/service.cc
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
    add_executable(android.hardware.contexthub-1.1-service.generic ${SOURCE_FILES})

    target_include_directories(android.hardware.contexthub-1.1-service.generic
            PUBLIC
            ${ANDROID_ROOT}/system/chre
            ${ANDROID_ROOT}/system/chre/host/hal_generic/V1_1
            ${ANDROID_ROOT}/system/chre/host/hal_generic/common
            ${ANDROID_ROOT}/system/chre/host/common/include
            ${ANDROID_ROOT}/system/chre/platform/shared/include
            ${ANDROID_ROOT}/system/chre/util/include
            ${ANDROID_ROOT}/system/chre/external/flatbuffers/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.contexthub-1.1-service.generic PRIVATE
            -Wall
            -Werror
            -Wall
            -Werror
    )

    target_link_libraries(android.hardware.contexthub-1.1-service.generic
            android.hardware.contexthub-1.0
            android.hardware.contexthub-1.1-common-impl
            android.hardware.contexthub-1.1
            android.hardware.contexthub-1.2
            chre_client
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libcutils
            libdl
            libhidlbase
            liblog
            libm
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.contexthub-1.1-service.generic ####################

#################### pw_rpc_chre ####################
set(FILES
        ${ANDROID_ROOT}/external/pigweed/pw_rpc/call.cc
        ${ANDROID_ROOT}/external/pigweed/pw_rpc/channel.cc
        ${ANDROID_ROOT}/external/pigweed/pw_rpc/channel_list.cc
        ${ANDROID_ROOT}/external/pigweed/pw_rpc/client.cc
        ${ANDROID_ROOT}/external/pigweed/pw_rpc/client_call.cc
        ${ANDROID_ROOT}/external/pigweed/pw_rpc/client_server.cc
        ${ANDROID_ROOT}/external/pigweed/pw_rpc/endpoint.cc
        ${ANDROID_ROOT}/external/pigweed/pw_rpc/packet.cc
        ${ANDROID_ROOT}/external/pigweed/pw_rpc/packet_meta.cc
        ${ANDROID_ROOT}/external/pigweed/pw_rpc/server.cc
        ${ANDROID_ROOT}/external/pigweed/pw_rpc/server_call.cc
        ${ANDROID_ROOT}/external/pigweed/pw_rpc/service.cc
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
    add_library(pw_rpc_chre SHARED ${SOURCE_FILES})

    target_include_directories(pw_rpc_chre
            PUBLIC
            ${ANDROID_ROOT}/system/chre
    )

    # Add c/c++ flags
    target_compile_options(pw_rpc_chre PRIVATE
            -Werror
    )

    target_link_libraries(pw_rpc_chre
            fuschia_sdk_lib_fit
            fuschia_sdk_lib_stdcompat
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            pw_assert_headers
            pw_assert_log_headers
            pw_bytes
            pw_containers
            pw_function_headers
            pw_log_headers
            pw_log_null_headers
            pw_polyfill_headers
            pw_preprocessor_headers
            pw_protobuf
            pw_result_headers
            pw_rpc_include_dirs
            pw_span_headers
            pw_status
            pw_stream
            pw_string
            pw_sync_baremetal_headers
            pw_sync_headers
            pw_toolchain
            pw_varint
    )
endif ()
#################### pw_rpc_chre ####################

#################### chre_aidl_hal_client ####################
set(FILES
        ${ANDROID_ROOT}/system/chre/host/common/file_stream.cc
        ${ANDROID_ROOT}/system/chre/host/common/chre_aidl_hal_client.cc
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
    add_executable(chre_aidl_hal_client ${SOURCE_FILES})

    target_include_directories(chre_aidl_hal_client
            PUBLIC
            ${ANDROID_ROOT}/system/chre/host/common/include
            ${ANDROID_ROOT}/system/chre/chre_api/include
            ${ANDROID_ROOT}/system/chre
    )

    # Add c/c++ flags
    target_compile_options(chre_aidl_hal_client PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(chre_aidl_hal_client
            android.hardware.contexthub-V2-ndk
            libbase
            libbinder_ndk
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libdl
            libjsoncpp
            liblog
            libm
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### chre_aidl_hal_client ####################

#################### chre_host_util ####################
set(FILES
        ${ANDROID_ROOT}/system/chre/host/common/file_stream.cc
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
    add_library(chre_host_util SHARED ${SOURCE_FILES})

    target_include_directories(chre_host_util
            PUBLIC
            ${ANDROID_ROOT}/system/chre/host/common/include
            ${ANDROID_ROOT}/system/chre
    )

    # Add c/c++ flags
    target_compile_options(chre_host_util PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(chre_host_util
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            liblog
            libm
    )
endif ()
#################### chre_host_util ####################
