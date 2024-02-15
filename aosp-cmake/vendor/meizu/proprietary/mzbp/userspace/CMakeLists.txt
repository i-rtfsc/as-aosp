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

project(mzbp_helper)

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

#################### mzbp_server ####################
set(FILES
        ${ANDROID_ROOT}/vendor/meizu/proprietary/mzbp/userspace/mzbp_logdir_mgr.c
        ${ANDROID_ROOT}/vendor/meizu/proprietary/mzbp/userspace/mzbp_plog_mgr.c
        ${ANDROID_ROOT}/vendor/meizu/proprietary/mzbp/userspace/mzbp_trigger_api.c
        ${ANDROID_ROOT}/vendor/meizu/proprietary/mzbp/userspace/mzbp_logger_api.c
        ${ANDROID_ROOT}/vendor/meizu/proprietary/mzbp/userspace/mzbp_server.c
        ${ANDROID_ROOT}/vendor/meizu/proprietary/mzbp/userspace/startup/mzbp_binary_format_mgr_startup.c
        ${ANDROID_ROOT}/vendor/meizu/proprietary/mzbp/userspace/startup/mzbp_logdir_manager_startup.c
        ${ANDROID_ROOT}/vendor/meizu/proprietary/mzbp/userspace/startup/mzbp_logger_api_startup.c
        ${ANDROID_ROOT}/vendor/meizu/proprietary/mzbp/userspace/startup/mzbp_logger_mgr_startup.c
        ${ANDROID_ROOT}/vendor/meizu/proprietary/mzbp/userspace/startup/mzbp_trigger_api_startup.c
        ${ANDROID_ROOT}/vendor/meizu/proprietary/mzbp/userspace/startup/mzbp_trigger_mgr_startup.c
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
    add_executable(mzbp_server ${SOURCE_FILES})

    target_include_directories(mzbp_server
            PUBLIC
            ${ANDROID_ROOT}/vendor/meizu/proprietary/mzbp/userspace
            ${ANDROID_ROOT}/vendor/meizu/proprietary/mzbp/userspace/include
    )

    target_link_libraries(mzbp_server
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libdl
            libhidlbase
            liblog
            libm
            libmzbp_comm
            libmzbp_kernel
            libmzbp_userdebug
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### mzbp_server ####################

#################### mzbp_kernel_helper ####################
set(FILES
        ${ANDROID_ROOT}/vendor/meizu/proprietary/mzbp/userspace/mzbp_logdir_mgr.c
        ${ANDROID_ROOT}/vendor/meizu/proprietary/mzbp/userspace/mzbp_plog_mgr.c
        ${ANDROID_ROOT}/vendor/meizu/proprietary/mzbp/userspace/mzbp_trigger_api.c
        ${ANDROID_ROOT}/vendor/meizu/proprietary/mzbp/userspace/mzbp_logger_api.c
        ${ANDROID_ROOT}/vendor/meizu/proprietary/mzbp/userspace/mzbp_kernel_helper.c
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
    add_executable(mzbp_kernel_helper ${SOURCE_FILES})

    target_include_directories(mzbp_kernel_helper
            PUBLIC
            ${ANDROID_ROOT}/vendor/meizu/proprietary/mzbp/userspace
            ${ANDROID_ROOT}/vendor/meizu/proprietary/mzbp/userspace/include
    )

    target_link_libraries(mzbp_kernel_helper
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libdl
            libhidlbase
            liblog
            libm
            libmzbp_comm
            libmzbp_kernel
            libmzbp_userdebug
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### mzbp_kernel_helper ####################

#################### mzbp_helper ####################
set(FILES
        ${ANDROID_ROOT}/vendor/meizu/proprietary/mzbp/userspace/mzbp_logdir_mgr.c
        ${ANDROID_ROOT}/vendor/meizu/proprietary/mzbp/userspace/mzbp_plog_mgr.c
        ${ANDROID_ROOT}/vendor/meizu/proprietary/mzbp/userspace/mzbp_trigger_api.c
        ${ANDROID_ROOT}/vendor/meizu/proprietary/mzbp/userspace/mzbp_logger_api.c
        ${ANDROID_ROOT}/vendor/meizu/proprietary/mzbp/userspace/mzbp_object_selecter.c
        ${ANDROID_ROOT}/vendor/meizu/proprietary/mzbp/userspace/mzbp_helper.c
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
    add_executable(mzbp_helper ${SOURCE_FILES})

    target_include_directories(mzbp_helper
            PUBLIC
            ${ANDROID_ROOT}/vendor/meizu/proprietary/mzbp/userspace
            ${ANDROID_ROOT}/vendor/meizu/proprietary/mzbp/userspace/include
    )

    target_link_libraries(mzbp_helper
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libdl
            libhidlbase
            liblog
            libm
            libmzbp_comm
            libmzbp_kernel
            libmzbp_userdebug
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### mzbp_helper ####################
