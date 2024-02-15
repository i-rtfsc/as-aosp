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

project(tflite_support_tokenizers)

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

#################### tflite_support_tokenizers ####################
set(FILES
        ${ANDROID_ROOT}/external/tflite-support/tensorflow_lite_support/cc/task/core/task_utils.cc
        ${ANDROID_ROOT}/external/tflite-support/tensorflow_lite_support/cc/text/tokenizers/regex_tokenizer.cc
        ${ANDROID_ROOT}/external/tflite-support/tensorflow_lite_support/cc/text/tokenizers/tokenizer_utils.cc
        ${ANDROID_ROOT}/external/tflite-support/tensorflow_lite_support/cc/utils/common_utils.cc
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
    add_library(tflite_support_tokenizers SHARED ${SOURCE_FILES})

    target_include_directories(tflite_support_tokenizers
            PUBLIC
            ${ANDROID_ROOT}/external/tflite-support
            ${OUT_INTERMEDIATES_ROOT}/external/tflite-support/tflite_support_metadata_schema/gen
    )

    # Add c/c++ flags
    target_compile_options(tflite_support_tokenizers PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(tflite_support_tokenizers
            flatbuffer_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
            libtextclassifier_abseil
            libtextclassifier_bert_tokenizer
            libtextclassifier_flatbuffer_headers
            libtflite_static
            tensorflow_headers
    )
endif ()
#################### tflite_support_tokenizers ####################

#################### tflite_support_metadata_extractor ####################
set(FILES
        ${ANDROID_ROOT}/external/tflite-support/tensorflow_lite_support/metadata/cc/metadata_extractor.cc
        ${ANDROID_ROOT}/external/tflite-support/tensorflow_lite_support/metadata/cc/utils/zip_readonly_mem_file.cc
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
    add_library(tflite_support_metadata_extractor SHARED ${SOURCE_FILES})

    target_include_directories(tflite_support_metadata_extractor
            PUBLIC
            ${ANDROID_ROOT}/external/tflite-support
            ${ANDROID_ROOT}/external/tflite-support/third_party/zlib
            ${OUT_INTERMEDIATES_ROOT}/external/tflite-support/tflite_support_metadata_schema/gen
    )

    # Add c/c++ flags
    target_compile_options(tflite_support_metadata_extractor PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(tflite_support_metadata_extractor
            flatbuffer_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libtextclassifier_abseil
            libtextclassifier_flatbuffer_headers
            tensorflow_headers
            tflite_support_libz
    )
endif ()
#################### tflite_support_metadata_extractor ####################

#################### tflite_support ####################
set(FILES
        ${ANDROID_ROOT}/external/tflite-support/tensorflow_lite_support/cc/utils/common_utils.cc
        ${ANDROID_ROOT}/external/tflite-support/tensorflow_lite_support/custom_ops/kernel/sentencepiece/sentencepiece_tokenizer_tflite.cc
        ${ANDROID_ROOT}/external/tflite-support/tensorflow_lite_support/custom_ops/kernel/sentencepiece/optimized_encoder.cc
        ${ANDROID_ROOT}/external/tflite-support/tensorflow_lite_support/custom_ops/kernel/ragged/ragged_range_tflite.cc
        ${ANDROID_ROOT}/external/tflite-support/tensorflow_lite_support/custom_ops/kernel/ragged/ragged_tensor_to_tensor_tflite.cc
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
    add_library(tflite_support SHARED ${SOURCE_FILES})

    target_include_directories(tflite_support
            PUBLIC
            ${ANDROID_ROOT}/external/tflite-support
            ${OUT_INTERMEDIATES_ROOT}/external/tflite-support/tflite_support_spm_config/gen
            ${OUT_INTERMEDIATES_ROOT}/external/tflite-support/tflite_support_spm_encoder_config/gen
    )

    # Add c/c++ flags
    target_compile_options(tflite_support PRIVATE
            -Werror
            -Wall
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(tflite_support
            flatbuffer_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
            libtextclassifier_abseil
            libtflite_static
            tensorflow_headers
    )
endif ()
#################### tflite_support ####################

#################### tflite_support_task_core_proto ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/external/tflite-support/tflite_support_task_core_proto/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow_lite_support/cc/task/core/proto/external_file.pb.cc
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
    add_library(tflite_support_task_core_proto SHARED ${SOURCE_FILES})

    target_include_directories(tflite_support_task_core_proto
            PUBLIC
            ${ANDROID_ROOT}/external/tflite-support
            ${OUT_INTERMEDIATES_ROOT}/external/tflite-support/tflite_support_task_core_proto/android_${OUT_ARCH_CPU}_static/gen/proto
    )

    # Add c/c++ flags
    target_compile_options(tflite_support_task_core_proto PRIVATE
            -Werror
    )

    target_link_libraries(tflite_support_task_core_proto
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libprotobuf-cpp-lite
    )
endif ()
#################### tflite_support_task_core_proto ####################

#################### libtflite_support_classifiers_native ####################
set(FILES
        ${ANDROID_ROOT}/external/tflite-support/tensorflow_lite_support/java/src/native/task/text/nlclassifier/bert_nlclassifier/bert_nl_classifier_jni.cc
        ${ANDROID_ROOT}/external/tflite-support/tensorflow_lite_support/java/src/native/task/text/nlclassifier/nl_classifier_jni.cc
        ${ANDROID_ROOT}/external/tflite-support/tensorflow_lite_support/java/src/native/task/text/nlclassifier/nl_classifier_jni_utils.cc
        ${ANDROID_ROOT}/external/tflite-support/tensorflow_lite_support/java/src/native/task/core/rbtml_op_resolver.cc
        ${ANDROID_ROOT}/external/tflite-support/tensorflow_lite_support/custom_ops/kernel/unsorted_segment.cc
        ${ANDROID_ROOT}/external/tflite-support/tensorflow_lite_support/cc/utils/jni_utils.cc
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
    add_library(libtflite_support_classifiers_native SHARED ${SOURCE_FILES})

    target_include_directories(libtflite_support_classifiers_native
            PUBLIC
            ${ANDROID_ROOT}/external/tflite-support
            ${ANDROID_ROOT}/external/tflite-support/third_party/zlib
            ${OUT_INTERMEDIATES_ROOT}/external/tflite-support/tflite_support_task_core_proto/android_${OUT_ARCH_CPU}_static/gen/proto
            ${OUT_INTERMEDIATES_ROOT}/external/tflite-support/tflite_support_metadata_schema/gen
    )

    # Add c/c++ flags
    target_compile_options(libtflite_support_classifiers_native PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(libtflite_support_classifiers_native
            flatbuffer_headers
            jni_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            liblog
            libm
            libprotobuf-cpp-lite-ndk
            libtextclassifier_abseil
            libtextclassifier_bert_tokenizer
            libtextclassifier_flatbuffer_headers
            libtflite_static
            tensorflow_headers
            tflite_configuration_proto
            tflite_support_classifiers
            tflite_support_libz
            tflite_support_metadata_extractor
            tflite_support_task_core
            tflite_support_task_core_proto
            tflite_support_tokenizers
    )
endif ()
#################### libtflite_support_classifiers_native ####################

#################### tflite_support_task_core ####################
set(FILES
        ${ANDROID_ROOT}/external/tflite-support/tensorflow_lite_support/cc/task/core/external_file_handler.cc
        ${ANDROID_ROOT}/external/tflite-support/tensorflow_lite_support/cc/task/core/task_utils.cc
        ${ANDROID_ROOT}/external/tflite-support/tensorflow_lite_support/cc/task/core/tflite_engine.cc
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
    add_library(tflite_support_task_core SHARED ${SOURCE_FILES})

    target_include_directories(tflite_support_task_core
            PUBLIC
            ${ANDROID_ROOT}/external/tflite-support
            ${OUT_INTERMEDIATES_ROOT}/external/tflite-support/tflite_support_task_core_proto/android_${OUT_ARCH_CPU}_static/gen/proto
            ${OUT_INTERMEDIATES_ROOT}/external/tflite-support/tflite_support_metadata_schema/gen
    )

    # Add c/c++ flags
    target_compile_options(tflite_support_task_core PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(tflite_support_task_core
            flatbuffer_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libtextclassifier_abseil
            tensorflow_headers
            tflite_configuration_proto
            tflite_support_task_core_proto
    )
endif ()
#################### tflite_support_task_core ####################

#################### tflite_support_classifiers ####################
set(FILES
        ${ANDROID_ROOT}/external/tflite-support/tensorflow_lite_support/cc/task/text/nlclassifier/bert_nl_classifier.cc
        ${ANDROID_ROOT}/external/tflite-support/tensorflow_lite_support/cc/task/text/nlclassifier/bert_nl_classifier_c_api.cc
        ${ANDROID_ROOT}/external/tflite-support/tensorflow_lite_support/cc/task/text/nlclassifier/nl_classifier.cc
        ${ANDROID_ROOT}/external/tflite-support/tensorflow_lite_support/cc/task/text/nlclassifier/nl_classifier_c_api.cc
        ${ANDROID_ROOT}/external/tflite-support/tensorflow_lite_support/cc/task/text/nlclassifier/nl_classifier_c_api_common.cc
        ${ANDROID_ROOT}/external/tflite-support/tensorflow_lite_support/cc/common.cc
        ${ANDROID_ROOT}/external/tflite-support/tensorflow_lite_support/cc/port/default/tflite_wrapper.cc
        ${ANDROID_ROOT}/external/tflite-support/tensorflow_lite_support/cc/port/default/statusor.cc
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
    add_library(tflite_support_classifiers SHARED ${SOURCE_FILES})

    target_include_directories(tflite_support_classifiers
            PUBLIC
            ${ANDROID_ROOT}/external/tflite-support
            ${OUT_INTERMEDIATES_ROOT}/external/tflite-support/tflite_support_task_core_proto/android_${OUT_ARCH_CPU}_static/gen/proto
            ${OUT_INTERMEDIATES_ROOT}/external/tflite-support/tflite_support_metadata_schema/gen
    )

    # Add c/c++ flags
    target_compile_options(tflite_support_classifiers PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(tflite_support_classifiers
            flatbuffer_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
            libtextclassifier_abseil
            libtflite_static
            tensorflow_headers
            tflite_configuration_proto
            tflite_support_metadata_extractor
            tflite_support_task_core_proto
            tflite_support_tokenizers
    )
endif ()
#################### tflite_support_classifiers ####################
