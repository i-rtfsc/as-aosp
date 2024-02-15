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

project(tracing_vm_protos)

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

#################### tracing_vm_protos ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/device/google/trout/tools/tracing/proto/TracingVMProtoStub_cc/gen/TracingVM.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/device/google/trout/tools/tracing/proto/TracingVMProtoStub_cc/gen/TracingVM.grpc.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/device/google/trout/tools/tracing/proto/perfetto_trace_gen_cc/gen/perfetto_trace.pb.cc
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
    add_library(tracing_vm_protos SHARED ${SOURCE_FILES})

    target_include_directories(tracing_vm_protos
            PUBLIC
            ${ANDROID_ROOT}/device/google/trout/tools/tracing/proto
            ${OUT_INTERMEDIATES_ROOT}/device/google/trout/tools/tracing/proto/TracingVMProtoStub_h/gen
            ${OUT_INTERMEDIATES_ROOT}/device/google/trout/tools/tracing/proto/perfetto_trace_gen_h/gen
    )

    # Add c/c++ flags
    target_compile_options(tracing_vm_protos PRIVATE
            -Wno-unused-parameter
    )

    target_link_libraries(tracing_vm_protos
            libc++_static
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.builtins
            libdl
            libgrpc++
            libgrpc++
            libm
            libprotobuf-cpp-lite
            libprotobuf-cpp-lite
            libc
            libdl
            libm
    )
endif ()
#################### tracing_vm_protos ####################
