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

project(tensorflow_abseil)

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

#################### tensorflow_abseil ####################
set(FILES
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/base/inline_variable_test_a.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/base/inline_variable_test_b.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/base/log_severity.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/base/internal/atomic_hook_test_helper.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/base/internal/cycleclock.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/base/internal/exception_safety_testing.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/base/internal/low_level_alloc.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/base/internal/raw_logging.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/base/internal/scoped_set_env.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/base/internal/spinlock.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/base/internal/spinlock_wait.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/base/internal/strerror.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/base/internal/sysinfo.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/base/internal/thread_identity.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/base/internal/throw_delegate.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/base/internal/unscaledcycleclock.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/container/internal/hash_generator_testing.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/container/internal/hashtablez_sampler.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/container/internal/hashtablez_sampler_force_weak_definition.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/container/internal/raw_hash_set.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/container/internal/test_instance_tracker.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/debugging/failure_signal_handler.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/debugging/leak_check.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/debugging/stacktrace.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/debugging/symbolize.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/debugging/internal/address_is_readable.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/debugging/internal/demangle.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/debugging/internal/elf_mem_image.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/debugging/internal/examine_stack.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/debugging/internal/stack_consumption.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/debugging/internal/vdso_support.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/flags/commandlineflag.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/flags/flag.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/flags/flag_test_defs.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/flags/marshalling.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/flags/parse.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/flags/reflection.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/flags/usage.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/flags/usage_config.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/flags/internal/commandlineflag.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/flags/internal/flag.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/flags/internal/private_handle_accessor.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/flags/internal/program_name.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/flags/internal/usage.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/hash/internal/city.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/hash/internal/hash.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/hash/internal/low_level_hash.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/numeric/int128.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/profiling/internal/exponential_biased.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/profiling/internal/periodic_sampler.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/random/discrete_distribution.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/random/gaussian_distribution.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/random/seed_gen_exception.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/random/seed_sequences.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/random/internal/chi_square.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/random/internal/gaussian_distribution_gentables.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/random/internal/pool_urbg.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/random/internal/randen.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/random/internal/randen_benchmarks.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/random/internal/randen_detect.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/random/internal/randen_hwaes.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/random/internal/randen_round_keys.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/random/internal/randen_slow.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/random/internal/seed_material.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/status/status.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/status/status_payload_printer.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/status/statusor.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/strings/ascii.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/strings/charconv.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/strings/cord.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/strings/cord_analysis.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/strings/cord_buffer.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/strings/escaping.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/strings/match.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/strings/numbers.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/strings/str_cat.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/strings/str_replace.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/strings/str_split.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/strings/string_view.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/strings/substitute.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/strings/internal/charconv_bigint.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/strings/internal/charconv_parse.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/strings/internal/cord_internal.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/strings/internal/cord_rep_btree.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/strings/internal/cord_rep_btree_navigator.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/strings/internal/cord_rep_btree_reader.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/strings/internal/cord_rep_consume.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/strings/internal/cord_rep_crc.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/strings/internal/cord_rep_ring.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/strings/internal/cordz_functions.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/strings/internal/cordz_handle.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/strings/internal/cordz_info.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/strings/internal/cordz_sample_token.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/strings/internal/escaping.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/strings/internal/memutil.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/strings/internal/ostringstream.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/strings/internal/pow10_helper.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/strings/internal/utf8.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/strings/internal/str_format/arg.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/strings/internal/str_format/bind.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/strings/internal/str_format/extension.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/strings/internal/str_format/float_conversion.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/strings/internal/str_format/output.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/strings/internal/str_format/parser.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/synchronization/barrier.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/synchronization/blocking_counter.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/synchronization/mutex.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/synchronization/notification.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/synchronization/internal/create_thread_identity.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/synchronization/internal/graphcycles.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/synchronization/internal/per_thread_sem.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/synchronization/internal/waiter.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/time/civil_time.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/time/clock.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/time/duration.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/time/format.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/time/time.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/time/internal/cctz/src/civil_time_detail.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/time/internal/cctz/src/time_zone_fixed.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/time/internal/cctz/src/time_zone_format.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/time/internal/cctz/src/time_zone_if.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/time/internal/cctz/src/time_zone_impl.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/time/internal/cctz/src/time_zone_info.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/time/internal/cctz/src/time_zone_libc.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/time/internal/cctz/src/time_zone_lookup.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/time/internal/cctz/src/time_zone_posix.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/time/internal/cctz/src/zone_info_source.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/types/bad_any_cast.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/types/bad_optional_access.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/types/bad_variant_access.cc
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
    add_library(tensorflow_abseil SHARED ${SOURCE_FILES})

    target_include_directories(tensorflow_abseil
            PUBLIC
            ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp
    )

    # Add c/c++ flags
    target_compile_options(tensorflow_abseil PRIVATE
            -Werror
    )

    target_link_libraries(tensorflow_abseil
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### tensorflow_abseil ####################
