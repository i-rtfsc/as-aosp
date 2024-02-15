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

project(lpdumpd)

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

#################### lpdump ####################
set(FILES
        ${ANDROID_ROOT}/system/extras/partition_tools/lpdump_target.cc
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
    add_executable(lpdump ${SOURCE_FILES})

    target_include_directories(lpdump
            PUBLIC
            ${ANDROID_ROOT}/system/extras/partition_tools
            ${OUT_INTERMEDIATES_ROOT}/system/extras/partition_tools/aidl/liblpdump_interface-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(lpdump PRIVATE
            -Werror
            -Wextra
    )

    target_link_libraries(lpdump
            libbase
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libjsonpbparse
            liblog
            liblp
            liblpdump
            liblpdump_interface-cpp
            libm
            libutils
    )
endif ()
#################### lpdump ####################

#################### liblpdump ####################
set(FILES
        ${ANDROID_ROOT}/system/extras/partition_tools/lpdump.cc
        ${OUT_INTERMEDIATES_ROOT}/system/extras/partition_tools/liblpdump/android_${OUT_ARCH_CPU}_shared/gen/proto/system/extras/partition_tools/dynamic_partitions_device_info.pb.cc
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
    add_library(liblpdump SHARED ${SOURCE_FILES})

    target_include_directories(liblpdump
            PUBLIC
            ${ANDROID_ROOT}/system/extras/partition_tools
            ${OUT_INTERMEDIATES_ROOT}/system/extras/partition_tools/liblpdump/android_${OUT_ARCH_CPU}_shared/gen/proto/system/extras/partition_tools
            ${OUT_INTERMEDIATES_ROOT}/system/extras/partition_tools/liblpdump/android_${OUT_ARCH_CPU}_shared/gen/proto
    )

    # Add c/c++ flags
    target_compile_options(liblpdump PRIVATE
            -Werror
            -Wextra
    )

    target_link_libraries(liblpdump
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libfs_mgr
            libjsonpbparse
            liblog
            liblp
            libm
            libprotobuf-cpp-full
    )
endif ()
#################### liblpdump ####################

#################### lpflash ####################
set(FILES
        ${ANDROID_ROOT}/system/extras/partition_tools/lpflash.cc
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
    add_executable(lpflash ${SOURCE_FILES})

    target_include_directories(lpflash
            PUBLIC
            ${ANDROID_ROOT}/system/extras/partition_tools
    )

    # Add c/c++ flags
    target_compile_options(lpflash PRIVATE
            -Werror
            -Wextra
    )

    target_link_libraries(lpflash
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            liblog
            liblp
            libm
    )
endif ()
#################### lpflash ####################

#################### lpdumpd ####################
set(FILES
        ${ANDROID_ROOT}/system/extras/partition_tools/lpdumpd.cc
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
    add_executable(lpdumpd ${SOURCE_FILES})

    target_include_directories(lpdumpd
            PUBLIC
            ${ANDROID_ROOT}/system/extras/partition_tools
            ${OUT_INTERMEDIATES_ROOT}/system/extras/partition_tools/aidl/liblpdump_interface-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(lpdumpd PRIVATE
            -Werror
            -Wextra
    )

    target_link_libraries(lpdumpd
            libbase
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            liblog
            liblp
            liblpdump
            liblpdump_interface-cpp
            libm
            libutils
    )
endif ()
#################### lpdumpd ####################
