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

project(libtextclassifier_bert_tokenizer)

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

#################### libtextclassifier_bert_tokenizer ####################
set(FILES
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/base/logging.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/base/logging_raw.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/bert_tokenizer.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/strings/utf8.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/tokenizer-utils.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/utf8/unilib-common.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/utf8/unicodetext.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/wordpiece_tokenizer.cc
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
    add_library(libtextclassifier_bert_tokenizer SHARED ${SOURCE_FILES})

    target_include_directories(libtextclassifier_bert_tokenizer
            PUBLIC
            ${ANDROID_ROOT}/external/libtextclassifier/native
            ${OUT_INTERMEDIATES_ROOT}/external/libtextclassifier/native/libtextclassifier_fbgen_utils_i18n_language-tag/gen
            ${OUT_INTERMEDIATES_ROOT}/external/libtextclassifier/native/libtextclassifier_fbgen_utils_tflite_text_encoder_config/gen
            ${OUT_INTERMEDIATES_ROOT}/external/libtextclassifier/native/libtextclassifier_fbgen_utils_resources/gen
            ${OUT_INTERMEDIATES_ROOT}/external/libtextclassifier/native/libtextclassifier_fbgen_utils_grammar_rules/gen
            ${OUT_INTERMEDIATES_ROOT}/external/libtextclassifier/native/libtextclassifier_fbgen_utils_grammar_semantics_expression/gen
            ${OUT_INTERMEDIATES_ROOT}/external/libtextclassifier/native/libtextclassifier_fbgen_utils_zlib_buffer/gen
            ${OUT_INTERMEDIATES_ROOT}/external/libtextclassifier/native/libtextclassifier_fbgen_utils_normalization/gen
            ${OUT_INTERMEDIATES_ROOT}/external/libtextclassifier/native/libtextclassifier_fbgen_utils_intents_intent-config/gen
            ${OUT_INTERMEDIATES_ROOT}/external/libtextclassifier/native/libtextclassifier_fbgen_utils_container_bit-vector/gen
            ${OUT_INTERMEDIATES_ROOT}/external/libtextclassifier/native/libtextclassifier_fbgen_utils_codepoint-range/gen
            ${OUT_INTERMEDIATES_ROOT}/external/libtextclassifier/native/libtextclassifier_fbgen_utils_tokenizer/gen
            ${OUT_INTERMEDIATES_ROOT}/external/libtextclassifier/native/libtextclassifier_fbgen_utils_flatbuffers_flatbuffers/gen
            ${OUT_INTERMEDIATES_ROOT}/external/libtextclassifier/native/libtextclassifier_fbgen_actions_actions_model/gen
            ${OUT_INTERMEDIATES_ROOT}/external/libtextclassifier/native/libtextclassifier_fbgen_actions_actions-entity-data/gen
            ${OUT_INTERMEDIATES_ROOT}/external/libtextclassifier/native/libtextclassifier_fbgen_lang_id_common_flatbuffers_embedding-network/gen
            ${OUT_INTERMEDIATES_ROOT}/external/libtextclassifier/native/libtextclassifier_fbgen_lang_id_common_flatbuffers_model/gen
            ${OUT_INTERMEDIATES_ROOT}/external/libtextclassifier/native/libtextclassifier_fbgen_annotator_entity-data/gen
            ${OUT_INTERMEDIATES_ROOT}/external/libtextclassifier/native/libtextclassifier_fbgen_annotator_person_name_person_name_model/gen
            ${OUT_INTERMEDIATES_ROOT}/external/libtextclassifier/native/libtextclassifier_fbgen_annotator_experimental_experimental/gen
            ${OUT_INTERMEDIATES_ROOT}/external/libtextclassifier/native/libtextclassifier_fbgen_annotator_model/gen
            ${OUT_INTERMEDIATES_ROOT}/external/libtextclassifier/native/libtextclassifier_fbgen_annotator_datetime_datetime/gen
    )

    # Add c/c++ flags
    target_compile_options(libtextclassifier_bert_tokenizer PRIVATE
            -Werror
            -Wno-unused-parameter
            -DTC3_DEBUG_LOGGING=1
    )

    target_link_libraries(libtextclassifier_bert_tokenizer
            flatbuffer_headers
            jni_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
            libtextclassifier_abseil
            libtextclassifier_flatbuffer_headers
            tensorflow_headers
            tflite_support
    )
endif ()
#################### libtextclassifier_bert_tokenizer ####################

#################### libtextclassifier_hash ####################
set(FILES
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/hash/farmhash.cc
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
    add_library(libtextclassifier_hash SHARED ${SOURCE_FILES})

    target_include_directories(libtextclassifier_hash
            PUBLIC
            ${ANDROID_ROOT}/external/libtextclassifier/native
    )

    # Add c/c++ flags
    target_compile_options(libtextclassifier_hash PRIVATE
            -Wall
            -Werror
            -Wno-unused-function
    )

    target_link_libraries(libtextclassifier_hash
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libtextclassifier_hash ####################

#################### libtextclassifier_hash_static ####################
set(FILES
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/hash/farmhash.cc
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
    add_library(libtextclassifier_hash_static SHARED ${SOURCE_FILES})

    target_include_directories(libtextclassifier_hash_static
            PUBLIC
            ${ANDROID_ROOT}/external/libtextclassifier/native
    )

    # Add c/c++ flags
    target_compile_options(libtextclassifier_hash_static PRIVATE
            -Wall
            -Werror
            -Wno-unused-function
    )

    target_link_libraries(libtextclassifier_hash_static
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libtextclassifier_hash_static ####################

#################### libtextclassifier ####################
set(FILES
        ${ANDROID_ROOT}/external/libtextclassifier/native/actions/actions-suggestions.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/actions/actions_jni.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/actions/feature-processor.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/actions/grammar-actions.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/actions/lua-actions.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/actions/lua-ranker.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/actions/ngram-model.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/actions/ranker.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/actions/regex-actions.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/actions/tflite-sensitive-model.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/actions/utils.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/actions/zlib-utils.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/annotator/annotator.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/annotator/annotator_jni.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/annotator/annotator_jni_common.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/annotator/cached-features.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/annotator/feature-processor.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/annotator/flatbuffer-utils.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/annotator/model-executor.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/annotator/quantization.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/annotator/strip-unpaired-brackets.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/annotator/types.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/annotator/zlib-utils.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/annotator/datetime/datetime-grounder.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/annotator/datetime/extractor.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/annotator/datetime/grammar-parser.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/annotator/datetime/regex-parser.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/annotator/datetime/utils.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/annotator/duration/duration.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/annotator/grammar/grammar-annotator.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/annotator/grammar/utils.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/annotator/number/number.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/annotator/pod_ner/pod-ner-impl.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/annotator/pod_ner/utils.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/annotator/translate/translate.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/annotator/vocab/vocab-annotator-impl.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/annotator/vocab/vocab-level-table.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/lang_id/custom-tokenizer.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/lang_id/lang-id-wrapper.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/lang_id/lang-id.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/lang_id/lang-id_jni.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/lang_id/common/embedding-feature-extractor.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/lang_id/common/embedding-network-params.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/lang_id/common/embedding-network.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/lang_id/common/utf8.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/lang_id/common/fel/feature-descriptors.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/lang_id/common/fel/feature-extractor.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/lang_id/common/fel/fel-parser.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/lang_id/common/fel/task-context.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/lang_id/common/fel/workspace.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/lang_id/common/file/file-utils.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/lang_id/common/file/mmap.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/lang_id/common/flatbuffers/embedding-network-params-from-flatbuffer.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/lang_id/common/flatbuffers/model-utils.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/lang_id/common/lite_base/compact-logging-raw.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/lang_id/common/lite_base/compact-logging.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/lang_id/common/lite_strings/numbers.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/lang_id/common/lite_strings/str-split.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/lang_id/common/math/checksum.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/lang_id/common/math/fastexp.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/lang_id/common/math/hash.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/lang_id/common/math/softmax.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/lang_id/fb_model/lang-id-from-fb.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/lang_id/fb_model/model-provider-from-fb.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/lang_id/features/char-ngram-feature.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/lang_id/features/light-sentence-features.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/lang_id/features/relevant-script-feature.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/lang_id/script/approx-script-data.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/lang_id/script/approx-script.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/lang_id/script/script-detector.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/lang_id/script/tiny-script-detector.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/tensorflow_models/seq_flow_lite/tflite_ops/layer_norm.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/bert_tokenizer.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/checksum.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/codepoint-range.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/lua-utils.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/normalization.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/regex-match.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/resources.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/tensor-view.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/tflite-model-executor.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/token-feature-extractor.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/tokenizer-utils.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/tokenizer.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/variant.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/wordpiece_tokenizer.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/base/arena.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/base/logging.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/base/logging_raw.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/base/status.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/calendar/calendar-javaicu.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/container/bit-vector.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/container/double-array-trie.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/container/sorted-strings-table.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/flatbuffers/mutable.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/flatbuffers/reflection.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/grammar/analyzer.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/grammar/rules-utils.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/grammar/parsing/derivation.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/grammar/parsing/lexer.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/grammar/parsing/matcher.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/grammar/parsing/parse-tree.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/grammar/parsing/parser.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/grammar/semantics/composer.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/grammar/semantics/evaluators/arithmetic-eval.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/grammar/semantics/evaluators/compose-eval.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/grammar/semantics/evaluators/merge-values-eval.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/grammar/utils/ir.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/grammar/utils/locale-shard-map.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/grammar/utils/rules.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/hash/cityhash.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/hash/farmhash.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/i18n/locale-list.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/i18n/locale.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/intents/intent-generator.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/intents/jni-lua.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/intents/jni.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/intents/zlib-utils.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/java/jni-base.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/java/jni-cache.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/java/jni-helper.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/math/fastexp.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/math/softmax.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/memory/mmap.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/sentencepiece/encoder.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/sentencepiece/normalizer.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/strings/append.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/strings/numbers.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/strings/split.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/strings/substitute.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/strings/utf8.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/tflite/blacklist.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/tflite/blacklist_base.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/tflite/dist_diversification.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/tflite/encoder_common.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/tflite/skipgram_finder.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/tflite/string_projection.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/tflite/string_projection_base.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/tflite/text_encoder.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/tflite/text_encoder3s.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/tflite/token_encoder.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/tokenfree/byte_encoder.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/utf8/unicodetext.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/utf8/unilib-common.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/utf8/unilib-javaicu.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/zlib/zlib.cc
        ${ANDROID_ROOT}/external/libtextclassifier/native/utils/zlib/zlib_regex.cc
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
    add_library(libtextclassifier SHARED ${SOURCE_FILES})

    target_include_directories(libtextclassifier
            PUBLIC
            ${ANDROID_ROOT}/external/libtextclassifier/native
            ${OUT_INTERMEDIATES_ROOT}/external/libtextclassifier/native/libtextclassifier_fbgen_utils_i18n_language-tag/gen
            ${OUT_INTERMEDIATES_ROOT}/external/libtextclassifier/native/libtextclassifier_fbgen_utils_tflite_text_encoder_config/gen
            ${OUT_INTERMEDIATES_ROOT}/external/libtextclassifier/native/libtextclassifier_fbgen_utils_resources/gen
            ${OUT_INTERMEDIATES_ROOT}/external/libtextclassifier/native/libtextclassifier_fbgen_utils_grammar_rules/gen
            ${OUT_INTERMEDIATES_ROOT}/external/libtextclassifier/native/libtextclassifier_fbgen_utils_grammar_semantics_expression/gen
            ${OUT_INTERMEDIATES_ROOT}/external/libtextclassifier/native/libtextclassifier_fbgen_utils_zlib_buffer/gen
            ${OUT_INTERMEDIATES_ROOT}/external/libtextclassifier/native/libtextclassifier_fbgen_utils_normalization/gen
            ${OUT_INTERMEDIATES_ROOT}/external/libtextclassifier/native/libtextclassifier_fbgen_utils_intents_intent-config/gen
            ${OUT_INTERMEDIATES_ROOT}/external/libtextclassifier/native/libtextclassifier_fbgen_utils_container_bit-vector/gen
            ${OUT_INTERMEDIATES_ROOT}/external/libtextclassifier/native/libtextclassifier_fbgen_utils_codepoint-range/gen
            ${OUT_INTERMEDIATES_ROOT}/external/libtextclassifier/native/libtextclassifier_fbgen_utils_tokenizer/gen
            ${OUT_INTERMEDIATES_ROOT}/external/libtextclassifier/native/libtextclassifier_fbgen_utils_flatbuffers_flatbuffers/gen
            ${OUT_INTERMEDIATES_ROOT}/external/libtextclassifier/native/libtextclassifier_fbgen_actions_actions_model/gen
            ${OUT_INTERMEDIATES_ROOT}/external/libtextclassifier/native/libtextclassifier_fbgen_actions_actions-entity-data/gen
            ${OUT_INTERMEDIATES_ROOT}/external/libtextclassifier/native/libtextclassifier_fbgen_lang_id_common_flatbuffers_embedding-network/gen
            ${OUT_INTERMEDIATES_ROOT}/external/libtextclassifier/native/libtextclassifier_fbgen_lang_id_common_flatbuffers_model/gen
            ${OUT_INTERMEDIATES_ROOT}/external/libtextclassifier/native/libtextclassifier_fbgen_annotator_entity-data/gen
            ${OUT_INTERMEDIATES_ROOT}/external/libtextclassifier/native/libtextclassifier_fbgen_annotator_person_name_person_name_model/gen
            ${OUT_INTERMEDIATES_ROOT}/external/libtextclassifier/native/libtextclassifier_fbgen_annotator_experimental_experimental/gen
            ${OUT_INTERMEDIATES_ROOT}/external/libtextclassifier/native/libtextclassifier_fbgen_annotator_model/gen
            ${OUT_INTERMEDIATES_ROOT}/external/libtextclassifier/native/libtextclassifier_fbgen_annotator_datetime_datetime/gen
    )

    # Add c/c++ flags
    target_compile_options(libtextclassifier PRIVATE
            -Wall
            -Werror
            -Wno-unused-function
            -Wno-unused-parameter
            -DTC3_DEBUG_LOGGING=1
    )

    target_link_libraries(libtextclassifier
            flatbuffer_headers
            jni_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libflatbuffers-cpp
            liblog
            liblua
            libm
            libtextclassifier_abseil
            libtextclassifier_flatbuffer_headers
            libtflite_static
            libutf
            libz
            marisa-trie
            tensorflow_headers
            tflite_support
    )
endif ()
#################### libtextclassifier ####################
