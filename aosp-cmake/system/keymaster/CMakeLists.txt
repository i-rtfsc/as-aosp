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

project(libkeymaster_fuzz_buffer)

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

#################### libkeymaster_messages ####################
set(FILES
        ${ANDROID_ROOT}/system/keymaster/android_keymaster/android_keymaster_messages.cpp
        ${ANDROID_ROOT}/system/keymaster/android_keymaster/android_keymaster_utils.cpp
        ${ANDROID_ROOT}/system/keymaster/android_keymaster/authorization_set.cpp
        ${ANDROID_ROOT}/system/keymaster/android_keymaster/keymaster_tags.cpp
        ${ANDROID_ROOT}/system/keymaster/android_keymaster/logger.cpp
        ${ANDROID_ROOT}/system/keymaster/android_keymaster/serializable.cpp
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
    add_library(libkeymaster_messages SHARED ${SOURCE_FILES})

    target_include_directories(libkeymaster_messages
            PUBLIC
            ${ANDROID_ROOT}/system/keymaster/include
            ${ANDROID_ROOT}/system/keymaster
    )

    # Add c/c++ flags
    target_compile_options(libkeymaster_messages PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libkeymaster_messages
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libhardware_headers
            libm
            libwinpthread
    )
endif ()
#################### libkeymaster_messages ####################

#################### libkeymaster_fuzz_serializable ####################
set(FILES
        ${ANDROID_ROOT}/system/keymaster/tests/fuzzers/message_serializable_fuzz.cpp
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
    add_executable(libkeymaster_fuzz_serializable ${SOURCE_FILES})

    target_include_directories(libkeymaster_fuzz_serializable
            PUBLIC
            ${ANDROID_ROOT}/system/keymaster
            ${ANDROID_ROOT}/system/keymaster/include
    )

    # Add c/c++ flags
    target_compile_options(libkeymaster_fuzz_serializable PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libkeymaster_fuzz_serializable
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            libhardware_headers
            libkeymaster_messages
            libm
    )
endif ()
#################### libkeymaster_fuzz_serializable ####################

#################### lib_android_keymaster_keymint_utils ####################
set(FILES
        ${ANDROID_ROOT}/system/keymaster/ng/KeyMintUtils.cpp
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
    add_library(lib_android_keymaster_keymint_utils SHARED ${SOURCE_FILES})

    target_include_directories(lib_android_keymaster_keymint_utils
            PUBLIC
            ${ANDROID_ROOT}/system/keymaster/ng/include
            ${ANDROID_ROOT}/system/keymaster/include
            ${ANDROID_ROOT}/system/keymaster
    )

    # Add c/c++ flags
    target_compile_options(lib_android_keymaster_keymint_utils PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(lib_android_keymaster_keymint_utils
            android.hardware.security.keymint-V3-ndk
            libbase
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libhardware
            libm
    )
endif ()
#################### lib_android_keymaster_keymint_utils ####################

#################### libpuresoftkeymasterdevice ####################
set(FILES
        ${ANDROID_ROOT}/system/keymaster/android_keymaster/keymaster_configuration.cpp
        ${ANDROID_ROOT}/system/keymaster/contexts/soft_attestation_context.cpp
        ${ANDROID_ROOT}/system/keymaster/contexts/pure_soft_keymaster_context.cpp
        ${ANDROID_ROOT}/system/keymaster/contexts/pure_soft_remote_provisioning_context.cpp
        ${ANDROID_ROOT}/system/keymaster/contexts/soft_keymaster_logger.cpp
        ${ANDROID_ROOT}/system/keymaster/km_openssl/soft_keymaster_enforcement.cpp
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
    add_library(libpuresoftkeymasterdevice SHARED ${SOURCE_FILES})

    target_include_directories(libpuresoftkeymasterdevice
            PUBLIC
            ${ANDROID_ROOT}/system/keymaster/include
            ${ANDROID_ROOT}/system/keymaster
    )

    # Add c/c++ flags
    target_compile_options(libpuresoftkeymasterdevice PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libpuresoftkeymasterdevice
            libbase
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcppbor_external
            libcppcose_rkp
            libcrypto
            libcutils
            libdl
            libkeymaster_messages
            libkeymaster_portable
            liblog
            libm
            libsoft_attestation_cert
    )
endif ()
#################### libpuresoftkeymasterdevice ####################

#################### libkeymaster_fuzz_buffer ####################
set(FILES
        ${ANDROID_ROOT}/system/keymaster/tests/fuzzers/buffer_fuzz.cpp
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
    add_executable(libkeymaster_fuzz_buffer ${SOURCE_FILES})

    target_include_directories(libkeymaster_fuzz_buffer
            PUBLIC
            ${ANDROID_ROOT}/system/keymaster
            ${ANDROID_ROOT}/system/keymaster/include
    )

    # Add c/c++ flags
    target_compile_options(libkeymaster_fuzz_buffer PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libkeymaster_fuzz_buffer
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            libhardware_headers
            libkeymaster_messages
            libm
    )
endif ()
#################### libkeymaster_fuzz_buffer ####################

#################### libsoftkeymasterdevice ####################
set(FILES
        ${ANDROID_ROOT}/system/keymaster/android_keymaster/keymaster_configuration.cpp
        ${ANDROID_ROOT}/system/keymaster/contexts/pure_soft_keymaster_context.cpp
        ${ANDROID_ROOT}/system/keymaster/contexts/pure_soft_remote_provisioning_context.cpp
        ${ANDROID_ROOT}/system/keymaster/contexts/soft_attestation_context.cpp
        ${ANDROID_ROOT}/system/keymaster/contexts/soft_keymaster_context.cpp
        ${ANDROID_ROOT}/system/keymaster/contexts/soft_keymaster_device.cpp
        ${ANDROID_ROOT}/system/keymaster/contexts/soft_keymaster_logger.cpp
        ${ANDROID_ROOT}/system/keymaster/km_openssl/soft_keymaster_enforcement.cpp
        ${ANDROID_ROOT}/system/keymaster/legacy_support/ec_keymaster1_key.cpp
        ${ANDROID_ROOT}/system/keymaster/legacy_support/ecdsa_keymaster1_operation.cpp
        ${ANDROID_ROOT}/system/keymaster/legacy_support/keymaster1_engine.cpp
        ${ANDROID_ROOT}/system/keymaster/legacy_support/keymaster1_legacy_support.cpp
        ${ANDROID_ROOT}/system/keymaster/legacy_support/rsa_keymaster1_key.cpp
        ${ANDROID_ROOT}/system/keymaster/legacy_support/rsa_keymaster1_operation.cpp
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
    add_library(libsoftkeymasterdevice SHARED ${SOURCE_FILES})

    target_include_directories(libsoftkeymasterdevice
            PUBLIC
            ${ANDROID_ROOT}/system/keymaster/include
            ${ANDROID_ROOT}/system/keymaster
    )

    # Add c/c++ flags
    target_compile_options(libsoftkeymasterdevice PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libsoftkeymasterdevice
            libbase
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcppbor_external
            libcppcose_rkp
            libcrypto
            libcutils
            libdl
            libkeymaster_messages
            libkeymaster_portable
            liblog
            libm
            libsoft_attestation_cert
    )
endif ()
#################### libsoftkeymasterdevice ####################

#################### libkeymaster3device ####################
set(FILES
        ${ANDROID_ROOT}/system/keymaster/legacy_support/keymaster_passthrough_key.cpp
        ${ANDROID_ROOT}/system/keymaster/legacy_support/keymaster_passthrough_engine.cpp
        ${ANDROID_ROOT}/system/keymaster/legacy_support/keymaster_passthrough_operation.cpp
        ${ANDROID_ROOT}/system/keymaster/contexts/keymaster1_passthrough_context.cpp
        ${ANDROID_ROOT}/system/keymaster/contexts/keymaster2_passthrough_context.cpp
        ${ANDROID_ROOT}/system/keymaster/ng/AndroidKeymaster3Device.cpp
        ${ANDROID_ROOT}/system/keymaster/android_keymaster/keymaster_configuration.cpp
        ${ANDROID_ROOT}/system/keymaster/legacy_support/ec_keymaster1_key.cpp
        ${ANDROID_ROOT}/system/keymaster/legacy_support/ecdsa_keymaster1_operation.cpp
        ${ANDROID_ROOT}/system/keymaster/legacy_support/keymaster1_engine.cpp
        ${ANDROID_ROOT}/system/keymaster/legacy_support/keymaster1_legacy_support.cpp
        ${ANDROID_ROOT}/system/keymaster/legacy_support/rsa_keymaster1_key.cpp
        ${ANDROID_ROOT}/system/keymaster/legacy_support/rsa_keymaster1_operation.cpp
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
    add_library(libkeymaster3device SHARED ${SOURCE_FILES})

    target_include_directories(libkeymaster3device
            PUBLIC
            ${ANDROID_ROOT}/system/keymaster/include
            ${ANDROID_ROOT}/system/keymaster/ng/include
            ${ANDROID_ROOT}/system/keymaster
    )

    # Add c/c++ flags
    target_compile_options(libkeymaster3device PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libkeymaster3device
            android.hardware.keymaster-3.0
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcrypto
            libcutils
            libdl
            libhidlbase
            libkeymaster_messages
            libkeymaster_portable
            liblog
            libm
            libpuresoftkeymasterdevice
            libsoft_attestation_cert
            libutils
    )
endif ()
#################### libkeymaster3device ####################

#################### libkeymaster41 ####################
set(FILES
        ${ANDROID_ROOT}/system/keymaster/ng/AndroidKeymaster41Device.cpp
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
    add_library(libkeymaster41 SHARED ${SOURCE_FILES})

    target_include_directories(libkeymaster41
            PUBLIC
            ${ANDROID_ROOT}/system/keymaster/ng/include
            ${ANDROID_ROOT}/system/keymaster
            ${ANDROID_ROOT}/system/keymaster/include
    )

    # Add c/c++ flags
    target_compile_options(libkeymaster41 PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libkeymaster41
            android.hardware.keymaster-4.0
            android.hardware.keymaster-4.1
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcrypto
            libcutils
            libdl
            libhidlbase
            libkeymaster4
            libkeymaster4_1support
            libkeymaster4support
            libkeymaster_messages
            libkeymaster_portable
            liblog
            libm
            libpuresoftkeymasterdevice
            libutils
    )
endif ()
#################### libkeymaster41 ####################

#################### libkeymint ####################
set(FILES
        ${ANDROID_ROOT}/system/keymaster/android_keymaster/keymaster_configuration.cpp
        ${ANDROID_ROOT}/system/keymaster/legacy_support/keymaster_passthrough_engine.cpp
        ${ANDROID_ROOT}/system/keymaster/legacy_support/keymaster_passthrough_key.cpp
        ${ANDROID_ROOT}/system/keymaster/legacy_support/keymaster_passthrough_operation.cpp
        ${ANDROID_ROOT}/system/keymaster/ng/AndroidKeyMintDevice.cpp
        ${ANDROID_ROOT}/system/keymaster/ng/AndroidKeyMintOperation.cpp
        ${ANDROID_ROOT}/system/keymaster/ng/AndroidRemotelyProvisionedComponentDevice.cpp
        ${ANDROID_ROOT}/system/keymaster/ng/AndroidSharedSecret.cpp
        ${ANDROID_ROOT}/system/keymaster/ng/AndroidSecureClock.cpp
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
    add_library(libkeymint SHARED ${SOURCE_FILES})

    target_include_directories(libkeymint
            PUBLIC
            ${ANDROID_ROOT}/system/keymaster/include
            ${ANDROID_ROOT}/system/keymaster/ng/include
            ${ANDROID_ROOT}/system/keymaster
    )

    # Add c/c++ flags
    target_compile_options(libkeymint PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libkeymint
            android.hardware.security.keymint-V3-ndk
            android.hardware.security.rkp-V3-ndk
            android.hardware.security.secureclock-V1-ndk
            android.hardware.security.sharedsecret-V1-ndk
            lib_android_keymaster_keymint_utils
            libbase
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcppbor_external
            libcrypto
            libcutils
            libdl
            libhidlbase
            libkeymaster_messages
            libkeymaster_portable
            liblog
            libm
            libpuresoftkeymasterdevice
            libutils
    )
endif ()
#################### libkeymint ####################

#################### libcppcose_rkp ####################
set(FILES
        ${ANDROID_ROOT}/system/keymaster/cppcose/cppcose.cpp
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
    add_library(libcppcose_rkp SHARED ${SOURCE_FILES})

    target_include_directories(libcppcose_rkp
            PUBLIC
            ${ANDROID_ROOT}/system/keymaster/include
            ${ANDROID_ROOT}/system/keymaster
    )

    # Add c/c++ flags
    target_compile_options(libcppcose_rkp PRIVATE
            -Werror
    )

    target_link_libraries(libcppcose_rkp
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcppbor_external
            libcrypto
            libdl
            liblog
            libm
            libwinpthread
    )
endif ()
#################### libcppcose_rkp ####################

#################### libkeymaster4 ####################
set(FILES
        ${ANDROID_ROOT}/system/keymaster/legacy_support/keymaster_passthrough_key.cpp
        ${ANDROID_ROOT}/system/keymaster/legacy_support/keymaster_passthrough_engine.cpp
        ${ANDROID_ROOT}/system/keymaster/legacy_support/keymaster_passthrough_operation.cpp
        ${ANDROID_ROOT}/system/keymaster/ng/AndroidKeymaster4Device.cpp
        ${ANDROID_ROOT}/system/keymaster/android_keymaster/keymaster_configuration.cpp
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
    add_library(libkeymaster4 SHARED ${SOURCE_FILES})

    target_include_directories(libkeymaster4
            PUBLIC
            ${ANDROID_ROOT}/system/keymaster/ng/include
            ${ANDROID_ROOT}/system/keymaster/include
            ${ANDROID_ROOT}/system/keymaster
    )

    # Add c/c++ flags
    target_compile_options(libkeymaster4 PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libkeymaster4
            android.hardware.keymaster-4.0
            libbase
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcrypto
            libcutils
            libdl
            libhidlbase
            libkeymaster4support
            libkeymaster_messages
            libkeymaster_portable
            liblog
            libm
            libpuresoftkeymasterdevice
            libutils
    )
endif ()
#################### libkeymaster4 ####################

#################### libkeymaster_portable ####################
set(FILES
        ${ANDROID_ROOT}/system/keymaster/android_keymaster/android_keymaster.cpp
        ${ANDROID_ROOT}/system/keymaster/android_keymaster/android_keymaster_messages.cpp
        ${ANDROID_ROOT}/system/keymaster/android_keymaster/android_keymaster_utils.cpp
        ${ANDROID_ROOT}/system/keymaster/android_keymaster/authorization_set.cpp
        ${ANDROID_ROOT}/system/keymaster/android_keymaster/keymaster_enforcement.cpp
        ${ANDROID_ROOT}/system/keymaster/android_keymaster/keymaster_tags.cpp
        ${ANDROID_ROOT}/system/keymaster/android_keymaster/logger.cpp
        ${ANDROID_ROOT}/system/keymaster/android_keymaster/operation.cpp
        ${ANDROID_ROOT}/system/keymaster/android_keymaster/operation_table.cpp
        ${ANDROID_ROOT}/system/keymaster/android_keymaster/pure_soft_secure_key_storage.cpp
        ${ANDROID_ROOT}/system/keymaster/android_keymaster/remote_provisioning_utils.cpp
        ${ANDROID_ROOT}/system/keymaster/android_keymaster/serializable.cpp
        ${ANDROID_ROOT}/system/keymaster/key_blob_utils/auth_encrypted_key_blob.cpp
        ${ANDROID_ROOT}/system/keymaster/key_blob_utils/integrity_assured_key_blob.cpp
        ${ANDROID_ROOT}/system/keymaster/key_blob_utils/ocb.c
        ${ANDROID_ROOT}/system/keymaster/key_blob_utils/ocb_utils.cpp
        ${ANDROID_ROOT}/system/keymaster/key_blob_utils/software_keyblobs.cpp
        ${ANDROID_ROOT}/system/keymaster/km_openssl/aes_key.cpp
        ${ANDROID_ROOT}/system/keymaster/km_openssl/aes_operation.cpp
        ${ANDROID_ROOT}/system/keymaster/km_openssl/asymmetric_key.cpp
        ${ANDROID_ROOT}/system/keymaster/km_openssl/asymmetric_key_factory.cpp
        ${ANDROID_ROOT}/system/keymaster/km_openssl/attestation_record.cpp
        ${ANDROID_ROOT}/system/keymaster/km_openssl/attestation_utils.cpp
        ${ANDROID_ROOT}/system/keymaster/km_openssl/block_cipher_operation.cpp
        ${ANDROID_ROOT}/system/keymaster/km_openssl/certificate_utils.cpp
        ${ANDROID_ROOT}/system/keymaster/km_openssl/ckdf.cpp
        ${ANDROID_ROOT}/system/keymaster/km_openssl/curve25519_key.cpp
        ${ANDROID_ROOT}/system/keymaster/km_openssl/ec_key.cpp
        ${ANDROID_ROOT}/system/keymaster/km_openssl/ec_key_factory.cpp
        ${ANDROID_ROOT}/system/keymaster/km_openssl/ecdh_operation.cpp
        ${ANDROID_ROOT}/system/keymaster/km_openssl/ecdsa_operation.cpp
        ${ANDROID_ROOT}/system/keymaster/km_openssl/ecies_kem.cpp
        ${ANDROID_ROOT}/system/keymaster/km_openssl/hkdf.cpp
        ${ANDROID_ROOT}/system/keymaster/km_openssl/hmac.cpp
        ${ANDROID_ROOT}/system/keymaster/km_openssl/hmac_key.cpp
        ${ANDROID_ROOT}/system/keymaster/km_openssl/hmac_operation.cpp
        ${ANDROID_ROOT}/system/keymaster/km_openssl/iso18033kdf.cpp
        ${ANDROID_ROOT}/system/keymaster/km_openssl/kdf.cpp
        ${ANDROID_ROOT}/system/keymaster/km_openssl/nist_curve_key_exchange.cpp
        ${ANDROID_ROOT}/system/keymaster/km_openssl/openssl_err.cpp
        ${ANDROID_ROOT}/system/keymaster/km_openssl/openssl_utils.cpp
        ${ANDROID_ROOT}/system/keymaster/km_openssl/rsa_key.cpp
        ${ANDROID_ROOT}/system/keymaster/km_openssl/rsa_key_factory.cpp
        ${ANDROID_ROOT}/system/keymaster/km_openssl/rsa_operation.cpp
        ${ANDROID_ROOT}/system/keymaster/km_openssl/software_random_source.cpp
        ${ANDROID_ROOT}/system/keymaster/km_openssl/symmetric_key.cpp
        ${ANDROID_ROOT}/system/keymaster/km_openssl/triple_des_key.cpp
        ${ANDROID_ROOT}/system/keymaster/km_openssl/triple_des_operation.cpp
        ${ANDROID_ROOT}/system/keymaster/km_openssl/wrapped_key.cpp
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
    add_library(libkeymaster_portable SHARED ${SOURCE_FILES})

    target_include_directories(libkeymaster_portable
            PUBLIC
            ${ANDROID_ROOT}/system/keymaster/include
            ${ANDROID_ROOT}/system/keymaster
    )

    # Add c/c++ flags
    target_compile_options(libkeymaster_portable PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libkeymaster_portable
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcppbor_external
            libcppcose_rkp
            libcrypto
            libdl
            libhardware_headers
            libm
            libwinpthread
    )
endif ()
#################### libkeymaster_portable ####################
