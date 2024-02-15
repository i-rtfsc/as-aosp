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

project(deqp_amber)

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

#################### deqp_amber ####################
set(FILES
        ${ANDROID_ROOT}/external/deqp-deps/amber/src/amber.cc
        ${ANDROID_ROOT}/external/deqp-deps/amber/src/amberscript/parser.cc
        ${ANDROID_ROOT}/external/deqp-deps/amber/src/buffer.cc
        ${ANDROID_ROOT}/external/deqp-deps/amber/src/command_data.cc
        ${ANDROID_ROOT}/external/deqp-deps/amber/src/command.cc
        ${ANDROID_ROOT}/external/deqp-deps/amber/src/debug.cc
        ${ANDROID_ROOT}/external/deqp-deps/amber/src/descriptor_set_and_binding_parser.cc
        ${ANDROID_ROOT}/external/deqp-deps/amber/src/engine.cc
        ${ANDROID_ROOT}/external/deqp-deps/amber/src/executor.cc
        ${ANDROID_ROOT}/external/deqp-deps/amber/src/float16_helper.cc
        ${ANDROID_ROOT}/external/deqp-deps/amber/src/format.cc
        ${ANDROID_ROOT}/external/deqp-deps/amber/src/parser.cc
        ${ANDROID_ROOT}/external/deqp-deps/amber/src/pipeline_data.cc
        ${ANDROID_ROOT}/external/deqp-deps/amber/src/pipeline.cc
        ${ANDROID_ROOT}/external/deqp-deps/amber/src/recipe.cc
        ${ANDROID_ROOT}/external/deqp-deps/amber/src/result.cc
        ${ANDROID_ROOT}/external/deqp-deps/amber/src/sampler.cc
        ${ANDROID_ROOT}/external/deqp-deps/amber/src/script.cc
        ${ANDROID_ROOT}/external/deqp-deps/amber/src/shader_compiler.cc
        ${ANDROID_ROOT}/external/deqp-deps/amber/src/shader.cc
        ${ANDROID_ROOT}/external/deqp-deps/amber/src/sleep.cc
        ${ANDROID_ROOT}/external/deqp-deps/amber/src/tokenizer.cc
        ${ANDROID_ROOT}/external/deqp-deps/amber/src/type_parser.cc
        ${ANDROID_ROOT}/external/deqp-deps/amber/src/type.cc
        ${ANDROID_ROOT}/external/deqp-deps/amber/src/value.cc
        ${ANDROID_ROOT}/external/deqp-deps/amber/src/verifier.cc
        ${ANDROID_ROOT}/external/deqp-deps/amber/src/virtual_file_store.cc
        ${ANDROID_ROOT}/external/deqp-deps/amber/src/vkscript/command_parser.cc
        ${ANDROID_ROOT}/external/deqp-deps/amber/src/vkscript/datum_type_parser.cc
        ${ANDROID_ROOT}/external/deqp-deps/amber/src/vkscript/parser.cc
        ${ANDROID_ROOT}/external/deqp-deps/amber/src/vkscript/section_parser.cc
        ${ANDROID_ROOT}/external/deqp-deps/amber/src/vulkan_engine_config.cc
        ${ANDROID_ROOT}/external/deqp-deps/amber/src/vulkan/buffer_backed_descriptor.cc
        ${ANDROID_ROOT}/external/deqp-deps/amber/src/vulkan/buffer_descriptor.cc
        ${ANDROID_ROOT}/external/deqp-deps/amber/src/vulkan/command_buffer.cc
        ${ANDROID_ROOT}/external/deqp-deps/amber/src/vulkan/command_pool.cc
        ${ANDROID_ROOT}/external/deqp-deps/amber/src/vulkan/compute_pipeline.cc
        ${ANDROID_ROOT}/external/deqp-deps/amber/src/vulkan/descriptor.cc
        ${ANDROID_ROOT}/external/deqp-deps/amber/src/vulkan/device.cc
        ${ANDROID_ROOT}/external/deqp-deps/amber/src/vulkan/engine_vulkan_debugger.cc
        ${ANDROID_ROOT}/external/deqp-deps/amber/src/vulkan/engine_vulkan.cc
        ${ANDROID_ROOT}/external/deqp-deps/amber/src/vulkan/frame_buffer.cc
        ${ANDROID_ROOT}/external/deqp-deps/amber/src/vulkan/graphics_pipeline.cc
        ${ANDROID_ROOT}/external/deqp-deps/amber/src/vulkan/image_descriptor.cc
        ${ANDROID_ROOT}/external/deqp-deps/amber/src/vulkan/index_buffer.cc
        ${ANDROID_ROOT}/external/deqp-deps/amber/src/vulkan/pipeline.cc
        ${ANDROID_ROOT}/external/deqp-deps/amber/src/vulkan/push_constant.cc
        ${ANDROID_ROOT}/external/deqp-deps/amber/src/vulkan/resource.cc
        ${ANDROID_ROOT}/external/deqp-deps/amber/src/vulkan/sampler_descriptor.cc
        ${ANDROID_ROOT}/external/deqp-deps/amber/src/vulkan/sampler.cc
        ${ANDROID_ROOT}/external/deqp-deps/amber/src/vulkan/transfer_buffer.cc
        ${ANDROID_ROOT}/external/deqp-deps/amber/src/vulkan/transfer_image.cc
        ${ANDROID_ROOT}/external/deqp-deps/amber/src/vulkan/vertex_buffer.cc
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
    add_library(deqp_amber SHARED ${SOURCE_FILES})

    target_include_directories(deqp_amber
            PUBLIC
            ${ANDROID_ROOT}/external/deqp-deps/amber/include
            ${ANDROID_ROOT}/external/deqp-deps/amber
            ${OUT_INTERMEDIATES_ROOT}/external/deqp-deps/amber/deqp_amber_gen_vk_wrappers/gen
    )

    # Add c/c++ flags
    target_compile_options(deqp_amber PRIVATE
            -DAMBER_ENABLE_VK_DEBUGGING=0
            -Wall
            -Werror
    )

    target_link_libraries(deqp_amber
            deqp_vulkan_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### deqp_amber ####################
