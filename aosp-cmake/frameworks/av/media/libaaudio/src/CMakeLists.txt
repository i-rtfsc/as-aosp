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

project(libaaudio)

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

#################### libaaudio ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/core/AAudioAudio.cpp
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
    add_library(libaaudio SHARED ${SOURCE_FILES})

    target_include_directories(libaaudio
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/binding
            ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/client
            ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/core
            ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/fifo
            ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/flowgraph
            ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/legacy
            ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/utility
            ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaaudio/src/aaudio-aidl-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(libaaudio PRIVATE
            -Wno-unused-parameter
            -Wall
            -Werror
    )

    target_link_libraries(libaaudio
            framework-permission-aidl-cpp
            libaaudio_headers
            libaaudio_internal
            libaudioclient
            libaudioutils
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            liblog
            libm
            libmedia_helper
            libmediametrics
            libmediautils
            libutils
    )
endif ()
#################### libaaudio ####################

#################### aaudio-aidl-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaaudio/src/aaudio-aidl-ndk-source/gen/aaudio/Endpoint.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaaudio/src/aaudio-aidl-ndk-source/gen/aaudio/RingBuffer.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaaudio/src/aaudio-aidl-ndk-source/gen/aaudio/SharedRegion.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaaudio/src/aaudio-aidl-ndk-source/gen/aaudio/StreamParameters.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaaudio/src/aaudio-aidl-ndk-source/gen/aaudio/StreamRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaaudio/src/aaudio-aidl-ndk-source/gen/aaudio/IAAudioClient.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaaudio/src/aaudio-aidl-ndk-source/gen/aaudio/IAAudioService.cpp
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
    add_library(aaudio-aidl-ndk SHARED ${SOURCE_FILES})

    target_include_directories(aaudio-aidl-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaaudio/src/aaudio-aidl-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(aaudio-aidl-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(aaudio-aidl-ndk
            android.media.audio.common.types-V2-ndk
            audioclient-types-aidl-ndk
            framework-permission-aidl-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            shared-file-region-aidl-ndk
    )
endif ()
#################### aaudio-aidl-ndk ####################

#################### aaudio-aidl-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaaudio/src/aaudio-aidl-ndk_platform-source/gen/aaudio/Endpoint.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaaudio/src/aaudio-aidl-ndk_platform-source/gen/aaudio/RingBuffer.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaaudio/src/aaudio-aidl-ndk_platform-source/gen/aaudio/SharedRegion.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaaudio/src/aaudio-aidl-ndk_platform-source/gen/aaudio/StreamParameters.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaaudio/src/aaudio-aidl-ndk_platform-source/gen/aaudio/StreamRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaaudio/src/aaudio-aidl-ndk_platform-source/gen/aaudio/IAAudioClient.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaaudio/src/aaudio-aidl-ndk_platform-source/gen/aaudio/IAAudioService.cpp
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
    add_library(aaudio-aidl-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(aaudio-aidl-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaaudio/src/aaudio-aidl-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(aaudio-aidl-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(aaudio-aidl-ndk_platform
            android.media.audio.common.types-V2-ndk_platform
            audioclient-types-aidl-ndk_platform
            framework-permission-aidl-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            shared-file-region-aidl-ndk_platform
    )
endif ()
#################### aaudio-aidl-ndk_platform ####################

#################### aaudio-aidl-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaaudio/src/aaudio-aidl-cpp-source/gen/aaudio/Endpoint.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaaudio/src/aaudio-aidl-cpp-source/gen/aaudio/RingBuffer.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaaudio/src/aaudio-aidl-cpp-source/gen/aaudio/SharedRegion.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaaudio/src/aaudio-aidl-cpp-source/gen/aaudio/StreamParameters.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaaudio/src/aaudio-aidl-cpp-source/gen/aaudio/StreamRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaaudio/src/aaudio-aidl-cpp-source/gen/aaudio/IAAudioClient.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaaudio/src/aaudio-aidl-cpp-source/gen/aaudio/IAAudioService.cpp
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
    add_library(aaudio-aidl-cpp SHARED ${SOURCE_FILES})

    target_include_directories(aaudio-aidl-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaaudio/src/aaudio-aidl-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(aaudio-aidl-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(aaudio-aidl-cpp
            android.media.audio.common.types-V2-cpp
            audioclient-types-aidl-cpp
            framework-permission-aidl-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libm
            libutils
            shared-file-region-aidl-cpp
    )
endif ()
#################### aaudio-aidl-cpp ####################

#################### libaaudio_internal ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/core/AudioGlobal.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/core/AudioStream.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/core/AudioStreamBuilder.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/core/AAudioStreamParameters.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/legacy/AudioStreamLegacy.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/legacy/AudioStreamRecord.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/legacy/AudioStreamTrack.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/utility/AAudioUtilities.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/utility/FixedBlockAdapter.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/utility/FixedBlockReader.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/utility/FixedBlockWriter.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/fifo/FifoBuffer.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/fifo/FifoControllerBase.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/client/AAudioFlowGraph.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/client/AudioEndpoint.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/client/AudioStreamInternal.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/client/AudioStreamInternalCapture.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/client/AudioStreamInternalPlay.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/client/IsochronousClockModel.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/binding/AudioEndpointParcelable.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/binding/AAudioBinderAdapter.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/binding/AAudioBinderClient.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/binding/AAudioStreamRequest.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/binding/AAudioStreamConfiguration.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/binding/RingBufferParcelable.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/binding/SharedMemoryParcelable.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/binding/SharedRegionParcelable.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/flowgraph/ChannelCountConverter.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/flowgraph/ClipToRange.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/flowgraph/FlowGraphNode.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/flowgraph/Limiter.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/flowgraph/ManyToMultiConverter.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/flowgraph/MonoBlend.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/flowgraph/MonoToMultiConverter.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/flowgraph/MultiToMonoConverter.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/flowgraph/MultiToManyConverter.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/flowgraph/RampLinear.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/flowgraph/SampleRateConverter.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/flowgraph/SinkFloat.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/flowgraph/SinkI16.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/flowgraph/SinkI24.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/flowgraph/SinkI32.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/flowgraph/SourceFloat.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/flowgraph/SourceI16.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/flowgraph/SourceI24.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/flowgraph/SourceI32.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/flowgraph/resampler/IntegerRatio.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/flowgraph/resampler/LinearResampler.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/flowgraph/resampler/MultiChannelResampler.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/flowgraph/resampler/PolyphaseResampler.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/flowgraph/resampler/PolyphaseResamplerMono.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/flowgraph/resampler/PolyphaseResamplerStereo.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/flowgraph/resampler/SincResampler.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/flowgraph/resampler/SincResamplerStereo.cpp
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
    add_library(libaaudio_internal SHARED ${SOURCE_FILES})

    target_include_directories(libaaudio_internal
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src
            ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/binding
            ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/client
            ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/core
            ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/fifo
            ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/legacy
            ${ANDROID_ROOT}/frameworks/av/media/libaaudio/src/utility
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaaudio/src/aaudio-aidl-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(libaaudio_internal PRIVATE
            -Wno-unused-parameter
            -Wall
            -Werror
    )

    target_link_libraries(libaaudio_internal
            aaudio-aidl-cpp
            android.media.audio.common.types-V2-cpp
            audioclient-types-aidl-cpp
            framework-permission-aidl-cpp
            libaaudio_headers
            libaudioclient
            libaudioclient_aidl_conversion
            libaudioutils
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            liblog
            libm
            libmedia_headers
            libmedia_helper
            libmediametrics
            libmediametrics_headers
            libmediautils
            libutils
    )
endif ()
#################### libaaudio_internal ####################
