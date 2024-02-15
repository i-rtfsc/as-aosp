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

project(libqtiguiextension)

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

#################### libqtiguiextension ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/libs/gui/QtiExtension/QtiBLASTBufferQueueExtension.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/QtiExtension/QtiBufferQueueProducerExtension.cpp
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
    add_library(libqtiguiextension SHARED ${SOURCE_FILES})

    target_include_directories(libqtiguiextension
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/libs/gui/include
            ${ANDROID_ROOT}/frameworks/native/libs/gui/QtiExtension
            ${ANDROID_ROOT}/frameworks/native/libs/gui
            ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_aidl_static/android_${OUT_ARCH_CPU}_static/gen/aidl
            ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/sysprop/libLibGuiProperties/android_${OUT_ARCH_CPU}_static/gen/sysprop/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_window_info_static/android_${OUT_ARCH_CPU}_static/gen/aidl
    )

    # Add c/c++ flags
    target_compile_options(libqtiguiextension PRIVATE
            -Wall
            -Werror
            -DQTI_DISPLAY_EXTENSION
    )

    target_link_libraries(libqtiguiextension
            android.hardware.graphics.bufferqueue-1.0
            android.hardware.graphics.bufferqueue-2.0
            android.hardware.graphics.common-V4-ndk
            android.hardware.graphics.common-1.1
            android.hardware.graphics.common-1.2
            android.hidl.token-1.0-utils
            libEGL
            libGLESv2
            libLibGuiProperties
            libbase
            libbinderthreadstateutils
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libgui_aidl_headers
            libgui_headers
            libgui_window_info_static
            libhidlbase
            liblog
            libm
            libnativebase_headers
            libnativewindow
            libsync
            libui
            libutils
            libvndksupport
    )
endif ()
#################### libqtiguiextension ####################

#################### libgui_bufferqueue_static ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_bufferqueue_static/android_${OUT_ARCH_CPU}_static/gen/aidl/android/gui/FrameEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_bufferqueue_static/android_${OUT_ARCH_CPU}_static/gen/aidl/android/os/IInputConstants.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_bufferqueue_static/android_${OUT_ARCH_CPU}_static/gen/aidl/android/os/InputEventInjectionResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_bufferqueue_static/android_${OUT_ARCH_CPU}_static/gen/aidl/android/os/InputEventInjectionSync.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_bufferqueue_static/android_${OUT_ARCH_CPU}_static/gen/aidl/android/os/InputConfig.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/BatchBufferOps.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/BufferItem.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/BufferQueue.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/BufferQueueConsumer.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/BufferQueueCore.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/BufferQueueProducer.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/BufferQueueThreadState.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/BufferSlot.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/FrameTimestamps.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/GLConsumerUtils.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/HdrMetadata.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/IGraphicBufferProducerFlattenables.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/bufferqueue/1.0/Conversion.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/bufferqueue/1.0/H2BProducerListener.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/bufferqueue/1.0/WProducerListener.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/bufferqueue/2.0/B2HGraphicBufferProducer.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/bufferqueue/2.0/H2BProducerListener.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/bufferqueue/2.0/types.cpp
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
    add_library(libgui_bufferqueue_static SHARED ${SOURCE_FILES})

    target_include_directories(libgui_bufferqueue_static
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/libs/gui/include
            ${ANDROID_ROOT}/frameworks/native/libs/gui
            ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_bufferqueue_static/android_${OUT_ARCH_CPU}_static/gen/aidl
            ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/sysprop/libLibGuiProperties/android_${OUT_ARCH_CPU}_static/gen/sysprop/include
    )

    # Add c/c++ flags
    target_compile_options(libgui_bufferqueue_static PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libgui_bufferqueue_static
            android.hardware.graphics.bufferqueue-1.0
            android.hardware.graphics.bufferqueue-2.0
            android.hardware.graphics.common-V4-ndk
            android.hardware.graphics.common-1.1
            android.hardware.graphics.common-1.2
            android.hidl.token-1.0-utils
            libEGL
            libGLESv2
            libLibGuiProperties
            libbase
            libbinderthreadstateutils
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libgui_headers
            libhidlbase
            liblog
            libm
            libnativebase_headers
            libnativewindow
            libsync
            libui
            libutils
            libvndksupport
    )
endif ()
#################### libgui_bufferqueue_static ####################

