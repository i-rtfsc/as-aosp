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

project(libhwui_static)

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

#################### libhwui_static ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/canvas/CanvasFrontend.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/canvas/CanvasOpBuffer.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/canvas/CanvasOpRasterizer.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/effects/StretchEffect.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/effects/GainmapRenderer.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/pipeline/skia/HolePunch.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/pipeline/skia/SkiaDisplayList.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/pipeline/skia/SkiaRecordingCanvas.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/pipeline/skia/StretchMask.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/pipeline/skia/RenderNodeDrawable.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/pipeline/skia/ReorderBarrierDrawables.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/pipeline/skia/TransformCanvas.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/renderthread/Frame.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/renderthread/RenderTask.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/renderthread/TimeLord.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/hwui/AnimatedImageDrawable.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/hwui/Bitmap.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/hwui/BlurDrawLooper.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/hwui/Canvas.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/hwui/ImageDecoder.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/hwui/MinikinSkia.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/hwui/MinikinUtils.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/hwui/PaintImpl.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/hwui/Typeface.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/utils/Blur.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/utils/Color.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/utils/LinearAllocator.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/utils/VectorDrawableUtils.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/AnimationContext.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/Animator.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/AnimatorManager.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/CanvasTransform.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/DamageAccumulator.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/Gainmap.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/Interpolator.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/LightingInfo.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/Matrix.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/Mesh.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/MemoryPolicy.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/PathParser.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/Properties.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/PropertyValuesAnimatorSet.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/PropertyValuesHolder.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/RecordingCanvas.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/RenderNode.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/RenderProperties.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/RootRenderNode.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/SkiaCanvas.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/SkiaInterpolator.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/Tonemapper.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/VectorDrawable.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/flyme/BitmapAllocationTracker.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/flyme/BitmapDumper.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/flyme/GraphicBufferDumper.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/flyme/TextureDumperSkiaGL.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/flyme/Utils.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/flyme/flyme_draw_functor.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/flyme/FlymeBlurImageFilter.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/hwui/AnimatedImageThread.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/pipeline/skia/ATraceMemoryDump.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/pipeline/skia/GLFunctorDrawable.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/pipeline/skia/LayerDrawable.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/pipeline/skia/ShaderCache.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/pipeline/skia/SkiaMemoryTracer.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/pipeline/skia/SkiaOpenGLPipeline.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/pipeline/skia/SkiaPipeline.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/pipeline/skia/SkiaProfileRenderer.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/pipeline/skia/SkiaVulkanPipeline.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/pipeline/skia/VkFunctorDrawable.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/pipeline/skia/VkInteropFunctorDrawable.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/renderstate/RenderState.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/renderthread/CacheManager.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/renderthread/CanvasContext.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/renderthread/DrawFrameTask.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/renderthread/EglManager.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/renderthread/ReliableSurface.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/renderthread/RenderEffectCapabilityQuery.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/renderthread/VulkanManager.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/renderthread/VulkanSurface.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/renderthread/RenderProxy.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/renderthread/RenderThread.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/renderthread/HintSessionWrapper.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/service/GraphicsStatsService.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/thread/CommonPool.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/utils/GLUtils.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/utils/NdkUtils.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/utils/StringUtils.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/AutoBackendTextureRelease.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/DeferredLayerUpdater.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/DeviceInfo.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/FrameInfo.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/FrameInfoVisualizer.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/HardwareBitmapUploader.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libs/hwui/libhwui_static/android_${OUT_ARCH_CPU}_static/gen/sysprop/HWUIProperties.sysprop.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/JankTracker.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/FrameMetricsReporter.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/Layer.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/LayerUpdateQueue.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/ProfileData.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/ProfileDataContainer.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/Readback.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/TreeInfo.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/WebViewFunctorManager.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libs/hwui/libhwui_static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/libs/hwui/protos/graphicsstats.pb.cc
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
    add_library(libhwui_static SHARED ${SOURCE_FILES})

    target_include_directories(libhwui_static
            PUBLIC
            ${ANDROID_ROOT}/frameworks/base/libs/hwui
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libs/hwui/libhwui_static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/libs/hwui
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libs/hwui/libhwui_static/android_${OUT_ARCH_CPU}_static/gen/proto
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libs/hwui/libhwui_static/android_${OUT_ARCH_CPU}_static/gen/sysprop/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libs/hwui/statslog_hwui.h/gen
    )

    # Add c/c++ flags
    target_compile_options(libhwui_static PRIVATE
            -Wall
            -Wno-unused-parameter
            -Werror
            -DFLYME_GRAPHICS_EXTEND
    )

    target_link_libraries(libhwui_static
            android.hardware.graphics.common-V4-ndk
            android.hardware.graphics.common-1.2
            android.hardware.graphics.composer3-V2-ndk
            libEGL
            libEGL_blobCache
            libGLESv1_CM
            libGLESv2
            libGLESv3
            libandroid_headers_private
            libandroidfw
            libarect
            libbase
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto
            libcutils
            libdl
            libdng_sdk
            libem
            libexpat
            libft2
            libgui
            libharfbuzz_ng
            libheif
            libjpeg
            liblog
            libm
            libminikin
            libnativedisplay
            libnativewindow
            libperfetto_client_experimental
            libpiex
            libpng
            libprotobuf-cpp-lite
            libprotoutil
            libsfntly
            libshaders
            libskia
            libstatslog_hwui
            libstatspull_lazy
            libstatssocket_lazy
            libsync
            libtiff
            libtonemap
            libtonemap_headers
            libui
            libui-types
            libutils
            libvulkan
            libwebp-decode
            libwebp-encode
            libwuffs_mirror_release_c
            libz
    )
