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

project(vehicle_binding_util)

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

#################### libvehicle_binding_util ####################
set(FILES
        ${ANDROID_ROOT}/packages/services/Car/cpp/security/vehicle_binding_util/src/VehicleBindingUtil.cpp
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
    add_library(libvehicle_binding_util SHARED ${SOURCE_FILES})

    target_include_directories(libvehicle_binding_util
            PUBLIC
            ${ANDROID_ROOT}/packages/services/Car/cpp/security/vehicle_binding_util/src
            ${ANDROID_ROOT}/packages/services/Car/cpp/security/vehicle_binding_util
    )

    # Add c/c++ flags
    target_compile_options(libvehicle_binding_util PRIVATE
            -Wall
            -Wextra
            -Werror
            -Wall
            -Werror
            -Wno-unused-variable
            -Wall
            -Werror
            -Wno-unused-variable
    )

    target_link_libraries(libvehicle_binding_util
            VehicleHalUtils
            android-automotive-large-parcelable-lib
            android.hardware.automotive.vehicle-V2-ndk
            android.hardware.automotive.vehicle.property-V2-ndk
            android.hardware.automotive.vehicle-2.0
            libbase
            libbinder
            libbinder_ndk
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libhidlbase
            liblog
            liblogwrap
            libm
            libmath
            libutils
            libvhalclient
    )
endif ()
#################### libvehicle_binding_util ####################

#################### vehicle_binding_util ####################
set(FILES
        ${ANDROID_ROOT}/packages/services/Car/cpp/security/vehicle_binding_util/src/main.cpp
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
    add_executable(vehicle_binding_util ${SOURCE_FILES})

    target_include_directories(vehicle_binding_util
            PUBLIC
            ${ANDROID_ROOT}/packages/services/Car/cpp/security/vehicle_binding_util
            ${ANDROID_ROOT}/packages/services/Car/cpp/security/vehicle_binding_util/src
    )

    # Add c/c++ flags
    target_compile_options(vehicle_binding_util PRIVATE
            -Wall
            -Wextra
            -Werror
            -Wall
            -Werror
            -Wno-unused-variable
            -Wall
            -Werror
            -Wno-unused-variable
    )

    target_link_libraries(vehicle_binding_util
            VehicleHalUtils
            android-automotive-large-parcelable-lib
            android.hardware.automotive.vehicle-V2-ndk
            android.hardware.automotive.vehicle.property-V2-ndk
            android.hardware.automotive.vehicle-2.0
            libbase
            libbinder
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libhidlbase
            liblog
            liblogwrap
            libm
            libmath
            libutils
            libvehicle_binding_util
            libvhalclient
    )
endif ()
#################### vehicle_binding_util ####################