#################### libgui_window_info_static ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_window_info_static/android_${OUT_ARCH_CPU}_static/gen/aidl/android/gui/DropInputMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_window_info_static/android_${OUT_ARCH_CPU}_static/gen/aidl/android/gui/TouchOcclusionMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_window_info_static/android_${OUT_ARCH_CPU}_static/gen/aidl/android/os/IInputConstants.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_window_info_static/android_${OUT_ARCH_CPU}_static/gen/aidl/android/os/InputEventInjectionResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_window_info_static/android_${OUT_ARCH_CPU}_static/gen/aidl/android/os/InputEventInjectionSync.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_window_info_static/android_${OUT_ARCH_CPU}_static/gen/aidl/android/os/InputConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_window_info_static/android_${OUT_ARCH_CPU}_static/gen/aidl/android/gui/DisplayInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_window_info_static/android_${OUT_ARCH_CPU}_static/gen/aidl/android/gui/FocusRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_window_info_static/android_${OUT_ARCH_CPU}_static/gen/aidl/android/gui/InputApplicationInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_window_info_static/android_${OUT_ARCH_CPU}_static/gen/aidl/android/gui/IWindowInfosListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_window_info_static/android_${OUT_ARCH_CPU}_static/gen/aidl/android/gui/IWindowInfosPublisher.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_window_info_static/android_${OUT_ARCH_CPU}_static/gen/aidl/android/gui/IWindowInfosReportedListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_window_info_static/android_${OUT_ARCH_CPU}_static/gen/aidl/android/gui/WindowInfosUpdate.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_window_info_static/android_${OUT_ARCH_CPU}_static/gen/aidl/android/gui/WindowInfo.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/DisplayInfo.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/WindowInfo.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/WindowInfosUpdate.cpp
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
    add_library(libgui_window_info_static SHARED ${SOURCE_FILES})

    target_include_directories(libgui_window_info_static
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/libs/gui/include
            ${ANDROID_ROOT}/frameworks/native/libs/gui
            ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_window_info_static/android_${OUT_ARCH_CPU}_static/gen/aidl
    )

    # Add c/c++ flags
    target_compile_options(libgui_window_info_static PRIVATE
            -Werror
    )

    target_link_libraries(libgui_window_info_static
            libbinder
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
            libui-types
    )
endif ()
#################### libgui_window_info_static ####################

#################### libgui_mocks ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_mocks/android_${OUT_ARCH_CPU}_static/gen/aidl/android/gui/FrameEvent.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/mock/GraphicBufferConsumer.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/mock/GraphicBufferProducer.cpp
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
    add_library(libgui_mocks SHARED ${SOURCE_FILES})

    target_include_directories(libgui_mocks
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/libs/gui/include
            ${ANDROID_ROOT}/frameworks/native/libs/gui
            ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_mocks/android_${OUT_ARCH_CPU}_static/gen/aidl
            ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/sysprop/libLibGuiProperties/android_${OUT_ARCH_CPU}_static/gen/sysprop/include
    )

    # Add c/c++ flags
    target_compile_options(libgui_mocks PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libgui_mocks
            android.hardware.graphics.bufferqueue-1.0
            android.hardware.graphics.bufferqueue-2.0
            android.hardware.graphics.common-V4-ndk
            android.hardware.graphics.common-1.1
            android.hardware.graphics.common-1.2
            android.hidl.token-1.0-utils
            libEGL
            libGLESv2
            libLibGuiProperties
            libbase
            libbinderthreadstateutils
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libgmock
            libgtest
            libgui_headers
            libhidlbase
            liblog
            libm
            libnativebase_headers
            libnativewindow
            libsync
            libui
            libutils
            libvndksupport
    )
endif ()
#################### libgui_mocks ####################

