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

project(neuralnetworks_types)

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

#################### neuralnetworks_types ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Activation.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/ArgMinMax.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/BatchMatmul.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/BidirectionalSequenceLSTM.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/BidirectionalSequenceRNN.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Broadcast.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Cast.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/ChannelShuffle.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Comparisons.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Concatenation.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Conv2D.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Densify.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/DepthwiseConv2D.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Dequantize.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Elementwise.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Elu.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/EmbeddingLookup.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/ExpandDims.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Fill.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/FullyConnected.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Gather.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/GenerateProposals.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/GroupedConv2D.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/HashtableLookup.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/HeatmapMaxKeypoint.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/InstanceNormalization.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/L2Normalization.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/LSHProjection.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/LSTM.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/LocalResponseNormalization.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/LogSoftmax.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/LogicalAndOr.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/LogicalNot.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/MaximumMinimum.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/MirrorPad.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Multinomial.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Neg.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/PRelu.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Pack.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Pooling.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Pow.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/QLSTM.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Quantize.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/QuantizedLSTM.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/RNN.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Rank.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Reduce.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Reshape.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/ResizeImageOps.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Reverse.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/RoiAlign.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/RoiPooling.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/SVDF.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Select.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/SimpleMath.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Slice.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Softmax.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Split.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Squeeze.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/StridedSlice.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Tile.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/TopK_V2.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Transpose.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/TransposeConv2D.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/UnidirectionalSequenceLSTM.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/UnidirectionalSequenceRNN.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/src/OperationsUtils.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/src/OperationsValidationUtils.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/src/SharedMemory.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/src/SharedMemoryAndroid.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/src/TypeUtils.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/src/Types.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/src/Validation.cpp
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
    add_library(neuralnetworks_types SHARED ${SOURCE_FILES})

    target_include_directories(neuralnetworks_types
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/include
            ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/include/nnapi
            ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/include
            ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types
    )

    # Add c/c++ flags
    target_compile_options(neuralnetworks_types PRIVATE
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(neuralnetworks_types
            libarect
            libbase
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libm
            libnativewindow
            libutils
    )
endif ()
#################### neuralnetworks_types ####################

#################### neuralnetworks_types_cl ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Activation.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/ArgMinMax.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/BatchMatmul.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/BidirectionalSequenceLSTM.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/BidirectionalSequenceRNN.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Broadcast.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Cast.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/ChannelShuffle.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Comparisons.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Concatenation.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Conv2D.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Densify.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/DepthwiseConv2D.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Dequantize.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Elementwise.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Elu.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/EmbeddingLookup.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/ExpandDims.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Fill.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/FullyConnected.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Gather.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/GenerateProposals.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/GroupedConv2D.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/HashtableLookup.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/HeatmapMaxKeypoint.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/InstanceNormalization.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/L2Normalization.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/LSHProjection.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/LSTM.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/LocalResponseNormalization.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/LogSoftmax.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/LogicalAndOr.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/LogicalNot.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/MaximumMinimum.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/MirrorPad.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Multinomial.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Neg.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/PRelu.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Pack.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Pooling.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Pow.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/QLSTM.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Quantize.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/QuantizedLSTM.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/RNN.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Rank.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Reduce.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Reshape.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/ResizeImageOps.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Reverse.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/RoiAlign.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/RoiPooling.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/SVDF.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Select.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/SimpleMath.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Slice.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Softmax.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Split.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Squeeze.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/StridedSlice.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Tile.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/TopK_V2.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Transpose.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/TransposeConv2D.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/UnidirectionalSequenceLSTM.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/UnidirectionalSequenceRNN.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/src/DynamicCLDeps.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/src/OperationsUtils.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/src/OperationsValidationUtils.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/src/SharedMemory.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/src/SharedMemoryAndroid.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/src/TypeUtils.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/src/Types.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/src/Validation.cpp
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
    add_library(neuralnetworks_types_cl SHARED ${SOURCE_FILES})

    target_include_directories(neuralnetworks_types_cl
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/include
            ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/include/nnapi
            ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/include
            ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types
    )

    # Add c/c++ flags
    target_compile_options(neuralnetworks_types_cl PRIVATE
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(neuralnetworks_types_cl
            libandroid
            libbase_ndk
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### neuralnetworks_types_cl ####################

#################### neuralnetworks_types_experimental ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Activation.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/ArgMinMax.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/BatchMatmul.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/BidirectionalSequenceLSTM.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/BidirectionalSequenceRNN.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Broadcast.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Cast.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/ChannelShuffle.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Comparisons.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Concatenation.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Conv2D.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Densify.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/DepthwiseConv2D.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Dequantize.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Elementwise.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Elu.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/EmbeddingLookup.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/ExpandDims.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Fill.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/FullyConnected.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Gather.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/GenerateProposals.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/GroupedConv2D.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/HashtableLookup.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/HeatmapMaxKeypoint.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/InstanceNormalization.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/L2Normalization.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/LSHProjection.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/LSTM.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/LocalResponseNormalization.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/LogSoftmax.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/LogicalAndOr.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/LogicalNot.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/MaximumMinimum.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/MirrorPad.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Multinomial.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Neg.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/PRelu.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Pack.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Pooling.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Pow.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/QLSTM.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Quantize.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/QuantizedLSTM.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/RNN.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Rank.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Reduce.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Reshape.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/ResizeImageOps.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Reverse.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/RoiAlign.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/RoiPooling.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/SVDF.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Select.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/SimpleMath.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Slice.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Softmax.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Split.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Squeeze.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/StridedSlice.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Tile.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/TopK_V2.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/Transpose.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/TransposeConv2D.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/UnidirectionalSequenceLSTM.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/UnidirectionalSequenceRNN.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/src/OperationsUtils.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/src/OperationsValidationUtils.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/src/SharedMemory.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/src/SharedMemoryAndroid.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/src/TypeUtils.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/src/Types.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/src/Validation.cpp
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
    add_library(neuralnetworks_types_experimental SHARED ${SOURCE_FILES})

    target_include_directories(neuralnetworks_types_experimental
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/include
            ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/include/nnapi
            ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/include
            ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types
    )

    # Add c/c++ flags
    target_compile_options(neuralnetworks_types_experimental PRIVATE
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(neuralnetworks_types_experimental
            libarect
            libbase
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libm
            libnativewindow
            libutils
    )
endif ()
#################### neuralnetworks_types_experimental ####################
