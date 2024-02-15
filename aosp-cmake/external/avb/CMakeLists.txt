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

project(avbctl)

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

#################### libavb_user ####################
set(FILES
        ${ANDROID_ROOT}/external/avb/libavb/boringssl/sha.c
        ${ANDROID_ROOT}/external/avb/libavb/avb_chain_partition_descriptor.c
        ${ANDROID_ROOT}/external/avb/libavb/avb_cmdline.c
        ${ANDROID_ROOT}/external/avb/libavb/avb_crc32.c
        ${ANDROID_ROOT}/external/avb/libavb/avb_crypto.c
        ${ANDROID_ROOT}/external/avb/libavb/avb_descriptor.c
        ${ANDROID_ROOT}/external/avb/libavb/avb_footer.c
        ${ANDROID_ROOT}/external/avb/libavb/avb_hash_descriptor.c
        ${ANDROID_ROOT}/external/avb/libavb/avb_hashtree_descriptor.c
        ${ANDROID_ROOT}/external/avb/libavb/avb_kernel_cmdline_descriptor.c
        ${ANDROID_ROOT}/external/avb/libavb/avb_property_descriptor.c
        ${ANDROID_ROOT}/external/avb/libavb/avb_rsa.c
        ${ANDROID_ROOT}/external/avb/libavb/avb_slot_verify.c
        ${ANDROID_ROOT}/external/avb/libavb/avb_util.c
        ${ANDROID_ROOT}/external/avb/libavb/avb_vbmeta_image.c
        ${ANDROID_ROOT}/external/avb/libavb/avb_version.c
        ${ANDROID_ROOT}/external/avb/libavb/avb_sysdeps_posix.c
        ${ANDROID_ROOT}/external/avb/libavb_ab/avb_ab_flow.c
        ${ANDROID_ROOT}/external/avb/libavb_user/avb_ops_user.cpp
        ${ANDROID_ROOT}/external/avb/libavb_user/avb_user_verity.c
        ${ANDROID_ROOT}/external/avb/libavb_user/avb_user_verification.c
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
    add_library(libavb_user SHARED ${SOURCE_FILES})

    target_include_directories(libavb_user
            PUBLIC
            ${ANDROID_ROOT}/external/avb/libavb/boringssl
            ${ANDROID_ROOT}/external/avb
    )

    # Add c/c++ flags
    target_compile_options(libavb_user PRIVATE
            -Werror
            -Wall
            -Wextra
            -Wno-unused-parameter
            -DAVB_ENABLE_DEBUG
    )

    target_link_libraries(libavb_user
            avb_headers
            libbase
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto
            libdl
            libfs_mgr
            libm
    )
endif ()
#################### libavb_user ####################

#################### bootctrl.avb ####################
set(FILES
        ${ANDROID_ROOT}/external/avb/boot_control/boot_control_avb.c
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
    add_library(bootctrl.avb SHARED ${SOURCE_FILES})

    target_include_directories(bootctrl.avb
            PUBLIC
            ${ANDROID_ROOT}/external/avb
    )

    # Add c/c++ flags
    target_compile_options(bootctrl.avb PRIVATE
            -Werror
            -Wall
            -Wextra
            -Wno-unused-parameter
            -DAVB_ENABLE_DEBUG
    )

    target_link_libraries(bootctrl.avb
            libavb_user
            libbase
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto
            libcutils
            libdl
            libfs_mgr
            libm
    )
endif ()
#################### bootctrl.avb ####################

#################### libavb ####################
set(FILES
        ${ANDROID_ROOT}/external/avb/libavb/boringssl/sha.c
        ${ANDROID_ROOT}/external/avb/libavb/avb_chain_partition_descriptor.c
        ${ANDROID_ROOT}/external/avb/libavb/avb_cmdline.c
        ${ANDROID_ROOT}/external/avb/libavb/avb_crc32.c
        ${ANDROID_ROOT}/external/avb/libavb/avb_crypto.c
        ${ANDROID_ROOT}/external/avb/libavb/avb_descriptor.c
        ${ANDROID_ROOT}/external/avb/libavb/avb_footer.c
        ${ANDROID_ROOT}/external/avb/libavb/avb_hash_descriptor.c
        ${ANDROID_ROOT}/external/avb/libavb/avb_hashtree_descriptor.c
        ${ANDROID_ROOT}/external/avb/libavb/avb_kernel_cmdline_descriptor.c
        ${ANDROID_ROOT}/external/avb/libavb/avb_property_descriptor.c
        ${ANDROID_ROOT}/external/avb/libavb/avb_rsa.c
        ${ANDROID_ROOT}/external/avb/libavb/avb_slot_verify.c
        ${ANDROID_ROOT}/external/avb/libavb/avb_util.c
        ${ANDROID_ROOT}/external/avb/libavb/avb_vbmeta_image.c
        ${ANDROID_ROOT}/external/avb/libavb/avb_version.c
        ${ANDROID_ROOT}/external/avb/libavb/avb_sysdeps_posix.c
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
    add_library(libavb SHARED ${SOURCE_FILES})

    target_include_directories(libavb
            PUBLIC
            ${ANDROID_ROOT}/external/avb/libavb/boringssl
            ${ANDROID_ROOT}/external/avb
    )

    # Add c/c++ flags
    target_compile_options(libavb PRIVATE
            -Werror
            -Wall
            -Wextra
            -Wno-unused-parameter
            -DAVB_ENABLE_DEBUG
    )

    target_link_libraries(libavb
            avb_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcrypto
            libdl
            libm
    )
endif ()
#################### libavb ####################

#################### libavb_baremetal ####################
set(FILES
        ${ANDROID_ROOT}/external/avb/libavb/boringssl/sha.c
        ${ANDROID_ROOT}/external/avb/libavb/avb_chain_partition_descriptor.c
        ${ANDROID_ROOT}/external/avb/libavb/avb_cmdline.c
        ${ANDROID_ROOT}/external/avb/libavb/avb_crc32.c
        ${ANDROID_ROOT}/external/avb/libavb/avb_crypto.c
        ${ANDROID_ROOT}/external/avb/libavb/avb_descriptor.c
        ${ANDROID_ROOT}/external/avb/libavb/avb_footer.c
        ${ANDROID_ROOT}/external/avb/libavb/avb_hash_descriptor.c
        ${ANDROID_ROOT}/external/avb/libavb/avb_hashtree_descriptor.c
        ${ANDROID_ROOT}/external/avb/libavb/avb_kernel_cmdline_descriptor.c
        ${ANDROID_ROOT}/external/avb/libavb/avb_property_descriptor.c
        ${ANDROID_ROOT}/external/avb/libavb/avb_rsa.c
        ${ANDROID_ROOT}/external/avb/libavb/avb_slot_verify.c
        ${ANDROID_ROOT}/external/avb/libavb/avb_util.c
        ${ANDROID_ROOT}/external/avb/libavb/avb_vbmeta_image.c
        ${ANDROID_ROOT}/external/avb/libavb/avb_version.c
        ${ANDROID_ROOT}/external/avb/libavb/avb_sysdeps_posix.c
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
    add_library(libavb_baremetal SHARED ${SOURCE_FILES})

    target_include_directories(libavb_baremetal
            PUBLIC
            ${ANDROID_ROOT}/external/avb/libavb/boringssl
            ${ANDROID_ROOT}/external/avb
    )

    # Add c/c++ flags
    target_compile_options(libavb_baremetal PRIVATE
            -Werror
            -Wall
            -Wextra
            -Wno-unused-parameter
            -DAVB_ENABLE_DEBUG
            -UAVB_ENABLE_DEBUG
    )

    target_link_libraries(libavb_baremetal
            avb_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcrypto_baremetal
            libdl
            libm
    )
endif ()
#################### libavb_baremetal ####################

#################### avbctl ####################
set(FILES
        ${ANDROID_ROOT}/external/avb/tools/avbctl/avbctl.cc
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
    add_executable(avbctl ${SOURCE_FILES})

    target_include_directories(avbctl
            PUBLIC
            ${ANDROID_ROOT}/external/avb
    )

    # Add c/c++ flags
    target_compile_options(avbctl PRIVATE
            -Werror
            -Wall
            -Wextra
            -Wno-unused-parameter
            -DAVB_ENABLE_DEBUG
    )

    target_link_libraries(avbctl
            libavb_user
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto
            libdl
            libfs_mgr
            libm
    )
endif ()
#################### avbctl ####################