#################### libgui ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/libs/gui/flyme/FlymeBufferQueueTracker.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui/android_${OUT_ARCH_CPU}_shared/gen/aidl/android/os/PersistableBundle.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui/android_${OUT_ARCH_CPU}_shared/gen/aidl/android/view/LayerMetadataKey.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui/android_${OUT_ARCH_CPU}_shared/gen/aidl/android/view/Surface.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui/android_${OUT_ARCH_CPU}_shared/gen/aidl/android/os/IInputConstants.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui/android_${OUT_ARCH_CPU}_shared/gen/aidl/android/os/InputEventInjectionResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui/android_${OUT_ARCH_CPU}_shared/gen/aidl/android/os/InputEventInjectionSync.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui/android_${OUT_ARCH_CPU}_shared/gen/aidl/android/os/InputConfig.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/BatchBufferOps.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/BufferItem.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/BufferQueue.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/BufferQueueConsumer.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/BufferQueueCore.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/BufferQueueProducer.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/BufferQueueThreadState.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/BufferSlot.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/FrameTimestamps.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/GLConsumerUtils.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/HdrMetadata.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/IGraphicBufferProducerFlattenables.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/bufferqueue/1.0/Conversion.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/bufferqueue/1.0/H2BProducerListener.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/bufferqueue/1.0/WProducerListener.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/bufferqueue/2.0/B2HGraphicBufferProducer.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/bufferqueue/2.0/H2BProducerListener.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/bufferqueue/2.0/types.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/BitTube.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/BLASTBufferQueue.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/BufferItemConsumer.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/Choreographer.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/CompositorTiming.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/ConsumerBase.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/CpuConsumer.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/DebugEGLImageTracker.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/DisplayEventDispatcher.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/DisplayEventReceiver.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/FenceMonitor.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/GLConsumer.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/IConsumerListener.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/IGraphicBufferConsumer.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/IGraphicBufferProducer.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/IProducerListener.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/ISurfaceComposer.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/ITransactionCompletedListener.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/LayerDebugInfo.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/LayerMetadata.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/LayerStatePermissions.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/LayerState.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/OccupancyTracker.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/StreamSplitter.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/ScreenCaptureResults.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/Surface.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/SurfaceControl.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/SurfaceComposerClient.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/SyncFeatures.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/VsyncEventData.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/view/Surface.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/WindowInfosListenerReporter.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/bufferqueue/1.0/B2HProducerListener.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/bufferqueue/1.0/H2BGraphicBufferProducer.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/bufferqueue/2.0/B2HProducerListener.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/gui/bufferqueue/2.0/H2BGraphicBufferProducer.cpp
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
    add_library(libgui SHARED ${SOURCE_FILES})

    target_include_directories(libgui
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/libs/gui/flyme
            ${ANDROID_ROOT}/frameworks/native/libs/gui/include
            ${ANDROID_ROOT}/frameworks/native/libs/gui
            ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui/android_${OUT_ARCH_CPU}_shared/gen/aidl
            ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_aidl_static/android_${OUT_ARCH_CPU}_static/gen/aidl
            ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/sysprop/libLibGuiProperties/android_${OUT_ARCH_CPU}_static_afdo-libgui_lto-thin/gen/sysprop/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_aidl_static/android_${OUT_ARCH_CPU}_static_afdo-libgui_lto-thin/gen/aidl
            ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_window_info_static/android_${OUT_ARCH_CPU}_static_afdo-libgui_lto-thin/gen/aidl
            ${ANDROID_ROOT}/frameworks/native/libs/gui/QtiExtension
    )

    # Add c/c++ flags
    target_compile_options(libgui PRIVATE
            -DFLYME_GRAPHICS_EXTEND
            -DFLYME_SUPPORT_LTPO_V2
            -DQTI_DISPLAY_EXTENSION
            -Wall
            -Werror
    )

    target_link_libraries(libgui
            android.hardware.graphics.bufferqueue-1.0
            android.hardware.graphics.bufferqueue-2.0
            android.hardware.graphics.common-V4-ndk
            android.hardware.graphics.common-1.1
            android.hardware.graphics.common-1.2
            android.hidl.token-1.0-utils
            jni_headers
            libEGL
            libGLESv2
            libLibGuiProperties
            libbase
            libbinder
            libbinderthreadstateutils
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libdvr_headers
            libgui_aidl_headers
            libgui_aidl_static
            libgui_headers
            libgui_window_info_static
            libhidlbase
            liblog
            libm
            libnativebase_headers
            libnativewindow
            libpdx_headers
            libqtiguiextension
            libsync
            libui
            libutils
            libvndksupport
    )
endif ()
#################### libgui ####################

