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

project(libchrome-gestures)

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

#################### libchrome-gestures ####################
set(FILES
        ${ANDROID_ROOT}/external/libchrome-gestures/src/accel_filter_interpreter.cc
        ${ANDROID_ROOT}/external/libchrome-gestures/src/activity_log.cc
        ${ANDROID_ROOT}/external/libchrome-gestures/src/box_filter_interpreter.cc
        ${ANDROID_ROOT}/external/libchrome-gestures/src/click_wiggle_filter_interpreter.cc
        ${ANDROID_ROOT}/external/libchrome-gestures/src/file_util.cc
        ${ANDROID_ROOT}/external/libchrome-gestures/src/filter_interpreter.cc
        ${ANDROID_ROOT}/external/libchrome-gestures/src/finger_merge_filter_interpreter.cc
        ${ANDROID_ROOT}/external/libchrome-gestures/src/finger_metrics.cc
        ${ANDROID_ROOT}/external/libchrome-gestures/src/fling_stop_filter_interpreter.cc
        ${ANDROID_ROOT}/external/libchrome-gestures/src/gestures.cc
        ${ANDROID_ROOT}/external/libchrome-gestures/src/haptic_button_generator_filter_interpreter.cc
        ${ANDROID_ROOT}/external/libchrome-gestures/src/iir_filter_interpreter.cc
        ${ANDROID_ROOT}/external/libchrome-gestures/src/immediate_interpreter.cc
        ${ANDROID_ROOT}/external/libchrome-gestures/src/integral_gesture_filter_interpreter.cc
        ${ANDROID_ROOT}/external/libchrome-gestures/src/interpreter.cc
        ${ANDROID_ROOT}/external/libchrome-gestures/src/logging_filter_interpreter.cc
        ${ANDROID_ROOT}/external/libchrome-gestures/src/lookahead_filter_interpreter.cc
        ${ANDROID_ROOT}/external/libchrome-gestures/src/metrics_filter_interpreter.cc
        ${ANDROID_ROOT}/external/libchrome-gestures/src/mouse_interpreter.cc
        ${ANDROID_ROOT}/external/libchrome-gestures/src/multitouch_mouse_interpreter.cc
        ${ANDROID_ROOT}/external/libchrome-gestures/src/non_linearity_filter_interpreter.cc
        ${ANDROID_ROOT}/external/libchrome-gestures/src/palm_classifying_filter_interpreter.cc
        ${ANDROID_ROOT}/external/libchrome-gestures/src/prop_registry.cc
        ${ANDROID_ROOT}/external/libchrome-gestures/src/scaling_filter_interpreter.cc
        ${ANDROID_ROOT}/external/libchrome-gestures/src/sensor_jump_filter_interpreter.cc
        ${ANDROID_ROOT}/external/libchrome-gestures/src/split_correcting_filter_interpreter.cc
        ${ANDROID_ROOT}/external/libchrome-gestures/src/stationary_wiggle_filter_interpreter.cc
        ${ANDROID_ROOT}/external/libchrome-gestures/src/string_util.cc
        ${ANDROID_ROOT}/external/libchrome-gestures/src/stuck_button_inhibitor_filter_interpreter.cc
        ${ANDROID_ROOT}/external/libchrome-gestures/src/t5r2_correcting_filter_interpreter.cc
        ${ANDROID_ROOT}/external/libchrome-gestures/src/timestamp_filter_interpreter.cc
        ${ANDROID_ROOT}/external/libchrome-gestures/src/trace_marker.cc
        ${ANDROID_ROOT}/external/libchrome-gestures/src/tracer.cc
        ${ANDROID_ROOT}/external/libchrome-gestures/src/trend_classifying_filter_interpreter.cc
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
    add_library(libchrome-gestures SHARED ${SOURCE_FILES})

    target_include_directories(libchrome-gestures
            PUBLIC
            ${ANDROID_ROOT}/external/libchrome-gestures
    )

    # Add c/c++ flags
    target_compile_options(libchrome-gestures PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libchrome-gestures
            libc
            libc++_static
            libc++demangle
            libchrome-gestures_headers
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libgtest_prod_headers
            libjsoncpp
            libm
    )
endif ()
#################### libchrome-gestures ####################
