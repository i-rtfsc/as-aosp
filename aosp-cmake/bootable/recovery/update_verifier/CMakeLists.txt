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

project(update_verifier)

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

#################### libupdate_verifier ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/bootable/recovery/update_verifier/libupdate_verifier/android_${OUT_ARCH_CPU}_static/gen/proto/bootable/recovery/update_verifier/care_map.pb.cc
        ${ANDROID_ROOT}/bootable/recovery/update_verifier/update_verifier.cpp
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
    add_library(libupdate_verifier SHARED ${SOURCE_FILES})

    target_include_directories(libupdate_verifier
            PUBLIC
            ${ANDROID_ROOT}/bootable/recovery/update_verifier/include
            ${ANDROID_ROOT}/bootable/recovery/update_verifier
            ${OUT_INTERMEDIATES_ROOT}/bootable/recovery/update_verifier/libupdate_verifier/android_${OUT_ARCH_CPU}_static/gen/proto/bootable/recovery/update_verifier
            ${OUT_INTERMEDIATES_ROOT}/bootable/recovery/update_verifier/libupdate_verifier/android_${OUT_ARCH_CPU}_static/gen/proto
    )

    # Add c/c++ flags
    target_compile_options(libupdate_verifier PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libupdate_verifier
            android.hardware.boot-1.0
            libbase
            libbinder
            libboot_control_client
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libm
            libotautil
            libprotobuf-cpp-lite
            libsnapshot_snapuserd
            libutils
            libvold_binder
    )
endif ()
#################### libupdate_verifier ####################

#################### update_verifier ####################
set(FILES
        ${ANDROID_ROOT}/bootable/recovery/update_verifier/update_verifier_main.cpp
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
    add_executable(update_verifier ${SOURCE_FILES})

    target_include_directories(update_verifier
            PUBLIC
            ${ANDROID_ROOT}/bootable/recovery/update_verifier/include
            ${ANDROID_ROOT}/bootable/recovery/update_verifier
            ${OUT_INTERMEDIATES_ROOT}/bootable/recovery/update_verifier/libupdate_verifier/android_${OUT_ARCH_CPU}_static/gen/proto/bootable/recovery/update_verifier
            ${OUT_INTERMEDIATES_ROOT}/bootable/recovery/update_verifier/libupdate_verifier/android_${OUT_ARCH_CPU}_static/gen/proto
    )

    # Add c/c++ flags
    target_compile_options(update_verifier PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(update_verifier
            android.hardware.boot-1.0
            libbase
            libbinder
            libboot_control_client
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libhardware
            libhidlbase
            liblog
            libm
            libotautil
            libprotobuf-cpp-lite
            libupdate_verifier
            libutils
            libvold_binder
    )
endif ()
#################### update_verifier ####################
