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

project(libjni_latinime)

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

#################### libjni_latinime ####################
set(FILES
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/com_android_inputmethod_keyboard_ProximityInfo.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/com_android_inputmethod_latin_BinaryDictionary.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/com_android_inputmethod_latin_BinaryDictionaryUtils.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/com_android_inputmethod_latin_DicTraverseSession.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/jni_common.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/dictionary/header/header_policy.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/dictionary/header/header_read_write_utils.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/dictionary/property/ngram_context.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/dictionary/structure/dictionary_structure_with_buffer_policy_factory.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/dictionary/structure/pt_common/bigram/bigram_list_read_write_utils.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/dictionary/structure/pt_common/dynamic_pt_gc_event_listeners.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/dictionary/structure/pt_common/dynamic_pt_reading_helper.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/dictionary/structure/pt_common/dynamic_pt_reading_utils.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/dictionary/structure/pt_common/dynamic_pt_updating_helper.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/dictionary/structure/pt_common/dynamic_pt_writing_utils.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/dictionary/structure/pt_common/patricia_trie_reading_utils.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/dictionary/structure/pt_common/shortcut/shortcut_list_reading_utils.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/dictionary/structure/v2/patricia_trie_policy.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/dictionary/structure/v2/ver2_patricia_trie_node_reader.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/dictionary/structure/v2/ver2_pt_node_array_reader.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/dictionary/structure/v4/ver4_dict_buffers.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/dictionary/structure/v4/ver4_dict_constants.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/dictionary/structure/v4/ver4_patricia_trie_node_reader.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/dictionary/structure/v4/ver4_patricia_trie_node_writer.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/dictionary/structure/v4/ver4_patricia_trie_policy.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/dictionary/structure/v4/ver4_patricia_trie_reading_utils.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/dictionary/structure/v4/ver4_patricia_trie_writing_helper.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/dictionary/structure/v4/ver4_pt_node_array_reader.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/dictionary/structure/v4/content/dynamic_language_model_probability_utils.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/dictionary/structure/v4/content/language_model_dict_content.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/dictionary/structure/v4/content/language_model_dict_content_global_counters.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/dictionary/structure/v4/content/shortcut_dict_content.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/dictionary/structure/v4/content/sparse_table_dict_content.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/dictionary/structure/v4/content/terminal_position_lookup_table.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/dictionary/utils/buffer_with_extendable_buffer.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/dictionary/utils/byte_array_utils.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/dictionary/utils/dict_file_writing_utils.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/dictionary/utils/file_utils.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/dictionary/utils/forgetting_curve_utils.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/dictionary/utils/format_utils.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/dictionary/utils/mmapped_buffer.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/dictionary/utils/multi_bigram_map.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/dictionary/utils/probability_utils.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/dictionary/utils/sparse_table.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/dictionary/utils/trie_map.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/suggest/core/suggest.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/suggest/core/dicnode/dic_node.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/suggest/core/dicnode/dic_node_utils.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/suggest/core/dicnode/dic_nodes_cache.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/suggest/core/dictionary/dictionary.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/suggest/core/dictionary/dictionary_utils.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/suggest/core/dictionary/digraph_utils.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/suggest/core/dictionary/error_type_utils.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/suggest/core/layout/additional_proximity_chars.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/suggest/core/layout/proximity_info.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/suggest/core/layout/proximity_info_params.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/suggest/core/layout/proximity_info_state.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/suggest/core/layout/proximity_info_state_utils.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/suggest/core/policy/weighting.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/suggest/core/session/dic_traverse_session.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/suggest/core/result/suggestion_results.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/suggest/core/result/suggestions_output_utils.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/suggest/policyimpl/gesture/gesture_suggest_policy_factory.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/suggest/policyimpl/typing/scoring_params.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/suggest/policyimpl/typing/typing_scoring.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/suggest/policyimpl/typing/typing_suggest_policy.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/suggest/policyimpl/typing/typing_traversal.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/suggest/policyimpl/typing/typing_weighting.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/utils/autocorrection_threshold_utils.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/utils/char_utils.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/utils/jni_data_utils.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/utils/log_utils.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/utils/time_keeper.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/dictionary/structure/backward/v402/ver4_dict_buffers.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/dictionary/structure/backward/v402/ver4_dict_constants.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/dictionary/structure/backward/v402/ver4_patricia_trie_node_reader.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/dictionary/structure/backward/v402/ver4_patricia_trie_node_writer.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/dictionary/structure/backward/v402/ver4_patricia_trie_policy.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/dictionary/structure/backward/v402/ver4_patricia_trie_reading_utils.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/dictionary/structure/backward/v402/ver4_patricia_trie_writing_helper.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/dictionary/structure/backward/v402/ver4_pt_node_array_reader.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/dictionary/structure/backward/v402/content/bigram_dict_content.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/dictionary/structure/backward/v402/content/probability_dict_content.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/dictionary/structure/backward/v402/content/shortcut_dict_content.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/dictionary/structure/backward/v402/content/sparse_table_dict_content.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/dictionary/structure/backward/v402/content/terminal_position_lookup_table.cpp
        ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src/dictionary/structure/backward/v402/bigram/ver4_bigram_list_policy.cpp
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
    add_library(libjni_latinime SHARED ${SOURCE_FILES})

    target_include_directories(libjni_latinime
            PUBLIC
            ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni/src
            ${ANDROID_ROOT}/packages/inputmethods/LatinIME/native/jni
    )

    # Add c/c++ flags
    target_compile_options(libjni_latinime PRIVATE
            -Werror
            -Wall
            -Wextra
            -Wno-unused-parameter
            -Wno-unused-function
    )

    target_link_libraries(libjni_latinime
            jni_headers
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libjni_latinime ####################
