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

project(oboe)

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

#################### oboe ####################
set(FILES
        ${ANDROID_ROOT}/external/oboe/src/aaudio/AAudioLoader.cpp
        ${ANDROID_ROOT}/external/oboe/src/aaudio/AudioStreamAAudio.cpp
        ${ANDROID_ROOT}/external/oboe/src/common/AudioSourceCaller.cpp
        ${ANDROID_ROOT}/external/oboe/src/common/AudioStream.cpp
        ${ANDROID_ROOT}/external/oboe/src/common/AudioStreamBuilder.cpp
        ${ANDROID_ROOT}/external/oboe/src/common/DataConversionFlowGraph.cpp
        ${ANDROID_ROOT}/external/oboe/src/common/FilterAudioStream.cpp
        ${ANDROID_ROOT}/external/oboe/src/common/FixedBlockAdapter.cpp
        ${ANDROID_ROOT}/external/oboe/src/common/FixedBlockReader.cpp
        ${ANDROID_ROOT}/external/oboe/src/common/FixedBlockWriter.cpp
        ${ANDROID_ROOT}/external/oboe/src/common/LatencyTuner.cpp
        ${ANDROID_ROOT}/external/oboe/src/common/OboeExtensions.cpp
        ${ANDROID_ROOT}/external/oboe/src/common/SourceFloatCaller.cpp
        ${ANDROID_ROOT}/external/oboe/src/common/SourceI16Caller.cpp
        ${ANDROID_ROOT}/external/oboe/src/common/SourceI24Caller.cpp
        ${ANDROID_ROOT}/external/oboe/src/common/SourceI32Caller.cpp
        ${ANDROID_ROOT}/external/oboe/src/common/Utilities.cpp
        ${ANDROID_ROOT}/external/oboe/src/common/QuirksManager.cpp
        ${ANDROID_ROOT}/external/oboe/src/fifo/FifoBuffer.cpp
        ${ANDROID_ROOT}/external/oboe/src/fifo/FifoController.cpp
        ${ANDROID_ROOT}/external/oboe/src/fifo/FifoControllerBase.cpp
        ${ANDROID_ROOT}/external/oboe/src/fifo/FifoControllerIndirect.cpp
        ${ANDROID_ROOT}/external/oboe/src/flowgraph/FlowGraphNode.cpp
        ${ANDROID_ROOT}/external/oboe/src/flowgraph/ChannelCountConverter.cpp
        ${ANDROID_ROOT}/external/oboe/src/flowgraph/ClipToRange.cpp
        ${ANDROID_ROOT}/external/oboe/src/flowgraph/Limiter.cpp
        ${ANDROID_ROOT}/external/oboe/src/flowgraph/ManyToMultiConverter.cpp
        ${ANDROID_ROOT}/external/oboe/src/flowgraph/MonoBlend.cpp
        ${ANDROID_ROOT}/external/oboe/src/flowgraph/MonoToMultiConverter.cpp
        ${ANDROID_ROOT}/external/oboe/src/flowgraph/MultiToManyConverter.cpp
        ${ANDROID_ROOT}/external/oboe/src/flowgraph/MultiToMonoConverter.cpp
        ${ANDROID_ROOT}/external/oboe/src/flowgraph/RampLinear.cpp
        ${ANDROID_ROOT}/external/oboe/src/flowgraph/SampleRateConverter.cpp
        ${ANDROID_ROOT}/external/oboe/src/flowgraph/SinkFloat.cpp
        ${ANDROID_ROOT}/external/oboe/src/flowgraph/SinkI16.cpp
        ${ANDROID_ROOT}/external/oboe/src/flowgraph/SinkI24.cpp
        ${ANDROID_ROOT}/external/oboe/src/flowgraph/SinkI32.cpp
        ${ANDROID_ROOT}/external/oboe/src/flowgraph/SourceFloat.cpp
        ${ANDROID_ROOT}/external/oboe/src/flowgraph/SourceI16.cpp
        ${ANDROID_ROOT}/external/oboe/src/flowgraph/SourceI24.cpp
        ${ANDROID_ROOT}/external/oboe/src/flowgraph/SourceI32.cpp
        ${ANDROID_ROOT}/external/oboe/src/flowgraph/resampler/IntegerRatio.cpp
        ${ANDROID_ROOT}/external/oboe/src/flowgraph/resampler/LinearResampler.cpp
        ${ANDROID_ROOT}/external/oboe/src/flowgraph/resampler/MultiChannelResampler.cpp
        ${ANDROID_ROOT}/external/oboe/src/flowgraph/resampler/PolyphaseResampler.cpp
        ${ANDROID_ROOT}/external/oboe/src/flowgraph/resampler/PolyphaseResamplerMono.cpp
        ${ANDROID_ROOT}/external/oboe/src/flowgraph/resampler/PolyphaseResamplerStereo.cpp
        ${ANDROID_ROOT}/external/oboe/src/flowgraph/resampler/SincResampler.cpp
        ${ANDROID_ROOT}/external/oboe/src/flowgraph/resampler/SincResamplerStereo.cpp
        ${ANDROID_ROOT}/external/oboe/src/opensles/AudioInputStreamOpenSLES.cpp
        ${ANDROID_ROOT}/external/oboe/src/opensles/AudioOutputStreamOpenSLES.cpp
        ${ANDROID_ROOT}/external/oboe/src/opensles/AudioStreamBuffered.cpp
        ${ANDROID_ROOT}/external/oboe/src/opensles/AudioStreamOpenSLES.cpp
        ${ANDROID_ROOT}/external/oboe/src/opensles/EngineOpenSLES.cpp
        ${ANDROID_ROOT}/external/oboe/src/opensles/OpenSLESUtilities.cpp
        ${ANDROID_ROOT}/external/oboe/src/opensles/OutputMixerOpenSLES.cpp
        ${ANDROID_ROOT}/external/oboe/src/common/StabilizedCallback.cpp
        ${ANDROID_ROOT}/external/oboe/src/common/Trace.cpp
        ${ANDROID_ROOT}/external/oboe/src/common/Version.cpp
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
    add_library(oboe SHARED ${SOURCE_FILES})

    target_include_directories(oboe
            PUBLIC
            ${ANDROID_ROOT}/external/oboe/include
            ${ANDROID_ROOT}/external/oboe/src
            ${ANDROID_ROOT}/external/oboe
    )

    # Add c/c++ flags
    target_compile_options(oboe PRIVATE
            -Werror
            -Wall
            -Wextra-semi
            -Wno-unused-parameter
    )

    target_link_libraries(oboe
            jni_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### oboe ####################
