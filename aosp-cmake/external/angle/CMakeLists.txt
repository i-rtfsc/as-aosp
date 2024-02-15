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

project(libGLESv2_angle)

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

#################### libGLESv2_angle ####################
set(FILES
        ${ANDROID_ROOT}/external/angle/third_party/vulkan_memory_allocator/vma.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/vulkan/vk_mem_alloc_wrapper.cpp
        ${ANDROID_ROOT}/external/angle/src/common/vulkan/vulkan_icd.cpp
        ${ANDROID_ROOT}/external/angle/src/third_party/volk/volk.c
        ${ANDROID_ROOT}/external/angle/src/common/vulkan/libvulkan_loader.cpp
        ${ANDROID_ROOT}/external/angle/src/common/spirv/spirv_instruction_parser_autogen.cpp
        ${ANDROID_ROOT}/external/angle/src/common/spirv/spirv_instruction_builder_autogen.cpp
        ${ANDROID_ROOT}/external/angle/src/common/spirv/angle_spirv_utils.cpp
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/container/internal/hashtablez_sampler.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/container/internal/hashtablez_sampler_force_weak_definition.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/container/internal/raw_hash_set.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/time/internal/cctz/src/time_zone_fixed.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/time/internal/cctz/src/time_zone_format.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/time/internal/cctz/src/time_zone_if.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/time/internal/cctz/src/time_zone_impl.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/time/internal/cctz/src/time_zone_info.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/time/internal/cctz/src/time_zone_libc.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/time/internal/cctz/src/time_zone_lookup.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/time/internal/cctz/src/time_zone_posix.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/time/internal/cctz/src/zone_info_source.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/time/internal/cctz/src/civil_time_detail.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/time/civil_time.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/time/clock.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/time/duration.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/time/format.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/time/time.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/synchronization/internal/kernel_timeout.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/synchronization/internal/graphcycles.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/debugging/internal/demangle.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/debugging/symbolize.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/base/internal/low_level_alloc.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/synchronization/barrier.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/synchronization/blocking_counter.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/synchronization/internal/create_thread_identity.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/synchronization/internal/futex_waiter.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/synchronization/internal/per_thread_sem.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/synchronization/internal/pthread_waiter.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/synchronization/internal/sem_waiter.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/synchronization/internal/stdcpp_waiter.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/synchronization/internal/waiter_base.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/synchronization/internal/win32_waiter.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/synchronization/mutex.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/synchronization/notification.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/internal/cordz_handle.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/debugging/internal/address_is_readable.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/debugging/internal/elf_mem_image.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/debugging/internal/vdso_support.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/debugging/stacktrace.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/internal/cordz_info.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/profiling/internal/exponential_biased.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/internal/cordz_functions.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/internal/cord_internal.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/internal/cord_rep_btree.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/internal/cord_rep_btree_navigator.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/internal/cord_rep_btree_reader.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/internal/cord_rep_consume.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/internal/cord_rep_crc.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/internal/cord_rep_ring.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/crc/internal/crc.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/crc/internal/crc_x86_arm_combined.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/crc/internal/cpu_detect.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/crc/crc32c.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/crc/internal/crc_memcpy_fallback.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/crc/internal/crc_memcpy_x86_64.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/crc/internal/crc_non_temporal_memcpy.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/crc/internal/crc_cord_state.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/cord.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/cord_analysis.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/cord_buffer.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/types/bad_variant_access.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/types/bad_optional_access.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/internal/escaping.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/internal/ostringstream.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/internal/utf8.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/ascii.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/charconv.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/escaping.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/internal/charconv_bigint.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/internal/charconv_parse.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/internal/damerau_levenshtein_distance.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/internal/memutil.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/internal/stringify_sink.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/match.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/numbers.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/str_cat.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/str_replace.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/str_split.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/string_view.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/substitute.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/numeric/int128.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/hash/internal/low_level_hash.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/hash/internal/city.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/base/internal/throw_delegate.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/base/internal/spinlock_wait.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/base/internal/raw_logging.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/base/log_severity.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/base/internal/cycleclock.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/base/internal/spinlock.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/base/internal/sysinfo.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/base/internal/thread_identity.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/base/internal/unscaledcycleclock.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/hash/internal/hash.cc
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/vulkan/AllocatorHelperPool.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/vulkan/AllocatorHelperRing.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/vulkan/BufferVk.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/vulkan/CommandProcessor.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/vulkan/CompilerVk.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/vulkan/ContextVk.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/vulkan/DebugAnnotatorVk.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/vulkan/DeviceVk.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/vulkan/DisplayVk.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/vulkan/FenceNVVk.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/vulkan/FramebufferVk.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/vulkan/ImageVk.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/vulkan/MemoryObjectVk.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/vulkan/MemoryTracking.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/vulkan/OverlayVk.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/vulkan/PersistentCommandPool.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/vulkan/ProgramExecutableVk.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/vulkan/ProgramPipelineVk.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/vulkan/ProgramVk.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/vulkan/QueryVk.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/vulkan/RenderTargetVk.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/vulkan/RenderbufferVk.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/vulkan/RendererVk.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/vulkan/ResourceVk.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/vulkan/SamplerVk.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/vulkan/SecondaryCommandBuffer.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/vulkan/SecondaryCommandPool.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/vulkan/SemaphoreVk.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/vulkan/ShaderInterfaceVariableInfoMap.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/vulkan/ShaderVk.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/vulkan/Suballocation.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/vulkan/SurfaceVk.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/vulkan/SyncVk.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/vulkan/TextureVk.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/vulkan/TransformFeedbackVk.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/vulkan/UtilsVk.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/vulkan/VertexArrayVk.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/vulkan/VkImageImageSiblingVk.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/vulkan/VulkanSecondaryCommandBuffer.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/vulkan/android/AHBFunctions.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/vulkan/android/DisplayVkAndroid.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/vulkan/android/HardwareBufferImageSiblingVkAndroid.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/vulkan/android/WindowSurfaceVkAndroid.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/vulkan/android/vk_android_utils.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/vulkan/spv_utils.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/vulkan/vk_cache_utils.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/vulkan/vk_caps_utils.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/vulkan/vk_format_table_autogen.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/vulkan/vk_format_utils.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/vulkan/vk_helpers.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/vulkan/vk_internal_shaders_autogen.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/vulkan/vk_mandatory_format_support_table_autogen.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/vulkan/vk_utils.cpp
        ${ANDROID_ROOT}/external/angle/src/common/angle_version_info.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/AttributeMap.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/BlobCache.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/Buffer.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/Caps.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/Compiler.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/Config.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/Context.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/Context_gl.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/Context_gles_1_0.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/Debug.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/Device.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/Display.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/EGLSync.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/Error.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/Fence.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/Framebuffer.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/FramebufferAttachment.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/GLES1Renderer.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/GLES1State.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/HandleAllocator.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/Image.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/ImageIndex.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/IndexRangeCache.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/LoggingAnnotator.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/MemoryObject.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/MemoryProgramCache.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/MemoryShaderCache.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/Observer.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/Overlay.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/OverlayWidgets.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/Overlay_autogen.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/Overlay_font_autogen.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/PixelLocalStorage.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/Platform.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/Program.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/ProgramExecutable.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/ProgramLinkedResources.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/ProgramPipeline.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/Query.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/Renderbuffer.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/ResourceManager.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/Sampler.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/Semaphore.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/Shader.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/State.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/Stream.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/Surface.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/Texture.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/Thread.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/TransformFeedback.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/Uniform.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/VaryingPacking.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/VertexArray.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/VertexAttribute.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/angletypes.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/es3_copy_conversion_table_autogen.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/format_map_autogen.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/format_map_desktop.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/formatutils.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/gles_extensions_autogen.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/queryconversions.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/queryutils.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/BufferImpl.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/ContextImpl.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/DeviceImpl.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/DisplayImpl.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/EGLReusableSync.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/EGLSyncImpl.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/Format_table_autogen.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/FramebufferImpl.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/ImageImpl.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/ProgramImpl.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/ProgramPipelineImpl.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/QueryImpl.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/RenderbufferImpl.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/ShaderImpl.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/SurfaceImpl.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/TextureImpl.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/TransformFeedbackImpl.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/VertexArrayImpl.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/driver_utils.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/load_functions_table_autogen.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/renderer/renderer_utils.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/validationEGL.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/validationES.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/validationES1.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/validationES2.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/validationES3.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/validationES31.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/validationES32.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/validationESEXT.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/validationGL1.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/validationGL2.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/validationGL3.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/validationGL4.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/entry_points_utils.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/capture/FrameCapture_mock.cpp
        ${ANDROID_ROOT}/external/angle/src/libANGLE/capture/serialize_mock.cpp
        ${ANDROID_ROOT}/external/angle/src/libGLESv2/egl_ext_stubs.cpp
        ${ANDROID_ROOT}/external/angle/src/libGLESv2/egl_stubs.cpp
        ${ANDROID_ROOT}/external/angle/src/libGLESv2/entry_points_egl_autogen.cpp
        ${ANDROID_ROOT}/external/angle/src/libGLESv2/entry_points_egl_ext_autogen.cpp
        ${ANDROID_ROOT}/external/angle/src/libGLESv2/entry_points_gles_1_0_autogen.cpp
        ${ANDROID_ROOT}/external/angle/src/libGLESv2/entry_points_gles_2_0_autogen.cpp
        ${ANDROID_ROOT}/external/angle/src/libGLESv2/entry_points_gles_3_0_autogen.cpp
        ${ANDROID_ROOT}/external/angle/src/libGLESv2/entry_points_gles_3_1_autogen.cpp
        ${ANDROID_ROOT}/external/angle/src/libGLESv2/entry_points_gles_3_2_autogen.cpp
        ${ANDROID_ROOT}/external/angle/src/libGLESv2/entry_points_gles_ext_autogen.cpp
        ${ANDROID_ROOT}/external/angle/src/libGLESv2/global_state.cpp
        ${ANDROID_ROOT}/external/angle/src/libGLESv2/libGLESv2_autogen.cpp
        ${ANDROID_ROOT}/external/angle/src/libGLESv2/proc_table_egl_autogen.cpp
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
    add_library(libGLESv2_angle SHARED ${SOURCE_FILES})

    target_include_directories(libGLESv2_angle
            PUBLIC
            ${ANDROID_ROOT}/external/angle
            ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/vulkan-headers/src/include
            ${ANDROID_ROOT}/external/angle/third_party/vulkan_memory_allocator/include
            ${ANDROID_ROOT}/external/angle/include
            ${ANDROID_ROOT}/external/angle/src
            ${ANDROID_ROOT}/external/angle/src/third_party/volk
            ${ANDROID_ROOT}/external/angle/src/common/base
            ${ANDROID_ROOT}/external/angle/src/common/third_party/xxhash
            ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp
            ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-headers/src/include
            ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/include
            ${OUT_INTERMEDIATES_ROOT}/external/angle/angle_commit_id/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/angle_program_version_id/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_core_enums_unified1/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_core_tables_unified1/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_generators_inc/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_glsl_tables_glsl1-0/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_language_header_cldebuginfo100/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_language_header_debuginfo/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_language_header_vkdebuginfo100/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_opencl_tables_opencl1-0/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_vendor_tables_debuginfo/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_vendor_tables_nonsemantic.clspvreflection/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_vendor_tables_nonsemantic.shader.debuginfo.100/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_vendor_tables_opencl.debuginfo.100/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_vendor_tables_spv-amd-gcn-shader/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_vendor_tables_spv-amd-shader-ballot/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_vendor_tables_spv-amd-shader-explicit-vertex-parameter/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_vendor_tables_spv-amd-shader-trinary-minmax/android_${OUT_ARCH_CPU}/gen
    )

    # Add c/c++ flags
    target_compile_options(libGLESv2_angle PRIVATE
            -Werror
            -DNDEBUG
            -Wno-unused-parameter
            -Wno-unused-private-field
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-parameter
            -Wno-unused-private-field
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-parameter
            -DNDEBUG
            -Wno-unused-parameter
            -DNDEBUG
            -Wno-unused-parameter
            -DNDEBUG
            -Wno-unused-parameter
            -DNDEBUG
            -Wno-unused-parameter
            -DNDEBUG
            -Wno-unused-parameter
            -DNDEBUG
            -Wno-unused-parameter
            -DNDEBUG
            -Wno-unused-parameter
            -DNDEBUG
            -Wno-unused-parameter
            -DNDEBUG
            -Wno-unused-parameter
            -DNDEBUG
            -Wno-unused-parameter
            -DNDEBUG
            -Wno-unused-parameter
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-parameter
            -DNDEBUG
            -Wno-unused-parameter
            -DNDEBUG
            -Wno-unused-parameter
            -DNDEBUG
            -Wno-unused-parameter
            -DNDEBUG
            -Wno-unused-parameter
            -DNDEBUG
            -Wno-unused-parameter
            -DNDEBUG
            -Wno-unused-parameter
            -DNDEBUG
            -Wno-unused-parameter
            -DNDEBUG
            -Wno-unused-parameter
            -DNDEBUG
            -Wno-unused-parameter
    )

    target_link_libraries(libGLESv2_angle
            angle_common
            angle_common_shader_state
            angle_gpu_info_util
            angle_image_util
            angle_preprocessor
            angle_translator
            cpufeatures
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            liblog
            libm
            libnativewindow
            libz_static
            third_party_vulkan_deps_spirv_tools_src_spvtools
            third_party_vulkan_deps_spirv_tools_src_spvtools_val
            zlib_google_compression_utils_portable
    )
endif ()
#################### libGLESv2_angle ####################

#################### angle_preprocessor ####################
set(FILES
        ${ANDROID_ROOT}/external/angle/src/compiler/preprocessor/DiagnosticsBase.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/preprocessor/DirectiveHandlerBase.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/preprocessor/DirectiveParser.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/preprocessor/Input.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/preprocessor/Lexer.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/preprocessor/Macro.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/preprocessor/MacroExpander.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/preprocessor/Preprocessor.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/preprocessor/Token.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/preprocessor/preprocessor_lex_autogen.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/preprocessor/preprocessor_tab_autogen.cpp
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
    add_library(angle_preprocessor SHARED ${SOURCE_FILES})

    target_include_directories(angle_preprocessor
            PUBLIC
            ${ANDROID_ROOT}/external/angle/include
            ${ANDROID_ROOT}/external/angle/src
            ${ANDROID_ROOT}/external/angle/src/common/base
            ${ANDROID_ROOT}/external/angle/src/common/third_party/xxhash
            ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp
            ${ANDROID_ROOT}/external/angle
    )

    # Add c/c++ flags
    target_compile_options(angle_preprocessor PRIVATE
            -Werror
            -DNDEBUG
            -Wno-unused-parameter
            -DNDEBUG
            -Wno-unused-parameter
            -DNDEBUG
            -Wno-unused-parameter
    )

    target_link_libraries(angle_preprocessor
            angle_common
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            liblog
            libm
            libnativewindow
    )
endif ()
#################### angle_preprocessor ####################

#################### third_party_vulkan_deps_spirv_tools_src_spvtools ####################
set(FILES
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/assembly_grammar.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/binary.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/diagnostic.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/disassemble.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/enum_string_mapping.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/ext_inst.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/extensions.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/libspirv.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/name_mapper.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/opcode.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/operand.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/parsed_operand.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/print.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/spirv_endian.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/spirv_fuzzer_options.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/spirv_optimizer_options.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/spirv_reducer_options.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/spirv_target_env.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/spirv_validator_options.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/table.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/text.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/text_handler.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/util/bit_vector.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/util/parse_number.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/util/string_utils.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/util/timer.cpp
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
    add_library(third_party_vulkan_deps_spirv_tools_src_spvtools SHARED ${SOURCE_FILES})

    target_include_directories(third_party_vulkan_deps_spirv_tools_src_spvtools
            PUBLIC
            ${ANDROID_ROOT}/external/angle
            ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/include
            ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-headers/src/include
            ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_core_enums_unified1/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_core_tables_unified1/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_generators_inc/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_glsl_tables_glsl1-0/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_language_header_cldebuginfo100/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_language_header_debuginfo/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_language_header_vkdebuginfo100/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_opencl_tables_opencl1-0/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_vendor_tables_debuginfo/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_vendor_tables_nonsemantic.clspvreflection/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_vendor_tables_nonsemantic.shader.debuginfo.100/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_vendor_tables_opencl.debuginfo.100/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_vendor_tables_spv-amd-gcn-shader/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_vendor_tables_spv-amd-shader-ballot/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_vendor_tables_spv-amd-shader-explicit-vertex-parameter/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_vendor_tables_spv-amd-shader-trinary-minmax/android_${OUT_ARCH_CPU}/gen
    )

    # Add c/c++ flags
    target_compile_options(third_party_vulkan_deps_spirv_tools_src_spvtools PRIVATE
            -Werror
            -DNDEBUG
            -Wno-unused-parameter
            -DNDEBUG
            -Wno-unused-parameter
            -DNDEBUG
            -Wno-unused-parameter
    )

    target_link_libraries(third_party_vulkan_deps_spirv_tools_src_spvtools
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### third_party_vulkan_deps_spirv_tools_src_spvtools ####################

#################### angle_common ####################
set(FILES
        ${ANDROID_ROOT}/external/angle/src/common/third_party/xxhash/xxhash.c
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/container/internal/hashtablez_sampler.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/container/internal/hashtablez_sampler_force_weak_definition.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/container/internal/raw_hash_set.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/time/internal/cctz/src/time_zone_fixed.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/time/internal/cctz/src/time_zone_format.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/time/internal/cctz/src/time_zone_if.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/time/internal/cctz/src/time_zone_impl.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/time/internal/cctz/src/time_zone_info.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/time/internal/cctz/src/time_zone_libc.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/time/internal/cctz/src/time_zone_lookup.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/time/internal/cctz/src/time_zone_posix.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/time/internal/cctz/src/zone_info_source.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/time/internal/cctz/src/civil_time_detail.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/time/civil_time.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/time/clock.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/time/duration.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/time/format.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/time/time.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/synchronization/internal/kernel_timeout.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/synchronization/internal/graphcycles.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/debugging/internal/demangle.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/debugging/symbolize.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/base/internal/low_level_alloc.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/synchronization/barrier.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/synchronization/blocking_counter.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/synchronization/internal/create_thread_identity.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/synchronization/internal/futex_waiter.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/synchronization/internal/per_thread_sem.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/synchronization/internal/pthread_waiter.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/synchronization/internal/sem_waiter.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/synchronization/internal/stdcpp_waiter.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/synchronization/internal/waiter_base.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/synchronization/internal/win32_waiter.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/synchronization/mutex.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/synchronization/notification.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/internal/cordz_handle.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/debugging/internal/address_is_readable.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/debugging/internal/elf_mem_image.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/debugging/internal/vdso_support.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/debugging/stacktrace.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/internal/cordz_info.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/profiling/internal/exponential_biased.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/internal/cordz_functions.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/internal/cord_internal.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/internal/cord_rep_btree.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/internal/cord_rep_btree_navigator.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/internal/cord_rep_btree_reader.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/internal/cord_rep_consume.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/internal/cord_rep_crc.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/internal/cord_rep_ring.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/crc/internal/crc.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/crc/internal/crc_x86_arm_combined.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/crc/internal/cpu_detect.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/crc/crc32c.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/crc/internal/crc_memcpy_fallback.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/crc/internal/crc_memcpy_x86_64.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/crc/internal/crc_non_temporal_memcpy.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/crc/internal/crc_cord_state.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/cord.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/cord_analysis.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/cord_buffer.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/types/bad_variant_access.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/types/bad_optional_access.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/internal/escaping.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/internal/ostringstream.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/internal/utf8.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/ascii.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/charconv.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/escaping.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/internal/charconv_bigint.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/internal/charconv_parse.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/internal/damerau_levenshtein_distance.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/internal/memutil.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/internal/stringify_sink.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/match.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/numbers.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/str_cat.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/str_replace.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/str_split.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/string_view.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/substitute.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/numeric/int128.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/hash/internal/low_level_hash.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/hash/internal/city.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/base/internal/throw_delegate.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/base/internal/spinlock_wait.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/base/internal/raw_logging.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/base/log_severity.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/base/internal/cycleclock.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/base/internal/spinlock.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/base/internal/sysinfo.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/base/internal/thread_identity.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/base/internal/unscaledcycleclock.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/hash/internal/hash.cc
        ${ANDROID_ROOT}/external/angle/src/common/Float16ToFloat32.cpp
        ${ANDROID_ROOT}/external/angle/src/common/MemoryBuffer.cpp
        ${ANDROID_ROOT}/external/angle/src/common/PackedEGLEnums_autogen.cpp
        ${ANDROID_ROOT}/external/angle/src/common/PackedEnums.cpp
        ${ANDROID_ROOT}/external/angle/src/common/PackedGLEnums_autogen.cpp
        ${ANDROID_ROOT}/external/angle/src/common/PoolAlloc.cpp
        ${ANDROID_ROOT}/external/angle/src/common/RingBufferAllocator.cpp
        ${ANDROID_ROOT}/external/angle/src/common/WorkerThread.cpp
        ${ANDROID_ROOT}/external/angle/src/common/aligned_memory.cpp
        ${ANDROID_ROOT}/external/angle/src/common/android_util.cpp
        ${ANDROID_ROOT}/external/angle/src/common/angleutils.cpp
        ${ANDROID_ROOT}/external/angle/src/common/backtrace_utils_noop.cpp
        ${ANDROID_ROOT}/external/angle/src/common/base/anglebase/sha1.cc
        ${ANDROID_ROOT}/external/angle/src/common/debug.cpp
        ${ANDROID_ROOT}/external/angle/src/common/entry_points_enum_autogen.cpp
        ${ANDROID_ROOT}/external/angle/src/common/event_tracer.cpp
        ${ANDROID_ROOT}/external/angle/src/common/mathutil.cpp
        ${ANDROID_ROOT}/external/angle/src/common/matrix_utils.cpp
        ${ANDROID_ROOT}/external/angle/src/common/string_utils.cpp
        ${ANDROID_ROOT}/external/angle/src/common/system_utils.cpp
        ${ANDROID_ROOT}/external/angle/src/common/system_utils_linux.cpp
        ${ANDROID_ROOT}/external/angle/src/common/system_utils_posix.cpp
        ${ANDROID_ROOT}/external/angle/src/common/tls.cpp
        ${ANDROID_ROOT}/external/angle/src/common/uniform_type_info_autogen.cpp
        ${ANDROID_ROOT}/external/angle/src/common/utilities.cpp
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
    add_library(angle_common SHARED ${SOURCE_FILES})

    target_include_directories(angle_common
            PUBLIC
            ${ANDROID_ROOT}/external/angle/include
            ${ANDROID_ROOT}/external/angle/src
            ${ANDROID_ROOT}/external/angle
            ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp
            ${ANDROID_ROOT}/external/angle/src/common/base
            ${ANDROID_ROOT}/external/angle/src/common/third_party/xxhash
    )

    # Add c/c++ flags
    target_compile_options(angle_common PRIVATE
            -Werror
            -DNDEBUG
            -Wno-unused-parameter
            -DNDEBUG
            -Wno-unused-parameter
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-parameter
    )

    target_link_libraries(angle_common
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            liblog
            libm
            libnativewindow
    )
endif ()
#################### angle_common ####################

#################### angle_image_util ####################
set(FILES
        ${ANDROID_ROOT}/external/angle/src/image_util/AstcDecompressorNoOp.cpp
        ${ANDROID_ROOT}/external/angle/src/image_util/copyimage.cpp
        ${ANDROID_ROOT}/external/angle/src/image_util/imageformats.cpp
        ${ANDROID_ROOT}/external/angle/src/image_util/loadimage.cpp
        ${ANDROID_ROOT}/external/angle/src/image_util/loadimage_astc.cpp
        ${ANDROID_ROOT}/external/angle/src/image_util/loadimage_etc.cpp
        ${ANDROID_ROOT}/external/angle/src/image_util/loadimage_paletted.cpp
        ${ANDROID_ROOT}/external/angle/src/image_util/storeimage_paletted.cpp
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
    add_library(angle_image_util SHARED ${SOURCE_FILES})

    target_include_directories(angle_image_util
            PUBLIC
            ${ANDROID_ROOT}/external/angle/include
            ${ANDROID_ROOT}/external/angle/src
            ${ANDROID_ROOT}/external/angle/src/common/base
            ${ANDROID_ROOT}/external/angle/src/common/third_party/xxhash
            ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp
            ${ANDROID_ROOT}/external/angle
    )

    # Add c/c++ flags
    target_compile_options(angle_image_util PRIVATE
            -Werror
            -DNDEBUG
            -Wno-unused-parameter
            -DNDEBUG
            -Wno-unused-parameter
    )

    target_link_libraries(angle_image_util
            angle_common
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            liblog
            libm
            libnativewindow
    )
endif ()
#################### angle_image_util ####################

#################### libGLESv1_CM_angle ####################
set(FILES
        ${ANDROID_ROOT}/external/angle/src/libGLESv1_CM/libGLESv1_CM.cpp
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
    add_library(libGLESv1_CM_angle SHARED ${SOURCE_FILES})

    target_include_directories(libGLESv1_CM_angle
            PUBLIC
            ${ANDROID_ROOT}/external/angle/include
            ${ANDROID_ROOT}/external/angle/src
            ${ANDROID_ROOT}/external/angle
            ${OUT_INTERMEDIATES_ROOT}/external/angle/angle_commit_id/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/angle_program_version_id/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_core_enums_unified1/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_core_tables_unified1/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_generators_inc/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_glsl_tables_glsl1-0/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_language_header_cldebuginfo100/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_language_header_debuginfo/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_language_header_vkdebuginfo100/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_opencl_tables_opencl1-0/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_vendor_tables_debuginfo/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_vendor_tables_nonsemantic.clspvreflection/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_vendor_tables_nonsemantic.shader.debuginfo.100/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_vendor_tables_opencl.debuginfo.100/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_vendor_tables_spv-amd-gcn-shader/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_vendor_tables_spv-amd-shader-ballot/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_vendor_tables_spv-amd-shader-explicit-vertex-parameter/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_vendor_tables_spv-amd-shader-trinary-minmax/android_${OUT_ARCH_CPU}/gen
    )

    # Add c/c++ flags
    target_compile_options(libGLESv1_CM_angle PRIVATE
            -Werror
            -DNDEBUG
            -Wno-unused-parameter
            -DNDEBUG
            -Wno-unused-parameter
    )

    target_link_libraries(libGLESv1_CM_angle
            angle_common
            angle_common_shader_state
            angle_gpu_info_util
            angle_image_util
            angle_preprocessor
            angle_translator
            cpufeatures
            libGLESv2_angle
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libz_static
            third_party_vulkan_deps_spirv_tools_src_spvtools
            third_party_vulkan_deps_spirv_tools_src_spvtools_val
            zlib_google_compression_utils_portable
    )
endif ()
#################### libGLESv1_CM_angle ####################

#################### libEGL_angle ####################
set(FILES
        ${ANDROID_ROOT}/external/angle/src/libEGL/egl_loader_autogen.cpp
        ${ANDROID_ROOT}/external/angle/src/libEGL/libEGL_autogen.cpp
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
    add_library(libEGL_angle SHARED ${SOURCE_FILES})

    target_include_directories(libEGL_angle
            PUBLIC
            ${ANDROID_ROOT}/external/angle/include
            ${ANDROID_ROOT}/external/angle/src
            ${ANDROID_ROOT}/external/angle/src/common/base
            ${ANDROID_ROOT}/external/angle/src/common/third_party/xxhash
            ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp
            ${ANDROID_ROOT}/external/angle
            ${OUT_INTERMEDIATES_ROOT}/external/angle/angle_commit_id/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/angle_program_version_id/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_core_enums_unified1/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_core_tables_unified1/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_generators_inc/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_glsl_tables_glsl1-0/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_language_header_cldebuginfo100/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_language_header_debuginfo/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_language_header_vkdebuginfo100/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_opencl_tables_opencl1-0/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_vendor_tables_debuginfo/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_vendor_tables_nonsemantic.clspvreflection/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_vendor_tables_nonsemantic.shader.debuginfo.100/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_vendor_tables_opencl.debuginfo.100/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_vendor_tables_spv-amd-gcn-shader/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_vendor_tables_spv-amd-shader-ballot/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_vendor_tables_spv-amd-shader-explicit-vertex-parameter/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_vendor_tables_spv-amd-shader-trinary-minmax/android_${OUT_ARCH_CPU}/gen
    )

    # Add c/c++ flags
    target_compile_options(libEGL_angle PRIVATE
            -Werror
            -DNDEBUG
            -Wno-unused-parameter
            -DNDEBUG
            -Wno-unused-parameter
            -DNDEBUG
            -Wno-unused-parameter
    )

    target_link_libraries(libEGL_angle
            angle_common
            angle_common_shader_state
            angle_gpu_info_util
            angle_image_util
            angle_preprocessor
            angle_translator
            cpufeatures
            libGLESv2_angle
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            liblog
            libm
            libnativewindow
            libz_static
            third_party_vulkan_deps_spirv_tools_src_spvtools
            third_party_vulkan_deps_spirv_tools_src_spvtools_val
            zlib_google_compression_utils_portable
    )
endif ()
#################### libEGL_angle ####################

#################### angle_translator ####################
set(FILES
        ${ANDROID_ROOT}/external/angle/src/common/spirv/spirv_instruction_builder_autogen.cpp
        ${ANDROID_ROOT}/external/angle/src/common/spirv/angle_spirv_utils.cpp
        ${ANDROID_ROOT}/external/angle/src/common/angle_version_info.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/BaseTypes.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/BuildSPIRV.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/BuiltInFunctionEmulator.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/BuiltinsWorkaroundGLSL.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/CallDAG.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/CodeGen.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/CollectVariables.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/Compiler.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/ConstantUnion.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/Declarator.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/Diagnostics.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/DirectiveHandler.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/ExtensionBehavior.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/FlagStd140Structs.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/FunctionLookup.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/HashNames.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/ImmutableStringBuilder.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/ImmutableString_ESSL_autogen.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/InfoSink.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/Initialize.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/InitializeDll.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/IntermNode.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/IsASTDepthBelowLimit.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/Operator.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/OutputGLSL.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/OutputGLSLBase.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/OutputSPIRV.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/OutputTree.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/ParseContext.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/PoolAlloc.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/QualifierTypes.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/ShaderLang.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/ShaderVars.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/Symbol.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/SymbolTable.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/SymbolTable_ESSL_autogen.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/SymbolUniqueId.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/TranslatorVulkan.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/Types.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/ValidateAST.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/ValidateBarrierFunctionCall.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/ValidateClipCullDistance.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/ValidateGlobalInitializer.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/ValidateLimitations.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/ValidateMaxParameters.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/ValidateOutputs.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/ValidateSwitch.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/ValidateTypeSizeLimitations.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/ValidateVaryingLocations.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/VariablePacker.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/blocklayout.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/glslang_lex_autogen.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/glslang_tab_autogen.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/tree_ops/ClampIndirectIndices.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/tree_ops/ClampPointSize.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/tree_ops/ConvertUnsupportedConstructorsToFunctionCalls.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/tree_ops/DeclareAndInitBuiltinsForInstancedMultiview.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/tree_ops/DeclarePerVertexBlocks.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/tree_ops/DeferGlobalInitializers.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/tree_ops/EmulateGLFragColorBroadcast.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/tree_ops/EmulateMultiDrawShaderBuiltins.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/tree_ops/FoldExpressions.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/tree_ops/ForcePrecisionQualifier.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/tree_ops/InitializeVariables.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/tree_ops/MonomorphizeUnsupportedFunctions.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/tree_ops/PruneEmptyCases.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/tree_ops/PruneNoOps.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/tree_ops/RecordConstantPrecision.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/tree_ops/RemoveArrayLengthMethod.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/tree_ops/RemoveAtomicCounterBuiltins.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/tree_ops/RemoveDynamicIndexing.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/tree_ops/RemoveInactiveInterfaceVariables.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/tree_ops/RemoveInvariantDeclaration.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/tree_ops/RemoveUnreferencedVariables.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/tree_ops/RewriteArrayOfArrayOfOpaqueUniforms.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/tree_ops/RewriteAtomicCounters.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/tree_ops/RewriteCubeMapSamplersAs2DArray.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/tree_ops/RewriteDfdy.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/tree_ops/RewritePixelLocalStorage.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/tree_ops/RewriteStructSamplers.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/tree_ops/RewriteTexelFetchOffset.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/tree_ops/ScalarizeVecAndMatConstructorArgs.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/tree_ops/SeparateDeclarations.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/tree_ops/SeparateStructFromUniformDeclarations.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/tree_ops/SimplifyLoopConditions.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/tree_ops/SplitSequenceOperator.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/tree_ops/vulkan/EmulateAdvancedBlendEquations.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/tree_ops/vulkan/EmulateDithering.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/tree_ops/vulkan/EmulateFragColorData.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/tree_ops/vulkan/EmulateYUVBuiltIns.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/tree_ops/vulkan/FlagSamplersWithTexelFetch.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/tree_ops/vulkan/ReplaceForShaderFramebufferFetch.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/tree_ops/vulkan/RewriteInterpolateAtOffset.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/tree_ops/vulkan/RewriteR32fImages.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/tree_util/DriverUniform.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/tree_util/FindFunction.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/tree_util/FindMain.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/tree_util/FindPreciseNodes.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/tree_util/FindSymbolNode.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/tree_util/IntermNodePatternMatcher.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/tree_util/IntermNode_util.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/tree_util/IntermRebuild.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/tree_util/IntermTraverse.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/tree_util/ReplaceArrayOfMatrixVarying.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/tree_util/ReplaceClipCullDistanceVariable.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/tree_util/ReplaceShadowingVariables.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/tree_util/ReplaceVariable.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/tree_util/RewriteSampleMaskVariable.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/tree_util/RunAtTheBeginningOfShader.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/tree_util/RunAtTheEndOfShader.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/tree_util/SpecializationConstant.cpp
        ${ANDROID_ROOT}/external/angle/src/compiler/translator/util.cpp
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
    add_library(angle_translator SHARED ${SOURCE_FILES})

    target_include_directories(angle_translator
            PUBLIC
            ${ANDROID_ROOT}/external/angle/include
            ${ANDROID_ROOT}/external/angle/src
            ${ANDROID_ROOT}/external/angle/src/common/base
            ${ANDROID_ROOT}/external/angle/src/common/third_party/xxhash
            ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp
            ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-headers/src/include
            ${ANDROID_ROOT}/external/angle
            ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/include
            ${OUT_INTERMEDIATES_ROOT}/external/angle/angle_commit_id/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/angle_program_version_id/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_core_enums_unified1/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_core_tables_unified1/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_generators_inc/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_glsl_tables_glsl1-0/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_language_header_cldebuginfo100/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_language_header_debuginfo/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_language_header_vkdebuginfo100/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_opencl_tables_opencl1-0/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_vendor_tables_debuginfo/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_vendor_tables_nonsemantic.clspvreflection/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_vendor_tables_nonsemantic.shader.debuginfo.100/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_vendor_tables_opencl.debuginfo.100/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_vendor_tables_spv-amd-gcn-shader/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_vendor_tables_spv-amd-shader-ballot/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_vendor_tables_spv-amd-shader-explicit-vertex-parameter/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_vendor_tables_spv-amd-shader-trinary-minmax/android_${OUT_ARCH_CPU}/gen
    )

    # Add c/c++ flags
    target_compile_options(angle_translator PRIVATE
            -Werror
            -DNDEBUG
            -Wno-unused-parameter
            -DNDEBUG
            -Wno-unused-parameter
            -DNDEBUG
            -Wno-unused-parameter
            -DNDEBUG
            -Wno-unused-parameter
            -DNDEBUG
            -Wno-unused-parameter
            -DNDEBUG
            -Wno-unused-parameter
            -DNDEBUG
            -Wno-unused-parameter
            -DNDEBUG
            -Wno-unused-parameter
            -DNDEBUG
            -Wno-unused-parameter
            -DNDEBUG
            -Wno-unused-parameter
    )

    target_link_libraries(angle_translator
            angle_common
            angle_common_shader_state
            angle_preprocessor
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            liblog
            libm
            libnativewindow
            third_party_vulkan_deps_spirv_tools_src_spvtools
            third_party_vulkan_deps_spirv_tools_src_spvtools_val
    )
endif ()
#################### angle_translator ####################

#################### angle_common_shader_state ####################
set(FILES
        ${ANDROID_ROOT}/external/angle/src/common/third_party/xxhash/xxhash.c
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/container/internal/hashtablez_sampler.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/container/internal/hashtablez_sampler_force_weak_definition.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/container/internal/raw_hash_set.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/time/internal/cctz/src/time_zone_fixed.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/time/internal/cctz/src/time_zone_format.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/time/internal/cctz/src/time_zone_if.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/time/internal/cctz/src/time_zone_impl.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/time/internal/cctz/src/time_zone_info.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/time/internal/cctz/src/time_zone_libc.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/time/internal/cctz/src/time_zone_lookup.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/time/internal/cctz/src/time_zone_posix.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/time/internal/cctz/src/zone_info_source.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/time/internal/cctz/src/civil_time_detail.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/time/civil_time.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/time/clock.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/time/duration.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/time/format.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/time/time.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/synchronization/internal/kernel_timeout.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/synchronization/internal/graphcycles.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/debugging/internal/demangle.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/debugging/symbolize.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/base/internal/low_level_alloc.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/synchronization/barrier.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/synchronization/blocking_counter.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/synchronization/internal/create_thread_identity.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/synchronization/internal/futex_waiter.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/synchronization/internal/per_thread_sem.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/synchronization/internal/pthread_waiter.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/synchronization/internal/sem_waiter.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/synchronization/internal/stdcpp_waiter.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/synchronization/internal/waiter_base.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/synchronization/internal/win32_waiter.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/synchronization/mutex.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/synchronization/notification.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/internal/cordz_handle.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/debugging/internal/address_is_readable.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/debugging/internal/elf_mem_image.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/debugging/internal/vdso_support.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/debugging/stacktrace.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/internal/cordz_info.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/profiling/internal/exponential_biased.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/internal/cordz_functions.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/internal/cord_internal.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/internal/cord_rep_btree.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/internal/cord_rep_btree_navigator.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/internal/cord_rep_btree_reader.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/internal/cord_rep_consume.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/internal/cord_rep_crc.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/internal/cord_rep_ring.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/crc/internal/crc.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/crc/internal/crc_x86_arm_combined.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/crc/internal/cpu_detect.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/crc/crc32c.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/crc/internal/crc_memcpy_fallback.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/crc/internal/crc_memcpy_x86_64.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/crc/internal/crc_non_temporal_memcpy.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/crc/internal/crc_cord_state.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/cord.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/cord_analysis.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/cord_buffer.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/types/bad_variant_access.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/types/bad_optional_access.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/internal/escaping.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/internal/ostringstream.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/internal/utf8.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/ascii.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/charconv.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/escaping.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/internal/charconv_bigint.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/internal/charconv_parse.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/internal/damerau_levenshtein_distance.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/internal/memutil.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/internal/stringify_sink.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/match.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/numbers.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/str_cat.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/str_replace.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/str_split.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/string_view.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/strings/substitute.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/numeric/int128.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/hash/internal/low_level_hash.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/hash/internal/city.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/base/internal/throw_delegate.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/base/internal/spinlock_wait.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/base/internal/raw_logging.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/base/log_severity.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/base/internal/cycleclock.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/base/internal/spinlock.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/base/internal/sysinfo.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/base/internal/thread_identity.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/base/internal/unscaledcycleclock.cc
        ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp/absl/hash/internal/hash.cc
        ${ANDROID_ROOT}/external/angle/src/common/CompiledShaderState.cpp
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
    add_library(angle_common_shader_state SHARED ${SOURCE_FILES})

    target_include_directories(angle_common_shader_state
            PUBLIC
            ${ANDROID_ROOT}/external/angle/include
            ${ANDROID_ROOT}/external/angle/src
            ${ANDROID_ROOT}/external/angle
            ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp
            ${ANDROID_ROOT}/external/angle/src/common/base
            ${ANDROID_ROOT}/external/angle/src/common/third_party/xxhash
    )

    # Add c/c++ flags
    target_compile_options(angle_common_shader_state PRIVATE
            -Werror
            -DNDEBUG
            -Wno-unused-parameter
            -DNDEBUG
            -Wno-unused-parameter
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
            -Wno-unused-variable
            -DNDEBUG
            -Wno-unused-parameter
    )

    target_link_libraries(angle_common_shader_state
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            liblog
            libm
            libnativewindow
    )
endif ()
#################### angle_common_shader_state ####################

#################### angle_gpu_info_util ####################
set(FILES
        ${ANDROID_ROOT}/external/angle/src/third_party/volk/volk.c
        ${ANDROID_ROOT}/external/angle/src/common/vulkan/vulkan_icd.cpp
        ${ANDROID_ROOT}/external/angle/src/common/vulkan/libvulkan_loader.cpp
        ${ANDROID_ROOT}/external/angle/src/gpu_info_util/SystemInfo.cpp
        ${ANDROID_ROOT}/external/angle/src/gpu_info_util/SystemInfo_android.cpp
        ${ANDROID_ROOT}/external/angle/src/gpu_info_util/SystemInfo_vulkan.cpp
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
    add_library(angle_gpu_info_util SHARED ${SOURCE_FILES})

    target_include_directories(angle_gpu_info_util
            PUBLIC
            ${ANDROID_ROOT}/external/angle
            ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/vulkan-headers/src/include
            ${ANDROID_ROOT}/external/angle/src/third_party/volk
            ${ANDROID_ROOT}/external/angle/include
            ${ANDROID_ROOT}/external/angle/src
            ${ANDROID_ROOT}/external/angle/src/common/base
            ${ANDROID_ROOT}/external/angle/src/common/third_party/xxhash
            ${ANDROID_ROOT}/external/angle/third_party/abseil-cpp
    )

    # Add c/c++ flags
    target_compile_options(angle_gpu_info_util PRIVATE
            -Werror
            -DNDEBUG
            -Wno-unused-parameter
            -DNDEBUG
            -Wno-unused-parameter
            -DNDEBUG
            -Wno-unused-parameter
            -DNDEBUG
            -Wno-unused-parameter
            -DNDEBUG
            -Wno-unused-parameter
            -DNDEBUG
            -Wno-unused-parameter
    )

    target_link_libraries(angle_gpu_info_util
            angle_common
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            liblog
            libm
            libnativewindow
    )
endif ()
#################### angle_gpu_info_util ####################

#################### third_party_vulkan_deps_spirv_tools_src_spvtools_val ####################
set(FILES
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/val/basic_block.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/val/construct.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/val/function.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/val/instruction.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/val/validate.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/val/validate_adjacency.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/val/validate_annotation.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/val/validate_arithmetics.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/val/validate_atomics.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/val/validate_barriers.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/val/validate_bitwise.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/val/validate_builtins.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/val/validate_capability.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/val/validate_cfg.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/val/validate_composites.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/val/validate_constants.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/val/validate_conversion.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/val/validate_debug.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/val/validate_decorations.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/val/validate_derivatives.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/val/validate_execution_limitations.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/val/validate_extensions.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/val/validate_function.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/val/validate_id.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/val/validate_image.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/val/validate_instruction.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/val/validate_interfaces.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/val/validate_layout.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/val/validate_literals.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/val/validate_logicals.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/val/validate_memory.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/val/validate_memory_semantics.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/val/validate_mesh_shading.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/val/validate_misc.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/val/validate_mode_setting.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/val/validate_non_uniform.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/val/validate_primitives.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/val/validate_ray_query.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/val/validate_ray_tracing.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/val/validate_ray_tracing_reorder.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/val/validate_scopes.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/val/validate_small_type_uses.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/val/validate_type.cpp
        ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/source/val/validation_state.cpp
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
    add_library(third_party_vulkan_deps_spirv_tools_src_spvtools_val SHARED ${SOURCE_FILES})

    target_include_directories(third_party_vulkan_deps_spirv_tools_src_spvtools_val
            PUBLIC
            ${ANDROID_ROOT}/external/angle
            ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src/include
            ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-headers/src/include
            ${ANDROID_ROOT}/external/angle/third_party/vulkan-deps/spirv-tools/src
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_core_enums_unified1/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_core_tables_unified1/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_generators_inc/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_glsl_tables_glsl1-0/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_language_header_cldebuginfo100/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_language_header_debuginfo/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_language_header_vkdebuginfo100/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_opencl_tables_opencl1-0/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_vendor_tables_debuginfo/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_vendor_tables_nonsemantic.clspvreflection/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_vendor_tables_nonsemantic.shader.debuginfo.100/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_vendor_tables_opencl.debuginfo.100/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_vendor_tables_spv-amd-gcn-shader/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_vendor_tables_spv-amd-shader-ballot/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_vendor_tables_spv-amd-shader-explicit-vertex-parameter/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/angle/third_party_vulkan_deps_spirv_tools_src_spvtools_vendor_tables_spv-amd-shader-trinary-minmax/android_${OUT_ARCH_CPU}/gen
    )

    # Add c/c++ flags
    target_compile_options(third_party_vulkan_deps_spirv_tools_src_spvtools_val PRIVATE
            -Werror
            -DNDEBUG
            -Wno-unused-parameter
            -DNDEBUG
            -Wno-unused-parameter
    )

    target_link_libraries(third_party_vulkan_deps_spirv_tools_src_spvtools_val
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            third_party_vulkan_deps_spirv_tools_src_spvtools
    )
endif ()
#################### third_party_vulkan_deps_spirv_tools_src_spvtools_val ####################
