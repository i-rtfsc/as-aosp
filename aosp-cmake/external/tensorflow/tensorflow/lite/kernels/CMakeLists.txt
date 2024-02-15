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

project(libtflite_kernel_utils)

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

#################### libtflite_kernel_utils ####################
set(FILES
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/internal/optimized/neon_tensor_utils.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/internal/optimized/sse_tensor_utils.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/internal/quantization_util.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/internal/reference/portable_tensor_utils.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/internal/tensor_utils.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/internal/transpose_utils.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/kernel_util.cc
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
    add_library(libtflite_kernel_utils SHARED ${SOURCE_FILES})

    target_include_directories(libtflite_kernel_utils
            PUBLIC
            ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels
    )

    # Add c/c++ flags
    target_compile_options(libtflite_kernel_utils PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
            -Wno-unused-function
    )

    target_link_libraries(libtflite_kernel_utils
            flatbuffer_headers
            gemmlowp_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libeigen
            liblog
            libm
            libruy_static
            tensorflow_headers
    )
endif ()
#################### libtflite_kernel_utils ####################

#################### libtflite_kernels ####################
set(FILES
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/activations.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/add.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/add_n.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/arg_min_max.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/atan2.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/basic_rnn.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/batch_matmul.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/batch_to_space_nd.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/bidirectional_sequence_lstm.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/bidirectional_sequence_rnn.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/broadcast_to.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/bucketize.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/call_once.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/cast.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/ceil.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/comparisons.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/complex_support.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/concatenation.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/conv.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/conv3d.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/cumsum.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/densify.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/depth_to_space.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/depthwise_conv.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/dequantize.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/detection_postprocess.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/div.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/dynamic_update_slice.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/elementwise.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/embedding_lookup.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/embedding_lookup_sparse.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/exp.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/expand_dims.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/fake_quant.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/fill.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/floor.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/floor_div.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/floor_mod.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/fully_connected.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/gather.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/gather_nd.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/hashtable_lookup.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/if.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/internal/kernel_utils.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/internal/optimized/cpu_check.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/internal/optimized/neon_tensor_utils.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/internal/optimized/sse_tensor_utils.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/internal/quantization_util.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/internal/reference/portable_tensor_utils.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/internal/tensor_utils.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/internal/transpose_utils.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/kernel_util.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/l2norm.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/local_response_norm.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/logical.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/lsh_projection.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/lstm.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/lstm_eval.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/matrix_diag.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/matrix_set_diag.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/maximum_minimum.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/mirror_pad.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/mul.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/neg.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/non_max_suppression.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/numeric_verify.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/one_hot.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/pack.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/pad.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/pooling.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/pow.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/quantize.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/random_uniform.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/random_ops.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/range.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/rank.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/reduce.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/register.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/reshape.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/resize_bilinear.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/resize_nearest_neighbor.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/reverse.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/reverse_sequence.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/rfft2d.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/round.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/scatter_nd.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/segment_sum.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/select.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/shape.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/sign.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/skip_gram.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/slice.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/space_to_batch_nd.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/space_to_depth.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/sparse_to_dense.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/split.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/split_v.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/squared_difference.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/squeeze.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/strided_slice.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/sub.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/svdf.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/tile.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/topk_v2.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/transpose.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/transpose_conv.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/unidirectional_sequence_lstm.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/unidirectional_sequence_rnn.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/unique.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/unpack.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/unsorted_segment.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/where.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/while.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/zeros_like.cc
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
    add_library(libtflite_kernels SHARED ${SOURCE_FILES})

    target_include_directories(libtflite_kernels
            PUBLIC
            ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels
    )

    # Add c/c++ flags
    target_compile_options(libtflite_kernels PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
            -Wno-unused-function
            -Wno-unused-local-typedef
            -Wno-unused-private-field
            -Wno-unused-variable
    )

    target_link_libraries(libtflite_kernels
            flatbuffer_headers
            gemmlowp_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libeigen
            libfft2d
            liblog
            libm
            libruy_headers
            libtextclassifier_hash_headers
            tensorflow_headers
    )
endif ()
#################### libtflite_kernels ####################
