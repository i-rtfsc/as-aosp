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

project(assemble_cvd)

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

#################### assemble_cvd ####################
set(FILES
        ${ANDROID_ROOT}/device/google/cuttlefish/host/commands/assemble_cvd/alloc.cc
        ${ANDROID_ROOT}/device/google/cuttlefish/host/commands/assemble_cvd/assemble_cvd.cc
        ${ANDROID_ROOT}/device/google/cuttlefish/host/commands/assemble_cvd/boot_config.cc
        ${ANDROID_ROOT}/device/google/cuttlefish/host/commands/assemble_cvd/boot_image_utils.cc
        ${ANDROID_ROOT}/device/google/cuttlefish/host/commands/assemble_cvd/clean.cc
        ${ANDROID_ROOT}/device/google/cuttlefish/host/commands/assemble_cvd/disk_builder.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/host/commands/assemble_cvd/disk_flags.cc
        ${ANDROID_ROOT}/device/google/cuttlefish/host/commands/assemble_cvd/flags.cc
        ${ANDROID_ROOT}/device/google/cuttlefish/host/commands/assemble_cvd/flag_feature.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/host/commands/assemble_cvd/misc_info.cc
        ${ANDROID_ROOT}/device/google/cuttlefish/host/commands/assemble_cvd/super_image_mixer.cc
        ${ANDROID_ROOT}/device/google/cuttlefish/host/commands/assemble_cvd/vendor_dlkm_utils.cc
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
    add_executable(assemble_cvd ${SOURCE_FILES})

    target_include_directories(assemble_cvd
            PUBLIC
            ${ANDROID_ROOT}/device/google/cuttlefish/host/commands/assemble_cvd
            ${OUT_INTERMEDIATES_ROOT}/device/google/cuttlefish/host/commands/assemble_cvd/proto/libcuttlefish_launch_cvd_proto/android_${OUT_ARCH_CPU}_static/gen/proto/device/google/cuttlefish/host/commands/assemble_cvd/proto
            ${OUT_INTERMEDIATES_ROOT}/device/google/cuttlefish/host/commands/assemble_cvd/proto/libcuttlefish_launch_cvd_proto/android_${OUT_ARCH_CPU}_static/gen/proto
    )

    # Add c/c++ flags
    target_compile_options(assemble_cvd PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(assemble_cvd
            bootimg_headers
            cuttlefish_common_headers
            libbase
            libc
            libc++
            libc++demangle
            libcdisk_spec
            libclang_rt.builtins
            libcuttlefish_allocd_utils
            libcuttlefish_display_flags
            libcuttlefish_fs
            libcuttlefish_graphics_configuration
            libcuttlefish_graphics_detector
            libcuttlefish_host_config
            libcuttlefish_host_config_adb
            libcuttlefish_host_config_fastboot
            libcuttlefish_launch_cvd_proto
            libcuttlefish_utils
            libcuttlefish_vm_manager
            libdl
            libext2_blkid
            libext2_uuid
            libfruit
            libgflags
            libicuuc
            libimage_aggregator
            libjsoncpp
            libm
            libnl
            libprotobuf-cpp-full
            libsparse
            libz
            libziparchive
    )
endif ()
#################### assemble_cvd ####################

#################### libcuttlefish_display_flags ####################
set(FILES
        ${ANDROID_ROOT}/device/google/cuttlefish/host/commands/assemble_cvd/display_flags.cpp
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
    add_library(libcuttlefish_display_flags SHARED ${SOURCE_FILES})

    target_include_directories(libcuttlefish_display_flags
            PUBLIC
            ${ANDROID_ROOT}/device/google/cuttlefish/host/commands/assemble_cvd
    )

    # Add c/c++ flags
    target_compile_options(libcuttlefish_display_flags PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(libcuttlefish_display_flags
            cuttlefish_common_headers
            libbase
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcuttlefish_host_config
            libdl
            libjsoncpp
            libm
    )
endif ()
#################### libcuttlefish_display_flags ####################
