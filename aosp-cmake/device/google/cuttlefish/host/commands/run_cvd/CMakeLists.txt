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

project(run_cvd)

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

#################### run_cvd ####################
set(FILES
        ${ANDROID_ROOT}/device/google/cuttlefish/host/commands/run_cvd/boot_state_machine.cc
        ${ANDROID_ROOT}/device/google/cuttlefish/host/commands/run_cvd/launch/bluetooth_connector.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/host/commands/run_cvd/launch/uwb_connector.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/host/commands/run_cvd/launch/config_server.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/host/commands/run_cvd/launch/console_forwarder.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/host/commands/run_cvd/launch/gnss_grpc_proxy.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/host/commands/run_cvd/launch/kernel_log_monitor.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/host/commands/run_cvd/launch/logcat_receiver.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/host/commands/run_cvd/launch/log_tee_creator.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/host/commands/run_cvd/launch/grpc_socket_creator.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/host/commands/run_cvd/launch/modem.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/host/commands/run_cvd/launch/metrics.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/host/commands/run_cvd/launch/open_wrt.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/host/commands/run_cvd/launch/openwrt_control_server.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/host/commands/run_cvd/launch/echo_server.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/host/commands/run_cvd/launch/root_canal.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/host/commands/run_cvd/launch/pica.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/host/commands/run_cvd/launch/secure_env.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/host/commands/run_cvd/launch/streamer.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/host/commands/run_cvd/launch/tombstone_receiver.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/host/commands/run_cvd/launch/wmediumd_server.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/host/commands/run_cvd/launch/netsim_server.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/host/commands/run_cvd/main.cc
        ${ANDROID_ROOT}/device/google/cuttlefish/host/commands/run_cvd/reporting.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/host/commands/run_cvd/process_monitor.cc
        ${ANDROID_ROOT}/device/google/cuttlefish/host/commands/run_cvd/server_loop.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/host/commands/run_cvd/validate.cpp
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
    add_executable(run_cvd ${SOURCE_FILES})

    target_include_directories(run_cvd
            PUBLIC
            ${ANDROID_ROOT}/device/google/cuttlefish/host/commands/run_cvd
    )

    # Add c/c++ flags
    target_compile_options(run_cvd PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(run_cvd
            cuttlefish_common_headers
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcuttlefish_fs
            libcuttlefish_host_config
            libcuttlefish_host_config_adb
            libcuttlefish_host_config_fastboot
            libcuttlefish_kernel_log_monitor_utils
            libcuttlefish_metrics
            libcuttlefish_msg_queue
            libcuttlefish_utils
            libcuttlefish_vm_manager
            libdl
            libext2_blkid
            libfruit
            libgflags
            libicuuc
            libjsoncpp
            libm
            libnl
    )
endif ()
#################### run_cvd ####################
