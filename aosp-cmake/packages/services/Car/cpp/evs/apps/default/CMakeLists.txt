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

project(evs_app)

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

#################### libcartelemetry-evs-proto ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/evs/apps/default/libcartelemetry-evs-proto/android_${OUT_ARCH_CPU}_static/gen/proto/packages/services/Car/cpp/telemetry/proto/evs.pb.cc
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
    add_library(libcartelemetry-evs-proto SHARED ${SOURCE_FILES})

    target_include_directories(libcartelemetry-evs-proto
            PUBLIC
            ${ANDROID_ROOT}/packages/services/Car/cpp/evs/apps/default
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/evs/apps/default/libcartelemetry-evs-proto/android_${OUT_ARCH_CPU}_static/gen/proto/packages/services/Car/cpp/evs/apps/default
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/evs/apps/default/libcartelemetry-evs-proto/android_${OUT_ARCH_CPU}_static/gen/proto
    )

    # Add c/c++ flags
    target_compile_options(libcartelemetry-evs-proto PRIVATE
            -Werror
    )

    target_link_libraries(libcartelemetry-evs-proto
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libprotobuf-cpp-lite
    )
endif ()
#################### libcartelemetry-evs-proto ####################

#################### evs_app ####################
set(FILES
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/apps/default/src/ConfigManager.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/apps/default/src/EvsStateControl.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/apps/default/src/EvsStats.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/apps/default/src/FormatConvert.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/apps/default/src/RenderBase.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/apps/default/src/RenderDirectView.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/apps/default/src/RenderPixelCopy.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/apps/default/src/RenderTopView.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/apps/default/src/StreamHandler.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/apps/default/src/TexWrapper.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/apps/default/src/Utils.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/apps/default/src/VideoTex.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/apps/default/src/evs_app.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/apps/default/src/glError.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/apps/default/src/shader.cpp
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
    add_executable(evs_app ${SOURCE_FILES})

    target_include_directories(evs_app
            PUBLIC
            ${ANDROID_ROOT}/packages/services/Car/cpp/evs/apps/default/inc
            ${ANDROID_ROOT}/packages/services/Car/cpp/evs/apps/default
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/evs/apps/default/libcartelemetry-evs-proto/android_${OUT_ARCH_CPU}_static/gen/proto/packages/services/Car/cpp/evs/apps/default
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/evs/apps/default/libcartelemetry-evs-proto/android_${OUT_ARCH_CPU}_static/gen/proto
    )

    # Add c/c++ flags
    target_compile_options(evs_app PRIVATE
            -Wall
            -Wextra
            -Werror
            -Wall
            -Werror
            -Wno-unused-variable
            -Wall
            -Werror
    )

    target_link_libraries(evs_app
            VehicleHalUtils
            android-automotive-large-parcelable-lib
            android.frameworks.automotive.telemetry-V1-ndk
            android.hardware.automotive.evs-V2-ndk
            android.hardware.automotive.vehicle-V2-ndk
            android.hardware.automotive.vehicle.property-V2-ndk
            android.hardware.automotive.vehicle-2.0
            android.hardware.camera.device-3.2
            android.hardware.common-V2-ndk
            android.hardware.graphics.common-V4-ndk
            libEGL
            libGLESv2
            libaidlcommonsupport
            libbase
            libbinder
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libcamera_metadata
            libcartelemetry-evs-proto
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libhardware
            libhidlbase
            libjsoncpp
            liblog
            libm
            libmath
            libpng
            libprotobuf-cpp-lite
            libui
            libutils
            libvhalclient
    )
endif ()
#################### evs_app ####################
