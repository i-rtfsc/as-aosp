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

project(svc_dec_fuzzer)

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

#################### avc_dec_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/external/libavc/fuzzer/avc_dec_fuzzer.cpp
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
    add_executable(avc_dec_fuzzer ${SOURCE_FILES})

    target_include_directories(avc_dec_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/external/libavc/fuzzer
    )

    # Add c/c++ flags
    target_compile_options(avc_dec_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(avc_dec_fuzzer
            libavcdec
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            liblog
            libm
    )
endif ()
#################### avc_dec_fuzzer ####################

#################### mvc_dec_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/external/libavc/fuzzer/mvc_dec_fuzzer.cpp
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
    add_executable(mvc_dec_fuzzer ${SOURCE_FILES})

    target_include_directories(mvc_dec_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/external/libavc/fuzzer
    )

    # Add c/c++ flags
    target_compile_options(mvc_dec_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(mvc_dec_fuzzer
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            liblog
            libm
            libmvcdec
    )
endif ()
#################### mvc_dec_fuzzer ####################

#################### svc_enc_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/external/libavc/fuzzer/svc_enc_fuzzer.cpp
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
    add_executable(svc_enc_fuzzer ${SOURCE_FILES})

    target_include_directories(svc_enc_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/external/libavc/fuzzer
    )

    # Add c/c++ flags
    target_compile_options(svc_enc_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(svc_enc_fuzzer
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            liblog
            libm
            libsvcenc
    )
endif ()
#################### svc_enc_fuzzer ####################

#################### avc_enc_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/external/libavc/fuzzer/avc_enc_fuzzer.cpp
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
    add_executable(avc_enc_fuzzer ${SOURCE_FILES})

    target_include_directories(avc_enc_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/external/libavc/fuzzer
    )

    # Add c/c++ flags
    target_compile_options(avc_enc_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(avc_enc_fuzzer
            libavcenc
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            liblog
            libm
    )
endif ()
#################### avc_enc_fuzzer ####################

#################### svc_dec_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/external/libavc/fuzzer/svc_dec_fuzzer.cpp
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
    add_executable(svc_dec_fuzzer ${SOURCE_FILES})

    target_include_directories(svc_dec_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/external/libavc/fuzzer
    )

    # Add c/c++ flags
    target_compile_options(svc_dec_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(svc_dec_fuzzer
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            liblog
            libm
            libsvcdec
    )
endif ()
#################### svc_dec_fuzzer ####################
