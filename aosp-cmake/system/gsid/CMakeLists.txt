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

project(gsi_aidl_interface-cpp-analyzer)

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

#################### gsid ####################
set(FILES
        ${ANDROID_ROOT}/system/gsid/daemon.cpp
        ${ANDROID_ROOT}/system/gsid/gsi_service.cpp
        ${ANDROID_ROOT}/system/gsid/partition_installer.cpp
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
    add_executable(gsid ${SOURCE_FILES})

    target_include_directories(gsid
            PUBLIC
            ${ANDROID_ROOT}/system/gsid/include
            ${ANDROID_ROOT}/system/gsid
            ${OUT_INTERMEDIATES_ROOT}/system/gsid/gsi_aidl_interface-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(gsid PRIVATE
            -Werror
    )

    target_link_libraries(gsid
            gsi_aidl_interface-cpp
            libavb
            libbase
            libbinder
            libc
            libc++
            libc++demangle
            libc++fs
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto
            libcutils
            libdl
            libdm
            libext4_utils
            libfs_mgr
            libgsi
            libgsid
            liblog
            liblp
            libm
            libprocessgroup
            libselinux
            libstorage_literals_headers
            libutils
            libvndksupport
            libvold_binder
    )
endif ()
#################### gsid ####################

#################### gsi_aidl_interface-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/gsid/gsi_aidl_interface-cpp-source/gen/android/gsi/AvbPublicKey.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/gsid/gsi_aidl_interface-cpp-source/gen/android/gsi/GsiProgress.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/gsid/gsi_aidl_interface-cpp-source/gen/android/gsi/IGsiService.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/gsid/gsi_aidl_interface-cpp-source/gen/android/gsi/IGsiServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/gsid/gsi_aidl_interface-cpp-source/gen/android/gsi/IImageService.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/gsid/gsi_aidl_interface-cpp-source/gen/android/gsi/IProgressCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/gsid/gsi_aidl_interface-cpp-source/gen/android/gsi/MappedImage.cpp
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
    add_library(gsi_aidl_interface-cpp SHARED ${SOURCE_FILES})

    target_include_directories(gsi_aidl_interface-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/system/gsid/gsi_aidl_interface-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(gsi_aidl_interface-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(gsi_aidl_interface-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libm
            libutils
    )
endif ()
#################### gsi_aidl_interface-cpp ####################

#################### libgsid ####################
set(FILES
        ${ANDROID_ROOT}/system/gsid/libgsid.cpp
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
    add_library(libgsid SHARED ${SOURCE_FILES})

    target_include_directories(libgsid
            PUBLIC
            ${ANDROID_ROOT}/system/gsid/include
            ${ANDROID_ROOT}/system/gsid
            ${OUT_INTERMEDIATES_ROOT}/system/gsid/gsi_aidl_interface-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(libgsid PRIVATE
            -Werror
    )

    target_link_libraries(libgsid
            gsi_aidl_interface-cpp
            libbase
            libbinder
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libutils
    )
endif ()
#################### libgsid ####################

#################### gsi_tool ####################
set(FILES
        ${ANDROID_ROOT}/system/gsid/gsi_tool.cpp
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
    add_executable(gsi_tool ${SOURCE_FILES})

    target_include_directories(gsi_tool
            PUBLIC
            ${ANDROID_ROOT}/system/gsid
            ${ANDROID_ROOT}/system/gsid/include
            ${OUT_INTERMEDIATES_ROOT}/system/gsid/gsi_aidl_interface-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(gsi_tool PRIVATE
            -Werror
    )

    target_link_libraries(gsi_tool
            gsi_aidl_interface-cpp
            libbase
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libgsi
            libgsid
            liblog
            libm
            libutils
    )
endif ()
#################### gsi_tool ####################

#################### gsi_aidl_interface-cpp-analyzer ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/gsid/gsi_aidl_interface-cpp-analyzer-source/gen/android/gsi/AvbPublicKey.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/gsid/gsi_aidl_interface-cpp-analyzer-source/gen/android/gsi/GsiProgress.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/gsid/gsi_aidl_interface-cpp-analyzer-source/gen/android/gsi/IGsiService.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/gsid/gsi_aidl_interface-cpp-analyzer-source/gen/android/gsi/IGsiServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/gsid/gsi_aidl_interface-cpp-analyzer-source/gen/android/gsi/IImageService.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/gsid/gsi_aidl_interface-cpp-analyzer-source/gen/android/gsi/IProgressCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/gsid/gsi_aidl_interface-cpp-analyzer-source/gen/android/gsi/MappedImage.cpp
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
    add_executable(gsi_aidl_interface-cpp-analyzer ${SOURCE_FILES})

    target_include_directories(gsi_aidl_interface-cpp-analyzer
            PUBLIC
            ${ANDROID_ROOT}/system/gsid
            ${OUT_INTERMEDIATES_ROOT}/system/gsid/gsi_aidl_interface-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/system/gsid/gsi_aidl_interface-cpp-analyzer-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(gsi_aidl_interface-cpp-analyzer PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(gsi_aidl_interface-cpp-analyzer
            aidl-analyzer-main
            gsi_aidl_interface-cpp
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libutils
    )
endif ()
#################### gsi_aidl_interface-cpp-analyzer ####################

#################### libgsi ####################
set(FILES
        ${ANDROID_ROOT}/system/gsid/libgsi.cpp
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
    add_library(libgsi SHARED ${SOURCE_FILES})

    target_include_directories(libgsi
            PUBLIC
            ${ANDROID_ROOT}/system/gsid/include
            ${ANDROID_ROOT}/system/gsid
    )

    # Add c/c++ flags
    target_compile_options(libgsi PRIVATE
            -Werror
    )

    target_link_libraries(libgsi
            libbase
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libgsi ####################
