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

project(sensors.dynamic_sensor_hal)

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

#################### sensors.dynamic_sensor_hal ####################
set(FILES
        ${ANDROID_ROOT}/hardware/libhardware/modules/sensors/dynamic_sensor/DynamicSensorsSubHal.cpp
        ${ANDROID_ROOT}/hardware/libhardware/modules/sensors/dynamic_sensor/sensors.cpp
        ${ANDROID_ROOT}/hardware/libhardware/modules/sensors/dynamic_sensor/BaseDynamicSensorDaemon.cpp
        ${ANDROID_ROOT}/hardware/libhardware/modules/sensors/dynamic_sensor/BaseSensorObject.cpp
        ${ANDROID_ROOT}/hardware/libhardware/modules/sensors/dynamic_sensor/ConnectionDetector.cpp
        ${ANDROID_ROOT}/hardware/libhardware/modules/sensors/dynamic_sensor/DummyDynamicAccelDaemon.cpp
        ${ANDROID_ROOT}/hardware/libhardware/modules/sensors/dynamic_sensor/DynamicSensorManager.cpp
        ${ANDROID_ROOT}/hardware/libhardware/modules/sensors/dynamic_sensor/HidRawDevice.cpp
        ${ANDROID_ROOT}/hardware/libhardware/modules/sensors/dynamic_sensor/HidRawSensor.cpp
        ${ANDROID_ROOT}/hardware/libhardware/modules/sensors/dynamic_sensor/HidRawSensorDaemon.cpp
        ${ANDROID_ROOT}/hardware/libhardware/modules/sensors/dynamic_sensor/HidRawSensorDevice.cpp
        ${ANDROID_ROOT}/hardware/libhardware/modules/sensors/dynamic_sensor/RingBuffer.cpp
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
    add_library(sensors.dynamic_sensor_hal SHARED ${SOURCE_FILES})

    target_include_directories(sensors.dynamic_sensor_hal
            PUBLIC
            ${ANDROID_ROOT}/hardware/libhardware/modules/sensors/dynamic_sensor
            ${ANDROID_ROOT}/hardware/libhardware/modules/sensors/dynamic_sensor/HidUtils
    )

    # Add c/c++ flags
    target_compile_options(sensors.dynamic_sensor_hal PRIVATE
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(sensors.dynamic_sensor_hal
            android.hardware.sensors-1.0-convert
            android.hardware.sensors-2.0-ScopedWakelock
            android.hardware.sensors-2.0
            android.hardware.sensors-2.1
            android.hardware.sensors-2.X-multihal.header
            android.hardware.sensors-2.X-shared-utils
            libbase
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libhardware_headers
            libhidlbase
            libhidparser
            liblog
            libm
            libstagefright_foundation_headers
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### sensors.dynamic_sensor_hal ####################

#################### libdynamic_sensor_ext ####################
set(FILES
        ${ANDROID_ROOT}/hardware/libhardware/modules/sensors/dynamic_sensor/BaseDynamicSensorDaemon.cpp
        ${ANDROID_ROOT}/hardware/libhardware/modules/sensors/dynamic_sensor/BaseSensorObject.cpp
        ${ANDROID_ROOT}/hardware/libhardware/modules/sensors/dynamic_sensor/ConnectionDetector.cpp
        ${ANDROID_ROOT}/hardware/libhardware/modules/sensors/dynamic_sensor/DummyDynamicAccelDaemon.cpp
        ${ANDROID_ROOT}/hardware/libhardware/modules/sensors/dynamic_sensor/DynamicSensorManager.cpp
        ${ANDROID_ROOT}/hardware/libhardware/modules/sensors/dynamic_sensor/HidRawDevice.cpp
        ${ANDROID_ROOT}/hardware/libhardware/modules/sensors/dynamic_sensor/HidRawSensor.cpp
        ${ANDROID_ROOT}/hardware/libhardware/modules/sensors/dynamic_sensor/HidRawSensorDaemon.cpp
        ${ANDROID_ROOT}/hardware/libhardware/modules/sensors/dynamic_sensor/HidRawSensorDevice.cpp
        ${ANDROID_ROOT}/hardware/libhardware/modules/sensors/dynamic_sensor/RingBuffer.cpp
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
    add_library(libdynamic_sensor_ext SHARED ${SOURCE_FILES})

    target_include_directories(libdynamic_sensor_ext
            PUBLIC
            ${ANDROID_ROOT}/hardware/libhardware/modules/sensors/dynamic_sensor
            ${ANDROID_ROOT}/hardware/libhardware/modules/sensors/dynamic_sensor/HidUtils
    )

    # Add c/c++ flags
    target_compile_options(libdynamic_sensor_ext PRIVATE
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(libdynamic_sensor_ext
            libbase
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libcutils
            libdl
            libhardware_headers
            libhidparser
            liblog
            libm
            libstagefright_foundation_headers
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### libdynamic_sensor_ext ####################
