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

project(libcxxbridge05)

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

#################### libcxxbridge05 ####################
set(FILES
        ${ANDROID_ROOT}/external/rust/cxx/src/cxx.cc
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
    add_library(libcxxbridge05 SHARED ${SOURCE_FILES})

    target_include_directories(libcxxbridge05
            PUBLIC
            ${ANDROID_ROOT}/external/rust/cxx
    )

    # Add c/c++ flags
    target_compile_options(libcxxbridge05 PRIVATE
            -Werror
    )

    target_link_libraries(libcxxbridge05
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libcxxbridge05 ####################

#################### libcxx-demo-blobstore ####################
set(FILES
        ${ANDROID_ROOT}/external/rust/cxx/demo/src/blobstore.cc
        ${OUT_INTERMEDIATES_ROOT}/external/rust/cxx/demo/cxx-demo-bridge-code/gen/demo/generated.cc
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
    add_library(libcxx-demo-blobstore SHARED ${SOURCE_FILES})

    target_include_directories(libcxx-demo-blobstore
            PUBLIC
            ${ANDROID_ROOT}/external/rust/cxx
            ${OUT_INTERMEDIATES_ROOT}/external/rust/cxx/demo/cxx-demo-bridge-code/gen
            ${OUT_INTERMEDIATES_ROOT}/external/rust/cxx/demo/cxx-demo-bridge-header/gen
            ${OUT_INTERMEDIATES_ROOT}/external/rust/cxx/cxx-bridge-header/gen
    )

    # Add c/c++ flags
    target_compile_options(libcxx-demo-blobstore PRIVATE
            -Werror
    )

    target_link_libraries(libcxx-demo-blobstore
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libcxx-demo-blobstore ####################
