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

project(libservices.core-gnss)

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

#################### libservices.core-gnss ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/gnss/AGnss.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/gnss/AGnssCallback.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/gnss/AGnssRil.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/gnss/AGnssRilCallback.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/gnss/GnssAntennaInfo.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/gnss/Gnss.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/gnss/GnssCallback.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/gnss/GnssAntennaInfoCallback.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/gnss/GnssBatching.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/gnss/GnssBatchingCallback.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/gnss/GnssConfiguration.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/gnss/GnssDebug.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/gnss/GnssGeofence.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/gnss/GnssGeofenceCallback.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/gnss/GnssMeasurement.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/gnss/GnssMeasurementCallback.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/gnss/GnssNavigationMessage.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/gnss/GnssNavigationMessageCallback.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/gnss/GnssPsds.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/gnss/GnssPsdsCallback.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/gnss/GnssVisibilityControl.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/gnss/GnssVisibilityControlCallback.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/gnss/MeasurementCorrections.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/gnss/MeasurementCorrectionsCallback.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/gnss/Utils.cpp
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
    add_library(libservices.core-gnss SHARED ${SOURCE_FILES})

    target_include_directories(libservices.core-gnss
            PUBLIC
            ${ANDROID_ROOT}/frameworks/base/services/core/jni/gnss
    )

    # Add c/c++ flags
    target_compile_options(libservices.core-gnss PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libservices.core-gnss
            android.hardware.gnss-V3-cpp
            android.hardware.gnss.measurement_corrections-1.0
            android.hardware.gnss.visibility_control-1.0
            android.hardware.gnss-1.0
            android.hardware.gnss-1.1
            android.hardware.gnss-2.0
            android.hardware.gnss-2.1
            libandroid_runtime
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libhardware_legacy
            libhidlbase
            liblog
            libm
            libnativehelper
            libutils
    )
endif ()
#################### libservices.core-gnss ####################
