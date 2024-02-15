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

project(android.hardware.automotive.vehicle@V1-default-service)

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

#################### android.hardware.automotive.vehicle-V1-default-service_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/vhal/src/fuzzer.cpp
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
    add_executable(android.hardware.automotive.vehicle-V1-default-service_fuzzer ${SOURCE_FILES})

    target_include_directories(android.hardware.automotive.vehicle-V1-default-service_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/vhal
            ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/vhal/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.automotive.vehicle-V1-default-service_fuzzer PRIVATE
            -Wall
            -Wextra
            -Werror
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(android.hardware.automotive.vehicle-V1-default-service_fuzzer
            DefaultVehicleHal
            FakeObd2Frame
            FakeUserHal
            FakeVehicleHalValueGenerators
            FakeVehicleHardware
            IVehicleHardware
            VehicleHalJsonConfigLoaderEnableTestProperties
            VehicleHalTestUtilHeaders
            VehicleHalUtils
            android-automotive-large-parcelable-lib
            android.hardware.automotive.vehicle-V2-ndk
            android.hardware.automotive.vehicle.property-V2-ndk
            libbase
            libbinder
            libbinder_ndk
            libbinder_random_parcel
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcutils
            libdl
            libjsoncpp
            liblog
            libm
            libmath
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.automotive.vehicle-V1-default-service_fuzzer ####################

#################### DefaultVehicleHal ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/vhal/src/ConnectedClient.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/vhal/src/DefaultVehicleHal.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/vhal/src/SubscriptionManager.cpp
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
    add_library(DefaultVehicleHal SHARED ${SOURCE_FILES})

    target_include_directories(DefaultVehicleHal
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/vhal/include
            ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/vhal
    )

    # Add c/c++ flags
    target_compile_options(DefaultVehicleHal PRIVATE
            -Wall
            -Wextra
            -Werror
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(DefaultVehicleHal
            IVehicleHardware
            VehicleHalUtils
            android-automotive-large-parcelable-lib
            android.hardware.automotive.vehicle-V2-ndk
            android.hardware.automotive.vehicle.property-V2-ndk
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
            liblog
            libm
            libmath
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### DefaultVehicleHal ####################

#################### android.hardware.automotive.vehicle-V1-default-service ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/vhal/src/VehicleService.cpp
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
    add_executable(android.hardware.automotive.vehicle-V1-default-service ${SOURCE_FILES})

    target_include_directories(android.hardware.automotive.vehicle-V1-default-service
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/vhal
            ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/vhal/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.automotive.vehicle-V1-default-service PRIVATE
            -Wall
            -Wextra
            -Werror
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(android.hardware.automotive.vehicle-V1-default-service
            DefaultVehicleHal
            FakeObd2Frame
            FakeUserHal
            FakeVehicleHalValueGenerators
            FakeVehicleHardware
            IVehicleHardware
            VehicleHalJsonConfigLoaderEnableTestProperties
            VehicleHalTestUtilHeaders
            VehicleHalUtils
            android-automotive-large-parcelable-lib
            android.hardware.automotive.vehicle-V2-ndk
            android.hardware.automotive.vehicle.property-V2-ndk
            libbase
            libbinder_ndk
            libc++
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
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.automotive.vehicle-V1-default-service ####################
