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

project(libselinux_android_setcontext_fuzzer)

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

#################### libselinux_selinux_check_access_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/external/selinux/libselinux/fuzzers/selinux_check_access_fuzzer.cpp
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
    add_executable(libselinux_selinux_check_access_fuzzer ${SOURCE_FILES})

    target_include_directories(libselinux_selinux_check_access_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/external/selinux/libselinux/fuzzers
    )

    # Add c/c++ flags
    target_compile_options(libselinux_selinux_check_access_fuzzer PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(libselinux_selinux_check_access_fuzzer
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            liblog
            libm
            libselinux
    )
endif ()
#################### libselinux_selinux_check_access_fuzzer ####################

#################### libselinux_selabel_lookup_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/external/selinux/libselinux/fuzzers/selabel_lookup_fuzzer.cpp
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
    add_executable(libselinux_selabel_lookup_fuzzer ${SOURCE_FILES})

    target_include_directories(libselinux_selabel_lookup_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/external/selinux/libselinux/fuzzers
    )

    # Add c/c++ flags
    target_compile_options(libselinux_selabel_lookup_fuzzer PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(libselinux_selabel_lookup_fuzzer
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            liblog
            libm
            libselinux
    )
endif ()
#################### libselinux_selabel_lookup_fuzzer ####################

#################### libselinux_setfilecon_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/external/selinux/libselinux/fuzzers/setfilecon_fuzzer.cpp
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
    add_executable(libselinux_setfilecon_fuzzer ${SOURCE_FILES})

    target_include_directories(libselinux_setfilecon_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/external/selinux/libselinux/fuzzers
    )

    # Add c/c++ flags
    target_compile_options(libselinux_setfilecon_fuzzer PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(libselinux_setfilecon_fuzzer
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            liblog
            libm
            libselinux
    )
endif ()
#################### libselinux_setfilecon_fuzzer ####################

#################### libselinux_lsetfilecon_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/external/selinux/libselinux/fuzzers/lsetfilecon_fuzzer.cpp
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
    add_executable(libselinux_lsetfilecon_fuzzer ${SOURCE_FILES})

    target_include_directories(libselinux_lsetfilecon_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/external/selinux/libselinux/fuzzers
    )

    # Add c/c++ flags
    target_compile_options(libselinux_lsetfilecon_fuzzer PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(libselinux_lsetfilecon_fuzzer
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            liblog
            libm
            libselinux
    )
endif ()
#################### libselinux_lsetfilecon_fuzzer ####################

#################### libselinux_context_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/external/selinux/libselinux/fuzzers/ContextFuzzer.cpp
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
    add_executable(libselinux_context_fuzzer ${SOURCE_FILES})

    target_include_directories(libselinux_context_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/external/selinux/libselinux/fuzzers
    )

    # Add c/c++ flags
    target_compile_options(libselinux_context_fuzzer PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(libselinux_context_fuzzer
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            liblog
            libm
            libselinux
    )
endif ()
#################### libselinux_context_fuzzer ####################

#################### libselinux_selinux_android_restorecon_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/external/selinux/libselinux/fuzzers/selinux_android_restorecon_fuzzer.cpp
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
    add_executable(libselinux_selinux_android_restorecon_fuzzer ${SOURCE_FILES})

    target_include_directories(libselinux_selinux_android_restorecon_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/external/selinux/libselinux/fuzzers
    )

    # Add c/c++ flags
    target_compile_options(libselinux_selinux_android_restorecon_fuzzer PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(libselinux_selinux_android_restorecon_fuzzer
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            liblog
            libm
            libselinux
    )
endif ()
#################### libselinux_selinux_android_restorecon_fuzzer ####################

#################### libselinux_string_to_security_class_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/external/selinux/libselinux/fuzzers/string_to_security_class_fuzzer.cpp
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
    add_executable(libselinux_string_to_security_class_fuzzer ${SOURCE_FILES})

    target_include_directories(libselinux_string_to_security_class_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/external/selinux/libselinux/fuzzers
    )

    # Add c/c++ flags
    target_compile_options(libselinux_string_to_security_class_fuzzer PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(libselinux_string_to_security_class_fuzzer
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            liblog
            libm
            libselinux
    )
endif ()
#################### libselinux_string_to_security_class_fuzzer ####################

#################### libselinux_selinux_android_setcon_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/external/selinux/libselinux/fuzzers/selinux_android_setcon_fuzzer.cpp
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
    add_executable(libselinux_selinux_android_setcon_fuzzer ${SOURCE_FILES})

    target_include_directories(libselinux_selinux_android_setcon_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/external/selinux/libselinux/fuzzers
    )

    # Add c/c++ flags
    target_compile_options(libselinux_selinux_android_setcon_fuzzer PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(libselinux_selinux_android_setcon_fuzzer
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            liblog
            libm
            libselinux
    )
endif ()
#################### libselinux_selinux_android_setcon_fuzzer ####################

#################### libselinux_android_setcontext_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/external/selinux/libselinux/fuzzers/AndroidSetcontextFuzzer.cpp
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
    add_executable(libselinux_android_setcontext_fuzzer ${SOURCE_FILES})

    target_include_directories(libselinux_android_setcontext_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/external/selinux/libselinux/fuzzers
    )

    # Add c/c++ flags
    target_compile_options(libselinux_android_setcontext_fuzzer PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(libselinux_android_setcontext_fuzzer
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            liblog
            libm
            libselinux
    )
endif ()
#################### libselinux_android_setcontext_fuzzer ####################
