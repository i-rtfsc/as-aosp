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

project(libtflite_flex_delegate_only_runtime)

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

#################### libtflite_flex_delegate_only_runtime ####################
set(FILES
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/delegates/flex/buffer_map.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/delegates/flex/buffer_map_util.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/delegates/flex/delegate.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/delegates/flex/delegate_data.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/delegates/flex/kernel.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/delegates/flex/util.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/delegates/utils/simple_delegate.cc
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
    add_library(libtflite_flex_delegate_only_runtime SHARED ${SOURCE_FILES})

    target_include_directories(libtflite_flex_delegate_only_runtime
            PUBLIC
            ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/delegates
    )

    # Add c/c++ flags
    target_compile_options(libtflite_flex_delegate_only_runtime PRIVATE
            -Werror
            -Wno-unused-const-variable
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
    )

    target_link_libraries(libtflite_flex_delegate_only_runtime
            flatbuffer_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libeigen
            libm
            libtensorflow_core_mobile_srcs_only_runtime
            libtextclassifier_hash_headers
            tensorflow_abseil
            tensorflow_core_proto_cpp_lite
            tensorflow_headers
    )
endif ()
#################### libtflite_flex_delegate_only_runtime ####################
