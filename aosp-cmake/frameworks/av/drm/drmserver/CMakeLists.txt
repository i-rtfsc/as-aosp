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

project(drmserver_fuzzer)

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

#################### drmserver ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/drm/drmserver/main_drmserver.cpp
        ${ANDROID_ROOT}/frameworks/av/drm/drmserver/DrmManager.cpp
        ${ANDROID_ROOT}/frameworks/av/drm/drmserver/DrmManagerService.cpp
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
    add_executable(drmserver ${SOURCE_FILES})

    target_include_directories(drmserver
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/drm/drmserver
    )

    # Add c/c++ flags
    target_compile_options(drmserver PRIVATE
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(drmserver
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libdrmframeworkcommon
            liblog
            libm
            libmediametrics
            libmediautils
            libselinux
            libstagefright_foundation
            libutils
    )
endif ()
#################### drmserver ####################

#################### drmserver_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/drm/drmserver/fuzzer/DrmFuzzer.cpp
        ${ANDROID_ROOT}/frameworks/av/drm/drmserver/DrmManagerService.cpp
        ${ANDROID_ROOT}/frameworks/av/drm/drmserver/DrmManager.cpp
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
    add_executable(drmserver_fuzzer ${SOURCE_FILES})

    target_include_directories(drmserver_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/drm/drmserver
    )

    # Add c/c++ flags
    target_compile_options(drmserver_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(drmserver_fuzzer
            libbase
            libbinder
            libbinder_ndk
            libbinder_random_parcel
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_minimal
            libclang_rt.ubsan_standalone
            libcutils
            libdl
            libdrmframeworkcommon
            liblog
            libm
            libmediametrics
            libmediautils
            libselinux
            libstagefright_foundation
            libutils
    )
endif ()
#################### drmserver_fuzzer ####################
