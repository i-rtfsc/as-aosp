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

project(fbase_defrag)

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

#################### fbase_defrag ####################
set(FILES
        ${ANDROID_ROOT}/vendor/meizu/proprietary/fbase_defrag/fbase_defrag/main.c
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
    add_executable(fbase_defrag ${SOURCE_FILES})

    target_include_directories(fbase_defrag
            PUBLIC
            ${ANDROID_ROOT}/vendor/meizu/proprietary/fbase_defrag/fbase_defrag
    )

    target_link_libraries(fbase_defrag
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libdl
            libfbase_defrag
            liblog
            libm
            libc
            libdl
            libm
    )
endif ()
#################### fbase_defrag ####################

#################### libfbase_defrag ####################
set(FILES
        ${ANDROID_ROOT}/vendor/meizu/proprietary/fbase_defrag/fbase_defrag/file_map.c
        ${ANDROID_ROOT}/vendor/meizu/proprietary/fbase_defrag/fbase_defrag/ufs_fbo.c
        ${ANDROID_ROOT}/vendor/meizu/proprietary/fbase_defrag/fbase_defrag/ufs_fbo_controller.c
        ${ANDROID_ROOT}/vendor/meizu/proprietary/fbase_defrag/fbase_defrag/ufs_cmds.c
        ${ANDROID_ROOT}/vendor/meizu/proprietary/fbase_defrag/fbase_defrag/scsi_bsg_util.c
        ${ANDROID_ROOT}/vendor/meizu/proprietary/fbase_defrag/fbase_defrag/ufs.c
        ${ANDROID_ROOT}/vendor/meizu/proprietary/fbase_defrag/fbase_defrag/dump.c
        ${ANDROID_ROOT}/vendor/meizu/proprietary/fbase_defrag/fbase_defrag/vfs_stat.c
        ${ANDROID_ROOT}/vendor/meizu/proprietary/fbase_defrag/fbase_defrag/hashmap.c
        ${ANDROID_ROOT}/vendor/meizu/proprietary/fbase_defrag/fbase_defrag/crc32.c
        ${ANDROID_ROOT}/vendor/meizu/proprietary/fbase_defrag/fbase_defrag/vfs_stat_mgr.c
        ${ANDROID_ROOT}/vendor/meizu/proprietary/fbase_defrag/fbase_defrag/fbo_perf_mgr.c
        ${ANDROID_ROOT}/vendor/meizu/proprietary/fbase_defrag/fbase_defrag/dir_iter.c
        ${ANDROID_ROOT}/vendor/meizu/proprietary/fbase_defrag/fbase_defrag/regex_dir_iter.c
        ${ANDROID_ROOT}/vendor/meizu/proprietary/fbase_defrag/fbase_defrag/fbase_defrag_thread.c
        ${ANDROID_ROOT}/vendor/meizu/proprietary/fbase_defrag/fbase_defrag/scan_config.c
        ${ANDROID_ROOT}/vendor/meizu/proprietary/fbase_defrag/fbase_defrag/ufs_hardware.c
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
    add_library(libfbase_defrag SHARED ${SOURCE_FILES})

    target_include_directories(libfbase_defrag
            PUBLIC
            ${ANDROID_ROOT}/vendor/meizu/proprietary/fbase_defrag/fbase_defrag
    )

    target_link_libraries(libfbase_defrag
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libdl
            liblog
            libm
            libc
            libdl
            libm
    )
endif ()
#################### libfbase_defrag ####################
