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

project(libopen_dice_cbor_p384_fuzzer)

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

#################### libopen_dice_cbor_p384 ####################
set(FILES
        ${ANDROID_ROOT}/external/open-dice/src/boringssl_hash_kdf_ops.c
        ${ANDROID_ROOT}/external/open-dice/src/boringssl_ecdsa_utils.c
        ${ANDROID_ROOT}/external/open-dice/src/boringssl_p384_ops.c
        ${ANDROID_ROOT}/external/open-dice/src/cbor_cert_op.c
        ${ANDROID_ROOT}/external/open-dice/src/cbor_p384_cert_op.c
        ${ANDROID_ROOT}/external/open-dice/src/cbor_writer.c
        ${ANDROID_ROOT}/external/open-dice/src/clear_memory.c
        ${ANDROID_ROOT}/external/open-dice/src/dice.c
        ${ANDROID_ROOT}/external/open-dice/src/utils.c
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
    add_library(libopen_dice_cbor_p384 SHARED ${SOURCE_FILES})

    target_include_directories(libopen_dice_cbor_p384
            PUBLIC
            ${ANDROID_ROOT}/external/open-dice
            ${ANDROID_ROOT}/external/open-dice/include/dice/config/boringssl_ecdsa_p384
            ${ANDROID_ROOT}/external/open-dice/include
    )

    # Add c/c++ flags
    target_compile_options(libopen_dice_cbor_p384 PRIVATE
            -Werror
    )

    target_link_libraries(libopen_dice_cbor_p384
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcrypto
            libdl
            libm
            libopen_dice_boringssl_ecdsa_p384_headers
            libopen_dice_headers
    )
endif ()
#################### libopen_dice_cbor_p384 ####################

#################### libopen_dice_bcc_baremetal ####################
set(FILES
        ${ANDROID_ROOT}/external/open-dice/src/cbor_reader.c
        ${ANDROID_ROOT}/external/open-dice/src/android/bcc.c
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
    add_library(libopen_dice_bcc_baremetal SHARED ${SOURCE_FILES})

    target_include_directories(libopen_dice_bcc_baremetal
            PUBLIC
            ${ANDROID_ROOT}/external/open-dice
            ${ANDROID_ROOT}/external/open-dice/include/dice/config/boringssl_ed25519
            ${ANDROID_ROOT}/external/open-dice/include
    )

    # Add c/c++ flags
    target_compile_options(libopen_dice_bcc_baremetal PRIVATE
            -Werror
    )

    target_link_libraries(libopen_dice_bcc_baremetal
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcrypto_baremetal
            libdl
            libm
            libopen_dice_cbor_baremetal
    )
endif ()
#################### libopen_dice_bcc_baremetal ####################

#################### libopen_dice_cbor ####################
set(FILES
        ${ANDROID_ROOT}/external/open-dice/src/boringssl_hash_kdf_ops.c
        ${ANDROID_ROOT}/external/open-dice/src/boringssl_ed25519_ops.c
        ${ANDROID_ROOT}/external/open-dice/src/cbor_cert_op.c
        ${ANDROID_ROOT}/external/open-dice/src/cbor_ed25519_cert_op.c
        ${ANDROID_ROOT}/external/open-dice/src/cbor_writer.c
        ${ANDROID_ROOT}/external/open-dice/src/dice.c
        ${ANDROID_ROOT}/external/open-dice/src/utils.c
        ${ANDROID_ROOT}/external/open-dice/src/clear_memory.c
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
    add_library(libopen_dice_cbor SHARED ${SOURCE_FILES})

    target_include_directories(libopen_dice_cbor
            PUBLIC
            ${ANDROID_ROOT}/external/open-dice
            ${ANDROID_ROOT}/external/open-dice/include/dice/config/boringssl_ed25519
            ${ANDROID_ROOT}/external/open-dice/include
    )

    # Add c/c++ flags
    target_compile_options(libopen_dice_cbor PRIVATE
            -Werror
    )

    target_link_libraries(libopen_dice_cbor
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcrypto
            libdl
            libm
            libopen_dice_boringssl_ed25519_headers
            libopen_dice_headers
    )
endif ()
#################### libopen_dice_cbor ####################

#################### libopen_dice_cbor_writer_p384_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/external/open-dice/src/cbor_writer_fuzzer.cc
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
    add_executable(libopen_dice_cbor_writer_p384_fuzzer ${SOURCE_FILES})

    target_include_directories(libopen_dice_cbor_writer_p384_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/external/open-dice
            ${ANDROID_ROOT}/external/open-dice/include/dice/config/boringssl_ecdsa_p384
            ${ANDROID_ROOT}/external/open-dice/include
    )

    # Add c/c++ flags
    target_compile_options(libopen_dice_cbor_writer_p384_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(libopen_dice_cbor_writer_p384_fuzzer
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            libm
            libopen_dice_cbor_p384
    )
endif ()
#################### libopen_dice_cbor_writer_p384_fuzzer ####################

#################### libopen_dice_bcc_p384_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/external/open-dice/src/android/bcc_fuzzer.cc
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
    add_executable(libopen_dice_bcc_p384_fuzzer ${SOURCE_FILES})

    target_include_directories(libopen_dice_bcc_p384_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/external/open-dice
            ${ANDROID_ROOT}/external/open-dice/include/dice/config/boringssl_ecdsa_p384
            ${ANDROID_ROOT}/external/open-dice/include
    )

    # Add c/c++ flags
    target_compile_options(libopen_dice_bcc_p384_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(libopen_dice_bcc_p384_fuzzer
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcrypto
            libdl
            libm
            libopen_dice_bcc_p384
            libopen_dice_cbor_p384
    )
endif ()
#################### libopen_dice_bcc_p384_fuzzer ####################

#################### libopen_dice_bcc_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/external/open-dice/src/android/bcc_fuzzer.cc
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
    add_executable(libopen_dice_bcc_fuzzer ${SOURCE_FILES})

    target_include_directories(libopen_dice_bcc_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/external/open-dice
            ${ANDROID_ROOT}/external/open-dice/include/dice/config/boringssl_ed25519
            ${ANDROID_ROOT}/external/open-dice/include
    )

    # Add c/c++ flags
    target_compile_options(libopen_dice_bcc_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(libopen_dice_bcc_fuzzer
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcrypto
            libdl
            libm
            libopen_dice_bcc
            libopen_dice_cbor
    )
endif ()
#################### libopen_dice_bcc_fuzzer ####################

#################### libopen_dice_bcc_p384 ####################
set(FILES
        ${ANDROID_ROOT}/external/open-dice/src/cbor_reader.c
        ${ANDROID_ROOT}/external/open-dice/src/android/bcc.c
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
    add_library(libopen_dice_bcc_p384 SHARED ${SOURCE_FILES})

    target_include_directories(libopen_dice_bcc_p384
            PUBLIC
            ${ANDROID_ROOT}/external/open-dice
            ${ANDROID_ROOT}/external/open-dice/include/dice/config/boringssl_ecdsa_p384
            ${ANDROID_ROOT}/external/open-dice/include
    )

    # Add c/c++ flags
    target_compile_options(libopen_dice_bcc_p384 PRIVATE
            -Werror
    )

    target_link_libraries(libopen_dice_bcc_p384
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcrypto
            libdl
            libm
            libopen_dice_cbor_p384
    )
endif ()
#################### libopen_dice_bcc_p384 ####################

#################### libopen_dice_bcc ####################
set(FILES
        ${ANDROID_ROOT}/external/open-dice/src/cbor_reader.c
        ${ANDROID_ROOT}/external/open-dice/src/android/bcc.c
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
    add_library(libopen_dice_bcc SHARED ${SOURCE_FILES})

    target_include_directories(libopen_dice_bcc
            PUBLIC
            ${ANDROID_ROOT}/external/open-dice
            ${ANDROID_ROOT}/external/open-dice/include/dice/config/boringssl_ed25519
            ${ANDROID_ROOT}/external/open-dice/include
    )

    # Add c/c++ flags
    target_compile_options(libopen_dice_bcc PRIVATE
            -Werror
    )

    target_link_libraries(libopen_dice_bcc
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcrypto
            libdl
            libm
            libopen_dice_cbor
    )
endif ()
#################### libopen_dice_bcc ####################

#################### libopen_dice_cbor_reader_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/external/open-dice/src/cbor_reader_fuzzer.cc
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
    add_executable(libopen_dice_cbor_reader_fuzzer ${SOURCE_FILES})

    target_include_directories(libopen_dice_cbor_reader_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/external/open-dice
            ${ANDROID_ROOT}/external/open-dice/include/dice/config/boringssl_ed25519
            ${ANDROID_ROOT}/external/open-dice/include
    )

    # Add c/c++ flags
    target_compile_options(libopen_dice_cbor_reader_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(libopen_dice_cbor_reader_fuzzer
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            libm
            libopen_dice_bcc
    )
endif ()
#################### libopen_dice_cbor_reader_fuzzer ####################

#################### libopen_dice_cbor_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/external/open-dice/src/fuzzer.cc
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
    add_executable(libopen_dice_cbor_fuzzer ${SOURCE_FILES})

    target_include_directories(libopen_dice_cbor_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/external/open-dice
            ${ANDROID_ROOT}/external/open-dice/include/dice/config/boringssl_ed25519
            ${ANDROID_ROOT}/external/open-dice/include
    )

    # Add c/c++ flags
    target_compile_options(libopen_dice_cbor_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(libopen_dice_cbor_fuzzer
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcrypto
            libdl
            libm
            libopen_dice_cbor
    )
endif ()
#################### libopen_dice_cbor_fuzzer ####################

#################### libopen_dice_cbor_writer_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/external/open-dice/src/cbor_writer_fuzzer.cc
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
    add_executable(libopen_dice_cbor_writer_fuzzer ${SOURCE_FILES})

    target_include_directories(libopen_dice_cbor_writer_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/external/open-dice
            ${ANDROID_ROOT}/external/open-dice/include/dice/config/boringssl_ed25519
            ${ANDROID_ROOT}/external/open-dice/include
    )

    # Add c/c++ flags
    target_compile_options(libopen_dice_cbor_writer_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(libopen_dice_cbor_writer_fuzzer
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            libm
            libopen_dice_cbor
    )
endif ()
#################### libopen_dice_cbor_writer_fuzzer ####################

#################### libopen_dice_cbor_reader_p384_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/external/open-dice/src/cbor_reader_fuzzer.cc
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
    add_executable(libopen_dice_cbor_reader_p384_fuzzer ${SOURCE_FILES})

    target_include_directories(libopen_dice_cbor_reader_p384_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/external/open-dice
            ${ANDROID_ROOT}/external/open-dice/include/dice/config/boringssl_ecdsa_p384
            ${ANDROID_ROOT}/external/open-dice/include
    )

    # Add c/c++ flags
    target_compile_options(libopen_dice_cbor_reader_p384_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(libopen_dice_cbor_reader_p384_fuzzer
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            libm
            libopen_dice_bcc_p384
    )
endif ()
#################### libopen_dice_cbor_reader_p384_fuzzer ####################

#################### libopen_dice_cbor_p384_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/external/open-dice/src/fuzzer.cc
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
    add_executable(libopen_dice_cbor_p384_fuzzer ${SOURCE_FILES})

    target_include_directories(libopen_dice_cbor_p384_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/external/open-dice
            ${ANDROID_ROOT}/external/open-dice/include/dice/config/boringssl_ecdsa_p384
            ${ANDROID_ROOT}/external/open-dice/include
    )

    # Add c/c++ flags
    target_compile_options(libopen_dice_cbor_p384_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(libopen_dice_cbor_p384_fuzzer
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcrypto
            libdl
            libm
            libopen_dice_cbor_p384
    )
endif ()
#################### libopen_dice_cbor_p384_fuzzer ####################

#################### libopen_dice_cbor_baremetal ####################
set(FILES
        ${ANDROID_ROOT}/external/open-dice/src/boringssl_hash_kdf_ops.c
        ${ANDROID_ROOT}/external/open-dice/src/boringssl_ed25519_ops.c
        ${ANDROID_ROOT}/external/open-dice/src/cbor_cert_op.c
        ${ANDROID_ROOT}/external/open-dice/src/cbor_ed25519_cert_op.c
        ${ANDROID_ROOT}/external/open-dice/src/cbor_writer.c
        ${ANDROID_ROOT}/external/open-dice/src/dice.c
        ${ANDROID_ROOT}/external/open-dice/src/utils.c
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
    add_library(libopen_dice_cbor_baremetal SHARED ${SOURCE_FILES})

    target_include_directories(libopen_dice_cbor_baremetal
            PUBLIC
            ${ANDROID_ROOT}/external/open-dice
            ${ANDROID_ROOT}/external/open-dice/include/dice/config/boringssl_ed25519
            ${ANDROID_ROOT}/external/open-dice/include
    )

    # Add c/c++ flags
    target_compile_options(libopen_dice_cbor_baremetal PRIVATE
            -Werror
    )

    target_link_libraries(libopen_dice_cbor_baremetal
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcrypto_baremetal
            libdl
            libm
            libopen_dice_boringssl_ed25519_headers
            libopen_dice_headers
    )
endif ()
#################### libopen_dice_cbor_baremetal ####################
