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

project(android.hardware.automotive.vehicle@default-grpc-libgrpc)

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

#################### android.hardware.automotive.vehicle-default-grpc-libgrpc ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/grpc/VehicleServerProtoStub_cc@default-grpc/gen/VehicleServer.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/grpc/VehicleServerProtoStub_cc@default-grpc/gen/VehicleServer.grpc.pb.cc
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
    add_library(android.hardware.automotive.vehicle-default-grpc-libgrpc SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.automotive.vehicle-default-grpc-libgrpc
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/grpc
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/grpc/VehicleServerProtoStub_cc@default-grpc/gen
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/grpc/VehicleServerProtoStub_h@default-grpc/gen
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.automotive.vehicle-default-grpc-libgrpc PRIVATE
            -Wno-unused-parameter
    )

    target_link_libraries(android.hardware.automotive.vehicle-default-grpc-libgrpc
            VehicleHalProtos
            libc++_static
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libgrpc++
            libgrpc++
            libm
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.automotive.vehicle-default-grpc-libgrpc ####################

#################### android.hardware.automotive.vehicle-default-grpc-hardware-lib ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/grpc/GRPCVehicleHardware.cpp
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
    add_library(android.hardware.automotive.vehicle-default-grpc-hardware-lib SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.automotive.vehicle-default-grpc-hardware-lib
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/grpc
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/grpc/VehicleServerProtoStub_h@default-grpc/gen
            ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/grpc/utils/proto_message_converter/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.automotive.vehicle-default-grpc-hardware-lib PRIVATE
            -Wall
            -Wextra
            -Werror
            -Wall
            -Wextra
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(android.hardware.automotive.vehicle-default-grpc-hardware-lib
            IVehicleHardware
            VehicleHalProtoMessageConverter
            android-automotive-large-parcelable-lib
            android.hardware.automotive.vehicle-V2-ndk
            android.hardware.automotive.vehicle.property-V2-ndk
            android.hardware.automotive.vehicle-default-grpc-libgrpc
            libbase
            libbinder_ndk
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libgrpc++
            liblog
            libm
            libmath
            libprotobuf-cpp-full
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.automotive.vehicle-default-grpc-hardware-lib ####################

#################### android.hardware.automotive.vehicle-default-grpc-server-lib ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/grpc/GRPCVehicleProxyServer.cpp
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
    add_library(android.hardware.automotive.vehicle-default-grpc-server-lib SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.automotive.vehicle-default-grpc-server-lib
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/grpc
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/grpc/VehicleServerProtoStub_h@default-grpc/gen
            ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/grpc/utils/proto_message_converter/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.automotive.vehicle-default-grpc-server-lib PRIVATE
            -Wall
            -Wextra
            -Werror
            -Wall
            -Wextra
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(android.hardware.automotive.vehicle-default-grpc-server-lib
            IVehicleHardware
            VehicleHalProtoMessageConverter
            android-automotive-large-parcelable-lib
            android.hardware.automotive.vehicle-V2-ndk
            android.hardware.automotive.vehicle.property-V2-ndk
            android.hardware.automotive.vehicle-default-grpc-libgrpc
            libbase
            libbinder_ndk
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libgrpc++
            liblog
            libm
            libmath
            libprotobuf-cpp-full
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.automotive.vehicle-default-grpc-server-lib ####################
