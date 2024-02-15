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

project(vulkan.pastel.debug)

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

#################### vulkan.pastel.debug ####################
set(FILES
        ${ANDROID_ROOT}/external/swiftshader/src/System/Build.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/System/Configurator.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/System/CPUID.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/System/GrallocAndroid.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/System/Half.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/System/Linux/MemFd.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/System/Math.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/System/Memory.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/System/Socket.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/System/SwiftConfig.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/System/Timer.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Device/ASTC_Decoder.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Device/BC_Decoder.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Device/Blitter.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Device/Clipper.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Device/Context.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Device/ETC_Decoder.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Device/PixelProcessor.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Device/QuadRasterizer.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Device/Renderer.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Device/SetupProcessor.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Device/VertexProcessor.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Pipeline/ComputeProgram.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Pipeline/Constants.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Pipeline/PixelProgram.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Pipeline/PixelRoutine.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Pipeline/SamplerCore.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Pipeline/SetupRoutine.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Pipeline/ShaderCore.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Pipeline/SpirvBinary.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Pipeline/SpirvProfiler.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Pipeline/SpirvShader.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Pipeline/SpirvShaderArithmetic.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Pipeline/SpirvShaderControlFlow.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Pipeline/SpirvShaderDebugger.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Pipeline/SpirvShaderGLSLstd450.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Pipeline/SpirvShaderGroup.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Pipeline/SpirvShaderImage.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Pipeline/SpirvShaderInstructions.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Pipeline/SpirvShaderMemory.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Pipeline/SpirvShaderSampling.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Pipeline/SpirvShaderSpec.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Pipeline/VertexProgram.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Pipeline/VertexRoutine.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkBuffer.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkBufferView.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkCommandBuffer.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkCommandPool.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkDebugUtilsMessenger.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkDescriptorPool.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkDescriptorSet.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkDescriptorSetLayout.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkDescriptorUpdateTemplate.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkDevice.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkDeviceMemory.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkDeviceMemoryExternalAndroid.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkDeviceMemoryExternalHost.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkFormat.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkFramebuffer.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkGetProcAddress.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkImage.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkImageView.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkInstance.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkMemory.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkPhysicalDevice.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkPipeline.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkPipelineCache.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkPipelineLayout.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkPromotedExtensions.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkQueryPool.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkQueue.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkRenderPass.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkSampler.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkSemaphore.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkShaderModule.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkSpecializationInfo.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkStringify.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkTimelineSemaphore.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/libVulkan.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/main.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/WSI/HeadlessSurfaceKHR.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/WSI/VkSurfaceKHR.cpp
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
    add_library(vulkan.pastel.debug SHARED ${SOURCE_FILES})

    target_include_directories(vulkan.pastel.debug
            PUBLIC
            ${ANDROID_ROOT}/external/swiftshader/src/Vulkan
            ${ANDROID_ROOT}/external/swiftshader/src
            ${OUT_INTERMEDIATES_ROOT}/external/swiftshader/src/commit_header/gen
    )

    # Add c/c++ flags
    target_compile_options(vulkan.pastel.debug PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-local-typedef
            -UNDEBUG
    )

    target_link_libraries(vulkan.pastel.debug
            android.hardware.graphics.mapper-3.0
            android.hardware.graphics.mapper-4.0
            hwvulkan_headers
            libLLVM10_swiftshader_debug
            libarect
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libcutils
            libdl
            libhardware
            libhidlbase
            liblog
            libm
            libnativewindow
            libswiftshadervk_llvm_debug
            libsync
            libutils
            swiftshader_astc
            swiftshader_debug
            swiftshader_marl
            swiftshader_platform_headers
            swiftshader_spirv-tools
            vulkan_headers
            libc
            libdl
            libm
    )
endif ()
#################### vulkan.pastel.debug ####################

#################### swiftshader_debug ####################
set(FILES
        ${ANDROID_ROOT}/external/swiftshader/src/System/Debug.cpp
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
    add_library(swiftshader_debug SHARED ${SOURCE_FILES})

    target_include_directories(swiftshader_debug
            PUBLIC
            ${ANDROID_ROOT}/external/swiftshader/src
    )

    # Add c/c++ flags
    target_compile_options(swiftshader_debug PRIVATE
            -Werror
    )

    target_link_libraries(swiftshader_debug
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            liblog
            libm
    )
endif ()
#################### swiftshader_debug ####################

#################### libswiftshadervk_llvm_debug ####################
set(FILES
        ${ANDROID_ROOT}/external/swiftshader/src/Reactor/Assert.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Reactor/CPUID.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Reactor/Debug.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Reactor/ExecutableMemory.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Reactor/LLVMJIT.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Reactor/LLVMReactor.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Reactor/Pragma.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Reactor/Reactor.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Reactor/SIMD.cpp
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
    add_library(libswiftshadervk_llvm_debug SHARED ${SOURCE_FILES})

    target_include_directories(libswiftshadervk_llvm_debug
            PUBLIC
            ${ANDROID_ROOT}/external/swiftshader/src
    )

    # Add c/c++ flags
    target_compile_options(libswiftshadervk_llvm_debug PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libswiftshadervk_llvm_debug
            libLLVM10_swiftshader
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libcutils
            libdl
            liblog
            libm
            swiftshader_platform_headers
            libc
            libdl
            libm
    )
endif ()
#################### libswiftshadervk_llvm_debug ####################

#################### libswiftshadervk_llvm ####################
set(FILES
        ${ANDROID_ROOT}/external/swiftshader/src/Reactor/Assert.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Reactor/CPUID.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Reactor/Debug.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Reactor/ExecutableMemory.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Reactor/LLVMJIT.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Reactor/LLVMReactor.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Reactor/Pragma.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Reactor/Reactor.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Reactor/SIMD.cpp
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
    add_library(libswiftshadervk_llvm SHARED ${SOURCE_FILES})

    target_include_directories(libswiftshadervk_llvm
            PUBLIC
            ${ANDROID_ROOT}/external/swiftshader/src
    )

    # Add c/c++ flags
    target_compile_options(libswiftshadervk_llvm PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libswiftshadervk_llvm
            libLLVM10_swiftshader
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libcutils
            libdl
            liblog
            libm
            swiftshader_platform_headers
            libc
            libdl
            libm
    )
endif ()
#################### libswiftshadervk_llvm ####################

#################### vulkan.pastel ####################
set(FILES
        ${ANDROID_ROOT}/external/swiftshader/src/System/Build.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/System/Configurator.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/System/CPUID.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/System/GrallocAndroid.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/System/Half.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/System/Linux/MemFd.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/System/Math.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/System/Memory.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/System/Socket.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/System/SwiftConfig.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/System/Timer.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Device/ASTC_Decoder.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Device/BC_Decoder.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Device/Blitter.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Device/Clipper.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Device/Context.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Device/ETC_Decoder.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Device/PixelProcessor.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Device/QuadRasterizer.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Device/Renderer.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Device/SetupProcessor.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Device/VertexProcessor.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Pipeline/ComputeProgram.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Pipeline/Constants.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Pipeline/PixelProgram.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Pipeline/PixelRoutine.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Pipeline/SamplerCore.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Pipeline/SetupRoutine.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Pipeline/ShaderCore.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Pipeline/SpirvBinary.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Pipeline/SpirvProfiler.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Pipeline/SpirvShader.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Pipeline/SpirvShaderArithmetic.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Pipeline/SpirvShaderControlFlow.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Pipeline/SpirvShaderDebugger.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Pipeline/SpirvShaderGLSLstd450.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Pipeline/SpirvShaderGroup.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Pipeline/SpirvShaderImage.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Pipeline/SpirvShaderInstructions.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Pipeline/SpirvShaderMemory.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Pipeline/SpirvShaderSampling.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Pipeline/SpirvShaderSpec.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Pipeline/VertexProgram.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Pipeline/VertexRoutine.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkBuffer.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkBufferView.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkCommandBuffer.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkCommandPool.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkDebugUtilsMessenger.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkDescriptorPool.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkDescriptorSet.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkDescriptorSetLayout.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkDescriptorUpdateTemplate.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkDevice.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkDeviceMemory.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkDeviceMemoryExternalAndroid.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkDeviceMemoryExternalHost.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkFormat.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkFramebuffer.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkGetProcAddress.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkImage.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkImageView.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkInstance.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkMemory.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkPhysicalDevice.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkPipeline.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkPipelineCache.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkPipelineLayout.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkPromotedExtensions.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkQueryPool.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkQueue.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkRenderPass.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkSampler.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkSemaphore.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkShaderModule.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkSpecializationInfo.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkStringify.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/VkTimelineSemaphore.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/libVulkan.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/Vulkan/main.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/WSI/HeadlessSurfaceKHR.cpp
        ${ANDROID_ROOT}/external/swiftshader/src/WSI/VkSurfaceKHR.cpp
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
    add_library(vulkan.pastel SHARED ${SOURCE_FILES})

    target_include_directories(vulkan.pastel
            PUBLIC
            ${ANDROID_ROOT}/external/swiftshader/src/Vulkan
            ${ANDROID_ROOT}/external/swiftshader/src
            ${OUT_INTERMEDIATES_ROOT}/external/swiftshader/src/commit_header/gen
    )

    # Add c/c++ flags
    target_compile_options(vulkan.pastel PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-local-typedef
    )

    target_link_libraries(vulkan.pastel
            android.hardware.graphics.mapper-3.0
            android.hardware.graphics.mapper-4.0
            hwvulkan_headers
            libLLVM10_swiftshader
            libarect
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libcutils
            libdl
            libhardware
            libhidlbase
            liblog
            libm
            libnativewindow
            libswiftshadervk_llvm_debug
            libsync
            libutils
            swiftshader_astc
            swiftshader_debug
            swiftshader_marl
            swiftshader_platform_headers
            swiftshader_spirv-tools
            vulkan_headers
            libc
            libdl
            libm
    )
endif ()
#################### vulkan.pastel ####################
