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

project(android.hardware.automotive.remoteaccess@V1-default-service)

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

#################### android.hardware.automotive.remoteaccess-V1-default-service.aidl_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/automotive/remoteaccess/hal/default/fuzzer/fuzzer.cpp
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
    add_executable(android.hardware.automotive.remoteaccess-V1-default-service.aidl_fuzzer ${SOURCE_FILES})

    target_include_directories(android.hardware.automotive.remoteaccess-V1-default-service.aidl_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/automotive/remoteaccess/hal/default
            ${ANDROID_ROOT}/hardware/interfaces/automotive/remoteaccess/hal/default/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/remoteaccess/hal/default/proto/wakeup_client_pb_h/gen
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.automotive.remoteaccess-V1-default-service.aidl_fuzzer PRIVATE
            -Wall
            -Wextra
            -Werror
            -Wall
            -Werror
            -Wno-unused-variable
            -Wno-unused-parameter
    )

    target_link_libraries(android.hardware.automotive.remoteaccess-V1-default-service.aidl_fuzzer
            RemoteAccessService
            VehicleHalUtils
            android-automotive-large-parcelable-lib
            android.hardware.automotive.vehicle-V2-ndk
            android.hardware.automotive.vehicle.property-V2-ndk
            android.hardware.automotive.vehicle-2.0
            libbase
            libbinder
            libbinder_ndk
            libbinder_random_parcel
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcutils
            libdl
            libgmock
            libgrpc++
            libgtest
            libhidlbase
            liblog
            libm
            libmath
            libprotobuf-cpp-full
            libutils
    )
endif ()
#################### android.hardware.automotive.remoteaccess-V1-default-service.aidl_fuzzer ####################

#################### android.hardware.automotive.remoteaccess-V1-default-service ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/automotive/remoteaccess/hal/default/src/RemoteAccessImpl.cpp
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
    add_executable(android.hardware.automotive.remoteaccess-V1-default-service ${SOURCE_FILES})

    target_include_directories(android.hardware.automotive.remoteaccess-V1-default-service
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/automotive/remoteaccess/hal/default
            ${ANDROID_ROOT}/hardware/interfaces/automotive/remoteaccess/hal/default/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/remoteaccess/hal/default/proto/wakeup_client_pb_h/gen
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.automotive.remoteaccess-V1-default-service PRIVATE
            -Wno-unused-parameter
            -Wall
            -Wextra
            -Werror
            -Wall
            -Werror
            -Wno-unused-variable
            -Wno-unused-parameter
    )

    target_link_libraries(android.hardware.automotive.remoteaccess-V1-default-service
            BindToDeviceSocketMutatorLib
            RemoteAccessService
            VehicleHalUtils
            android-automotive-large-parcelable-lib
            android.hardware.automotive.can-libnetdevice
            android.hardware.automotive.vehicle-V2-ndk
            android.hardware.automotive.vehicle.property-V2-ndk
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
            liblog
            libm
            libmath
            libnl++
            libprotobuf-cpp-full
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.automotive.remoteaccess-V1-default-service ####################

#################### RemoteAccessService ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/automotive/remoteaccess/hal/default/src/RemoteAccessService.cpp
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
    add_library(RemoteAccessService SHARED ${SOURCE_FILES})

    target_include_directories(RemoteAccessService
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/automotive/remoteaccess/hal/default/include
            ${ANDROID_ROOT}/hardware/interfaces/automotive/remoteaccess/hal/default
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/remoteaccess/hal/default/proto/wakeup_client_pb_h/gen
    )

    # Add c/c++ flags
    target_compile_options(RemoteAccessService PRIVATE
            -Wall
            -Wextra
            -Werror
            -Wall
            -Werror
            -Wno-unused-variable
            -Wno-unused-parameter
    )

    target_link_libraries(RemoteAccessService
            VehicleHalUtils
            android-automotive-large-parcelable-lib
            android.hardware.automotive.remoteaccess-V1-ndk
            android.hardware.automotive.vehicle-V2-ndk
            android.hardware.automotive.vehicle.property-V2-ndk
            android.hardware.automotive.vehicle-2.0
            libbase
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libgrpc++
            libhidlbase
            liblog
            libm
            libmath
            libprotobuf-cpp-full
            libutils
            libvhalclient
            wakeup_client_protos
    )
endif ()
#################### RemoteAccessService ####################
