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

project(android.frameworks.automotive.powerpolicy.CarPowerPolicyServer.fuzzer)

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

#################### lib_carpowerpolicyserver ####################
set(FILES
        ${ANDROID_ROOT}/packages/services/Car/cpp/powerpolicy/server/src/CarPowerPolicyServer.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/powerpolicy/server/src/PolicyManager.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/powerpolicy/server/src/PowerComponentHandler.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/powerpolicy/server/src/SilentModeHandler.cpp
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
    add_library(lib_carpowerpolicyserver SHARED ${SOURCE_FILES})

    target_include_directories(lib_carpowerpolicyserver
            PUBLIC
            ${ANDROID_ROOT}/packages/services/Car/cpp/powerpolicy/server/src
            ${ANDROID_ROOT}/packages/services/Car/cpp/powerpolicy/server
    )

    # Add c/c++ flags
    target_compile_options(lib_carpowerpolicyserver PRIVATE
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

    target_link_libraries(lib_carpowerpolicyserver
            VehicleHalUtils
            android-automotive-large-parcelable-lib
            android.frameworks.automotive.powerpolicy-V2-ndk
            android.frameworks.automotive.powerpolicy.internal-V1-cpp
            android.frameworks.automotive.powerpolicy.internal-V1-ndk
            android.hardware.automotive.vehicle-V2-ndk
            android.hardware.automotive.vehicle.property-V2-ndk
            android.hardware.automotive.vehicle-2.0
            libbase
            libbinder
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libc++fs
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libhidlbase
            liblog
            libm
            libmath
            libsysfsmonitor
            libtinyxml2
            libutils
            libvhalclient
    )
endif ()
#################### lib_carpowerpolicyserver ####################

#################### carpowerpolicyd ####################
set(FILES
        ${ANDROID_ROOT}/packages/services/Car/cpp/powerpolicy/server/src/main.cpp
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
    add_executable(carpowerpolicyd ${SOURCE_FILES})

    target_include_directories(carpowerpolicyd
            PUBLIC
            ${ANDROID_ROOT}/packages/services/Car/cpp/powerpolicy/server
            ${ANDROID_ROOT}/packages/services/Car/cpp/powerpolicy/server/src
    )

    # Add c/c++ flags
    target_compile_options(carpowerpolicyd PRIVATE
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

    target_link_libraries(carpowerpolicyd
            VehicleHalUtils
            android-automotive-large-parcelable-lib
            android.frameworks.automotive.powerpolicy-V2-ndk
            android.frameworks.automotive.powerpolicy.internal-V1-cpp
            android.frameworks.automotive.powerpolicy.internal-V1-ndk
            android.hardware.automotive.vehicle-V2-ndk
            android.hardware.automotive.vehicle.property-V2-ndk
            android.hardware.automotive.vehicle-2.0
            lib_carpowerpolicyserver
            libbase
            libbinder
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libc++fs
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libhidlbase
            liblog
            libm
            libmath
            libsysfsmonitor
            libtinyxml2
            libutils
            libvhalclient
    )
endif ()
#################### carpowerpolicyd ####################

#################### android.frameworks.automotive.powerpolicy.CarPowerPolicyServer.fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/packages/services/Car/cpp/powerpolicy/server/src/fuzzer.cpp
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
    add_executable(android.frameworks.automotive.powerpolicy.CarPowerPolicyServer.fuzzer ${SOURCE_FILES})

    target_include_directories(android.frameworks.automotive.powerpolicy.CarPowerPolicyServer.fuzzer
            PUBLIC
            ${ANDROID_ROOT}/packages/services/Car/cpp/powerpolicy/server
            ${ANDROID_ROOT}/packages/services/Car/cpp/powerpolicy/server/src
    )

    # Add c/c++ flags
    target_compile_options(android.frameworks.automotive.powerpolicy.CarPowerPolicyServer.fuzzer PRIVATE
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

    target_link_libraries(android.frameworks.automotive.powerpolicy.CarPowerPolicyServer.fuzzer
            VehicleHalUtils
            android-automotive-large-parcelable-lib
            android.frameworks.automotive.powerpolicy-V2-ndk
            android.frameworks.automotive.powerpolicy.internal-V1-cpp
            android.frameworks.automotive.powerpolicy.internal-V1-ndk
            android.hardware.automotive.vehicle-V2-ndk
            android.hardware.automotive.vehicle.property-V2-ndk
            android.hardware.automotive.vehicle-2.0
            lib_carpowerpolicyserver
            libbase
            libbinder
            libbinder_ndk
            libbinder_random_parcel
            libc
            libc++
            libc++demangle
            libc++fs
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcutils
            libdl
            libhidlbase
            liblog
            libm
            libmath
            libsysfsmonitor
            libtinyxml2
            libutils
            libvhalclient
    )
endif ()
#################### android.frameworks.automotive.powerpolicy.CarPowerPolicyServer.fuzzer ####################
