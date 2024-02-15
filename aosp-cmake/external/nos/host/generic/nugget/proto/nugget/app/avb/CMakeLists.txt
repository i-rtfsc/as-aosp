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

project(nos_app_avb)

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

#################### nos_app_avb ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/external/nos/host/generic/nugget/proto/nugget/app/avb/nos_app_avb_service_genc++/gen/Avb.client.cpp
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
    add_library(nos_app_avb SHARED ${SOURCE_FILES})

    target_include_directories(nos_app_avb
            PUBLIC
            ${ANDROID_ROOT}/external/nos/host/generic/nugget/proto/nugget/app/avb
            ${OUT_INTERMEDIATES_ROOT}/external/nos/host/generic/nugget/proto/nugget/app/avb/nos_app_avb_service_genc++/gen
            ${OUT_INTERMEDIATES_ROOT}/external/nos/host/generic/nugget/proto/nugget/app/avb/nos_app_avb_service_genc++_headers/gen
    )

    # Add c/c++ flags
    target_compile_options(nos_app_avb PRIVATE
            -Wall
            -Wextra
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(nos_app_avb
            libc++
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libdl
            libm
            libnos
            libnosprotos
            libprotobuf-cpp-full
            nos_headers
            libc
            libdl
            libm
    )
endif ()
#################### nos_app_avb ####################