#################### libgui_aidl_static ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_aidl_static/android_${OUT_ARCH_CPU}_static/gen/aidl/android/gui/ARect.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_aidl_static/android_${OUT_ARCH_CPU}_static/gen/aidl/android/gui/BitTube.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_aidl_static/android_${OUT_ARCH_CPU}_static/gen/aidl/android/gui/CachingHint.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_aidl_static/android_${OUT_ARCH_CPU}_static/gen/aidl/android/gui/Color.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_aidl_static/android_${OUT_ARCH_CPU}_static/gen/aidl/android/gui/CompositionPreference.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_aidl_static/android_${OUT_ARCH_CPU}_static/gen/aidl/android/gui/ContentSamplingAttributes.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_aidl_static/android_${OUT_ARCH_CPU}_static/gen/aidl/android/gui/CreateSurfaceResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_aidl_static/android_${OUT_ARCH_CPU}_static/gen/aidl/android/gui/DeviceProductInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_aidl_static/android_${OUT_ARCH_CPU}_static/gen/aidl/android/gui/DisplayBrightness.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_aidl_static/android_${OUT_ARCH_CPU}_static/gen/aidl/android/gui/DisplayCaptureArgs.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_aidl_static/android_${OUT_ARCH_CPU}_static/gen/aidl/android/gui/DisplayConnectionType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_aidl_static/android_${OUT_ARCH_CPU}_static/gen/aidl/android/gui/DisplayDecorationSupport.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_aidl_static/android_${OUT_ARCH_CPU}_static/gen/aidl/android/gui/DisplayMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_aidl_static/android_${OUT_ARCH_CPU}_static/gen/aidl/android/gui/DisplayModeSpecs.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_aidl_static/android_${OUT_ARCH_CPU}_static/gen/aidl/android/gui/DisplayPrimaries.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_aidl_static/android_${OUT_ARCH_CPU}_static/gen/aidl/android/gui/DisplayStatInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_aidl_static/android_${OUT_ARCH_CPU}_static/gen/aidl/android/gui/DisplayState.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_aidl_static/android_${OUT_ARCH_CPU}_static/gen/aidl/android/gui/DisplayedFrameStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_aidl_static/android_${OUT_ARCH_CPU}_static/gen/aidl/android/gui/DynamicDisplayInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_aidl_static/android_${OUT_ARCH_CPU}_static/gen/aidl/android/gui/FrameEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_aidl_static/android_${OUT_ARCH_CPU}_static/gen/aidl/android/gui/FrameStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_aidl_static/android_${OUT_ARCH_CPU}_static/gen/aidl/android/gui/FrameTimelineInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_aidl_static/android_${OUT_ARCH_CPU}_static/gen/aidl/android/gui/HdrCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_aidl_static/android_${OUT_ARCH_CPU}_static/gen/aidl/android/gui/HdrConversionCapability.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_aidl_static/android_${OUT_ARCH_CPU}_static/gen/aidl/android/gui/HdrConversionStrategy.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_aidl_static/android_${OUT_ARCH_CPU}_static/gen/aidl/android/gui/IDisplayEventConnection.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_aidl_static/android_${OUT_ARCH_CPU}_static/gen/aidl/android/gui/IFpsListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_aidl_static/android_${OUT_ARCH_CPU}_static/gen/aidl/android/gui/IHdrConversionConstants.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_aidl_static/android_${OUT_ARCH_CPU}_static/gen/aidl/android/gui/IHdrLayerInfoListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_aidl_static/android_${OUT_ARCH_CPU}_static/gen/aidl/android/gui/IRegionSamplingListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_aidl_static/android_${OUT_ARCH_CPU}_static/gen/aidl/android/gui/IScreenCaptureListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_aidl_static/android_${OUT_ARCH_CPU}_static/gen/aidl/android/gui/ISurfaceComposer.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_aidl_static/android_${OUT_ARCH_CPU}_static/gen/aidl/android/gui/ISurfaceComposerClient.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_aidl_static/android_${OUT_ARCH_CPU}_static/gen/aidl/android/gui/ITunnelModeEnabledListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_aidl_static/android_${OUT_ARCH_CPU}_static/gen/aidl/android/gui/LayerCaptureArgs.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_aidl_static/android_${OUT_ARCH_CPU}_static/gen/aidl/android/gui/LayerDebugInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_aidl_static/android_${OUT_ARCH_CPU}_static/gen/aidl/android/gui/LayerMetadata.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_aidl_static/android_${OUT_ARCH_CPU}_static/gen/aidl/android/gui/OverlayProperties.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_aidl_static/android_${OUT_ARCH_CPU}_static/gen/aidl/android/gui/ParcelableVsyncEventData.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_aidl_static/android_${OUT_ARCH_CPU}_static/gen/aidl/android/gui/PullAtomData.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_aidl_static/android_${OUT_ARCH_CPU}_static/gen/aidl/android/gui/Rotation.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_aidl_static/android_${OUT_ARCH_CPU}_static/gen/aidl/android/gui/ScreenCaptureResults.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_aidl_static/android_${OUT_ARCH_CPU}_static/gen/aidl/android/gui/Size.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_aidl_static/android_${OUT_ARCH_CPU}_static/gen/aidl/android/gui/StaticDisplayInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_aidl_static/android_${OUT_ARCH_CPU}_static/gen/aidl/android/gui/TrustedPresentationThresholds.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_aidl_static/android_${OUT_ARCH_CPU}_static/gen/aidl/android/gui/WindowInfosListenerInfo.cpp
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
    add_library(libgui_aidl_static SHARED ${SOURCE_FILES})

    target_include_directories(libgui_aidl_static
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/libs/gui/include
            ${ANDROID_ROOT}/frameworks/native/libs/gui
            ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_aidl_static/android_${OUT_ARCH_CPU}_static/gen/aidl
            ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/gui/libgui_window_info_static/android_${OUT_ARCH_CPU}_static/gen/aidl
    )

    # Add c/c++ flags
    target_compile_options(libgui_aidl_static PRIVATE
            -Werror
    )

    target_link_libraries(libgui_aidl_static
            libbinder
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libgui_window_info_static
            libm
            libui
            libui-types
    )
endif ()
#################### libgui_aidl_static ####################