endif ()
#################### libhwui_static ####################

#################### libhwui ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/android_graphics_animation_NativeInterpolatorFactory.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/android_graphics_animation_RenderNodeAnimator.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/android_graphics_Canvas.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/android_graphics_ColorSpace.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/android_graphics_drawable_AnimatedVectorDrawable.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/android_graphics_drawable_VectorDrawable.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/android_graphics_HardwareRendererObserver.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/android_graphics_Matrix.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/android_graphics_Picture.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/android_graphics_DisplayListCanvas.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/android_graphics_Mesh.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/android_graphics_RenderNode.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/android_nio_utils.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/android_util_PathParser.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/Bitmap.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/BufferUtils.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/HardwareBufferHelpers.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/BitmapFactory.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/ByteBufferStreamAdaptor.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/Camera.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/CanvasProperty.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/ColorFilter.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/CreateJavaOutputStreamAdaptor.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/FontFamily.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/FontUtils.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/Gainmap.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/Graphics.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/ImageDecoder.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/Interpolator.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/MeshSpecification.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/MaskFilter.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/NinePatch.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/NinePatchPeeker.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/Paint.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/PaintFilter.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/Path.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/PathEffect.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/PathIterator.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/PathMeasure.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/Picture.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/Region.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/ScopedParcel.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/Shader.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/RenderEffect.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/Typeface.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/Utils.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/YuvToJpegEncoder.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/fonts/Font.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/fonts/FontFamily.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/text/LineBreaker.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/text/MeasuredText.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/text/TextShaper.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/text/GraphemeBreak.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/apex/android_matrix.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/apex/android_paint.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/apex/android_region.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/apex/properties.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/canvas/CanvasFrontend.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/canvas/CanvasOpBuffer.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/canvas/CanvasOpRasterizer.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/effects/StretchEffect.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/effects/GainmapRenderer.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/pipeline/skia/HolePunch.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/pipeline/skia/SkiaDisplayList.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/pipeline/skia/SkiaRecordingCanvas.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/pipeline/skia/StretchMask.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/pipeline/skia/RenderNodeDrawable.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/pipeline/skia/ReorderBarrierDrawables.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/pipeline/skia/TransformCanvas.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/renderthread/Frame.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/renderthread/RenderTask.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/renderthread/TimeLord.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/hwui/AnimatedImageDrawable.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/hwui/Bitmap.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/hwui/BlurDrawLooper.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/hwui/Canvas.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/hwui/ImageDecoder.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/hwui/MinikinSkia.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/hwui/MinikinUtils.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/hwui/PaintImpl.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/hwui/Typeface.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/utils/Blur.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/utils/Color.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/utils/LinearAllocator.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/utils/VectorDrawableUtils.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/AnimationContext.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/Animator.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/AnimatorManager.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/CanvasTransform.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/DamageAccumulator.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/Gainmap.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/Interpolator.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/LightingInfo.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/Matrix.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/Mesh.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/MemoryPolicy.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/PathParser.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/Properties.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/PropertyValuesAnimatorSet.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/PropertyValuesHolder.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/RecordingCanvas.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/RenderNode.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/RenderProperties.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/RootRenderNode.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/SkiaCanvas.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/SkiaInterpolator.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/Tonemapper.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/VectorDrawable.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/AnimatedImageDrawable.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/android_graphics_TextureLayer.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/android_graphics_HardwareRenderer.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/android_graphics_HardwareBufferRenderer.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/BitmapRegionDecoder.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/GIFMovie.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/GraphicsStatsService.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/Movie.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/MovieImpl.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/pdf/PdfDocument.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/pdf/PdfEditor.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/pdf/PdfRenderer.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/pdf/PdfUtils.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/apex/android_bitmap.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/apex/android_canvas.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/apex/jni_runtime.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/flyme/BitmapAllocationTracker.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/flyme/BitmapDumper.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/flyme/GraphicBufferDumper.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/flyme/TextureDumperSkiaGL.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/flyme/Utils.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/flyme/flyme_draw_functor.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/flyme/FlymeBlurImageFilter.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/hwui/AnimatedImageThread.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/pipeline/skia/ATraceMemoryDump.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/pipeline/skia/GLFunctorDrawable.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/pipeline/skia/LayerDrawable.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/pipeline/skia/ShaderCache.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/pipeline/skia/SkiaMemoryTracer.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/pipeline/skia/SkiaOpenGLPipeline.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/pipeline/skia/SkiaPipeline.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/pipeline/skia/SkiaProfileRenderer.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/pipeline/skia/SkiaVulkanPipeline.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/pipeline/skia/VkFunctorDrawable.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/pipeline/skia/VkInteropFunctorDrawable.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/renderstate/RenderState.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/renderthread/CacheManager.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/renderthread/CanvasContext.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/renderthread/DrawFrameTask.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/renderthread/EglManager.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/renderthread/ReliableSurface.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/renderthread/RenderEffectCapabilityQuery.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/renderthread/VulkanManager.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/renderthread/VulkanSurface.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/renderthread/RenderProxy.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/renderthread/RenderThread.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/renderthread/HintSessionWrapper.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/service/GraphicsStatsService.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/thread/CommonPool.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/utils/GLUtils.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/utils/NdkUtils.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/utils/StringUtils.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/AutoBackendTextureRelease.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/DeferredLayerUpdater.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/DeviceInfo.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/FrameInfo.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/FrameInfoVisualizer.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/HardwareBitmapUploader.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libs/hwui/libhwui/android_${OUT_ARCH_CPU}_static/gen/sysprop/HWUIProperties.sysprop.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/JankTracker.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/FrameMetricsReporter.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/Layer.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/LayerUpdateQueue.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/ProfileData.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/ProfileDataContainer.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/Readback.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/TreeInfo.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/WebViewFunctorManager.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libs/hwui/libhwui/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/libs/hwui/protos/graphicsstats.pb.cc
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
    add_library(libhwui SHARED ${SOURCE_FILES})

    target_include_directories(libhwui
            PUBLIC
            ${ANDROID_ROOT}/frameworks/base/libs/hwui
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libs/hwui/libhwui/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/libs/hwui
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libs/hwui/libhwui/android_${OUT_ARCH_CPU}_static/gen/proto
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libs/hwui/libhwui/android_${OUT_ARCH_CPU}_static/gen/sysprop/include
            ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni
            ${ANDROID_ROOT}/frameworks/base/libs/hwui/apex/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libs/hwui/statslog_hwui.h/gen
    )

    # Add c/c++ flags
    target_compile_options(libhwui PRIVATE
            -Wno-unused-parameter
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wall
            -Werror
            -Wall
            -Wno-unused-parameter
            -Werror
            -DFLYME_GRAPHICS_EXTEND
    )

    target_link_libraries(libhwui
            android.hardware.graphics.common-V4-ndk
            android.hardware.graphics.common-1.2
            android.hardware.graphics.composer3-V2-ndk
            android_graphics_apex_headers
            android_graphics_jni_headers
            libEGL
            libEGL_blobCache
            libGLESv1_CM
            libGLESv2
            libGLESv3
            libandroid_headers_private
            libandroidfw
            libarect
            libbase
            libbinder
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto
            libcutils
            libdl
            libdng_sdk
            libem
            libexpat
            libft2
            libgif
            libgui
            libharfbuzz_ng
            libheif
            libimage_io
            libjpeg
            liblog
            libm
            libmediandk
            libminikin
            libnativedisplay
            libnativehelper_lazy
            libnativewindow
            libnativewindow_headers
            libpdfium
            libperfetto_client_experimental
            libpiex
            libpng
            libprotobuf-cpp-lite
            libprotoutil
            libsfntly
            libshaders
            libskia
            libstatslog_hwui
            libstatspull_lazy
            libstatssocket_lazy
            libsync
            libtiff
            libtonemap
            libtonemap_headers
            libui
            libui-types
            libultrahdr
            libutils
            libvulkan
            libwebp-decode
            libwebp-encode
            libwuffs_mirror_release_c
            libz
            libziparchive_for_incfs
    )
endif ()
#################### libhwui ####################

#################### libstatslog_hwui ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libs/hwui/statslog_hwui.cpp/gen/statslog_hwui.cpp
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
    add_library(libstatslog_hwui SHARED ${SOURCE_FILES})

    target_include_directories(libstatslog_hwui
            PUBLIC
            ${ANDROID_ROOT}/frameworks/base/libs/hwui
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libs/hwui/statslog_hwui.cpp/gen
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libs/hwui/statslog_hwui.h/gen
    )

    target_link_libraries(libstatslog_hwui
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libstatspull_lazy
            libstatssocket_lazy
    )
endif ()
#################### libstatslog_hwui ####################
