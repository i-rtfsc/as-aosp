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

project(libpowermanager)

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

#################### libpowermanager ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/services/powermanager/BatterySaverPolicyConfig.cpp
        ${ANDROID_ROOT}/frameworks/native/services/powermanager/CoolingDevice.cpp
        ${ANDROID_ROOT}/frameworks/native/services/powermanager/ParcelDuration.cpp
        ${ANDROID_ROOT}/frameworks/native/services/powermanager/PowerHalController.cpp
        ${ANDROID_ROOT}/frameworks/native/services/powermanager/PowerHalLoader.cpp
        ${ANDROID_ROOT}/frameworks/native/services/powermanager/PowerHalWrapper.cpp
        ${ANDROID_ROOT}/frameworks/native/services/powermanager/PowerSaveState.cpp
        ${ANDROID_ROOT}/frameworks/native/services/powermanager/Temperature.cpp
        ${ANDROID_ROOT}/frameworks/native/services/powermanager/WorkSource.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/services/powermanager/libpowermanager/android_${OUT_ARCH_CPU}_shared/gen/aidl/android/os/Temperature.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/services/powermanager/libpowermanager/android_${OUT_ARCH_CPU}_shared/gen/aidl/android/os/CoolingDevice.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/services/powermanager/libpowermanager/android_${OUT_ARCH_CPU}_shared/gen/aidl/android/os/IHintManager.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/services/powermanager/libpowermanager/android_${OUT_ARCH_CPU}_shared/gen/aidl/android/os/IHintSession.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/services/powermanager/libpowermanager/android_${OUT_ARCH_CPU}_shared/gen/aidl/android/os/IThermalEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/services/powermanager/libpowermanager/android_${OUT_ARCH_CPU}_shared/gen/aidl/android/os/IThermalStatusListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/services/powermanager/libpowermanager/android_${OUT_ARCH_CPU}_shared/gen/aidl/android/os/IThermalService.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/services/powermanager/libpowermanager/android_${OUT_ARCH_CPU}_shared/gen/aidl/android/os/IPowerManager.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/services/powermanager/libpowermanager/android_${OUT_ARCH_CPU}_shared/gen/aidl/android/os/IWakeLockCallback.cpp
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
    add_library(libpowermanager SHARED ${SOURCE_FILES})

    target_include_directories(libpowermanager
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/services/powermanager/include
            ${ANDROID_ROOT}/frameworks/native/services/powermanager
            ${OUT_INTERMEDIATES_ROOT}/frameworks/native/services/powermanager/libpowermanager/android_${OUT_ARCH_CPU}_shared/gen/aidl
    )

    # Add c/c++ flags
    target_compile_options(libpowermanager PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libpowermanager
            android.hardware.power-V4-cpp
            android.hardware.power-1.0
            android.hardware.power-1.1
            android.hardware.power-1.2
            android.hardware.power-1.3
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libhidlbase
            liblog
            libm
            libutils
    )
endif ()
#################### libpowermanager ####################
