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

project(libtensorflowlite_jni)

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

#################### libtensorflowlite_jni ####################
set(FILES
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/java/src/main/native/interpreter_factory_impl_jni.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/java/src/main/native/jni_utils.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/java/src/main/native/nativeinterpreterwrapper_jni.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/java/src/main/native/nativeinterpreterwrapperexperimental_jni.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/java/src/main/native/nativesignaturerunner_jni.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/java/src/main/native/op_resolver_lazy_delegate_proxy.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/java/src/main/native/tensor_jni.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/java/src/main/native/tensorflow_lite_jni.cc
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
    add_library(libtensorflowlite_jni SHARED ${SOURCE_FILES})

    target_include_directories(libtensorflowlite_jni
            PUBLIC
            ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/java/src/main/native
            ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/java
    )

    # Add c/c++ flags
    target_compile_options(libtensorflowlite_jni PRIVATE
            -Werror
            -Wno-unused-function
            -Wno-unused-parameter
    )

    target_link_libraries(libtensorflowlite_jni
            flatbuffer_headers
            jni_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            liblog
            libm
            libruy_static
            libtflite_static
            tensorflow_headers
    )
endif ()
#################### libtensorflowlite_jni ####################
