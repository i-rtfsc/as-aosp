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

project(android.hardware.automotive.vehicle@2.0-virtualization-grpc-server)

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

#################### android.hardware.automotive.vehicle-2.0-virtualization-utils ####################
set(FILES
        ${ANDROID_ROOT}/device/google/trout/hal/vehicle/2.0/Utils.cpp
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
    add_library(android.hardware.automotive.vehicle-2.0-virtualization-utils SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.automotive.vehicle-2.0-virtualization-utils
            PUBLIC
            ${ANDROID_ROOT}/device/google/trout/hal/vehicle/2.0
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.automotive.vehicle-2.0-virtualization-utils PRIVATE
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(android.hardware.automotive.vehicle-2.0-virtualization-utils
            android.hardware.automotive.utils.vsockinfo
            android.hardware.automotive.vehicle-2.0
            android.hardware.automotive.vehicle-2.0
            libbase
            libbase
            libc++_static
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.builtins
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
#################### android.hardware.automotive.vehicle-2.0-virtualization-utils ####################

#################### android.hardware.automotive.vehicle-2.0-virtualization-service ####################
set(FILES
        ${ANDROID_ROOT}/device/google/trout/hal/vehicle/2.0/VirtualizedVehicleService.cpp
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
    add_executable(android.hardware.automotive.vehicle-2.0-virtualization-service ${SOURCE_FILES})

    target_include_directories(android.hardware.automotive.vehicle-2.0-virtualization-service
            PUBLIC
            ${ANDROID_ROOT}/device/google/trout/hal/vehicle/2.0
            ${OUT_INTERMEDIATES_ROOT}/device/google/trout/hal/vehicle/2.0/VehicleServerProtoStub_h@2.0-grpc-trout/gen
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.automotive.vehicle-2.0-virtualization-service PRIVATE
            -Wall
            -Wextra
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(android.hardware.automotive.vehicle-2.0-virtualization-service
            android.automotive.watchdog-V2-ndk
            android.hardware.automotive.vehicle-2.0-virtualization-client-lib
            android.hardware.automotive.vehicle-2.0
            libbase
            libbinder_ndk
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libgrpc++
            libhidlbase
            libjsoncpp
            liblog
            libm
            libprotobuf-cpp-full
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.automotive.vehicle-2.0-virtualization-service ####################

#################### android.hardware.automotive.vehicle-2.0-virtualization-grpc-server-lib ####################
set(FILES
        ${ANDROID_ROOT}/device/google/trout/hal/vehicle/2.0/GarageModeServerSideHandler.cpp
        ${ANDROID_ROOT}/device/google/trout/hal/vehicle/2.0/GrpcVehicleServer.cpp
        ${ANDROID_ROOT}/device/google/trout/hal/vehicle/2.0/PowerStateListener.cpp
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
    add_library(android.hardware.automotive.vehicle-2.0-virtualization-grpc-server-lib SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.automotive.vehicle-2.0-virtualization-grpc-server-lib
            PUBLIC
            ${ANDROID_ROOT}/device/google/trout/hal/vehicle/2.0
            ${OUT_INTERMEDIATES_ROOT}/device/google/trout/hal/vehicle/2.0/VehicleServerProtoStub_h@2.0-grpc-trout/gen
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.automotive.vehicle-2.0-virtualization-grpc-server-lib PRIVATE
            -Wall
            -Wextra
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(android.hardware.automotive.vehicle-2.0-virtualization-grpc-server-lib
            android.hardware.automotive.utils.vsockinfo
            android.hardware.automotive.vehicle-2.0
            android.hardware.automotive.vehicle-2.0-grpc-trout
            android.hardware.automotive.vehicle-2.0-server-impl-lib
            android.hardware.automotive.vehicle-2.0-virtualization-utils
            android.hardware.automotive.vehicle-2.0
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
            libhidlbase
            libhidlbase
            liblog
            liblog
            libm
            libprotobuf-cpp-full
            libprotobuf-cpp-full
            libutils
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.automotive.vehicle-2.0-virtualization-grpc-server-lib ####################

#################### android.hardware.automotive.vehicle-2.0-virtualization-client-lib ####################
set(FILES
        ${ANDROID_ROOT}/device/google/trout/hal/vehicle/2.0/GrpcVehicleClient.cpp
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
    add_library(android.hardware.automotive.vehicle-2.0-virtualization-client-lib SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.automotive.vehicle-2.0-virtualization-client-lib
            PUBLIC
            ${ANDROID_ROOT}/device/google/trout/hal/vehicle/2.0
            ${OUT_INTERMEDIATES_ROOT}/device/google/trout/hal/vehicle/2.0/VehicleServerProtoStub_h@2.0-grpc-trout/gen
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.automotive.vehicle-2.0-virtualization-client-lib PRIVATE
            -Wall
            -Wextra
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(android.hardware.automotive.vehicle-2.0-virtualization-client-lib
            android.automotive.watchdog-V2-ndk
            android.hardware.automotive.utils.vsockinfo
            android.hardware.automotive.vehicle-2.0-default-impl-lib
            android.hardware.automotive.vehicle-2.0-grpc-trout
            android.hardware.automotive.vehicle-2.0-manager-lib
            android.hardware.automotive.vehicle-2.0-virtualization-utils
            android.hardware.automotive.vehicle-2.0
            libbinder_ndk
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libgrpc++
            libhidlbase
            liblog
            libm
            libprotobuf-cpp-full
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.automotive.vehicle-2.0-virtualization-client-lib ####################

#################### android.hardware.automotive.vehicle-2.0-grpc-trout ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/proto/DefaultVehicleHalProtoStub_cc/gen/VehicleHalProto.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/proto/DefaultVehicleHalProtoStub_cc/gen/VehicleHalProto.grpc.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/device/google/trout/hal/vehicle/2.0/VehicleServerProtoStub_cc@2.0-grpc-trout/gen/VehicleServer.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/device/google/trout/hal/vehicle/2.0/VehicleServerProtoStub_cc@2.0-grpc-trout/gen/VehicleServer.grpc.pb.cc
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
    add_library(android.hardware.automotive.vehicle-2.0-grpc-trout SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.automotive.vehicle-2.0-grpc-trout
            PUBLIC
            ${ANDROID_ROOT}/device/google/trout/hal/vehicle/2.0
            ${OUT_INTERMEDIATES_ROOT}/device/google/trout/hal/vehicle/2.0/VehicleServerProtoStub_cc@2.0-grpc-trout/gen
            ${OUT_INTERMEDIATES_ROOT}/device/google/trout/hal/vehicle/2.0/VehicleServerProtoStub_h@2.0-grpc-trout/gen
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.automotive.vehicle-2.0-grpc-trout PRIVATE
            -Wno-unused-parameter
    )

    target_link_libraries(android.hardware.automotive.vehicle-2.0-grpc-trout
            libc++_static
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.builtins
            libdl
            libgrpc++_unsecure
            libgrpc++_unsecure
            libm
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.automotive.vehicle-2.0-grpc-trout ####################

#################### android.hardware.automotive.vehicle-2.0-virtualization-grpc-server ####################
set(FILES
        ${ANDROID_ROOT}/device/google/trout/hal/vehicle/2.0/VirtualizationGrpcServer.cpp
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
    add_executable(android.hardware.automotive.vehicle-2.0-virtualization-grpc-server ${SOURCE_FILES})

    target_include_directories(android.hardware.automotive.vehicle-2.0-virtualization-grpc-server
            PUBLIC
            ${ANDROID_ROOT}/device/google/trout/hal/vehicle/2.0
            ${OUT_INTERMEDIATES_ROOT}/device/google/trout/hal/vehicle/2.0/VehicleServerProtoStub_h@2.0-grpc-trout/gen
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.automotive.vehicle-2.0-virtualization-grpc-server PRIVATE
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(android.hardware.automotive.vehicle-2.0-virtualization-grpc-server
            android.hardware.automotive.vehicle-2.0
            android.hardware.automotive.vehicle-2.0-virtualization-grpc-server-lib
            android.hardware.automotive.vehicle-2.0
            libbase
            libbase
            libc++
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libgrpc++
            libgrpc++
            libhidlbase
            libhidlbase
            libjsoncpp
            libjsoncpp
            liblog
            liblog
            libm
            libprotobuf-cpp-full
            libprotobuf-cpp-full
            libutils
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.automotive.vehicle-2.0-virtualization-grpc-server ####################
