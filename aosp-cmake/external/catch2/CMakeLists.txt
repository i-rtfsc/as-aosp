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

project(libcatch2)

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

#################### libcatch2 ####################
set(FILES
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_approx.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_assertionhandler.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_assertionresult.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_capture_matchers.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_commandline.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_common.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_config.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_console_colour.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_context.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_debug_console.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_debugger.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_decomposer.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_enforce.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_enum_values_registry.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_errno_guard.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_exception_translator_registry.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_fatal_condition.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_generators.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_interfaces_capture.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_interfaces_config.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_interfaces_exception.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_interfaces_registry_hub.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_interfaces_reporter.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_interfaces_runner.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_interfaces_testcase.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_leak_detector.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_list.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_matchers.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_matchers_exception.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_matchers_floating.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_matchers_generic.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_matchers_string.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_message.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_output_redirect.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_polyfills.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_random_number_generator.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_registry_hub.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_reporter_registry.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_result_type.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_run_context.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_section.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_section_info.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_session.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_singletons.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_startup_exception_registry.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_stream.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_string_manip.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_stringref.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_tag_alias.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_tag_alias_autoregistrar.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_tag_alias_registry.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_test_case_info.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_test_case_registry_impl.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_test_case_tracker.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_test_registry.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_test_spec.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_test_spec_parser.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_timer.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_tostring.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_totals.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_uncaught_exceptions.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_version.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_wildcard_pattern.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_xmlwriter.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/benchmark/detail/catch_stats.cpp
        ${ANDROID_ROOT}/external/catch2/include/reporters/catch_reporter_bases.cpp
        ${ANDROID_ROOT}/external/catch2/include/reporters/catch_reporter_compact.cpp
        ${ANDROID_ROOT}/external/catch2/include/reporters/catch_reporter_console.cpp
        ${ANDROID_ROOT}/external/catch2/include/reporters/catch_reporter_junit.cpp
        ${ANDROID_ROOT}/external/catch2/include/reporters/catch_reporter_listening.cpp
        ${ANDROID_ROOT}/external/catch2/include/reporters/catch_reporter_xml.cpp
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
    add_library(libcatch2 SHARED ${SOURCE_FILES})

    target_include_directories(libcatch2
            PUBLIC
            ${ANDROID_ROOT}/external/catch2/include
            ${ANDROID_ROOT}/external/catch2
    )

    # Add c/c++ flags
    target_compile_options(libcatch2 PRIVATE
            -Werror
    )

    target_link_libraries(libcatch2
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libcatch2 ####################

#################### libcatch2-main ####################
set(FILES
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_approx.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_assertionhandler.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_assertionresult.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_capture_matchers.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_commandline.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_common.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_config.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_console_colour.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_context.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_debug_console.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_debugger.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_decomposer.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_enforce.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_enum_values_registry.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_errno_guard.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_exception_translator_registry.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_fatal_condition.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_generators.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_interfaces_capture.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_interfaces_config.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_interfaces_exception.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_interfaces_registry_hub.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_interfaces_reporter.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_interfaces_runner.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_interfaces_testcase.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_leak_detector.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_list.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_matchers.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_matchers_exception.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_matchers_floating.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_matchers_generic.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_matchers_string.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_message.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_output_redirect.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_polyfills.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_random_number_generator.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_registry_hub.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_reporter_registry.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_result_type.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_run_context.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_section.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_section_info.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_session.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_singletons.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_startup_exception_registry.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_stream.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_string_manip.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_stringref.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_tag_alias.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_tag_alias_autoregistrar.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_tag_alias_registry.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_test_case_info.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_test_case_registry_impl.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_test_case_tracker.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_test_registry.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_test_spec.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_test_spec_parser.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_timer.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_tostring.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_totals.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_uncaught_exceptions.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_version.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_wildcard_pattern.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_xmlwriter.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/benchmark/detail/catch_stats.cpp
        ${ANDROID_ROOT}/external/catch2/include/reporters/catch_reporter_bases.cpp
        ${ANDROID_ROOT}/external/catch2/include/reporters/catch_reporter_compact.cpp
        ${ANDROID_ROOT}/external/catch2/include/reporters/catch_reporter_console.cpp
        ${ANDROID_ROOT}/external/catch2/include/reporters/catch_reporter_junit.cpp
        ${ANDROID_ROOT}/external/catch2/include/reporters/catch_reporter_listening.cpp
        ${ANDROID_ROOT}/external/catch2/include/reporters/catch_reporter_xml.cpp
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
    add_library(libcatch2-main SHARED ${SOURCE_FILES})

    target_include_directories(libcatch2-main
            PUBLIC
            ${ANDROID_ROOT}/external/catch2/include
            ${ANDROID_ROOT}/external/catch2
    )

    # Add c/c++ flags
    target_compile_options(libcatch2-main PRIVATE
            -Werror
    )

    target_link_libraries(libcatch2-main
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libcatch2-main ####################

#################### libcatch2-upstream ####################
set(FILES
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_approx.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_assertionhandler.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_assertionresult.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_capture_matchers.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_commandline.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_common.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_config.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_console_colour.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_context.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_debug_console.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_debugger.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_decomposer.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_enforce.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_enum_values_registry.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_errno_guard.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_exception_translator_registry.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_fatal_condition.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_generators.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_interfaces_capture.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_interfaces_config.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_interfaces_exception.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_interfaces_registry_hub.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_interfaces_reporter.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_interfaces_runner.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_interfaces_testcase.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_leak_detector.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_list.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_matchers.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_matchers_exception.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_matchers_floating.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_matchers_generic.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_matchers_string.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_message.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_output_redirect.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_polyfills.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_random_number_generator.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_registry_hub.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_reporter_registry.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_result_type.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_run_context.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_section.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_section_info.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_session.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_singletons.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_startup_exception_registry.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_stream.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_string_manip.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_stringref.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_tag_alias.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_tag_alias_autoregistrar.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_tag_alias_registry.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_test_case_info.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_test_case_registry_impl.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_test_case_tracker.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_test_registry.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_test_spec.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_test_spec_parser.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_timer.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_tostring.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_totals.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_uncaught_exceptions.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_version.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_wildcard_pattern.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/catch_xmlwriter.cpp
        ${ANDROID_ROOT}/external/catch2/include/internal/benchmark/detail/catch_stats.cpp
        ${ANDROID_ROOT}/external/catch2/include/reporters/catch_reporter_bases.cpp
        ${ANDROID_ROOT}/external/catch2/include/reporters/catch_reporter_compact.cpp
        ${ANDROID_ROOT}/external/catch2/include/reporters/catch_reporter_console.cpp
        ${ANDROID_ROOT}/external/catch2/include/reporters/catch_reporter_junit.cpp
        ${ANDROID_ROOT}/external/catch2/include/reporters/catch_reporter_listening.cpp
        ${ANDROID_ROOT}/external/catch2/include/reporters/catch_reporter_xml.cpp
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
    add_library(libcatch2-upstream SHARED ${SOURCE_FILES})

    target_include_directories(libcatch2-upstream
            PUBLIC
            ${ANDROID_ROOT}/external/catch2/include
            ${ANDROID_ROOT}/external/catch2
    )

    # Add c/c++ flags
    target_compile_options(libcatch2-upstream PRIVATE
            -Werror
    )

    target_link_libraries(libcatch2-upstream
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libcatch2-upstream ####################
