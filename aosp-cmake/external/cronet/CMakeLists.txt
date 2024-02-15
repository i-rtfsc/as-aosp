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

project(cronet_aml_crypto_crypto)

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

#################### cronet_aml_crypto_crypto ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/crypto/aead.cc
        ${ANDROID_ROOT}/external/cronet/crypto/ec_private_key.cc
        ${ANDROID_ROOT}/external/cronet/crypto/ec_signature_creator.cc
        ${ANDROID_ROOT}/external/cronet/crypto/ec_signature_creator_impl.cc
        ${ANDROID_ROOT}/external/cronet/crypto/encryptor.cc
        ${ANDROID_ROOT}/external/cronet/crypto/hkdf.cc
        ${ANDROID_ROOT}/external/cronet/crypto/hmac.cc
        ${ANDROID_ROOT}/external/cronet/crypto/openssl_util.cc
        ${ANDROID_ROOT}/external/cronet/crypto/p224_spake.cc
        ${ANDROID_ROOT}/external/cronet/crypto/random.cc
        ${ANDROID_ROOT}/external/cronet/crypto/rsa_private_key.cc
        ${ANDROID_ROOT}/external/cronet/crypto/secure_hash.cc
        ${ANDROID_ROOT}/external/cronet/crypto/secure_util.cc
        ${ANDROID_ROOT}/external/cronet/crypto/sha2.cc
        ${ANDROID_ROOT}/external/cronet/crypto/signature_creator.cc
        ${ANDROID_ROOT}/external/cronet/crypto/signature_verifier.cc
        ${ANDROID_ROOT}/external/cronet/crypto/symmetric_key.cc
        ${ANDROID_ROOT}/external/cronet/crypto/unexportable_key.cc
        ${ANDROID_ROOT}/external/cronet/crypto/unexportable_key_metrics.cc
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
    add_library(cronet_aml_crypto_crypto SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_crypto_crypto
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
            ${ANDROID_ROOT}/external/cronet/third_party/boringssl/src/include
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_chromecast_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_chromeos_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_debugging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_logging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_partition_alloc_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_anchor_functions_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_android_runtime_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_android_runtime_unchecked_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_base_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_build_date/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_cfi_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_clang_profiling_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_debugging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_feature_list_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_ios_cronet_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_logging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_message_pump_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_orderfile_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_power_monitor_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_profiler_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_rust_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_sanitizer_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_synchronization_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_tracing_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_blink_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_branding_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_chromecast_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_chromeos_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_config_compiler_compiler_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_crypto_buildflags/android_${OUT_ARCH_CPU}/gen
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_crypto_crypto PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_crypto_crypto
            cronet_aml_base_allocator_partition_allocator_partition_alloc
            cronet_aml_base_base
            cronet_aml_base_base_static
            cronet_aml_base_nodebug_assertion
            cronet_aml_base_third_party_double_conversion_double_conversion
            cronet_aml_base_third_party_dynamic_annotations_dynamic_annotations
            cronet_aml_third_party_abseil_cpp_absl_base_base
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_malloc_internal
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait
            cronet_aml_third_party_abseil_cpp_absl_base_strerror
            cronet_aml_third_party_abseil_cpp_absl_base_throw_delegate
            cronet_aml_third_party_abseil_cpp_absl_container_hashtablez_sampler
            cronet_aml_third_party_abseil_cpp_absl_container_raw_hash_set
            cronet_aml_third_party_abseil_cpp_absl_crc_cpu_detect
            cronet_aml_third_party_abseil_cpp_absl_crc_crc32c
            cronet_aml_third_party_abseil_cpp_absl_crc_crc_cord_state
            cronet_aml_third_party_abseil_cpp_absl_crc_crc_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_debugging_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_demangle_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_examine_stack
            cronet_aml_third_party_abseil_cpp_absl_debugging_failure_signal_handler
            cronet_aml_third_party_abseil_cpp_absl_debugging_stacktrace
            cronet_aml_third_party_abseil_cpp_absl_debugging_symbolize
            cronet_aml_third_party_abseil_cpp_absl_hash_city
            cronet_aml_third_party_abseil_cpp_absl_hash_hash
            cronet_aml_third_party_abseil_cpp_absl_hash_low_level_hash
            cronet_aml_third_party_abseil_cpp_absl_log_die_if_null
            cronet_aml_third_party_abseil_cpp_absl_log_globals
            cronet_aml_third_party_abseil_cpp_absl_log_internal_check_op
            cronet_aml_third_party_abseil_cpp_absl_log_internal_conditions
            cronet_aml_third_party_abseil_cpp_absl_log_internal_format
            cronet_aml_third_party_abseil_cpp_absl_log_internal_globals
            cronet_aml_third_party_abseil_cpp_absl_log_internal_log_message
            cronet_aml_third_party_abseil_cpp_absl_log_internal_log_sink_set
            cronet_aml_third_party_abseil_cpp_absl_log_internal_nullguard
            cronet_aml_third_party_abseil_cpp_absl_log_internal_proto
            cronet_aml_third_party_abseil_cpp_absl_log_log_entry
            cronet_aml_third_party_abseil_cpp_absl_log_log_sink
            cronet_aml_third_party_abseil_cpp_absl_numeric_int128
            cronet_aml_third_party_abseil_cpp_absl_profiling_exponential_biased
            cronet_aml_third_party_abseil_cpp_absl_random_distributions
            cronet_aml_third_party_abseil_cpp_absl_random_internal_platform
            cronet_aml_third_party_abseil_cpp_absl_random_internal_pool_urbg
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_hwaes
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_hwaes_impl
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_slow
            cronet_aml_third_party_abseil_cpp_absl_random_internal_seed_material
            cronet_aml_third_party_abseil_cpp_absl_random_seed_gen_exception
            cronet_aml_third_party_abseil_cpp_absl_random_seed_sequences
            cronet_aml_third_party_abseil_cpp_absl_status_status
            cronet_aml_third_party_abseil_cpp_absl_status_statusor
            cronet_aml_third_party_abseil_cpp_absl_strings_cord
            cronet_aml_third_party_abseil_cpp_absl_strings_cord_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_functions
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_handle
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_info
            cronet_aml_third_party_abseil_cpp_absl_strings_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_str_format_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_strings
            cronet_aml_third_party_abseil_cpp_absl_synchronization_graphcycles_internal
            cronet_aml_third_party_abseil_cpp_absl_synchronization_kernel_timeout_internal
            cronet_aml_third_party_abseil_cpp_absl_synchronization_synchronization
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_civil_time
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_time_zone
            cronet_aml_third_party_abseil_cpp_absl_time_time
            cronet_aml_third_party_abseil_cpp_absl_types_bad_optional_access
            cronet_aml_third_party_abseil_cpp_absl_types_bad_variant_access
            cronet_aml_third_party_android_ndk_cpu_features
            cronet_aml_third_party_ashmem_ashmem
            cronet_aml_third_party_icu_icui18n
            cronet_aml_third_party_icu_icuuc_private
            cronet_aml_third_party_libevent_libevent
            cronet_aml_third_party_modp_b64_modp_b64
            jni_headers
            libandroid
            libc
            libclang_rt.builtins
            libcrypto
            libdl
            liblog
            libm
            libmediandk
            libssl
    )
endif ()
#################### cronet_aml_crypto_crypto ####################

