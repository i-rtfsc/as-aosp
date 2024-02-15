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

project(vehicleManager_fuzzer)

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

#################### android.hardware.automotive.vehicle-2.0-default-impl-lib ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default/impl/vhal_v2_0/DefaultVehicleHal.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default/impl/vhal_v2_0/VehicleHalClient.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default/impl/vhal_v2_0/DefaultVehicleHalServer.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default/impl/vhal_v2_0/LinearFakeValueGenerator.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default/impl/vhal_v2_0/JsonFakeValueGenerator.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default/impl/vhal_v2_0/GeneratorHub.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default/impl/vhal_v2_0/FakeObd2Frame.cpp
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
    add_library(android.hardware.automotive.vehicle-2.0-default-impl-lib SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.automotive.vehicle-2.0-default-impl-lib
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default/impl
            ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default/common/include/vhal_v2_0
            ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default
            ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default/impl/vhal_v2_0/userhal/include
            ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default/common/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.automotive.vehicle-2.0-default-impl-lib PRIVATE
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(android.hardware.automotive.vehicle-2.0-default-impl-lib
            android.automotive.watchdog-V2-ndk
            android.hardware.automotive.vehicle-2.0-fake-user-hal-lib
            android.hardware.automotive.vehicle-2.0-libproto-native
            android.hardware.automotive.vehicle-2.0-manager-lib
            android.hardware.automotive.vehicle-2.0
            libbase
            libbinder_ndk
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libhidlbase
            libjsoncpp
            liblog
            libm
            libprotobuf-cpp-lite
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.automotive.vehicle-2.0-default-impl-lib ####################

#################### android.hardware.automotive.vehicle-2.0-manager-lib ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default/common/src/Obd2SensorStore.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default/common/src/ProtoMessageConverter.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default/common/src/SubscriptionManager.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default/common/src/VehicleHalManager.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default/common/src/VehicleObjectPool.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default/common/src/VehiclePropertyStore.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default/common/src/VehicleUtils.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default/common/src/VmsUtils.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default/common/src/WatchdogClient.cpp
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
    add_library(android.hardware.automotive.vehicle-2.0-manager-lib SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.automotive.vehicle-2.0-manager-lib
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default/common/include
            ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default/common/include/vhal_v2_0
            ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.automotive.vehicle-2.0-manager-lib PRIVATE
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(android.hardware.automotive.vehicle-2.0-manager-lib
            android.automotive.watchdog-V2-ndk
            android.hardware.automotive.vehicle-2.0-libproto-native
            android.hardware.automotive.vehicle-2.0
            libbase
            libbinder_ndk
            libc++
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libhidlbase
            liblog
            libm
            libprotobuf-cpp-lite
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.automotive.vehicle-2.0-manager-lib ####################

#################### android.hardware.automotive.vehicle-2.0-default-service ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default/VehicleService.cpp
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
    add_executable(android.hardware.automotive.vehicle-2.0-default-service ${SOURCE_FILES})

    target_include_directories(android.hardware.automotive.vehicle-2.0-default-service
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default
            ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default/common/include
            ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default/impl
            ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default/impl/vhal_v2_0/userhal/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.automotive.vehicle-2.0-default-service PRIVATE
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(android.hardware.automotive.vehicle-2.0-default-service
            android.automotive.watchdog-V2-ndk
            android.hardware.automotive.vehicle-2.0-default-impl-lib
            android.hardware.automotive.vehicle-2.0-libproto-native
            android.hardware.automotive.vehicle-2.0-manager-lib
            android.hardware.automotive.vehicle-2.0
            libbase
            libbinder_ndk
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libhidlbase
            libjsoncpp
            liblog
            libm
            libprotobuf-cpp-lite
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.automotive.vehicle-2.0-default-service ####################

#################### android.hardware.automotive.vehicle-2.0-server-common-lib ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default/common/src/Obd2SensorStore.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default/common/src/ProtoMessageConverter.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default/common/src/VehicleObjectPool.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default/common/src/VehiclePropertyStore.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default/common/src/VehicleUtils.cpp
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
    add_library(android.hardware.automotive.vehicle-2.0-server-common-lib SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.automotive.vehicle-2.0-server-common-lib
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default/common/include
            ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default/common/include/vhal_v2_0
            ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.automotive.vehicle-2.0-server-common-lib PRIVATE
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(android.hardware.automotive.vehicle-2.0-server-common-lib
            android.hardware.automotive.vehicle-2.0
            android.hardware.automotive.vehicle-2.0-libproto-native
            android.hardware.automotive.vehicle-2.0
            libc++_static
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libhidlbase
            libhidlbase
            liblog
            liblog
            libm
            libutils
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.automotive.vehicle-2.0-server-common-lib ####################

#################### vehicleManager_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default/tests/fuzzer/VehicleManager_fuzzer.cpp
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
    add_executable(vehicleManager_fuzzer ${SOURCE_FILES})

    target_include_directories(vehicleManager_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default/common/include
            ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default/tests
            ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default
    )

    # Add c/c++ flags
    target_compile_options(vehicleManager_fuzzer PRIVATE
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(vehicleManager_fuzzer
            android.automotive.watchdog-V2-ndk
            android.hardware.automotive.vehicle-2.0-manager-lib
            android.hardware.automotive.vehicle-2.0
            libbase
            libbase_headers
            libbinder_ndk
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
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
#################### vehicleManager_fuzzer ####################

#################### android.hardware.automotive.vehicle-2.0-server-impl-lib ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default/impl/vhal_v2_0/GeneratorHub.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default/impl/vhal_v2_0/JsonFakeValueGenerator.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default/impl/vhal_v2_0/LinearFakeValueGenerator.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default/impl/vhal_v2_0/DefaultVehicleHalServer.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default/impl/vhal_v2_0/FakeObd2Frame.cpp
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
    add_library(android.hardware.automotive.vehicle-2.0-server-impl-lib SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.automotive.vehicle-2.0-server-impl-lib
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default/impl
            ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default/common/include/vhal_v2_0
            ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default
            ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default/common/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.automotive.vehicle-2.0-server-impl-lib PRIVATE
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(android.hardware.automotive.vehicle-2.0-server-impl-lib
            android.hardware.automotive.vehicle-2.0
            android.hardware.automotive.vehicle-2.0-libproto-native
            android.hardware.automotive.vehicle-2.0-server-common-lib
            android.hardware.automotive.vehicle-2.0
            libbase
            libbase
            libc++_static
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libhidlbase
            libhidlbase
            libjsoncpp
            libjsoncpp
            liblog
            liblog
            libm
            libutils
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.automotive.vehicle-2.0-server-impl-lib ####################
