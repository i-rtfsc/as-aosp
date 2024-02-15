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

project(VehicleEmulator)

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

#################### VehicleEmulator ####################
set(FILES
        ${ANDROID_ROOT}/device/generic/car/emulator/vhal_aidl/VehicleEmulator/EmulatedVehicleHardware.cpp
        ${ANDROID_ROOT}/device/generic/car/emulator/vhal_aidl/VehicleEmulator/VehicleEmulator.cpp
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
    add_library(VehicleEmulator SHARED ${SOURCE_FILES})

    target_include_directories(VehicleEmulator
            PUBLIC
            ${ANDROID_ROOT}/device/generic/car/emulator/vhal_aidl/VehicleEmulator/include
            ${ANDROID_ROOT}/device/generic/car/emulator/vhal_aidl/VehicleEmulator
    )

    # Add c/c++ flags
    target_compile_options(VehicleEmulator PRIVATE
            -Wall
            -Wextra
            -Werror
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(VehicleEmulator
            EmulatorCommConn
            EmulatorPipeComm
            EmulatorSocketComm
            FakeObd2Frame
            FakeUserHal
            FakeVehicleHalValueGenerators
            FakeVehicleHardware
            IVehicleHardware
            VehicleHalJsonConfigLoaderEnableTestProperties
            VehicleHalProtos
            VehicleHalTestUtilHeaders
            VehicleHalUtils
            android-automotive-large-parcelable-lib
            android.hardware.automotive.vehicle-V2-ndk
            android.hardware.automotive.vehicle.property-V2-ndk
            android.hardware.automotive.vehicle-2.0-libproto-native
            device.generic.car.emulator-aidl-V1-ndk
            libbase
            libbinder_ndk
            libc++
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libjsoncpp
            liblog
            libm
            libmath
            libprotobuf-cpp-lite
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### VehicleEmulator ####################
