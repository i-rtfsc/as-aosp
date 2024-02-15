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

project(drop_privs)

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

#################### libminijail_gen_syscall_llvmir ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/external/minijail/libminijail_gen_syscall/android_${OUT_ARCH_CPU}/gen/libsyscalls.c
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
    add_library(libminijail_gen_syscall_llvmir SHARED ${SOURCE_FILES})

    target_include_directories(libminijail_gen_syscall_llvmir
            PUBLIC
            ${ANDROID_ROOT}/external/minijail
            ${OUT_INTERMEDIATES_ROOT}/external/minijail/libminijail_gen_syscall/android_${OUT_ARCH_CPU}/gen
    )

    target_link_libraries(libminijail_gen_syscall_llvmir
            libc
            libc++
            libc++demangle
            libdl
            libm
    )
endif ()
#################### libminijail_gen_syscall_llvmir ####################

#################### libminijail_generated ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/external/minijail/libminijail_gen_syscall/android_${OUT_ARCH_CPU}/gen/libsyscalls.c
        ${OUT_INTERMEDIATES_ROOT}/external/minijail/libminijail_gen_constants/android_${OUT_ARCH_CPU}/gen/libconstants.c
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
    add_library(libminijail_generated SHARED ${SOURCE_FILES})

    target_include_directories(libminijail_generated
            PUBLIC
            ${ANDROID_ROOT}/external/minijail
            ${OUT_INTERMEDIATES_ROOT}/external/minijail/libminijail_gen_syscall/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/external/minijail/libminijail_gen_constants/android_${OUT_ARCH_CPU}/gen
    )

    # Add c/c++ flags
    target_compile_options(libminijail_generated PRIVATE
            -DALLOW_DEBUG_LOGGING
            -Wall
            -Werror
    )

    target_link_libraries(libminijail_generated
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libminijail_generated ####################

#################### libminijail_gen_constants_obj ####################
set(FILES
        ${ANDROID_ROOT}/external/minijail/gen_constants.c
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
    add_library(libminijail_gen_constants_obj SHARED ${SOURCE_FILES})

    target_include_directories(libminijail_gen_constants_obj
            PUBLIC
            ${ANDROID_ROOT}/external/minijail
    )

    # Add c/c++ flags
    target_compile_options(libminijail_gen_constants_obj PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libminijail_gen_constants_obj
            libc
            libc++
            libc++demangle
            libdl
            libm
    )
endif ()
#################### libminijail_gen_constants_obj ####################

#################### minijail0 ####################
set(FILES
        ${ANDROID_ROOT}/external/minijail/config_parser.c
        ${ANDROID_ROOT}/external/minijail/elfparse.c
        ${ANDROID_ROOT}/external/minijail/minijail0.c
        ${ANDROID_ROOT}/external/minijail/minijail0_cli.c
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
    add_executable(minijail0 ${SOURCE_FILES})

    target_include_directories(minijail0
            PUBLIC
            ${ANDROID_ROOT}/external/minijail
    )

    # Add c/c++ flags
    target_compile_options(minijail0 PRIVATE
            -DALLOW_DEBUG_LOGGING
            -Wall
            -Werror
    )

    target_link_libraries(minijail0
            libc
            libc++
            libc++demangle
            libcap
            libclang_rt.builtins
            libdl
            libm
            libminijail
            libminijail_generated
    )
endif ()
#################### minijail0 ####################

#################### libminijail ####################
set(FILES
        ${ANDROID_ROOT}/external/minijail/bpf.c
        ${ANDROID_ROOT}/external/minijail/landlock_util.c
        ${ANDROID_ROOT}/external/minijail/libminijail.c
        ${ANDROID_ROOT}/external/minijail/signal_handler.c
        ${ANDROID_ROOT}/external/minijail/syscall_filter.c
        ${ANDROID_ROOT}/external/minijail/syscall_wrapper.c
        ${ANDROID_ROOT}/external/minijail/system.c
        ${ANDROID_ROOT}/external/minijail/util.c
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
    add_library(libminijail SHARED ${SOURCE_FILES})

    target_include_directories(libminijail
            PUBLIC
            ${ANDROID_ROOT}/external/minijail
    )

    # Add c/c++ flags
    target_compile_options(libminijail PRIVATE
            -DALLOW_DEBUG_LOGGING
            -Wall
            -Werror
    )

    target_link_libraries(libminijail
            libc
            libc++
            libc++_static
            libc++demangle
            libcap
            libclang_rt.builtins
            libdl
            libm
            libminijail_generated
    )
endif ()
#################### libminijail ####################

#################### drop_privs ####################
set(FILES
        ${ANDROID_ROOT}/external/minijail/examples/drop_privs.cpp
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
    add_executable(drop_privs ${SOURCE_FILES})

    target_include_directories(drop_privs
            PUBLIC
            ${ANDROID_ROOT}/external/minijail
    )

    # Add c/c++ flags
    target_compile_options(drop_privs PRIVATE
            -DALLOW_DEBUG_LOGGING
            -Wall
            -Werror
    )

    target_link_libraries(drop_privs
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libminijail
    )
endif ()
#################### drop_privs ####################

#################### libminijail_gen_syscall_obj ####################
set(FILES
        ${ANDROID_ROOT}/external/minijail/gen_syscalls.c
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
    add_library(libminijail_gen_syscall_obj SHARED ${SOURCE_FILES})

    target_include_directories(libminijail_gen_syscall_obj
            PUBLIC
            ${ANDROID_ROOT}/external/minijail
    )

    # Add c/c++ flags
    target_compile_options(libminijail_gen_syscall_obj PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libminijail_gen_syscall_obj
            libc
            libc++
            libc++demangle
            libdl
            libm
    )
endif ()
#################### libminijail_gen_syscall_obj ####################

#################### libminijail_gen_constants_llvmir ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/external/minijail/libminijail_gen_constants/android_${OUT_ARCH_CPU}/gen/libconstants.c
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
    add_library(libminijail_gen_constants_llvmir SHARED ${SOURCE_FILES})

    target_include_directories(libminijail_gen_constants_llvmir
            PUBLIC
            ${ANDROID_ROOT}/external/minijail
            ${OUT_INTERMEDIATES_ROOT}/external/minijail/libminijail_gen_constants/android_${OUT_ARCH_CPU}/gen
    )

    target_link_libraries(libminijail_gen_constants_llvmir
            libc
            libc++
            libc++demangle
            libdl
            libm
    )
endif ()
#################### libminijail_gen_constants_llvmir ####################
