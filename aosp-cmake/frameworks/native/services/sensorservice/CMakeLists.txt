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

project(sensorservice)

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

#################### libsensorservice ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/services/sensorservice/AidlSensorHalWrapper.cpp
        ${ANDROID_ROOT}/frameworks/native/services/sensorservice/BatteryService.cpp
        ${ANDROID_ROOT}/frameworks/native/services/sensorservice/CorrectedGyroSensor.cpp
        ${ANDROID_ROOT}/frameworks/native/services/sensorservice/Fusion.cpp
        ${ANDROID_ROOT}/frameworks/native/services/sensorservice/GravitySensor.cpp
        ${ANDROID_ROOT}/frameworks/native/services/sensorservice/HidlSensorHalWrapper.cpp
        ${ANDROID_ROOT}/frameworks/native/services/sensorservice/LimitedAxesImuSensor.cpp
        ${ANDROID_ROOT}/frameworks/native/services/sensorservice/LinearAccelerationSensor.cpp
        ${ANDROID_ROOT}/frameworks/native/services/sensorservice/OrientationSensor.cpp
        ${ANDROID_ROOT}/frameworks/native/services/sensorservice/RecentEventLogger.cpp
        ${ANDROID_ROOT}/frameworks/native/services/sensorservice/RotationVectorSensor.cpp
        ${ANDROID_ROOT}/frameworks/native/services/sensorservice/SensorDevice.cpp
        ${ANDROID_ROOT}/frameworks/native/services/sensorservice/SensorDeviceUtils.cpp
        ${ANDROID_ROOT}/frameworks/native/services/sensorservice/SensorDirectConnection.cpp
        ${ANDROID_ROOT}/frameworks/native/services/sensorservice/SensorEventConnection.cpp
        ${ANDROID_ROOT}/frameworks/native/services/sensorservice/SensorFusion.cpp
        ${ANDROID_ROOT}/frameworks/native/services/sensorservice/SensorInterface.cpp
        ${ANDROID_ROOT}/frameworks/native/services/sensorservice/SensorList.cpp
        ${ANDROID_ROOT}/frameworks/native/services/sensorservice/SensorRecord.cpp
        ${ANDROID_ROOT}/frameworks/native/services/sensorservice/SensorService.cpp
        ${ANDROID_ROOT}/frameworks/native/services/sensorservice/SensorServiceUtils.cpp
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
    add_library(libsensorservice SHARED ${SOURCE_FILES})

    target_include_directories(libsensorservice
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/services/sensorservice
    )

    # Add c/c++ flags
    target_compile_options(libsensorservice PRIVATE
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(libsensorservice
            android.hardware.common-V2-ndk
            android.hardware.common.fmq-V1-ndk
            android.hardware.sensors-V1-convert
            android.hardware.sensors-V2-ndk
            android.hardware.sensors-1.0
            android.hardware.sensors-1.0-convert
            android.hardware.sensors-2.0
            android.hardware.sensors-2.1
            android.hardware.sensors-2.X-shared-utils
            libactivitymanager_aidl
            libaidlcommonsupport
            libbase
            libbatterystats_aidl
            libbinder
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto
            libcutils
            libdl
            libfmq
            libhardware
            libhardware_legacy
            libhidlbase
            liblog
            libm
            libpermission
            libprotoutil
            libsensor
            libsensorprivacy
            libutils
            packagemanager_aidl-cpp
    )
endif ()
#################### libsensorservice ####################

#################### sensorservice ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/services/sensorservice/main_sensorservice.cpp
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
    add_executable(sensorservice ${SOURCE_FILES})

    target_include_directories(sensorservice
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/services/sensorservice
    )

    # Add c/c++ flags
    target_compile_options(sensorservice PRIVATE
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(sensorservice
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libsensorprivacy
            libsensorservice
            libutils
    )
endif ()
#################### sensorservice ####################
