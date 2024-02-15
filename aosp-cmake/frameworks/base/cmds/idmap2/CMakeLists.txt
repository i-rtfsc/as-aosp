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

project(idmap2d)

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

#################### libidmap2 ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/base/cmds/idmap2/libidmap2/BinaryStreamVisitor.cpp
        ${ANDROID_ROOT}/frameworks/base/cmds/idmap2/libidmap2/CommandLineOptions.cpp
        ${ANDROID_ROOT}/frameworks/base/cmds/idmap2/libidmap2/FabricatedOverlay.cpp
        ${ANDROID_ROOT}/frameworks/base/cmds/idmap2/libidmap2/FileUtils.cpp
        ${ANDROID_ROOT}/frameworks/base/cmds/idmap2/libidmap2/Idmap.cpp
        ${ANDROID_ROOT}/frameworks/base/cmds/idmap2/libidmap2/PolicyUtils.cpp
        ${ANDROID_ROOT}/frameworks/base/cmds/idmap2/libidmap2/PrettyPrintVisitor.cpp
        ${ANDROID_ROOT}/frameworks/base/cmds/idmap2/libidmap2/RawPrintVisitor.cpp
        ${ANDROID_ROOT}/frameworks/base/cmds/idmap2/libidmap2/ResourceContainer.cpp
        ${ANDROID_ROOT}/frameworks/base/cmds/idmap2/libidmap2/ResourceMapping.cpp
        ${ANDROID_ROOT}/frameworks/base/cmds/idmap2/libidmap2/ResourceUtils.cpp
        ${ANDROID_ROOT}/frameworks/base/cmds/idmap2/libidmap2/Result.cpp
        ${ANDROID_ROOT}/frameworks/base/cmds/idmap2/libidmap2/SysTrace.cpp
        ${ANDROID_ROOT}/frameworks/base/cmds/idmap2/libidmap2/XmlParser.cpp
        ${ANDROID_ROOT}/frameworks/base/cmds/idmap2/self_targeting/SelfTargeting.cpp
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
    add_library(libidmap2 SHARED ${SOURCE_FILES})

    target_include_directories(libidmap2
            PUBLIC
            ${ANDROID_ROOT}/frameworks/base/cmds/idmap2/include
            ${ANDROID_ROOT}/frameworks/base/cmds/idmap2
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/cmds/idmap2/libidmap2_protos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/cmds/idmap2
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/cmds/idmap2/libidmap2_protos/android_${OUT_ARCH_CPU}_static/gen/proto
            ${ANDROID_ROOT}/frameworks/base/cmds/idmap2/libidmap2_policies/include
    )

    # Add c/c++ flags
    target_compile_options(libidmap2 PRIVATE
            -Werror
    )

    target_link_libraries(libidmap2
            libandroidfw
            libbase
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libidmap2_policies
            libidmap2_protos
            libm
            libprotobuf-cpp-lite
            libutils
            libz
            libziparchive
    )
endif ()
#################### libidmap2 ####################

