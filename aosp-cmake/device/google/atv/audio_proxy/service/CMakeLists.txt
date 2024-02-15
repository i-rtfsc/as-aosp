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

project(device.google.atv.audio_proxy@7.1-service)

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

#################### audio_proxy_service_util ####################
set(FILES
        ${ANDROID_ROOT}/device/google/atv/audio_proxy/service/RingBufferUtil.cpp
        ${ANDROID_ROOT}/device/google/atv/audio_proxy/service/ServiceConfig.cpp
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
    add_library(audio_proxy_service_util SHARED ${SOURCE_FILES})

    target_include_directories(audio_proxy_service_util
            PUBLIC
            ${ANDROID_ROOT}/device/google/atv/audio_proxy/service
    )

    target_link_libraries(audio_proxy_service_util
            libbase
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### audio_proxy_service_util ####################

#################### device.google.atv.audio_proxy-7.0-service ####################
set(FILES
        ${ANDROID_ROOT}/device/google/atv/audio_proxy/service/AudioProxyImpl.cpp
        ${ANDROID_ROOT}/device/google/atv/audio_proxy/service/AudioUtil.cpp
        ${ANDROID_ROOT}/device/google/atv/audio_proxy/service/BusOutputStream.cpp
        ${ANDROID_ROOT}/device/google/atv/audio_proxy/service/BusStreamProvider.cpp
        ${ANDROID_ROOT}/device/google/atv/audio_proxy/service/DeviceImpl.cpp
        ${ANDROID_ROOT}/device/google/atv/audio_proxy/service/DevicesFactoryImpl.cpp
        ${ANDROID_ROOT}/device/google/atv/audio_proxy/service/DummyBusOutputStream.cpp
        ${ANDROID_ROOT}/device/google/atv/audio_proxy/service/RemoteBusOutputStream.cpp
        ${ANDROID_ROOT}/device/google/atv/audio_proxy/service/StreamOutImpl.cpp
        ${ANDROID_ROOT}/device/google/atv/audio_proxy/service/WriteThread.cpp
        ${ANDROID_ROOT}/device/google/atv/audio_proxy/service/main.cpp
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
    add_executable(device.google.atv.audio_proxy-7.0-service ${SOURCE_FILES})

    target_include_directories(device.google.atv.audio_proxy-7.0-service
            PUBLIC
            ${ANDROID_ROOT}/device/google/atv/audio_proxy/service
    )

    # Add c/c++ flags
    target_compile_options(device.google.atv.audio_proxy-7.0-service PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(device.google.atv.audio_proxy-7.0-service
            android.hardware.audio.common-7.0
            android.hardware.audio-7.0
            audio_proxy_service_util
            device.google.atv.audio_proxy-aidl-V3-ndk
            libaudio_system_headers
            libbase
            libbinder_ndk
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libcutils
            libdl
            libfmq
            libhidlbase
            libm
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### device.google.atv.audio_proxy-7.0-service ####################

#################### device.google.atv.audio_proxy-5.1-service ####################
set(FILES
        ${ANDROID_ROOT}/device/google/atv/audio_proxy/service/AudioProxyImpl.cpp
        ${ANDROID_ROOT}/device/google/atv/audio_proxy/service/AudioUtil.cpp
        ${ANDROID_ROOT}/device/google/atv/audio_proxy/service/BusOutputStream.cpp
        ${ANDROID_ROOT}/device/google/atv/audio_proxy/service/BusStreamProvider.cpp
        ${ANDROID_ROOT}/device/google/atv/audio_proxy/service/DeviceImpl.cpp
        ${ANDROID_ROOT}/device/google/atv/audio_proxy/service/DevicesFactoryImpl.cpp
        ${ANDROID_ROOT}/device/google/atv/audio_proxy/service/DummyBusOutputStream.cpp
        ${ANDROID_ROOT}/device/google/atv/audio_proxy/service/RemoteBusOutputStream.cpp
        ${ANDROID_ROOT}/device/google/atv/audio_proxy/service/StreamOutImpl.cpp
        ${ANDROID_ROOT}/device/google/atv/audio_proxy/service/WriteThread.cpp
        ${ANDROID_ROOT}/device/google/atv/audio_proxy/service/main.cpp
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
    add_executable(device.google.atv.audio_proxy-5.1-service ${SOURCE_FILES})

    target_include_directories(device.google.atv.audio_proxy-5.1-service
            PUBLIC
            ${ANDROID_ROOT}/device/google/atv/audio_proxy/service
    )

    # Add c/c++ flags
    target_compile_options(device.google.atv.audio_proxy-5.1-service PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(device.google.atv.audio_proxy-5.1-service
            android.hardware.audio.common-5.0
            android.hardware.audio-5.0
            audio_proxy_service_util
            device.google.atv.audio_proxy-aidl-V3-ndk
            libaudio_system_headers
            libbase
            libbinder_ndk
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libcutils
            libdl
            libfmq
            libhidlbase
            libm
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### device.google.atv.audio_proxy-5.1-service ####################

#################### device.google.atv.audio_proxy-6.0-service ####################
set(FILES
        ${ANDROID_ROOT}/device/google/atv/audio_proxy/service/AudioProxyImpl.cpp
        ${ANDROID_ROOT}/device/google/atv/audio_proxy/service/AudioUtil.cpp
        ${ANDROID_ROOT}/device/google/atv/audio_proxy/service/BusOutputStream.cpp
        ${ANDROID_ROOT}/device/google/atv/audio_proxy/service/BusStreamProvider.cpp
        ${ANDROID_ROOT}/device/google/atv/audio_proxy/service/DeviceImpl.cpp
        ${ANDROID_ROOT}/device/google/atv/audio_proxy/service/DevicesFactoryImpl.cpp
        ${ANDROID_ROOT}/device/google/atv/audio_proxy/service/DummyBusOutputStream.cpp
        ${ANDROID_ROOT}/device/google/atv/audio_proxy/service/RemoteBusOutputStream.cpp
        ${ANDROID_ROOT}/device/google/atv/audio_proxy/service/StreamOutImpl.cpp
        ${ANDROID_ROOT}/device/google/atv/audio_proxy/service/WriteThread.cpp
        ${ANDROID_ROOT}/device/google/atv/audio_proxy/service/main.cpp
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
    add_executable(device.google.atv.audio_proxy-6.0-service ${SOURCE_FILES})

    target_include_directories(device.google.atv.audio_proxy-6.0-service
            PUBLIC
            ${ANDROID_ROOT}/device/google/atv/audio_proxy/service
    )

    # Add c/c++ flags
    target_compile_options(device.google.atv.audio_proxy-6.0-service PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(device.google.atv.audio_proxy-6.0-service
            android.hardware.audio.common-6.0
            android.hardware.audio-6.0
            audio_proxy_service_util
            device.google.atv.audio_proxy-aidl-V3-ndk
            libaudio_system_headers
            libbase
            libbinder_ndk
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libcutils
            libdl
            libfmq
            libhidlbase
            libm
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### device.google.atv.audio_proxy-6.0-service ####################

#################### device.google.atv.audio_proxy-7.1-service ####################
set(FILES
        ${ANDROID_ROOT}/device/google/atv/audio_proxy/service/AudioProxyImpl.cpp
        ${ANDROID_ROOT}/device/google/atv/audio_proxy/service/AudioUtil.cpp
        ${ANDROID_ROOT}/device/google/atv/audio_proxy/service/BusOutputStream.cpp
        ${ANDROID_ROOT}/device/google/atv/audio_proxy/service/BusStreamProvider.cpp
        ${ANDROID_ROOT}/device/google/atv/audio_proxy/service/DeviceImpl.cpp
        ${ANDROID_ROOT}/device/google/atv/audio_proxy/service/DevicesFactoryImpl.cpp
        ${ANDROID_ROOT}/device/google/atv/audio_proxy/service/DummyBusOutputStream.cpp
        ${ANDROID_ROOT}/device/google/atv/audio_proxy/service/RemoteBusOutputStream.cpp
        ${ANDROID_ROOT}/device/google/atv/audio_proxy/service/StreamOutImpl.cpp
        ${ANDROID_ROOT}/device/google/atv/audio_proxy/service/WriteThread.cpp
        ${ANDROID_ROOT}/device/google/atv/audio_proxy/service/main.cpp
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
    add_executable(device.google.atv.audio_proxy-7.1-service ${SOURCE_FILES})

    target_include_directories(device.google.atv.audio_proxy-7.1-service
            PUBLIC
            ${ANDROID_ROOT}/device/google/atv/audio_proxy/service
    )

    # Add c/c++ flags
    target_compile_options(device.google.atv.audio_proxy-7.1-service PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(device.google.atv.audio_proxy-7.1-service
            android.hardware.audio.common-7.0
            android.hardware.audio-7.0
            android.hardware.audio-7.1
            audio_proxy_service_util
            device.google.atv.audio_proxy-aidl-V3-ndk
            libaudio_system_headers
            libbase
            libbinder_ndk
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libcutils
            libdl
            libfmq
            libhidlbase
            libm
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### device.google.atv.audio_proxy-7.1-service ####################
