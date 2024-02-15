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

project(librecovery_utils)

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

#################### librecovery_utils ####################
set(FILES
        ${ANDROID_ROOT}/bootable/recovery/recovery_utils/battery_utils.cpp
        ${ANDROID_ROOT}/bootable/recovery/recovery_utils/logging.cpp
        ${ANDROID_ROOT}/bootable/recovery/recovery_utils/parse_install_logs.cpp
        ${ANDROID_ROOT}/bootable/recovery/recovery_utils/roots.cpp
        ${ANDROID_ROOT}/bootable/recovery/recovery_utils/thermalutil.cpp
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
    add_library(librecovery_utils SHARED ${SOURCE_FILES})

    target_include_directories(librecovery_utils
            PUBLIC
            ${ANDROID_ROOT}/bootable/recovery/recovery_utils/include
            ${ANDROID_ROOT}/bootable/recovery/recovery_utils
    )

    # Add c/c++ flags
    target_compile_options(librecovery_utils PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(librecovery_utils
            android.hardware.health-V2-ndk
            android.hardware.health-translate-ndk
            android.hardware.health-2.0
            libbase
            libbinder_ndk
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libext4_utils
            libfs_mgr
            libfstab
            libhealthhalutils
            libhealthshim
            libhidlbase
            libm
            libotautil
            libselinux
            libutils
            libvold_headers
    )
endif ()
#################### librecovery_utils ####################
