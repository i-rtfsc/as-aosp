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

project(list_cpu_features)

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

#################### list_cpu_features ####################
set(FILES
        ${ANDROID_ROOT}/external/cpu_features/src/utils/list_cpu_features.c
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
    add_executable(list_cpu_features ${SOURCE_FILES})

    target_include_directories(list_cpu_features
            PUBLIC
            ${ANDROID_ROOT}/external/cpu_features/include
            ${ANDROID_ROOT}/external/cpu_features
    )

    # Add c/c++ flags
    target_compile_options(list_cpu_features PRIVATE
            -Werror
            -Wno-unused-function
    )

    target_link_libraries(list_cpu_features
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcpu_features
            libdl
            libm
    )
endif ()
#################### list_cpu_features ####################

#################### libcpu_features-stack_line_reader ####################
set(FILES
        ${ANDROID_ROOT}/external/cpu_features/src/stack_line_reader.c
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
    add_library(libcpu_features-stack_line_reader SHARED ${SOURCE_FILES})

    target_include_directories(libcpu_features-stack_line_reader
            PUBLIC
            ${ANDROID_ROOT}/external/cpu_features/include
            ${ANDROID_ROOT}/external/cpu_features
    )

    # Add c/c++ flags
    target_compile_options(libcpu_features-stack_line_reader PRIVATE
            -Werror
    )

    target_link_libraries(libcpu_features-stack_line_reader
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcpu_features-filesystem_for_testing
            libcpu_features-string_view
            libdl
            libgtest
            libgtest_main
            libm
    )
endif ()
#################### libcpu_features-stack_line_reader ####################

#################### libcpu_features-hwcaps ####################
set(FILES
        ${ANDROID_ROOT}/external/cpu_features/src/hwcaps.c
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
    add_library(libcpu_features-hwcaps SHARED ${SOURCE_FILES})

    target_include_directories(libcpu_features-hwcaps
            PUBLIC
            ${ANDROID_ROOT}/external/cpu_features/include
            ${ANDROID_ROOT}/external/cpu_features
    )

    # Add c/c++ flags
    target_compile_options(libcpu_features-hwcaps PRIVATE
            -Werror
    )

    target_link_libraries(libcpu_features-hwcaps
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcpu_features-utils
            libdl
            libm
    )
endif ()
#################### libcpu_features-hwcaps ####################

#################### libcpu_features-string_view ####################
set(FILES
        ${ANDROID_ROOT}/external/cpu_features/src/string_view.c
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
    add_library(libcpu_features-string_view SHARED ${SOURCE_FILES})

    target_include_directories(libcpu_features-string_view
            PUBLIC
            ${ANDROID_ROOT}/external/cpu_features/include
            ${ANDROID_ROOT}/external/cpu_features
    )

    # Add c/c++ flags
    target_compile_options(libcpu_features-string_view PRIVATE
            -Werror
    )

    target_link_libraries(libcpu_features-string_view
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libgtest
            libgtest_main
            libm
    )
endif ()
#################### libcpu_features-string_view ####################

#################### libcpu_features ####################
set(FILES
        ${ANDROID_ROOT}/external/cpu_features/src/impl_aarch64_linux_or_android.c
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
    add_library(libcpu_features SHARED ${SOURCE_FILES})

    target_include_directories(libcpu_features
            PUBLIC
            ${ANDROID_ROOT}/external/cpu_features/include
            ${ANDROID_ROOT}/external/cpu_features
    )

    # Add c/c++ flags
    target_compile_options(libcpu_features PRIVATE
            -Werror
    )

    target_link_libraries(libcpu_features
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcpu_features-hwcaps
            libcpu_features-utils
            libdl
            libm
            libwinpthread
    )
endif ()
#################### libcpu_features ####################

#################### libcpu_features-utils ####################
set(FILES
        ${ANDROID_ROOT}/external/cpu_features/src/filesystem.c
        ${ANDROID_ROOT}/external/cpu_features/src/stack_line_reader.c
        ${ANDROID_ROOT}/external/cpu_features/src/string_view.c
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
    add_library(libcpu_features-utils SHARED ${SOURCE_FILES})

    target_include_directories(libcpu_features-utils
            PUBLIC
            ${ANDROID_ROOT}/external/cpu_features/include
            ${ANDROID_ROOT}/external/cpu_features
    )

    # Add c/c++ flags
    target_compile_options(libcpu_features-utils PRIVATE
            -Werror
    )

    target_link_libraries(libcpu_features-utils
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libwinpthread
    )
endif ()
#################### libcpu_features-utils ####################
