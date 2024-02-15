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

project(rkp_factory_extraction_tool)

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

#################### rkp_factory_extraction_tool ####################
set(FILES
        ${ANDROID_ROOT}/system/security/provisioner/rkp_factory_extraction_tool.cpp
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
    add_executable(rkp_factory_extraction_tool ${SOURCE_FILES})

    target_include_directories(rkp_factory_extraction_tool
            PUBLIC
            ${ANDROID_ROOT}/system/security/provisioner
    )

    # Add c/c++ flags
    target_compile_options(rkp_factory_extraction_tool PRIVATE
            -Werror
    )

    target_link_libraries(rkp_factory_extraction_tool
            android.hardware.security.keymint-V3-ndk
            android.hardware.security.rkp-V3-ndk
            libbase
            libbinder
            libbinder_ndk
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcppbor_external
            libcppcose_rkp
            libcrypto
            libdl
            libgflags
            libjsoncpp
            libkeymint_remote_prov_support
            liblog
            libm
            librkp_factory_extraction
            libc
            libdl
            libm
    )
endif ()
#################### rkp_factory_extraction_tool ####################

#################### librkp_factory_extraction ####################
set(FILES
        ${ANDROID_ROOT}/system/security/provisioner/rkp_factory_extraction_lib.cpp
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
    add_library(librkp_factory_extraction SHARED ${SOURCE_FILES})

    target_include_directories(librkp_factory_extraction
            PUBLIC
            ${ANDROID_ROOT}/system/security/provisioner
    )

    # Add c/c++ flags
    target_compile_options(librkp_factory_extraction PRIVATE
            -Werror
    )

    target_link_libraries(librkp_factory_extraction
            android.hardware.security.keymint-V3-ndk
            android.hardware.security.rkp-V3-ndk
            libbase
            libbinder
            libbinder_ndk
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcppbor_external
            libcppcose_rkp
            libcrypto
            libdl
            libjsoncpp
            libkeymint_remote_prov_support
            liblog
            libm
    )
endif ()
#################### librkp_factory_extraction ####################
