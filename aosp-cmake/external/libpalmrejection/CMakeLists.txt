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

project(libpalmrejection)

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

#################### libpalmrejection ####################
set(FILES
        ${ANDROID_ROOT}/external/libpalmrejection/chrome_to_android_compatibility.cc
        ${ANDROID_ROOT}/external/libpalmrejection/ui/events/ozone/features.cc
        ${ANDROID_ROOT}/external/libpalmrejection/ui/events/ozone/evdev/touch_evdev_types.cc
        ${ANDROID_ROOT}/external/libpalmrejection/ui/events/ozone/evdev/touch_filter/neural_stylus_palm_detection_filter.cc
        ${ANDROID_ROOT}/external/libpalmrejection/ui/events/ozone/evdev/touch_filter/neural_stylus_palm_detection_filter_model.cc
        ${ANDROID_ROOT}/external/libpalmrejection/ui/events/ozone/evdev/touch_filter/neural_stylus_palm_detection_filter_util.cc
        ${ANDROID_ROOT}/external/libpalmrejection/ui/events/ozone/evdev/touch_filter/palm_detection_filter.cc
        ${ANDROID_ROOT}/external/libpalmrejection/ui/events/ozone/evdev/touch_filter/shared_palm_detection_filter_state.cc
        ${ANDROID_ROOT}/external/libpalmrejection/ui/events/ozone/evdev/touch_filter/palm_model/onedevice_train_palm_detection_filter_model.cc
        ${ANDROID_ROOT}/external/libpalmrejection/ui/events/ozone/evdev/touch_filter/palm_model/onedevice_train_palm_detection_filter_inference.cc
        ${ANDROID_ROOT}/external/libpalmrejection/ui/events/ozone/evdev/touch_filter/palm_model/onedevice_train_palm_detection_filter_inference_beta.cc
        ${ANDROID_ROOT}/external/libpalmrejection/ui/events/ozone/evdev/touch_filter/palm_model/onedevice_train_palm_detection_filter_inference_v2.cc
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
    add_library(libpalmrejection SHARED ${SOURCE_FILES})

    target_include_directories(libpalmrejection
            PUBLIC
            ${ANDROID_ROOT}/external/libpalmrejection
    )

    # Add c/c++ flags
    target_compile_options(libpalmrejection PRIVATE
            -Wall
            -Wextra
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libpalmrejection
            libc
            libc++_static
            libc++demangle
            libchrome
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### libpalmrejection ####################