#################### overlayable_policy_aidl-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/cmds/idmap2/overlayable_policy_aidl-ndk_platform-source/gen/android/os/OverlayablePolicy.cpp
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
    add_library(overlayable_policy_aidl-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(overlayable_policy_aidl-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/cmds/idmap2/overlayable_policy_aidl-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(overlayable_policy_aidl-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(overlayable_policy_aidl-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### overlayable_policy_aidl-ndk_platform ####################

#################### overlayable_policy_aidl-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/cmds/idmap2/overlayable_policy_aidl-ndk-source/gen/android/os/OverlayablePolicy.cpp
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
    add_library(overlayable_policy_aidl-ndk SHARED ${SOURCE_FILES})

    target_include_directories(overlayable_policy_aidl-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/cmds/idmap2/overlayable_policy_aidl-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(overlayable_policy_aidl-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(overlayable_policy_aidl-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### overlayable_policy_aidl-ndk ####################

#################### overlayable_policy_aidl-cpp-analyzer ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/cmds/idmap2/overlayable_policy_aidl-cpp-analyzer-source/gen/android/os/OverlayablePolicy.cpp
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
    add_executable(overlayable_policy_aidl-cpp-analyzer ${SOURCE_FILES})

    target_include_directories(overlayable_policy_aidl-cpp-analyzer
            PUBLIC
            ${ANDROID_ROOT}/frameworks/base/cmds/idmap2
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/cmds/idmap2/overlayable_policy_aidl-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/cmds/idmap2/overlayable_policy_aidl-cpp-analyzer-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(overlayable_policy_aidl-cpp-analyzer PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(overlayable_policy_aidl-cpp-analyzer
            aidl-analyzer-main
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libutils
            overlayable_policy_aidl-cpp
    )
endif ()
#################### overlayable_policy_aidl-cpp-analyzer ####################

#################### idmap2 ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/base/cmds/idmap2/idmap2/CommandUtils.cpp
        ${ANDROID_ROOT}/frameworks/base/cmds/idmap2/idmap2/Create.cpp
        ${ANDROID_ROOT}/frameworks/base/cmds/idmap2/idmap2/CreateMultiple.cpp
        ${ANDROID_ROOT}/frameworks/base/cmds/idmap2/idmap2/Dump.cpp
        ${ANDROID_ROOT}/frameworks/base/cmds/idmap2/idmap2/Lookup.cpp
        ${ANDROID_ROOT}/frameworks/base/cmds/idmap2/idmap2/Main.cpp
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
    add_executable(idmap2 ${SOURCE_FILES})

    target_include_directories(idmap2
            PUBLIC
            ${ANDROID_ROOT}/frameworks/base/cmds/idmap2
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/cmds/idmap2/libidmap2_protos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/cmds/idmap2
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/cmds/idmap2/libidmap2_protos/android_${OUT_ARCH_CPU}_static/gen/proto
            ${ANDROID_ROOT}/frameworks/base/cmds/idmap2/include
            ${ANDROID_ROOT}/frameworks/base/cmds/idmap2/libidmap2_policies/include
    )

    # Add c/c++ flags
    target_compile_options(idmap2 PRIVATE
            -Werror
    )

    target_link_libraries(idmap2
            libandroidfw
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libidmap2
            libidmap2_policies
            libidmap2_protos
            liblog
            libm
            libprotobuf-cpp-lite
            libutils
            libz
            libziparchive
    )
endif ()
#################### idmap2 ####################

#################### libidmap2_protos ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/cmds/idmap2/libidmap2_protos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/cmds/idmap2/libidmap2/proto/fabricated_v1.pb.cc
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
    add_library(libidmap2_protos SHARED ${SOURCE_FILES})

    target_include_directories(libidmap2_protos
            PUBLIC
            ${ANDROID_ROOT}/frameworks/base/cmds/idmap2
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/cmds/idmap2/libidmap2_protos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/cmds/idmap2
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/cmds/idmap2/libidmap2_protos/android_${OUT_ARCH_CPU}_static/gen/proto
    )

    # Add c/c++ flags
    target_compile_options(libidmap2_protos PRIVATE
            -Werror
    )

    target_link_libraries(libidmap2_protos
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libprotobuf-cpp-lite
    )
endif ()
#################### libidmap2_protos ####################

#################### idmap2d ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/base/cmds/idmap2/idmap2d/Idmap2Service.cpp
        ${ANDROID_ROOT}/frameworks/base/cmds/idmap2/idmap2d/Main.cpp
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
    add_executable(idmap2d ${SOURCE_FILES})

    target_include_directories(idmap2d
            PUBLIC
            ${ANDROID_ROOT}/frameworks/base/cmds/idmap2
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/cmds/idmap2/libidmap2_protos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/cmds/idmap2
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/cmds/idmap2/libidmap2_protos/android_${OUT_ARCH_CPU}_static/gen/proto
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/cmds/idmap2/libidmap2daidl/android_${OUT_ARCH_CPU}_static/gen/aidl
            ${ANDROID_ROOT}/frameworks/base/cmds/idmap2/include
            ${ANDROID_ROOT}/frameworks/base/cmds/idmap2/libidmap2_policies/include
    )

    # Add c/c++ flags
    target_compile_options(idmap2d PRIVATE
            -Werror
    )

    target_link_libraries(idmap2d
            libandroidfw
            libbase
            libbinder
            libc
            libc++
            libc++demangle
            libc++fs
            libclang_rt.builtins
            libcutils
            libdl
            libidmap2
            libidmap2_policies
            libidmap2_protos
            libidmap2daidl
            libm
            libprotobuf-cpp-lite
            libutils
            libziparchive
    )
endif ()
#################### idmap2d ####################

#################### overlayable_policy_aidl-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/cmds/idmap2/overlayable_policy_aidl-cpp-source/gen/android/os/OverlayablePolicy.cpp
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
    add_library(overlayable_policy_aidl-cpp SHARED ${SOURCE_FILES})

    target_include_directories(overlayable_policy_aidl-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/cmds/idmap2/overlayable_policy_aidl-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(overlayable_policy_aidl-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(overlayable_policy_aidl-cpp
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
#################### overlayable_policy_aidl-cpp ####################

#################### libidmap2daidl ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/cmds/idmap2/libidmap2daidl/android_${OUT_ARCH_CPU}_static/gen/aidl/android/os/IIdmap2.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/cmds/idmap2/libidmap2daidl/android_${OUT_ARCH_CPU}_static/gen/aidl/android/os/FabricatedOverlayInternal.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/cmds/idmap2/libidmap2daidl/android_${OUT_ARCH_CPU}_static/gen/aidl/android/os/FabricatedOverlayInternalEntry.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/cmds/idmap2/libidmap2daidl/android_${OUT_ARCH_CPU}_static/gen/aidl/android/os/FabricatedOverlayInfo.cpp
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
    add_library(libidmap2daidl SHARED ${SOURCE_FILES})

    target_include_directories(libidmap2daidl
            PUBLIC
            ${ANDROID_ROOT}/frameworks/base/cmds/idmap2
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/cmds/idmap2/libidmap2daidl/android_${OUT_ARCH_CPU}_static/gen/aidl
    )

    # Add c/c++ flags
    target_compile_options(libidmap2daidl PRIVATE
            -Werror
    )

    target_link_libraries(libidmap2daidl
            libbase
            libbinder_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libidmap2daidl ####################
