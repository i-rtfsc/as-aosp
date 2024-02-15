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

project(libneuralnetworks_common)

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

#################### libneuralnetworks_common ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/OperationResolver.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Activation.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/BatchMatmul.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/BidirectionalSequenceRNN.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Broadcast.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/ChannelShuffle.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Comparisons.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Concatenation.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Conv2D.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Densify.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/DepthwiseConv2D.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Dequantize.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Elementwise.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Elu.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Fill.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/FullyConnected.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Gather.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/GenerateProposals.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/HeatmapMaxKeypoint.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/InstanceNormalization.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/L2Normalization.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/LocalResponseNormalization.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/LogSoftmax.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/LogicalAndOr.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/LogicalNot.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/MirrorPad.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Neg.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/PRelu.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Pack.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Pooling.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/QLSTM.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Quantize.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Rank.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Reduce.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/ResizeImageOps.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Reverse.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/RoiAlign.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/RoiPooling.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Select.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Slice.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Softmax.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Squeeze.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/StridedSlice.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/TopK_V2.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Transpose.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/TransposeConv2D.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/UnidirectionalSequenceLSTM.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/UnidirectionalSequenceRNN.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/ActivationFunctor.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/BufferTracker.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/CpuExecutor.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/ExecutionBurstController.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/ExecutionBurstServer.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/GraphDump.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/HalBufferTracker.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/IndexedShapeWrapper.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/LegacyHalUtils.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/LegacyUtils.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/MemoryUtils.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/MetaModel.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/ModelUtils.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/OperationsExecutionUtils.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/QuantUtils.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/TokenHasher.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/ValidateHal.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/ArgMinMax.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/BidirectionalSequenceLSTM.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Cast.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/EmbeddingLookup.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/ExpandDims.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/GroupedConv2D.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/HashtableLookup.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/LSHProjection.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/LSTM.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/MaximumMinimum.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Multinomial.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Pow.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/QuantizedLSTM.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/RNN.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Reshape.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/SVDF.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/SimpleMath.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Split.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Tile.cpp
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
    add_library(libneuralnetworks_common SHARED ${SOURCE_FILES})

    target_include_directories(libneuralnetworks_common
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/include
            ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/include
            ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common
            ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/random
            ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/include
    )

    # Add c/c++ flags
    target_compile_options(libneuralnetworks_common PRIVATE
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(libneuralnetworks_common
            android.hardware.neuralnetworks-1.0
            android.hardware.neuralnetworks-1.1
            android.hardware.neuralnetworks-1.2
            android.hardware.neuralnetworks-1.3
            android.hidl.allocator-1.0
            android.hidl.memory-1.0
            gemmlowp_headers
            libarect
            libbase
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto_static
            libcutils
            libdl
            libeigen
            libfmq
            libhidlbase
            libhidlmemory
            liblog
            libm
            libnativewindow
            libneuralnetworks_headers
            libruy_static
            libtextclassifier_hash_headers
            libtextclassifier_hash_static
            libtflite_kernel_utils
            libutils
            neuralnetworks_types
            neuralnetworks_utils_hal_1_0
            neuralnetworks_utils_hal_1_1
            neuralnetworks_utils_hal_1_2
            neuralnetworks_utils_hal_1_3
            neuralnetworks_utils_hal_common
            philox_random
            philox_random_headers
            tensorflow_headers
    )
endif ()
#################### libneuralnetworks_common ####################

#################### libneuralnetworks_common_cl ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/OperationResolver.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Activation.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/BatchMatmul.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/BidirectionalSequenceRNN.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Broadcast.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/ChannelShuffle.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Comparisons.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Concatenation.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Conv2D.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Densify.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/DepthwiseConv2D.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Dequantize.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Elementwise.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Elu.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Fill.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/FullyConnected.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Gather.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/GenerateProposals.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/HeatmapMaxKeypoint.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/InstanceNormalization.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/L2Normalization.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/LocalResponseNormalization.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/LogSoftmax.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/LogicalAndOr.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/LogicalNot.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/MirrorPad.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Neg.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/PRelu.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Pack.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Pooling.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/QLSTM.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Quantize.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Rank.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Reduce.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/ResizeImageOps.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Reverse.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/RoiAlign.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/RoiPooling.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Select.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Slice.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Softmax.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Squeeze.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/StridedSlice.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/TopK_V2.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Transpose.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/TransposeConv2D.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/UnidirectionalSequenceLSTM.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/UnidirectionalSequenceRNN.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/BufferTracker.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/CpuExecutor.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/GraphDump.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/IndexedShapeWrapper.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/LegacyUtils.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/MetaModel.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/ModelUtils.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/OperationsExecutionUtils.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/TokenHasher.cpp
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
    add_library(libneuralnetworks_common_cl SHARED ${SOURCE_FILES})

    target_include_directories(libneuralnetworks_common_cl
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/include
            ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/include
            ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common
            ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/include
    )

    # Add c/c++ flags
    target_compile_options(libneuralnetworks_common_cl PRIVATE
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(libneuralnetworks_common_cl
            libbase_ndk
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto_static
            libdl
            libm
            libnativewindow
            libneuralnetworks_headers_ndk
            neuralnetworks_types_cl
    )
endif ()
#################### libneuralnetworks_common_cl ####################

#################### libneuralnetworks_common_experimental ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/OperationResolver.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Activation.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/BatchMatmul.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/BidirectionalSequenceRNN.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Broadcast.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/ChannelShuffle.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Comparisons.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Concatenation.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Conv2D.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Densify.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/DepthwiseConv2D.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Dequantize.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Elementwise.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Elu.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Fill.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/FullyConnected.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Gather.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/GenerateProposals.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/HeatmapMaxKeypoint.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/InstanceNormalization.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/L2Normalization.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/LocalResponseNormalization.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/LogSoftmax.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/LogicalAndOr.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/LogicalNot.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/MirrorPad.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Neg.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/PRelu.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Pack.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Pooling.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/QLSTM.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Quantize.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Rank.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Reduce.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/ResizeImageOps.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Reverse.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/RoiAlign.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/RoiPooling.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Select.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Slice.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Softmax.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Squeeze.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/StridedSlice.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/TopK_V2.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Transpose.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/TransposeConv2D.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/UnidirectionalSequenceLSTM.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/UnidirectionalSequenceRNN.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/ActivationFunctor.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/BufferTracker.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/CpuExecutor.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/ExecutionBurstController.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/ExecutionBurstServer.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/GraphDump.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/HalBufferTracker.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/IndexedShapeWrapper.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/LegacyHalUtils.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/LegacyUtils.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/MemoryUtils.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/MetaModel.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/ModelUtils.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/OperationsExecutionUtils.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/QuantUtils.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/TokenHasher.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/ValidateHal.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/ArgMinMax.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/BidirectionalSequenceLSTM.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Cast.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/EmbeddingLookup.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/ExpandDims.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/GroupedConv2D.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/HashtableLookup.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/LSHProjection.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/LSTM.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/MaximumMinimum.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Multinomial.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Pow.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/QuantizedLSTM.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/RNN.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Reshape.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/SVDF.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/SimpleMath.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Split.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Tile.cpp
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
    add_library(libneuralnetworks_common_experimental SHARED ${SOURCE_FILES})

    target_include_directories(libneuralnetworks_common_experimental
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/include
            ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/include
            ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common
            ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/random
            ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/include
    )

    # Add c/c++ flags
    target_compile_options(libneuralnetworks_common_experimental PRIVATE
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(libneuralnetworks_common_experimental
            android.hardware.neuralnetworks-1.0
            android.hardware.neuralnetworks-1.1
            android.hardware.neuralnetworks-1.2
            android.hardware.neuralnetworks-1.3
            android.hidl.allocator-1.0
            android.hidl.memory-1.0
            gemmlowp_headers
            libarect
            libbase
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto_static
            libcutils
            libdl
            libeigen
            libfmq
            libhidlbase
            libhidlmemory
            liblog
            libm
            libnativewindow
            libneuralnetworks_headers
            libruy_static
            libtextclassifier_hash_headers
            libtextclassifier_hash_static
            libtflite_kernel_utils
            libutils
            neuralnetworks_types_experimental
            neuralnetworks_utils_hal_1_0
            neuralnetworks_utils_hal_1_1
            neuralnetworks_utils_hal_1_2
            neuralnetworks_utils_hal_1_3
            neuralnetworks_utils_hal_common
            philox_random
            philox_random_headers
            tensorflow_headers
    )
endif ()
#################### libneuralnetworks_common_experimental ####################

#################### libneuralnetworks_common_cl_cpu ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/OperationResolver.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Activation.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/BatchMatmul.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/BidirectionalSequenceRNN.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Broadcast.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/ChannelShuffle.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Comparisons.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Concatenation.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Conv2D.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Densify.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/DepthwiseConv2D.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Dequantize.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Elementwise.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Elu.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Fill.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/FullyConnected.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Gather.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/GenerateProposals.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/HeatmapMaxKeypoint.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/InstanceNormalization.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/L2Normalization.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/LocalResponseNormalization.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/LogSoftmax.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/LogicalAndOr.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/LogicalNot.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/MirrorPad.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Neg.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/PRelu.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Pack.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Pooling.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/QLSTM.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Quantize.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Rank.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Reduce.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/ResizeImageOps.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Reverse.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/RoiAlign.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/RoiPooling.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Select.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Slice.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Softmax.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Squeeze.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/StridedSlice.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/TopK_V2.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Transpose.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/TransposeConv2D.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/UnidirectionalSequenceLSTM.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/UnidirectionalSequenceRNN.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/BufferTracker.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/CpuExecutor.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/GraphDump.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/IndexedShapeWrapper.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/LegacyUtils.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/MetaModel.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/ModelUtils.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/OperationsExecutionUtils.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/TokenHasher.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/ActivationFunctor.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/QuantUtils.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/ArgMinMax.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/BidirectionalSequenceLSTM.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Cast.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/EmbeddingLookup.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/ExpandDims.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/GroupedConv2D.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/HashtableLookup.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/LSHProjection.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/LSTM.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/MaximumMinimum.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Multinomial.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Pow.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/QuantizedLSTM.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/RNN.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Reshape.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/SVDF.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/SimpleMath.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Split.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/Tile.cpp
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
    add_library(libneuralnetworks_common_cl_cpu SHARED ${SOURCE_FILES})

    target_include_directories(libneuralnetworks_common_cl_cpu
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/include
            ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/include
            ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common
            ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/random
            ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/include
    )

    # Add c/c++ flags
    target_compile_options(libneuralnetworks_common_cl_cpu PRIVATE
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(libneuralnetworks_common_cl_cpu
            gemmlowp_headers
            libbase_ndk
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto_static
            libdl
            libeigen
            libm
            libnativewindow
            libneuralnetworks_headers_ndk
            libruy_headers
            libtextclassifier_hash_headers
            libtextclassifier_hash_static
            neuralnetworks_types_cl
            philox_random_cl
            philox_random_headers
            tensorflow_headers
    )
endif ()
#################### libneuralnetworks_common_cl_cpu ####################
