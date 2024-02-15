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

project(libbt_audio_hal_hearing_aid_software_encoding_fuzzer)

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

#################### libbt_audio_hal_a2dp_encoding_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/audio_hal_interface/fuzzer/libbt_audio_hal_a2dp_encoding_fuzzer.cpp
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
    add_executable(libbt_audio_hal_a2dp_encoding_fuzzer ${SOURCE_FILES})

    target_include_directories(libbt_audio_hal_a2dp_encoding_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/audio_hal_interface/fuzzer
    )

    # Add c/c++ flags
    target_compile_options(libbt_audio_hal_a2dp_encoding_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(libbt_audio_hal_a2dp_encoding_fuzzer
            android.hardware.bluetooth.a2dp-1.0
            android.hardware.bluetooth.audio-V3-ndk
            android.hardware.bluetooth.audio-2.0
            android.hardware.bluetooth.audio-2.1
            android.hardware.bluetooth-1.0
            android.hardware.bluetooth-1.1
            android.system.suspend-V1-ndk
            android.system.suspend.control-V1-ndk
            avrcp_headers
            libFraunhoferAAC
            libPlatformProperties
            libaaudio
            libaudio-a2dp-hw-utils
            libbase
            libbinder_ndk
            libbluetooth-dumpsys
            libbluetooth-protos
            libbluetooth-types
            libbluetooth_core_rs
            libbluetooth_gd
            libbluetooth_headers
            libbluetooth_rust_interop
            libbt-audio-hal-interface
            libbt-bta
            libbt-bta-core
            libbt-common
            libbt-hci
            libbt-sbc-decoder
            libbt-sbc-encoder
            libbt-stack
            libbt-stack-core
            libbt_shim_bridge
            libbtcore
            libbtdevice
            libbte
            libbtif
            libbtif-core
            libc
            libc++
            libc++demangle
            libcgrouprc
            libcgrouprc_format
            libchrome
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcrypto
            libcutils
            libdl
            libevent
            libflatbuffers-cpp
            libfmq
            libg722codec
            libhidlbase
            libjsoncpp
            liblc3
            liblog
            libm
            libopus
            libosi
            libprotobuf-cpp-lite
            libstatslog
            libstatslog_bt
            libstatssocket
            libudrv-uipc
            libutils
            libvndksupport
            libz
    )
endif ()
#################### libbt_audio_hal_a2dp_encoding_fuzzer ####################

#################### libbt_audio_hal_le_audio_software_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/audio_hal_interface/fuzzer/libbt_audio_hal_le_audio_software_fuzzer.cpp
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
    add_executable(libbt_audio_hal_le_audio_software_fuzzer ${SOURCE_FILES})

    target_include_directories(libbt_audio_hal_le_audio_software_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/audio_hal_interface/fuzzer
    )

    # Add c/c++ flags
    target_compile_options(libbt_audio_hal_le_audio_software_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(libbt_audio_hal_le_audio_software_fuzzer
            android.hardware.bluetooth.a2dp-1.0
            android.hardware.bluetooth.audio-V3-ndk
            android.hardware.bluetooth.audio-2.0
            android.hardware.bluetooth.audio-2.1
            android.hardware.bluetooth-1.0
            android.hardware.bluetooth-1.1
            android.system.suspend-V1-ndk
            android.system.suspend.control-V1-ndk
            avrcp_headers
            libFraunhoferAAC
            libPlatformProperties
            libaaudio
            libaudio-a2dp-hw-utils
            libbase
            libbinder_ndk
            libbluetooth-dumpsys
            libbluetooth-protos
            libbluetooth-types
            libbluetooth_core_rs
            libbluetooth_gd
            libbluetooth_headers
            libbluetooth_rust_interop
            libbt-audio-hal-interface
            libbt-bta
            libbt-bta-core
            libbt-common
            libbt-hci
            libbt-sbc-decoder
            libbt-sbc-encoder
            libbt-stack
            libbt-stack-core
            libbt_shim_bridge
            libbtcore
            libbtdevice
            libbte
            libbtif
            libbtif-core
            libc
            libc++
            libc++demangle
            libcgrouprc
            libcgrouprc_format
            libchrome
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcrypto
            libcutils
            libdl
            libevent
            libflatbuffers-cpp
            libfmq
            libg722codec
            libhidlbase
            libjsoncpp
            liblc3
            liblog
            libm
            libopus
            libosi
            libprotobuf-cpp-lite
            libstatslog
            libstatslog_bt
            libstatssocket
            libudrv-uipc
            libutils
            libvndksupport
            libz
    )
endif ()
#################### libbt_audio_hal_le_audio_software_fuzzer ####################

#################### libbt_audio_hal_client_interface_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/audio_hal_interface/fuzzer/libbt_audio_hal_client_interface_fuzzer.cpp
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
    add_executable(libbt_audio_hal_client_interface_fuzzer ${SOURCE_FILES})

    target_include_directories(libbt_audio_hal_client_interface_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/audio_hal_interface/fuzzer
    )

    # Add c/c++ flags
    target_compile_options(libbt_audio_hal_client_interface_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(libbt_audio_hal_client_interface_fuzzer
            android.hardware.bluetooth.a2dp-1.0
            android.hardware.bluetooth.audio-V3-ndk
            android.hardware.bluetooth.audio-2.0
            android.hardware.bluetooth.audio-2.1
            android.hardware.bluetooth-1.0
            android.hardware.bluetooth-1.1
            android.system.suspend-V1-ndk
            android.system.suspend.control-V1-ndk
            avrcp_headers
            libFraunhoferAAC
            libPlatformProperties
            libaaudio
            libaudio-a2dp-hw-utils
            libbase
            libbinder_ndk
            libbluetooth-dumpsys
            libbluetooth-protos
            libbluetooth-types
            libbluetooth_core_rs
            libbluetooth_gd
            libbluetooth_headers
            libbluetooth_rust_interop
            libbt-audio-hal-interface
            libbt-bta
            libbt-bta-core
            libbt-common
            libbt-hci
            libbt-sbc-decoder
            libbt-sbc-encoder
            libbt-stack
            libbt-stack-core
            libbt_shim_bridge
            libbtcore
            libbtdevice
            libbte
            libbtif
            libbtif-core
            libc
            libc++
            libc++demangle
            libcgrouprc
            libcgrouprc_format
            libchrome
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcrypto
            libcutils
            libdl
            libevent
            libflatbuffers-cpp
            libfmq
            libg722codec
            libhidlbase
            libjsoncpp
            liblc3
            liblog
            libm
            libopus
            libosi
            libprotobuf-cpp-lite
            libstatslog
            libstatslog_bt
            libstatssocket
            libudrv-uipc
            libutils
            libvndksupport
            libz
    )
endif ()
#################### libbt_audio_hal_client_interface_fuzzer ####################

#################### libbt_audio_hal_hearing_aid_software_encoding_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/audio_hal_interface/fuzzer/libbt_audio_hal_hearing_aid_software_encoding_fuzzer.cpp
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
    add_executable(libbt_audio_hal_hearing_aid_software_encoding_fuzzer ${SOURCE_FILES})

    target_include_directories(libbt_audio_hal_hearing_aid_software_encoding_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/audio_hal_interface/fuzzer
    )

    # Add c/c++ flags
    target_compile_options(libbt_audio_hal_hearing_aid_software_encoding_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(libbt_audio_hal_hearing_aid_software_encoding_fuzzer
            android.hardware.bluetooth.a2dp-1.0
            android.hardware.bluetooth.audio-V3-ndk
            android.hardware.bluetooth.audio-2.0
            android.hardware.bluetooth.audio-2.1
            android.hardware.bluetooth-1.0
            android.hardware.bluetooth-1.1
            android.system.suspend-V1-ndk
            android.system.suspend.control-V1-ndk
            avrcp_headers
            libFraunhoferAAC
            libPlatformProperties
            libaaudio
            libaudio-a2dp-hw-utils
            libbase
            libbinder_ndk
            libbluetooth-dumpsys
            libbluetooth-protos
            libbluetooth-types
            libbluetooth_core_rs
            libbluetooth_gd
            libbluetooth_headers
            libbluetooth_rust_interop
            libbt-audio-hal-interface
            libbt-bta
            libbt-bta-core
            libbt-common
            libbt-hci
            libbt-sbc-decoder
            libbt-sbc-encoder
            libbt-stack
            libbt-stack-core
            libbt_shim_bridge
            libbtcore
            libbtdevice
            libbte
            libbtif
            libbtif-core
            libc
            libc++
            libc++demangle
            libcgrouprc
            libcgrouprc_format
            libchrome
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcrypto
            libcutils
            libdl
            libevent
            libflatbuffers-cpp
            libfmq
            libg722codec
            libhidlbase
            libjsoncpp
            liblc3
            liblog
            libm
            libopus
            libosi
            libprotobuf-cpp-lite
            libstatslog
            libstatslog_bt
            libstatssocket
            libudrv-uipc
            libutils
            libvndksupport
            libz
    )
endif ()
#################### libbt_audio_hal_hearing_aid_software_encoding_fuzzer ####################