#################### cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_hwaes_impl ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/random/internal/randen_hwaes.cc
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
    add_library(cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_hwaes_impl SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_hwaes_impl
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_chromeos_buildflags/android_${OUT_ARCH_CPU}/gen
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_hwaes_impl PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_hwaes_impl
            cronet_aml_third_party_abseil_cpp_absl_numeric_int128
            cronet_aml_third_party_abseil_cpp_absl_random_internal_platform
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_hwaes_impl ####################

#################### cronet_aml_third_party_abseil_cpp_absl_base_base ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/base/internal/cycleclock.cc
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/base/internal/spinlock.cc
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/base/internal/sysinfo.cc
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/base/internal/thread_identity.cc
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/base/internal/unscaledcycleclock.cc
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
    add_library(cronet_aml_third_party_abseil_cpp_absl_base_base SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_abseil_cpp_absl_base_base
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_abseil_cpp_absl_base_base PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_abseil_cpp_absl_base_base
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_abseil_cpp_absl_base_base ####################

#################### cronet_aml_third_party_abseil_cpp_absl_log_internal_check_op ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/log/internal/check_op.cc
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
    add_library(cronet_aml_third_party_abseil_cpp_absl_log_internal_check_op SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_abseil_cpp_absl_log_internal_check_op
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_abseil_cpp_absl_log_internal_check_op PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_abseil_cpp_absl_log_internal_check_op
            cronet_aml_third_party_abseil_cpp_absl_base_base
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_malloc_internal
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait
            cronet_aml_third_party_abseil_cpp_absl_base_strerror
            cronet_aml_third_party_abseil_cpp_absl_base_throw_delegate
            cronet_aml_third_party_abseil_cpp_absl_debugging_debugging_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_demangle_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_examine_stack
            cronet_aml_third_party_abseil_cpp_absl_debugging_stacktrace
            cronet_aml_third_party_abseil_cpp_absl_debugging_symbolize
            cronet_aml_third_party_abseil_cpp_absl_hash_city
            cronet_aml_third_party_abseil_cpp_absl_hash_hash
            cronet_aml_third_party_abseil_cpp_absl_hash_low_level_hash
            cronet_aml_third_party_abseil_cpp_absl_log_globals
            cronet_aml_third_party_abseil_cpp_absl_log_internal_format
            cronet_aml_third_party_abseil_cpp_absl_log_internal_globals
            cronet_aml_third_party_abseil_cpp_absl_log_internal_log_message
            cronet_aml_third_party_abseil_cpp_absl_log_internal_log_sink_set
            cronet_aml_third_party_abseil_cpp_absl_log_internal_nullguard
            cronet_aml_third_party_abseil_cpp_absl_log_internal_proto
            cronet_aml_third_party_abseil_cpp_absl_log_log_entry
            cronet_aml_third_party_abseil_cpp_absl_log_log_sink
            cronet_aml_third_party_abseil_cpp_absl_numeric_int128
            cronet_aml_third_party_abseil_cpp_absl_strings_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_str_format_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_strings
            cronet_aml_third_party_abseil_cpp_absl_synchronization_graphcycles_internal
            cronet_aml_third_party_abseil_cpp_absl_synchronization_kernel_timeout_internal
            cronet_aml_third_party_abseil_cpp_absl_synchronization_synchronization
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_civil_time
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_time_zone
            cronet_aml_third_party_abseil_cpp_absl_time_time
            cronet_aml_third_party_abseil_cpp_absl_types_bad_optional_access
            cronet_aml_third_party_abseil_cpp_absl_types_bad_variant_access
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            liblog
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_abseil_cpp_absl_log_internal_check_op ####################

#################### cronet_aml_net_http_transport_security_state_generated_files ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/net/http/transport_security_state.cc
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
    add_library(cronet_aml_net_http_transport_security_state_generated_files SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_net_http_transport_security_state_generated_files
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/overrides
            ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src
            ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/common/platform/default
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
            ${ANDROID_ROOT}/external/cronet/third_party/boringssl/src/include
            ${ANDROID_ROOT}/external/cronet/third_party/brotli/include
            ${ANDROID_ROOT}/external/cronet/third_party/protobuf/src
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_chromecast_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_chromeos_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_debugging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_logging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_partition_alloc_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_anchor_functions_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_android_runtime_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_android_runtime_unchecked_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_base_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_build_date/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_cfi_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_clang_profiling_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_debugging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_feature_list_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_ios_cronet_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_logging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_message_pump_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_orderfile_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_power_monitor_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_profiler_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_rust_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_sanitizer_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_synchronization_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_tracing_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_blink_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_branding_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_chromecast_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_chromeos_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_config_compiler_compiler_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_crypto_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_base_registry_controlled_domains_registry_controlled_domains/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_cronet_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_isolation_info_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_isolation_info_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/net/base
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_isolation_info_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/protos
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_net_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_url_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_net_nqe_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_net_nqe_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/net/nqe/proto
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_net_nqe_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/protos
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_third_party_quiche_net_quic_test_tools_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_third_party_quiche_net_quic_test_tools_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/net/third_party/quiche/src
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_third_party_quiche_net_quic_test_tools_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/net/third_party/quiche/src/quiche/quic/test_tools
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_third_party_quiche_net_quic_test_tools_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/protos
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_third_party_quiche_net_quic_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_third_party_quiche_net_quic_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/net/third_party/quiche/src
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_third_party_quiche_net_quic_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/protos
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_url_url_jni_headers/android_${OUT_ARCH_CPU}/gen
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_net_http_transport_security_state_generated_files PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_net_http_transport_security_state_generated_files
            cronet_aml_base_allocator_partition_allocator_partition_alloc
            cronet_aml_base_base
            cronet_aml_base_base_static
            cronet_aml_base_nodebug_assertion
            cronet_aml_base_third_party_double_conversion_double_conversion
            cronet_aml_base_third_party_dynamic_annotations_dynamic_annotations
            cronet_aml_crypto_crypto
            cronet_aml_net_dns_dns
            cronet_aml_net_dns_public_public
            cronet_aml_net_net_deps
            cronet_aml_net_net_public_deps
            cronet_aml_net_preload_decoder
            cronet_aml_net_third_party_quiche_quiche
            cronet_aml_net_traffic_annotation_traffic_annotation
            cronet_aml_net_uri_template
            cronet_aml_third_party_abseil_cpp_absl_base_base
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_malloc_internal
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait
            cronet_aml_third_party_abseil_cpp_absl_base_strerror
            cronet_aml_third_party_abseil_cpp_absl_base_throw_delegate
            cronet_aml_third_party_abseil_cpp_absl_container_hashtablez_sampler
            cronet_aml_third_party_abseil_cpp_absl_container_raw_hash_set
            cronet_aml_third_party_abseil_cpp_absl_crc_cpu_detect
            cronet_aml_third_party_abseil_cpp_absl_crc_crc32c
            cronet_aml_third_party_abseil_cpp_absl_crc_crc_cord_state
            cronet_aml_third_party_abseil_cpp_absl_crc_crc_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_debugging_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_demangle_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_examine_stack
            cronet_aml_third_party_abseil_cpp_absl_debugging_failure_signal_handler
            cronet_aml_third_party_abseil_cpp_absl_debugging_stacktrace
            cronet_aml_third_party_abseil_cpp_absl_debugging_symbolize
            cronet_aml_third_party_abseil_cpp_absl_hash_city
            cronet_aml_third_party_abseil_cpp_absl_hash_hash
            cronet_aml_third_party_abseil_cpp_absl_hash_low_level_hash
            cronet_aml_third_party_abseil_cpp_absl_log_die_if_null
            cronet_aml_third_party_abseil_cpp_absl_log_globals
            cronet_aml_third_party_abseil_cpp_absl_log_internal_check_op
            cronet_aml_third_party_abseil_cpp_absl_log_internal_conditions
            cronet_aml_third_party_abseil_cpp_absl_log_internal_format
            cronet_aml_third_party_abseil_cpp_absl_log_internal_globals
            cronet_aml_third_party_abseil_cpp_absl_log_internal_log_message
            cronet_aml_third_party_abseil_cpp_absl_log_internal_log_sink_set
            cronet_aml_third_party_abseil_cpp_absl_log_internal_nullguard
            cronet_aml_third_party_abseil_cpp_absl_log_internal_proto
            cronet_aml_third_party_abseil_cpp_absl_log_log_entry
            cronet_aml_third_party_abseil_cpp_absl_log_log_sink
            cronet_aml_third_party_abseil_cpp_absl_numeric_int128
            cronet_aml_third_party_abseil_cpp_absl_profiling_exponential_biased
            cronet_aml_third_party_abseil_cpp_absl_random_distributions
            cronet_aml_third_party_abseil_cpp_absl_random_internal_platform
            cronet_aml_third_party_abseil_cpp_absl_random_internal_pool_urbg
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_hwaes
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_hwaes_impl
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_slow
            cronet_aml_third_party_abseil_cpp_absl_random_internal_seed_material
            cronet_aml_third_party_abseil_cpp_absl_random_seed_gen_exception
            cronet_aml_third_party_abseil_cpp_absl_random_seed_sequences
            cronet_aml_third_party_abseil_cpp_absl_status_status
            cronet_aml_third_party_abseil_cpp_absl_status_statusor
            cronet_aml_third_party_abseil_cpp_absl_strings_cord
            cronet_aml_third_party_abseil_cpp_absl_strings_cord_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_functions
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_handle
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_info
            cronet_aml_third_party_abseil_cpp_absl_strings_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_str_format_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_strings
            cronet_aml_third_party_abseil_cpp_absl_synchronization_graphcycles_internal
            cronet_aml_third_party_abseil_cpp_absl_synchronization_kernel_timeout_internal
            cronet_aml_third_party_abseil_cpp_absl_synchronization_synchronization
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_civil_time
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_time_zone
            cronet_aml_third_party_abseil_cpp_absl_time_time
            cronet_aml_third_party_abseil_cpp_absl_types_bad_optional_access
            cronet_aml_third_party_abseil_cpp_absl_types_bad_variant_access
            cronet_aml_third_party_android_ndk_cpu_features
            cronet_aml_third_party_ashmem_ashmem
            cronet_aml_third_party_brotli_common
            cronet_aml_third_party_brotli_dec
            cronet_aml_third_party_icu_icui18n
            cronet_aml_third_party_icu_icuuc_private
            cronet_aml_third_party_libevent_libevent
            cronet_aml_third_party_modp_b64_modp_b64
            cronet_aml_third_party_protobuf_protobuf_lite
            cronet_aml_url_url
            jni_headers
            libandroid
            libc
            libclang_rt.builtins
            libcrypto
            libdl
            liblog
            libm
            libmediandk
            libssl
            libz
    )
endif ()
#################### cronet_aml_net_http_transport_security_state_generated_files ####################

#################### cronet_aml_third_party_abseil_cpp_absl_time_time ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/time/civil_time.cc
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/time/clock.cc
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/time/duration.cc
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/time/format.cc
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/time/time.cc
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
    add_library(cronet_aml_third_party_abseil_cpp_absl_time_time SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_abseil_cpp_absl_time_time
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_abseil_cpp_absl_time_time PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_abseil_cpp_absl_time_time
            cronet_aml_third_party_abseil_cpp_absl_base_base
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait
            cronet_aml_third_party_abseil_cpp_absl_base_throw_delegate
            cronet_aml_third_party_abseil_cpp_absl_numeric_int128
            cronet_aml_third_party_abseil_cpp_absl_strings_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_strings
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_civil_time
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_time_zone
            cronet_aml_third_party_abseil_cpp_absl_types_bad_optional_access
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_abseil_cpp_absl_time_time ####################

#################### cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_slow ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/random/internal/randen_slow.cc
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
    add_library(cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_slow SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_slow
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_chromeos_buildflags/android_${OUT_ARCH_CPU}/gen
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_slow PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_slow
            cronet_aml_third_party_abseil_cpp_absl_base_base
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait
            cronet_aml_third_party_abseil_cpp_absl_numeric_int128
            cronet_aml_third_party_abseil_cpp_absl_random_internal_platform
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_slow ####################

#################### cronet_aml_third_party_abseil_cpp_absl_hash_city ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/hash/internal/city.cc
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
    add_library(cronet_aml_third_party_abseil_cpp_absl_hash_city SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_abseil_cpp_absl_hash_city
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_abseil_cpp_absl_hash_city PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_abseil_cpp_absl_hash_city
            cronet_aml_third_party_abseil_cpp_absl_base_base
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_abseil_cpp_absl_hash_city ####################

#################### cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_hwaes ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/random/internal/randen_detect.cc
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
    add_library(cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_hwaes SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_hwaes
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_chromeos_buildflags/android_${OUT_ARCH_CPU}/gen
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_hwaes PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_hwaes
            cronet_aml_third_party_abseil_cpp_absl_numeric_int128
            cronet_aml_third_party_abseil_cpp_absl_random_internal_platform
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_hwaes_impl
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_hwaes ####################

#################### cronet_aml_third_party_abseil_cpp_absl_log_internal_conditions ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/log/internal/conditions.cc
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
    add_library(cronet_aml_third_party_abseil_cpp_absl_log_internal_conditions SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_abseil_cpp_absl_log_internal_conditions
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_abseil_cpp_absl_log_internal_conditions PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_abseil_cpp_absl_log_internal_conditions
            cronet_aml_third_party_abseil_cpp_absl_base_base
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_abseil_cpp_absl_log_internal_conditions ####################

#################### cronet_aml_third_party_abseil_cpp_absl_strings_internal ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/strings/internal/escaping.cc
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/strings/internal/ostringstream.cc
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/strings/internal/utf8.cc
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
    add_library(cronet_aml_third_party_abseil_cpp_absl_strings_internal SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_abseil_cpp_absl_strings_internal
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_abseil_cpp_absl_strings_internal PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_abseil_cpp_absl_strings_internal
            cronet_aml_third_party_abseil_cpp_absl_base_base
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_abseil_cpp_absl_strings_internal ####################

#################### cronet_aml_third_party_abseil_cpp_absl_log_globals ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/log/globals.cc
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
    add_library(cronet_aml_third_party_abseil_cpp_absl_log_globals SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_abseil_cpp_absl_log_globals
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_abseil_cpp_absl_log_globals PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_abseil_cpp_absl_log_globals
            cronet_aml_third_party_abseil_cpp_absl_base_base
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait
            cronet_aml_third_party_abseil_cpp_absl_base_throw_delegate
            cronet_aml_third_party_abseil_cpp_absl_hash_city
            cronet_aml_third_party_abseil_cpp_absl_hash_hash
            cronet_aml_third_party_abseil_cpp_absl_hash_low_level_hash
            cronet_aml_third_party_abseil_cpp_absl_numeric_int128
            cronet_aml_third_party_abseil_cpp_absl_strings_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_strings
            cronet_aml_third_party_abseil_cpp_absl_types_bad_optional_access
            cronet_aml_third_party_abseil_cpp_absl_types_bad_variant_access
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_abseil_cpp_absl_log_globals ####################

#################### cronet_aml_third_party_abseil_cpp_absl_log_log_sink ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/log/log_sink.cc
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
    add_library(cronet_aml_third_party_abseil_cpp_absl_log_log_sink SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_abseil_cpp_absl_log_log_sink
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_abseil_cpp_absl_log_log_sink PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_abseil_cpp_absl_log_log_sink
            cronet_aml_third_party_abseil_cpp_absl_base_base
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait
            cronet_aml_third_party_abseil_cpp_absl_base_throw_delegate
            cronet_aml_third_party_abseil_cpp_absl_log_log_entry
            cronet_aml_third_party_abseil_cpp_absl_numeric_int128
            cronet_aml_third_party_abseil_cpp_absl_strings_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_strings
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_civil_time
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_time_zone
            cronet_aml_third_party_abseil_cpp_absl_time_time
            cronet_aml_third_party_abseil_cpp_absl_types_bad_optional_access
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_abseil_cpp_absl_log_log_sink ####################

#################### cronet_aml_net_dns_public_public ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/net/dns/public/dns_config_overrides.cc
        ${ANDROID_ROOT}/external/cronet/net/dns/public/dns_over_https_config.cc
        ${ANDROID_ROOT}/external/cronet/net/dns/public/dns_over_https_server_config.cc
        ${ANDROID_ROOT}/external/cronet/net/dns/public/dns_query_type.cc
        ${ANDROID_ROOT}/external/cronet/net/dns/public/doh_provider_entry.cc
        ${ANDROID_ROOT}/external/cronet/net/dns/public/host_resolver_results.cc
        ${ANDROID_ROOT}/external/cronet/net/dns/public/resolve_error_info.cc
        ${ANDROID_ROOT}/external/cronet/net/dns/public/util.cc
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
    add_library(cronet_aml_net_dns_public_public SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_net_dns_public_public
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/overrides
            ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src
            ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/common/platform/default
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
            ${ANDROID_ROOT}/external/cronet/third_party/boringssl/src/include
            ${ANDROID_ROOT}/external/cronet/third_party/brotli/include
            ${ANDROID_ROOT}/external/cronet/third_party/protobuf/src
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_chromecast_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_chromeos_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_debugging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_logging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_partition_alloc_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_anchor_functions_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_android_runtime_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_android_runtime_unchecked_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_base_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_build_date/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_cfi_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_clang_profiling_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_debugging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_feature_list_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_ios_cronet_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_logging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_message_pump_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_orderfile_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_power_monitor_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_profiler_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_rust_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_sanitizer_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_synchronization_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_tracing_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_blink_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_branding_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_chromecast_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_chromeos_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_config_compiler_compiler_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_crypto_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_base_registry_controlled_domains_registry_controlled_domains/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_cronet_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_isolation_info_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_isolation_info_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/net/base
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_isolation_info_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/protos
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_net_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_url_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_net_nqe_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_net_nqe_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/net/nqe/proto
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_net_nqe_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/protos
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_third_party_quiche_net_quic_test_tools_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_third_party_quiche_net_quic_test_tools_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/net/third_party/quiche/src
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_third_party_quiche_net_quic_test_tools_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/net/third_party/quiche/src/quiche/quic/test_tools
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_third_party_quiche_net_quic_test_tools_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/protos
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_third_party_quiche_net_quic_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_third_party_quiche_net_quic_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/net/third_party/quiche/src
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_third_party_quiche_net_quic_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/protos
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_url_url_jni_headers/android_${OUT_ARCH_CPU}/gen
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_net_dns_public_public PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_net_dns_public_public
            cronet_aml_base_allocator_partition_allocator_partition_alloc
            cronet_aml_base_base
            cronet_aml_base_base_static
            cronet_aml_base_nodebug_assertion
            cronet_aml_base_third_party_double_conversion_double_conversion
            cronet_aml_base_third_party_dynamic_annotations_dynamic_annotations
            cronet_aml_crypto_crypto
            cronet_aml_net_net_deps
            cronet_aml_net_net_public_deps
            cronet_aml_net_preload_decoder
            cronet_aml_net_third_party_quiche_quiche
            cronet_aml_net_traffic_annotation_traffic_annotation
            cronet_aml_net_uri_template
            cronet_aml_third_party_abseil_cpp_absl_base_base
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_malloc_internal
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait
            cronet_aml_third_party_abseil_cpp_absl_base_strerror
            cronet_aml_third_party_abseil_cpp_absl_base_throw_delegate
            cronet_aml_third_party_abseil_cpp_absl_container_hashtablez_sampler
            cronet_aml_third_party_abseil_cpp_absl_container_raw_hash_set
            cronet_aml_third_party_abseil_cpp_absl_crc_cpu_detect
            cronet_aml_third_party_abseil_cpp_absl_crc_crc32c
            cronet_aml_third_party_abseil_cpp_absl_crc_crc_cord_state
            cronet_aml_third_party_abseil_cpp_absl_crc_crc_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_debugging_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_demangle_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_examine_stack
            cronet_aml_third_party_abseil_cpp_absl_debugging_failure_signal_handler
            cronet_aml_third_party_abseil_cpp_absl_debugging_stacktrace
            cronet_aml_third_party_abseil_cpp_absl_debugging_symbolize
            cronet_aml_third_party_abseil_cpp_absl_hash_city
            cronet_aml_third_party_abseil_cpp_absl_hash_hash
            cronet_aml_third_party_abseil_cpp_absl_hash_low_level_hash
            cronet_aml_third_party_abseil_cpp_absl_log_die_if_null
            cronet_aml_third_party_abseil_cpp_absl_log_globals
            cronet_aml_third_party_abseil_cpp_absl_log_internal_check_op
            cronet_aml_third_party_abseil_cpp_absl_log_internal_conditions
            cronet_aml_third_party_abseil_cpp_absl_log_internal_format
            cronet_aml_third_party_abseil_cpp_absl_log_internal_globals
            cronet_aml_third_party_abseil_cpp_absl_log_internal_log_message
            cronet_aml_third_party_abseil_cpp_absl_log_internal_log_sink_set
            cronet_aml_third_party_abseil_cpp_absl_log_internal_nullguard
            cronet_aml_third_party_abseil_cpp_absl_log_internal_proto
            cronet_aml_third_party_abseil_cpp_absl_log_log_entry
            cronet_aml_third_party_abseil_cpp_absl_log_log_sink
            cronet_aml_third_party_abseil_cpp_absl_numeric_int128
            cronet_aml_third_party_abseil_cpp_absl_profiling_exponential_biased
            cronet_aml_third_party_abseil_cpp_absl_random_distributions
            cronet_aml_third_party_abseil_cpp_absl_random_internal_platform
            cronet_aml_third_party_abseil_cpp_absl_random_internal_pool_urbg
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_hwaes
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_hwaes_impl
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_slow
            cronet_aml_third_party_abseil_cpp_absl_random_internal_seed_material
            cronet_aml_third_party_abseil_cpp_absl_random_seed_gen_exception
            cronet_aml_third_party_abseil_cpp_absl_random_seed_sequences
            cronet_aml_third_party_abseil_cpp_absl_status_status
            cronet_aml_third_party_abseil_cpp_absl_status_statusor
            cronet_aml_third_party_abseil_cpp_absl_strings_cord
            cronet_aml_third_party_abseil_cpp_absl_strings_cord_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_functions
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_handle
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_info
            cronet_aml_third_party_abseil_cpp_absl_strings_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_str_format_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_strings
            cronet_aml_third_party_abseil_cpp_absl_synchronization_graphcycles_internal
            cronet_aml_third_party_abseil_cpp_absl_synchronization_kernel_timeout_internal
            cronet_aml_third_party_abseil_cpp_absl_synchronization_synchronization
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_civil_time
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_time_zone
            cronet_aml_third_party_abseil_cpp_absl_time_time
            cronet_aml_third_party_abseil_cpp_absl_types_bad_optional_access
            cronet_aml_third_party_abseil_cpp_absl_types_bad_variant_access
            cronet_aml_third_party_android_ndk_cpu_features
            cronet_aml_third_party_ashmem_ashmem
            cronet_aml_third_party_brotli_common
            cronet_aml_third_party_brotli_dec
            cronet_aml_third_party_icu_icui18n
            cronet_aml_third_party_icu_icuuc_private
            cronet_aml_third_party_libevent_libevent
            cronet_aml_third_party_modp_b64_modp_b64
            cronet_aml_third_party_protobuf_protobuf_lite
            cronet_aml_url_url
            jni_headers
            libandroid
            libc
            libclang_rt.builtins
            libcrypto
            libdl
            liblog
            libm
            libmediandk
            libssl
            libz
    )
endif ()
#################### cronet_aml_net_dns_public_public ####################

#################### cronet_aml_components_cronet_android_integrated_mode_state_preprocess ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_components_cronet_android_integrated_mode_state_rename/gen/IntegratedModeState.cc
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
    add_library(cronet_aml_components_cronet_android_integrated_mode_state_preprocess SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_components_cronet_android_integrated_mode_state_preprocess
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_components_cronet_android_integrated_mode_state_preprocess PRIVATE
            -Werror
    )

    target_link_libraries(cronet_aml_components_cronet_android_integrated_mode_state_preprocess
            libc
            libc++
            libc++demangle
            libdl
            libm
    )
endif ()
#################### cronet_aml_components_cronet_android_integrated_mode_state_preprocess ####################

#################### cronet_aml_third_party_abseil_cpp_absl_hash_hash ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/hash/internal/hash.cc
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
    add_library(cronet_aml_third_party_abseil_cpp_absl_hash_hash SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_abseil_cpp_absl_hash_hash
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_abseil_cpp_absl_hash_hash PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_abseil_cpp_absl_hash_hash
            cronet_aml_third_party_abseil_cpp_absl_base_base
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait
            cronet_aml_third_party_abseil_cpp_absl_base_throw_delegate
            cronet_aml_third_party_abseil_cpp_absl_hash_city
            cronet_aml_third_party_abseil_cpp_absl_hash_low_level_hash
            cronet_aml_third_party_abseil_cpp_absl_numeric_int128
            cronet_aml_third_party_abseil_cpp_absl_strings_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_strings
            cronet_aml_third_party_abseil_cpp_absl_types_bad_optional_access
            cronet_aml_third_party_abseil_cpp_absl_types_bad_variant_access
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_abseil_cpp_absl_hash_hash ####################

#################### cronet_aml_third_party_abseil_cpp_absl_crc_cpu_detect ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/crc/internal/cpu_detect.cc
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
    add_library(cronet_aml_third_party_abseil_cpp_absl_crc_cpu_detect SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_abseil_cpp_absl_crc_cpu_detect
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_abseil_cpp_absl_crc_cpu_detect PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_abseil_cpp_absl_crc_cpu_detect
            cronet_aml_third_party_abseil_cpp_absl_base_base
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_abseil_cpp_absl_crc_cpu_detect ####################

#################### cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/base/internal/spinlock_wait.cc
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
    add_library(cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait ####################

#################### cronet_aml_third_party_abseil_cpp_absl_synchronization_synchronization ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/synchronization/barrier.cc
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/synchronization/blocking_counter.cc
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/synchronization/internal/create_thread_identity.cc
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/synchronization/internal/futex_waiter.cc
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/synchronization/internal/per_thread_sem.cc
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/synchronization/internal/pthread_waiter.cc
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/synchronization/internal/sem_waiter.cc
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/synchronization/internal/stdcpp_waiter.cc
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/synchronization/internal/waiter_base.cc
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/synchronization/internal/win32_waiter.cc
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/synchronization/mutex.cc
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/synchronization/notification.cc
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
    add_library(cronet_aml_third_party_abseil_cpp_absl_synchronization_synchronization SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_abseil_cpp_absl_synchronization_synchronization
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_abseil_cpp_absl_synchronization_synchronization PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_abseil_cpp_absl_synchronization_synchronization
            cronet_aml_third_party_abseil_cpp_absl_base_base
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_malloc_internal
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait
            cronet_aml_third_party_abseil_cpp_absl_base_throw_delegate
            cronet_aml_third_party_abseil_cpp_absl_debugging_debugging_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_demangle_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_stacktrace
            cronet_aml_third_party_abseil_cpp_absl_debugging_symbolize
            cronet_aml_third_party_abseil_cpp_absl_numeric_int128
            cronet_aml_third_party_abseil_cpp_absl_strings_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_strings
            cronet_aml_third_party_abseil_cpp_absl_synchronization_graphcycles_internal
            cronet_aml_third_party_abseil_cpp_absl_synchronization_kernel_timeout_internal
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_civil_time
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_time_zone
            cronet_aml_third_party_abseil_cpp_absl_time_time
            cronet_aml_third_party_abseil_cpp_absl_types_bad_optional_access
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_abseil_cpp_absl_synchronization_synchronization ####################

#################### cronet_aml_third_party_abseil_cpp_absl_random_internal_platform ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/random/internal/randen_round_keys.cc
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
    add_library(cronet_aml_third_party_abseil_cpp_absl_random_internal_platform SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_abseil_cpp_absl_random_internal_platform
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_chromeos_buildflags/android_${OUT_ARCH_CPU}/gen
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_abseil_cpp_absl_random_internal_platform PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_abseil_cpp_absl_random_internal_platform
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_abseil_cpp_absl_random_internal_platform ####################

#################### cronet_aml_buildtools_third_party_libc___libc__ ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/src/algorithm.cpp
        ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/src/any.cpp
        ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/src/atomic.cpp
        ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/src/barrier.cpp
        ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/src/bind.cpp
        ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/src/charconv.cpp
        ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/src/chrono.cpp
        ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/src/condition_variable.cpp
        ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/src/condition_variable_destructor.cpp
        ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/src/exception.cpp
        ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/src/format.cpp
        ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/src/functional.cpp
        ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/src/future.cpp
        ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/src/hash.cpp
        ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/src/ios.cpp
        ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/src/ios.instantiations.cpp
        ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/src/iostream.cpp
        ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/src/legacy_pointer_safety.cpp
        ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/src/locale.cpp
        ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/src/memory.cpp
        ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/src/mutex.cpp
        ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/src/mutex_destructor.cpp
        ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/src/new.cpp
        ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/src/optional.cpp
        ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/src/random.cpp
        ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/src/random_shuffle.cpp
        ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/src/regex.cpp
        ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/src/ryu/d2fixed.cpp
        ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/src/ryu/d2s.cpp
        ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/src/ryu/f2s.cpp
        ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/src/shared_mutex.cpp
        ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/src/stdexcept.cpp
        ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/src/string.cpp
        ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/src/strstream.cpp
        ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/src/system_error.cpp
        ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/src/thread.cpp
        ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/src/typeinfo.cpp
        ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/src/utility.cpp
        ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/src/valarray.cpp
        ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/src/variant.cpp
        ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/src/vector.cpp
        ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/src/verbose_abort.cpp
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
    add_library(cronet_aml_buildtools_third_party_libc___libc__ SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_buildtools_third_party_libc___libc__
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/src
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_buildtools_third_party_libc___libc__ PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_buildtools_third_party_libc___libc__
            cronet_aml_buildtools_third_party_libc__abi_libc__abi
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_buildtools_third_party_libc___libc__ ####################

#################### cronet_aml_third_party_icu_icui18n ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/alphaindex.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/anytrans.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/astro.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/basictz.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/bocsu.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/brktrans.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/buddhcal.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/calendar.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/casetrn.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/cecal.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/chnsecal.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/choicfmt.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/coleitr.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/coll.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/collation.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/collationbuilder.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/collationcompare.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/collationdata.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/collationdatabuilder.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/collationdatareader.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/collationdatawriter.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/collationfastlatin.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/collationfastlatinbuilder.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/collationfcd.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/collationiterator.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/collationkeys.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/collationroot.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/collationrootelements.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/collationruleparser.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/collationsets.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/collationsettings.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/collationtailoring.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/collationweights.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/compactdecimalformat.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/coptccal.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/cpdtrans.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/csdetect.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/csmatch.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/csr2022.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/csrecog.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/csrmbcs.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/csrsbcs.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/csrucode.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/csrutf8.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/curramt.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/currfmt.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/currpinf.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/currunit.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/dangical.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/datefmt.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/dayperiodrules.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/dcfmtsym.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/decContext.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/decNumber.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/decimfmt.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/displayoptions.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/double-conversion-bignum-dtoa.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/double-conversion-bignum.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/double-conversion-cached-powers.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/double-conversion-double-to-string.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/double-conversion-fast-dtoa.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/double-conversion-string-to-double.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/double-conversion-strtod.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/dtfmtsym.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/dtitvfmt.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/dtitvinf.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/dtptngen.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/dtrule.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/erarules.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/esctrn.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/ethpccal.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/fmtable.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/fmtable_cnv.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/format.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/formatted_string_builder.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/formattedval_iterimpl.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/formattedval_sbimpl.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/formattedvalue.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/fphdlimp.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/fpositer.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/funcrepl.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/gender.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/gregocal.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/gregoimp.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/hebrwcal.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/indiancal.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/inputext.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/islamcal.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/japancal.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/listformatter.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/measfmt.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/measunit.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/measunit_extra.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/measure.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/msgfmt.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/name2uni.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/nfrs.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/nfrule.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/nfsubs.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/nortrans.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/nultrans.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/number_affixutils.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/number_asformat.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/number_capi.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/number_compact.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/number_currencysymbols.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/number_decimalquantity.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/number_decimfmtprops.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/number_fluent.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/number_formatimpl.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/number_grouping.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/number_integerwidth.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/number_longnames.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/number_mapper.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/number_modifiers.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/number_multiplier.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/number_notation.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/number_output.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/number_padding.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/number_patternmodifier.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/number_patternstring.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/number_rounding.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/number_scientific.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/number_skeletons.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/number_symbolswrapper.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/number_usageprefs.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/number_utils.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/numfmt.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/numparse_affixes.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/numparse_compositions.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/numparse_currency.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/numparse_decimal.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/numparse_impl.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/numparse_parsednumber.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/numparse_scientific.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/numparse_symbols.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/numparse_validators.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/numrange_capi.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/numrange_fluent.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/numrange_impl.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/numsys.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/olsontz.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/persncal.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/pluralranges.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/plurfmt.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/plurrule.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/quant.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/quantityformatter.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/rbnf.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/rbt.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/rbt_data.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/rbt_pars.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/rbt_rule.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/rbt_set.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/rbtz.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/regexcmp.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/regeximp.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/regexst.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/regextxt.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/region.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/reldatefmt.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/reldtfmt.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/rematch.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/remtrans.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/repattrn.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/rulebasedcollator.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/scientificnumberformatter.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/scriptset.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/search.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/selfmt.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/sharedbreakiterator.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/simpletz.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/smpdtfmt.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/smpdtfst.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/sortkey.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/standardplural.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/string_segment.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/strmatch.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/strrepl.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/stsearch.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/taiwncal.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/timezone.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/titletrn.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/tmunit.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/tmutamt.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/tmutfmt.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/tolowtrn.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/toupptrn.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/translit.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/transreg.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/tridpars.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/tzfmt.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/tzgnames.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/tznames.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/tznames_impl.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/tzrule.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/tztrans.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/ucal.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/ucln_in.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/ucol.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/ucol_res.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/ucol_sit.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/ucoleitr.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/ucsdet.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/udat.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/udateintervalformat.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/udatpg.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/ufieldpositer.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/uitercollationiterator.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/ulistformatter.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/ulocdata.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/umsg.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/unesctrn.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/uni2name.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/units_complexconverter.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/units_converter.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/units_data.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/units_router.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/unum.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/unumsys.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/upluralrules.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/uregex.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/uregexc.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/uregion.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/usearch.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/uspoof.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/uspoof_build.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/uspoof_conf.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/uspoof_impl.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/utf16collationiterator.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/utf8collationiterator.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/utmscale.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/utrans.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/vtzone.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/vzone.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/windtfmt.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/winnmfmt.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/wintzimpl.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/zonemeta.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/zrule.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n/ztrans.cpp
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
    add_library(cronet_aml_third_party_icu_icui18n SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_icu_icui18n
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common
            ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_icu_icui18n PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_icu_icui18n
            cronet_aml_third_party_icu_icuuc_private
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_icu_icui18n ####################

#################### cronet_aml_third_party_brotli_common ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/brotli/common/constants.c
        ${ANDROID_ROOT}/external/cronet/third_party/brotli/common/context.c
        ${ANDROID_ROOT}/external/cronet/third_party/brotli/common/dictionary.c
        ${ANDROID_ROOT}/external/cronet/third_party/brotli/common/platform.c
        ${ANDROID_ROOT}/external/cronet/third_party/brotli/common/shared_dictionary.c
        ${ANDROID_ROOT}/external/cronet/third_party/brotli/common/transform.c
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
    add_library(cronet_aml_third_party_brotli_common SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_brotli_common
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/brotli/include
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_brotli_common PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_brotli_common
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_brotli_common ####################

#################### cronet_aml_third_party_abseil_cpp_absl_debugging_symbolize ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/debugging/symbolize.cc
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
    add_library(cronet_aml_third_party_abseil_cpp_absl_debugging_symbolize SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_abseil_cpp_absl_debugging_symbolize
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_abseil_cpp_absl_debugging_symbolize PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_abseil_cpp_absl_debugging_symbolize
            cronet_aml_third_party_abseil_cpp_absl_base_base
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_malloc_internal
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait
            cronet_aml_third_party_abseil_cpp_absl_base_throw_delegate
            cronet_aml_third_party_abseil_cpp_absl_debugging_debugging_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_demangle_internal
            cronet_aml_third_party_abseil_cpp_absl_numeric_int128
            cronet_aml_third_party_abseil_cpp_absl_strings_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_strings
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_abseil_cpp_absl_debugging_symbolize ####################

#################### cronet_aml_components_metrics_library_support ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_third_party_metrics_proto_metrics_proto_gen/android_${OUT_ARCH_CPU}/gen/external/cronet/third_party/metrics_proto/call_stack_profile.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_third_party_metrics_proto_metrics_proto_gen/android_${OUT_ARCH_CPU}/gen/external/cronet/third_party/metrics_proto/cast_logs.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_third_party_metrics_proto_metrics_proto_gen/android_${OUT_ARCH_CPU}/gen/external/cronet/third_party/metrics_proto/chrome_os_app_list_launch_event.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_third_party_metrics_proto_metrics_proto_gen/android_${OUT_ARCH_CPU}/gen/external/cronet/third_party/metrics_proto/chrome_user_metrics_extension.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_third_party_metrics_proto_metrics_proto_gen/android_${OUT_ARCH_CPU}/gen/external/cronet/third_party/metrics_proto/custom_tab_session.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_third_party_metrics_proto_metrics_proto_gen/android_${OUT_ARCH_CPU}/gen/external/cronet/third_party/metrics_proto/execution_context.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_third_party_metrics_proto_metrics_proto_gen/android_${OUT_ARCH_CPU}/gen/external/cronet/third_party/metrics_proto/extension_install.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_third_party_metrics_proto_metrics_proto_gen/android_${OUT_ARCH_CPU}/gen/external/cronet/third_party/metrics_proto/histogram_event.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_third_party_metrics_proto_metrics_proto_gen/android_${OUT_ARCH_CPU}/gen/external/cronet/third_party/metrics_proto/omnibox_event.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_third_party_metrics_proto_metrics_proto_gen/android_${OUT_ARCH_CPU}/gen/external/cronet/third_party/metrics_proto/omnibox_focus_type.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_third_party_metrics_proto_metrics_proto_gen/android_${OUT_ARCH_CPU}/gen/external/cronet/third_party/metrics_proto/omnibox_input_type.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_third_party_metrics_proto_metrics_proto_gen/android_${OUT_ARCH_CPU}/gen/external/cronet/third_party/metrics_proto/perf_data.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_third_party_metrics_proto_metrics_proto_gen/android_${OUT_ARCH_CPU}/gen/external/cronet/third_party/metrics_proto/perf_stat.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_third_party_metrics_proto_metrics_proto_gen/android_${OUT_ARCH_CPU}/gen/external/cronet/third_party/metrics_proto/printer_event.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_third_party_metrics_proto_metrics_proto_gen/android_${OUT_ARCH_CPU}/gen/external/cronet/third_party/metrics_proto/reporting_info.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_third_party_metrics_proto_metrics_proto_gen/android_${OUT_ARCH_CPU}/gen/external/cronet/third_party/metrics_proto/sampled_profile.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_third_party_metrics_proto_metrics_proto_gen/android_${OUT_ARCH_CPU}/gen/external/cronet/third_party/metrics_proto/structured_data.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_third_party_metrics_proto_metrics_proto_gen/android_${OUT_ARCH_CPU}/gen/external/cronet/third_party/metrics_proto/system_profile.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_third_party_metrics_proto_metrics_proto_gen/android_${OUT_ARCH_CPU}/gen/external/cronet/third_party/metrics_proto/trace_log.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_third_party_metrics_proto_metrics_proto_gen/android_${OUT_ARCH_CPU}/gen/external/cronet/third_party/metrics_proto/translate_event.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_third_party_metrics_proto_metrics_proto_gen/android_${OUT_ARCH_CPU}/gen/external/cronet/third_party/metrics_proto/ukm/aggregate.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_third_party_metrics_proto_metrics_proto_gen/android_${OUT_ARCH_CPU}/gen/external/cronet/third_party/metrics_proto/ukm/entry.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_third_party_metrics_proto_metrics_proto_gen/android_${OUT_ARCH_CPU}/gen/external/cronet/third_party/metrics_proto/ukm/report.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_third_party_metrics_proto_metrics_proto_gen/android_${OUT_ARCH_CPU}/gen/external/cronet/third_party/metrics_proto/ukm/source.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_third_party_metrics_proto_metrics_proto_gen/android_${OUT_ARCH_CPU}/gen/external/cronet/third_party/metrics_proto/user_action_event.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_third_party_metrics_proto_metrics_proto_gen/android_${OUT_ARCH_CPU}/gen/external/cronet/third_party/metrics_proto/user_demographics.pb.cc
        ${ANDROID_ROOT}/external/cronet/components/metrics/histogram_encoder.cc
        ${ANDROID_ROOT}/external/cronet/components/metrics/library_support/histogram_manager.cc
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
    add_library(cronet_aml_components_metrics_library_support SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_components_metrics_library_support
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
            ${ANDROID_ROOT}/external/cronet/third_party/boringssl/src/include
            ${ANDROID_ROOT}/external/cronet/third_party/protobuf/src
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_chromecast_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_chromeos_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_debugging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_logging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_partition_alloc_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_anchor_functions_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_android_runtime_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_android_runtime_unchecked_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_base_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_build_date/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_cfi_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_clang_profiling_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_debugging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_feature_list_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_ios_cronet_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_logging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_message_pump_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_orderfile_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_power_monitor_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_profiler_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_rust_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_sanitizer_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_synchronization_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_tracing_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_blink_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_branding_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_chromecast_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_chromeos_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_config_compiler_compiler_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_third_party_metrics_proto_metrics_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_third_party_metrics_proto_metrics_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/protos
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_third_party_metrics_proto_metrics_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/third_party/metrics_proto
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_components_metrics_library_support PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_components_metrics_library_support
            cronet_aml_base_allocator_partition_allocator_partition_alloc
            cronet_aml_base_base
            cronet_aml_base_base_static
            cronet_aml_base_nodebug_assertion
            cronet_aml_base_third_party_double_conversion_double_conversion
            cronet_aml_base_third_party_dynamic_annotations_dynamic_annotations
            cronet_aml_third_party_abseil_cpp_absl_base_base
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_malloc_internal
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait
            cronet_aml_third_party_abseil_cpp_absl_base_strerror
            cronet_aml_third_party_abseil_cpp_absl_base_throw_delegate
            cronet_aml_third_party_abseil_cpp_absl_container_hashtablez_sampler
            cronet_aml_third_party_abseil_cpp_absl_container_raw_hash_set
            cronet_aml_third_party_abseil_cpp_absl_crc_cpu_detect
            cronet_aml_third_party_abseil_cpp_absl_crc_crc32c
            cronet_aml_third_party_abseil_cpp_absl_crc_crc_cord_state
            cronet_aml_third_party_abseil_cpp_absl_crc_crc_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_debugging_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_demangle_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_examine_stack
            cronet_aml_third_party_abseil_cpp_absl_debugging_failure_signal_handler
            cronet_aml_third_party_abseil_cpp_absl_debugging_stacktrace
            cronet_aml_third_party_abseil_cpp_absl_debugging_symbolize
            cronet_aml_third_party_abseil_cpp_absl_hash_city
            cronet_aml_third_party_abseil_cpp_absl_hash_hash
            cronet_aml_third_party_abseil_cpp_absl_hash_low_level_hash
            cronet_aml_third_party_abseil_cpp_absl_log_die_if_null
            cronet_aml_third_party_abseil_cpp_absl_log_globals
            cronet_aml_third_party_abseil_cpp_absl_log_internal_check_op
            cronet_aml_third_party_abseil_cpp_absl_log_internal_conditions
            cronet_aml_third_party_abseil_cpp_absl_log_internal_format
            cronet_aml_third_party_abseil_cpp_absl_log_internal_globals
            cronet_aml_third_party_abseil_cpp_absl_log_internal_log_message
            cronet_aml_third_party_abseil_cpp_absl_log_internal_log_sink_set
            cronet_aml_third_party_abseil_cpp_absl_log_internal_nullguard
            cronet_aml_third_party_abseil_cpp_absl_log_internal_proto
            cronet_aml_third_party_abseil_cpp_absl_log_log_entry
            cronet_aml_third_party_abseil_cpp_absl_log_log_sink
            cronet_aml_third_party_abseil_cpp_absl_numeric_int128
            cronet_aml_third_party_abseil_cpp_absl_profiling_exponential_biased
            cronet_aml_third_party_abseil_cpp_absl_random_distributions
            cronet_aml_third_party_abseil_cpp_absl_random_internal_platform
            cronet_aml_third_party_abseil_cpp_absl_random_internal_pool_urbg
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_hwaes
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_hwaes_impl
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_slow
            cronet_aml_third_party_abseil_cpp_absl_random_internal_seed_material
            cronet_aml_third_party_abseil_cpp_absl_random_seed_gen_exception
            cronet_aml_third_party_abseil_cpp_absl_random_seed_sequences
            cronet_aml_third_party_abseil_cpp_absl_status_status
            cronet_aml_third_party_abseil_cpp_absl_status_statusor
            cronet_aml_third_party_abseil_cpp_absl_strings_cord
            cronet_aml_third_party_abseil_cpp_absl_strings_cord_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_functions
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_handle
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_info
            cronet_aml_third_party_abseil_cpp_absl_strings_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_str_format_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_strings
            cronet_aml_third_party_abseil_cpp_absl_synchronization_graphcycles_internal
            cronet_aml_third_party_abseil_cpp_absl_synchronization_kernel_timeout_internal
            cronet_aml_third_party_abseil_cpp_absl_synchronization_synchronization
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_civil_time
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_time_zone
            cronet_aml_third_party_abseil_cpp_absl_time_time
            cronet_aml_third_party_abseil_cpp_absl_types_bad_optional_access
            cronet_aml_third_party_abseil_cpp_absl_types_bad_variant_access
            cronet_aml_third_party_android_ndk_cpu_features
            cronet_aml_third_party_ashmem_ashmem
            cronet_aml_third_party_icu_icui18n
            cronet_aml_third_party_icu_icuuc_private
            cronet_aml_third_party_libevent_libevent
            cronet_aml_third_party_modp_b64_modp_b64
            cronet_aml_third_party_protobuf_protobuf_lite
            jni_headers
            libandroid
            libc
            libclang_rt.builtins
            libcrypto
            libdl
            liblog
            libm
            libmediandk
            libprotobuf-cpp-lite
            libssl
    )
endif ()
#################### cronet_aml_components_metrics_library_support ####################

#################### cronet_aml_net_dns_dns ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/net/dns/address_info.cc
        ${ANDROID_ROOT}/external/cronet/net/dns/address_sorter_posix.cc
        ${ANDROID_ROOT}/external/cronet/net/dns/context_host_resolver.cc
        ${ANDROID_ROOT}/external/cronet/net/dns/dns_alias_utility.cc
        ${ANDROID_ROOT}/external/cronet/net/dns/dns_client.cc
        ${ANDROID_ROOT}/external/cronet/net/dns/dns_config.cc
        ${ANDROID_ROOT}/external/cronet/net/dns/dns_config_service.cc
        ${ANDROID_ROOT}/external/cronet/net/dns/dns_config_service_android.cc
        ${ANDROID_ROOT}/external/cronet/net/dns/dns_hosts.cc
        ${ANDROID_ROOT}/external/cronet/net/dns/dns_names_util.cc
        ${ANDROID_ROOT}/external/cronet/net/dns/dns_query.cc
        ${ANDROID_ROOT}/external/cronet/net/dns/dns_reloader.cc
        ${ANDROID_ROOT}/external/cronet/net/dns/dns_response.cc
        ${ANDROID_ROOT}/external/cronet/net/dns/dns_response_result_extractor.cc
        ${ANDROID_ROOT}/external/cronet/net/dns/dns_server_iterator.cc
        ${ANDROID_ROOT}/external/cronet/net/dns/dns_session.cc
        ${ANDROID_ROOT}/external/cronet/net/dns/dns_transaction.cc
        ${ANDROID_ROOT}/external/cronet/net/dns/dns_udp_tracker.cc
        ${ANDROID_ROOT}/external/cronet/net/dns/dns_util.cc
        ${ANDROID_ROOT}/external/cronet/net/dns/host_cache.cc
        ${ANDROID_ROOT}/external/cronet/net/dns/host_resolver.cc
        ${ANDROID_ROOT}/external/cronet/net/dns/host_resolver_internal_result.cc
        ${ANDROID_ROOT}/external/cronet/net/dns/host_resolver_manager.cc
        ${ANDROID_ROOT}/external/cronet/net/dns/host_resolver_mdns_listener_impl.cc
        ${ANDROID_ROOT}/external/cronet/net/dns/host_resolver_mdns_task.cc
        ${ANDROID_ROOT}/external/cronet/net/dns/host_resolver_nat64_task.cc
        ${ANDROID_ROOT}/external/cronet/net/dns/host_resolver_proc.cc
        ${ANDROID_ROOT}/external/cronet/net/dns/host_resolver_system_task.cc
        ${ANDROID_ROOT}/external/cronet/net/dns/https_record_rdata.cc
        ${ANDROID_ROOT}/external/cronet/net/dns/httpssvc_metrics.cc
        ${ANDROID_ROOT}/external/cronet/net/dns/mapped_host_resolver.cc
        ${ANDROID_ROOT}/external/cronet/net/dns/nsswitch_reader.cc
        ${ANDROID_ROOT}/external/cronet/net/dns/opt_record_rdata.cc
        ${ANDROID_ROOT}/external/cronet/net/dns/record_parsed.cc
        ${ANDROID_ROOT}/external/cronet/net/dns/record_rdata.cc
        ${ANDROID_ROOT}/external/cronet/net/dns/resolve_context.cc
        ${ANDROID_ROOT}/external/cronet/net/dns/serial_worker.cc
        ${ANDROID_ROOT}/external/cronet/net/dns/system_dns_config_change_notifier.cc
        ${ANDROID_ROOT}/external/cronet/net/dns/test_dns_config_service.cc
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
    add_library(cronet_aml_net_dns_dns SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_net_dns_dns
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/overrides
            ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src
            ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/common/platform/default
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
            ${ANDROID_ROOT}/external/cronet/third_party/boringssl/src/include
            ${ANDROID_ROOT}/external/cronet/third_party/brotli/include
            ${ANDROID_ROOT}/external/cronet/third_party/protobuf/src
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_chromecast_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_chromeos_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_debugging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_logging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_partition_alloc_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_anchor_functions_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_android_runtime_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_android_runtime_unchecked_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_base_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_build_date/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_cfi_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_clang_profiling_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_debugging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_feature_list_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_ios_cronet_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_logging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_message_pump_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_orderfile_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_power_monitor_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_profiler_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_rust_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_sanitizer_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_synchronization_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_tracing_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_blink_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_branding_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_chromecast_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_chromeos_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_config_compiler_compiler_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_crypto_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_base_registry_controlled_domains_registry_controlled_domains/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_cronet_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_isolation_info_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_isolation_info_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/net/base
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_isolation_info_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/protos
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_net_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_url_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_net_nqe_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_net_nqe_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/net/nqe/proto
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_net_nqe_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/protos
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_third_party_quiche_net_quic_test_tools_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_third_party_quiche_net_quic_test_tools_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/net/third_party/quiche/src
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_third_party_quiche_net_quic_test_tools_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/net/third_party/quiche/src/quiche/quic/test_tools
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_third_party_quiche_net_quic_test_tools_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/protos
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_third_party_quiche_net_quic_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_third_party_quiche_net_quic_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/net/third_party/quiche/src
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_third_party_quiche_net_quic_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/protos
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_url_url_jni_headers/android_${OUT_ARCH_CPU}/gen
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_net_dns_dns PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_net_dns_dns
            cronet_aml_base_allocator_partition_allocator_partition_alloc
            cronet_aml_base_base
            cronet_aml_base_base_static
            cronet_aml_base_nodebug_assertion
            cronet_aml_base_third_party_double_conversion_double_conversion
            cronet_aml_base_third_party_dynamic_annotations_dynamic_annotations
            cronet_aml_crypto_crypto
            cronet_aml_net_dns_public_public
            cronet_aml_net_net_deps
            cronet_aml_net_net_public_deps
            cronet_aml_net_preload_decoder
            cronet_aml_net_third_party_quiche_quiche
            cronet_aml_net_traffic_annotation_traffic_annotation
            cronet_aml_net_uri_template
            cronet_aml_third_party_abseil_cpp_absl_base_base
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_malloc_internal
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait
            cronet_aml_third_party_abseil_cpp_absl_base_strerror
            cronet_aml_third_party_abseil_cpp_absl_base_throw_delegate
            cronet_aml_third_party_abseil_cpp_absl_container_hashtablez_sampler
            cronet_aml_third_party_abseil_cpp_absl_container_raw_hash_set
            cronet_aml_third_party_abseil_cpp_absl_crc_cpu_detect
            cronet_aml_third_party_abseil_cpp_absl_crc_crc32c
            cronet_aml_third_party_abseil_cpp_absl_crc_crc_cord_state
            cronet_aml_third_party_abseil_cpp_absl_crc_crc_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_debugging_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_demangle_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_examine_stack
            cronet_aml_third_party_abseil_cpp_absl_debugging_failure_signal_handler
            cronet_aml_third_party_abseil_cpp_absl_debugging_stacktrace
            cronet_aml_third_party_abseil_cpp_absl_debugging_symbolize
            cronet_aml_third_party_abseil_cpp_absl_hash_city
            cronet_aml_third_party_abseil_cpp_absl_hash_hash
            cronet_aml_third_party_abseil_cpp_absl_hash_low_level_hash
            cronet_aml_third_party_abseil_cpp_absl_log_die_if_null
            cronet_aml_third_party_abseil_cpp_absl_log_globals
            cronet_aml_third_party_abseil_cpp_absl_log_internal_check_op
            cronet_aml_third_party_abseil_cpp_absl_log_internal_conditions
            cronet_aml_third_party_abseil_cpp_absl_log_internal_format
            cronet_aml_third_party_abseil_cpp_absl_log_internal_globals
            cronet_aml_third_party_abseil_cpp_absl_log_internal_log_message
            cronet_aml_third_party_abseil_cpp_absl_log_internal_log_sink_set
            cronet_aml_third_party_abseil_cpp_absl_log_internal_nullguard
            cronet_aml_third_party_abseil_cpp_absl_log_internal_proto
            cronet_aml_third_party_abseil_cpp_absl_log_log_entry
            cronet_aml_third_party_abseil_cpp_absl_log_log_sink
            cronet_aml_third_party_abseil_cpp_absl_numeric_int128
            cronet_aml_third_party_abseil_cpp_absl_profiling_exponential_biased
            cronet_aml_third_party_abseil_cpp_absl_random_distributions
            cronet_aml_third_party_abseil_cpp_absl_random_internal_platform
            cronet_aml_third_party_abseil_cpp_absl_random_internal_pool_urbg
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_hwaes
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_hwaes_impl
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_slow
            cronet_aml_third_party_abseil_cpp_absl_random_internal_seed_material
            cronet_aml_third_party_abseil_cpp_absl_random_seed_gen_exception
            cronet_aml_third_party_abseil_cpp_absl_random_seed_sequences
            cronet_aml_third_party_abseil_cpp_absl_status_status
            cronet_aml_third_party_abseil_cpp_absl_status_statusor
            cronet_aml_third_party_abseil_cpp_absl_strings_cord
            cronet_aml_third_party_abseil_cpp_absl_strings_cord_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_functions
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_handle
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_info
            cronet_aml_third_party_abseil_cpp_absl_strings_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_str_format_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_strings
            cronet_aml_third_party_abseil_cpp_absl_synchronization_graphcycles_internal
            cronet_aml_third_party_abseil_cpp_absl_synchronization_kernel_timeout_internal
            cronet_aml_third_party_abseil_cpp_absl_synchronization_synchronization
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_civil_time
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_time_zone
            cronet_aml_third_party_abseil_cpp_absl_time_time
            cronet_aml_third_party_abseil_cpp_absl_types_bad_optional_access
            cronet_aml_third_party_abseil_cpp_absl_types_bad_variant_access
            cronet_aml_third_party_android_ndk_cpu_features
            cronet_aml_third_party_ashmem_ashmem
            cronet_aml_third_party_brotli_common
            cronet_aml_third_party_brotli_dec
            cronet_aml_third_party_icu_icui18n
            cronet_aml_third_party_icu_icuuc_private
            cronet_aml_third_party_libevent_libevent
            cronet_aml_third_party_modp_b64_modp_b64
            cronet_aml_third_party_protobuf_protobuf_lite
            cronet_aml_url_url
            jni_headers
            libandroid
            libc
            libclang_rt.builtins
            libcrypto
            libdl
            liblog
            libm
            libmediandk
            libssl
            libz
    )
endif ()
#################### cronet_aml_net_dns_dns ####################

#################### cronet_aml_third_party_abseil_cpp_absl_log_internal_globals ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/log/internal/globals.cc
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
    add_library(cronet_aml_third_party_abseil_cpp_absl_log_internal_globals SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_abseil_cpp_absl_log_internal_globals
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_abseil_cpp_absl_log_internal_globals PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_abseil_cpp_absl_log_internal_globals
            cronet_aml_third_party_abseil_cpp_absl_base_base
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait
            cronet_aml_third_party_abseil_cpp_absl_base_throw_delegate
            cronet_aml_third_party_abseil_cpp_absl_numeric_int128
            cronet_aml_third_party_abseil_cpp_absl_strings_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_strings
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_civil_time
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_time_zone
            cronet_aml_third_party_abseil_cpp_absl_time_time
            cronet_aml_third_party_abseil_cpp_absl_types_bad_optional_access
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_abseil_cpp_absl_log_internal_globals ####################

#################### cronet_aml_net_net ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/net/android/android_http_util.cc
        ${ANDROID_ROOT}/external/cronet/net/android/cert_verify_result_android.cc
        ${ANDROID_ROOT}/external/cronet/net/android/gurl_utils.cc
        ${ANDROID_ROOT}/external/cronet/net/android/http_auth_negotiate_android.cc
        ${ANDROID_ROOT}/external/cronet/net/android/keystore.cc
        ${ANDROID_ROOT}/external/cronet/net/android/network_change_notifier_android.cc
        ${ANDROID_ROOT}/external/cronet/net/android/network_change_notifier_delegate_android.cc
        ${ANDROID_ROOT}/external/cronet/net/android/network_change_notifier_factory_android.cc
        ${ANDROID_ROOT}/external/cronet/net/android/network_library.cc
        ${ANDROID_ROOT}/external/cronet/net/android/traffic_stats.cc
        ${ANDROID_ROOT}/external/cronet/net/base/address_family.cc
        ${ANDROID_ROOT}/external/cronet/net/base/address_list.cc
        ${ANDROID_ROOT}/external/cronet/net/base/address_map_linux.cc
        ${ANDROID_ROOT}/external/cronet/net/base/address_tracker_linux.cc
        ${ANDROID_ROOT}/external/cronet/net/base/auth.cc
        ${ANDROID_ROOT}/external/cronet/net/base/backoff_entry.cc
        ${ANDROID_ROOT}/external/cronet/net/base/backoff_entry_serializer.cc
        ${ANDROID_ROOT}/external/cronet/net/base/cache_metrics.cc
        ${ANDROID_ROOT}/external/cronet/net/base/chunked_upload_data_stream.cc
        ${ANDROID_ROOT}/external/cronet/net/base/connection_endpoint_metadata.cc
        ${ANDROID_ROOT}/external/cronet/net/base/data_url.cc
        ${ANDROID_ROOT}/external/cronet/net/base/datagram_buffer.cc
        ${ANDROID_ROOT}/external/cronet/net/base/elements_upload_data_stream.cc
        ${ANDROID_ROOT}/external/cronet/net/base/features.cc
        ${ANDROID_ROOT}/external/cronet/net/base/file_stream.cc
        ${ANDROID_ROOT}/external/cronet/net/base/file_stream_context.cc
        ${ANDROID_ROOT}/external/cronet/net/base/file_stream_context_posix.cc
        ${ANDROID_ROOT}/external/cronet/net/base/filename_util.cc
        ${ANDROID_ROOT}/external/cronet/net/base/filename_util_internal.cc
        ${ANDROID_ROOT}/external/cronet/net/base/hash_value.cc
        ${ANDROID_ROOT}/external/cronet/net/base/hex_utils.cc
        ${ANDROID_ROOT}/external/cronet/net/base/host_mapping_rules.cc
        ${ANDROID_ROOT}/external/cronet/net/base/host_port_pair.cc
        ${ANDROID_ROOT}/external/cronet/net/base/io_buffer.cc
        ${ANDROID_ROOT}/external/cronet/net/base/ip_address.cc
        ${ANDROID_ROOT}/external/cronet/net/base/ip_endpoint.cc
        ${ANDROID_ROOT}/external/cronet/net/base/isolation_info.cc
        ${ANDROID_ROOT}/external/cronet/net/base/load_timing_info.cc
        ${ANDROID_ROOT}/external/cronet/net/base/logging_network_change_observer.cc
        ${ANDROID_ROOT}/external/cronet/net/base/lookup_string_in_fixed_set.cc
        ${ANDROID_ROOT}/external/cronet/net/base/mime_sniffer.cc
        ${ANDROID_ROOT}/external/cronet/net/base/mime_util.cc
        ${ANDROID_ROOT}/external/cronet/net/base/net_errors.cc
        ${ANDROID_ROOT}/external/cronet/net/base/net_errors_posix.cc
        ${ANDROID_ROOT}/external/cronet/net/base/net_module.cc
        ${ANDROID_ROOT}/external/cronet/net/base/net_string_util_icu_alternatives_android.cc
        ${ANDROID_ROOT}/external/cronet/net/base/network_activity_monitor.cc
        ${ANDROID_ROOT}/external/cronet/net/base/network_anonymization_key.cc
        ${ANDROID_ROOT}/external/cronet/net/base/network_change_notifier.cc
        ${ANDROID_ROOT}/external/cronet/net/base/network_change_notifier_passive.cc
        ${ANDROID_ROOT}/external/cronet/net/base/network_delegate.cc
        ${ANDROID_ROOT}/external/cronet/net/base/network_delegate_impl.cc
        ${ANDROID_ROOT}/external/cronet/net/base/network_interfaces.cc
        ${ANDROID_ROOT}/external/cronet/net/base/network_interfaces_getifaddrs.cc
        ${ANDROID_ROOT}/external/cronet/net/base/network_interfaces_getifaddrs_android.cc
        ${ANDROID_ROOT}/external/cronet/net/base/network_interfaces_linux.cc
        ${ANDROID_ROOT}/external/cronet/net/base/network_interfaces_posix.cc
        ${ANDROID_ROOT}/external/cronet/net/base/network_isolation_key.cc
        ${ANDROID_ROOT}/external/cronet/net/base/parse_number.cc
        ${ANDROID_ROOT}/external/cronet/net/base/platform_mime_util_linux.cc
        ${ANDROID_ROOT}/external/cronet/net/base/port_util.cc
        ${ANDROID_ROOT}/external/cronet/net/base/prioritized_dispatcher.cc
        ${ANDROID_ROOT}/external/cronet/net/base/prioritized_task_runner.cc
        ${ANDROID_ROOT}/external/cronet/net/base/privacy_mode.cc
        ${ANDROID_ROOT}/external/cronet/net/base/proxy_server.cc
        ${ANDROID_ROOT}/external/cronet/net/base/proxy_string_util.cc
        ${ANDROID_ROOT}/external/cronet/net/base/registry_controlled_domains/registry_controlled_domain.cc
        ${ANDROID_ROOT}/external/cronet/net/base/request_priority.cc
        ${ANDROID_ROOT}/external/cronet/net/base/scheme_host_port_matcher.cc
        ${ANDROID_ROOT}/external/cronet/net/base/scheme_host_port_matcher_rule.cc
        ${ANDROID_ROOT}/external/cronet/net/base/schemeful_site.cc
        ${ANDROID_ROOT}/external/cronet/net/base/sockaddr_storage.cc
        ${ANDROID_ROOT}/external/cronet/net/base/sockaddr_util_posix.cc
        ${ANDROID_ROOT}/external/cronet/net/base/transport_info.cc
        ${ANDROID_ROOT}/external/cronet/net/base/upload_bytes_element_reader.cc
        ${ANDROID_ROOT}/external/cronet/net/base/upload_data_stream.cc
        ${ANDROID_ROOT}/external/cronet/net/base/upload_element_reader.cc
        ${ANDROID_ROOT}/external/cronet/net/base/upload_file_element_reader.cc
        ${ANDROID_ROOT}/external/cronet/net/base/url_search_params.cc
        ${ANDROID_ROOT}/external/cronet/net/base/url_util.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/asn1_util.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/caching_cert_verifier.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/cert_and_ct_verifier.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/cert_database.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/cert_status_flags.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/cert_verifier.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/cert_verify_proc.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/cert_verify_proc_android.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/cert_verify_proc_builtin.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/cert_verify_result.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/coalescing_cert_verifier.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/crl_set.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/ct_log_response_parser.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/ct_log_verifier.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/ct_log_verifier_util.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/ct_objects_extractor.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/ct_policy_enforcer.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/ct_sct_to_string.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/ct_serialization.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/ct_signed_certificate_timestamp_log_param.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/do_nothing_ct_verifier.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/ev_root_ca_metadata.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/internal/cert_issuer_source_aia.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/internal/revocation_checker.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/internal/system_trust_store.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/internal/trust_store_features.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/known_roots.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/merkle_audit_proof.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/merkle_consistency_proof.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/merkle_tree_leaf.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/multi_log_ct_verifier.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/multi_threaded_cert_verifier.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/ocsp_verify_result.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/pem.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/pki/cert_error_id.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/pki/cert_error_params.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/pki/cert_errors.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/pki/cert_issuer_source_static.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/pki/certificate_policies.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/pki/common_cert_errors.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/pki/crl.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/pki/extended_key_usage.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/pki/general_names.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/pki/name_constraints.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/pki/ocsp.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/pki/parse_certificate.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/pki/parse_name.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/pki/parsed_certificate.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/pki/path_builder.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/pki/revocation_util.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/pki/signature_algorithm.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/pki/simple_path_builder_delegate.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/pki/string_util.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/pki/trust_store.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/pki/trust_store_collection.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/pki/trust_store_in_memory.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/pki/verify_certificate_chain.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/pki/verify_name_match.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/pki/verify_signed_data.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/sct_status_flags.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/signed_certificate_timestamp.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/signed_certificate_timestamp_and_status.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/signed_tree_head.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/symantec_certs.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/test_root_certs.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/test_root_certs_android.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/trial_comparison_cert_verifier_util.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/x509_cert_types.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/x509_certificate.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/x509_certificate_net_log_param.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/x509_util.cc
        ${ANDROID_ROOT}/external/cronet/net/cert/x509_util_android.cc
        ${ANDROID_ROOT}/external/cronet/net/cert_net/cert_net_fetcher_url_request.cc
        ${ANDROID_ROOT}/external/cronet/net/cookies/canonical_cookie.cc
        ${ANDROID_ROOT}/external/cronet/net/cookies/cookie_access_delegate.cc
        ${ANDROID_ROOT}/external/cronet/net/cookies/cookie_access_result.cc
        ${ANDROID_ROOT}/external/cronet/net/cookies/cookie_change_dispatcher.cc
        ${ANDROID_ROOT}/external/cronet/net/cookies/cookie_constants.cc
        ${ANDROID_ROOT}/external/cronet/net/cookies/cookie_deletion_info.cc
        ${ANDROID_ROOT}/external/cronet/net/cookies/cookie_inclusion_status.cc
        ${ANDROID_ROOT}/external/cronet/net/cookies/cookie_monster.cc
        ${ANDROID_ROOT}/external/cronet/net/cookies/cookie_monster_change_dispatcher.cc
        ${ANDROID_ROOT}/external/cronet/net/cookies/cookie_monster_netlog_params.cc
        ${ANDROID_ROOT}/external/cronet/net/cookies/cookie_options.cc
        ${ANDROID_ROOT}/external/cronet/net/cookies/cookie_partition_key.cc
        ${ANDROID_ROOT}/external/cronet/net/cookies/cookie_partition_key_collection.cc
        ${ANDROID_ROOT}/external/cronet/net/cookies/cookie_store.cc
        ${ANDROID_ROOT}/external/cronet/net/cookies/cookie_util.cc
        ${ANDROID_ROOT}/external/cronet/net/cookies/parsed_cookie.cc
        ${ANDROID_ROOT}/external/cronet/net/cookies/site_for_cookies.cc
        ${ANDROID_ROOT}/external/cronet/net/cookies/static_cookie_policy.cc
        ${ANDROID_ROOT}/external/cronet/net/der/encode_values.cc
        ${ANDROID_ROOT}/external/cronet/net/der/input.cc
        ${ANDROID_ROOT}/external/cronet/net/der/parse_values.cc
        ${ANDROID_ROOT}/external/cronet/net/der/parser.cc
        ${ANDROID_ROOT}/external/cronet/net/der/tag.cc
        ${ANDROID_ROOT}/external/cronet/net/disk_cache/backend_cleanup_tracker.cc
        ${ANDROID_ROOT}/external/cronet/net/disk_cache/blockfile/addr.cc
        ${ANDROID_ROOT}/external/cronet/net/disk_cache/blockfile/backend_impl.cc
        ${ANDROID_ROOT}/external/cronet/net/disk_cache/blockfile/bitmap.cc
        ${ANDROID_ROOT}/external/cronet/net/disk_cache/blockfile/block_files.cc
        ${ANDROID_ROOT}/external/cronet/net/disk_cache/blockfile/disk_format.cc
        ${ANDROID_ROOT}/external/cronet/net/disk_cache/blockfile/entry_impl.cc
        ${ANDROID_ROOT}/external/cronet/net/disk_cache/blockfile/eviction.cc
        ${ANDROID_ROOT}/external/cronet/net/disk_cache/blockfile/file.cc
        ${ANDROID_ROOT}/external/cronet/net/disk_cache/blockfile/file_lock.cc
        ${ANDROID_ROOT}/external/cronet/net/disk_cache/blockfile/file_posix.cc
        ${ANDROID_ROOT}/external/cronet/net/disk_cache/blockfile/in_flight_backend_io.cc
        ${ANDROID_ROOT}/external/cronet/net/disk_cache/blockfile/in_flight_io.cc
        ${ANDROID_ROOT}/external/cronet/net/disk_cache/blockfile/mapped_file.cc
        ${ANDROID_ROOT}/external/cronet/net/disk_cache/blockfile/rankings.cc
        ${ANDROID_ROOT}/external/cronet/net/disk_cache/blockfile/sparse_control.cc
        ${ANDROID_ROOT}/external/cronet/net/disk_cache/blockfile/stats.cc
        ${ANDROID_ROOT}/external/cronet/net/disk_cache/cache_util.cc
        ${ANDROID_ROOT}/external/cronet/net/disk_cache/cache_util_posix.cc
        ${ANDROID_ROOT}/external/cronet/net/disk_cache/disk_cache.cc
        ${ANDROID_ROOT}/external/cronet/net/disk_cache/memory/mem_backend_impl.cc
        ${ANDROID_ROOT}/external/cronet/net/disk_cache/memory/mem_entry_impl.cc
        ${ANDROID_ROOT}/external/cronet/net/disk_cache/net_log_parameters.cc
        ${ANDROID_ROOT}/external/cronet/net/disk_cache/simple/post_doom_waiter.cc
        ${ANDROID_ROOT}/external/cronet/net/disk_cache/simple/simple_backend_impl.cc
        ${ANDROID_ROOT}/external/cronet/net/disk_cache/simple/simple_entry_format.cc
        ${ANDROID_ROOT}/external/cronet/net/disk_cache/simple/simple_entry_impl.cc
        ${ANDROID_ROOT}/external/cronet/net/disk_cache/simple/simple_entry_operation.cc
        ${ANDROID_ROOT}/external/cronet/net/disk_cache/simple/simple_file_enumerator.cc
        ${ANDROID_ROOT}/external/cronet/net/disk_cache/simple/simple_file_tracker.cc
        ${ANDROID_ROOT}/external/cronet/net/disk_cache/simple/simple_index.cc
        ${ANDROID_ROOT}/external/cronet/net/disk_cache/simple/simple_index_file.cc
        ${ANDROID_ROOT}/external/cronet/net/disk_cache/simple/simple_net_log_parameters.cc
        ${ANDROID_ROOT}/external/cronet/net/disk_cache/simple/simple_synchronous_entry.cc
        ${ANDROID_ROOT}/external/cronet/net/disk_cache/simple/simple_util.cc
        ${ANDROID_ROOT}/external/cronet/net/disk_cache/simple/simple_util_posix.cc
        ${ANDROID_ROOT}/external/cronet/net/disk_cache/simple/simple_version_upgrade.cc
        ${ANDROID_ROOT}/external/cronet/net/filter/brotli_source_stream.cc
        ${ANDROID_ROOT}/external/cronet/net/filter/filter_source_stream.cc
        ${ANDROID_ROOT}/external/cronet/net/filter/gzip_header.cc
        ${ANDROID_ROOT}/external/cronet/net/filter/gzip_source_stream.cc
        ${ANDROID_ROOT}/external/cronet/net/filter/source_stream.cc
        ${ANDROID_ROOT}/external/cronet/net/first_party_sets/addition_overlaps_union_find.cc
        ${ANDROID_ROOT}/external/cronet/net/first_party_sets/first_party_set_entry.cc
        ${ANDROID_ROOT}/external/cronet/net/first_party_sets/first_party_set_entry_override.cc
        ${ANDROID_ROOT}/external/cronet/net/first_party_sets/first_party_set_metadata.cc
        ${ANDROID_ROOT}/external/cronet/net/first_party_sets/first_party_sets_cache_filter.cc
        ${ANDROID_ROOT}/external/cronet/net/first_party_sets/first_party_sets_context_config.cc
        ${ANDROID_ROOT}/external/cronet/net/first_party_sets/global_first_party_sets.cc
        ${ANDROID_ROOT}/external/cronet/net/first_party_sets/same_party_context.cc
        ${ANDROID_ROOT}/external/cronet/net/http/alternative_service.cc
        ${ANDROID_ROOT}/external/cronet/net/http/bidirectional_stream.cc
        ${ANDROID_ROOT}/external/cronet/net/http/bidirectional_stream_impl.cc
        ${ANDROID_ROOT}/external/cronet/net/http/bidirectional_stream_request_info.cc
        ${ANDROID_ROOT}/external/cronet/net/http/broken_alternative_services.cc
        ${ANDROID_ROOT}/external/cronet/net/http/http_auth.cc
        ${ANDROID_ROOT}/external/cronet/net/http/http_auth_cache.cc
        ${ANDROID_ROOT}/external/cronet/net/http/http_auth_challenge_tokenizer.cc
        ${ANDROID_ROOT}/external/cronet/net/http/http_auth_controller.cc
        ${ANDROID_ROOT}/external/cronet/net/http/http_auth_filter.cc
        ${ANDROID_ROOT}/external/cronet/net/http/http_auth_handler.cc
        ${ANDROID_ROOT}/external/cronet/net/http/http_auth_handler_basic.cc
        ${ANDROID_ROOT}/external/cronet/net/http/http_auth_handler_digest.cc
        ${ANDROID_ROOT}/external/cronet/net/http/http_auth_handler_factory.cc
        ${ANDROID_ROOT}/external/cronet/net/http/http_auth_handler_negotiate.cc
        ${ANDROID_ROOT}/external/cronet/net/http/http_auth_handler_ntlm.cc
        ${ANDROID_ROOT}/external/cronet/net/http/http_auth_handler_ntlm_portable.cc
        ${ANDROID_ROOT}/external/cronet/net/http/http_auth_multi_round_parse.cc
        ${ANDROID_ROOT}/external/cronet/net/http/http_auth_ntlm_mechanism.cc
        ${ANDROID_ROOT}/external/cronet/net/http/http_auth_preferences.cc
        ${ANDROID_ROOT}/external/cronet/net/http/http_auth_scheme.cc
        ${ANDROID_ROOT}/external/cronet/net/http/http_basic_state.cc
        ${ANDROID_ROOT}/external/cronet/net/http/http_basic_stream.cc
        ${ANDROID_ROOT}/external/cronet/net/http/http_byte_range.cc
        ${ANDROID_ROOT}/external/cronet/net/http/http_cache.cc
        ${ANDROID_ROOT}/external/cronet/net/http/http_cache_lookup_manager.cc
        ${ANDROID_ROOT}/external/cronet/net/http/http_cache_transaction.cc
        ${ANDROID_ROOT}/external/cronet/net/http/http_cache_writers.cc
        ${ANDROID_ROOT}/external/cronet/net/http/http_chunked_decoder.cc
        ${ANDROID_ROOT}/external/cronet/net/http/http_content_disposition.cc
        ${ANDROID_ROOT}/external/cronet/net/http/http_log_util.cc
        ${ANDROID_ROOT}/external/cronet/net/http/http_network_layer.cc
        ${ANDROID_ROOT}/external/cronet/net/http/http_network_session.cc
        ${ANDROID_ROOT}/external/cronet/net/http/http_network_session_peer.cc
        ${ANDROID_ROOT}/external/cronet/net/http/http_network_transaction.cc
        ${ANDROID_ROOT}/external/cronet/net/http/http_no_vary_search_data.cc
        ${ANDROID_ROOT}/external/cronet/net/http/http_proxy_client_socket.cc
        ${ANDROID_ROOT}/external/cronet/net/http/http_proxy_connect_job.cc
        ${ANDROID_ROOT}/external/cronet/net/http/http_raw_request_headers.cc
        ${ANDROID_ROOT}/external/cronet/net/http/http_request_headers.cc
        ${ANDROID_ROOT}/external/cronet/net/http/http_request_info.cc
        ${ANDROID_ROOT}/external/cronet/net/http/http_response_body_drainer.cc
        ${ANDROID_ROOT}/external/cronet/net/http/http_response_headers.cc
        ${ANDROID_ROOT}/external/cronet/net/http/http_response_info.cc
        ${ANDROID_ROOT}/external/cronet/net/http/http_security_headers.cc
        ${ANDROID_ROOT}/external/cronet/net/http/http_server_properties.cc
        ${ANDROID_ROOT}/external/cronet/net/http/http_server_properties_manager.cc
        ${ANDROID_ROOT}/external/cronet/net/http/http_status_code.cc
        ${ANDROID_ROOT}/external/cronet/net/http/http_stream.cc
        ${ANDROID_ROOT}/external/cronet/net/http/http_stream_factory.cc
        ${ANDROID_ROOT}/external/cronet/net/http/http_stream_factory_job.cc
        ${ANDROID_ROOT}/external/cronet/net/http/http_stream_factory_job_controller.cc
        ${ANDROID_ROOT}/external/cronet/net/http/http_stream_parser.cc
        ${ANDROID_ROOT}/external/cronet/net/http/http_stream_request.cc
        ${ANDROID_ROOT}/external/cronet/net/http/http_util.cc
        ${ANDROID_ROOT}/external/cronet/net/http/http_vary_data.cc
        ${ANDROID_ROOT}/external/cronet/net/http/partial_data.cc
        ${ANDROID_ROOT}/external/cronet/net/http/proxy_client_socket.cc
        ${ANDROID_ROOT}/external/cronet/net/http/proxy_fallback.cc
        ${ANDROID_ROOT}/external/cronet/net/http/transport_security_persister.cc
        ${ANDROID_ROOT}/external/cronet/net/http/transport_security_state_source.cc
        ${ANDROID_ROOT}/external/cronet/net/http/url_security_manager.cc
        ${ANDROID_ROOT}/external/cronet/net/http/url_security_manager_posix.cc
        ${ANDROID_ROOT}/external/cronet/net/http/webfonts_histogram.cc
        ${ANDROID_ROOT}/external/cronet/net/log/file_net_log_observer.cc
        ${ANDROID_ROOT}/external/cronet/net/log/net_log.cc
        ${ANDROID_ROOT}/external/cronet/net/log/net_log_capture_mode.cc
        ${ANDROID_ROOT}/external/cronet/net/log/net_log_entry.cc
        ${ANDROID_ROOT}/external/cronet/net/log/net_log_event_type.cc
        ${ANDROID_ROOT}/external/cronet/net/log/net_log_source.cc
        ${ANDROID_ROOT}/external/cronet/net/log/net_log_util.cc
        ${ANDROID_ROOT}/external/cronet/net/log/net_log_values.cc
        ${ANDROID_ROOT}/external/cronet/net/log/net_log_with_source.cc
        ${ANDROID_ROOT}/external/cronet/net/log/trace_net_log_observer.cc
        ${ANDROID_ROOT}/external/cronet/net/network_error_logging/network_error_logging_service.cc
        ${ANDROID_ROOT}/external/cronet/net/nqe/cached_network_quality.cc
        ${ANDROID_ROOT}/external/cronet/net/nqe/effective_connection_type.cc
        ${ANDROID_ROOT}/external/cronet/net/nqe/event_creator.cc
        ${ANDROID_ROOT}/external/cronet/net/nqe/network_id.cc
        ${ANDROID_ROOT}/external/cronet/net/nqe/network_qualities_prefs_manager.cc
        ${ANDROID_ROOT}/external/cronet/net/nqe/network_quality.cc
        ${ANDROID_ROOT}/external/cronet/net/nqe/network_quality_estimator.cc
        ${ANDROID_ROOT}/external/cronet/net/nqe/network_quality_estimator_params.cc
        ${ANDROID_ROOT}/external/cronet/net/nqe/network_quality_estimator_util.cc
        ${ANDROID_ROOT}/external/cronet/net/nqe/network_quality_observation.cc
        ${ANDROID_ROOT}/external/cronet/net/nqe/network_quality_store.cc
        ${ANDROID_ROOT}/external/cronet/net/nqe/observation_buffer.cc
        ${ANDROID_ROOT}/external/cronet/net/nqe/pref_names.cc
        ${ANDROID_ROOT}/external/cronet/net/nqe/socket_watcher.cc
        ${ANDROID_ROOT}/external/cronet/net/nqe/socket_watcher_factory.cc
        ${ANDROID_ROOT}/external/cronet/net/nqe/throughput_analyzer.cc
        ${ANDROID_ROOT}/external/cronet/net/ntlm/ntlm.cc
        ${ANDROID_ROOT}/external/cronet/net/ntlm/ntlm_buffer_reader.cc
        ${ANDROID_ROOT}/external/cronet/net/ntlm/ntlm_buffer_writer.cc
        ${ANDROID_ROOT}/external/cronet/net/ntlm/ntlm_client.cc
        ${ANDROID_ROOT}/external/cronet/net/ntlm/ntlm_constants.cc
        ${ANDROID_ROOT}/external/cronet/net/proxy_resolution/configured_proxy_resolution_request.cc
        ${ANDROID_ROOT}/external/cronet/net/proxy_resolution/configured_proxy_resolution_service.cc
        ${ANDROID_ROOT}/external/cronet/net/proxy_resolution/dhcp_pac_file_fetcher.cc
        ${ANDROID_ROOT}/external/cronet/net/proxy_resolution/multi_threaded_proxy_resolver.cc
        ${ANDROID_ROOT}/external/cronet/net/proxy_resolution/network_delegate_error_observer.cc
        ${ANDROID_ROOT}/external/cronet/net/proxy_resolution/pac_file_data.cc
        ${ANDROID_ROOT}/external/cronet/net/proxy_resolution/pac_file_decider.cc
        ${ANDROID_ROOT}/external/cronet/net/proxy_resolution/pac_file_fetcher.cc
        ${ANDROID_ROOT}/external/cronet/net/proxy_resolution/pac_file_fetcher_impl.cc
        ${ANDROID_ROOT}/external/cronet/net/proxy_resolution/polling_proxy_config_service.cc
        ${ANDROID_ROOT}/external/cronet/net/proxy_resolution/proxy_bypass_rules.cc
        ${ANDROID_ROOT}/external/cronet/net/proxy_resolution/proxy_config.cc
        ${ANDROID_ROOT}/external/cronet/net/proxy_resolution/proxy_config_service.cc
        ${ANDROID_ROOT}/external/cronet/net/proxy_resolution/proxy_config_service_android.cc
        ${ANDROID_ROOT}/external/cronet/net/proxy_resolution/proxy_config_service_fixed.cc
        ${ANDROID_ROOT}/external/cronet/net/proxy_resolution/proxy_config_with_annotation.cc
        ${ANDROID_ROOT}/external/cronet/net/proxy_resolution/proxy_info.cc
        ${ANDROID_ROOT}/external/cronet/net/proxy_resolution/proxy_list.cc
        ${ANDROID_ROOT}/external/cronet/net/proxy_resolution/proxy_resolver_factory.cc
        ${ANDROID_ROOT}/external/cronet/net/quic/bidirectional_stream_quic_impl.cc
        ${ANDROID_ROOT}/external/cronet/net/quic/crypto/proof_source_chromium.cc
        ${ANDROID_ROOT}/external/cronet/net/quic/crypto/proof_verifier_chromium.cc
        ${ANDROID_ROOT}/external/cronet/net/quic/dedicated_web_transport_http3_client.cc
        ${ANDROID_ROOT}/external/cronet/net/quic/network_connection.cc
        ${ANDROID_ROOT}/external/cronet/net/quic/platform/impl/quic_chromium_clock.cc
        ${ANDROID_ROOT}/external/cronet/net/quic/properties_based_quic_server_info.cc
        ${ANDROID_ROOT}/external/cronet/net/quic/quic_address_mismatch.cc
        ${ANDROID_ROOT}/external/cronet/net/quic/quic_chromium_alarm_factory.cc
        ${ANDROID_ROOT}/external/cronet/net/quic/quic_chromium_client_session.cc
        ${ANDROID_ROOT}/external/cronet/net/quic/quic_chromium_client_stream.cc
        ${ANDROID_ROOT}/external/cronet/net/quic/quic_chromium_connection_helper.cc
        ${ANDROID_ROOT}/external/cronet/net/quic/quic_chromium_packet_reader.cc
        ${ANDROID_ROOT}/external/cronet/net/quic/quic_chromium_packet_writer.cc
        ${ANDROID_ROOT}/external/cronet/net/quic/quic_clock_skew_detector.cc
        ${ANDROID_ROOT}/external/cronet/net/quic/quic_connection_logger.cc
        ${ANDROID_ROOT}/external/cronet/net/quic/quic_connectivity_monitor.cc
        ${ANDROID_ROOT}/external/cronet/net/quic/quic_context.cc
        ${ANDROID_ROOT}/external/cronet/net/quic/quic_crypto_client_config_handle.cc
        ${ANDROID_ROOT}/external/cronet/net/quic/quic_crypto_client_stream_factory.cc
        ${ANDROID_ROOT}/external/cronet/net/quic/quic_event_logger.cc
        ${ANDROID_ROOT}/external/cronet/net/quic/quic_http3_logger.cc
        ${ANDROID_ROOT}/external/cronet/net/quic/quic_http_stream.cc
        ${ANDROID_ROOT}/external/cronet/net/quic/quic_http_utils.cc
        ${ANDROID_ROOT}/external/cronet/net/quic/quic_proxy_client_socket.cc
        ${ANDROID_ROOT}/external/cronet/net/quic/quic_server_info.cc
        ${ANDROID_ROOT}/external/cronet/net/quic/quic_session_key.cc
        ${ANDROID_ROOT}/external/cronet/net/quic/quic_stream_factory.cc
        ${ANDROID_ROOT}/external/cronet/net/quic/set_quic_flag.cc
        ${ANDROID_ROOT}/external/cronet/net/quic/web_transport_client.cc
        ${ANDROID_ROOT}/external/cronet/net/quic/web_transport_error.cc
        ${ANDROID_ROOT}/external/cronet/net/reporting/reporting_browsing_data_remover.cc
        ${ANDROID_ROOT}/external/cronet/net/reporting/reporting_cache.cc
        ${ANDROID_ROOT}/external/cronet/net/reporting/reporting_cache_impl.cc
        ${ANDROID_ROOT}/external/cronet/net/reporting/reporting_cache_observer.cc
        ${ANDROID_ROOT}/external/cronet/net/reporting/reporting_context.cc
        ${ANDROID_ROOT}/external/cronet/net/reporting/reporting_delegate.cc
        ${ANDROID_ROOT}/external/cronet/net/reporting/reporting_delivery_agent.cc
        ${ANDROID_ROOT}/external/cronet/net/reporting/reporting_endpoint.cc
        ${ANDROID_ROOT}/external/cronet/net/reporting/reporting_endpoint_manager.cc
        ${ANDROID_ROOT}/external/cronet/net/reporting/reporting_garbage_collector.cc
        ${ANDROID_ROOT}/external/cronet/net/reporting/reporting_header_parser.cc
        ${ANDROID_ROOT}/external/cronet/net/reporting/reporting_network_change_observer.cc
        ${ANDROID_ROOT}/external/cronet/net/reporting/reporting_policy.cc
        ${ANDROID_ROOT}/external/cronet/net/reporting/reporting_report.cc
        ${ANDROID_ROOT}/external/cronet/net/reporting/reporting_service.cc
        ${ANDROID_ROOT}/external/cronet/net/reporting/reporting_uploader.cc
        ${ANDROID_ROOT}/external/cronet/net/socket/client_socket_factory.cc
        ${ANDROID_ROOT}/external/cronet/net/socket/client_socket_handle.cc
        ${ANDROID_ROOT}/external/cronet/net/socket/client_socket_pool.cc
        ${ANDROID_ROOT}/external/cronet/net/socket/client_socket_pool_manager.cc
        ${ANDROID_ROOT}/external/cronet/net/socket/client_socket_pool_manager_impl.cc
        ${ANDROID_ROOT}/external/cronet/net/socket/connect_job.cc
        ${ANDROID_ROOT}/external/cronet/net/socket/connect_job_factory.cc
        ${ANDROID_ROOT}/external/cronet/net/socket/network_binding_client_socket_factory.cc
        ${ANDROID_ROOT}/external/cronet/net/socket/next_proto.cc
        ${ANDROID_ROOT}/external/cronet/net/socket/server_socket.cc
        ${ANDROID_ROOT}/external/cronet/net/socket/socket.cc
        ${ANDROID_ROOT}/external/cronet/net/socket/socket_bio_adapter.cc
        ${ANDROID_ROOT}/external/cronet/net/socket/socket_descriptor.cc
        ${ANDROID_ROOT}/external/cronet/net/socket/socket_net_log_params.cc
        ${ANDROID_ROOT}/external/cronet/net/socket/socket_options.cc
        ${ANDROID_ROOT}/external/cronet/net/socket/socket_posix.cc
        ${ANDROID_ROOT}/external/cronet/net/socket/socket_tag.cc
        ${ANDROID_ROOT}/external/cronet/net/socket/socks5_client_socket.cc
        ${ANDROID_ROOT}/external/cronet/net/socket/socks_client_socket.cc
        ${ANDROID_ROOT}/external/cronet/net/socket/socks_connect_job.cc
        ${ANDROID_ROOT}/external/cronet/net/socket/ssl_client_socket.cc
        ${ANDROID_ROOT}/external/cronet/net/socket/ssl_client_socket_impl.cc
        ${ANDROID_ROOT}/external/cronet/net/socket/ssl_connect_job.cc
        ${ANDROID_ROOT}/external/cronet/net/socket/ssl_server_socket_impl.cc
        ${ANDROID_ROOT}/external/cronet/net/socket/stream_socket.cc
        ${ANDROID_ROOT}/external/cronet/net/socket/tcp_client_socket.cc
        ${ANDROID_ROOT}/external/cronet/net/socket/tcp_server_socket.cc
        ${ANDROID_ROOT}/external/cronet/net/socket/tcp_socket_posix.cc
        ${ANDROID_ROOT}/external/cronet/net/socket/transport_client_socket.cc
        ${ANDROID_ROOT}/external/cronet/net/socket/transport_client_socket_pool.cc
        ${ANDROID_ROOT}/external/cronet/net/socket/transport_connect_job.cc
        ${ANDROID_ROOT}/external/cronet/net/socket/transport_connect_sub_job.cc
        ${ANDROID_ROOT}/external/cronet/net/socket/udp_client_socket.cc
        ${ANDROID_ROOT}/external/cronet/net/socket/udp_net_log_parameters.cc
        ${ANDROID_ROOT}/external/cronet/net/socket/udp_server_socket.cc
        ${ANDROID_ROOT}/external/cronet/net/socket/udp_socket_global_limits.cc
        ${ANDROID_ROOT}/external/cronet/net/socket/udp_socket_posix.cc
        ${ANDROID_ROOT}/external/cronet/net/socket/unix_domain_client_socket_posix.cc
        ${ANDROID_ROOT}/external/cronet/net/socket/unix_domain_server_socket_posix.cc
        ${ANDROID_ROOT}/external/cronet/net/socket/websocket_endpoint_lock_manager.cc
        ${ANDROID_ROOT}/external/cronet/net/socket/websocket_transport_client_socket_pool.cc
        ${ANDROID_ROOT}/external/cronet/net/spdy/alps_decoder.cc
        ${ANDROID_ROOT}/external/cronet/net/spdy/bidirectional_stream_spdy_impl.cc
        ${ANDROID_ROOT}/external/cronet/net/spdy/buffered_spdy_framer.cc
        ${ANDROID_ROOT}/external/cronet/net/spdy/header_coalescer.cc
        ${ANDROID_ROOT}/external/cronet/net/spdy/http2_priority_dependencies.cc
        ${ANDROID_ROOT}/external/cronet/net/spdy/http2_push_promise_index.cc
        ${ANDROID_ROOT}/external/cronet/net/spdy/multiplexed_http_stream.cc
        ${ANDROID_ROOT}/external/cronet/net/spdy/multiplexed_session.cc
        ${ANDROID_ROOT}/external/cronet/net/spdy/spdy_buffer.cc
        ${ANDROID_ROOT}/external/cronet/net/spdy/spdy_buffer_producer.cc
        ${ANDROID_ROOT}/external/cronet/net/spdy/spdy_http_stream.cc
        ${ANDROID_ROOT}/external/cronet/net/spdy/spdy_http_utils.cc
        ${ANDROID_ROOT}/external/cronet/net/spdy/spdy_log_util.cc
        ${ANDROID_ROOT}/external/cronet/net/spdy/spdy_proxy_client_socket.cc
        ${ANDROID_ROOT}/external/cronet/net/spdy/spdy_read_queue.cc
        ${ANDROID_ROOT}/external/cronet/net/spdy/spdy_session.cc
        ${ANDROID_ROOT}/external/cronet/net/spdy/spdy_session_key.cc
        ${ANDROID_ROOT}/external/cronet/net/spdy/spdy_session_pool.cc
        ${ANDROID_ROOT}/external/cronet/net/spdy/spdy_stream.cc
        ${ANDROID_ROOT}/external/cronet/net/spdy/spdy_write_queue.cc
        ${ANDROID_ROOT}/external/cronet/net/ssl/cert_compression.cc
        ${ANDROID_ROOT}/external/cronet/net/ssl/client_cert_identity.cc
        ${ANDROID_ROOT}/external/cronet/net/ssl/openssl_ssl_util.cc
        ${ANDROID_ROOT}/external/cronet/net/ssl/ssl_cert_request_info.cc
        ${ANDROID_ROOT}/external/cronet/net/ssl/ssl_cipher_suite_names.cc
        ${ANDROID_ROOT}/external/cronet/net/ssl/ssl_client_auth_cache.cc
        ${ANDROID_ROOT}/external/cronet/net/ssl/ssl_client_session_cache.cc
        ${ANDROID_ROOT}/external/cronet/net/ssl/ssl_config.cc
        ${ANDROID_ROOT}/external/cronet/net/ssl/ssl_config_service.cc
        ${ANDROID_ROOT}/external/cronet/net/ssl/ssl_config_service_defaults.cc
        ${ANDROID_ROOT}/external/cronet/net/ssl/ssl_info.cc
        ${ANDROID_ROOT}/external/cronet/net/ssl/ssl_key_logger.cc
        ${ANDROID_ROOT}/external/cronet/net/ssl/ssl_key_logger_impl.cc
        ${ANDROID_ROOT}/external/cronet/net/ssl/ssl_platform_key_android.cc
        ${ANDROID_ROOT}/external/cronet/net/ssl/ssl_platform_key_util.cc
        ${ANDROID_ROOT}/external/cronet/net/ssl/ssl_private_key.cc
        ${ANDROID_ROOT}/external/cronet/net/ssl/ssl_server_config.cc
        ${ANDROID_ROOT}/external/cronet/net/ssl/threaded_ssl_private_key.cc
        ${ANDROID_ROOT}/external/cronet/net/url_request/redirect_info.cc
        ${ANDROID_ROOT}/external/cronet/net/url_request/redirect_util.cc
        ${ANDROID_ROOT}/external/cronet/net/url_request/report_sender.cc
        ${ANDROID_ROOT}/external/cronet/net/url_request/static_http_user_agent_settings.cc
        ${ANDROID_ROOT}/external/cronet/net/url_request/url_request.cc
        ${ANDROID_ROOT}/external/cronet/net/url_request/url_request_context.cc
        ${ANDROID_ROOT}/external/cronet/net/url_request/url_request_context_builder.cc
        ${ANDROID_ROOT}/external/cronet/net/url_request/url_request_context_getter.cc
        ${ANDROID_ROOT}/external/cronet/net/url_request/url_request_error_job.cc
        ${ANDROID_ROOT}/external/cronet/net/url_request/url_request_filter.cc
        ${ANDROID_ROOT}/external/cronet/net/url_request/url_request_http_job.cc
        ${ANDROID_ROOT}/external/cronet/net/url_request/url_request_interceptor.cc
        ${ANDROID_ROOT}/external/cronet/net/url_request/url_request_job.cc
        ${ANDROID_ROOT}/external/cronet/net/url_request/url_request_job_factory.cc
        ${ANDROID_ROOT}/external/cronet/net/url_request/url_request_netlog_params.cc
        ${ANDROID_ROOT}/external/cronet/net/url_request/url_request_redirect_job.cc
        ${ANDROID_ROOT}/external/cronet/net/url_request/url_request_throttler_entry.cc
        ${ANDROID_ROOT}/external/cronet/net/url_request/url_request_throttler_manager.cc
        ${ANDROID_ROOT}/external/cronet/net/url_request/view_cache_helper.cc
        ${ANDROID_ROOT}/external/cronet/net/url_request/websocket_handshake_userdata_key.cc
        ${ANDROID_ROOT}/external/cronet/net/disk_cache/blockfile/mapped_file_bypass_mmap_posix.cc
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
    add_library(cronet_aml_net_net SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_net_net
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/overrides
            ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src
            ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/common/platform/default
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
            ${ANDROID_ROOT}/external/cronet/third_party/boringssl/src/include
            ${ANDROID_ROOT}/external/cronet/third_party/brotli/include
            ${ANDROID_ROOT}/external/cronet/third_party/protobuf/src
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_chromecast_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_chromeos_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_debugging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_logging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_partition_alloc_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_anchor_functions_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_android_runtime_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_android_runtime_unchecked_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_base_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_build_date/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_cfi_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_clang_profiling_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_debugging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_feature_list_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_ios_cronet_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_logging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_message_pump_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_orderfile_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_power_monitor_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_profiler_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_rust_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_sanitizer_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_synchronization_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_tracing_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_blink_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_branding_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_chromecast_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_chromeos_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_config_compiler_compiler_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_crypto_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_base_registry_controlled_domains_registry_controlled_domains/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_cronet_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_isolation_info_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_isolation_info_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/net/base
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_isolation_info_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/protos
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_net_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_url_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_net_nqe_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_net_nqe_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/net/nqe/proto
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_net_nqe_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/protos
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_third_party_quiche_net_quic_test_tools_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_third_party_quiche_net_quic_test_tools_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/net/third_party/quiche/src
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_third_party_quiche_net_quic_test_tools_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/net/third_party/quiche/src/quiche/quic/test_tools
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_third_party_quiche_net_quic_test_tools_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/protos
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_third_party_quiche_net_quic_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_third_party_quiche_net_quic_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/net/third_party/quiche/src
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_third_party_quiche_net_quic_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/protos
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_url_url_jni_headers/android_${OUT_ARCH_CPU}/gen
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_net_net PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_net_net
            cronet_aml_base_allocator_partition_allocator_partition_alloc
            cronet_aml_base_base
            cronet_aml_base_base_static
            cronet_aml_base_nodebug_assertion
            cronet_aml_base_third_party_double_conversion_double_conversion
            cronet_aml_base_third_party_dynamic_annotations_dynamic_annotations
            cronet_aml_crypto_crypto
            cronet_aml_net_dns_dns
            cronet_aml_net_dns_public_public
            cronet_aml_net_http_transport_security_state_generated_files
            cronet_aml_net_net_deps
            cronet_aml_net_net_public_deps
            cronet_aml_net_preload_decoder
            cronet_aml_net_third_party_quiche_quiche
            cronet_aml_net_traffic_annotation_traffic_annotation
            cronet_aml_net_uri_template
            cronet_aml_third_party_abseil_cpp_absl_base_base
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_malloc_internal
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait
            cronet_aml_third_party_abseil_cpp_absl_base_strerror
            cronet_aml_third_party_abseil_cpp_absl_base_throw_delegate
            cronet_aml_third_party_abseil_cpp_absl_container_hashtablez_sampler
            cronet_aml_third_party_abseil_cpp_absl_container_raw_hash_set
            cronet_aml_third_party_abseil_cpp_absl_crc_cpu_detect
            cronet_aml_third_party_abseil_cpp_absl_crc_crc32c
            cronet_aml_third_party_abseil_cpp_absl_crc_crc_cord_state
            cronet_aml_third_party_abseil_cpp_absl_crc_crc_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_debugging_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_demangle_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_examine_stack
            cronet_aml_third_party_abseil_cpp_absl_debugging_failure_signal_handler
            cronet_aml_third_party_abseil_cpp_absl_debugging_stacktrace
            cronet_aml_third_party_abseil_cpp_absl_debugging_symbolize
            cronet_aml_third_party_abseil_cpp_absl_hash_city
            cronet_aml_third_party_abseil_cpp_absl_hash_hash
            cronet_aml_third_party_abseil_cpp_absl_hash_low_level_hash
            cronet_aml_third_party_abseil_cpp_absl_log_die_if_null
            cronet_aml_third_party_abseil_cpp_absl_log_globals
            cronet_aml_third_party_abseil_cpp_absl_log_internal_check_op
            cronet_aml_third_party_abseil_cpp_absl_log_internal_conditions
            cronet_aml_third_party_abseil_cpp_absl_log_internal_format
            cronet_aml_third_party_abseil_cpp_absl_log_internal_globals
            cronet_aml_third_party_abseil_cpp_absl_log_internal_log_message
            cronet_aml_third_party_abseil_cpp_absl_log_internal_log_sink_set
            cronet_aml_third_party_abseil_cpp_absl_log_internal_nullguard
            cronet_aml_third_party_abseil_cpp_absl_log_internal_proto
            cronet_aml_third_party_abseil_cpp_absl_log_log_entry
            cronet_aml_third_party_abseil_cpp_absl_log_log_sink
            cronet_aml_third_party_abseil_cpp_absl_numeric_int128
            cronet_aml_third_party_abseil_cpp_absl_profiling_exponential_biased
            cronet_aml_third_party_abseil_cpp_absl_random_distributions
            cronet_aml_third_party_abseil_cpp_absl_random_internal_platform
            cronet_aml_third_party_abseil_cpp_absl_random_internal_pool_urbg
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_hwaes
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_hwaes_impl
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_slow
            cronet_aml_third_party_abseil_cpp_absl_random_internal_seed_material
            cronet_aml_third_party_abseil_cpp_absl_random_seed_gen_exception
            cronet_aml_third_party_abseil_cpp_absl_random_seed_sequences
            cronet_aml_third_party_abseil_cpp_absl_status_status
            cronet_aml_third_party_abseil_cpp_absl_status_statusor
            cronet_aml_third_party_abseil_cpp_absl_strings_cord
            cronet_aml_third_party_abseil_cpp_absl_strings_cord_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_functions
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_handle
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_info
            cronet_aml_third_party_abseil_cpp_absl_strings_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_str_format_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_strings
            cronet_aml_third_party_abseil_cpp_absl_synchronization_graphcycles_internal
            cronet_aml_third_party_abseil_cpp_absl_synchronization_kernel_timeout_internal
            cronet_aml_third_party_abseil_cpp_absl_synchronization_synchronization
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_civil_time
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_time_zone
            cronet_aml_third_party_abseil_cpp_absl_time_time
            cronet_aml_third_party_abseil_cpp_absl_types_bad_optional_access
            cronet_aml_third_party_abseil_cpp_absl_types_bad_variant_access
            cronet_aml_third_party_android_ndk_cpu_features
            cronet_aml_third_party_ashmem_ashmem
            cronet_aml_third_party_brotli_common
            cronet_aml_third_party_brotli_dec
            cronet_aml_third_party_icu_icui18n
            cronet_aml_third_party_icu_icuuc_private
            cronet_aml_third_party_libevent_libevent
            cronet_aml_third_party_modp_b64_modp_b64
            cronet_aml_third_party_protobuf_protobuf_lite
            cronet_aml_url_url
            jni_headers
            libandroid
            libc
            libclang_rt.builtins
            libcrypto
            libdl
            liblog
            libm
            libmediandk
            libssl
            libz
    )
endif ()
#################### cronet_aml_net_net ####################

#################### cronet_aml_third_party_abseil_cpp_absl_debugging_debugging_internal ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/debugging/internal/address_is_readable.cc
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/debugging/internal/elf_mem_image.cc
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/debugging/internal/vdso_support.cc
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
    add_library(cronet_aml_third_party_abseil_cpp_absl_debugging_debugging_internal SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_abseil_cpp_absl_debugging_debugging_internal
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_abseil_cpp_absl_debugging_debugging_internal PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_abseil_cpp_absl_debugging_debugging_internal
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_abseil_cpp_absl_debugging_debugging_internal ####################

#################### cronet_aml_net_net_deps ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_isolation_info_proto_gen/android_${OUT_ARCH_CPU}/gen/external/cronet/net/base/isolation_info.pb.cc
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
    add_library(cronet_aml_net_net_deps SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_net_net_deps
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
            ${ANDROID_ROOT}/external/cronet/third_party/boringssl/src/include
            ${ANDROID_ROOT}/external/cronet/third_party/brotli/include
            ${ANDROID_ROOT}/external/cronet/third_party/protobuf/src
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_chromecast_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_chromeos_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_debugging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_logging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_partition_alloc_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_anchor_functions_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_android_runtime_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_android_runtime_unchecked_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_base_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_build_date/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_cfi_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_clang_profiling_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_debugging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_feature_list_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_ios_cronet_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_logging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_message_pump_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_orderfile_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_power_monitor_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_profiler_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_rust_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_sanitizer_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_synchronization_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_tracing_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_blink_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_branding_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_chromecast_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_chromeos_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_config_compiler_compiler_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_base_registry_controlled_domains_registry_controlled_domains/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_cronet_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_isolation_info_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_isolation_info_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/net/base
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_isolation_info_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/protos
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_net_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_url_buildflags/android_${OUT_ARCH_CPU}/gen
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_net_net_deps PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_net_net_deps
            cronet_aml_base_allocator_partition_allocator_partition_alloc
            cronet_aml_base_base
            cronet_aml_base_base_static
            cronet_aml_base_nodebug_assertion
            cronet_aml_base_third_party_double_conversion_double_conversion
            cronet_aml_base_third_party_dynamic_annotations_dynamic_annotations
            cronet_aml_net_preload_decoder
            cronet_aml_third_party_abseil_cpp_absl_base_base
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_malloc_internal
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait
            cronet_aml_third_party_abseil_cpp_absl_base_strerror
            cronet_aml_third_party_abseil_cpp_absl_base_throw_delegate
            cronet_aml_third_party_abseil_cpp_absl_container_hashtablez_sampler
            cronet_aml_third_party_abseil_cpp_absl_container_raw_hash_set
            cronet_aml_third_party_abseil_cpp_absl_crc_cpu_detect
            cronet_aml_third_party_abseil_cpp_absl_crc_crc32c
            cronet_aml_third_party_abseil_cpp_absl_crc_crc_cord_state
            cronet_aml_third_party_abseil_cpp_absl_crc_crc_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_debugging_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_demangle_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_examine_stack
            cronet_aml_third_party_abseil_cpp_absl_debugging_failure_signal_handler
            cronet_aml_third_party_abseil_cpp_absl_debugging_stacktrace
            cronet_aml_third_party_abseil_cpp_absl_debugging_symbolize
            cronet_aml_third_party_abseil_cpp_absl_hash_city
            cronet_aml_third_party_abseil_cpp_absl_hash_hash
            cronet_aml_third_party_abseil_cpp_absl_hash_low_level_hash
            cronet_aml_third_party_abseil_cpp_absl_log_die_if_null
            cronet_aml_third_party_abseil_cpp_absl_log_globals
            cronet_aml_third_party_abseil_cpp_absl_log_internal_check_op
            cronet_aml_third_party_abseil_cpp_absl_log_internal_conditions
            cronet_aml_third_party_abseil_cpp_absl_log_internal_format
            cronet_aml_third_party_abseil_cpp_absl_log_internal_globals
            cronet_aml_third_party_abseil_cpp_absl_log_internal_log_message
            cronet_aml_third_party_abseil_cpp_absl_log_internal_log_sink_set
            cronet_aml_third_party_abseil_cpp_absl_log_internal_nullguard
            cronet_aml_third_party_abseil_cpp_absl_log_internal_proto
            cronet_aml_third_party_abseil_cpp_absl_log_log_entry
            cronet_aml_third_party_abseil_cpp_absl_log_log_sink
            cronet_aml_third_party_abseil_cpp_absl_numeric_int128
            cronet_aml_third_party_abseil_cpp_absl_profiling_exponential_biased
            cronet_aml_third_party_abseil_cpp_absl_random_distributions
            cronet_aml_third_party_abseil_cpp_absl_random_internal_platform
            cronet_aml_third_party_abseil_cpp_absl_random_internal_pool_urbg
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_hwaes
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_hwaes_impl
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_slow
            cronet_aml_third_party_abseil_cpp_absl_random_internal_seed_material
            cronet_aml_third_party_abseil_cpp_absl_random_seed_gen_exception
            cronet_aml_third_party_abseil_cpp_absl_random_seed_sequences
            cronet_aml_third_party_abseil_cpp_absl_status_status
            cronet_aml_third_party_abseil_cpp_absl_status_statusor
            cronet_aml_third_party_abseil_cpp_absl_strings_cord
            cronet_aml_third_party_abseil_cpp_absl_strings_cord_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_functions
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_handle
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_info
            cronet_aml_third_party_abseil_cpp_absl_strings_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_str_format_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_strings
            cronet_aml_third_party_abseil_cpp_absl_synchronization_graphcycles_internal
            cronet_aml_third_party_abseil_cpp_absl_synchronization_kernel_timeout_internal
            cronet_aml_third_party_abseil_cpp_absl_synchronization_synchronization
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_civil_time
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_time_zone
            cronet_aml_third_party_abseil_cpp_absl_time_time
            cronet_aml_third_party_abseil_cpp_absl_types_bad_optional_access
            cronet_aml_third_party_abseil_cpp_absl_types_bad_variant_access
            cronet_aml_third_party_android_ndk_cpu_features
            cronet_aml_third_party_ashmem_ashmem
            cronet_aml_third_party_brotli_common
            cronet_aml_third_party_brotli_dec
            cronet_aml_third_party_icu_icui18n
            cronet_aml_third_party_icu_icuuc_private
            cronet_aml_third_party_libevent_libevent
            cronet_aml_third_party_modp_b64_modp_b64
            cronet_aml_third_party_protobuf_protobuf_lite
            jni_headers
            libandroid
            libc
            libclang_rt.builtins
            libcrypto
            libdl
            liblog
            libm
            libmediandk
            libprotobuf-cpp-lite
            libssl
            libz
    )
endif ()
#################### cronet_aml_net_net_deps ####################

#################### cronet_aml_third_party_abseil_cpp_absl_profiling_exponential_biased ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/profiling/internal/exponential_biased.cc
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
    add_library(cronet_aml_third_party_abseil_cpp_absl_profiling_exponential_biased SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_abseil_cpp_absl_profiling_exponential_biased
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_abseil_cpp_absl_profiling_exponential_biased PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_abseil_cpp_absl_profiling_exponential_biased
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_abseil_cpp_absl_profiling_exponential_biased ####################

#################### cronet_aml_third_party_abseil_cpp_absl_synchronization_graphcycles_internal ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/synchronization/internal/graphcycles.cc
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
    add_library(cronet_aml_third_party_abseil_cpp_absl_synchronization_graphcycles_internal SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_abseil_cpp_absl_synchronization_graphcycles_internal
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_abseil_cpp_absl_synchronization_graphcycles_internal PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_abseil_cpp_absl_synchronization_graphcycles_internal
            cronet_aml_third_party_abseil_cpp_absl_base_base
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_malloc_internal
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_abseil_cpp_absl_synchronization_graphcycles_internal ####################

#################### cronet_aml_third_party_abseil_cpp_absl_log_internal_nullguard ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/log/internal/nullguard.cc
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
    add_library(cronet_aml_third_party_abseil_cpp_absl_log_internal_nullguard SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_abseil_cpp_absl_log_internal_nullguard
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_abseil_cpp_absl_log_internal_nullguard PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_abseil_cpp_absl_log_internal_nullguard
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_abseil_cpp_absl_log_internal_nullguard ####################

#################### cronet_aml_third_party_abseil_cpp_absl_random_internal_seed_material ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/random/internal/seed_material.cc
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
    add_library(cronet_aml_third_party_abseil_cpp_absl_random_internal_seed_material SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_abseil_cpp_absl_random_internal_seed_material
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_abseil_cpp_absl_random_internal_seed_material PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_abseil_cpp_absl_random_internal_seed_material
            cronet_aml_third_party_abseil_cpp_absl_base_base
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait
            cronet_aml_third_party_abseil_cpp_absl_base_throw_delegate
            cronet_aml_third_party_abseil_cpp_absl_numeric_int128
            cronet_aml_third_party_abseil_cpp_absl_strings_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_strings
            cronet_aml_third_party_abseil_cpp_absl_types_bad_optional_access
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_abseil_cpp_absl_random_internal_seed_material ####################

#################### cronet_aml_third_party_abseil_cpp_absl_types_bad_variant_access ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/types/bad_variant_access.cc
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
    add_library(cronet_aml_third_party_abseil_cpp_absl_types_bad_variant_access SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_abseil_cpp_absl_types_bad_variant_access
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_abseil_cpp_absl_types_bad_variant_access PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_abseil_cpp_absl_types_bad_variant_access
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_abseil_cpp_absl_types_bad_variant_access ####################

#################### cronet_aml_net_net_public_deps ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_net_nqe_proto_gen/android_${OUT_ARCH_CPU}/gen/external/cronet/net/nqe/proto/network_id_proto.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_third_party_quiche_net_quic_test_tools_proto_gen/android_${OUT_ARCH_CPU}/gen/external/cronet/net/third_party/quiche/src/quiche/quic/test_tools/send_algorithm_test_result.pb.cc
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
    add_library(cronet_aml_net_net_public_deps SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_net_net_public_deps
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/overrides
            ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src
            ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/common/platform/default
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
            ${ANDROID_ROOT}/external/cronet/third_party/boringssl/src/include
            ${ANDROID_ROOT}/external/cronet/third_party/protobuf/src
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_chromecast_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_chromeos_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_debugging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_logging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_partition_alloc_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_anchor_functions_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_android_runtime_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_android_runtime_unchecked_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_base_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_build_date/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_cfi_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_clang_profiling_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_debugging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_feature_list_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_ios_cronet_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_logging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_message_pump_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_orderfile_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_power_monitor_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_profiler_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_rust_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_sanitizer_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_synchronization_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_tracing_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_blink_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_branding_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_chromecast_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_chromeos_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_config_compiler_compiler_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_crypto_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_third_party_quiche_net_quic_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_third_party_quiche_net_quic_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/net/third_party/quiche/src
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_third_party_quiche_net_quic_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/protos
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_url_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_url_url_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_net_nqe_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_net_nqe_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/net/nqe/proto
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_net_nqe_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/protos
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_third_party_quiche_net_quic_test_tools_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_third_party_quiche_net_quic_test_tools_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/net/third_party/quiche/src
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_third_party_quiche_net_quic_test_tools_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/net/third_party/quiche/src/quiche/quic/test_tools
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_third_party_quiche_net_quic_test_tools_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/protos
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_net_net_public_deps PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_net_net_public_deps
            cronet_aml_base_allocator_partition_allocator_partition_alloc
            cronet_aml_base_base
            cronet_aml_base_base_static
            cronet_aml_base_nodebug_assertion
            cronet_aml_base_third_party_double_conversion_double_conversion
            cronet_aml_base_third_party_dynamic_annotations_dynamic_annotations
            cronet_aml_crypto_crypto
            cronet_aml_net_third_party_quiche_quiche
            cronet_aml_net_traffic_annotation_traffic_annotation
            cronet_aml_net_uri_template
            cronet_aml_third_party_abseil_cpp_absl_base_base
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_malloc_internal
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait
            cronet_aml_third_party_abseil_cpp_absl_base_strerror
            cronet_aml_third_party_abseil_cpp_absl_base_throw_delegate
            cronet_aml_third_party_abseil_cpp_absl_container_hashtablez_sampler
            cronet_aml_third_party_abseil_cpp_absl_container_raw_hash_set
            cronet_aml_third_party_abseil_cpp_absl_crc_cpu_detect
            cronet_aml_third_party_abseil_cpp_absl_crc_crc32c
            cronet_aml_third_party_abseil_cpp_absl_crc_crc_cord_state
            cronet_aml_third_party_abseil_cpp_absl_crc_crc_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_debugging_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_demangle_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_examine_stack
            cronet_aml_third_party_abseil_cpp_absl_debugging_failure_signal_handler
            cronet_aml_third_party_abseil_cpp_absl_debugging_stacktrace
            cronet_aml_third_party_abseil_cpp_absl_debugging_symbolize
            cronet_aml_third_party_abseil_cpp_absl_hash_city
            cronet_aml_third_party_abseil_cpp_absl_hash_hash
            cronet_aml_third_party_abseil_cpp_absl_hash_low_level_hash
            cronet_aml_third_party_abseil_cpp_absl_log_die_if_null
            cronet_aml_third_party_abseil_cpp_absl_log_globals
            cronet_aml_third_party_abseil_cpp_absl_log_internal_check_op
            cronet_aml_third_party_abseil_cpp_absl_log_internal_conditions
            cronet_aml_third_party_abseil_cpp_absl_log_internal_format
            cronet_aml_third_party_abseil_cpp_absl_log_internal_globals
            cronet_aml_third_party_abseil_cpp_absl_log_internal_log_message
            cronet_aml_third_party_abseil_cpp_absl_log_internal_log_sink_set
            cronet_aml_third_party_abseil_cpp_absl_log_internal_nullguard
            cronet_aml_third_party_abseil_cpp_absl_log_internal_proto
            cronet_aml_third_party_abseil_cpp_absl_log_log_entry
            cronet_aml_third_party_abseil_cpp_absl_log_log_sink
            cronet_aml_third_party_abseil_cpp_absl_numeric_int128
            cronet_aml_third_party_abseil_cpp_absl_profiling_exponential_biased
            cronet_aml_third_party_abseil_cpp_absl_random_distributions
            cronet_aml_third_party_abseil_cpp_absl_random_internal_platform
            cronet_aml_third_party_abseil_cpp_absl_random_internal_pool_urbg
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_hwaes
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_hwaes_impl
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_slow
            cronet_aml_third_party_abseil_cpp_absl_random_internal_seed_material
            cronet_aml_third_party_abseil_cpp_absl_random_seed_gen_exception
            cronet_aml_third_party_abseil_cpp_absl_random_seed_sequences
            cronet_aml_third_party_abseil_cpp_absl_status_status
            cronet_aml_third_party_abseil_cpp_absl_status_statusor
            cronet_aml_third_party_abseil_cpp_absl_strings_cord
            cronet_aml_third_party_abseil_cpp_absl_strings_cord_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_functions
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_handle
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_info
            cronet_aml_third_party_abseil_cpp_absl_strings_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_str_format_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_strings
            cronet_aml_third_party_abseil_cpp_absl_synchronization_graphcycles_internal
            cronet_aml_third_party_abseil_cpp_absl_synchronization_kernel_timeout_internal
            cronet_aml_third_party_abseil_cpp_absl_synchronization_synchronization
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_civil_time
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_time_zone
            cronet_aml_third_party_abseil_cpp_absl_time_time
            cronet_aml_third_party_abseil_cpp_absl_types_bad_optional_access
            cronet_aml_third_party_abseil_cpp_absl_types_bad_variant_access
            cronet_aml_third_party_android_ndk_cpu_features
            cronet_aml_third_party_ashmem_ashmem
            cronet_aml_third_party_icu_icui18n
            cronet_aml_third_party_icu_icuuc_private
            cronet_aml_third_party_libevent_libevent
            cronet_aml_third_party_modp_b64_modp_b64
            cronet_aml_third_party_protobuf_protobuf_lite
            cronet_aml_url_url
            jni_headers
            libandroid
            libc
            libclang_rt.builtins
            libcrypto
            libdl
            liblog
            libm
            libmediandk
            libprotobuf-cpp-lite
            libssl
            libz
    )
endif ()
#################### cronet_aml_net_net_public_deps ####################

#################### cronet_aml_net_uri_template ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/net/third_party/uri_template/uri_template.cc
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
    add_library(cronet_aml_net_uri_template SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_net_uri_template
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
            ${ANDROID_ROOT}/external/cronet/third_party/boringssl/src/include
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_chromecast_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_chromeos_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_debugging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_logging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_partition_alloc_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_anchor_functions_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_android_runtime_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_android_runtime_unchecked_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_base_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_build_date/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_cfi_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_clang_profiling_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_debugging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_feature_list_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_ios_cronet_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_logging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_message_pump_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_orderfile_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_power_monitor_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_profiler_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_rust_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_sanitizer_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_synchronization_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_tracing_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_blink_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_branding_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_chromecast_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_chromeos_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_config_compiler_compiler_buildflags/android_${OUT_ARCH_CPU}/gen
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_net_uri_template PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_net_uri_template
            cronet_aml_base_allocator_partition_allocator_partition_alloc
            cronet_aml_base_base
            cronet_aml_base_base_static
            cronet_aml_base_nodebug_assertion
            cronet_aml_base_third_party_double_conversion_double_conversion
            cronet_aml_base_third_party_dynamic_annotations_dynamic_annotations
            cronet_aml_third_party_abseil_cpp_absl_base_base
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_malloc_internal
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait
            cronet_aml_third_party_abseil_cpp_absl_base_strerror
            cronet_aml_third_party_abseil_cpp_absl_base_throw_delegate
            cronet_aml_third_party_abseil_cpp_absl_container_hashtablez_sampler
            cronet_aml_third_party_abseil_cpp_absl_container_raw_hash_set
            cronet_aml_third_party_abseil_cpp_absl_crc_cpu_detect
            cronet_aml_third_party_abseil_cpp_absl_crc_crc32c
            cronet_aml_third_party_abseil_cpp_absl_crc_crc_cord_state
            cronet_aml_third_party_abseil_cpp_absl_crc_crc_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_debugging_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_demangle_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_examine_stack
            cronet_aml_third_party_abseil_cpp_absl_debugging_failure_signal_handler
            cronet_aml_third_party_abseil_cpp_absl_debugging_stacktrace
            cronet_aml_third_party_abseil_cpp_absl_debugging_symbolize
            cronet_aml_third_party_abseil_cpp_absl_hash_city
            cronet_aml_third_party_abseil_cpp_absl_hash_hash
            cronet_aml_third_party_abseil_cpp_absl_hash_low_level_hash
            cronet_aml_third_party_abseil_cpp_absl_log_die_if_null
            cronet_aml_third_party_abseil_cpp_absl_log_globals
            cronet_aml_third_party_abseil_cpp_absl_log_internal_check_op
            cronet_aml_third_party_abseil_cpp_absl_log_internal_conditions
            cronet_aml_third_party_abseil_cpp_absl_log_internal_format
            cronet_aml_third_party_abseil_cpp_absl_log_internal_globals
            cronet_aml_third_party_abseil_cpp_absl_log_internal_log_message
            cronet_aml_third_party_abseil_cpp_absl_log_internal_log_sink_set
            cronet_aml_third_party_abseil_cpp_absl_log_internal_nullguard
            cronet_aml_third_party_abseil_cpp_absl_log_internal_proto
            cronet_aml_third_party_abseil_cpp_absl_log_log_entry
            cronet_aml_third_party_abseil_cpp_absl_log_log_sink
            cronet_aml_third_party_abseil_cpp_absl_numeric_int128
            cronet_aml_third_party_abseil_cpp_absl_profiling_exponential_biased
            cronet_aml_third_party_abseil_cpp_absl_random_distributions
            cronet_aml_third_party_abseil_cpp_absl_random_internal_platform
            cronet_aml_third_party_abseil_cpp_absl_random_internal_pool_urbg
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_hwaes
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_hwaes_impl
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_slow
            cronet_aml_third_party_abseil_cpp_absl_random_internal_seed_material
            cronet_aml_third_party_abseil_cpp_absl_random_seed_gen_exception
            cronet_aml_third_party_abseil_cpp_absl_random_seed_sequences
            cronet_aml_third_party_abseil_cpp_absl_status_status
            cronet_aml_third_party_abseil_cpp_absl_status_statusor
            cronet_aml_third_party_abseil_cpp_absl_strings_cord
            cronet_aml_third_party_abseil_cpp_absl_strings_cord_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_functions
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_handle
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_info
            cronet_aml_third_party_abseil_cpp_absl_strings_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_str_format_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_strings
            cronet_aml_third_party_abseil_cpp_absl_synchronization_graphcycles_internal
            cronet_aml_third_party_abseil_cpp_absl_synchronization_kernel_timeout_internal
            cronet_aml_third_party_abseil_cpp_absl_synchronization_synchronization
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_civil_time
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_time_zone
            cronet_aml_third_party_abseil_cpp_absl_time_time
            cronet_aml_third_party_abseil_cpp_absl_types_bad_optional_access
            cronet_aml_third_party_abseil_cpp_absl_types_bad_variant_access
            cronet_aml_third_party_android_ndk_cpu_features
            cronet_aml_third_party_ashmem_ashmem
            cronet_aml_third_party_icu_icui18n
            cronet_aml_third_party_icu_icuuc_private
            cronet_aml_third_party_libevent_libevent
            cronet_aml_third_party_modp_b64_modp_b64
            jni_headers
            libandroid
            libc
            libclang_rt.builtins
            libcrypto
            libdl
            liblog
            libm
            libmediandk
            libssl
    )
endif ()
#################### cronet_aml_net_uri_template ####################

#################### cronet_aml_third_party_android_ndk_cpu_features ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/android_ndk/sources/android/cpufeatures/cpu-features.c
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
    add_library(cronet_aml_third_party_android_ndk_cpu_features SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_android_ndk_cpu_features
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/android_ndk/sources/android/cpufeatures
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_android_ndk_cpu_features PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_android_ndk_cpu_features
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_android_ndk_cpu_features ####################

#################### cronet_aml_third_party_abseil_cpp_absl_random_internal_pool_urbg ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/random/internal/pool_urbg.cc
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
    add_library(cronet_aml_third_party_abseil_cpp_absl_random_internal_pool_urbg SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_abseil_cpp_absl_random_internal_pool_urbg
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_chromeos_buildflags/android_${OUT_ARCH_CPU}/gen
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_abseil_cpp_absl_random_internal_pool_urbg PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_abseil_cpp_absl_random_internal_pool_urbg
            cronet_aml_third_party_abseil_cpp_absl_base_base
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait
            cronet_aml_third_party_abseil_cpp_absl_base_throw_delegate
            cronet_aml_third_party_abseil_cpp_absl_numeric_int128
            cronet_aml_third_party_abseil_cpp_absl_random_internal_platform
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_hwaes
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_hwaes_impl
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_slow
            cronet_aml_third_party_abseil_cpp_absl_random_internal_seed_material
            cronet_aml_third_party_abseil_cpp_absl_random_seed_gen_exception
            cronet_aml_third_party_abseil_cpp_absl_strings_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_strings
            cronet_aml_third_party_abseil_cpp_absl_types_bad_optional_access
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_abseil_cpp_absl_random_internal_pool_urbg ####################

#################### cronet_aml_third_party_icu_icuuc_private ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/appendable.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/bmpset.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/brkeng.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/brkiter.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/bytesinkutil.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/bytestream.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/bytestrie.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/bytestriebuilder.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/bytestrieiterator.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/caniter.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/characterproperties.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/chariter.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/charstr.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/cmemory.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/cstr.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/cstring.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/cwchar.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/dictbe.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/dictionarydata.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/dtintrv.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/edits.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/emojiprops.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/errorcode.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/filteredbrk.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/filterednormalizer2.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/icudataver.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/icuplug.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/loadednormalizer2impl.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/localebuilder.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/localematcher.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/localeprioritylist.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/locavailable.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/locbased.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/locdispnames.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/locdistance.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/locdspnm.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/locid.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/loclikely.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/loclikelysubtags.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/locmap.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/locresdata.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/locutil.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/lsr.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/lstmbe.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/messagepattern.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/normalizer2.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/normalizer2impl.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/normlzr.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/parsepos.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/patternprops.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/pluralmap.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/propname.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/propsvec.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/punycode.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/putil.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/rbbi.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/rbbi_cache.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/rbbidata.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/rbbinode.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/rbbirb.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/rbbiscan.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/rbbisetb.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/rbbistbl.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/rbbitblb.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/resbund.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/resbund_cnv.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/resource.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/restrace.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/ruleiter.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/schriter.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/serv.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/servlk.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/servlkf.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/servls.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/servnotf.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/servrbf.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/servslkf.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/sharedobject.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/simpleformatter.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/static_unicode_sets.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/stringpiece.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/stringtriebuilder.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/uarrsort.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/ubidi.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/ubidi_props.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/ubidiln.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/ubiditransform.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/ubidiwrt.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/ubrk.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/ucase.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/ucasemap.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/ucasemap_titlecase_brkiter.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/ucat.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/uchar.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/ucharstrie.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/ucharstriebuilder.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/ucharstrieiterator.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/uchriter.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/ucln_cmn.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/ucmndata.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/ucnv.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/ucnv2022.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/ucnv_bld.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/ucnv_cb.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/ucnv_cnv.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/ucnv_ct.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/ucnv_err.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/ucnv_ext.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/ucnv_io.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/ucnv_lmb.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/ucnv_set.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/ucnv_u16.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/ucnv_u32.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/ucnv_u7.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/ucnv_u8.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/ucnvbocu.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/ucnvdisp.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/ucnvhz.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/ucnvisci.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/ucnvlat1.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/ucnvmbcs.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/ucnvscsu.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/ucnvsel.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/ucol_swp.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/ucptrie.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/ucurr.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/udata.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/udatamem.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/udataswp.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/uenum.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/uhash.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/uhash_us.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/uidna.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/uinit.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/uinvchar.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/uiter.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/ulist.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/uloc.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/uloc_keytype.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/uloc_tag.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/umapfile.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/umath.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/umutablecptrie.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/umutex.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/unames.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/unifiedcache.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/unifilt.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/unifunct.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/uniset.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/uniset_closure.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/uniset_props.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/unisetspan.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/unistr.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/unistr_case.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/unistr_case_locale.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/unistr_cnv.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/unistr_props.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/unistr_titlecase_brkiter.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/unorm.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/unormcmp.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/uobject.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/uprops.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/ures_cnv.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/uresbund.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/uresdata.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/usc_impl.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/uscript.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/uscript_props.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/uset.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/uset_props.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/usetiter.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/ushape.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/usprep.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/ustack.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/ustr_cnv.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/ustr_titlecase_brkiter.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/ustr_wcs.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/ustrcase.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/ustrcase_locale.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/ustrenum.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/ustrfmt.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/ustring.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/ustrtrns.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/utext.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/utf_impl.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/util.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/util_props.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/utrace.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/utrie.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/utrie2.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/utrie2_builder.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/utrie_swap.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/uts46.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/utypes.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/uvector.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/uvectr32.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/uvectr64.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common/wintz.cpp
        ${ANDROID_ROOT}/external/cronet/third_party/icu/source/stubdata/stubdata.cpp
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
    add_library(cronet_aml_third_party_icu_icuuc_private SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_icu_icuuc_private
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common
            ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_icu_icuuc_private PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_icu_icuuc_private
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_icu_icuuc_private ####################

#################### cronet_aml_base_nodebug_assertion ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/base/nodebug_assertion.cc
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
    add_library(cronet_aml_base_nodebug_assertion SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_base_nodebug_assertion
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_chromeos_buildflags/android_${OUT_ARCH_CPU}/gen
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_base_nodebug_assertion PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_base_nodebug_assertion
            cronet_aml_base_base_static
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_base_nodebug_assertion ####################

#################### cronet_aml_third_party_abseil_cpp_absl_synchronization_kernel_timeout_internal ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/synchronization/internal/kernel_timeout.cc
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
    add_library(cronet_aml_third_party_abseil_cpp_absl_synchronization_kernel_timeout_internal SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_abseil_cpp_absl_synchronization_kernel_timeout_internal
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_abseil_cpp_absl_synchronization_kernel_timeout_internal PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_abseil_cpp_absl_synchronization_kernel_timeout_internal
            cronet_aml_third_party_abseil_cpp_absl_base_base
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait
            cronet_aml_third_party_abseil_cpp_absl_base_throw_delegate
            cronet_aml_third_party_abseil_cpp_absl_numeric_int128
            cronet_aml_third_party_abseil_cpp_absl_strings_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_strings
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_civil_time
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_time_zone
            cronet_aml_third_party_abseil_cpp_absl_time_time
            cronet_aml_third_party_abseil_cpp_absl_types_bad_optional_access
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_abseil_cpp_absl_synchronization_kernel_timeout_internal ####################

#################### cronet_aml_components_prefs_prefs_features ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/components/prefs/prefs_features.cc
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
    add_library(cronet_aml_components_prefs_prefs_features SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_components_prefs_prefs_features
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
            ${ANDROID_ROOT}/external/cronet/third_party/boringssl/src/include
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_chromecast_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_chromeos_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_debugging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_logging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_partition_alloc_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_anchor_functions_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_android_runtime_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_android_runtime_unchecked_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_base_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_build_date/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_cfi_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_clang_profiling_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_debugging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_feature_list_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_ios_cronet_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_logging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_message_pump_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_orderfile_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_power_monitor_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_profiler_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_rust_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_sanitizer_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_synchronization_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_tracing_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_blink_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_branding_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_chromecast_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_chromeos_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_config_compiler_compiler_buildflags/android_${OUT_ARCH_CPU}/gen
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_components_prefs_prefs_features PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_components_prefs_prefs_features
            cronet_aml_base_allocator_partition_allocator_partition_alloc
            cronet_aml_base_base
            cronet_aml_base_base_static
            cronet_aml_base_nodebug_assertion
            cronet_aml_base_third_party_double_conversion_double_conversion
            cronet_aml_base_third_party_dynamic_annotations_dynamic_annotations
            cronet_aml_third_party_abseil_cpp_absl_base_base
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_malloc_internal
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait
            cronet_aml_third_party_abseil_cpp_absl_base_strerror
            cronet_aml_third_party_abseil_cpp_absl_base_throw_delegate
            cronet_aml_third_party_abseil_cpp_absl_container_hashtablez_sampler
            cronet_aml_third_party_abseil_cpp_absl_container_raw_hash_set
            cronet_aml_third_party_abseil_cpp_absl_crc_cpu_detect
            cronet_aml_third_party_abseil_cpp_absl_crc_crc32c
            cronet_aml_third_party_abseil_cpp_absl_crc_crc_cord_state
            cronet_aml_third_party_abseil_cpp_absl_crc_crc_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_debugging_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_demangle_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_examine_stack
            cronet_aml_third_party_abseil_cpp_absl_debugging_failure_signal_handler
            cronet_aml_third_party_abseil_cpp_absl_debugging_stacktrace
            cronet_aml_third_party_abseil_cpp_absl_debugging_symbolize
            cronet_aml_third_party_abseil_cpp_absl_hash_city
            cronet_aml_third_party_abseil_cpp_absl_hash_hash
            cronet_aml_third_party_abseil_cpp_absl_hash_low_level_hash
            cronet_aml_third_party_abseil_cpp_absl_log_die_if_null
            cronet_aml_third_party_abseil_cpp_absl_log_globals
            cronet_aml_third_party_abseil_cpp_absl_log_internal_check_op
            cronet_aml_third_party_abseil_cpp_absl_log_internal_conditions
            cronet_aml_third_party_abseil_cpp_absl_log_internal_format
            cronet_aml_third_party_abseil_cpp_absl_log_internal_globals
            cronet_aml_third_party_abseil_cpp_absl_log_internal_log_message
            cronet_aml_third_party_abseil_cpp_absl_log_internal_log_sink_set
            cronet_aml_third_party_abseil_cpp_absl_log_internal_nullguard
            cronet_aml_third_party_abseil_cpp_absl_log_internal_proto
            cronet_aml_third_party_abseil_cpp_absl_log_log_entry
            cronet_aml_third_party_abseil_cpp_absl_log_log_sink
            cronet_aml_third_party_abseil_cpp_absl_numeric_int128
            cronet_aml_third_party_abseil_cpp_absl_profiling_exponential_biased
            cronet_aml_third_party_abseil_cpp_absl_random_distributions
            cronet_aml_third_party_abseil_cpp_absl_random_internal_platform
            cronet_aml_third_party_abseil_cpp_absl_random_internal_pool_urbg
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_hwaes
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_hwaes_impl
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_slow
            cronet_aml_third_party_abseil_cpp_absl_random_internal_seed_material
            cronet_aml_third_party_abseil_cpp_absl_random_seed_gen_exception
            cronet_aml_third_party_abseil_cpp_absl_random_seed_sequences
            cronet_aml_third_party_abseil_cpp_absl_status_status
            cronet_aml_third_party_abseil_cpp_absl_status_statusor
            cronet_aml_third_party_abseil_cpp_absl_strings_cord
            cronet_aml_third_party_abseil_cpp_absl_strings_cord_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_functions
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_handle
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_info
            cronet_aml_third_party_abseil_cpp_absl_strings_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_str_format_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_strings
            cronet_aml_third_party_abseil_cpp_absl_synchronization_graphcycles_internal
            cronet_aml_third_party_abseil_cpp_absl_synchronization_kernel_timeout_internal
            cronet_aml_third_party_abseil_cpp_absl_synchronization_synchronization
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_civil_time
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_time_zone
            cronet_aml_third_party_abseil_cpp_absl_time_time
            cronet_aml_third_party_abseil_cpp_absl_types_bad_optional_access
            cronet_aml_third_party_abseil_cpp_absl_types_bad_variant_access
            cronet_aml_third_party_android_ndk_cpu_features
            cronet_aml_third_party_ashmem_ashmem
            cronet_aml_third_party_icu_icui18n
            cronet_aml_third_party_icu_icuuc_private
            cronet_aml_third_party_libevent_libevent
            cronet_aml_third_party_modp_b64_modp_b64
            jni_headers
            libandroid
            libc
            libclang_rt.builtins
            libcrypto
            libdl
            liblog
            libm
            libmediandk
            libssl
    )
endif ()
#################### cronet_aml_components_prefs_prefs_features ####################

#################### cronet_aml_third_party_abseil_cpp_absl_strings_cordz_handle ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/strings/internal/cordz_handle.cc
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
    add_library(cronet_aml_third_party_abseil_cpp_absl_strings_cordz_handle SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_abseil_cpp_absl_strings_cordz_handle
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_abseil_cpp_absl_strings_cordz_handle PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_abseil_cpp_absl_strings_cordz_handle
            cronet_aml_third_party_abseil_cpp_absl_base_base
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_malloc_internal
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait
            cronet_aml_third_party_abseil_cpp_absl_base_throw_delegate
            cronet_aml_third_party_abseil_cpp_absl_debugging_debugging_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_demangle_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_stacktrace
            cronet_aml_third_party_abseil_cpp_absl_debugging_symbolize
            cronet_aml_third_party_abseil_cpp_absl_numeric_int128
            cronet_aml_third_party_abseil_cpp_absl_strings_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_strings
            cronet_aml_third_party_abseil_cpp_absl_synchronization_graphcycles_internal
            cronet_aml_third_party_abseil_cpp_absl_synchronization_kernel_timeout_internal
            cronet_aml_third_party_abseil_cpp_absl_synchronization_synchronization
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_civil_time
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_time_zone
            cronet_aml_third_party_abseil_cpp_absl_time_time
            cronet_aml_third_party_abseil_cpp_absl_types_bad_optional_access
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_abseil_cpp_absl_strings_cordz_handle ####################

#################### cronet_aml_third_party_abseil_cpp_absl_crc_crc32c ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/crc/crc32c.cc
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/crc/internal/crc_memcpy_fallback.cc
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/crc/internal/crc_memcpy_x86_64.cc
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/crc/internal/crc_non_temporal_memcpy.cc
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
    add_library(cronet_aml_third_party_abseil_cpp_absl_crc_crc32c SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_abseil_cpp_absl_crc_crc32c
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_abseil_cpp_absl_crc_crc32c PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_abseil_cpp_absl_crc_crc32c
            cronet_aml_third_party_abseil_cpp_absl_base_base
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait
            cronet_aml_third_party_abseil_cpp_absl_base_throw_delegate
            cronet_aml_third_party_abseil_cpp_absl_crc_cpu_detect
            cronet_aml_third_party_abseil_cpp_absl_crc_crc_internal
            cronet_aml_third_party_abseil_cpp_absl_numeric_int128
            cronet_aml_third_party_abseil_cpp_absl_strings_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_strings
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_abseil_cpp_absl_crc_crc32c ####################

#################### cronet_aml_third_party_abseil_cpp_absl_container_hashtablez_sampler ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/container/internal/hashtablez_sampler.cc
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/container/internal/hashtablez_sampler_force_weak_definition.cc
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
    add_library(cronet_aml_third_party_abseil_cpp_absl_container_hashtablez_sampler SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_abseil_cpp_absl_container_hashtablez_sampler
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_abseil_cpp_absl_container_hashtablez_sampler PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_abseil_cpp_absl_container_hashtablez_sampler
            cronet_aml_third_party_abseil_cpp_absl_base_base
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_malloc_internal
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait
            cronet_aml_third_party_abseil_cpp_absl_base_throw_delegate
            cronet_aml_third_party_abseil_cpp_absl_debugging_debugging_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_demangle_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_stacktrace
            cronet_aml_third_party_abseil_cpp_absl_debugging_symbolize
            cronet_aml_third_party_abseil_cpp_absl_numeric_int128
            cronet_aml_third_party_abseil_cpp_absl_profiling_exponential_biased
            cronet_aml_third_party_abseil_cpp_absl_strings_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_strings
            cronet_aml_third_party_abseil_cpp_absl_synchronization_graphcycles_internal
            cronet_aml_third_party_abseil_cpp_absl_synchronization_kernel_timeout_internal
            cronet_aml_third_party_abseil_cpp_absl_synchronization_synchronization
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_civil_time
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_time_zone
            cronet_aml_third_party_abseil_cpp_absl_time_time
            cronet_aml_third_party_abseil_cpp_absl_types_bad_optional_access
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_abseil_cpp_absl_container_hashtablez_sampler ####################

#################### cronet_aml_third_party_abseil_cpp_absl_base_strerror ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/base/internal/strerror.cc
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
    add_library(cronet_aml_third_party_abseil_cpp_absl_base_strerror SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_abseil_cpp_absl_base_strerror
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_abseil_cpp_absl_base_strerror PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_abseil_cpp_absl_base_strerror
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_abseil_cpp_absl_base_strerror ####################

#################### cronet_aml_base_third_party_dynamic_annotations_dynamic_annotations ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/base/third_party/dynamic_annotations/dynamic_annotations.c
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
    add_library(cronet_aml_base_third_party_dynamic_annotations_dynamic_annotations SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_base_third_party_dynamic_annotations_dynamic_annotations
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_base_third_party_dynamic_annotations_dynamic_annotations PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_base_third_party_dynamic_annotations_dynamic_annotations
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_base_third_party_dynamic_annotations_dynamic_annotations ####################

#################### cronet_aml_third_party_abseil_cpp_absl_status_statusor ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/status/statusor.cc
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
    add_library(cronet_aml_third_party_abseil_cpp_absl_status_statusor SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_abseil_cpp_absl_status_statusor
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_abseil_cpp_absl_status_statusor PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_abseil_cpp_absl_status_statusor
            cronet_aml_third_party_abseil_cpp_absl_base_base
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_malloc_internal
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait
            cronet_aml_third_party_abseil_cpp_absl_base_strerror
            cronet_aml_third_party_abseil_cpp_absl_base_throw_delegate
            cronet_aml_third_party_abseil_cpp_absl_crc_cpu_detect
            cronet_aml_third_party_abseil_cpp_absl_crc_crc32c
            cronet_aml_third_party_abseil_cpp_absl_crc_crc_cord_state
            cronet_aml_third_party_abseil_cpp_absl_crc_crc_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_debugging_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_demangle_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_stacktrace
            cronet_aml_third_party_abseil_cpp_absl_debugging_symbolize
            cronet_aml_third_party_abseil_cpp_absl_numeric_int128
            cronet_aml_third_party_abseil_cpp_absl_profiling_exponential_biased
            cronet_aml_third_party_abseil_cpp_absl_status_status
            cronet_aml_third_party_abseil_cpp_absl_strings_cord
            cronet_aml_third_party_abseil_cpp_absl_strings_cord_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_functions
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_handle
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_info
            cronet_aml_third_party_abseil_cpp_absl_strings_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_str_format_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_strings
            cronet_aml_third_party_abseil_cpp_absl_synchronization_graphcycles_internal
            cronet_aml_third_party_abseil_cpp_absl_synchronization_kernel_timeout_internal
            cronet_aml_third_party_abseil_cpp_absl_synchronization_synchronization
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_civil_time
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_time_zone
            cronet_aml_third_party_abseil_cpp_absl_time_time
            cronet_aml_third_party_abseil_cpp_absl_types_bad_optional_access
            cronet_aml_third_party_abseil_cpp_absl_types_bad_variant_access
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_abseil_cpp_absl_status_statusor ####################

#################### cronet_aml_third_party_modp_b64_modp_b64 ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/modp_b64/modp_b64.cc
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
    add_library(cronet_aml_third_party_modp_b64_modp_b64 SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_modp_b64_modp_b64
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_modp_b64_modp_b64 PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_modp_b64_modp_b64
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_modp_b64_modp_b64 ####################

#################### cronet_aml_third_party_abseil_cpp_absl_random_seed_gen_exception ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/random/seed_gen_exception.cc
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
    add_library(cronet_aml_third_party_abseil_cpp_absl_random_seed_gen_exception SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_abseil_cpp_absl_random_seed_gen_exception
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_abseil_cpp_absl_random_seed_gen_exception PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_abseil_cpp_absl_random_seed_gen_exception
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_abseil_cpp_absl_random_seed_gen_exception ####################

#################### cronet_aml_third_party_ashmem_ashmem ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/ashmem/ashmem-dev.c
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
    add_library(cronet_aml_third_party_ashmem_ashmem SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_ashmem_ashmem
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_ashmem_ashmem PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_ashmem_ashmem
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_ashmem_ashmem ####################

#################### cronet_aml_third_party_abseil_cpp_absl_log_internal_log_message ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/log/internal/log_message.cc
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
    add_library(cronet_aml_third_party_abseil_cpp_absl_log_internal_log_message SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_abseil_cpp_absl_log_internal_log_message
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_abseil_cpp_absl_log_internal_log_message PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_abseil_cpp_absl_log_internal_log_message
            cronet_aml_third_party_abseil_cpp_absl_base_base
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_malloc_internal
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait
            cronet_aml_third_party_abseil_cpp_absl_base_strerror
            cronet_aml_third_party_abseil_cpp_absl_base_throw_delegate
            cronet_aml_third_party_abseil_cpp_absl_debugging_debugging_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_demangle_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_examine_stack
            cronet_aml_third_party_abseil_cpp_absl_debugging_stacktrace
            cronet_aml_third_party_abseil_cpp_absl_debugging_symbolize
            cronet_aml_third_party_abseil_cpp_absl_hash_city
            cronet_aml_third_party_abseil_cpp_absl_hash_hash
            cronet_aml_third_party_abseil_cpp_absl_hash_low_level_hash
            cronet_aml_third_party_abseil_cpp_absl_log_globals
            cronet_aml_third_party_abseil_cpp_absl_log_internal_format
            cronet_aml_third_party_abseil_cpp_absl_log_internal_globals
            cronet_aml_third_party_abseil_cpp_absl_log_internal_log_sink_set
            cronet_aml_third_party_abseil_cpp_absl_log_internal_nullguard
            cronet_aml_third_party_abseil_cpp_absl_log_internal_proto
            cronet_aml_third_party_abseil_cpp_absl_log_log_entry
            cronet_aml_third_party_abseil_cpp_absl_log_log_sink
            cronet_aml_third_party_abseil_cpp_absl_numeric_int128
            cronet_aml_third_party_abseil_cpp_absl_strings_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_str_format_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_strings
            cronet_aml_third_party_abseil_cpp_absl_synchronization_graphcycles_internal
            cronet_aml_third_party_abseil_cpp_absl_synchronization_kernel_timeout_internal
            cronet_aml_third_party_abseil_cpp_absl_synchronization_synchronization
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_civil_time
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_time_zone
            cronet_aml_third_party_abseil_cpp_absl_time_time
            cronet_aml_third_party_abseil_cpp_absl_types_bad_optional_access
            cronet_aml_third_party_abseil_cpp_absl_types_bad_variant_access
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            liblog
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_abseil_cpp_absl_log_internal_log_message ####################

#################### cronet_aml_base_third_party_double_conversion_double_conversion ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/base/third_party/double_conversion/double-conversion/bignum-dtoa.cc
        ${ANDROID_ROOT}/external/cronet/base/third_party/double_conversion/double-conversion/bignum.cc
        ${ANDROID_ROOT}/external/cronet/base/third_party/double_conversion/double-conversion/cached-powers.cc
        ${ANDROID_ROOT}/external/cronet/base/third_party/double_conversion/double-conversion/double-to-string.cc
        ${ANDROID_ROOT}/external/cronet/base/third_party/double_conversion/double-conversion/fast-dtoa.cc
        ${ANDROID_ROOT}/external/cronet/base/third_party/double_conversion/double-conversion/fixed-dtoa.cc
        ${ANDROID_ROOT}/external/cronet/base/third_party/double_conversion/double-conversion/string-to-double.cc
        ${ANDROID_ROOT}/external/cronet/base/third_party/double_conversion/double-conversion/strtod.cc
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
    add_library(cronet_aml_base_third_party_double_conversion_double_conversion SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_base_third_party_double_conversion_double_conversion
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_base_third_party_double_conversion_double_conversion PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_base_third_party_double_conversion_double_conversion
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_base_third_party_double_conversion_double_conversion ####################

#################### cronet_aml_third_party_abseil_cpp_absl_strings_cordz_info ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/strings/internal/cordz_info.cc
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
    add_library(cronet_aml_third_party_abseil_cpp_absl_strings_cordz_info SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_abseil_cpp_absl_strings_cordz_info
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_abseil_cpp_absl_strings_cordz_info PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_abseil_cpp_absl_strings_cordz_info
            cronet_aml_third_party_abseil_cpp_absl_base_base
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_malloc_internal
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait
            cronet_aml_third_party_abseil_cpp_absl_base_throw_delegate
            cronet_aml_third_party_abseil_cpp_absl_crc_cpu_detect
            cronet_aml_third_party_abseil_cpp_absl_crc_crc32c
            cronet_aml_third_party_abseil_cpp_absl_crc_crc_cord_state
            cronet_aml_third_party_abseil_cpp_absl_crc_crc_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_debugging_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_demangle_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_stacktrace
            cronet_aml_third_party_abseil_cpp_absl_debugging_symbolize
            cronet_aml_third_party_abseil_cpp_absl_numeric_int128
            cronet_aml_third_party_abseil_cpp_absl_profiling_exponential_biased
            cronet_aml_third_party_abseil_cpp_absl_strings_cord_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_functions
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_handle
            cronet_aml_third_party_abseil_cpp_absl_strings_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_strings
            cronet_aml_third_party_abseil_cpp_absl_synchronization_graphcycles_internal
            cronet_aml_third_party_abseil_cpp_absl_synchronization_kernel_timeout_internal
            cronet_aml_third_party_abseil_cpp_absl_synchronization_synchronization
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_civil_time
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_time_zone
            cronet_aml_third_party_abseil_cpp_absl_time_time
            cronet_aml_third_party_abseil_cpp_absl_types_bad_optional_access
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_abseil_cpp_absl_strings_cordz_info ####################

#################### cronet_aml_third_party_abseil_cpp_absl_numeric_int128 ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/numeric/int128.cc
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
    add_library(cronet_aml_third_party_abseil_cpp_absl_numeric_int128 SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_abseil_cpp_absl_numeric_int128
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_abseil_cpp_absl_numeric_int128 PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_abseil_cpp_absl_numeric_int128
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_abseil_cpp_absl_numeric_int128 ####################

#################### cronet_aml_third_party_abseil_cpp_absl_container_raw_hash_set ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/container/internal/raw_hash_set.cc
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
    add_library(cronet_aml_third_party_abseil_cpp_absl_container_raw_hash_set SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_abseil_cpp_absl_container_raw_hash_set
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_abseil_cpp_absl_container_raw_hash_set PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_abseil_cpp_absl_container_raw_hash_set
            cronet_aml_third_party_abseil_cpp_absl_base_base
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_malloc_internal
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait
            cronet_aml_third_party_abseil_cpp_absl_base_throw_delegate
            cronet_aml_third_party_abseil_cpp_absl_container_hashtablez_sampler
            cronet_aml_third_party_abseil_cpp_absl_debugging_debugging_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_demangle_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_stacktrace
            cronet_aml_third_party_abseil_cpp_absl_debugging_symbolize
            cronet_aml_third_party_abseil_cpp_absl_hash_city
            cronet_aml_third_party_abseil_cpp_absl_hash_hash
            cronet_aml_third_party_abseil_cpp_absl_hash_low_level_hash
            cronet_aml_third_party_abseil_cpp_absl_numeric_int128
            cronet_aml_third_party_abseil_cpp_absl_profiling_exponential_biased
            cronet_aml_third_party_abseil_cpp_absl_strings_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_strings
            cronet_aml_third_party_abseil_cpp_absl_synchronization_graphcycles_internal
            cronet_aml_third_party_abseil_cpp_absl_synchronization_kernel_timeout_internal
            cronet_aml_third_party_abseil_cpp_absl_synchronization_synchronization
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_civil_time
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_time_zone
            cronet_aml_third_party_abseil_cpp_absl_time_time
            cronet_aml_third_party_abseil_cpp_absl_types_bad_optional_access
            cronet_aml_third_party_abseil_cpp_absl_types_bad_variant_access
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_abseil_cpp_absl_container_raw_hash_set ####################

#################### cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_civil_time ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/time/internal/cctz/src/civil_time_detail.cc
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
    add_library(cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_civil_time SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_civil_time
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_civil_time PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_civil_time
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_civil_time ####################

#################### cronet_aml_third_party_abseil_cpp_absl_random_distributions ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/random/discrete_distribution.cc
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/random/gaussian_distribution.cc
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
    add_library(cronet_aml_third_party_abseil_cpp_absl_random_distributions SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_abseil_cpp_absl_random_distributions
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_abseil_cpp_absl_random_distributions PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_abseil_cpp_absl_random_distributions
            cronet_aml_third_party_abseil_cpp_absl_base_base
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait
            cronet_aml_third_party_abseil_cpp_absl_base_throw_delegate
            cronet_aml_third_party_abseil_cpp_absl_numeric_int128
            cronet_aml_third_party_abseil_cpp_absl_strings_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_strings
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_abseil_cpp_absl_random_distributions ####################

#################### cronet_aml_buildtools_third_party_libc__abi_libc__abi ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/src/abort_message.cpp
        ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/src/cxa_aux_runtime.cpp
        ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/src/cxa_default_handlers.cpp
        ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/src/cxa_exception.cpp
        ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/src/cxa_exception_storage.cpp
        ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/src/cxa_guard.cpp
        ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/src/cxa_handlers.cpp
        ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/src/cxa_personality.cpp
        ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/src/cxa_thread_atexit.cpp
        ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/src/cxa_vector.cpp
        ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/src/cxa_virtual.cpp
        ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/src/fallback_malloc.cpp
        ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/src/private_typeinfo.cpp
        ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/src/stdlib_exception.cpp
        ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/src/stdlib_stdexcept.cpp
        ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/src/stdlib_typeinfo.cpp
        ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/cxa_demangle_stub.cc
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
    add_library(cronet_aml_buildtools_third_party_libc__abi_libc__abi SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_buildtools_third_party_libc__abi_libc__abi
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/src
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_buildtools_third_party_libc__abi_libc__abi PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_buildtools_third_party_libc__abi_libc__abi
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_buildtools_third_party_libc__abi_libc__abi ####################

#################### cronet_aml_third_party_abseil_cpp_absl_log_log_entry ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/log/log_entry.cc
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
    add_library(cronet_aml_third_party_abseil_cpp_absl_log_log_entry SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_abseil_cpp_absl_log_log_entry
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_abseil_cpp_absl_log_log_entry PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_abseil_cpp_absl_log_log_entry
            cronet_aml_third_party_abseil_cpp_absl_base_base
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait
            cronet_aml_third_party_abseil_cpp_absl_base_throw_delegate
            cronet_aml_third_party_abseil_cpp_absl_numeric_int128
            cronet_aml_third_party_abseil_cpp_absl_strings_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_strings
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_civil_time
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_time_zone
            cronet_aml_third_party_abseil_cpp_absl_time_time
            cronet_aml_third_party_abseil_cpp_absl_types_bad_optional_access
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_abseil_cpp_absl_log_log_entry ####################

#################### cronet_aml_third_party_abseil_cpp_absl_base_log_severity ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/base/log_severity.cc
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
    add_library(cronet_aml_third_party_abseil_cpp_absl_base_log_severity SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_abseil_cpp_absl_base_log_severity PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_abseil_cpp_absl_base_log_severity ####################

#################### cronet_aml_components_prefs_prefs ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/components/prefs/android/pref_service_android.cc
        ${ANDROID_ROOT}/external/cronet/components/prefs/command_line_pref_store.cc
        ${ANDROID_ROOT}/external/cronet/components/prefs/default_pref_store.cc
        ${ANDROID_ROOT}/external/cronet/components/prefs/in_memory_pref_store.cc
        ${ANDROID_ROOT}/external/cronet/components/prefs/json_pref_store.cc
        ${ANDROID_ROOT}/external/cronet/components/prefs/overlay_user_pref_store.cc
        ${ANDROID_ROOT}/external/cronet/components/prefs/persistent_pref_store.cc
        ${ANDROID_ROOT}/external/cronet/components/prefs/pref_change_registrar.cc
        ${ANDROID_ROOT}/external/cronet/components/prefs/pref_member.cc
        ${ANDROID_ROOT}/external/cronet/components/prefs/pref_notifier_impl.cc
        ${ANDROID_ROOT}/external/cronet/components/prefs/pref_registry.cc
        ${ANDROID_ROOT}/external/cronet/components/prefs/pref_registry_simple.cc
        ${ANDROID_ROOT}/external/cronet/components/prefs/pref_service.cc
        ${ANDROID_ROOT}/external/cronet/components/prefs/pref_service_factory.cc
        ${ANDROID_ROOT}/external/cronet/components/prefs/pref_store.cc
        ${ANDROID_ROOT}/external/cronet/components/prefs/pref_value_map.cc
        ${ANDROID_ROOT}/external/cronet/components/prefs/pref_value_store.cc
        ${ANDROID_ROOT}/external/cronet/components/prefs/scoped_user_pref_update.cc
        ${ANDROID_ROOT}/external/cronet/components/prefs/segregated_pref_store.cc
        ${ANDROID_ROOT}/external/cronet/components/prefs/value_map_pref_store.cc
        ${ANDROID_ROOT}/external/cronet/components/prefs/writeable_pref_store.cc
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
    add_library(cronet_aml_components_prefs_prefs SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_components_prefs_prefs
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
            ${ANDROID_ROOT}/external/cronet/third_party/boringssl/src/include
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_chromecast_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_chromeos_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_debugging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_logging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_partition_alloc_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_anchor_functions_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_android_runtime_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_android_runtime_unchecked_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_base_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_build_date/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_cfi_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_clang_profiling_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_debugging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_feature_list_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_ios_cronet_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_logging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_message_pump_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_orderfile_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_power_monitor_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_profiler_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_rust_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_sanitizer_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_synchronization_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_tracing_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_blink_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_branding_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_chromecast_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_chromeos_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_config_compiler_compiler_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_components_prefs_android_jni_headers/android_${OUT_ARCH_CPU}/gen
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_components_prefs_prefs PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_components_prefs_prefs
            cronet_aml_base_allocator_partition_allocator_partition_alloc
            cronet_aml_base_base
            cronet_aml_base_base_static
            cronet_aml_base_nodebug_assertion
            cronet_aml_base_third_party_double_conversion_double_conversion
            cronet_aml_base_third_party_dynamic_annotations_dynamic_annotations
            cronet_aml_components_prefs_prefs_features
            cronet_aml_third_party_abseil_cpp_absl_base_base
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_malloc_internal
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait
            cronet_aml_third_party_abseil_cpp_absl_base_strerror
            cronet_aml_third_party_abseil_cpp_absl_base_throw_delegate
            cronet_aml_third_party_abseil_cpp_absl_container_hashtablez_sampler
            cronet_aml_third_party_abseil_cpp_absl_container_raw_hash_set
            cronet_aml_third_party_abseil_cpp_absl_crc_cpu_detect
            cronet_aml_third_party_abseil_cpp_absl_crc_crc32c
            cronet_aml_third_party_abseil_cpp_absl_crc_crc_cord_state
            cronet_aml_third_party_abseil_cpp_absl_crc_crc_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_debugging_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_demangle_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_examine_stack
            cronet_aml_third_party_abseil_cpp_absl_debugging_failure_signal_handler
            cronet_aml_third_party_abseil_cpp_absl_debugging_stacktrace
            cronet_aml_third_party_abseil_cpp_absl_debugging_symbolize
            cronet_aml_third_party_abseil_cpp_absl_hash_city
            cronet_aml_third_party_abseil_cpp_absl_hash_hash
            cronet_aml_third_party_abseil_cpp_absl_hash_low_level_hash
            cronet_aml_third_party_abseil_cpp_absl_log_die_if_null
            cronet_aml_third_party_abseil_cpp_absl_log_globals
            cronet_aml_third_party_abseil_cpp_absl_log_internal_check_op
            cronet_aml_third_party_abseil_cpp_absl_log_internal_conditions
            cronet_aml_third_party_abseil_cpp_absl_log_internal_format
            cronet_aml_third_party_abseil_cpp_absl_log_internal_globals
            cronet_aml_third_party_abseil_cpp_absl_log_internal_log_message
            cronet_aml_third_party_abseil_cpp_absl_log_internal_log_sink_set
            cronet_aml_third_party_abseil_cpp_absl_log_internal_nullguard
            cronet_aml_third_party_abseil_cpp_absl_log_internal_proto
            cronet_aml_third_party_abseil_cpp_absl_log_log_entry
            cronet_aml_third_party_abseil_cpp_absl_log_log_sink
            cronet_aml_third_party_abseil_cpp_absl_numeric_int128
            cronet_aml_third_party_abseil_cpp_absl_profiling_exponential_biased
            cronet_aml_third_party_abseil_cpp_absl_random_distributions
            cronet_aml_third_party_abseil_cpp_absl_random_internal_platform
            cronet_aml_third_party_abseil_cpp_absl_random_internal_pool_urbg
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_hwaes
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_hwaes_impl
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_slow
            cronet_aml_third_party_abseil_cpp_absl_random_internal_seed_material
            cronet_aml_third_party_abseil_cpp_absl_random_seed_gen_exception
            cronet_aml_third_party_abseil_cpp_absl_random_seed_sequences
            cronet_aml_third_party_abseil_cpp_absl_status_status
            cronet_aml_third_party_abseil_cpp_absl_status_statusor
            cronet_aml_third_party_abseil_cpp_absl_strings_cord
            cronet_aml_third_party_abseil_cpp_absl_strings_cord_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_functions
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_handle
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_info
            cronet_aml_third_party_abseil_cpp_absl_strings_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_str_format_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_strings
            cronet_aml_third_party_abseil_cpp_absl_synchronization_graphcycles_internal
            cronet_aml_third_party_abseil_cpp_absl_synchronization_kernel_timeout_internal
            cronet_aml_third_party_abseil_cpp_absl_synchronization_synchronization
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_civil_time
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_time_zone
            cronet_aml_third_party_abseil_cpp_absl_time_time
            cronet_aml_third_party_abseil_cpp_absl_types_bad_optional_access
            cronet_aml_third_party_abseil_cpp_absl_types_bad_variant_access
            cronet_aml_third_party_android_ndk_cpu_features
            cronet_aml_third_party_ashmem_ashmem
            cronet_aml_third_party_icu_icui18n
            cronet_aml_third_party_icu_icuuc_private
            cronet_aml_third_party_libevent_libevent
            cronet_aml_third_party_modp_b64_modp_b64
            jni_headers
            libandroid
            libc
            libclang_rt.builtins
            libcrypto
            libdl
            liblog
            libm
            libmediandk
            libssl
    )
endif ()
#################### cronet_aml_components_prefs_prefs ####################

#################### cronet_aml_third_party_abseil_cpp_absl_crc_crc_internal ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/crc/internal/crc.cc
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/crc/internal/crc_x86_arm_combined.cc
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
    add_library(cronet_aml_third_party_abseil_cpp_absl_crc_crc_internal SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_abseil_cpp_absl_crc_crc_internal
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_abseil_cpp_absl_crc_crc_internal PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_abseil_cpp_absl_crc_crc_internal
            cronet_aml_third_party_abseil_cpp_absl_base_base
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait
            cronet_aml_third_party_abseil_cpp_absl_crc_cpu_detect
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_abseil_cpp_absl_crc_crc_internal ####################

#################### cronet_aml_base_base ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/base/allocator/allocator_check.cc
        ${ANDROID_ROOT}/external/cronet/base/allocator/allocator_extension.cc
        ${ANDROID_ROOT}/external/cronet/base/allocator/dispatcher/dispatcher.cc
        ${ANDROID_ROOT}/external/cronet/base/allocator/dispatcher/internal/dispatch_data.cc
        ${ANDROID_ROOT}/external/cronet/base/allocator/dispatcher/reentry_guard.cc
        ${ANDROID_ROOT}/external/cronet/base/allocator/dispatcher/tls.cc
        ${ANDROID_ROOT}/external/cronet/base/allocator/partition_allocator/shim/allocator_shim_default_dispatch_to_linker_wrapped_symbols.cc
        ${ANDROID_ROOT}/external/cronet/base/android/android_hardware_buffer_compat.cc
        ${ANDROID_ROOT}/external/cronet/base/android/android_image_reader_compat.cc
        ${ANDROID_ROOT}/external/cronet/base/android/apk_assets.cc
        ${ANDROID_ROOT}/external/cronet/base/android/application_status_listener.cc
        ${ANDROID_ROOT}/external/cronet/base/android/base_feature_list.cc
        ${ANDROID_ROOT}/external/cronet/base/android/base_features.cc
        ${ANDROID_ROOT}/external/cronet/base/android/base_jni_onload.cc
        ${ANDROID_ROOT}/external/cronet/base/android/build_info.cc
        ${ANDROID_ROOT}/external/cronet/base/android/bundle_utils.cc
        ${ANDROID_ROOT}/external/cronet/base/android/callback_android.cc
        ${ANDROID_ROOT}/external/cronet/base/android/child_process_service.cc
        ${ANDROID_ROOT}/external/cronet/base/android/command_line_android.cc
        ${ANDROID_ROOT}/external/cronet/base/android/content_uri_utils.cc
        ${ANDROID_ROOT}/external/cronet/base/android/cpu_features.cc
        ${ANDROID_ROOT}/external/cronet/base/android/early_trace_event_binding.cc
        ${ANDROID_ROOT}/external/cronet/base/android/event_log.cc
        ${ANDROID_ROOT}/external/cronet/base/android/feature_list_jni.cc
        ${ANDROID_ROOT}/external/cronet/base/android/features_jni.cc
        ${ANDROID_ROOT}/external/cronet/base/android/field_trial_list.cc
        ${ANDROID_ROOT}/external/cronet/base/android/important_file_writer_android.cc
        ${ANDROID_ROOT}/external/cronet/base/android/int_string_callback.cc
        ${ANDROID_ROOT}/external/cronet/base/android/java_exception_reporter.cc
        ${ANDROID_ROOT}/external/cronet/base/android/java_handler_thread.cc
        ${ANDROID_ROOT}/external/cronet/base/android/java_heap_dump_generator.cc
        ${ANDROID_ROOT}/external/cronet/base/android/java_runtime.cc
        ${ANDROID_ROOT}/external/cronet/base/android/jni_android.cc
        ${ANDROID_ROOT}/external/cronet/base/android/jni_array.cc
        ${ANDROID_ROOT}/external/cronet/base/android/jni_registrar.cc
        ${ANDROID_ROOT}/external/cronet/base/android/jni_string.cc
        ${ANDROID_ROOT}/external/cronet/base/android/jni_utils.cc
        ${ANDROID_ROOT}/external/cronet/base/android/jni_weak_ref.cc
        ${ANDROID_ROOT}/external/cronet/base/android/library_loader/anchor_functions.cc
        ${ANDROID_ROOT}/external/cronet/base/android/library_loader/library_loader_hooks.cc
        ${ANDROID_ROOT}/external/cronet/base/android/library_loader/library_prefetcher.cc
        ${ANDROID_ROOT}/external/cronet/base/android/library_loader/library_prefetcher_hooks.cc
        ${ANDROID_ROOT}/external/cronet/base/android/locale_utils.cc
        ${ANDROID_ROOT}/external/cronet/base/android/meminfo_dump_provider.cc
        ${ANDROID_ROOT}/external/cronet/base/android/memory_pressure_listener_android.cc
        ${ANDROID_ROOT}/external/cronet/base/android/native_uma_recorder.cc
        ${ANDROID_ROOT}/external/cronet/base/android/path_service_android.cc
        ${ANDROID_ROOT}/external/cronet/base/android/path_utils.cc
        ${ANDROID_ROOT}/external/cronet/base/android/radio_utils.cc
        ${ANDROID_ROOT}/external/cronet/base/android/reached_addresses_bitset.cc
        ${ANDROID_ROOT}/external/cronet/base/android/remove_stale_data.cc
        ${ANDROID_ROOT}/external/cronet/base/android/scoped_hardware_buffer_fence_sync.cc
        ${ANDROID_ROOT}/external/cronet/base/android/scoped_hardware_buffer_handle.cc
        ${ANDROID_ROOT}/external/cronet/base/android/scoped_java_ref.cc
        ${ANDROID_ROOT}/external/cronet/base/android/statistics_recorder_android.cc
        ${ANDROID_ROOT}/external/cronet/base/android/sys_utils.cc
        ${ANDROID_ROOT}/external/cronet/base/android/task_scheduler/post_task_android.cc
        ${ANDROID_ROOT}/external/cronet/base/android/task_scheduler/task_runner_android.cc
        ${ANDROID_ROOT}/external/cronet/base/android/thread_instruction_count.cc
        ${ANDROID_ROOT}/external/cronet/base/android/timezone_utils.cc
        ${ANDROID_ROOT}/external/cronet/base/android/trace_event_binding.cc
        ${ANDROID_ROOT}/external/cronet/base/android/unguessable_token_android.cc
        ${ANDROID_ROOT}/external/cronet/base/at_exit.cc
        ${ANDROID_ROOT}/external/cronet/base/barrier_closure.cc
        ${ANDROID_ROOT}/external/cronet/base/base64.cc
        ${ANDROID_ROOT}/external/cronet/base/base64url.cc
        ${ANDROID_ROOT}/external/cronet/base/base_paths.cc
        ${ANDROID_ROOT}/external/cronet/base/base_paths_android.cc
        ${ANDROID_ROOT}/external/cronet/base/big_endian.cc
        ${ANDROID_ROOT}/external/cronet/base/callback_list.cc
        ${ANDROID_ROOT}/external/cronet/base/check.cc
        ${ANDROID_ROOT}/external/cronet/base/check_is_test.cc
        ${ANDROID_ROOT}/external/cronet/base/check_op.cc
        ${ANDROID_ROOT}/external/cronet/base/command_line.cc
        ${ANDROID_ROOT}/external/cronet/base/containers/flat_tree.cc
        ${ANDROID_ROOT}/external/cronet/base/containers/intrusive_heap.cc
        ${ANDROID_ROOT}/external/cronet/base/containers/linked_list.cc
        ${ANDROID_ROOT}/external/cronet/base/cpu.cc
        ${ANDROID_ROOT}/external/cronet/base/cpu_reduction_experiment.cc
        ${ANDROID_ROOT}/external/cronet/base/debug/alias.cc
        ${ANDROID_ROOT}/external/cronet/base/debug/asan_invalid_access.cc
        ${ANDROID_ROOT}/external/cronet/base/debug/asan_service.cc
        ${ANDROID_ROOT}/external/cronet/base/debug/buffered_dwarf_reader.cc
        ${ANDROID_ROOT}/external/cronet/base/debug/crash_logging.cc
        ${ANDROID_ROOT}/external/cronet/base/debug/debugger.cc
        ${ANDROID_ROOT}/external/cronet/base/debug/debugger_posix.cc
        ${ANDROID_ROOT}/external/cronet/base/debug/dump_without_crashing.cc
        ${ANDROID_ROOT}/external/cronet/base/debug/dwarf_line_no.cc
        ${ANDROID_ROOT}/external/cronet/base/debug/elf_reader.cc
        ${ANDROID_ROOT}/external/cronet/base/debug/proc_maps_linux.cc
        ${ANDROID_ROOT}/external/cronet/base/debug/profiler.cc
        ${ANDROID_ROOT}/external/cronet/base/debug/stack_trace.cc
        ${ANDROID_ROOT}/external/cronet/base/debug/stack_trace_android.cc
        ${ANDROID_ROOT}/external/cronet/base/debug/task_trace.cc
        ${ANDROID_ROOT}/external/cronet/base/environment.cc
        ${ANDROID_ROOT}/external/cronet/base/feature_list.cc
        ${ANDROID_ROOT}/external/cronet/base/features.cc
        ${ANDROID_ROOT}/external/cronet/base/file_descriptor_posix.cc
        ${ANDROID_ROOT}/external/cronet/base/file_descriptor_store.cc
        ${ANDROID_ROOT}/external/cronet/base/files/file.cc
        ${ANDROID_ROOT}/external/cronet/base/files/file_descriptor_watcher_posix.cc
        ${ANDROID_ROOT}/external/cronet/base/files/file_enumerator.cc
        ${ANDROID_ROOT}/external/cronet/base/files/file_enumerator_posix.cc
        ${ANDROID_ROOT}/external/cronet/base/files/file_path.cc
        ${ANDROID_ROOT}/external/cronet/base/files/file_path_watcher.cc
        ${ANDROID_ROOT}/external/cronet/base/files/file_path_watcher_inotify.cc
        ${ANDROID_ROOT}/external/cronet/base/files/file_posix.cc
        ${ANDROID_ROOT}/external/cronet/base/files/file_proxy.cc
        ${ANDROID_ROOT}/external/cronet/base/files/file_tracing.cc
        ${ANDROID_ROOT}/external/cronet/base/files/file_util.cc
        ${ANDROID_ROOT}/external/cronet/base/files/file_util_android.cc
        ${ANDROID_ROOT}/external/cronet/base/files/file_util_posix.cc
        ${ANDROID_ROOT}/external/cronet/base/files/important_file_writer.cc
        ${ANDROID_ROOT}/external/cronet/base/files/important_file_writer_cleaner.cc
        ${ANDROID_ROOT}/external/cronet/base/files/memory_mapped_file.cc
        ${ANDROID_ROOT}/external/cronet/base/files/memory_mapped_file_posix.cc
        ${ANDROID_ROOT}/external/cronet/base/files/safe_base_name.cc
        ${ANDROID_ROOT}/external/cronet/base/files/scoped_file.cc
        ${ANDROID_ROOT}/external/cronet/base/files/scoped_file_android.cc
        ${ANDROID_ROOT}/external/cronet/base/files/scoped_temp_dir.cc
        ${ANDROID_ROOT}/external/cronet/base/functional/callback_helpers.cc
        ${ANDROID_ROOT}/external/cronet/base/functional/callback_internal.cc
        ${ANDROID_ROOT}/external/cronet/base/hash/hash.cc
        ${ANDROID_ROOT}/external/cronet/base/hash/legacy_hash.cc
        ${ANDROID_ROOT}/external/cronet/base/hash/md5_boringssl.cc
        ${ANDROID_ROOT}/external/cronet/base/hash/sha1_boringssl.cc
        ${ANDROID_ROOT}/external/cronet/base/json/json_file_value_serializer.cc
        ${ANDROID_ROOT}/external/cronet/base/json/json_parser.cc
        ${ANDROID_ROOT}/external/cronet/base/json/json_reader.cc
        ${ANDROID_ROOT}/external/cronet/base/json/json_string_value_serializer.cc
        ${ANDROID_ROOT}/external/cronet/base/json/json_value_converter.cc
        ${ANDROID_ROOT}/external/cronet/base/json/json_writer.cc
        ${ANDROID_ROOT}/external/cronet/base/json/string_escape.cc
        ${ANDROID_ROOT}/external/cronet/base/json/values_util.cc
        ${ANDROID_ROOT}/external/cronet/base/lazy_instance_helpers.cc
        ${ANDROID_ROOT}/external/cronet/base/linux_util.cc
        ${ANDROID_ROOT}/external/cronet/base/location.cc
        ${ANDROID_ROOT}/external/cronet/base/logging.cc
        ${ANDROID_ROOT}/external/cronet/base/memory/aligned_memory.cc
        ${ANDROID_ROOT}/external/cronet/base/memory/discardable_memory.cc
        ${ANDROID_ROOT}/external/cronet/base/memory/discardable_memory_allocator.cc
        ${ANDROID_ROOT}/external/cronet/base/memory/discardable_shared_memory.cc
        ${ANDROID_ROOT}/external/cronet/base/memory/madv_free_discardable_memory_allocator_posix.cc
        ${ANDROID_ROOT}/external/cronet/base/memory/madv_free_discardable_memory_posix.cc
        ${ANDROID_ROOT}/external/cronet/base/memory/memory_pressure_listener.cc
        ${ANDROID_ROOT}/external/cronet/base/memory/memory_pressure_monitor.cc
        ${ANDROID_ROOT}/external/cronet/base/memory/nonscannable_memory.cc
        ${ANDROID_ROOT}/external/cronet/base/memory/page_size_posix.cc
        ${ANDROID_ROOT}/external/cronet/base/memory/platform_shared_memory_handle.cc
        ${ANDROID_ROOT}/external/cronet/base/memory/platform_shared_memory_mapper_android.cc
        ${ANDROID_ROOT}/external/cronet/base/memory/platform_shared_memory_region.cc
        ${ANDROID_ROOT}/external/cronet/base/memory/platform_shared_memory_region_android.cc
        ${ANDROID_ROOT}/external/cronet/base/memory/raw_ptr_asan_bound_arg_tracker.cc
        ${ANDROID_ROOT}/external/cronet/base/memory/raw_ptr_asan_hooks.cc
        ${ANDROID_ROOT}/external/cronet/base/memory/raw_ptr_asan_service.cc
        ${ANDROID_ROOT}/external/cronet/base/memory/read_only_shared_memory_region.cc
        ${ANDROID_ROOT}/external/cronet/base/memory/ref_counted.cc
        ${ANDROID_ROOT}/external/cronet/base/memory/ref_counted_memory.cc
        ${ANDROID_ROOT}/external/cronet/base/memory/shared_memory_mapper.cc
        ${ANDROID_ROOT}/external/cronet/base/memory/shared_memory_mapping.cc
        ${ANDROID_ROOT}/external/cronet/base/memory/shared_memory_security_policy.cc
        ${ANDROID_ROOT}/external/cronet/base/memory/shared_memory_tracker.cc
        ${ANDROID_ROOT}/external/cronet/base/memory/unsafe_shared_memory_pool.cc
        ${ANDROID_ROOT}/external/cronet/base/memory/unsafe_shared_memory_region.cc
        ${ANDROID_ROOT}/external/cronet/base/memory/weak_ptr.cc
        ${ANDROID_ROOT}/external/cronet/base/memory/writable_shared_memory_region.cc
        ${ANDROID_ROOT}/external/cronet/base/message_loop/message_pump.cc
        ${ANDROID_ROOT}/external/cronet/base/message_loop/message_pump_android.cc
        ${ANDROID_ROOT}/external/cronet/base/message_loop/message_pump_default.cc
        ${ANDROID_ROOT}/external/cronet/base/message_loop/message_pump_epoll.cc
        ${ANDROID_ROOT}/external/cronet/base/message_loop/message_pump_libevent.cc
        ${ANDROID_ROOT}/external/cronet/base/message_loop/watchable_io_message_pump_posix.cc
        ${ANDROID_ROOT}/external/cronet/base/message_loop/work_id_provider.cc
        ${ANDROID_ROOT}/external/cronet/base/metrics/bucket_ranges.cc
        ${ANDROID_ROOT}/external/cronet/base/metrics/crc32.cc
        ${ANDROID_ROOT}/external/cronet/base/metrics/dummy_histogram.cc
        ${ANDROID_ROOT}/external/cronet/base/metrics/field_trial.cc
        ${ANDROID_ROOT}/external/cronet/base/metrics/field_trial_list_including_low_anonymity.cc
        ${ANDROID_ROOT}/external/cronet/base/metrics/field_trial_param_associator.cc
        ${ANDROID_ROOT}/external/cronet/base/metrics/field_trial_params.cc
        ${ANDROID_ROOT}/external/cronet/base/metrics/histogram.cc
        ${ANDROID_ROOT}/external/cronet/base/metrics/histogram_base.cc
        ${ANDROID_ROOT}/external/cronet/base/metrics/histogram_delta_serialization.cc
        ${ANDROID_ROOT}/external/cronet/base/metrics/histogram_functions.cc
        ${ANDROID_ROOT}/external/cronet/base/metrics/histogram_samples.cc
        ${ANDROID_ROOT}/external/cronet/base/metrics/histogram_snapshot_manager.cc
        ${ANDROID_ROOT}/external/cronet/base/metrics/metrics_hashes.cc
        ${ANDROID_ROOT}/external/cronet/base/metrics/persistent_histogram_allocator.cc
        ${ANDROID_ROOT}/external/cronet/base/metrics/persistent_histogram_storage.cc
        ${ANDROID_ROOT}/external/cronet/base/metrics/persistent_memory_allocator.cc
        ${ANDROID_ROOT}/external/cronet/base/metrics/persistent_sample_map.cc
        ${ANDROID_ROOT}/external/cronet/base/metrics/ranges_manager.cc
        ${ANDROID_ROOT}/external/cronet/base/metrics/sample_map.cc
        ${ANDROID_ROOT}/external/cronet/base/metrics/sample_vector.cc
        ${ANDROID_ROOT}/external/cronet/base/metrics/single_sample_metrics.cc
        ${ANDROID_ROOT}/external/cronet/base/metrics/sparse_histogram.cc
        ${ANDROID_ROOT}/external/cronet/base/metrics/statistics_recorder.cc
        ${ANDROID_ROOT}/external/cronet/base/metrics/user_metrics.cc
        ${ANDROID_ROOT}/external/cronet/base/native_library.cc
        ${ANDROID_ROOT}/external/cronet/base/native_library_posix.cc
        ${ANDROID_ROOT}/external/cronet/base/observer_list_internal.cc
        ${ANDROID_ROOT}/external/cronet/base/observer_list_threadsafe.cc
        ${ANDROID_ROOT}/external/cronet/base/observer_list_types.cc
        ${ANDROID_ROOT}/external/cronet/base/one_shot_event.cc
        ${ANDROID_ROOT}/external/cronet/base/os_compat_android.cc
        ${ANDROID_ROOT}/external/cronet/base/path_service.cc
        ${ANDROID_ROOT}/external/cronet/base/pending_task.cc
        ${ANDROID_ROOT}/external/cronet/base/pickle.cc
        ${ANDROID_ROOT}/external/cronet/base/posix/can_lower_nice_to.cc
        ${ANDROID_ROOT}/external/cronet/base/posix/file_descriptor_shuffle.cc
        ${ANDROID_ROOT}/external/cronet/base/posix/global_descriptors.cc
        ${ANDROID_ROOT}/external/cronet/base/posix/safe_strerror.cc
        ${ANDROID_ROOT}/external/cronet/base/posix/unix_domain_socket.cc
        ${ANDROID_ROOT}/external/cronet/base/power_monitor/battery_level_provider.cc
        ${ANDROID_ROOT}/external/cronet/base/power_monitor/battery_state_sampler.cc
        ${ANDROID_ROOT}/external/cronet/base/power_monitor/moving_average.cc
        ${ANDROID_ROOT}/external/cronet/base/power_monitor/power_monitor.cc
        ${ANDROID_ROOT}/external/cronet/base/power_monitor/power_monitor_device_source.cc
        ${ANDROID_ROOT}/external/cronet/base/power_monitor/power_monitor_device_source_android.cc
        ${ANDROID_ROOT}/external/cronet/base/power_monitor/power_monitor_features.cc
        ${ANDROID_ROOT}/external/cronet/base/power_monitor/power_monitor_source.cc
        ${ANDROID_ROOT}/external/cronet/base/power_monitor/sampling_event_source.cc
        ${ANDROID_ROOT}/external/cronet/base/power_monitor/timer_sampling_event_source.cc
        ${ANDROID_ROOT}/external/cronet/base/process/current_process.cc
        ${ANDROID_ROOT}/external/cronet/base/process/environment_internal.cc
        ${ANDROID_ROOT}/external/cronet/base/process/internal_linux.cc
        ${ANDROID_ROOT}/external/cronet/base/process/kill.cc
        ${ANDROID_ROOT}/external/cronet/base/process/kill_posix.cc
        ${ANDROID_ROOT}/external/cronet/base/process/launch.cc
        ${ANDROID_ROOT}/external/cronet/base/process/launch_posix.cc
        ${ANDROID_ROOT}/external/cronet/base/process/memory.cc
        ${ANDROID_ROOT}/external/cronet/base/process/memory_linux.cc
        ${ANDROID_ROOT}/external/cronet/base/process/process_android.cc
        ${ANDROID_ROOT}/external/cronet/base/process/process_handle.cc
        ${ANDROID_ROOT}/external/cronet/base/process/process_handle_linux.cc
        ${ANDROID_ROOT}/external/cronet/base/process/process_handle_posix.cc
        ${ANDROID_ROOT}/external/cronet/base/process/process_iterator.cc
        ${ANDROID_ROOT}/external/cronet/base/process/process_iterator_linux.cc
        ${ANDROID_ROOT}/external/cronet/base/process/process_metrics.cc
        ${ANDROID_ROOT}/external/cronet/base/process/process_metrics_linux.cc
        ${ANDROID_ROOT}/external/cronet/base/process/process_metrics_posix.cc
        ${ANDROID_ROOT}/external/cronet/base/process/process_posix.cc
        ${ANDROID_ROOT}/external/cronet/base/profiler/frame.cc
        ${ANDROID_ROOT}/external/cronet/base/profiler/metadata_recorder.cc
        ${ANDROID_ROOT}/external/cronet/base/profiler/module_cache.cc
        ${ANDROID_ROOT}/external/cronet/base/profiler/module_cache_posix.cc
        ${ANDROID_ROOT}/external/cronet/base/profiler/sample_metadata.cc
        ${ANDROID_ROOT}/external/cronet/base/profiler/sampling_profiler_thread_token.cc
        ${ANDROID_ROOT}/external/cronet/base/profiler/stack_base_address_posix.cc
        ${ANDROID_ROOT}/external/cronet/base/profiler/stack_buffer.cc
        ${ANDROID_ROOT}/external/cronet/base/profiler/stack_copier.cc
        ${ANDROID_ROOT}/external/cronet/base/profiler/stack_copier_signal.cc
        ${ANDROID_ROOT}/external/cronet/base/profiler/stack_copier_suspend.cc
        ${ANDROID_ROOT}/external/cronet/base/profiler/stack_sampler.cc
        ${ANDROID_ROOT}/external/cronet/base/profiler/stack_sampler_android.cc
        ${ANDROID_ROOT}/external/cronet/base/profiler/stack_sampling_profiler.cc
        ${ANDROID_ROOT}/external/cronet/base/profiler/thread_delegate_posix.cc
        ${ANDROID_ROOT}/external/cronet/base/profiler/unwinder.cc
        ${ANDROID_ROOT}/external/cronet/base/rand_util.cc
        ${ANDROID_ROOT}/external/cronet/base/rand_util_posix.cc
        ${ANDROID_ROOT}/external/cronet/base/run_loop.cc
        ${ANDROID_ROOT}/external/cronet/base/sampling_heap_profiler/lock_free_address_hash_set.cc
        ${ANDROID_ROOT}/external/cronet/base/sampling_heap_profiler/poisson_allocation_sampler.cc
        ${ANDROID_ROOT}/external/cronet/base/sampling_heap_profiler/sampling_heap_profiler.cc
        ${ANDROID_ROOT}/external/cronet/base/scoped_add_feature_flags.cc
        ${ANDROID_ROOT}/external/cronet/base/scoped_environment_variable_override.cc
        ${ANDROID_ROOT}/external/cronet/base/scoped_native_library.cc
        ${ANDROID_ROOT}/external/cronet/base/sequence_checker.cc
        ${ANDROID_ROOT}/external/cronet/base/sequence_checker_impl.cc
        ${ANDROID_ROOT}/external/cronet/base/sequence_token.cc
        ${ANDROID_ROOT}/external/cronet/base/strings/abseil_string_number_conversions.cc
        ${ANDROID_ROOT}/external/cronet/base/strings/escape.cc
        ${ANDROID_ROOT}/external/cronet/base/strings/latin1_string_conversions.cc
        ${ANDROID_ROOT}/external/cronet/base/strings/pattern.cc
        ${ANDROID_ROOT}/external/cronet/base/strings/safe_sprintf.cc
        ${ANDROID_ROOT}/external/cronet/base/strings/strcat.cc
        ${ANDROID_ROOT}/external/cronet/base/strings/string_number_conversions.cc
        ${ANDROID_ROOT}/external/cronet/base/strings/string_piece.cc
        ${ANDROID_ROOT}/external/cronet/base/strings/string_split.cc
        ${ANDROID_ROOT}/external/cronet/base/strings/string_util.cc
        ${ANDROID_ROOT}/external/cronet/base/strings/string_util_constants.cc
        ${ANDROID_ROOT}/external/cronet/base/strings/stringprintf.cc
        ${ANDROID_ROOT}/external/cronet/base/strings/sys_string_conversions_posix.cc
        ${ANDROID_ROOT}/external/cronet/base/strings/utf_offset_string_conversions.cc
        ${ANDROID_ROOT}/external/cronet/base/strings/utf_string_conversion_utils.cc
        ${ANDROID_ROOT}/external/cronet/base/strings/utf_string_conversions.cc
        ${ANDROID_ROOT}/external/cronet/base/substring_set_matcher/matcher_string_pattern.cc
        ${ANDROID_ROOT}/external/cronet/base/substring_set_matcher/substring_set_matcher.cc
        ${ANDROID_ROOT}/external/cronet/base/supports_user_data.cc
        ${ANDROID_ROOT}/external/cronet/base/sync_socket.cc
        ${ANDROID_ROOT}/external/cronet/base/sync_socket_posix.cc
        ${ANDROID_ROOT}/external/cronet/base/synchronization/atomic_flag.cc
        ${ANDROID_ROOT}/external/cronet/base/synchronization/condition_variable_posix.cc
        ${ANDROID_ROOT}/external/cronet/base/synchronization/lock.cc
        ${ANDROID_ROOT}/external/cronet/base/synchronization/lock_impl_posix.cc
        ${ANDROID_ROOT}/external/cronet/base/synchronization/waitable_event.cc
        ${ANDROID_ROOT}/external/cronet/base/synchronization/waitable_event_posix.cc
        ${ANDROID_ROOT}/external/cronet/base/synchronization/waitable_event_watcher_posix.cc
        ${ANDROID_ROOT}/external/cronet/base/syslog_logging.cc
        ${ANDROID_ROOT}/external/cronet/base/system/sys_info.cc
        ${ANDROID_ROOT}/external/cronet/base/system/sys_info_android.cc
        ${ANDROID_ROOT}/external/cronet/base/system/sys_info_linux.cc
        ${ANDROID_ROOT}/external/cronet/base/system/sys_info_posix.cc
        ${ANDROID_ROOT}/external/cronet/base/system/system_monitor.cc
        ${ANDROID_ROOT}/external/cronet/base/task/cancelable_task_tracker.cc
        ${ANDROID_ROOT}/external/cronet/base/task/common/checked_lock_impl.cc
        ${ANDROID_ROOT}/external/cronet/base/task/common/lazy_now.cc
        ${ANDROID_ROOT}/external/cronet/base/task/common/operations_controller.cc
        ${ANDROID_ROOT}/external/cronet/base/task/common/scoped_defer_task_posting.cc
        ${ANDROID_ROOT}/external/cronet/base/task/common/task_annotator.cc
        ${ANDROID_ROOT}/external/cronet/base/task/current_thread.cc
        ${ANDROID_ROOT}/external/cronet/base/task/default_delayed_task_handle_delegate.cc
        ${ANDROID_ROOT}/external/cronet/base/task/deferred_sequenced_task_runner.cc
        ${ANDROID_ROOT}/external/cronet/base/task/delayed_task_handle.cc
        ${ANDROID_ROOT}/external/cronet/base/task/lazy_thread_pool_task_runner.cc
        ${ANDROID_ROOT}/external/cronet/base/task/post_job.cc
        ${ANDROID_ROOT}/external/cronet/base/task/scoped_set_task_priority_for_current_thread.cc
        ${ANDROID_ROOT}/external/cronet/base/task/sequence_manager/associated_thread_id.cc
        ${ANDROID_ROOT}/external/cronet/base/task/sequence_manager/atomic_flag_set.cc
        ${ANDROID_ROOT}/external/cronet/base/task/sequence_manager/delayed_task_handle_delegate.cc
        ${ANDROID_ROOT}/external/cronet/base/task/sequence_manager/enqueue_order_generator.cc
        ${ANDROID_ROOT}/external/cronet/base/task/sequence_manager/fence.cc
        ${ANDROID_ROOT}/external/cronet/base/task/sequence_manager/hierarchical_timing_wheel.cc
        ${ANDROID_ROOT}/external/cronet/base/task/sequence_manager/sequence_manager.cc
        ${ANDROID_ROOT}/external/cronet/base/task/sequence_manager/sequence_manager_impl.cc
        ${ANDROID_ROOT}/external/cronet/base/task/sequence_manager/sequenced_task_source.cc
        ${ANDROID_ROOT}/external/cronet/base/task/sequence_manager/task_order.cc
        ${ANDROID_ROOT}/external/cronet/base/task/sequence_manager/task_queue.cc
        ${ANDROID_ROOT}/external/cronet/base/task/sequence_manager/task_queue_impl.cc
        ${ANDROID_ROOT}/external/cronet/base/task/sequence_manager/task_queue_selector.cc
        ${ANDROID_ROOT}/external/cronet/base/task/sequence_manager/tasks.cc
        ${ANDROID_ROOT}/external/cronet/base/task/sequence_manager/thread_controller.cc
        ${ANDROID_ROOT}/external/cronet/base/task/sequence_manager/thread_controller_impl.cc
        ${ANDROID_ROOT}/external/cronet/base/task/sequence_manager/thread_controller_power_monitor.cc
        ${ANDROID_ROOT}/external/cronet/base/task/sequence_manager/thread_controller_with_message_pump_impl.cc
        ${ANDROID_ROOT}/external/cronet/base/task/sequence_manager/time_domain.cc
        ${ANDROID_ROOT}/external/cronet/base/task/sequence_manager/timing_wheel.cc
        ${ANDROID_ROOT}/external/cronet/base/task/sequence_manager/wake_up_queue.cc
        ${ANDROID_ROOT}/external/cronet/base/task/sequence_manager/work_deduplicator.cc
        ${ANDROID_ROOT}/external/cronet/base/task/sequence_manager/work_queue.cc
        ${ANDROID_ROOT}/external/cronet/base/task/sequence_manager/work_queue_sets.cc
        ${ANDROID_ROOT}/external/cronet/base/task/sequenced_task_runner.cc
        ${ANDROID_ROOT}/external/cronet/base/task/single_thread_task_executor.cc
        ${ANDROID_ROOT}/external/cronet/base/task/single_thread_task_runner.cc
        ${ANDROID_ROOT}/external/cronet/base/task/task_features.cc
        ${ANDROID_ROOT}/external/cronet/base/task/task_runner.cc
        ${ANDROID_ROOT}/external/cronet/base/task/task_traits.cc
        ${ANDROID_ROOT}/external/cronet/base/task/thread_pool.cc
        ${ANDROID_ROOT}/external/cronet/base/task/thread_pool/delayed_priority_queue.cc
        ${ANDROID_ROOT}/external/cronet/base/task/thread_pool/delayed_task_manager.cc
        ${ANDROID_ROOT}/external/cronet/base/task/thread_pool/environment_config.cc
        ${ANDROID_ROOT}/external/cronet/base/task/thread_pool/initialization_util.cc
        ${ANDROID_ROOT}/external/cronet/base/task/thread_pool/job_task_source.cc
        ${ANDROID_ROOT}/external/cronet/base/task/thread_pool/pooled_parallel_task_runner.cc
        ${ANDROID_ROOT}/external/cronet/base/task/thread_pool/pooled_sequenced_task_runner.cc
        ${ANDROID_ROOT}/external/cronet/base/task/thread_pool/pooled_single_thread_task_runner_manager.cc
        ${ANDROID_ROOT}/external/cronet/base/task/thread_pool/pooled_task_runner_delegate.cc
        ${ANDROID_ROOT}/external/cronet/base/task/thread_pool/priority_queue.cc
        ${ANDROID_ROOT}/external/cronet/base/task/thread_pool/sequence.cc
        ${ANDROID_ROOT}/external/cronet/base/task/thread_pool/service_thread.cc
        ${ANDROID_ROOT}/external/cronet/base/task/thread_pool/task.cc
        ${ANDROID_ROOT}/external/cronet/base/task/thread_pool/task_source.cc
        ${ANDROID_ROOT}/external/cronet/base/task/thread_pool/task_source_sort_key.cc
        ${ANDROID_ROOT}/external/cronet/base/task/thread_pool/task_tracker.cc
        ${ANDROID_ROOT}/external/cronet/base/task/thread_pool/thread_group.cc
        ${ANDROID_ROOT}/external/cronet/base/task/thread_pool/thread_group_impl.cc
        ${ANDROID_ROOT}/external/cronet/base/task/thread_pool/thread_pool_impl.cc
        ${ANDROID_ROOT}/external/cronet/base/task/thread_pool/thread_pool_instance.cc
        ${ANDROID_ROOT}/external/cronet/base/task/thread_pool/worker_thread.cc
        ${ANDROID_ROOT}/external/cronet/base/task/thread_pool/worker_thread_set.cc
        ${ANDROID_ROOT}/external/cronet/base/third_party/cityhash/city.cc
        ${ANDROID_ROOT}/external/cronet/base/third_party/cityhash_v103/src/city_v103.cc
        ${ANDROID_ROOT}/external/cronet/base/third_party/nspr/prtime.cc
        ${ANDROID_ROOT}/external/cronet/base/third_party/superfasthash/superfasthash.c
        ${ANDROID_ROOT}/external/cronet/base/threading/hang_watcher.cc
        ${ANDROID_ROOT}/external/cronet/base/threading/platform_thread.cc
        ${ANDROID_ROOT}/external/cronet/base/threading/platform_thread_android.cc
        ${ANDROID_ROOT}/external/cronet/base/threading/platform_thread_internal_posix.cc
        ${ANDROID_ROOT}/external/cronet/base/threading/platform_thread_posix.cc
        ${ANDROID_ROOT}/external/cronet/base/threading/platform_thread_ref.cc
        ${ANDROID_ROOT}/external/cronet/base/threading/post_task_and_reply_impl.cc
        ${ANDROID_ROOT}/external/cronet/base/threading/scoped_blocking_call.cc
        ${ANDROID_ROOT}/external/cronet/base/threading/scoped_blocking_call_internal.cc
        ${ANDROID_ROOT}/external/cronet/base/threading/scoped_thread_priority.cc
        ${ANDROID_ROOT}/external/cronet/base/threading/sequence_local_storage_map.cc
        ${ANDROID_ROOT}/external/cronet/base/threading/sequence_local_storage_slot.cc
        ${ANDROID_ROOT}/external/cronet/base/threading/simple_thread.cc
        ${ANDROID_ROOT}/external/cronet/base/threading/thread.cc
        ${ANDROID_ROOT}/external/cronet/base/threading/thread_checker.cc
        ${ANDROID_ROOT}/external/cronet/base/threading/thread_checker_impl.cc
        ${ANDROID_ROOT}/external/cronet/base/threading/thread_collision_warner.cc
        ${ANDROID_ROOT}/external/cronet/base/threading/thread_id_name_manager.cc
        ${ANDROID_ROOT}/external/cronet/base/threading/thread_local_storage.cc
        ${ANDROID_ROOT}/external/cronet/base/threading/thread_local_storage_posix.cc
        ${ANDROID_ROOT}/external/cronet/base/threading/thread_restrictions.cc
        ${ANDROID_ROOT}/external/cronet/base/threading/watchdog.cc
        ${ANDROID_ROOT}/external/cronet/base/time/clock.cc
        ${ANDROID_ROOT}/external/cronet/base/time/default_clock.cc
        ${ANDROID_ROOT}/external/cronet/base/time/default_tick_clock.cc
        ${ANDROID_ROOT}/external/cronet/base/time/tick_clock.cc
        ${ANDROID_ROOT}/external/cronet/base/time/time.cc
        ${ANDROID_ROOT}/external/cronet/base/time/time_android.cc
        ${ANDROID_ROOT}/external/cronet/base/time/time_conversion_posix.cc
        ${ANDROID_ROOT}/external/cronet/base/time/time_delta_from_string.cc
        ${ANDROID_ROOT}/external/cronet/base/time/time_exploded_icu.cc
        ${ANDROID_ROOT}/external/cronet/base/time/time_exploded_posix.cc
        ${ANDROID_ROOT}/external/cronet/base/time/time_now_posix.cc
        ${ANDROID_ROOT}/external/cronet/base/time/time_override.cc
        ${ANDROID_ROOT}/external/cronet/base/time/time_to_iso8601.cc
        ${ANDROID_ROOT}/external/cronet/base/timer/elapsed_timer.cc
        ${ANDROID_ROOT}/external/cronet/base/timer/hi_res_timer_manager_posix.cc
        ${ANDROID_ROOT}/external/cronet/base/timer/lap_timer.cc
        ${ANDROID_ROOT}/external/cronet/base/timer/timer.cc
        ${ANDROID_ROOT}/external/cronet/base/timer/wall_clock_timer.cc
        ${ANDROID_ROOT}/external/cronet/base/token.cc
        ${ANDROID_ROOT}/external/cronet/base/trace_event/heap_profiler_allocation_context.cc
        ${ANDROID_ROOT}/external/cronet/base/trace_event/heap_profiler_allocation_context_tracker.cc
        ${ANDROID_ROOT}/external/cronet/base/trace_event/memory_allocator_dump_guid.cc
        ${ANDROID_ROOT}/external/cronet/base/trace_event/memory_usage_estimator.cc
        ${ANDROID_ROOT}/external/cronet/base/trace_event/trace_event_stub.cc
        ${ANDROID_ROOT}/external/cronet/base/trace_event/trace_id_helper.cc
        ${ANDROID_ROOT}/external/cronet/base/unguessable_token.cc
        ${ANDROID_ROOT}/external/cronet/base/uuid.cc
        ${ANDROID_ROOT}/external/cronet/base/value_iterators.cc
        ${ANDROID_ROOT}/external/cronet/base/values.cc
        ${ANDROID_ROOT}/external/cronet/base/version.cc
        ${ANDROID_ROOT}/external/cronet/base/vlog.cc
        ${ANDROID_ROOT}/external/cronet/base/allocator/partition_allocator/shim/allocator_shim.cc
        ${ANDROID_ROOT}/external/cronet/base/android/reached_code_profiler.cc
        ${ANDROID_ROOT}/external/cronet/base/profiler/frame_pointer_unwinder.cc
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
    add_library(cronet_aml_base_base SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_base_base
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
            ${ANDROID_ROOT}/external/cronet/third_party/android_ndk/sources/android/cpufeatures
            ${ANDROID_ROOT}/external/cronet/third_party/boringssl/src/include
            ${ANDROID_ROOT}/external/cronet/third_party/icu/source/common
            ${ANDROID_ROOT}/external/cronet/third_party/icu/source/i18n
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_chromecast_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_chromeos_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_debugging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_logging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_partition_alloc_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_chromeos_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_anchor_functions_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_android_runtime_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_android_runtime_unchecked_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_base_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_build_date/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_cfi_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_clang_profiling_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_debugging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_feature_list_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_ios_cronet_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_logging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_message_pump_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_orderfile_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_power_monitor_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_profiler_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_rust_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_sanitizer_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_synchronization_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_tracing_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_blink_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_branding_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_chromecast_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_config_compiler_compiler_buildflags/android_${OUT_ARCH_CPU}/gen
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_base_base PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_base_base
            cronet_aml_base_allocator_partition_allocator_partition_alloc
            cronet_aml_base_base_static
            cronet_aml_base_nodebug_assertion
            cronet_aml_base_third_party_double_conversion_double_conversion
            cronet_aml_base_third_party_dynamic_annotations_dynamic_annotations
            cronet_aml_third_party_abseil_cpp_absl_base_base
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_malloc_internal
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait
            cronet_aml_third_party_abseil_cpp_absl_base_strerror
            cronet_aml_third_party_abseil_cpp_absl_base_throw_delegate
            cronet_aml_third_party_abseil_cpp_absl_container_hashtablez_sampler
            cronet_aml_third_party_abseil_cpp_absl_container_raw_hash_set
            cronet_aml_third_party_abseil_cpp_absl_crc_cpu_detect
            cronet_aml_third_party_abseil_cpp_absl_crc_crc32c
            cronet_aml_third_party_abseil_cpp_absl_crc_crc_cord_state
            cronet_aml_third_party_abseil_cpp_absl_crc_crc_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_debugging_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_demangle_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_examine_stack
            cronet_aml_third_party_abseil_cpp_absl_debugging_failure_signal_handler
            cronet_aml_third_party_abseil_cpp_absl_debugging_stacktrace
            cronet_aml_third_party_abseil_cpp_absl_debugging_symbolize
            cronet_aml_third_party_abseil_cpp_absl_hash_city
            cronet_aml_third_party_abseil_cpp_absl_hash_hash
            cronet_aml_third_party_abseil_cpp_absl_hash_low_level_hash
            cronet_aml_third_party_abseil_cpp_absl_log_die_if_null
            cronet_aml_third_party_abseil_cpp_absl_log_globals
            cronet_aml_third_party_abseil_cpp_absl_log_internal_check_op
            cronet_aml_third_party_abseil_cpp_absl_log_internal_conditions
            cronet_aml_third_party_abseil_cpp_absl_log_internal_format
            cronet_aml_third_party_abseil_cpp_absl_log_internal_globals
            cronet_aml_third_party_abseil_cpp_absl_log_internal_log_message
            cronet_aml_third_party_abseil_cpp_absl_log_internal_log_sink_set
            cronet_aml_third_party_abseil_cpp_absl_log_internal_nullguard
            cronet_aml_third_party_abseil_cpp_absl_log_internal_proto
            cronet_aml_third_party_abseil_cpp_absl_log_log_entry
            cronet_aml_third_party_abseil_cpp_absl_log_log_sink
            cronet_aml_third_party_abseil_cpp_absl_numeric_int128
            cronet_aml_third_party_abseil_cpp_absl_profiling_exponential_biased
            cronet_aml_third_party_abseil_cpp_absl_random_distributions
            cronet_aml_third_party_abseil_cpp_absl_random_internal_platform
            cronet_aml_third_party_abseil_cpp_absl_random_internal_pool_urbg
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_hwaes
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_hwaes_impl
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_slow
            cronet_aml_third_party_abseil_cpp_absl_random_internal_seed_material
            cronet_aml_third_party_abseil_cpp_absl_random_seed_gen_exception
            cronet_aml_third_party_abseil_cpp_absl_random_seed_sequences
            cronet_aml_third_party_abseil_cpp_absl_status_status
            cronet_aml_third_party_abseil_cpp_absl_status_statusor
            cronet_aml_third_party_abseil_cpp_absl_strings_cord
            cronet_aml_third_party_abseil_cpp_absl_strings_cord_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_functions
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_handle
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_info
            cronet_aml_third_party_abseil_cpp_absl_strings_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_str_format_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_strings
            cronet_aml_third_party_abseil_cpp_absl_synchronization_graphcycles_internal
            cronet_aml_third_party_abseil_cpp_absl_synchronization_kernel_timeout_internal
            cronet_aml_third_party_abseil_cpp_absl_synchronization_synchronization
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_civil_time
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_time_zone
            cronet_aml_third_party_abseil_cpp_absl_time_time
            cronet_aml_third_party_abseil_cpp_absl_types_bad_optional_access
            cronet_aml_third_party_abseil_cpp_absl_types_bad_variant_access
            cronet_aml_third_party_android_ndk_cpu_features
            cronet_aml_third_party_ashmem_ashmem
            cronet_aml_third_party_icu_icui18n
            cronet_aml_third_party_icu_icuuc_private
            cronet_aml_third_party_libevent_libevent
            cronet_aml_third_party_modp_b64_modp_b64
            jni_headers
            libandroid
            libc
            libclang_rt.builtins
            libcrypto
            libdl
            libgtest_prod_headers
            liblog
            libm
            libmediandk
            libssl
    )
endif ()
#################### cronet_aml_base_base ####################

#################### cronet_aml_components_cronet_cronet_common ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/components/cronet/cronet_context.cc
        ${ANDROID_ROOT}/external/cronet/components/cronet/cronet_prefs_manager.cc
        ${ANDROID_ROOT}/external/cronet/components/cronet/cronet_upload_data_stream.cc
        ${ANDROID_ROOT}/external/cronet/components/cronet/cronet_url_request.cc
        ${ANDROID_ROOT}/external/cronet/components/cronet/host_cache_persistence_manager.cc
        ${ANDROID_ROOT}/external/cronet/components/cronet/stale_host_resolver.cc
        ${ANDROID_ROOT}/external/cronet/components/cronet/url_request_context_config.cc
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
    add_library(cronet_aml_components_cronet_cronet_common SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_components_cronet_cronet_common
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/overrides
            ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src
            ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/common/platform/default
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
            ${ANDROID_ROOT}/external/cronet/third_party/boringssl/src/include
            ${ANDROID_ROOT}/external/cronet/third_party/protobuf/src
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_chromecast_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_chromeos_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_debugging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_logging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_partition_alloc_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_anchor_functions_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_android_runtime_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_android_runtime_unchecked_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_base_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_build_date/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_cfi_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_clang_profiling_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_debugging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_feature_list_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_ios_cronet_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_logging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_message_pump_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_orderfile_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_power_monitor_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_profiler_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_rust_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_sanitizer_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_synchronization_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_tracing_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_blink_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_branding_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_chromecast_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_chromeos_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_config_compiler_compiler_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_components_cronet_cronet_version_header_action/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_third_party_metrics_proto_metrics_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_third_party_metrics_proto_metrics_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/protos
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_third_party_metrics_proto_metrics_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/third_party/metrics_proto
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_components_prefs_android_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_crypto_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_third_party_quiche_net_quic_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_third_party_quiche_net_quic_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/net/third_party/quiche/src
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_third_party_quiche_net_quic_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/protos
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_cronet_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_base_registry_controlled_domains_registry_controlled_domains/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_isolation_info_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_isolation_info_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/net/base
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_isolation_info_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/protos
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_net_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_url_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_net_nqe_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_net_nqe_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/net/nqe/proto
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_net_nqe_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/protos
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_third_party_quiche_net_quic_test_tools_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_third_party_quiche_net_quic_test_tools_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/net/third_party/quiche/src
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_third_party_quiche_net_quic_test_tools_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/net/third_party/quiche/src/quiche/quic/test_tools
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_third_party_quiche_net_quic_test_tools_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/protos
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_url_url_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_components_cronet_cronet_buildflags/android_${OUT_ARCH_CPU}/gen
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_components_cronet_cronet_common PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_components_cronet_cronet_common
            cronet_aml_base_allocator_partition_allocator_partition_alloc
            cronet_aml_base_base
            cronet_aml_base_base_static
            cronet_aml_base_nodebug_assertion
            cronet_aml_base_third_party_double_conversion_double_conversion
            cronet_aml_base_third_party_dynamic_annotations_dynamic_annotations
            cronet_aml_components_cronet_cronet_version_header
            cronet_aml_components_metrics_library_support
            cronet_aml_components_prefs_prefs
            cronet_aml_components_prefs_prefs_features
            cronet_aml_crypto_crypto
            cronet_aml_net_dns_dns
            cronet_aml_net_dns_public_public
            cronet_aml_net_http_transport_security_state_generated_files
            cronet_aml_net_net
            cronet_aml_net_net_deps
            cronet_aml_net_net_public_deps
            cronet_aml_net_preload_decoder
            cronet_aml_net_third_party_quiche_quiche
            cronet_aml_net_traffic_annotation_traffic_annotation
            cronet_aml_net_uri_template
            cronet_aml_third_party_abseil_cpp_absl_base_base
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_malloc_internal
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait
            cronet_aml_third_party_abseil_cpp_absl_base_strerror
            cronet_aml_third_party_abseil_cpp_absl_base_throw_delegate
            cronet_aml_third_party_abseil_cpp_absl_container_hashtablez_sampler
            cronet_aml_third_party_abseil_cpp_absl_container_raw_hash_set
            cronet_aml_third_party_abseil_cpp_absl_crc_cpu_detect
            cronet_aml_third_party_abseil_cpp_absl_crc_crc32c
            cronet_aml_third_party_abseil_cpp_absl_crc_crc_cord_state
            cronet_aml_third_party_abseil_cpp_absl_crc_crc_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_debugging_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_demangle_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_examine_stack
            cronet_aml_third_party_abseil_cpp_absl_debugging_failure_signal_handler
            cronet_aml_third_party_abseil_cpp_absl_debugging_stacktrace
            cronet_aml_third_party_abseil_cpp_absl_debugging_symbolize
            cronet_aml_third_party_abseil_cpp_absl_hash_city
            cronet_aml_third_party_abseil_cpp_absl_hash_hash
            cronet_aml_third_party_abseil_cpp_absl_hash_low_level_hash
            cronet_aml_third_party_abseil_cpp_absl_log_die_if_null
            cronet_aml_third_party_abseil_cpp_absl_log_globals
            cronet_aml_third_party_abseil_cpp_absl_log_internal_check_op
            cronet_aml_third_party_abseil_cpp_absl_log_internal_conditions
            cronet_aml_third_party_abseil_cpp_absl_log_internal_format
            cronet_aml_third_party_abseil_cpp_absl_log_internal_globals
            cronet_aml_third_party_abseil_cpp_absl_log_internal_log_message
            cronet_aml_third_party_abseil_cpp_absl_log_internal_log_sink_set
            cronet_aml_third_party_abseil_cpp_absl_log_internal_nullguard
            cronet_aml_third_party_abseil_cpp_absl_log_internal_proto
            cronet_aml_third_party_abseil_cpp_absl_log_log_entry
            cronet_aml_third_party_abseil_cpp_absl_log_log_sink
            cronet_aml_third_party_abseil_cpp_absl_numeric_int128
            cronet_aml_third_party_abseil_cpp_absl_profiling_exponential_biased
            cronet_aml_third_party_abseil_cpp_absl_random_distributions
            cronet_aml_third_party_abseil_cpp_absl_random_internal_platform
            cronet_aml_third_party_abseil_cpp_absl_random_internal_pool_urbg
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_hwaes
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_hwaes_impl
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_slow
            cronet_aml_third_party_abseil_cpp_absl_random_internal_seed_material
            cronet_aml_third_party_abseil_cpp_absl_random_seed_gen_exception
            cronet_aml_third_party_abseil_cpp_absl_random_seed_sequences
            cronet_aml_third_party_abseil_cpp_absl_status_status
            cronet_aml_third_party_abseil_cpp_absl_status_statusor
            cronet_aml_third_party_abseil_cpp_absl_strings_cord
            cronet_aml_third_party_abseil_cpp_absl_strings_cord_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_functions
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_handle
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_info
            cronet_aml_third_party_abseil_cpp_absl_strings_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_str_format_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_strings
            cronet_aml_third_party_abseil_cpp_absl_synchronization_graphcycles_internal
            cronet_aml_third_party_abseil_cpp_absl_synchronization_kernel_timeout_internal
            cronet_aml_third_party_abseil_cpp_absl_synchronization_synchronization
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_civil_time
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_time_zone
            cronet_aml_third_party_abseil_cpp_absl_time_time
            cronet_aml_third_party_abseil_cpp_absl_types_bad_optional_access
            cronet_aml_third_party_abseil_cpp_absl_types_bad_variant_access
            cronet_aml_third_party_android_ndk_cpu_features
            cronet_aml_third_party_ashmem_ashmem
            cronet_aml_third_party_brotli_common
            cronet_aml_third_party_brotli_dec
            cronet_aml_third_party_icu_icui18n
            cronet_aml_third_party_icu_icuuc_private
            cronet_aml_third_party_libevent_libevent
            cronet_aml_third_party_modp_b64_modp_b64
            cronet_aml_third_party_protobuf_protobuf_lite
            cronet_aml_url_url
            jni_headers
            libandroid
            libc
            libclang_rt.builtins
            libcrypto
            libdl
            liblog
            libm
            libmediandk
            libssl
            libz
    )
endif ()
#################### cronet_aml_components_cronet_cronet_common ####################

#################### cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_time_zone ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/time/internal/cctz/src/time_zone_fixed.cc
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/time/internal/cctz/src/time_zone_format.cc
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/time/internal/cctz/src/time_zone_if.cc
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/time/internal/cctz/src/time_zone_impl.cc
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/time/internal/cctz/src/time_zone_info.cc
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/time/internal/cctz/src/time_zone_libc.cc
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/time/internal/cctz/src/time_zone_lookup.cc
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/time/internal/cctz/src/time_zone_posix.cc
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/time/internal/cctz/src/zone_info_source.cc
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
    add_library(cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_time_zone SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_time_zone
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_time_zone PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_time_zone
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_civil_time
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_time_zone ####################

#################### cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/base/internal/raw_logging.cc
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
    add_library(cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal ####################

#################### cronet_aml_url_url ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/url/gurl.cc
        ${ANDROID_ROOT}/external/cronet/url/origin.cc
        ${ANDROID_ROOT}/external/cronet/url/scheme_host_port.cc
        ${ANDROID_ROOT}/external/cronet/url/third_party/mozilla/url_parse.cc
        ${ANDROID_ROOT}/external/cronet/url/url_canon.cc
        ${ANDROID_ROOT}/external/cronet/url/url_canon_etc.cc
        ${ANDROID_ROOT}/external/cronet/url/url_canon_filesystemurl.cc
        ${ANDROID_ROOT}/external/cronet/url/url_canon_fileurl.cc
        ${ANDROID_ROOT}/external/cronet/url/url_canon_host.cc
        ${ANDROID_ROOT}/external/cronet/url/url_canon_internal.cc
        ${ANDROID_ROOT}/external/cronet/url/url_canon_ip.cc
        ${ANDROID_ROOT}/external/cronet/url/url_canon_mailtourl.cc
        ${ANDROID_ROOT}/external/cronet/url/url_canon_path.cc
        ${ANDROID_ROOT}/external/cronet/url/url_canon_pathurl.cc
        ${ANDROID_ROOT}/external/cronet/url/url_canon_query.cc
        ${ANDROID_ROOT}/external/cronet/url/url_canon_relative.cc
        ${ANDROID_ROOT}/external/cronet/url/url_canon_stdstring.cc
        ${ANDROID_ROOT}/external/cronet/url/url_canon_stdurl.cc
        ${ANDROID_ROOT}/external/cronet/url/url_constants.cc
        ${ANDROID_ROOT}/external/cronet/url/url_features.cc
        ${ANDROID_ROOT}/external/cronet/url/url_idna_icu_alternatives_android.cc
        ${ANDROID_ROOT}/external/cronet/url/url_parse_file.cc
        ${ANDROID_ROOT}/external/cronet/url/url_util.cc
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
    add_library(cronet_aml_url_url SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_url_url
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
            ${ANDROID_ROOT}/external/cronet/third_party/boringssl/src/include
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_chromecast_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_chromeos_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_debugging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_logging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_partition_alloc_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_anchor_functions_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_android_runtime_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_android_runtime_unchecked_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_base_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_build_date/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_cfi_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_clang_profiling_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_debugging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_feature_list_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_ios_cronet_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_logging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_message_pump_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_orderfile_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_power_monitor_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_profiler_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_rust_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_sanitizer_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_synchronization_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_tracing_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_blink_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_branding_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_chromecast_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_chromeos_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_config_compiler_compiler_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_url_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_url_url_jni_headers/android_${OUT_ARCH_CPU}/gen
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_url_url PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_url_url
            cronet_aml_base_allocator_partition_allocator_partition_alloc
            cronet_aml_base_base
            cronet_aml_base_base_static
            cronet_aml_base_nodebug_assertion
            cronet_aml_base_third_party_double_conversion_double_conversion
            cronet_aml_base_third_party_dynamic_annotations_dynamic_annotations
            cronet_aml_third_party_abseil_cpp_absl_base_base
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_malloc_internal
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait
            cronet_aml_third_party_abseil_cpp_absl_base_strerror
            cronet_aml_third_party_abseil_cpp_absl_base_throw_delegate
            cronet_aml_third_party_abseil_cpp_absl_container_hashtablez_sampler
            cronet_aml_third_party_abseil_cpp_absl_container_raw_hash_set
            cronet_aml_third_party_abseil_cpp_absl_crc_cpu_detect
            cronet_aml_third_party_abseil_cpp_absl_crc_crc32c
            cronet_aml_third_party_abseil_cpp_absl_crc_crc_cord_state
            cronet_aml_third_party_abseil_cpp_absl_crc_crc_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_debugging_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_demangle_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_examine_stack
            cronet_aml_third_party_abseil_cpp_absl_debugging_failure_signal_handler
            cronet_aml_third_party_abseil_cpp_absl_debugging_stacktrace
            cronet_aml_third_party_abseil_cpp_absl_debugging_symbolize
            cronet_aml_third_party_abseil_cpp_absl_hash_city
            cronet_aml_third_party_abseil_cpp_absl_hash_hash
            cronet_aml_third_party_abseil_cpp_absl_hash_low_level_hash
            cronet_aml_third_party_abseil_cpp_absl_log_die_if_null
            cronet_aml_third_party_abseil_cpp_absl_log_globals
            cronet_aml_third_party_abseil_cpp_absl_log_internal_check_op
            cronet_aml_third_party_abseil_cpp_absl_log_internal_conditions
            cronet_aml_third_party_abseil_cpp_absl_log_internal_format
            cronet_aml_third_party_abseil_cpp_absl_log_internal_globals
            cronet_aml_third_party_abseil_cpp_absl_log_internal_log_message
            cronet_aml_third_party_abseil_cpp_absl_log_internal_log_sink_set
            cronet_aml_third_party_abseil_cpp_absl_log_internal_nullguard
            cronet_aml_third_party_abseil_cpp_absl_log_internal_proto
            cronet_aml_third_party_abseil_cpp_absl_log_log_entry
            cronet_aml_third_party_abseil_cpp_absl_log_log_sink
            cronet_aml_third_party_abseil_cpp_absl_numeric_int128
            cronet_aml_third_party_abseil_cpp_absl_profiling_exponential_biased
            cronet_aml_third_party_abseil_cpp_absl_random_distributions
            cronet_aml_third_party_abseil_cpp_absl_random_internal_platform
            cronet_aml_third_party_abseil_cpp_absl_random_internal_pool_urbg
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_hwaes
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_hwaes_impl
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_slow
            cronet_aml_third_party_abseil_cpp_absl_random_internal_seed_material
            cronet_aml_third_party_abseil_cpp_absl_random_seed_gen_exception
            cronet_aml_third_party_abseil_cpp_absl_random_seed_sequences
            cronet_aml_third_party_abseil_cpp_absl_status_status
            cronet_aml_third_party_abseil_cpp_absl_status_statusor
            cronet_aml_third_party_abseil_cpp_absl_strings_cord
            cronet_aml_third_party_abseil_cpp_absl_strings_cord_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_functions
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_handle
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_info
            cronet_aml_third_party_abseil_cpp_absl_strings_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_str_format_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_strings
            cronet_aml_third_party_abseil_cpp_absl_synchronization_graphcycles_internal
            cronet_aml_third_party_abseil_cpp_absl_synchronization_kernel_timeout_internal
            cronet_aml_third_party_abseil_cpp_absl_synchronization_synchronization
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_civil_time
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_time_zone
            cronet_aml_third_party_abseil_cpp_absl_time_time
            cronet_aml_third_party_abseil_cpp_absl_types_bad_optional_access
            cronet_aml_third_party_abseil_cpp_absl_types_bad_variant_access
            cronet_aml_third_party_android_ndk_cpu_features
            cronet_aml_third_party_ashmem_ashmem
            cronet_aml_third_party_icu_icui18n
            cronet_aml_third_party_icu_icuuc_private
            cronet_aml_third_party_libevent_libevent
            cronet_aml_third_party_modp_b64_modp_b64
            jni_headers
            libandroid
            libc
            libclang_rt.builtins
            libcrypto
            libdl
            liblog
            libm
            libmediandk
            libssl
    )
endif ()
#################### cronet_aml_url_url ####################

#################### cronet_aml_build_android_build_config_gen_preprocess ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_android_build_config_gen_rename/gen/BuildConfig.cc
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
    add_library(cronet_aml_build_android_build_config_gen_preprocess SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_build_android_build_config_gen_preprocess
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_build_android_build_config_gen_preprocess PRIVATE
            -Werror
    )

    target_link_libraries(cronet_aml_build_android_build_config_gen_preprocess
            libc
            libc++
            libc++demangle
            libdl
            libm
    )
endif ()
#################### cronet_aml_build_android_build_config_gen_preprocess ####################

#################### cronet_aml_third_party_abseil_cpp_absl_crc_crc_cord_state ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/crc/internal/crc_cord_state.cc
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
    add_library(cronet_aml_third_party_abseil_cpp_absl_crc_crc_cord_state SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_abseil_cpp_absl_crc_crc_cord_state
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_abseil_cpp_absl_crc_crc_cord_state PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_abseil_cpp_absl_crc_crc_cord_state
            cronet_aml_third_party_abseil_cpp_absl_base_base
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait
            cronet_aml_third_party_abseil_cpp_absl_base_throw_delegate
            cronet_aml_third_party_abseil_cpp_absl_crc_cpu_detect
            cronet_aml_third_party_abseil_cpp_absl_crc_crc32c
            cronet_aml_third_party_abseil_cpp_absl_crc_crc_internal
            cronet_aml_third_party_abseil_cpp_absl_numeric_int128
            cronet_aml_third_party_abseil_cpp_absl_strings_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_strings
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_abseil_cpp_absl_crc_crc_cord_state ####################

#################### cronet_aml_third_party_protobuf_protobuf_lite ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/protobuf/src/google/protobuf/any_lite.cc
        ${ANDROID_ROOT}/external/cronet/third_party/protobuf/src/google/protobuf/arena.cc
        ${ANDROID_ROOT}/external/cronet/third_party/protobuf/src/google/protobuf/arenastring.cc
        ${ANDROID_ROOT}/external/cronet/third_party/protobuf/src/google/protobuf/arenaz_sampler.cc
        ${ANDROID_ROOT}/external/cronet/third_party/protobuf/src/google/protobuf/extension_set.cc
        ${ANDROID_ROOT}/external/cronet/third_party/protobuf/src/google/protobuf/generated_enum_util.cc
        ${ANDROID_ROOT}/external/cronet/third_party/protobuf/src/google/protobuf/generated_message_tctable_lite.cc
        ${ANDROID_ROOT}/external/cronet/third_party/protobuf/src/google/protobuf/generated_message_util.cc
        ${ANDROID_ROOT}/external/cronet/third_party/protobuf/src/google/protobuf/implicit_weak_message.cc
        ${ANDROID_ROOT}/external/cronet/third_party/protobuf/src/google/protobuf/inlined_string_field.cc
        ${ANDROID_ROOT}/external/cronet/third_party/protobuf/src/google/protobuf/io/coded_stream.cc
        ${ANDROID_ROOT}/external/cronet/third_party/protobuf/src/google/protobuf/io/io_win32.cc
        ${ANDROID_ROOT}/external/cronet/third_party/protobuf/src/google/protobuf/io/strtod.cc
        ${ANDROID_ROOT}/external/cronet/third_party/protobuf/src/google/protobuf/io/zero_copy_stream.cc
        ${ANDROID_ROOT}/external/cronet/third_party/protobuf/src/google/protobuf/io/zero_copy_stream_impl.cc
        ${ANDROID_ROOT}/external/cronet/third_party/protobuf/src/google/protobuf/io/zero_copy_stream_impl_lite.cc
        ${ANDROID_ROOT}/external/cronet/third_party/protobuf/src/google/protobuf/map.cc
        ${ANDROID_ROOT}/external/cronet/third_party/protobuf/src/google/protobuf/message_lite.cc
        ${ANDROID_ROOT}/external/cronet/third_party/protobuf/src/google/protobuf/parse_context.cc
        ${ANDROID_ROOT}/external/cronet/third_party/protobuf/src/google/protobuf/repeated_field.cc
        ${ANDROID_ROOT}/external/cronet/third_party/protobuf/src/google/protobuf/repeated_ptr_field.cc
        ${ANDROID_ROOT}/external/cronet/third_party/protobuf/src/google/protobuf/stubs/bytestream.cc
        ${ANDROID_ROOT}/external/cronet/third_party/protobuf/src/google/protobuf/stubs/common.cc
        ${ANDROID_ROOT}/external/cronet/third_party/protobuf/src/google/protobuf/stubs/int128.cc
        ${ANDROID_ROOT}/external/cronet/third_party/protobuf/src/google/protobuf/stubs/status.cc
        ${ANDROID_ROOT}/external/cronet/third_party/protobuf/src/google/protobuf/stubs/statusor.cc
        ${ANDROID_ROOT}/external/cronet/third_party/protobuf/src/google/protobuf/stubs/stringpiece.cc
        ${ANDROID_ROOT}/external/cronet/third_party/protobuf/src/google/protobuf/stubs/stringprintf.cc
        ${ANDROID_ROOT}/external/cronet/third_party/protobuf/src/google/protobuf/stubs/structurally_valid.cc
        ${ANDROID_ROOT}/external/cronet/third_party/protobuf/src/google/protobuf/stubs/strutil.cc
        ${ANDROID_ROOT}/external/cronet/third_party/protobuf/src/google/protobuf/stubs/time.cc
        ${ANDROID_ROOT}/external/cronet/third_party/protobuf/src/google/protobuf/wire_format_lite.cc
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
    add_library(cronet_aml_third_party_protobuf_protobuf_lite SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_protobuf_protobuf_lite
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/protobuf/src
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_protobuf_protobuf_lite PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_protobuf_protobuf_lite
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            liblog
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_protobuf_protobuf_lite ####################

#################### cronet_aml_third_party_abseil_cpp_absl_strings_cordz_functions ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/strings/internal/cordz_functions.cc
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
    add_library(cronet_aml_third_party_abseil_cpp_absl_strings_cordz_functions SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_abseil_cpp_absl_strings_cordz_functions
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_abseil_cpp_absl_strings_cordz_functions PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_abseil_cpp_absl_strings_cordz_functions
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            cronet_aml_third_party_abseil_cpp_absl_profiling_exponential_biased
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_abseil_cpp_absl_strings_cordz_functions ####################

#################### cronet_aml_third_party_abseil_cpp_absl_random_seed_sequences ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/random/seed_sequences.cc
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
    add_library(cronet_aml_third_party_abseil_cpp_absl_random_seed_sequences SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_abseil_cpp_absl_random_seed_sequences
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_chromeos_buildflags/android_${OUT_ARCH_CPU}/gen
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_abseil_cpp_absl_random_seed_sequences PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_abseil_cpp_absl_random_seed_sequences
            cronet_aml_third_party_abseil_cpp_absl_base_base
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait
            cronet_aml_third_party_abseil_cpp_absl_base_throw_delegate
            cronet_aml_third_party_abseil_cpp_absl_numeric_int128
            cronet_aml_third_party_abseil_cpp_absl_random_internal_platform
            cronet_aml_third_party_abseil_cpp_absl_random_internal_pool_urbg
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_hwaes
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_hwaes_impl
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_slow
            cronet_aml_third_party_abseil_cpp_absl_random_internal_seed_material
            cronet_aml_third_party_abseil_cpp_absl_random_seed_gen_exception
            cronet_aml_third_party_abseil_cpp_absl_strings_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_strings
            cronet_aml_third_party_abseil_cpp_absl_types_bad_optional_access
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_abseil_cpp_absl_random_seed_sequences ####################

#################### cronet_aml_net_android_net_errors_java_preprocess ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_android_net_errors_java_rename/gen/NetError.cc
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
    add_library(cronet_aml_net_android_net_errors_java_preprocess SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_net_android_net_errors_java_preprocess
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_net_android_net_errors_java_preprocess PRIVATE
            -Werror
    )

    target_link_libraries(cronet_aml_net_android_net_errors_java_preprocess
            libc
            libc++
            libc++demangle
            libdl
            libm
    )
endif ()
#################### cronet_aml_net_android_net_errors_java_preprocess ####################

#################### cronet_aml_third_party_libevent_libevent ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/libevent/buffer.c
        ${ANDROID_ROOT}/external/cronet/third_party/libevent/epoll.c
        ${ANDROID_ROOT}/external/cronet/third_party/libevent/evbuffer.c
        ${ANDROID_ROOT}/external/cronet/third_party/libevent/evdns.c
        ${ANDROID_ROOT}/external/cronet/third_party/libevent/event.c
        ${ANDROID_ROOT}/external/cronet/third_party/libevent/event_tagging.c
        ${ANDROID_ROOT}/external/cronet/third_party/libevent/evrpc.c
        ${ANDROID_ROOT}/external/cronet/third_party/libevent/evutil.c
        ${ANDROID_ROOT}/external/cronet/third_party/libevent/http.c
        ${ANDROID_ROOT}/external/cronet/third_party/libevent/log.c
        ${ANDROID_ROOT}/external/cronet/third_party/libevent/poll.c
        ${ANDROID_ROOT}/external/cronet/third_party/libevent/select.c
        ${ANDROID_ROOT}/external/cronet/third_party/libevent/signal.c
        ${ANDROID_ROOT}/external/cronet/third_party/libevent/strlcpy.c
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
    add_library(cronet_aml_third_party_libevent_libevent SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_libevent_libevent
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/libevent/android
            ${ANDROID_ROOT}/external/cronet/third_party/libevent/compat
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_libevent_libevent PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_libevent_libevent
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_libevent_libevent ####################

#################### cronet_aml_base_base_static ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/base/base_switches.cc
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
    add_library(cronet_aml_base_base_static SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_base_base_static
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_chromeos_buildflags/android_${OUT_ARCH_CPU}/gen
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_base_base_static PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_base_base_static
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_base_base_static ####################

#################### cronet_aml_third_party_abseil_cpp_absl_log_die_if_null ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/log/die_if_null.cc
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
    add_library(cronet_aml_third_party_abseil_cpp_absl_log_die_if_null SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_abseil_cpp_absl_log_die_if_null
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_abseil_cpp_absl_log_die_if_null PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_abseil_cpp_absl_log_die_if_null
            cronet_aml_third_party_abseil_cpp_absl_base_base
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_malloc_internal
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait
            cronet_aml_third_party_abseil_cpp_absl_base_strerror
            cronet_aml_third_party_abseil_cpp_absl_base_throw_delegate
            cronet_aml_third_party_abseil_cpp_absl_debugging_debugging_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_demangle_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_examine_stack
            cronet_aml_third_party_abseil_cpp_absl_debugging_stacktrace
            cronet_aml_third_party_abseil_cpp_absl_debugging_symbolize
            cronet_aml_third_party_abseil_cpp_absl_hash_city
            cronet_aml_third_party_abseil_cpp_absl_hash_hash
            cronet_aml_third_party_abseil_cpp_absl_hash_low_level_hash
            cronet_aml_third_party_abseil_cpp_absl_log_globals
            cronet_aml_third_party_abseil_cpp_absl_log_internal_conditions
            cronet_aml_third_party_abseil_cpp_absl_log_internal_format
            cronet_aml_third_party_abseil_cpp_absl_log_internal_globals
            cronet_aml_third_party_abseil_cpp_absl_log_internal_log_message
            cronet_aml_third_party_abseil_cpp_absl_log_internal_log_sink_set
            cronet_aml_third_party_abseil_cpp_absl_log_internal_nullguard
            cronet_aml_third_party_abseil_cpp_absl_log_internal_proto
            cronet_aml_third_party_abseil_cpp_absl_log_log_entry
            cronet_aml_third_party_abseil_cpp_absl_log_log_sink
            cronet_aml_third_party_abseil_cpp_absl_numeric_int128
            cronet_aml_third_party_abseil_cpp_absl_strings_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_str_format_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_strings
            cronet_aml_third_party_abseil_cpp_absl_synchronization_graphcycles_internal
            cronet_aml_third_party_abseil_cpp_absl_synchronization_kernel_timeout_internal
            cronet_aml_third_party_abseil_cpp_absl_synchronization_synchronization
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_civil_time
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_time_zone
            cronet_aml_third_party_abseil_cpp_absl_time_time
            cronet_aml_third_party_abseil_cpp_absl_types_bad_optional_access
            cronet_aml_third_party_abseil_cpp_absl_types_bad_variant_access
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            liblog
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_abseil_cpp_absl_log_die_if_null ####################

#################### cronet_aml_third_party_abseil_cpp_absl_types_bad_optional_access ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/types/bad_optional_access.cc
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
    add_library(cronet_aml_third_party_abseil_cpp_absl_types_bad_optional_access SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_abseil_cpp_absl_types_bad_optional_access
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_abseil_cpp_absl_types_bad_optional_access PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_abseil_cpp_absl_types_bad_optional_access
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_abseil_cpp_absl_types_bad_optional_access ####################

#################### cronet_aml_third_party_abseil_cpp_absl_debugging_failure_signal_handler ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/debugging/failure_signal_handler.cc
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
    add_library(cronet_aml_third_party_abseil_cpp_absl_debugging_failure_signal_handler SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_abseil_cpp_absl_debugging_failure_signal_handler
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_abseil_cpp_absl_debugging_failure_signal_handler PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_abseil_cpp_absl_debugging_failure_signal_handler
            cronet_aml_third_party_abseil_cpp_absl_base_base
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_malloc_internal
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait
            cronet_aml_third_party_abseil_cpp_absl_base_throw_delegate
            cronet_aml_third_party_abseil_cpp_absl_debugging_debugging_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_demangle_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_examine_stack
            cronet_aml_third_party_abseil_cpp_absl_debugging_stacktrace
            cronet_aml_third_party_abseil_cpp_absl_debugging_symbolize
            cronet_aml_third_party_abseil_cpp_absl_numeric_int128
            cronet_aml_third_party_abseil_cpp_absl_strings_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_strings
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_abseil_cpp_absl_debugging_failure_signal_handler ####################

#################### cronet_aml_third_party_abseil_cpp_absl_strings_cord ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/strings/cord.cc
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/strings/cord_analysis.cc
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/strings/cord_buffer.cc
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
    add_library(cronet_aml_third_party_abseil_cpp_absl_strings_cord SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_abseil_cpp_absl_strings_cord
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_abseil_cpp_absl_strings_cord PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_abseil_cpp_absl_strings_cord
            cronet_aml_third_party_abseil_cpp_absl_base_base
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_malloc_internal
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait
            cronet_aml_third_party_abseil_cpp_absl_base_throw_delegate
            cronet_aml_third_party_abseil_cpp_absl_crc_cpu_detect
            cronet_aml_third_party_abseil_cpp_absl_crc_crc32c
            cronet_aml_third_party_abseil_cpp_absl_crc_crc_cord_state
            cronet_aml_third_party_abseil_cpp_absl_crc_crc_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_debugging_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_demangle_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_stacktrace
            cronet_aml_third_party_abseil_cpp_absl_debugging_symbolize
            cronet_aml_third_party_abseil_cpp_absl_numeric_int128
            cronet_aml_third_party_abseil_cpp_absl_profiling_exponential_biased
            cronet_aml_third_party_abseil_cpp_absl_strings_cord_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_functions
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_handle
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_info
            cronet_aml_third_party_abseil_cpp_absl_strings_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_strings
            cronet_aml_third_party_abseil_cpp_absl_synchronization_graphcycles_internal
            cronet_aml_third_party_abseil_cpp_absl_synchronization_kernel_timeout_internal
            cronet_aml_third_party_abseil_cpp_absl_synchronization_synchronization
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_civil_time
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_time_zone
            cronet_aml_third_party_abseil_cpp_absl_time_time
            cronet_aml_third_party_abseil_cpp_absl_types_bad_optional_access
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_abseil_cpp_absl_strings_cord ####################

#################### cronet_aml_third_party_abseil_cpp_absl_log_internal_format ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/log/internal/log_format.cc
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
    add_library(cronet_aml_third_party_abseil_cpp_absl_log_internal_format SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_abseil_cpp_absl_log_internal_format
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_abseil_cpp_absl_log_internal_format PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_abseil_cpp_absl_log_internal_format
            cronet_aml_third_party_abseil_cpp_absl_base_base
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait
            cronet_aml_third_party_abseil_cpp_absl_base_throw_delegate
            cronet_aml_third_party_abseil_cpp_absl_log_internal_globals
            cronet_aml_third_party_abseil_cpp_absl_numeric_int128
            cronet_aml_third_party_abseil_cpp_absl_strings_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_str_format_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_strings
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_civil_time
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_time_zone
            cronet_aml_third_party_abseil_cpp_absl_time_time
            cronet_aml_third_party_abseil_cpp_absl_types_bad_optional_access
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_abseil_cpp_absl_log_internal_format ####################

#################### cronet_aml_components_cronet_metrics_util ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/components/cronet/metrics_util.cc
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
    add_library(cronet_aml_components_cronet_metrics_util SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_components_cronet_metrics_util
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
            ${ANDROID_ROOT}/external/cronet/third_party/boringssl/src/include
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_chromecast_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_chromeos_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_debugging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_logging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_partition_alloc_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_anchor_functions_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_android_runtime_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_android_runtime_unchecked_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_base_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_build_date/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_cfi_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_clang_profiling_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_debugging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_feature_list_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_ios_cronet_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_logging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_message_pump_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_orderfile_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_power_monitor_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_profiler_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_rust_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_sanitizer_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_synchronization_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_tracing_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_blink_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_branding_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_chromecast_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_chromeos_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_config_compiler_compiler_buildflags/android_${OUT_ARCH_CPU}/gen
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_components_cronet_metrics_util PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_components_cronet_metrics_util
            cronet_aml_base_allocator_partition_allocator_partition_alloc
            cronet_aml_base_base
            cronet_aml_base_base_static
            cronet_aml_base_nodebug_assertion
            cronet_aml_base_third_party_double_conversion_double_conversion
            cronet_aml_base_third_party_dynamic_annotations_dynamic_annotations
            cronet_aml_third_party_abseil_cpp_absl_base_base
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_malloc_internal
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait
            cronet_aml_third_party_abseil_cpp_absl_base_strerror
            cronet_aml_third_party_abseil_cpp_absl_base_throw_delegate
            cronet_aml_third_party_abseil_cpp_absl_container_hashtablez_sampler
            cronet_aml_third_party_abseil_cpp_absl_container_raw_hash_set
            cronet_aml_third_party_abseil_cpp_absl_crc_cpu_detect
            cronet_aml_third_party_abseil_cpp_absl_crc_crc32c
            cronet_aml_third_party_abseil_cpp_absl_crc_crc_cord_state
            cronet_aml_third_party_abseil_cpp_absl_crc_crc_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_debugging_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_demangle_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_examine_stack
            cronet_aml_third_party_abseil_cpp_absl_debugging_failure_signal_handler
            cronet_aml_third_party_abseil_cpp_absl_debugging_stacktrace
            cronet_aml_third_party_abseil_cpp_absl_debugging_symbolize
            cronet_aml_third_party_abseil_cpp_absl_hash_city
            cronet_aml_third_party_abseil_cpp_absl_hash_hash
            cronet_aml_third_party_abseil_cpp_absl_hash_low_level_hash
            cronet_aml_third_party_abseil_cpp_absl_log_die_if_null
            cronet_aml_third_party_abseil_cpp_absl_log_globals
            cronet_aml_third_party_abseil_cpp_absl_log_internal_check_op
            cronet_aml_third_party_abseil_cpp_absl_log_internal_conditions
            cronet_aml_third_party_abseil_cpp_absl_log_internal_format
            cronet_aml_third_party_abseil_cpp_absl_log_internal_globals
            cronet_aml_third_party_abseil_cpp_absl_log_internal_log_message
            cronet_aml_third_party_abseil_cpp_absl_log_internal_log_sink_set
            cronet_aml_third_party_abseil_cpp_absl_log_internal_nullguard
            cronet_aml_third_party_abseil_cpp_absl_log_internal_proto
            cronet_aml_third_party_abseil_cpp_absl_log_log_entry
            cronet_aml_third_party_abseil_cpp_absl_log_log_sink
            cronet_aml_third_party_abseil_cpp_absl_numeric_int128
            cronet_aml_third_party_abseil_cpp_absl_profiling_exponential_biased
            cronet_aml_third_party_abseil_cpp_absl_random_distributions
            cronet_aml_third_party_abseil_cpp_absl_random_internal_platform
            cronet_aml_third_party_abseil_cpp_absl_random_internal_pool_urbg
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_hwaes
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_hwaes_impl
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_slow
            cronet_aml_third_party_abseil_cpp_absl_random_internal_seed_material
            cronet_aml_third_party_abseil_cpp_absl_random_seed_gen_exception
            cronet_aml_third_party_abseil_cpp_absl_random_seed_sequences
            cronet_aml_third_party_abseil_cpp_absl_status_status
            cronet_aml_third_party_abseil_cpp_absl_status_statusor
            cronet_aml_third_party_abseil_cpp_absl_strings_cord
            cronet_aml_third_party_abseil_cpp_absl_strings_cord_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_functions
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_handle
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_info
            cronet_aml_third_party_abseil_cpp_absl_strings_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_str_format_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_strings
            cronet_aml_third_party_abseil_cpp_absl_synchronization_graphcycles_internal
            cronet_aml_third_party_abseil_cpp_absl_synchronization_kernel_timeout_internal
            cronet_aml_third_party_abseil_cpp_absl_synchronization_synchronization
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_civil_time
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_time_zone
            cronet_aml_third_party_abseil_cpp_absl_time_time
            cronet_aml_third_party_abseil_cpp_absl_types_bad_optional_access
            cronet_aml_third_party_abseil_cpp_absl_types_bad_variant_access
            cronet_aml_third_party_android_ndk_cpu_features
            cronet_aml_third_party_ashmem_ashmem
            cronet_aml_third_party_icu_icui18n
            cronet_aml_third_party_icu_icuuc_private
            cronet_aml_third_party_libevent_libevent
            cronet_aml_third_party_modp_b64_modp_b64
            jni_headers
            libandroid
            libc
            libclang_rt.builtins
            libcrypto
            libdl
            liblog
            libm
            libmediandk
            libssl
    )
endif ()
#################### cronet_aml_components_cronet_metrics_util ####################

#################### cronet_aml_third_party_abseil_cpp_absl_hash_low_level_hash ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/hash/internal/low_level_hash.cc
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
    add_library(cronet_aml_third_party_abseil_cpp_absl_hash_low_level_hash SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_abseil_cpp_absl_hash_low_level_hash
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_abseil_cpp_absl_hash_low_level_hash PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_abseil_cpp_absl_hash_low_level_hash
            cronet_aml_third_party_abseil_cpp_absl_base_base
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait
            cronet_aml_third_party_abseil_cpp_absl_numeric_int128
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_abseil_cpp_absl_hash_low_level_hash ####################

#################### cronet_aml_third_party_abseil_cpp_absl_base_malloc_internal ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/base/internal/low_level_alloc.cc
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
    add_library(cronet_aml_third_party_abseil_cpp_absl_base_malloc_internal SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_abseil_cpp_absl_base_malloc_internal
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_abseil_cpp_absl_base_malloc_internal PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_abseil_cpp_absl_base_malloc_internal
            cronet_aml_third_party_abseil_cpp_absl_base_base
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_abseil_cpp_absl_base_malloc_internal ####################

#################### cronet_aml_third_party_abseil_cpp_absl_base_throw_delegate ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/base/internal/throw_delegate.cc
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
    add_library(cronet_aml_third_party_abseil_cpp_absl_base_throw_delegate SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_abseil_cpp_absl_base_throw_delegate
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_abseil_cpp_absl_base_throw_delegate PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_abseil_cpp_absl_base_throw_delegate
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_abseil_cpp_absl_base_throw_delegate ####################

#################### cronet_aml_components_cronet_android_cronet_static ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_components_cronet_android_request_context_config_proto_gen/android_${OUT_ARCH_CPU}/gen/external/cronet/components/cronet/android/proto/request_context_config.pb.cc
        ${ANDROID_ROOT}/external/cronet/components/cronet/android/cronet_bidirectional_stream_adapter.cc
        ${ANDROID_ROOT}/external/cronet/components/cronet/android/cronet_context_adapter.cc
        ${ANDROID_ROOT}/external/cronet/components/cronet/android/cronet_library_loader.cc
        ${ANDROID_ROOT}/external/cronet/components/cronet/android/cronet_upload_data_stream_adapter.cc
        ${ANDROID_ROOT}/external/cronet/components/cronet/android/cronet_url_request_adapter.cc
        ${ANDROID_ROOT}/external/cronet/components/cronet/android/io_buffer_with_byte_buffer.cc
        ${ANDROID_ROOT}/external/cronet/components/cronet/android/url_request_error.cc
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
    add_library(cronet_aml_components_cronet_android_cronet_static SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_components_cronet_android_cronet_static
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/overrides
            ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src
            ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/common/platform/default
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
            ${ANDROID_ROOT}/external/cronet/third_party/boringssl/src/include
            ${ANDROID_ROOT}/external/cronet/third_party/protobuf/src
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_chromecast_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_chromeos_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_debugging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_logging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_partition_alloc_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_anchor_functions_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_android_runtime_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_android_runtime_unchecked_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_base_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_build_date/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_cfi_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_clang_profiling_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_debugging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_feature_list_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_ios_cronet_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_logging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_message_pump_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_orderfile_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_power_monitor_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_profiler_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_rust_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_sanitizer_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_synchronization_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_tracing_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_blink_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_branding_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_chromecast_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_chromeos_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_config_compiler_compiler_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_components_cronet_cronet_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_components_cronet_cronet_version_header_action/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_third_party_metrics_proto_metrics_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_third_party_metrics_proto_metrics_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/protos
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_third_party_metrics_proto_metrics_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/third_party/metrics_proto
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_components_prefs_android_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_crypto_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_third_party_quiche_net_quic_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_third_party_quiche_net_quic_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/net/third_party/quiche/src
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_third_party_quiche_net_quic_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/protos
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_cronet_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_base_registry_controlled_domains_registry_controlled_domains/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_isolation_info_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_isolation_info_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/net/base
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_isolation_info_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/protos
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_net_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_url_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_net_nqe_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_net_nqe_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/net/nqe/proto
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_net_nqe_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/protos
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_third_party_quiche_net_quic_test_tools_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_third_party_quiche_net_quic_test_tools_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/net/third_party/quiche/src
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_third_party_quiche_net_quic_test_tools_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/net/third_party/quiche/src/quiche/quic/test_tools
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_third_party_quiche_net_quic_test_tools_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/protos
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_url_url_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_components_cronet_android_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_components_cronet_android_cronet_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_components_cronet_android_cronet_jni_registration/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_components_cronet_android_request_context_config_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_components_cronet_android_request_context_config_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/components/cronet/android/proto
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_components_cronet_android_request_context_config_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/protos
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_components_cronet_android_cronet_static PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_components_cronet_android_cronet_static
            cronet_aml_base_allocator_partition_allocator_partition_alloc
            cronet_aml_base_base
            cronet_aml_base_base_static
            cronet_aml_base_nodebug_assertion
            cronet_aml_base_third_party_double_conversion_double_conversion
            cronet_aml_base_third_party_dynamic_annotations_dynamic_annotations
            cronet_aml_components_cronet_cronet_common
            cronet_aml_components_cronet_cronet_version_header
            cronet_aml_components_cronet_metrics_util
            cronet_aml_components_metrics_library_support
            cronet_aml_components_prefs_prefs
            cronet_aml_components_prefs_prefs_features
            cronet_aml_crypto_crypto
            cronet_aml_net_dns_dns
            cronet_aml_net_dns_public_public
            cronet_aml_net_http_transport_security_state_generated_files
            cronet_aml_net_net
            cronet_aml_net_net_deps
            cronet_aml_net_net_public_deps
            cronet_aml_net_preload_decoder
            cronet_aml_net_third_party_quiche_quiche
            cronet_aml_net_traffic_annotation_traffic_annotation
            cronet_aml_net_uri_template
            cronet_aml_third_party_abseil_cpp_absl_base_base
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_malloc_internal
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait
            cronet_aml_third_party_abseil_cpp_absl_base_strerror
            cronet_aml_third_party_abseil_cpp_absl_base_throw_delegate
            cronet_aml_third_party_abseil_cpp_absl_container_hashtablez_sampler
            cronet_aml_third_party_abseil_cpp_absl_container_raw_hash_set
            cronet_aml_third_party_abseil_cpp_absl_crc_cpu_detect
            cronet_aml_third_party_abseil_cpp_absl_crc_crc32c
            cronet_aml_third_party_abseil_cpp_absl_crc_crc_cord_state
            cronet_aml_third_party_abseil_cpp_absl_crc_crc_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_debugging_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_demangle_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_examine_stack
            cronet_aml_third_party_abseil_cpp_absl_debugging_failure_signal_handler
            cronet_aml_third_party_abseil_cpp_absl_debugging_stacktrace
            cronet_aml_third_party_abseil_cpp_absl_debugging_symbolize
            cronet_aml_third_party_abseil_cpp_absl_hash_city
            cronet_aml_third_party_abseil_cpp_absl_hash_hash
            cronet_aml_third_party_abseil_cpp_absl_hash_low_level_hash
            cronet_aml_third_party_abseil_cpp_absl_log_die_if_null
            cronet_aml_third_party_abseil_cpp_absl_log_globals
            cronet_aml_third_party_abseil_cpp_absl_log_internal_check_op
            cronet_aml_third_party_abseil_cpp_absl_log_internal_conditions
            cronet_aml_third_party_abseil_cpp_absl_log_internal_format
            cronet_aml_third_party_abseil_cpp_absl_log_internal_globals
            cronet_aml_third_party_abseil_cpp_absl_log_internal_log_message
            cronet_aml_third_party_abseil_cpp_absl_log_internal_log_sink_set
            cronet_aml_third_party_abseil_cpp_absl_log_internal_nullguard
            cronet_aml_third_party_abseil_cpp_absl_log_internal_proto
            cronet_aml_third_party_abseil_cpp_absl_log_log_entry
            cronet_aml_third_party_abseil_cpp_absl_log_log_sink
            cronet_aml_third_party_abseil_cpp_absl_numeric_int128
            cronet_aml_third_party_abseil_cpp_absl_profiling_exponential_biased
            cronet_aml_third_party_abseil_cpp_absl_random_distributions
            cronet_aml_third_party_abseil_cpp_absl_random_internal_platform
            cronet_aml_third_party_abseil_cpp_absl_random_internal_pool_urbg
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_hwaes
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_hwaes_impl
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_slow
            cronet_aml_third_party_abseil_cpp_absl_random_internal_seed_material
            cronet_aml_third_party_abseil_cpp_absl_random_seed_gen_exception
            cronet_aml_third_party_abseil_cpp_absl_random_seed_sequences
            cronet_aml_third_party_abseil_cpp_absl_status_status
            cronet_aml_third_party_abseil_cpp_absl_status_statusor
            cronet_aml_third_party_abseil_cpp_absl_strings_cord
            cronet_aml_third_party_abseil_cpp_absl_strings_cord_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_functions
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_handle
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_info
            cronet_aml_third_party_abseil_cpp_absl_strings_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_str_format_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_strings
            cronet_aml_third_party_abseil_cpp_absl_synchronization_graphcycles_internal
            cronet_aml_third_party_abseil_cpp_absl_synchronization_kernel_timeout_internal
            cronet_aml_third_party_abseil_cpp_absl_synchronization_synchronization
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_civil_time
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_time_zone
            cronet_aml_third_party_abseil_cpp_absl_time_time
            cronet_aml_third_party_abseil_cpp_absl_types_bad_optional_access
            cronet_aml_third_party_abseil_cpp_absl_types_bad_variant_access
            cronet_aml_third_party_android_ndk_cpu_features
            cronet_aml_third_party_ashmem_ashmem
            cronet_aml_third_party_brotli_common
            cronet_aml_third_party_brotli_dec
            cronet_aml_third_party_icu_icui18n
            cronet_aml_third_party_icu_icuuc_private
            cronet_aml_third_party_libevent_libevent
            cronet_aml_third_party_modp_b64_modp_b64
            cronet_aml_third_party_protobuf_protobuf_lite
            cronet_aml_url_url
            jni_headers
            libandroid
            libc
            libclang_rt.builtins
            libcrypto
            libdl
            liblog
            libm
            libmediandk
            libprotobuf-cpp-lite
            libssl
            libz
    )
endif ()
#################### cronet_aml_components_cronet_android_cronet_static ####################

#################### cronet_aml_base_allocator_partition_allocator_partition_alloc ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/base/allocator/partition_allocator/address_pool_manager.cc
        ${ANDROID_ROOT}/external/cronet/base/allocator/partition_allocator/address_pool_manager_bitmap.cc
        ${ANDROID_ROOT}/external/cronet/base/allocator/partition_allocator/address_space_randomization.cc
        ${ANDROID_ROOT}/external/cronet/base/allocator/partition_allocator/allocation_guard.cc
        ${ANDROID_ROOT}/external/cronet/base/allocator/partition_allocator/compressed_pointer.cc
        ${ANDROID_ROOT}/external/cronet/base/allocator/partition_allocator/dangling_raw_ptr_checks.cc
        ${ANDROID_ROOT}/external/cronet/base/allocator/partition_allocator/gwp_asan_support.cc
        ${ANDROID_ROOT}/external/cronet/base/allocator/partition_allocator/memory_reclaimer.cc
        ${ANDROID_ROOT}/external/cronet/base/allocator/partition_allocator/oom.cc
        ${ANDROID_ROOT}/external/cronet/base/allocator/partition_allocator/oom_callback.cc
        ${ANDROID_ROOT}/external/cronet/base/allocator/partition_allocator/page_allocator.cc
        ${ANDROID_ROOT}/external/cronet/base/allocator/partition_allocator/page_allocator_internals_posix.cc
        ${ANDROID_ROOT}/external/cronet/base/allocator/partition_allocator/partition_address_space.cc
        ${ANDROID_ROOT}/external/cronet/base/allocator/partition_allocator/partition_alloc.cc
        ${ANDROID_ROOT}/external/cronet/base/allocator/partition_allocator/partition_alloc_base/check.cc
        ${ANDROID_ROOT}/external/cronet/base/allocator/partition_allocator/partition_alloc_base/cpu.cc
        ${ANDROID_ROOT}/external/cronet/base/allocator/partition_allocator/partition_alloc_base/debug/alias.cc
        ${ANDROID_ROOT}/external/cronet/base/allocator/partition_allocator/partition_alloc_base/files/file_path.cc
        ${ANDROID_ROOT}/external/cronet/base/allocator/partition_allocator/partition_alloc_base/files/file_util_posix.cc
        ${ANDROID_ROOT}/external/cronet/base/allocator/partition_allocator/partition_alloc_base/logging.cc
        ${ANDROID_ROOT}/external/cronet/base/allocator/partition_allocator/partition_alloc_base/memory/ref_counted.cc
        ${ANDROID_ROOT}/external/cronet/base/allocator/partition_allocator/partition_alloc_base/native_library.cc
        ${ANDROID_ROOT}/external/cronet/base/allocator/partition_allocator/partition_alloc_base/native_library_posix.cc
        ${ANDROID_ROOT}/external/cronet/base/allocator/partition_allocator/partition_alloc_base/posix/safe_strerror.cc
        ${ANDROID_ROOT}/external/cronet/base/allocator/partition_allocator/partition_alloc_base/rand_util.cc
        ${ANDROID_ROOT}/external/cronet/base/allocator/partition_allocator/partition_alloc_base/rand_util_posix.cc
        ${ANDROID_ROOT}/external/cronet/base/allocator/partition_allocator/partition_alloc_base/strings/stringprintf.cc
        ${ANDROID_ROOT}/external/cronet/base/allocator/partition_allocator/partition_alloc_base/threading/platform_thread.cc
        ${ANDROID_ROOT}/external/cronet/base/allocator/partition_allocator/partition_alloc_base/threading/platform_thread_posix.cc
        ${ANDROID_ROOT}/external/cronet/base/allocator/partition_allocator/partition_alloc_base/time/time.cc
        ${ANDROID_ROOT}/external/cronet/base/allocator/partition_allocator/partition_alloc_base/time/time_android.cc
        ${ANDROID_ROOT}/external/cronet/base/allocator/partition_allocator/partition_alloc_base/time/time_conversion_posix.cc
        ${ANDROID_ROOT}/external/cronet/base/allocator/partition_allocator/partition_alloc_base/time/time_now_posix.cc
        ${ANDROID_ROOT}/external/cronet/base/allocator/partition_allocator/partition_alloc_base/time/time_override.cc
        ${ANDROID_ROOT}/external/cronet/base/allocator/partition_allocator/partition_alloc_hooks.cc
        ${ANDROID_ROOT}/external/cronet/base/allocator/partition_allocator/partition_bucket.cc
        ${ANDROID_ROOT}/external/cronet/base/allocator/partition_allocator/partition_oom.cc
        ${ANDROID_ROOT}/external/cronet/base/allocator/partition_allocator/partition_page.cc
        ${ANDROID_ROOT}/external/cronet/base/allocator/partition_allocator/partition_root.cc
        ${ANDROID_ROOT}/external/cronet/base/allocator/partition_allocator/partition_stats.cc
        ${ANDROID_ROOT}/external/cronet/base/allocator/partition_allocator/pkey.cc
        ${ANDROID_ROOT}/external/cronet/base/allocator/partition_allocator/random.cc
        ${ANDROID_ROOT}/external/cronet/base/allocator/partition_allocator/reservation_offset_table.cc
        ${ANDROID_ROOT}/external/cronet/base/allocator/partition_allocator/spinning_mutex.cc
        ${ANDROID_ROOT}/external/cronet/base/allocator/partition_allocator/tagging.cc
        ${ANDROID_ROOT}/external/cronet/base/allocator/partition_allocator/thread_cache.cc
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
    add_library(cronet_aml_base_allocator_partition_allocator_partition_alloc SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_base_allocator_partition_allocator_partition_alloc
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/android_ndk/sources/android/cpufeatures
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_chromecast_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_chromeos_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_debugging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_logging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_partition_alloc_buildflags/android_${OUT_ARCH_CPU}/gen
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_base_allocator_partition_allocator_partition_alloc PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_base_allocator_partition_allocator_partition_alloc
            cronet_aml_third_party_android_ndk_cpu_features
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libgtest_prod_headers
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_base_allocator_partition_allocator_partition_alloc ####################

#################### cronet_aml_third_party_abseil_cpp_absl_debugging_stacktrace ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/debugging/stacktrace.cc
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
    add_library(cronet_aml_third_party_abseil_cpp_absl_debugging_stacktrace SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_abseil_cpp_absl_debugging_stacktrace
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_abseil_cpp_absl_debugging_stacktrace PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_abseil_cpp_absl_debugging_stacktrace
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_debugging_internal
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_abseil_cpp_absl_debugging_stacktrace ####################

#################### cronet_aml_net_traffic_annotation_traffic_annotation ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/net/traffic_annotation/network_traffic_annotation_android.cc
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
    add_library(cronet_aml_net_traffic_annotation_traffic_annotation SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_net_traffic_annotation_traffic_annotation
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
            ${ANDROID_ROOT}/external/cronet/third_party/boringssl/src/include
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_chromecast_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_chromeos_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_debugging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_logging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_partition_alloc_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_anchor_functions_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_android_runtime_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_android_runtime_unchecked_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_base_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_build_date/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_cfi_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_clang_profiling_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_debugging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_feature_list_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_ios_cronet_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_logging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_message_pump_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_orderfile_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_power_monitor_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_profiler_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_rust_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_sanitizer_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_synchronization_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_tracing_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_blink_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_branding_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_chromecast_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_chromeos_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_config_compiler_compiler_buildflags/android_${OUT_ARCH_CPU}/gen
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_net_traffic_annotation_traffic_annotation PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_net_traffic_annotation_traffic_annotation
            cronet_aml_base_allocator_partition_allocator_partition_alloc
            cronet_aml_base_base
            cronet_aml_base_base_static
            cronet_aml_base_nodebug_assertion
            cronet_aml_base_third_party_double_conversion_double_conversion
            cronet_aml_base_third_party_dynamic_annotations_dynamic_annotations
            cronet_aml_third_party_abseil_cpp_absl_base_base
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_malloc_internal
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait
            cronet_aml_third_party_abseil_cpp_absl_base_strerror
            cronet_aml_third_party_abseil_cpp_absl_base_throw_delegate
            cronet_aml_third_party_abseil_cpp_absl_container_hashtablez_sampler
            cronet_aml_third_party_abseil_cpp_absl_container_raw_hash_set
            cronet_aml_third_party_abseil_cpp_absl_crc_cpu_detect
            cronet_aml_third_party_abseil_cpp_absl_crc_crc32c
            cronet_aml_third_party_abseil_cpp_absl_crc_crc_cord_state
            cronet_aml_third_party_abseil_cpp_absl_crc_crc_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_debugging_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_demangle_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_examine_stack
            cronet_aml_third_party_abseil_cpp_absl_debugging_failure_signal_handler
            cronet_aml_third_party_abseil_cpp_absl_debugging_stacktrace
            cronet_aml_third_party_abseil_cpp_absl_debugging_symbolize
            cronet_aml_third_party_abseil_cpp_absl_hash_city
            cronet_aml_third_party_abseil_cpp_absl_hash_hash
            cronet_aml_third_party_abseil_cpp_absl_hash_low_level_hash
            cronet_aml_third_party_abseil_cpp_absl_log_die_if_null
            cronet_aml_third_party_abseil_cpp_absl_log_globals
            cronet_aml_third_party_abseil_cpp_absl_log_internal_check_op
            cronet_aml_third_party_abseil_cpp_absl_log_internal_conditions
            cronet_aml_third_party_abseil_cpp_absl_log_internal_format
            cronet_aml_third_party_abseil_cpp_absl_log_internal_globals
            cronet_aml_third_party_abseil_cpp_absl_log_internal_log_message
            cronet_aml_third_party_abseil_cpp_absl_log_internal_log_sink_set
            cronet_aml_third_party_abseil_cpp_absl_log_internal_nullguard
            cronet_aml_third_party_abseil_cpp_absl_log_internal_proto
            cronet_aml_third_party_abseil_cpp_absl_log_log_entry
            cronet_aml_third_party_abseil_cpp_absl_log_log_sink
            cronet_aml_third_party_abseil_cpp_absl_numeric_int128
            cronet_aml_third_party_abseil_cpp_absl_profiling_exponential_biased
            cronet_aml_third_party_abseil_cpp_absl_random_distributions
            cronet_aml_third_party_abseil_cpp_absl_random_internal_platform
            cronet_aml_third_party_abseil_cpp_absl_random_internal_pool_urbg
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_hwaes
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_hwaes_impl
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_slow
            cronet_aml_third_party_abseil_cpp_absl_random_internal_seed_material
            cronet_aml_third_party_abseil_cpp_absl_random_seed_gen_exception
            cronet_aml_third_party_abseil_cpp_absl_random_seed_sequences
            cronet_aml_third_party_abseil_cpp_absl_status_status
            cronet_aml_third_party_abseil_cpp_absl_status_statusor
            cronet_aml_third_party_abseil_cpp_absl_strings_cord
            cronet_aml_third_party_abseil_cpp_absl_strings_cord_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_functions
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_handle
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_info
            cronet_aml_third_party_abseil_cpp_absl_strings_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_str_format_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_strings
            cronet_aml_third_party_abseil_cpp_absl_synchronization_graphcycles_internal
            cronet_aml_third_party_abseil_cpp_absl_synchronization_kernel_timeout_internal
            cronet_aml_third_party_abseil_cpp_absl_synchronization_synchronization
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_civil_time
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_time_zone
            cronet_aml_third_party_abseil_cpp_absl_time_time
            cronet_aml_third_party_abseil_cpp_absl_types_bad_optional_access
            cronet_aml_third_party_abseil_cpp_absl_types_bad_variant_access
            cronet_aml_third_party_android_ndk_cpu_features
            cronet_aml_third_party_ashmem_ashmem
            cronet_aml_third_party_icu_icui18n
            cronet_aml_third_party_icu_icuuc_private
            cronet_aml_third_party_libevent_libevent
            cronet_aml_third_party_modp_b64_modp_b64
            jni_headers
            libandroid
            libc
            libclang_rt.builtins
            libcrypto
            libdl
            liblog
            libm
            libmediandk
            libssl
    )
endif ()
#################### cronet_aml_net_traffic_annotation_traffic_annotation ####################

#################### cronet_aml_net_preload_decoder ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/net/extras/preload_data/decoder.cc
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
    add_library(cronet_aml_net_preload_decoder SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_net_preload_decoder
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
            ${ANDROID_ROOT}/external/cronet/third_party/boringssl/src/include
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_chromecast_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_chromeos_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_debugging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_logging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_partition_alloc_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_anchor_functions_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_android_runtime_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_android_runtime_unchecked_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_base_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_build_date/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_cfi_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_clang_profiling_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_debugging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_feature_list_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_ios_cronet_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_logging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_message_pump_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_orderfile_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_power_monitor_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_profiler_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_rust_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_sanitizer_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_synchronization_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_tracing_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_blink_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_branding_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_chromecast_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_chromeos_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_config_compiler_compiler_buildflags/android_${OUT_ARCH_CPU}/gen
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_net_preload_decoder PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_net_preload_decoder
            cronet_aml_base_allocator_partition_allocator_partition_alloc
            cronet_aml_base_base
            cronet_aml_base_base_static
            cronet_aml_base_nodebug_assertion
            cronet_aml_base_third_party_double_conversion_double_conversion
            cronet_aml_base_third_party_dynamic_annotations_dynamic_annotations
            cronet_aml_third_party_abseil_cpp_absl_base_base
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_malloc_internal
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait
            cronet_aml_third_party_abseil_cpp_absl_base_strerror
            cronet_aml_third_party_abseil_cpp_absl_base_throw_delegate
            cronet_aml_third_party_abseil_cpp_absl_container_hashtablez_sampler
            cronet_aml_third_party_abseil_cpp_absl_container_raw_hash_set
            cronet_aml_third_party_abseil_cpp_absl_crc_cpu_detect
            cronet_aml_third_party_abseil_cpp_absl_crc_crc32c
            cronet_aml_third_party_abseil_cpp_absl_crc_crc_cord_state
            cronet_aml_third_party_abseil_cpp_absl_crc_crc_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_debugging_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_demangle_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_examine_stack
            cronet_aml_third_party_abseil_cpp_absl_debugging_failure_signal_handler
            cronet_aml_third_party_abseil_cpp_absl_debugging_stacktrace
            cronet_aml_third_party_abseil_cpp_absl_debugging_symbolize
            cronet_aml_third_party_abseil_cpp_absl_hash_city
            cronet_aml_third_party_abseil_cpp_absl_hash_hash
            cronet_aml_third_party_abseil_cpp_absl_hash_low_level_hash
            cronet_aml_third_party_abseil_cpp_absl_log_die_if_null
            cronet_aml_third_party_abseil_cpp_absl_log_globals
            cronet_aml_third_party_abseil_cpp_absl_log_internal_check_op
            cronet_aml_third_party_abseil_cpp_absl_log_internal_conditions
            cronet_aml_third_party_abseil_cpp_absl_log_internal_format
            cronet_aml_third_party_abseil_cpp_absl_log_internal_globals
            cronet_aml_third_party_abseil_cpp_absl_log_internal_log_message
            cronet_aml_third_party_abseil_cpp_absl_log_internal_log_sink_set
            cronet_aml_third_party_abseil_cpp_absl_log_internal_nullguard
            cronet_aml_third_party_abseil_cpp_absl_log_internal_proto
            cronet_aml_third_party_abseil_cpp_absl_log_log_entry
            cronet_aml_third_party_abseil_cpp_absl_log_log_sink
            cronet_aml_third_party_abseil_cpp_absl_numeric_int128
            cronet_aml_third_party_abseil_cpp_absl_profiling_exponential_biased
            cronet_aml_third_party_abseil_cpp_absl_random_distributions
            cronet_aml_third_party_abseil_cpp_absl_random_internal_platform
            cronet_aml_third_party_abseil_cpp_absl_random_internal_pool_urbg
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_hwaes
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_hwaes_impl
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_slow
            cronet_aml_third_party_abseil_cpp_absl_random_internal_seed_material
            cronet_aml_third_party_abseil_cpp_absl_random_seed_gen_exception
            cronet_aml_third_party_abseil_cpp_absl_random_seed_sequences
            cronet_aml_third_party_abseil_cpp_absl_status_status
            cronet_aml_third_party_abseil_cpp_absl_status_statusor
            cronet_aml_third_party_abseil_cpp_absl_strings_cord
            cronet_aml_third_party_abseil_cpp_absl_strings_cord_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_functions
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_handle
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_info
            cronet_aml_third_party_abseil_cpp_absl_strings_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_str_format_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_strings
            cronet_aml_third_party_abseil_cpp_absl_synchronization_graphcycles_internal
            cronet_aml_third_party_abseil_cpp_absl_synchronization_kernel_timeout_internal
            cronet_aml_third_party_abseil_cpp_absl_synchronization_synchronization
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_civil_time
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_time_zone
            cronet_aml_third_party_abseil_cpp_absl_time_time
            cronet_aml_third_party_abseil_cpp_absl_types_bad_optional_access
            cronet_aml_third_party_abseil_cpp_absl_types_bad_variant_access
            cronet_aml_third_party_android_ndk_cpu_features
            cronet_aml_third_party_ashmem_ashmem
            cronet_aml_third_party_icu_icui18n
            cronet_aml_third_party_icu_icuuc_private
            cronet_aml_third_party_libevent_libevent
            cronet_aml_third_party_modp_b64_modp_b64
            jni_headers
            libandroid
            libc
            libclang_rt.builtins
            libcrypto
            libdl
            liblog
            libm
            libmediandk
            libssl
    )
endif ()
#################### cronet_aml_net_preload_decoder ####################

#################### cronet_aml_third_party_abseil_cpp_absl_log_internal_log_sink_set ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/log/internal/log_sink_set.cc
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
    add_library(cronet_aml_third_party_abseil_cpp_absl_log_internal_log_sink_set SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_abseil_cpp_absl_log_internal_log_sink_set
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_abseil_cpp_absl_log_internal_log_sink_set PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_abseil_cpp_absl_log_internal_log_sink_set
            cronet_aml_third_party_abseil_cpp_absl_base_base
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_malloc_internal
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait
            cronet_aml_third_party_abseil_cpp_absl_base_throw_delegate
            cronet_aml_third_party_abseil_cpp_absl_debugging_debugging_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_demangle_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_stacktrace
            cronet_aml_third_party_abseil_cpp_absl_debugging_symbolize
            cronet_aml_third_party_abseil_cpp_absl_hash_city
            cronet_aml_third_party_abseil_cpp_absl_hash_hash
            cronet_aml_third_party_abseil_cpp_absl_hash_low_level_hash
            cronet_aml_third_party_abseil_cpp_absl_log_globals
            cronet_aml_third_party_abseil_cpp_absl_log_internal_globals
            cronet_aml_third_party_abseil_cpp_absl_log_log_entry
            cronet_aml_third_party_abseil_cpp_absl_log_log_sink
            cronet_aml_third_party_abseil_cpp_absl_numeric_int128
            cronet_aml_third_party_abseil_cpp_absl_strings_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_strings
            cronet_aml_third_party_abseil_cpp_absl_synchronization_graphcycles_internal
            cronet_aml_third_party_abseil_cpp_absl_synchronization_kernel_timeout_internal
            cronet_aml_third_party_abseil_cpp_absl_synchronization_synchronization
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_civil_time
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_time_zone
            cronet_aml_third_party_abseil_cpp_absl_time_time
            cronet_aml_third_party_abseil_cpp_absl_types_bad_optional_access
            cronet_aml_third_party_abseil_cpp_absl_types_bad_variant_access
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            liblog
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_abseil_cpp_absl_log_internal_log_sink_set ####################

#################### cronet_aml_components_cronet_android_load_states_list_preprocess ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_components_cronet_android_load_states_list_rename/gen/LoadState.cc
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
    add_library(cronet_aml_components_cronet_android_load_states_list_preprocess SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_components_cronet_android_load_states_list_preprocess
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_components_cronet_android_load_states_list_preprocess PRIVATE
            -Werror
    )

    target_link_libraries(cronet_aml_components_cronet_android_load_states_list_preprocess
            libc
            libc++
            libc++demangle
            libdl
            libm
    )
endif ()
#################### cronet_aml_components_cronet_android_load_states_list_preprocess ####################

#################### cronet_aml_third_party_brotli_dec ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/brotli/dec/bit_reader.c
        ${ANDROID_ROOT}/external/cronet/third_party/brotli/dec/decode.c
        ${ANDROID_ROOT}/external/cronet/third_party/brotli/dec/huffman.c
        ${ANDROID_ROOT}/external/cronet/third_party/brotli/dec/state.c
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
    add_library(cronet_aml_third_party_brotli_dec SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_brotli_dec
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/brotli/include
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_brotli_dec PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_brotli_dec
            cronet_aml_third_party_brotli_common
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_brotli_dec ####################

#################### cronet_aml_third_party_abseil_cpp_absl_strings_strings ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/strings/ascii.cc
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/strings/charconv.cc
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/strings/escaping.cc
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/strings/internal/charconv_bigint.cc
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/strings/internal/charconv_parse.cc
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/strings/internal/damerau_levenshtein_distance.cc
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/strings/internal/memutil.cc
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/strings/internal/stringify_sink.cc
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/strings/match.cc
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/strings/numbers.cc
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/strings/str_cat.cc
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/strings/str_replace.cc
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/strings/str_split.cc
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/strings/string_view.cc
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/strings/substitute.cc
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
    add_library(cronet_aml_third_party_abseil_cpp_absl_strings_strings SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_abseil_cpp_absl_strings_strings
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_abseil_cpp_absl_strings_strings PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_abseil_cpp_absl_strings_strings
            cronet_aml_third_party_abseil_cpp_absl_base_base
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait
            cronet_aml_third_party_abseil_cpp_absl_base_throw_delegate
            cronet_aml_third_party_abseil_cpp_absl_numeric_int128
            cronet_aml_third_party_abseil_cpp_absl_strings_internal
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_abseil_cpp_absl_strings_strings ####################

#################### cronet_aml_net_third_party_quiche_quiche ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_third_party_quiche_net_quic_proto_gen/android_${OUT_ARCH_CPU}/gen/external/cronet/net/third_party/quiche/src/quiche/quic/core/proto/cached_network_parameters.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_third_party_quiche_net_quic_proto_gen/android_${OUT_ARCH_CPU}/gen/external/cronet/net/third_party/quiche/src/quiche/quic/core/proto/crypto_server_config.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_third_party_quiche_net_quic_proto_gen/android_${OUT_ARCH_CPU}/gen/external/cronet/net/third_party/quiche/src/quiche/quic/core/proto/source_address_token.pb.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/overrides/quiche_platform_impl/quiche_mutex_impl.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/overrides/quiche_platform_impl/quiche_time_utils_impl.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/overrides/quiche_platform_impl/quiche_url_utils_impl.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/binary_http/binary_http_message.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/common/capsule.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/common/platform/api/quiche_hostname_utils.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/common/platform/api/quiche_mutex.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/common/platform/default/quiche_platform_impl/quiche_flags_impl.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/common/quiche_buffer_allocator.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/common/quiche_crypto_logging.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/common/quiche_data_reader.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/common/quiche_data_writer.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/common/quiche_ip_address.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/common/quiche_ip_address_family.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/common/quiche_mem_slice_storage.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/common/quiche_random.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/common/quiche_text_utils.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/common/simple_buffer_allocator.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/common/structured_headers.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/http2/adapter/event_forwarder.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/http2/adapter/header_validator.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/http2/adapter/http2_protocol.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/http2/adapter/http2_util.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/http2/adapter/noop_header_validator.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/http2/adapter/oghttp2_adapter.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/http2/adapter/oghttp2_session.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/http2/adapter/oghttp2_util.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/http2/adapter/window_manager.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/http2/core/http2_trace_logging.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/http2/decoder/decode_buffer.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/http2/decoder/decode_http2_structures.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/http2/decoder/decode_status.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/http2/decoder/frame_decoder_state.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/http2/decoder/http2_frame_decoder.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/http2/decoder/http2_frame_decoder_listener.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/http2/decoder/http2_structure_decoder.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/http2/decoder/payload_decoders/altsvc_payload_decoder.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/http2/decoder/payload_decoders/continuation_payload_decoder.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/http2/decoder/payload_decoders/data_payload_decoder.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/http2/decoder/payload_decoders/goaway_payload_decoder.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/http2/decoder/payload_decoders/headers_payload_decoder.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/http2/decoder/payload_decoders/ping_payload_decoder.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/http2/decoder/payload_decoders/priority_payload_decoder.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/http2/decoder/payload_decoders/priority_update_payload_decoder.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/http2/decoder/payload_decoders/push_promise_payload_decoder.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/http2/decoder/payload_decoders/rst_stream_payload_decoder.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/http2/decoder/payload_decoders/settings_payload_decoder.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/http2/decoder/payload_decoders/unknown_payload_decoder.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/http2/decoder/payload_decoders/window_update_payload_decoder.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/http2/hpack/decoder/hpack_block_decoder.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/http2/hpack/decoder/hpack_decoder.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/http2/hpack/decoder/hpack_decoder_listener.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/http2/hpack/decoder/hpack_decoder_state.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/http2/hpack/decoder/hpack_decoder_string_buffer.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/http2/hpack/decoder/hpack_decoder_tables.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/http2/hpack/decoder/hpack_decoding_error.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/http2/hpack/decoder/hpack_entry_decoder.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/http2/hpack/decoder/hpack_entry_decoder_listener.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/http2/hpack/decoder/hpack_entry_type_decoder.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/http2/hpack/decoder/hpack_string_decoder.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/http2/hpack/decoder/hpack_string_decoder_listener.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/http2/hpack/decoder/hpack_whole_entry_buffer.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/http2/hpack/decoder/hpack_whole_entry_listener.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/http2/hpack/http2_hpack_constants.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/http2/hpack/huffman/hpack_huffman_decoder.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/http2/hpack/huffman/hpack_huffman_encoder.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/http2/hpack/huffman/huffman_spec_tables.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/http2/hpack/varint/hpack_varint_decoder.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/http2/hpack/varint/hpack_varint_encoder.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/http2/http2_constants.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/http2/http2_structures.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/oblivious_http/buffers/oblivious_http_request.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/oblivious_http/buffers/oblivious_http_response.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/oblivious_http/common/oblivious_http_header_key_config.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/oblivious_http/oblivious_http_client.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/oblivious_http/oblivious_http_gateway.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/congestion_control/bandwidth_sampler.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/congestion_control/bbr2_drain.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/congestion_control/bbr2_misc.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/congestion_control/bbr2_probe_bw.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/congestion_control/bbr2_probe_rtt.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/congestion_control/bbr2_sender.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/congestion_control/bbr2_startup.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/congestion_control/bbr_sender.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/congestion_control/cubic_bytes.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/congestion_control/general_loss_algorithm.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/congestion_control/hybrid_slow_start.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/congestion_control/pacing_sender.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/congestion_control/prr_sender.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/congestion_control/rtt_stats.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/congestion_control/send_algorithm_interface.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/congestion_control/tcp_cubic_sender_bytes.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/congestion_control/uber_loss_algorithm.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/crypto/aead_base_decrypter.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/crypto/aead_base_encrypter.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/crypto/aes_128_gcm_12_decrypter.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/crypto/aes_128_gcm_12_encrypter.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/crypto/aes_128_gcm_decrypter.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/crypto/aes_128_gcm_encrypter.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/crypto/aes_256_gcm_decrypter.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/crypto/aes_256_gcm_encrypter.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/crypto/aes_base_decrypter.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/crypto/aes_base_encrypter.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/crypto/cert_compressor.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/crypto/certificate_util.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/crypto/certificate_view.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/crypto/chacha20_poly1305_decrypter.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/crypto/chacha20_poly1305_encrypter.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/crypto/chacha20_poly1305_tls_decrypter.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/crypto/chacha20_poly1305_tls_encrypter.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/crypto/chacha_base_decrypter.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/crypto/chacha_base_encrypter.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/crypto/channel_id.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/crypto/client_proof_source.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/crypto/crypto_framer.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/crypto/crypto_handshake.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/crypto/crypto_handshake_message.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/crypto/crypto_secret_boxer.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/crypto/crypto_utils.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/crypto/curve25519_key_exchange.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/crypto/key_exchange.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/crypto/null_decrypter.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/crypto/null_encrypter.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/crypto/p256_key_exchange.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/crypto/proof_source.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/crypto/proof_source_x509.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/crypto/quic_client_session_cache.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/crypto/quic_compressed_certs_cache.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/crypto/quic_crypter.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/crypto/quic_crypto_client_config.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/crypto/quic_crypto_proof.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/crypto/quic_crypto_server_config.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/crypto/quic_decrypter.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/crypto/quic_encrypter.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/crypto/quic_hkdf.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/crypto/tls_client_connection.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/crypto/tls_connection.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/crypto/tls_server_connection.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/crypto/transport_parameters.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/crypto/web_transport_fingerprint_proof_verifier.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/deterministic_connection_id_generator.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/frames/quic_ack_frame.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/frames/quic_ack_frequency_frame.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/frames/quic_blocked_frame.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/frames/quic_connection_close_frame.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/frames/quic_crypto_frame.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/frames/quic_frame.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/frames/quic_goaway_frame.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/frames/quic_handshake_done_frame.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/frames/quic_max_streams_frame.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/frames/quic_message_frame.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/frames/quic_new_connection_id_frame.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/frames/quic_new_token_frame.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/frames/quic_padding_frame.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/frames/quic_path_challenge_frame.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/frames/quic_path_response_frame.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/frames/quic_ping_frame.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/frames/quic_retire_connection_id_frame.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/frames/quic_rst_stream_frame.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/frames/quic_stop_sending_frame.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/frames/quic_stop_waiting_frame.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/frames/quic_stream_frame.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/frames/quic_streams_blocked_frame.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/frames/quic_window_update_frame.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/http/http_constants.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/http/http_decoder.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/http/http_encoder.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/http/quic_client_promised_info.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/http/quic_client_push_promise_index.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/http/quic_header_list.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/http/quic_headers_stream.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/http/quic_receive_control_stream.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/http/quic_send_control_stream.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/http/quic_server_initiated_spdy_stream.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/http/quic_server_session_base.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/http/quic_spdy_client_session.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/http/quic_spdy_client_session_base.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/http/quic_spdy_client_stream.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/http/quic_spdy_session.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/http/quic_spdy_stream.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/http/quic_spdy_stream_body_manager.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/http/spdy_server_push_utils.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/http/spdy_utils.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/http/web_transport_http3.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/http/web_transport_stream_adapter.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/legacy_quic_stream_id_manager.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/qpack/qpack_blocking_manager.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/qpack/qpack_decoded_headers_accumulator.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/qpack/qpack_decoder.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/qpack/qpack_decoder_stream_receiver.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/qpack/qpack_decoder_stream_sender.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/qpack/qpack_encoder.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/qpack/qpack_encoder_stream_receiver.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/qpack/qpack_encoder_stream_sender.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/qpack/qpack_header_table.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/qpack/qpack_index_conversions.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/qpack/qpack_instruction_decoder.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/qpack/qpack_instruction_encoder.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/qpack/qpack_instructions.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/qpack/qpack_progressive_decoder.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/qpack/qpack_receive_stream.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/qpack/qpack_required_insert_count.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/qpack/qpack_send_stream.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/qpack/qpack_static_table.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/qpack/value_splitting_header_list.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/quic_ack_listener_interface.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/quic_alarm.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/quic_bandwidth.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/quic_chaos_protector.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/quic_coalesced_packet.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/quic_config.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/quic_connection.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/quic_connection_context.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/quic_connection_id.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/quic_connection_id_manager.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/quic_connection_stats.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/quic_constants.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/quic_control_frame_manager.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/quic_crypto_client_handshaker.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/quic_crypto_client_stream.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/quic_crypto_handshaker.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/quic_crypto_server_stream.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/quic_crypto_server_stream_base.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/quic_crypto_stream.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/quic_data_reader.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/quic_data_writer.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/quic_datagram_queue.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/quic_error_codes.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/quic_flow_controller.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/quic_framer.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/quic_idle_network_detector.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/quic_mtu_discovery.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/quic_network_blackhole_detector.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/quic_packet_creator.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/quic_packet_number.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/quic_packets.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/quic_path_validator.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/quic_ping_manager.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/quic_received_packet_manager.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/quic_sent_packet_manager.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/quic_server_id.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/quic_session.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/quic_socket_address_coder.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/quic_stream.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/quic_stream_id_manager.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/quic_stream_priority.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/quic_stream_send_buffer.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/quic_stream_sequencer.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/quic_stream_sequencer_buffer.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/quic_sustained_bandwidth_recorder.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/quic_tag.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/quic_time.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/quic_transmission_info.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/quic_types.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/quic_unacked_packet_map.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/quic_utils.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/quic_version_manager.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/quic_versions.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/quic_write_blocked_list.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/tls_client_handshaker.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/tls_handshaker.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/tls_server_handshaker.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/uber_quic_stream_id_manager.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/core/uber_received_packet_manager.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/quic/platform/api/quic_socket_address.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/spdy/core/array_output_buffer.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/spdy/core/hpack/hpack_constants.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/spdy/core/hpack/hpack_decoder_adapter.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/spdy/core/hpack/hpack_encoder.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/spdy/core/hpack/hpack_entry.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/spdy/core/hpack/hpack_header_table.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/spdy/core/hpack/hpack_output_stream.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/spdy/core/hpack/hpack_static_table.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/spdy/core/http2_frame_decoder_adapter.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/spdy/core/http2_header_block.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/spdy/core/http2_header_storage.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/spdy/core/recording_headers_handler.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/spdy/core/spdy_alt_svc_wire_format.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/spdy/core/spdy_frame_builder.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/spdy/core/spdy_framer.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/spdy/core/spdy_no_op_visitor.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/spdy/core/spdy_pinnable_buffer_piece.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/spdy/core/spdy_prefixed_buffer_reader.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/spdy/core/spdy_protocol.cc
        ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/spdy/core/spdy_simple_arena.cc
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
    add_library(cronet_aml_net_third_party_quiche_quiche SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_net_third_party_quiche_quiche
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/overrides
            ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src
            ${ANDROID_ROOT}/external/cronet/net/third_party/quiche/src/quiche/common/platform/default
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
            ${ANDROID_ROOT}/external/cronet/third_party/boringssl/src/include
            ${ANDROID_ROOT}/external/cronet/third_party/protobuf/src
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_chromecast_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_chromeos_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_debugging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_logging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_partition_allocator_partition_alloc_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_allocator_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_anchor_functions_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_android_runtime_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_android_runtime_unchecked_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_base_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_build_date/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_cfi_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_clang_profiling_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_debugging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_feature_list_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_ios_cronet_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_logging_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_message_pump_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_orderfile_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_power_monitor_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_profiler_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_rust_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_sanitizer_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_synchronization_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_base_tracing_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_blink_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_branding_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_chromecast_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_chromeos_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_config_compiler_compiler_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_url_buildflags/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_url_url_jni_headers/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_third_party_quiche_net_quic_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_third_party_quiche_net_quic_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/net/third_party/quiche/src
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_net_third_party_quiche_net_quic_proto_gen_headers/android_${OUT_ARCH_CPU}/gen/external/cronet/protos
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_net_third_party_quiche_quiche PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_net_third_party_quiche_quiche
            cronet_aml_base_allocator_partition_allocator_partition_alloc
            cronet_aml_base_base
            cronet_aml_base_base_static
            cronet_aml_base_nodebug_assertion
            cronet_aml_base_third_party_double_conversion_double_conversion
            cronet_aml_base_third_party_dynamic_annotations_dynamic_annotations
            cronet_aml_net_uri_template
            cronet_aml_third_party_abseil_cpp_absl_base_base
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_malloc_internal
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait
            cronet_aml_third_party_abseil_cpp_absl_base_strerror
            cronet_aml_third_party_abseil_cpp_absl_base_throw_delegate
            cronet_aml_third_party_abseil_cpp_absl_container_hashtablez_sampler
            cronet_aml_third_party_abseil_cpp_absl_container_raw_hash_set
            cronet_aml_third_party_abseil_cpp_absl_crc_cpu_detect
            cronet_aml_third_party_abseil_cpp_absl_crc_crc32c
            cronet_aml_third_party_abseil_cpp_absl_crc_crc_cord_state
            cronet_aml_third_party_abseil_cpp_absl_crc_crc_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_debugging_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_demangle_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_examine_stack
            cronet_aml_third_party_abseil_cpp_absl_debugging_failure_signal_handler
            cronet_aml_third_party_abseil_cpp_absl_debugging_stacktrace
            cronet_aml_third_party_abseil_cpp_absl_debugging_symbolize
            cronet_aml_third_party_abseil_cpp_absl_hash_city
            cronet_aml_third_party_abseil_cpp_absl_hash_hash
            cronet_aml_third_party_abseil_cpp_absl_hash_low_level_hash
            cronet_aml_third_party_abseil_cpp_absl_log_die_if_null
            cronet_aml_third_party_abseil_cpp_absl_log_globals
            cronet_aml_third_party_abseil_cpp_absl_log_internal_check_op
            cronet_aml_third_party_abseil_cpp_absl_log_internal_conditions
            cronet_aml_third_party_abseil_cpp_absl_log_internal_format
            cronet_aml_third_party_abseil_cpp_absl_log_internal_globals
            cronet_aml_third_party_abseil_cpp_absl_log_internal_log_message
            cronet_aml_third_party_abseil_cpp_absl_log_internal_log_sink_set
            cronet_aml_third_party_abseil_cpp_absl_log_internal_nullguard
            cronet_aml_third_party_abseil_cpp_absl_log_internal_proto
            cronet_aml_third_party_abseil_cpp_absl_log_log_entry
            cronet_aml_third_party_abseil_cpp_absl_log_log_sink
            cronet_aml_third_party_abseil_cpp_absl_numeric_int128
            cronet_aml_third_party_abseil_cpp_absl_profiling_exponential_biased
            cronet_aml_third_party_abseil_cpp_absl_random_distributions
            cronet_aml_third_party_abseil_cpp_absl_random_internal_platform
            cronet_aml_third_party_abseil_cpp_absl_random_internal_pool_urbg
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_hwaes
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_hwaes_impl
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_slow
            cronet_aml_third_party_abseil_cpp_absl_random_internal_seed_material
            cronet_aml_third_party_abseil_cpp_absl_random_seed_gen_exception
            cronet_aml_third_party_abseil_cpp_absl_random_seed_sequences
            cronet_aml_third_party_abseil_cpp_absl_status_status
            cronet_aml_third_party_abseil_cpp_absl_status_statusor
            cronet_aml_third_party_abseil_cpp_absl_strings_cord
            cronet_aml_third_party_abseil_cpp_absl_strings_cord_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_functions
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_handle
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_info
            cronet_aml_third_party_abseil_cpp_absl_strings_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_str_format_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_strings
            cronet_aml_third_party_abseil_cpp_absl_synchronization_graphcycles_internal
            cronet_aml_third_party_abseil_cpp_absl_synchronization_kernel_timeout_internal
            cronet_aml_third_party_abseil_cpp_absl_synchronization_synchronization
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_civil_time
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_time_zone
            cronet_aml_third_party_abseil_cpp_absl_time_time
            cronet_aml_third_party_abseil_cpp_absl_types_bad_optional_access
            cronet_aml_third_party_abseil_cpp_absl_types_bad_variant_access
            cronet_aml_third_party_android_ndk_cpu_features
            cronet_aml_third_party_ashmem_ashmem
            cronet_aml_third_party_icu_icui18n
            cronet_aml_third_party_icu_icuuc_private
            cronet_aml_third_party_libevent_libevent
            cronet_aml_third_party_modp_b64_modp_b64
            cronet_aml_third_party_protobuf_protobuf_lite
            cronet_aml_url_url
            jni_headers
            libandroid
            libc
            libclang_rt.builtins
            libcrypto
            libdl
            liblog
            libm
            libmediandk
            libprotobuf-cpp-lite
            libssl
            libz
    )
endif ()
#################### cronet_aml_net_third_party_quiche_quiche ####################

#################### cronet_aml_third_party_abseil_cpp_absl_strings_str_format_internal ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/strings/internal/str_format/arg.cc
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/strings/internal/str_format/bind.cc
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/strings/internal/str_format/extension.cc
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/strings/internal/str_format/float_conversion.cc
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/strings/internal/str_format/output.cc
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/strings/internal/str_format/parser.cc
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
    add_library(cronet_aml_third_party_abseil_cpp_absl_strings_str_format_internal SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_abseil_cpp_absl_strings_str_format_internal
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_abseil_cpp_absl_strings_str_format_internal PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_abseil_cpp_absl_strings_str_format_internal
            cronet_aml_third_party_abseil_cpp_absl_base_base
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait
            cronet_aml_third_party_abseil_cpp_absl_base_throw_delegate
            cronet_aml_third_party_abseil_cpp_absl_numeric_int128
            cronet_aml_third_party_abseil_cpp_absl_strings_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_strings
            cronet_aml_third_party_abseil_cpp_absl_types_bad_optional_access
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_abseil_cpp_absl_strings_str_format_internal ####################

#################### cronet_aml_third_party_abseil_cpp_absl_strings_cord_internal ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/strings/internal/cord_internal.cc
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/strings/internal/cord_rep_btree.cc
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/strings/internal/cord_rep_btree_navigator.cc
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/strings/internal/cord_rep_btree_reader.cc
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/strings/internal/cord_rep_consume.cc
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/strings/internal/cord_rep_crc.cc
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/strings/internal/cord_rep_ring.cc
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
    add_library(cronet_aml_third_party_abseil_cpp_absl_strings_cord_internal SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_abseil_cpp_absl_strings_cord_internal
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_abseil_cpp_absl_strings_cord_internal PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_abseil_cpp_absl_strings_cord_internal
            cronet_aml_third_party_abseil_cpp_absl_base_base
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait
            cronet_aml_third_party_abseil_cpp_absl_base_throw_delegate
            cronet_aml_third_party_abseil_cpp_absl_crc_cpu_detect
            cronet_aml_third_party_abseil_cpp_absl_crc_crc32c
            cronet_aml_third_party_abseil_cpp_absl_crc_crc_cord_state
            cronet_aml_third_party_abseil_cpp_absl_crc_crc_internal
            cronet_aml_third_party_abseil_cpp_absl_numeric_int128
            cronet_aml_third_party_abseil_cpp_absl_strings_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_strings
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_abseil_cpp_absl_strings_cord_internal ####################

#################### cronet_aml_third_party_abseil_cpp_absl_debugging_demangle_internal ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/debugging/internal/demangle.cc
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
    add_library(cronet_aml_third_party_abseil_cpp_absl_debugging_demangle_internal SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_abseil_cpp_absl_debugging_demangle_internal
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_abseil_cpp_absl_debugging_demangle_internal PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_abseil_cpp_absl_debugging_demangle_internal
            cronet_aml_third_party_abseil_cpp_absl_base_base
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_abseil_cpp_absl_debugging_demangle_internal ####################

#################### cronet_aml_third_party_abseil_cpp_absl_random_internal_randen ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/random/internal/randen.cc
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
    add_library(cronet_aml_third_party_abseil_cpp_absl_random_internal_randen SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_abseil_cpp_absl_random_internal_randen
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
            ${OUT_INTERMEDIATES_ROOT}/external/cronet/cronet_aml_build_chromeos_buildflags/android_${OUT_ARCH_CPU}/gen
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_abseil_cpp_absl_random_internal_randen PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_abseil_cpp_absl_random_internal_randen
            cronet_aml_third_party_abseil_cpp_absl_base_base
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait
            cronet_aml_third_party_abseil_cpp_absl_numeric_int128
            cronet_aml_third_party_abseil_cpp_absl_random_internal_platform
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_hwaes
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_hwaes_impl
            cronet_aml_third_party_abseil_cpp_absl_random_internal_randen_slow
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_abseil_cpp_absl_random_internal_randen ####################

#################### cronet_aml_third_party_abseil_cpp_absl_log_internal_proto ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/log/internal/proto.cc
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
    add_library(cronet_aml_third_party_abseil_cpp_absl_log_internal_proto SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_abseil_cpp_absl_log_internal_proto
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_abseil_cpp_absl_log_internal_proto PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_abseil_cpp_absl_log_internal_proto
            cronet_aml_third_party_abseil_cpp_absl_base_base
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait
            cronet_aml_third_party_abseil_cpp_absl_base_throw_delegate
            cronet_aml_third_party_abseil_cpp_absl_numeric_int128
            cronet_aml_third_party_abseil_cpp_absl_strings_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_strings
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_abseil_cpp_absl_log_internal_proto ####################

#################### cronet_aml_third_party_abseil_cpp_absl_status_status ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/status/status.cc
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/status/status_payload_printer.cc
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
    add_library(cronet_aml_third_party_abseil_cpp_absl_status_status SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_abseil_cpp_absl_status_status
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_abseil_cpp_absl_status_status PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_abseil_cpp_absl_status_status
            cronet_aml_third_party_abseil_cpp_absl_base_base
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_malloc_internal
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait
            cronet_aml_third_party_abseil_cpp_absl_base_strerror
            cronet_aml_third_party_abseil_cpp_absl_base_throw_delegate
            cronet_aml_third_party_abseil_cpp_absl_crc_cpu_detect
            cronet_aml_third_party_abseil_cpp_absl_crc_crc32c
            cronet_aml_third_party_abseil_cpp_absl_crc_crc_cord_state
            cronet_aml_third_party_abseil_cpp_absl_crc_crc_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_debugging_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_demangle_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_stacktrace
            cronet_aml_third_party_abseil_cpp_absl_debugging_symbolize
            cronet_aml_third_party_abseil_cpp_absl_numeric_int128
            cronet_aml_third_party_abseil_cpp_absl_profiling_exponential_biased
            cronet_aml_third_party_abseil_cpp_absl_strings_cord
            cronet_aml_third_party_abseil_cpp_absl_strings_cord_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_functions
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_handle
            cronet_aml_third_party_abseil_cpp_absl_strings_cordz_info
            cronet_aml_third_party_abseil_cpp_absl_strings_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_str_format_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_strings
            cronet_aml_third_party_abseil_cpp_absl_synchronization_graphcycles_internal
            cronet_aml_third_party_abseil_cpp_absl_synchronization_kernel_timeout_internal
            cronet_aml_third_party_abseil_cpp_absl_synchronization_synchronization
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_civil_time
            cronet_aml_third_party_abseil_cpp_absl_time_internal_cctz_time_zone
            cronet_aml_third_party_abseil_cpp_absl_time_time
            cronet_aml_third_party_abseil_cpp_absl_types_bad_optional_access
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_abseil_cpp_absl_status_status ####################

#################### cronet_aml_third_party_abseil_cpp_absl_debugging_examine_stack ####################
set(FILES
        ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp/absl/debugging/internal/examine_stack.cc
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
    add_library(cronet_aml_third_party_abseil_cpp_absl_debugging_examine_stack SHARED ${SOURCE_FILES})

    target_include_directories(cronet_aml_third_party_abseil_cpp_absl_debugging_examine_stack
            PUBLIC
            ${ANDROID_ROOT}/external/cronet
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++/trunk/include
            ${ANDROID_ROOT}/external/cronet/buildtools/third_party/libc++abi/trunk/include
            ${ANDROID_ROOT}/external/cronet/third_party/abseil-cpp
    )

    # Add c/c++ flags
    target_compile_options(cronet_aml_third_party_abseil_cpp_absl_debugging_examine_stack PRIVATE
            -Werror
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(cronet_aml_third_party_abseil_cpp_absl_debugging_examine_stack
            cronet_aml_third_party_abseil_cpp_absl_base_base
            cronet_aml_third_party_abseil_cpp_absl_base_log_severity
            cronet_aml_third_party_abseil_cpp_absl_base_malloc_internal
            cronet_aml_third_party_abseil_cpp_absl_base_raw_logging_internal
            cronet_aml_third_party_abseil_cpp_absl_base_spinlock_wait
            cronet_aml_third_party_abseil_cpp_absl_base_throw_delegate
            cronet_aml_third_party_abseil_cpp_absl_debugging_debugging_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_demangle_internal
            cronet_aml_third_party_abseil_cpp_absl_debugging_stacktrace
            cronet_aml_third_party_abseil_cpp_absl_debugging_symbolize
            cronet_aml_third_party_abseil_cpp_absl_numeric_int128
            cronet_aml_third_party_abseil_cpp_absl_strings_internal
            cronet_aml_third_party_abseil_cpp_absl_strings_strings
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            libm
            libmediandk
    )
endif ()
#################### cronet_aml_third_party_abseil_cpp_absl_debugging_examine_stack ####################
