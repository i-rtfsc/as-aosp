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

project(libselinux)

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

#################### libselinux ####################
set(FILES
        ${ANDROID_ROOT}/external/selinux/libselinux/src/android/android.c
        ${ANDROID_ROOT}/external/selinux/libselinux/src/android/android_seapp.c
        ${ANDROID_ROOT}/external/selinux/libselinux/src/avc.c
        ${ANDROID_ROOT}/external/selinux/libselinux/src/avc_internal.c
        ${ANDROID_ROOT}/external/selinux/libselinux/src/avc_sidtab.c
        ${ANDROID_ROOT}/external/selinux/libselinux/src/booleans.c
        ${ANDROID_ROOT}/external/selinux/libselinux/src/callbacks.c
        ${ANDROID_ROOT}/external/selinux/libselinux/src/canonicalize_context.c
        ${ANDROID_ROOT}/external/selinux/libselinux/src/checkAccess.c
        ${ANDROID_ROOT}/external/selinux/libselinux/src/check_context.c
        ${ANDROID_ROOT}/external/selinux/libselinux/src/compute_av.c
        ${ANDROID_ROOT}/external/selinux/libselinux/src/compute_create.c
        ${ANDROID_ROOT}/external/selinux/libselinux/src/compute_member.c
        ${ANDROID_ROOT}/external/selinux/libselinux/src/context.c
        ${ANDROID_ROOT}/external/selinux/libselinux/src/deny_unknown.c
        ${ANDROID_ROOT}/external/selinux/libselinux/src/disable.c
        ${ANDROID_ROOT}/external/selinux/libselinux/src/enabled.c
        ${ANDROID_ROOT}/external/selinux/libselinux/src/fgetfilecon.c
        ${ANDROID_ROOT}/external/selinux/libselinux/src/freecon.c
        ${ANDROID_ROOT}/external/selinux/libselinux/src/fsetfilecon.c
        ${ANDROID_ROOT}/external/selinux/libselinux/src/get_initial_context.c
        ${ANDROID_ROOT}/external/selinux/libselinux/src/getenforce.c
        ${ANDROID_ROOT}/external/selinux/libselinux/src/getfilecon.c
        ${ANDROID_ROOT}/external/selinux/libselinux/src/getpeercon.c
        ${ANDROID_ROOT}/external/selinux/libselinux/src/init.c
        ${ANDROID_ROOT}/external/selinux/libselinux/src/label.c
        ${ANDROID_ROOT}/external/selinux/libselinux/src/label_backends_android.c
        ${ANDROID_ROOT}/external/selinux/libselinux/src/label_file.c
        ${ANDROID_ROOT}/external/selinux/libselinux/src/label_support.c
        ${ANDROID_ROOT}/external/selinux/libselinux/src/lgetfilecon.c
        ${ANDROID_ROOT}/external/selinux/libselinux/src/load_policy.c
        ${ANDROID_ROOT}/external/selinux/libselinux/src/lsetfilecon.c
        ${ANDROID_ROOT}/external/selinux/libselinux/src/mapping.c
        ${ANDROID_ROOT}/external/selinux/libselinux/src/matchpathcon.c
        ${ANDROID_ROOT}/external/selinux/libselinux/src/policyvers.c
        ${ANDROID_ROOT}/external/selinux/libselinux/src/procattr.c
        ${ANDROID_ROOT}/external/selinux/libselinux/src/regex.c
        ${ANDROID_ROOT}/external/selinux/libselinux/src/reject_unknown.c
        ${ANDROID_ROOT}/external/selinux/libselinux/src/selinux_internal.c
        ${ANDROID_ROOT}/external/selinux/libselinux/src/sestatus.c
        ${ANDROID_ROOT}/external/selinux/libselinux/src/setenforce.c
        ${ANDROID_ROOT}/external/selinux/libselinux/src/setfilecon.c
        ${ANDROID_ROOT}/external/selinux/libselinux/src/setrans_client.c
        ${ANDROID_ROOT}/external/selinux/libselinux/src/sha1.c
        ${ANDROID_ROOT}/external/selinux/libselinux/src/stringrep.c
        ${ANDROID_ROOT}/external/selinux/libselinux/src/android/android_device.c
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
    add_library(libselinux SHARED ${SOURCE_FILES})

    target_include_directories(libselinux
            PUBLIC
            ${ANDROID_ROOT}/external/selinux/libselinux/include
            ${ANDROID_ROOT}/external/selinux/libselinux/src
            ${ANDROID_ROOT}/external/selinux/libselinux
    )

    # Add c/c++ flags
    target_compile_options(libselinux PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libselinux
            libbase_headers
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils_headers
            liblog
            liblog_headers
            libpackagelistparser
            libpcre2
    )
endif ()
#################### libselinux ####################
